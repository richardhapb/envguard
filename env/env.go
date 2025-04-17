package env

import (
	"envguard/parser"
	"errors"
	"maps"
	"strings"
)

type Env struct {
	Path	  string
	Variables map[string]string
}

func New(path string) *Env {
	return &Env{
		Path:	   path,
		Variables: make(map[string]string),
	}
}

// Parse the variables, store them in the env instance, and return the variables
func (e *Env) Parse() (map[string]string, error) {

	if e.Path == "" {
		return nil, errors.New("path is not provided")
	}
	file, err := parser.ReadFile(e.Path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scn := file.Scanner()
	for scn.Scan() {
		line := scn.Text()

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)

		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			e.Variables[key] = value
		}
	}

	return e.Variables, file.Err()
}

func (e *Env) String() string {
	result := ""
	for key, value := range e.Variables {
		result += key + ": " + value + "\n"
	}
	return result
}

func (e *Env) Compare(other *Env) ([]string, []string) {

	var e1_unique []string
	var e2_unique []string

	for variable := range maps.Keys(e.Variables) {
		if _, exists := other.Variables[variable]; !exists {
			e1_unique = append(e1_unique, variable)
		}
	}

	for variable := range maps.Keys(other.Variables) {
		if _, exists := e.Variables[variable]; !exists {
			e2_unique = append(e2_unique, variable)
		}
	}

	return e1_unique, e2_unique
}
