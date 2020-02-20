package cli

import (
	"errors"
	"fate-cloud-agent/pkg/db"
	"fmt"
	"github.com/gosuri/uitable"
	"helm.sh/helm/v3/pkg/cli/output"
	"os"
)

type Cluster struct {
}

func (c *Cluster) getRequestPath() (Path string) {
	return "cluster/"
}

type ClusterResultList struct {
	Data []*db.Cluster
	Msg  string
}

type ClusterResult struct {
	Data *db.Cluster
	Msg  string
}

type ClusterResultMsg struct {
	Msg string
}

type ClusterResultErr struct {
	Error string
}

func (c *Cluster) getResult(Type int) (result interface{}, err error) {
	switch Type {
	case LIST:
		result = new(ClusterResultList)
	case INFO:
		result = new(ClusterResult)
	case MSG:
		result = new(ClusterResultMsg)
	case ERROR:
		result = new(ClusterResultErr)
	default:
		err = fmt.Errorf("no type %d", Type)
	}
	return
}

func (c *Cluster) outPut(result interface{}, Type int) error {
	switch Type {
	case LIST:
		return c.outPutList(result)
	case INFO:
		return c.outPutInfo(result)
	case MSG:
		return c.outPutMsg(result)
	case ERROR:
		return c.outPutErr(result)
	default:
		return fmt.Errorf("no type %d", Type)
	}
	return nil
}

func (c *Cluster) outPutList(result interface{}) error {
	if result == nil {
		return errors.New("no out put data")
	}

	table := uitable.New()
	table.AddRow("UUID", "NAME", "NAMESPACE", "REVISION", "STATUS", "CHART", "APP VERSION")
	for _, r := range result.(*ClusterResultList).Data {
		table.AddRow(r.Uuid, r.Name, r.NameSpaces, r.Version, r.Status, r.ChartName, r.ChartVersion)
	}
	return output.EncodeTable(os.Stdout, table)
}

func (c *Cluster) outPutMsg(result interface{}) error {
	if result == nil {
		return errors.New("no out put data")
	}
	item, ok := result.(*ClusterResult)
	if !ok {
		return errors.New("not ok")
	}

	_, err := fmt.Fprintf(os.Stdout, "%s", item.Msg)

	return err
}

func (c *Cluster) outPutErr(result interface{}) error {
	if result == nil {
		return errors.New("no out put data")
	}
	item, ok := result.(*ClusterResultErr)
	if !ok {
		return errors.New("not ok")
	}

	_, err := fmt.Fprintf(os.Stdout, "%s", item.Error)

	return err
}

func (c *Cluster) outPutInfo(result interface{}) error {
	if result == nil {
		return errors.New("no out put data")
	}

	item, ok := result.(*ClusterResult)
	if !ok {
		return errors.New("not ok")
	}

	cluster := item.Data

	table := uitable.New()

	table.AddRow("UUID", cluster.Uuid)
	table.AddRow("Name", cluster.Name)
	table.AddRow("NameSpaces", cluster.NameSpaces)
	table.AddRow("Version", cluster.Version)
	table.AddRow("Type", cluster.Type)
	table.AddRow("Status", cluster.Status)
	table.AddRow("Values", cluster.Values)
	table.AddRow("ChartName", cluster.ChartName)
	table.AddRow("ChartVersion", cluster.ChartVersion)
	table.AddRow("Backend", cluster.Backend)
	table.AddRow("BootstrapParties", cluster.BootstrapParties)

	return output.EncodeTable(os.Stdout, table)
}
