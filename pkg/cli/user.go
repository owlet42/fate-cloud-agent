package cli

import (
	"errors"
	"fate-cloud-agent/pkg/db"
	"fmt"
	"github.com/gosuri/uitable"
	"helm.sh/helm/v3/pkg/cli/output"
	"os"
)

type User struct {
}

func (c *User) getRequestPath() (Path string) {
	return "user/"
}

func (c *User) addArgs() (Args string) {
	return ""
}

type UserResultList struct {
	Data []*db.User
	Msg  string
}

type UserResult struct {
	Data *db.User
	Msg  string
}

type UserResultMsg struct {
	Msg string
}

type UserResultErr struct {
	Error string
}

func (c *User) getResult(Type int) (result interface{}, err error) {
	switch Type {
	case LIST:
		result = new(UserResultList)
	case INFO:
		result = new(UserResult)
	case MSG, JOB:
		result = new(UserResultMsg)
	case ERROR:
		result = new(UserResultErr)
	default:
		err = fmt.Errorf("no type %d", Type)
	}
	return
}

func (c *User) outPut(result interface{}, Type int) error {
	switch Type {
	case LIST:
		return c.outPutList(result)
	case INFO:
		return c.outPutInfo(result)
	case MSG, JOB:
		return c.outPutMsg(result)
	case ERROR:
		return c.outPutErr(result)
	default:
		return fmt.Errorf("no type %d", Type)
	}
}

func (c *User) outPutList(result interface{}) error {
	if result == nil {
		return errors.New("no out put data")
	}
	item, ok := result.(*UserResultList)
	if !ok {
		return errors.New("type userResultList not ok")
	}

	table := uitable.New()
	table.AddRow("UUID", "USERNAME", "EMAIL", "STATUS")
	for _, r := range item.Data {
		table.AddRow(r.Uuid, r.Username, r.Email, r.Status)
	}
	return output.EncodeTable(os.Stdout, table)
}

func (c *User) outPutMsg(result interface{}) error {
	if result == nil {
		return errors.New("no out put data")
	}
	item, ok := result.(*UserResultMsg)
	if !ok {
		return errors.New("type UserResultMsg not ok")
	}

	_, err := fmt.Fprintf(os.Stdout, "%s", item.Msg)

	return err
}

func (c *User) outPutErr(result interface{}) error {
	if result == nil {
		return errors.New("no out put data")
	}
	item, ok := result.(*UserResultErr)
	if !ok {
		return errors.New("type userResultErr not ok")
	}

	_, err := fmt.Fprintf(os.Stdout, "%s", item.Error)

	return err
}

func (c *User) outPutInfo(result interface{}) error {
	if result == nil {
		return errors.New("no out put data")
	}

	item, ok := result.(*UserResult)
	if !ok {
		return errors.New("type UserResult not ok")
	}

	user := item.Data

	table := uitable.New()

	table.AddRow("UUID", user.Uuid)
	table.AddRow("StartTime", user.Username)
	table.AddRow("EndTime", user.Email)
	table.AddRow("Status", user.Status)

	return output.EncodeTable(os.Stdout, table)
}
