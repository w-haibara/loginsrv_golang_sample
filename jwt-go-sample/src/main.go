package main

import (
	"jwt-go-sample/auth"
	"log"
	"net/http"
)

func main() {
	auth.Configure(auth.Config{
		LoginFormURI:  "/login?backTo=/sample",
		VerifyKeyFile: "secret.key",
	})

	mainHandler := http.HandlerFunc(mainHandler)

	http.Handle("/", auth.Handler(mainHandler))
	http.ListenAndServe(":8080", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("--- main handler ---")
	w.Write([]byte("OK"))
}
