package ssh

import (
	"fmt"
	"io"
	"log"

	db "../../pkg/database"
	"github.com/gliderlabs/ssh"
)

func main() {
	db.GetDatabase()

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
	// signer, err := ssh.ParsePublicKey(key)

	// if err != nil {
	// 	log.Println("error while parsing ssh key")
	// }
	// log.Println("authenticating user with key %s", signer)

	return true
}