package main

import (
	"testing"

	_ "github.com/polyswarm/polyswarm/tests"
	_ "github.com/polyswarm/polyswarm/migrations"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner
func Test(t *testing.T) { TestingT(t) }
