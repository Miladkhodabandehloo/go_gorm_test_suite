package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
)



func TestSome(t *testing.T) {
	s := new(ExampleSuite)
	suite.Run(t, s)
}
