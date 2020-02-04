package db

type HelmChart struct {
	Name     string `json:"name"` 
	Value    string `json:"value"` 
	Template string `json:"template"` 
}

func NewHelmChart(name string, value string, template string) *HelmChart {
	helm := &HelmChart{
		Name: name,
		Value: value,
		Template: template,
	}

	return helm
}