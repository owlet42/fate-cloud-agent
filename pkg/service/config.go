package service

import (
	"bufio"
	"fmt"
	"github.com/Masterminds/sprig/v3"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

func Config() error {

	values, err := ioutil.ReadFile("E:/machenlong/AI/gitlab/fate-cloud-agent/values.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	config, err := ioutil.ReadFile("E:/machenlong/AI/gitlab/fate-cloud-agent/config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal(config, &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("values_1.2.0").Funcs(funcMap()).Option("missingkey=zero").Parse(string(values)))
	// Execute the template for each recipient.
	for _, r := range m["PartyList"].([]interface{}) {
		filename := fmt.Sprintf("E:/machenlong/AI/gitlab/fate-cloud-agent/%d-values.yaml", r.(map[interface{}]interface{})["PartyId"])
		file, err := os.Create(filename)
		writer := bufio.NewWriter(file)

		f := make(map[interface{}]interface{})
		for k, v := range m {
			f[k] = v
		}
		for k, v := range r.(map[interface{}]interface{}) {
			f[k] = v
		}
		var buf strings.Builder
		err = t.Execute(&buf, f)
		if err != nil {
			log.Println("executing template:", err)
		}
		s := strings.ReplaceAll(buf.String(), "<no value>", "")
		_, _ = writer.WriteString(s)
		err = writer.Flush()
		if err != nil {
			log.Println("executing template:", err)
		}
	}
	return nil
}

func funcMap() template.FuncMap {
	f := sprig.TxtFuncMap()
	delete(f, "env")
	delete(f, "expandenv")

	return f
}