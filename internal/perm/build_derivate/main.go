package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() error {
	f, err := os.Open("../../permissions.yml")
	if err != nil {
		return fmt.Errorf("open permissions file: %w", err)
	}

	var d permFile
	if err := yaml.NewDecoder(f).Decode(&d); err != nil {
		return fmt.Errorf("decoding yaml: %w", err)
	}

	if err := write(os.Stdout, d); err != nil {
		return fmt.Errorf("writing data to stdout: %w", err)
	}
	return nil

}

type permFile map[string]permission

func (p permFile) derivate() map[string][]string {
	out := make(map[string][]string)

	for k, v := range map[string]permission(p) {
		v.derivate(out, k)
	}
	return out
}

type permission map[string]permission

func (p permission) derivate(data map[string][]string, collection string) {
	if len(p) == 0 {
		return
	}

	for k, v := range map[string]permission(p) {
		data[collection+"."+k] = v.subPerms(collection)
		v.derivate(data, collection)
	}
}

func (p permission) subPerms(collection string) []string {
	if len(p) == 0 {
		return nil
	}

	var out []string
	for k, v := range map[string]permission(p) {
		out = append(out, collection+"."+k)
		out = append(out, v.subPerms(collection)...)
	}
	return out
}

const tpl = `// Code generated with autogen.gen DO NOT EDIT.
package perm

var derivatePerms = map[string][]string{
	{{- range $key, $value := .Def}}
	"{{$key}}": { {{range $perm := $value}} "{{$perm}}", {{end}} },
	{{- end}}
}
`

func write(w io.Writer, data permFile) error {
	t := template.New("t")
	t, err := t.Parse(tpl)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	tdata := map[string]interface{}{
		"Def": data.derivate(),
	}

	if err := t.Execute(w, tdata); err != nil {
		return fmt.Errorf("writing template: %w", err)
	}
	return nil
}
