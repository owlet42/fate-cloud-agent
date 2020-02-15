package service

import (
	"io/ioutil"
	"strings"

	"fate-cloud-agent/pkg/db"

	"github.com/pkg/errors"

	"helm.sh/helm/v3/pkg/chart"
	"sigs.k8s.io/yaml"
)

// ReadFileToString read yaml file to string
func ReadFileToString(path string) (string, error) {
	dat, err := ioutil.ReadFile(path)
	return string(dat), err
}

// SaveChartFromPath read chart from path
func SaveChartFromPath(path string, name string) (*db.HelmChart, error) {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	// Get all template files
	templatePath := path + "templates/"
	files, err := ioutil.ReadDir(templatePath)
	if err != nil {
		return nil, err
	}
	templates := make([]*chart.File, len(files))
	for i, file := range files {
		modulePath := templatePath + file.Name()
		dataBytes, err := ioutil.ReadFile(modulePath)
		if err != nil {
			return nil, err
		}
		template := &chart.File{Name: file.Name(), Data: dataBytes}
		templates[i] = template
	}

	chartPath := path + "Chart.yaml"
	chartData, err := ioutil.ReadFile(chartPath)
	if err != nil {
		return nil, err
	}
	// get version from Chart.yaml
	metadata := new(chart.Metadata)
	if err := yaml.Unmarshal(chartData, metadata); err != nil {
		return nil, errors.Wrap(err, "cannot load Chart.yaml")
	}
	version := metadata.AppVersion

	valuePath := path + "values.yaml"
	valueString, err := ReadFileToString(valuePath)
	if err != nil {
		return nil, err
	}

	valueTemplatePath := path + "values-template.yaml"
	valueTemplateString, err := ReadFileToString(valueTemplatePath)
	if err != nil {
		return nil, err
	}

	helm := db.NewHelmChart(name, string(chartData), valueString, templates, version)
	helm.ValuesTemplate = valueTemplateString
	return helm, nil
}

// ConvertToChart convert database object to chart object
func ConvertToChart(helm *db.HelmChart) (*chart.Chart, error) {
	c := new(chart.Chart)

	templates := helm.Templates

	// Chart file
	chartData := []byte(helm.Chart)
	c.Raw = append(c.Raw, &chart.File{Name: "Chart.yaml", Data: chartData})
	if c.Metadata == nil {
		c.Metadata = new(chart.Metadata)
	}
	if err := yaml.Unmarshal(chartData, c.Metadata); err != nil {
		return c, errors.Wrap(err, "cannot load Chart.yaml")
	}
	if c.Metadata.APIVersion == "" {
		c.Metadata.APIVersion = chart.APIVersionV1
	}

	// Values file
	valuesData := []byte(helm.Values)
	c.Raw = append(c.Raw, &chart.File{Name: "values.yaml", Data: chartData})
	c.Values = make(map[string]interface{})
	if err := yaml.Unmarshal(valuesData, &c.Values); err != nil {
		return c, errors.Wrap(err, "cannot load values.yaml")
	}

	// Template files
	for _, template := range templates {
		c.Raw = append(c.Raw, template)
		c.Templates = append(c.Templates, &chart.File{Name: template.Name, Data: template.Data})
	}

	// TODO: Handling Chart.lock, values.schema.json, requirements.yaml, requirements.lock files

	return c, nil
}
