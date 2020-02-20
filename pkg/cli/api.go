package cli

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	LIST int = iota
	INFO
	ERROR
	MSG
)

type Item interface {
	getRequestPath() (Path string)
	getResult(Type int) (result interface{}, err error)
	outPut(result interface{}, Type int) error
}

func postItem(i Item, Body []byte) error {
	req := &request{
		Type: "POST",
		Path: i.getRequestPath(),
		Body: Body,
	}

	rep, err := Send(req)
	if err != nil {
		return err
	}

	msg, err := i.getResult(INFO)

	err = json.Unmarshal(rep.Body, &msg)
	if err != nil {
		return err
	}

	err = i.outPut(msg, INFO)
	if err != nil {
		return err
	}
	return nil
}
func putItem(i Item, Body []byte) error {
	req := &request{
		Type: "PUT",
		Path: i.getRequestPath(),
		Body: Body,
	}

	rep, err := Send(req)
	if err != nil {
		return err
	}

	msg, err := i.getResult(INFO)

	err = json.Unmarshal(rep.Body, &msg)
	if err != nil {
		return err
	}

	err = i.outPut(msg, INFO)
	if err != nil {
		return err
	}
	return nil
}

func getItem(i Item, UUID string) error {
	req := &request{
		Type: "GET",
		Path: i.getRequestPath() + UUID,
		Body: nil,
	}

	rep, err := Send(req)
	if err != nil {
		return err
	}

	msg, err := i.getResult(INFO)

	err = json.Unmarshal(rep.Body, &msg)
	if err != nil {
		return err
	}

	err = i.outPut(msg, INFO)
	if err != nil {
		return err
	}
	return nil
}

func getItemList(i Item) error {
	req := &request{
		Type: "GET",
		Path: i.getRequestPath(),
		Body: nil,
	}

	rep, err := Send(req)
	if err != nil {
		return err
	}

	msg, err := i.getResult(LIST)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rep.Body, &msg)
	if err != nil {
		return err
	}

	err = i.outPut(msg, LIST)
	if err != nil {
		return err
	}
	return nil
}

func deleteItem(i Item, UUID string) error {
	req := &request{
		Type: "DELETE",
		Path: i.getRequestPath() + UUID,
		Body: nil,
	}

	rep, err := Send(req)
	if err != nil {
		return err
	}

	msg, err := i.getResult(MSG)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rep.Body, &msg)
	if err != nil {
		return err
	}

	err = i.outPut(msg, MSG)
	if err != nil {
		return err
	}
	return nil
}

func errOutPut(err error) {
	out := os.Stdout
	_, _ = fmt.Fprintf(out, "%s", err)
}
