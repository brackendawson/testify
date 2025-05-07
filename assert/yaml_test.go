package assert_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type nopTestingT struct{}

func (nopTestingT) Errorf(string, ...any) {}

func FuzzYAMLEq(f *testing.F) {
	f.Add(`---
key: value # as yaml
`, `{"key":"value"}`)
	f.Add(`1: one`, `2: two`)
	f.Fuzz(func(t *testing.T, expexted, actual string) {
		gotOld := assert.YAMLEqOld(nopTestingT{}, expexted, actual)
		gotNew := assert.YAMLEqNew(nopTestingT{}, expexted, actual)
		assert.Equal(t, gotOld, gotNew)
	})
}
