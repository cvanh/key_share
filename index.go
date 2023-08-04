package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gliderlabs/ssh"
	"github.com/tyler-sommer/stick"
)

func main() {

	http_main()
	ssh_main()

	log.Println("new connection")
}

func ssh_main() {
	GetDatabase()

	ssh.Handle(func(s ssh.Session) {
		log.Println("new connection")
		io.WriteString(s, fmt.Sprintf("Hello %s \n", s.User()))

	})

	log.Println("starting ssh server on port 2222...")
	server := &ssh.Server{
		Addr:             ":2222",
		PublicKeyHandler: authHandlerkaas,
	}

	go log.Fatal(server.ListenAndServe())
}

func authHandlerkaas(ctx ssh.Context, key ssh.PublicKey) bool {
	// signer, err := ssh.ParsePublicKey(key)

	// if err != nil {
	// 	log.Println("error while parsing ssh key")
	// }
	// log.Println("authenticating user with key %s", signer)

	return true
}

func http_main() {
	fsRoot, _ := os.Getwd() // Templates are loaded relative to this directory.

	env := stick.New(stick.NewFilesystemLoader(fsRoot))
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		env.Execute("bar.html.twig", w, nil) // Loads "bar.html.twig" relative to fsRoot.
	})
	go log.Fatal(http.ListenAndServe(":8080", nil))
}
