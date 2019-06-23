package util

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGoroutineCount(t *testing.T) {

	if runtime.NumCPU() == 1 {
		assert.Equal(t, 1, getGoroutineCount())

	} else {
		assert.Equal(t, runtime.NumCPU()-1, getGoroutineCount())

	}

}
