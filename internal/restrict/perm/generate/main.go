package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

const permURL = "https://raw.githubusercontent.com/OpenSlides/OpenSlides/openslides4-dev/docs/permission.yml"

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() error {
	f, err := loadPermissions()
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

func loadPermissions() (io.ReadCloser, error) {
	r, err := http.Get(permURL)
	if err != nil {
		return nil, fmt.Errorf("request defition: %w", err)
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request returned status %s", r.Status)
	}
	return r.Body, nil
}

type permFile map[string]permission

func (p permFile) derivate() map[string][]string {
	out := make(map[string][]string)

	for k, v := range map[string]permission(p) {
		v.derivate(out, k)
	}

	for k := range out {
		sort.Strings(out[k])
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

	derivate := data.derivate()

	consts := make(map[string]string, len(derivate))
	for k := range derivate {
		consts[constName(k)] = k
	}

	tdata := map[string]interface{}{
		"Derivate": derivate,
		"Consts":   consts,
	}

	if err := t.Execute(w, tdata); err != nil {
		return fmt.Errorf("writing template: %w", err)
	}
	return nil
}
