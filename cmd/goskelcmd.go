package main

import (
	"flag"
	"io"
	"log"
	"net/http"

	"github.com/appleby/goskeleton"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()

	log.Println("word:", *wordPtr)
	log.Println("numb:", *numbPtr)
	log.Println("fork:", *boolPtr)
	log.Println("svar:", svar)
	log.Println("tail:", flag.Args())

	foo := goskeleton.Foo{Name: "appleby"}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		_, err := io.WriteString(w, foo.Greet())
		if err != nil {
			log.Fatal(err)
		}
	})
	log.Println("Listening on http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
