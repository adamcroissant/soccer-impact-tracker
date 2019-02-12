package p

import (
	"net/http"
	"encoding/json"
	"os"
)

type Game struct {
	Date		string
	Opponent	string
	Result		string
	DidAttend	bool
}

func GetGames(w http.ResponseWriter, r *http.Request) {
	game := Game{
		Date: "01/01/2019",
		Opponent: "Express",
		Result: "Win",
		DidAttend: true,
	}
	result := []Game{game}
	js, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
    		return
	}
	os.Stdout.Write(js)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
}
