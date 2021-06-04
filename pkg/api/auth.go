package api

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gleich/logoru"
	"golang.org/x/oauth2"
)

// Create a GitHub client.
func CreateClient() *http.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: strings.TrimSuffix(getToken(), "\n")},
	)
	return oauth2.NewClient(context.Background(), src)
}

// Get the GitHub PAT from the config.
func getToken() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logoru.Critical("Failed to get home dir:", err)
		os.Exit(1)
	}

	bin, err := ioutil.ReadFile(filepath.Join(homeDir, "cihat-config", "pat.txt"))
	if err != nil {
		logoru.Critical("Failed to read from config file:", err)
		os.Exit(1)
	}

	return string(bin)
}
