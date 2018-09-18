// Service to randomise the value of a salt variable on algolia index
package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

var (
	algoliaAppID  string
	algoliaAPIKey string
	algoliaIndex  string
)

func init() {
	algoliaAppID = getEnv("ALGOLIA_APP_ID", "")
	algoliaAPIKey = getEnv("ALGOLIA_API_KEY", "")
	algoliaIndex = getEnv("ALGOLIA_INDEX", "")

	http.HandleFunc("/", handleRoot)
}

func main() {
	appengine.Main()
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	client := algoliasearch.NewClient(algoliaAppID, algoliaAPIKey)
	index := client.InitIndex(algoliaIndex)

	fmt.Fprintln(w, "Endpoint hit success")
	log.Infof(ctx, "Getting index...")

	it, err := index.BrowseAll(algoliasearch.Map{"query": ""})
	if err != nil {
		log.Errorf(ctx, "Error retrieving results from index")
	}

	objects := []algoliasearch.Object{}
	var hit algoliasearch.Map
	for {
		if hit, err = it.Next(); err != nil {
			if err == algoliasearch.NoMoreHitsErr {
				log.Infof(ctx, "End of results")
			} else {
				log.Errorf(ctx, "Error while browsing results")
			}
			break
		}
		objectID := hit["objectID"]
		log.Debugf(ctx, fmt.Sprintf("Res %v", objectID))
		objects = append(objects, algoliasearch.Object{"objectID": objectID, "salt": rand.Uint32()})
	}

	log.Infof(ctx, fmt.Sprintf("Updating records: %d", len(objects)))
	_, err = index.PartialUpdateObjects(objects)
	if err != nil {
		log.Criticalf(ctx, fmt.Sprintf("Failed to update objects"))
	} else {
		log.Infof(ctx, fmt.Sprintf("Updated records"))
	}
}
