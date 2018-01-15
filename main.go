package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/perigord/migration"
	"github.com/polyswarm/perigord/network"

	"github.com/polyswarm/polyswarm/bindings"
	"github.com/polyswarm/polyswarm/bounty"
	_ "github.com/polyswarm/polyswarm/migrations"
)

var bountyRegistry *bounty.BountyRegistry
var upgrader = websocket.Upgrader{}

func eventHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error upgrading websocket connection:", err)
		return
	}
	defer c.Close()

	eventChan := make(chan *bounty.Event)
	if err := bountyRegistry.WatchForEvents(eventChan); err != nil {
		log.Println("error listening for incoming events:", err)
		return
	}

	for {
		event := <-eventChan
		j, err := json.Marshal(event)
		if err != nil {
			log.Println("error marshalling evevnt:", err)
			continue
		}

		if err := c.WriteMessage(websocket.TextMessage, j); err != nil {
			log.Println("error writing response:", err)
			continue
		}
	}
}

func main() {
	network.InitNetworks()

	nw, err := network.Dial("dev")
	if err != nil {
		log.Fatalln("could not connect to dev network:", err)
	}

	if err := migration.RunMigrations(context.Background(), nw, false); err != nil {
		log.Fatalln("error running migrations: ", err)
	}

	bountyRegistrySession, ok := contract.Session("BountyRegistry").(*bindings.BountyRegistrySession)
	if bountyRegistrySession == nil || !ok {
		log.Fatalln("Invalid bounty registry session")
	}

	bountyRegistry = bounty.NewBountyRegistry(bountyRegistrySession, nw.Client(), os.Getenv("IPFS_HOST"))

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/events", eventHandler)

	log.Println("Listening on :31337")
	log.Fatal(http.ListenAndServe(":31337", r))
}
