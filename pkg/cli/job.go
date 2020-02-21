package cli

import (
	"errors"
	"fate-cloud-agent/pkg/db"
	"fmt"
	"github.com/gosuri/uitable"
	"helm.sh/helm/v3/pkg/cli/output"
	"os"
)

type Job struct {
}

func (c *Job) getRequestPath() (Path string) {
	return "job/"
}

type JobResultList struct {
	Data []*db.Job
	Msg  string
}

type JobResult struct {
	Data *db.Job
	Msg  string
}

type JobResultMsg struct {
	Msg string
}

type JobResultErr struct {
	Error string
}

func (c *Job) getResult(Type int) (result interface{}, err error) {
	switch Type {
	case LIST:
		result = new(JobResultList)
	case INFO:
		result = new(JobResult)
	case MSG:
		result = new(JobResultMsg)
	case ERROR:
		result = new(JobResultErr)
	default:
		err = fmt.Errorf("no type %d", Type)
	}
	return
}

func (c *Job) outPut(result interface{}, Type int) error {
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

func (c *Job) outPutList(result interface{}) error {
	if result == nil {
		return errors.New("no out put data")
	}
	item, ok := result.(*JobResultList)
	if !ok {
		return errors.New("not ok")
	}

	table := uitable.New()
	table.AddRow("UUID", "CREATOR", "STARTTIME", "ENDTIME", "STATUS", "CLUSTERID", "RESULT")
	for _, r := range item.Data {
		table.AddRow(r.Uuid, r.Creator, r.StartTime.Format("2006-01-02 15:04:05"), r.EndTime.Format("2006-01-02 15:04:05"), r.Status, r.ClusterId, r.Result)
	}

	return output.EncodeTable(os.Stdout, table)
}

func (c *Job) outPutMsg(result interface{}) error {
	if result == nil {
		return errors.New("no out put data")
	}
	item, ok := result.(*JobResult)
	if !ok {
		return errors.New("not ok")
	}

	_, err := fmt.Fprintf(os.Stdout, "%s", item.Msg)

	return err
}

func (c *Job) outPutErr(result interface{}) error {
	if result == nil {
		return errors.New("no out put data")
	}
	item, ok := result.(*JobResultErr)
	if !ok {
		return errors.New("not ok")
	}

	_, err := fmt.Fprintf(os.Stdout, "%s", item.Error)

	return err
}

func (c *Job) outPutInfo(result interface{}) error {
	if result == nil {
		return errors.New("no out put data")
	}

	item, ok := result.(*JobResult)
	if !ok {
		return errors.New("not ok")
	}

	cluster := item.Data

	table := uitable.New()

	table.AddRow("UUID", cluster.Uuid)
	table.AddRow("StartTime", cluster.StartTime)
	table.AddRow("EndTime", cluster.EndTime)
	table.AddRow("Status", cluster.Status)
	table.AddRow("Creator", cluster.Creator)
	table.AddRow("ClusterId", cluster.ClusterId)
	table.AddRow("Result", cluster.Result)
	table.AddRow("SubJobs", cluster.SubJobs)

	return output.EncodeTable(os.Stdout, table)
}
