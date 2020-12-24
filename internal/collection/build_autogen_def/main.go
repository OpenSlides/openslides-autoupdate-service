package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

const backendURL = "http://localhost:9002/health"

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() error {
	resp, err := http.Get(backendURL)
	if err != nil {
		return fmt.Errorf("requesting data from backend: %w", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("backend returned with status %s", resp.Status)
	}
	defer resp.Body.Close()

	var respData backendResponce

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return fmt.Errorf("decoding responce body: %w", err)
	}

	routes := make(map[string]string)
	for name, a := range respData.Info.Actions {
		if a.Perm.Type == "Generic permission check" {
			routes[name] = a.Perm.Perm
		}
	}

	if err := write(os.Stdout, routes); err != nil {
		return fmt.Errorf("write to stdout: %w", err)
	}

	return nil
}

type backendResponce struct {
	Info struct {
		Actions map[string]action `json:"actions"`
	} `json:"healthinfo"`
}

type action struct {
	Perm struct {
		Type string `json:"type"`
		Perm string `json:"permission"`
	} `json:"permission"`
}

const tpl = `// Code generated with autogen.gen DO NOT EDIT.
package collection

var autogenDef = map[string]string{
	{{- range $key, $value := .Def}}
	"{{$key}}": "{{$value}}",
	{{- end}}
}
`

func write(w io.Writer, data map[string]string) error {
	t := template.New("t")
	t, err := t.Parse(tpl)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	tdata := map[string]interface{}{
		"Def": data,
	}

	if err := t.Execute(w, tdata); err != nil {
		return fmt.Errorf("writing template: %w", err)
	}
	return nil
}
