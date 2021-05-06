

package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"firebase.google.com/go/v4"
)

var (
	app *firebase.App
	ctx context.Context
)

type Profile struct {
	Name    string
	Hobbies []string
}

func main() {
	port			:= os.Getenv("REMOTE_PORT");	if port	==	""	{ port = "8080"	}
	defaultRoute	:= os.Getenv("ROUTE_BASE")

	ctx		= context.Background()
	app, _	= firebase.NewApp(ctx, nil);

	http.HandleFunc( defaultRoute, handler )

	log.Printf("Listening on port %s for %s", port, defaultRoute)
	if err	:= http.ListenAndServe(":"+port, nil);	err != nil { log.Fatal(err)}
}


func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != os.Getenv("ROUTE_BASE") {
		http.NotFound(w, r)
		return
	}

	profile := Profile{"Wes", []string{"skiing", "programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client, err		:= app.Auth(ctx); if err != nil { log.Fatal(err) }
	auth			:= r.Header.Get("Authorization")
	idToken 		:= strings.Split(auth, "Bearer ")
	log.Println("!!!!!!!!!!!!!!!!!!!!!", auth		)
	log.Println("!!!!!!!!!!!!!!!!!!!!!", idToken[1] )
	newToken, err	:= client.VerifyIDToken(ctx, idToken[1]);
	if err != nil {
		log.Fatalf("ERROR verifying ID token: %+v\nNew token: %+v\nClient: %+v", err, newToken, client)
		} else {
		log.Printf("Verified ID and received new token: %v", newToken)
	}


	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
