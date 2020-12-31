package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/blackcrw/wprecon-api/pkg/database"
	"github.com/gorilla/mux"
)

// GetPluginVuln ::
func GetPluginVuln(writer http.ResponseWriter, response *http.Request) {
	parameters := mux.Vars(response)

	vuln, err := database.GetVulnerabilities(parameters["Plugin"], parameters["Version"])

	if err != nil {
		log.Fatal(err)
	}

	var jsonData []byte

	jsonData, err = json.Marshal(vuln)

	if err != nil {
		log.Println(err)
	}

	writer.Write([]byte(jsonData))
}

// GetPluginNameList ::
func GetPluginNameList(writer http.ResponseWriter, response *http.Request) {
	writer.Write([]byte("Plugin List Name"))
}
