package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/perigord/migration"
	"github.com/polyswarm/perigord/network"

	"github.com/polyswarm/polyswarm/bindings"
	"github.com/polyswarm/polyswarm/bounty"
	_ "github.com/polyswarm/polyswarm/migrations"
)

var bountyRegistry *bounty.BountyRegistry

func main() {
	network.InitNetworks()

	nw, err := network.Dial("dev")
	if err != nil {
		log.Fatalln("could not connect to dev network: ", err)
	}

	if err := migration.RunMigrations(context.Background(), nw, false); err != nil {
		log.Fatalln("error running migrations: ", err)
	}

	bountyRegistrySession, ok := contract.Session("BountyRegistry").(*bindings.BountyRegistrySession)
	if bountyRegistrySession == nil || !ok {
		log.Fatalln("Invalid bounty registry session")
	}

	bountyRegistry = bounty.NewBountyRegistry(bountyRegistrySession, nw.Client(), os.Getenv("IPFS_HOST"))

	activeBounties := bountyRegistry.GetActiveBounties()
	for _, b := range activeBounties {
		j, err := json.Marshal(b)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(string(j))

		var nb bounty.Bounty
		err = json.Unmarshal(j, &nb)

		log.Println(nb)
	}

	//keystore_path := nw.KeystorePath()

	r := mux.NewRouter().StrictSlash(true)
	//	r.HandleFunc("/bounties", BountyIndexHandler)
	//	r.HandleFunc("/bounties/{id}", BountyHandler)
	//	r.HandleFunc("/assertions", AssertionIndexHandler)
	//	r.HandleFunc("/assertions/{id}", AssertionHandler)

	log.Println("Listening on :31337")
	log.Fatal(http.ListenAndServe(":31337", r))
}
