package taskfile

import (
	"testing"

	"github.com/go-task/task/pkg"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestCmdParse(t *testing.T) {
	const (
		yamlCmd      = `echo "a string command"`
		yamlDep      = `"task-name"`
		yamlTaskCall = `
task: another-task
vars:
  PARAM1: VALUE1
  PARAM2: VALUE2
`
	)
	tests := []struct {
		content  string
		v        interface{}
		expected interface{}
	}{
		{
			yamlCmd,
			&pkg.Cmd{},
			&pkg.Cmd{Cmd: `echo "a string command"`},
		},
		{
			yamlTaskCall,
			&pkg.Cmd{},
			&pkg.Cmd{Task: "another-task", Vars: pkg.Vars{
				"PARAM1": pkg.Var{Static: "VALUE1"},
				"PARAM2": pkg.Var{Static: "VALUE2"},
			}},
		},
		{
			yamlDep,
			&pkg.Dep{},
			&pkg.Dep{Task: "task-name"},
		},
		{
			yamlTaskCall,
			&pkg.Dep{},
			&pkg.Dep{Task: "another-task", Vars: pkg.Vars{
				"PARAM1": pkg.Var{Static: "VALUE1"},
				"PARAM2": pkg.Var{Static: "VALUE2"},
			}},
		},
	}
	for _, test := range tests {
		err := yaml.Unmarshal([]byte(test.content), test.v)
		assert.NoError(t, err)
		assert.Equal(t, test.expected, test.v)
	}
}
