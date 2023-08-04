package http

import (
	"log"
	"net/http"
	"os"

	"github.com/tyler-sommer/stick"
)

func main() {
	fsRoot, _ := os.Getwd() // Templates are loaded relative to this directory.

	env := stick.New(stick.NewFilesystemLoader(fsRoot))
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		env.Execute("bar.html.twig", w, nil) // Loads "bar.html.twig" relative to fsRoot.
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
