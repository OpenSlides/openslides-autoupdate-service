// Package environment implements helpers to handle environment varialbes and
// other function for startup.
package environment

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
	"golang.org/x/sys/unix"
)

// Environment variables used to configure the environment.
var (
	EnvDevelopment = NewVariable("OPENSLIDES_DEVELOPMENT", "false", "If set, the service uses the default secrets.")
)

// Variable represents a environment variable. It can be used by the packages
// for configuration.
//
// It is only allowed to use an environment variable at startup time.
type Variable struct {
	Key         string
	Default     string
	Description string
}

// NewVariable initializes a environment.Variable
func NewVariable(key, defaultValue, description string) Variable {
	return Variable{
		Key:         key,
		Default:     defaultValue,
		Description: description,
	}
}

// Value returns the value for an environment.Variable using a Getenver.
func (v Variable) Value(lookup Environmenter) string {
	lookup.UseVariable(v)

	if lookup == nil {
		return v.Default
	}

	val := lookup.Getenv(v.Key)
	if val == "" {
		return v.Default
	}
	return val
}

// Environmenter is an type, that can return the value for environment
// variables.
//
// It also saves the used environment variables.
type Environmenter interface {
	Getenv(key string) string
	UseVariable(v Variable)
}

// ForProduction is an environment used for production.
//
// It fetches the environment variables from os.Getenv().
type ForProduction struct{}

// Getenv calls os.Getenv.
func (e *ForProduction) Getenv(key string) string {
	return os.Getenv(key)
}

// UseVariable saves the used Variables
func (e *ForProduction) UseVariable(v Variable) {}

// ForDocu is an environment to gether the used environment variables.
//
// It does not access the real environment varialbes but returns an empty string
// so that the default value is used.
type ForDocu struct {
	usedVariables []Variable
}

// Getenv returns an empty string.
//
// But activates development mode to use the default values for secrets.
func (e *ForDocu) Getenv(key string) string {
	if key == EnvDevelopment.Key {
		return "true"
	}
	return ""
}

// UseVariable saves the used Variables
func (e *ForDocu) UseVariable(v Variable) {
	e.usedVariables = append(e.usedVariables, v)
}

// BuildDoc create the environment documentation with all used variables.
func (e *ForDocu) BuildDoc() (string, error) {
	var variables struct {
		Env []Variable
	}

	seen := set.New[string]()

	for _, v := range e.usedVariables {
		if seen.Has(v.Key) {
			continue
		}
		seen.Add(v.Key)

		variables.Env = append(variables.Env, v)
	}

	tmpl, err := template.New("Doc").Parse(tmplDoc)
	if err != nil {
		return "", fmt.Errorf("parsing template: %w", err)
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, variables); err != nil {
		return "", fmt.Errorf("executing template: %w", err)
	}

	return strings.ReplaceAll(buf.String(), "$", "`"), nil
}

// ReadSecret reads a secret from a file given by an environment variable.
//
// If OPENSLIDES_DEVELOPMENT is set, then this will always return the string
// 'openslides'
func ReadSecret(lookup Environmenter, pathVariable Variable) (string, error) {
	return ReadSecretWithDefault(lookup, pathVariable, "openslides")
}

// ReadSecretWithDefault is like ReadSecret, but it allows to set another
// default value then "openslides".
func ReadSecretWithDefault(lookup Environmenter, pathVariable Variable, defaultValue string) (string, error) {
	useDev, _ := strconv.ParseBool(EnvDevelopment.Value(lookup))
	path := pathVariable.Value(lookup)

	if useDev {
		return defaultValue, nil
	}

	secret, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read secret from %s: %w", path, err)
	}

	return string(secret), nil
}

// ForTests is a map that simulates environment variables.
type ForTests map[string]string

// Getenv returns the fake enironment variable for a key.
func (e ForTests) Getenv(key string) string {
	v := e[key]

	if key == EnvDevelopment.Key && v == "" {
		return "true"
	}

	return v
}

// UseVariable does nothing for tests.
func (e ForTests) UseVariable(v Variable) {}

// ParseDuration is like time.ParseDuration but uses second as default unit.
func ParseDuration(s string) (time.Duration, error) {
	sec, err := strconv.Atoi(s)
	if err == nil {
		return time.Duration(sec) * time.Second, nil
	}

	return time.ParseDuration(s)
}

// InterruptContext works like signal.NotifyContext. It returns a context that
// is canceled, when a signal is received.
//
// It listens on os.Interrupt and unix.SIGTERM. If the signal is received two
// times, os.Exit(2) is called.
func InterruptContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, unix.SIGTERM)
		<-sig
		cancel()
		<-sig
		os.Exit(2)
	}()
	return ctx, cancel
}

const tmplDoc = `<!--- Code generated with go generate ./... DO NOT EDIT. --->
# Configuration

## Environment Variables

The Service uses the following environment variables:
{{range .Env}}
* ${{.Key}}$: {{.Description}} The default is ${{.Default}}$.
{{- end}}`
