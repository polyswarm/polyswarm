package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/perigord/migration"
	"github.com/polyswarm/perigord/network"

	"github.com/polyswarm/polyswarm/bindings"
	"github.com/polyswarm/polyswarm/bountyregistry"
	_ "github.com/polyswarm/polyswarm/migrations"
)

var bountyRegistry *bountyregistry.BountyRegistry

func main() {
	network.InitNetworks()

	nw, err := network.Dial("dev")
	if err != nil {
		log.Fatalln("could not connect to dev network: ", err)
	}

	if err := migration.RunMigrations(context.Background(), nw); err != nil {
		log.Fatalln("error running migrations: ", err)
	}

	bountyRegistrySession, ok := contract.Session("BountyRegistry").(*bindings.BountyRegistrySession)
	if bountyRegistrySession == nil || !ok {
		log.Fatalln("Invalid bounty registry session")
	}

	bountyRegistry = bountyregistry.NewBountyRegistry(bountyRegistrySession, nw.Client())
	//keystore_path := nw.KeystorePath()

	r := mux.NewRouter().StrictSlash(true)
	//	r.HandleFunc("/bounties", BountyIndexHandler)
	//	r.HandleFunc("/bounties/{id}", BountyHandler)
	//	r.HandleFunc("/assertions", AssertionIndexHandler)
	//	r.HandleFunc("/assertions/{id}", AssertionHandler)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
