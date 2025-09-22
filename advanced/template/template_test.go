package template

import (
	"os"
	"strings"
	"testing"
	"text/template"
)

func Test_template_parse(t *testing.T) {
	envs := make(map[string]string)
	for _, env := range os.Environ() {
		kv := strings.SplitN(env, "=", 2)
		envs[kv[0]] = kv[1]
	}

	var args = struct {
		Envs map[string]string
	}{
		Envs: envs,
	}

	tmpl := template.Must(template.New("test_template_parse").Parse(`{{ .Envs.USER}} Envs:
{{ with .Envs }}
	{{ range $key, $val := . }}
		{{ $key }}: {{ $val -}}
	{{ end }}
{{ end }}
`))

	_ = tmpl.ExecuteTemplate(os.Stdout, "test_template_parse", args)
}

func Test_if_empty_string(t *testing.T) {
	var args = struct {
		Tag        string
		Prerelease string
	}{
		Tag: "v1.5.0",
		//Prerelease: "beta1",
	}

	tmpl := template.Must(template.New("test_template_parse").Parse(`{{ if .Prerelease }}Pre-release {{ .Tag }}{{ else }}Release {{ .Tag }}{{ end }}`))

	_ = tmpl.ExecuteTemplate(os.Stdout, "test_template_parse", args)
}
