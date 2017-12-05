// Example main file for a native dapp, replace with application code
package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/perigord/migration"
	"github.com/polyswarm/perigord/network"

	"github.com/polyswarm/polyswarm/bindings"
	"github.com/polyswarm/polyswarm/bounty"
	_ "github.com/polyswarm/polyswarm/migrations"
)

var bountyPoster *bounty.BountyPoster
var templates = template.Must(template.ParseFiles(ReadDir("templates")...))

func ReadDir(dirname string) []string {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil
	}

	filenames := make([]string, len(files))
	for i, file := range files {
		filenames[i] = filepath.Join(dirname, file.Name())
	}

	return filenames
}
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		bounties := bountyPoster.GetActiveBounties()
		templates.ExecuteTemplate(w, "bounties", bounties)
	} else {
		http.Error(w, "Invalid request method", 405)
	}
}

func BountyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templates.ExecuteTemplate(w, "bounty", nil)
	} else if r.Method == "POST" {
		r.ParseMultipartForm(1 << 20)

		// TODO: Further validation of these parameters
		fee, err := strconv.Atoi(r.FormValue("fee"))
		if err != nil {
			log.Println("Error parsing int: ", err)
			http.Error(w, "Error parsing int", 500)
		}

		bountyAmount, err := strconv.Atoi(r.FormValue("bountyAmount"))
		if err != nil {
			log.Println("Error parsing int: ", err)
			http.Error(w, "Error parsing int", 500)
		}

		blocksFromNow, err := strconv.Atoi(r.FormValue("blocksFromNow"))
		if err != nil {
			log.Println("Error parsing int: ", err)
			http.Error(w, "Error parsing int", 500)
		}

		file, handler, err := r.FormFile("artifact")
		if err != nil {
			log.Println("Error retrieving artifact: ", err)
			http.Error(w, "Error retrieving artifact", 500)
		}
		defer file.Close()

		// TODO: Change this logic, IPFS?
		pth := "/tmp/" + handler.Filename
		newFile, err := os.OpenFile(pth, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Println("Error creating file: ", err)
			http.Error(w, "Error creating file", 500)
		}
		defer newFile.Close()

		io.Copy(newFile, file)
		newBounty, err := bounty.NewBounty(pth, fee, bountyAmount, blocksFromNow)
		if err != nil {
			log.Println("Error creating bounty: ", err)
			http.Error(w, "Error creating bounty", 500)
		}

		newBounty.Upload()

		_, err = bountyPoster.PostBounty(context.Background(), newBounty)
		if err != nil {
			log.Println("Error posting bounty: ", err)
			http.Error(w, "Error posting bounty", 500)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	} else {
		http.Error(w, "Invalid request method", 405)
	}
}

func AssertionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templates.ExecuteTemplate(w, "assertion", nil)
	} else if r.Method == "POST" {
		r.ParseForm()

		// TODO: Further validation of these parameters
		malicious := r.FormValue("malicious") == "true"

		bid, err := strconv.Atoi(r.FormValue("bid"))
		if err != nil {
			log.Println("Error parsing int: ", err)
			http.Error(w, "Error parsing int", 500)
		}

		metadata := r.FormValue("metadata")

		guid := big.NewInt(0)
		if _, ok := guid.SetString(r.FormValue("guid"), 16); !ok {
			log.Println("Error parsing int: ", err)
			http.Error(w, "Error parsing int", 500)
		}

		bnty := bountyPoster.GetBountyByGuid(guid)
		if bnty == nil {
			log.Println("Error finding associated bounty: ", err)
			http.Error(w, "Error finding associated bounty", 500)
		}

		newAssertion, err := bounty.NewAssertion(malicious, bid, metadata)
		if err != nil {
			log.Println("Error creating assertion: ", err)
			http.Error(w, "Error creating assertion", 500)
		}

		newAssertion.SetGuid(guid)

		_, err = bountyPoster.PostAssertion(context.Background(), bnty, newAssertion)
		if err != nil {
			log.Println("Error posting assertion: ", err)
			http.Error(w, "Error posting assertion", 500)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	} else {
		http.Error(w, "Invalid request method", 405)
	}
}

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

	bountyPoster = bounty.NewBountyPoster(bountyRegistrySession, nw.Client())

	keystore_path := nw.KeystorePath()
	go func() {
		// Hardcode some of these parameters for now, these come from deploying
		// the raiden contracts onto our local test network
		cmd := exec.Command("raiden", "--eth-rpc-endpoint", "172.17.0.1", "--keystore-path", keystore_path,
			"--password-file", "scripts/password",
			"--address", "c3e33c2bb955d691f085544235a7e2698fbca900",
			"--discovery-contract-address", "e1562f18ff1278edbbf1201624512f9c4dceaac7",
			"--registry-contract-address", "191343ea2e338b7a061336fd224e83a83386439b",
		)
		cmd.Run()
	}()

	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/bounty", BountyHandler)
	r.HandleFunc("/assertion", AssertionHandler)
	http.Handle("/", r)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
