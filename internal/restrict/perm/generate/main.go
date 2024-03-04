package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/goccy/go-yaml"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() error {
	f, err := openPermissionYML()
	if err != nil {
		return fmt.Errorf("open permissions file: %w", err)
	}
	defer f.Close()

	var d permFile
	if err := yaml.NewDecoder(f).Decode(&d); err != nil {
		return fmt.Errorf("decoding yaml: %w", err)
	}

	if err := write(os.Stdout, d); err != nil {
		return fmt.Errorf("writing data to stdout: %w", err)
	}
	return nil
}

func openPermissionYML() (io.ReadCloser, error) {
	return os.Open("../../../meta/permission.yml")
}

type permFile map[string]permission

func (p permFile) derivative() map[string][]string {
	out := make(map[string][]string)

	for k, v := range map[string]permission(p) {
		v.derivative(out, k)
	}

	for k := range out {
		sort.Strings(out[k])
	}
	return out
}

type permission map[string]permission

func (p permission) derivative(data map[string][]string, collection string) {
	if len(p) == 0 {
		return
	}

	for k, v := range map[string]permission(p) {
		data[collection+"."+k] = v.subPerms(collection)
		v.derivative(data, collection)
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

func constName(perm string) string {
	bs := []byte(strings.ReplaceAll(perm, ".", "_"))

	bs[0] -= 'a' - 'A'

	for i := 1; i < len(bs); i++ {
		if bs[i-1] == '_' {
			bs[i] -= 'a' - 'A'
		}
	}

	bs = bytes.ReplaceAll(bs, []byte("_"), []byte(""))
	return string(bs)
}

const tpl = `// Code generated from models.yml DO NOT EDIT.
package perm

const (
	{{- range $key, $value := .Consts}}
	{{$key}} TPermission = "{{$value}}"
	{{- end}}
)

var derivatePerms = map[TPermission][]TPermission{
	{{- range $key, $value := .Derivate}}
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

	derivate := data.derivative()

	consts := make(map[string]string, len(derivate))
	for k := range derivate {
		consts[constName(k)] = k
	}

	tdata := map[string]interface{}{
		"Derivate": derivate,
		"Consts":   consts,
	}

	buf := new(bytes.Buffer)

	if err := t.Execute(buf, tdata); err != nil {
		return fmt.Errorf("writing template: %w", err)
	}

	formated, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("formating code: %w", err)
	}

	if _, err := w.Write(formated); err != nil {
		return fmt.Errorf("writing output: %w", err)
	}

	return nil
}
