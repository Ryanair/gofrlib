package errorUtils_test

import (
	"errors"
	"github.com/Ryanair/gofrlib/errorUtils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldJoinErrors(t *testing.T) {
	errs := []error{errors.New("err1"), errors.New("err2"), errors.New("err3")}

	mergedErr := errorUtils.MergeErrors(errs)

	assert.Equal(t, errors.New("err1\nerr2\nerr3"), mergedErr)
}

func Test_ShouldSkipNilsInErrorSlice(t *testing.T) {
	errs := []error{errors.New("err1"), nil, errors.New("err2"), nil}

	mergedErr := errorUtils.MergeErrors(errs)

	assert.Equal(t, errors.New("err1\nerr2"), mergedErr)
}

func Test_ShouldReturnNilWhenOnlyNilsInErrorSlice(t *testing.T) {
	errs := []error{nil, nil}

	mergedErr := errorUtils.MergeErrors(errs)

	assert.Nil(t, mergedErr)
}

func Test_ShouldReturnNilWhenNoErrors(t *testing.T) {
	var errs []error

	mergedErr := errorUtils.MergeErrors(errs)

	assert.Nil(t, mergedErr)
}
