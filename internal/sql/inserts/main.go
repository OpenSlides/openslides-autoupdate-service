package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/sql"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
)

func main() {
	if err := run(os.Stdout, os.Stdin); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run(out io.Writer, in io.Reader) error {
	yamlData, err := io.ReadAll(in)
	if err != nil {
		return fmt.Errorf("reading input: %w", err)
	}

	decoded := dsmock.YAMLData(string(yamlData))
	statements := sql.Insert(decoded)
	if _, err := out.Write([]byte(statements)); err != nil {
		return fmt.Errorf("writing statements: %w", err)
	}
	fmt.Fprintln(out)

	return nil
}
