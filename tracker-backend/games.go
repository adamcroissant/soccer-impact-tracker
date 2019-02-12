package p

import (
	"context"
	"net/http"
	"encoding/json"
	"os"
	"cloud.google.com/go/datastore"
	"time"
)

type Game struct {
	Date		time.Time
	Opponent	string
	Result		string
	DidAttend	bool
}

func GetGames(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    q := datastore.NewQuery("Game")

	var games []Game
    _, err = client.GetAll(ctx, q, &games)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	js, err := json.Marshal(games)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
    	return
	}
	os.Stdout.Write(js)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
}
