package main

import (
	"context"
	"encoding/json"
	"flag"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/perigord/migration"
	"github.com/polyswarm/perigord/network"

	"github.com/polyswarm/polyswarm/bindings"
	"github.com/polyswarm/polyswarm/bounty"
	_ "github.com/polyswarm/polyswarm/migrations"

	"github.com/satori/go.uuid"

	"github.com/tv42/base58"
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

func getArtifactsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ipfshash := vars["ipfshash"]

	// Verify that this is at least valid base58
	if _, err := base58.DecodeToBig([]byte(ipfshash)); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	links, err := bountyRegistry.ListArtifacts(ipfshash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	j, err := json.Marshal(links)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func getArtifactHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ipfshash := vars["ipfshash"]

	// Verify that this is at least valid base58
	if _, err := base58.DecodeToBig([]byte(ipfshash)); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil || id < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	links, err := bountyRegistry.ListArtifacts(ipfshash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if id >= len(links) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reader, err := bountyRegistry.DownloadArtifact(links[id])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=\"%v\"")
	io.Copy(w, reader)
}

func getArtifactStatHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ipfshash := vars["ipfshash"]

	// Verify that this is at least valid base58
	if _, err := base58.DecodeToBig([]byte(ipfshash)); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil || id < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	links, err := bountyRegistry.ListArtifacts(ipfshash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if id >= len(links) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stats, err := bountyRegistry.StatArtifact(ipfshash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	j, err := json.Marshal(stats)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func postArtifactsHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(1 << 12); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fs := make(map[string]io.Reader)
	for _, value := range r.MultipartForm.File {
		for _, fh := range value {
			f, err := fh.Open()
			defer f.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			fs[fh.Filename] = f
		}
	}

	uri, err := bountyRegistry.UploadArtifacts(fs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"uri": uri,
	}
	j, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func getBountiesHandler(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(bountyRegistry.GetBounties())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func postBountiesHandler(w http.ResponseWriter, r *http.Request) {
	var b struct {
		Amount      int    `json:"amount"`
		ArtifactURI string `json:"uri"`
		Duration    int    `json:"duration"`
	}

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	guid, err := bountyRegistry.PostBounty(context.Background(), b.ArtifactURI, b.Amount, b.Duration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"guid": uuid.FromBytesOrNil(guid.Bytes()),
	}
	j, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func getActiveBountiesHandler(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(bountyRegistry.GetActiveBounties())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func getBountyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	guid, err := uuid.FromString(vars["guid"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	guidInt := new(big.Int)
	guidInt.SetBytes(guid.Bytes())
	b := bountyRegistry.GetBountyByGuid(guidInt)
	if b == nil {
		http.Error(w, "invalid guid", http.StatusBadRequest)
		return
	}

	j, err := json.Marshal(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func getAssertionsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	guid, err := uuid.FromString(vars["guid"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	guidInt := new(big.Int)
	guidInt.SetBytes(guid.Bytes())
	j, err := json.Marshal(bountyRegistry.GetAssertionsByGuid(guidInt))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func postAssertionsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	guid, err := uuid.FromString(vars["guid"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	guidInt := new(big.Int)
	guidInt.SetBytes(guid.Bytes())

	var a struct {
		Bid      int    `json:"bid"`
		Mask     []bool `json:"mask"`
		Verdicts []bool `json:"verdicts"`
		Metadata string `json:"metadata"`
	}

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = bountyRegistry.PostAssertion(context.Background(), guidInt, a.Bid, a.Mask, a.Verdicts, a.Metadata)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{}
	j, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func getAssertionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	guid, err := uuid.FromString(vars["guid"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	guidInt := new(big.Int)
	guidInt.SetBytes(guid.Bytes())
	assertions := bountyRegistry.GetAssertionsByGuid(guidInt)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if id < 0 || id >= len(assertions) {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	j, err := json.Marshal(assertions[id])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

var networkFlag = flag.String("network", "dev", "which network to deploy to")

func main() {
	time.Sleep(15 * time.Second)

	flag.Parse()

	network.InitNetworks()

	nw, err := network.Dial(*networkFlag)
	if err != nil {
		log.Fatalln("could not connect to", *networkFlag, "network:", err)
	}

	for {
		if err := migration.RunMigrations(context.Background(), nw, false); err == nil {
			break
		}
		log.Println("Waiting on chain to sync...")
		time.Sleep(time.Second)
	}

	bountyRegistrySession, ok := contract.Session("BountyRegistry").(*bindings.BountyRegistrySession)
	if bountyRegistrySession == nil || !ok {
		log.Fatalln("Invalid bounty registry session")
	}

	bountyRegistry = bounty.NewBountyRegistry(bountyRegistrySession, nw.Client(), os.Getenv("IPFS_HOST"))

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/events", eventHandler)

	r.HandleFunc("/artifacts/{ipfshash}", getArtifactsHandler).Methods("GET")
	r.HandleFunc("/artifacts/{ipfshash}/{id}", getArtifactHandler).Methods("GET")
	r.HandleFunc("/artifacts/{ipfshash}/{id}/stat", getArtifactStatHandler).Methods("GET")
	r.HandleFunc("/artifacts", postArtifactsHandler).Methods("POST")

	r.HandleFunc("/bounties", getBountiesHandler).Methods("GET")
	r.HandleFunc("/bounties", postBountiesHandler).Methods("POST")
	r.HandleFunc("/bounties/active", getActiveBountiesHandler).Methods("GET")
	r.HandleFunc("/bounties/{guid}", getBountyHandler).Methods("GET")
	r.HandleFunc("/bounties/{guid}/assertions", getAssertionsHandler).Methods("GET")
	r.HandleFunc("/bounties/{guid}/assertions", postAssertionsHandler).Methods("POST")
	r.HandleFunc("/bounties/{guid}/assertions/{id}", getAssertionHandler).Methods("GET")

	log.Println("Listening on :31337")
	log.Fatal(http.ListenAndServe(":31337", Log(r)))
}
