package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormatDate(t *testing.T) {
	created := "12th July 12:04"
	expected := time.Date(0, 7, 12, 12, 4, 0, 0, time.UTC)

	res, err := parseDate(created)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}
