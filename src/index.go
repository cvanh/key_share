package main

import (
	"fmt"
	"io"
	"log"

	"github.com/gliderlabs/ssh"
)

func main() {
	ssh.Handle(func(s ssh.Session) {
		log.Println("new connection")
		io.WriteString(s, fmt.Sprintf("Hello %s \n", s.User()))

	})

	log.Println("starting ssh server on port 2222...")
	server := &ssh.Server{
		Addr:             ":2222",
		PublicKeyHandler: authHandlerkaas,
	}

	log.Fatal(server.ListenAndServe())
}

func authHandlerkaas(ctx ssh.Context, key ssh.PublicKey) bool {
	log.Println(ssh.ParsePublicKey(key))
	return true
}
