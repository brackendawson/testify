package assert_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/assert/yaml"
)

type nopTestingT struct{}

func (nopTestingT) Errorf(string, ...any) {}

func FuzzYAMLEq(f *testing.F) {
	f.Add(`---
key: value # as yaml
`, `{"key":"value"}`)
	f.Add(`1: one`, `2: two`)
	f.Fuzz(func(t *testing.T, expexted, actual string) {
		var v any
		if err := yaml.UnmarshalOld([]byte(expexted), &v); err != nil {
			t.Skip()
		}
		if err := yaml.UnmarshalOld([]byte(actual), &v); err != nil {
			t.Skip()
		}
		if err := yaml.UnmarshalNew([]byte(expexted), &v); err != nil {
			t.Skip()
		}
		if err := yaml.UnmarshalNew([]byte(actual), &v); err != nil {
			t.Skip()
		}

		gotOld := assert.YAMLEqOld(nopTestingT{}, expexted, actual)
		gotNew := assert.YAMLEqNew(nopTestingT{}, expexted, actual)
		assert.Equal(t, gotOld, gotNew)
	})
}
