package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProfessionalExist(t *testing.T) {
	p := &Professional{}
	assert.NotNil(t, p)
}

func TestCreateProfessional(t *testing.T) {
	p := GetProfessional("Jon10")
	assert.NotNil(t, p)
}

func TestGetSetAdjustScore(t *testing.T) {
	p := GetProfessional("Jon104")
	assert.NotNil(t, p)

	p.SetScore(110)
	assert.InEpsilon(t, 100, p.GetScore(), epsilon)
	p.AdjustScore(-0.5)
	assert.InEpsilon(t, 50, p.GetScore(), epsilon)
	p.SetScore(-100)
	assert.Equal(t, p.GetScore() < epsilon, true)
}
