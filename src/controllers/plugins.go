package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/blackcrw/wprecon-api/src/database"
	"github.com/gorilla/mux"
)

/*
GetPluginVuln :: This function will take the mux variables and will pass them to the database, to make a request and return the information in the api browser.
(In short: It will look for vulnerability in the plugin and version informed in the database and will print them on the screen.)

And to have access to these information just access the address: http://127.0.0.1:[port]/vulnerable/plugin/[Plugin Name]/[Version Plugin]
*/
func GetPluginVuln(writer http.ResponseWriter, response *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	parameters := mux.Vars(response)

	fmt.Println(parameters["Plugin"], parameters["Version"])

	vuln := database.GetVulnerabilities(parameters["Plugin"], parameters["Version"])

	jsonData, err := json.Marshal(vuln)

	if err != nil {
		log.Println("GetPluginVuln (plugins.go) :", err)
	}

	jsonData, _ = UnescapeCharactersJSON(jsonData)

	writer.Write([]byte(jsonData))
}

// GetPluginNameList ::
func GetPluginNameList(writer http.ResponseWriter, response *http.Request) {
	writer.Write([]byte("Plugin List Name"))
}

func UnescapeCharactersJSON(_jsonRaw json.RawMessage) (json.RawMessage, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(_jsonRaw)), `\\u`, `\u`, -1))

	if err != nil {
		return nil, err
	}

	return []byte(str), nil
}
