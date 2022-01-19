package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSampleFunc (t *testing.T) {
  actual := sampleFunc()
  expected := "Hello Universe. You are infinite..!"
  assert.Equal(t, actual, expected)  
}
