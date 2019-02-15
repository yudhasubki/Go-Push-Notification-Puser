package main

import (
	"html/template"
	"net/http"

	"github.com/pusher/pusher-http-go"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/push", pushNotification)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("index.html")

	tmpl.Execute(w, nil)
}

func pushNotification(w http.ResponseWriter, r *http.Request) {
	client := pusher.Client{
		AppId:   "your-app-id",     //pusher app id
		Key:     "your-key-id",     // pusher key id
		Secret:  "your-key-secret", //pusher key secret
		Cluster: "ap1",
		Secure:  true,
	}

	data := map[string]string{"message": "notif me!"}
	client.Trigger("your-channel", "your-event", data)
}
