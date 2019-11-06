package handler

import (
	"fmt"
	"github.com/pusher/pusher-http-go"
	"io/ioutil"
	"net/http"
	"os"
)

var client = pusher.Client{
	AppID:   os.Getenv("PUSHER_APP_ID"),
	Key:     os.Getenv("PUSHER_KEY"),
	Secret:  os.Getenv("PUSHER_SECRET"),
	Cluster: os.Getenv("PUSHER_CLUSTER"),
}

func Handler(w http.ResponseWriter, r *http.Request) {
	params, _ := ioutil.ReadAll(r.Body)

	response, err := client.AuthenticatePrivateChannel(params)
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()))
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	fmt.Fprintf(w, string(response))
}
