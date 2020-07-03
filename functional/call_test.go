package functional_test

import (
	"errors"
	"github.com/Ryanair/gofrlib/functional"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCall(t *testing.T) {

	var invoked bool
	fn1 := func() error { return errors.New("error") }
	fn2 := func() error { invoked = true; return nil }

	err := functional.Call(fn1, fn2)

	assert.Error(t, err)
	assert.False(t, invoked)

	err = functional.Call(fn2, fn1)
	assert.Error(t, err)
	assert.True(t, invoked)
}
