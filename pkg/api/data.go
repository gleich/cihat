package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Matt-Gleich/logoru"
)

// Get the checks from the API
func GetChecks(client *http.Client) QueryOutline {
	// Encoding the query
	jsonData, err := json.Marshal(map[string]string{"query": query})
	if err != nil {
		logoru.Error("Failed to encode json data:", err)
	}

	// Formulating the request
	request, err := http.NewRequest(
		"POST",
		"https://api.github.com/graphql",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		logoru.Error("Failed to form request:", err)
	}

	// Required for the checks preview
	request.Header.Set("Accept", "application/vnd.github.antiope-preview+json")

	// Making the request with the client
	response, err := client.Do(request)
	if err != nil {
		logoru.Error("Failed to make the request:", err)
	}
	defer response.Body.Close()

	// Reading response
	bin, _ := ioutil.ReadAll(response.Body)
	var data QueryOutline
	err = json.Unmarshal(bin, &data)
	if err != nil {
		logoru.Error("Failed to parse the data from GitHub:", err)
	}

	return data
}
