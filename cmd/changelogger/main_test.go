package main

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func fixtureFile(t *testing.T, contents []byte) string {
	t.Helper()
	f, err := os.CreateTemp("", "*")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { f.Close() })

	if _, err := f.Write(contents); err != nil {
		t.Fatal(err)
	}
	return f.Name()
}

type errorWriter struct{}

func (ew *errorWriter) Write(p []byte) (n int, err error) { return 0, errors.New("OOPS") }

func Test_run(t *testing.T) {
	markdown110, err := os.ReadFile(filepath.Join("testdata", "want", "1.1.0.md"))
	if err != nil {
		t.Fatal(err)
	}

	markdown090, err := os.ReadFile(filepath.Join("testdata", "want", "0.9.0.md"))
	if err != nil {
		t.Fatal(err)
	}

	markdownSimple, err := os.ReadFile(filepath.Join("testdata", "want", "simple.md"))
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		cfg config
	}
	tests := []struct {
		name      string
		args      args
		configure func() func()
		err       string
		want      string
	}{
		//  Empty config
		{
			name: "empty config returns the multierror validating the required options",
			args: args{cfg: config{}},
			err: `invalid flags: 4 errors occurred:
	* base-url cannot be empty
	* out io.writer cannot be nil
	* template cannot be empty
	* version cannot be empty

`,
		},

		// Failure paths.
		{
			name: "empty changelog writes nothing and returns an error",
			args: args{cfg: config{
				out:      new(bytes.Buffer),
				dir:      "testdata",
				version:  "empty",
				template: fixtureFile(t, []byte("")),
				baseURL:  "https://something",
			}},
			err: "folder testdata/empty has no changelog files",
		},
		{
			name: "attempts to parse an invalid template file",
			args: args{cfg: config{
				out:      new(bytes.Buffer),
				dir:      "testdata",
				version:  "0.9.0",
				template: fixtureFile(t, []byte("assadad {{")),
				baseURL:  "https://something",
			}},
			err: "failed parsing template file contents: template: changelog:1: unclosed action",
		},
		{
			name: "attempts to parse a non existent template file",
			args: args{cfg: config{
				out:      new(bytes.Buffer),
				dir:      "testdata",
				version:  "0.9.0",
				template: "where?",
				baseURL:  "https://something",
			}},
			err: "failed opening template file: open where?: no such file or directory",
		},
		{
			name: "attempts to walk an invalid path",
			args: args{cfg: config{
				out:      new(bytes.Buffer),
				dir:      "a",
				version:  "ww",
				template: filepath.Join("testdata", "templates", "markdown.gtpl"),
				baseURL:  "https://something",
			}},
			err: "failed walking the specified path: lstat a/ww: no such file or directory",
		},
		{
			name: "fails writing to the output io.Writer",
			args: args{cfg: config{
				out:      &errorWriter{},
				dir:      "testdata",
				version:  "v1.1.0",
				template: filepath.Join("testdata", "templates", "markdown.gtpl"),
				baseURL:  "https://something",
			}},
			err: "failed copying the template output: OOPS",
		},
		{
			name: "changelog 1.1.0 fails executing the template when it's invalid",
			args: args{cfg: config{
				out:      new(bytes.Buffer),
				dir:      "testdata",
				version:  "v1.1.0",
				template: filepath.Join("testdata", "templates", "failure.gtpl"),
				baseURL:  "https://something",
			}},
			err: "failed executing the changelog template: template: changelog:1:2: executing \"changelog\" at <.AAAAA>: can't evaluate field AAAAA in type changelogger.Changes",
		},
		{
			name: "changelog invalid fails parsing changelog config files",
			args: args{cfg: config{
				out:      new(bytes.Buffer),
				dir:      "testdata",
				version:  "invalid",
				template: filepath.Join("testdata", "templates", "simple.gtpl"),
				baseURL:  "https://something",
			}},
			err: "failed walking the specified path: failed decoding yaml file testdata/invalid/a.yml: error unmarshaling JSON: json: cannot unmarshal string into Go value of type changelogger.Change",
		},

		// Success cases.
		{
			name: "changelog 1.1.0 has a few changes",
			args: args{cfg: config{
				out:      new(bytes.Buffer),
				dir:      "testdata",
				version:  "v1.1.0",
				template: filepath.Join("testdata", "templates", "markdown.gtpl"),
				baseURL:  "https://something",
			}},
			err:  "folder testdata/empty has no changelog files",
			want: string(markdown110),
		},
		{
			name: "changelog 0.9.0 has a single change",
			configure: func() func() {
				old := os.Getenv("ECE_VERSION")
				os.Setenv("ECE_VERSION", "SOME_VERSION")
				return func() { os.Setenv("ECE_VERSION", old) }
			},
			args: args{cfg: config{
				out:      new(bytes.Buffer),
				dir:      "testdata",
				version:  "v0.9.0",
				template: filepath.Join("testdata", "templates", "markdown.gtpl"),
				baseURL:  "https://something/else",
			}},
			err:  "folder testdata/empty has no changelog files",
			want: string(markdown090),
		},
		{
			name: "uses a very simple template",
			args: args{cfg: config{
				out:      new(bytes.Buffer),
				dir:      "testdata",
				version:  "v0.9.0",
				template: filepath.Join("testdata", "templates", "simple.gtpl"),
				baseURL:  "https://link/to",
			}},
			err:  "folder testdata/empty has no changelog files",
			want: string(markdownSimple),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.configure != nil {
				t.Cleanup(tt.configure())
			}

			if err := run(tt.args.cfg); err != nil {
				assert.EqualError(t, err, tt.err)
			}

			if buf, ok := tt.args.cfg.out.(*bytes.Buffer); ok {
				assert.Equal(t, tt.want, buf.String())
			}
		})
	}
}
