// Service to randomise the value of a salt variable on algolia index
package main

import (
	"fmt"
	"net/http"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

var (
	client algoliasearch.Client
	index  algoliasearch.Index
)

func init() {

	algoliaAppID := getEnv("ALGOLIA_APP_ID", "")
	algoliaAPIKey := getEnv("ALGOLIA_API_KEY", "")
	algoliaIndex := getEnv("ALGOLIA_INDEX", "")

	client = algoliasearch.NewClient(algoliaAppID, algoliaAPIKey)
	index = client.InitIndex(algoliaIndex)

	http.HandleFunc("/", handleRoot)
}

func main() {
	appengine.Main()
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	log.Infof(ctx, "")
	fmt.Fprintln(w, fmt.Sprintf("%v", 1))
}
