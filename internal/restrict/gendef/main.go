// This tool generates the list of related fields in the file def.go.
// To call it, just call "go generate ./..." in the root folder of the repository
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"
)

const defURL = "https://raw.githubusercontent.com/OpenSlides/OpenSlides/openslides4-dev/docs/models.txt"

const pkgName = "restrict"

func main() {
	r, err := loadDefition()
	if err != nil {
		log.Fatalf("Can not load defition: %v", err)
	}
	defer r.Close()

	data, err := parse(r)
	if err != nil {
		log.Fatalf("Can not parse model definition: %v", err)
	}

	if err := writeFile(os.Stdout, data); err != nil {
		log.Fatalf("Can not write result: %v", err)
	}
}

func loadDefition() (io.ReadCloser, error) {
	r, err := http.Get(defURL)
	if err != nil {
		return nil, fmt.Errorf("request defition: %w", err)
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request returned status %s", r.Status)
	}
	return r.Body, nil
}

func parse(r io.Reader) (map[string]string, error) {
	reInterface := regexp.MustCompile("^Interface ([a-z_]+) {$")
	reRelationField := regexp.MustCompile(`^\s*([a-z0-9_]+(?:\$<[^>]+>)?(?:[a-z0-9_]+)?): \((\*|[a-z_]+)/[^)]*\)\[\];\s*(?://.*)?$`)
	reTemplateID := regexp.MustCompile(`<[^>]+>`)

	s := bufio.NewScanner(r)
	var currentModel string
	out := make(map[string]string)
	for s.Scan() {
		// Look for current Model.
		m := reInterface.FindStringSubmatch(s.Text())
		if len(m) == 2 {
			currentModel = m[1]
			continue
		}

		// If there was no current model yet, continue.
		if currentModel == "" {
			continue
		}

		m = reRelationField.FindStringSubmatch(s.Text())
		if len(m) == 3 {
			key := fmt.Sprintf("%s/%s", currentModel, m[1])
			key = reTemplateID.ReplaceAllString(key, "")
			out[key] = m[2]
		}

	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("scanning: %w", err)
	}
	return out, nil
}

const tpl = `// Code generated with models.txt DO NOT EDIT.
package {{.PackageName}}

var relationLists = map[string]string{
	{{- range $key, $value := .Def}}
	"{{$key}}": "{{$value}}",
	{{- end}}
}
`

func writeFile(w io.Writer, rlist map[string]string) error {
	t := template.New("t")
	t, err := t.Parse(tpl)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	data := map[string]interface{}{
		"PackageName": pkgName,
		"Def":         rlist,
	}

	if err := t.Execute(w, data); err != nil {
		return fmt.Errorf("writing template: %w", err)
	}
	return nil
}
