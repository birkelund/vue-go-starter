package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/rakyll/statik/fs"

	_ "github.com/birkelund/vue-go-starter/statik"
)

//go:generate statik -f -src=./ui/dist

var httpAddr = flag.String("http", ":8080", "address to listen on")

func main() {
	flag.Parse()

	statikfs, err := fs.New()
	if err != nil {
		log.Fatalf("could not create statik virtual file system: %v", err)
	}

	// we use a chi router
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})
	})

	r.Mount("/static/", http.FileServer(statikfs))

	// Since our app is a single page client-side app using HTML5 history mode,
	// users will get a 404 error if they access urls directly. We fix this by
	// adding a catch-all fallback route such that if the request did not match
	// anything in /api/ or /static/, we serve the main index.html from the statik
	// virtual file system.
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		f, err := statikfs.Open("/index.html")
		if err != nil {
			log.Printf("could not open index.html: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		fi, err := f.Stat()
		if err != nil {
			log.Printf("could not stat index.html: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		http.ServeContent(w, r, "", fi.ModTime(), f)
	})

	log.Fatal(http.ListenAndServe(*httpAddr, r))
}
