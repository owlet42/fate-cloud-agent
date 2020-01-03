package pkg

import (
	"bytes"
	"fmt"
	"github.com/gosuri/uitable"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli/output"
	"helm.sh/helm/v3/pkg/release"
	"io"
	"os"
	"strconv"
)

func List(namespaces string) (*releaseListWriter, error) {
	cfg := new(action.Configuration)
	client := action.NewList(cfg)
	settings.EnvVars()
	if err := cfg.Init(settings.RESTClientGetter(), Namespace(namespaces), os.Getenv("HELM_DRIVER"), debug); err != nil {
		debug("%+v", err)
		os.Exit(1)
	}

	client.SetStateMask()

	results, err := client.Run()
	if err != nil {
		return nil, err
	}

	res := newReleaseListWriter(results)
	s, _ := res.WriteToJSON()
	fmt.Println(s)
	return res, nil
}
func Namespace(namespace string) string {

	if namespace == "allnamespaces" {
		return ""
	}
	if namespace != "" {
		return namespace
	}
	if ns, _, err := settings.RESTClientGetter().ToRawKubeConfigLoader().Namespace(); err == nil {
		return ns
	}
	return "default"
}

type releaseElement struct {
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
	Revision   string `json:"revision"`
	Updated    string `json:"updated"`
	Status     string `json:"status"`
	Chart      string `json:"chart"`
	AppVersion string `json:"app_version"`
}
type releaseListWriter struct {
	Releases []releaseElement
}

func newReleaseListWriter(releases []*release.Release) *releaseListWriter {
	// Initialize the array so no results returns an empty array instead of null
	elements := make([]releaseElement, 0, len(releases))
	for _, r := range releases {
		element := releaseElement{
			Name:       r.Name,
			Namespace:  r.Namespace,
			Revision:   strconv.Itoa(r.Version),
			Status:     r.Info.Status.String(),
			Chart:      fmt.Sprintf("%s-%s", r.Chart.Metadata.Name, r.Chart.Metadata.Version),
			AppVersion: r.Chart.Metadata.AppVersion,
		}
		t := "-"
		if tspb := r.Info.LastDeployed; !tspb.IsZero() {
			t = tspb.String()
		}
		element.Updated = t
		elements = append(elements, element)
	}
	return &releaseListWriter{elements}
}

func (r *releaseListWriter) WriteTable(out io.Writer) error {
	table := uitable.New()
	table.AddRow("NAME", "NAMESPACE", "REVISION", "UPDATED", "STATUS", "CHART", "APP VERSION")
	for _, r := range r.Releases {
		table.AddRow(r.Name, r.Namespace, r.Revision, r.Updated, r.Status, r.Chart, r.AppVersion)
	}
	return output.EncodeTable(out, table)
}


func (r *releaseListWriter) WriteToJSON() (s string, err error) {
	buf := new(bytes.Buffer)
	err = output.EncodeJSON(buf, r.Releases)
	s = buf.String()
	return s, err
}

