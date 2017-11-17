// Invokes the perigord driver application

package main

import (
	_ "github.com/polyswarm/polyswarm/migrations"
	"github.com/polyswarm/perigord/stub"
)

func main() {
	stub.StubMain()
}
