package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingConfig(t *testing.T) {
	assert := assert.New(t)

	var p Plugin
	err := p.Exec()

	assert.Error(err, "Missing circle-ci configs")
}
