package main

import (
	"./controllers"
	"./utils"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
)

// Create views
var static = controllers.NewStatic()

// Home renders the Home page
func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	utils.Must(static.Home.Render(w, nil))
}

// Contact renders the Contact page
func Contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	utils.Must(static.Contact.Render(w, nil))
}

func main() {
	// Remove local DB on imageService start up
	os.Remove("./segmed.db")

	// Open DB connection
	database, err := sql.Open("sqlite3", "./segmed.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// DB Table
	statement, _ := database.Prepare(
		"CREATE TABLE IF NOT EXISTS taggedImages (id INTEGER PRIMARY KEY, imageId INTEGER, isTagged INTEGER)")
	statement.Exec()

	// MUX Router
	r := mux.NewRouter()

	// Create Image Service
	service := controllers.ImageService{
		Db:            database,
		ShowTableView: static.ShowTable,
	}

	// Routes
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/contact", Contact).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}", service.TagImage).Methods("POST")
	r.HandleFunc("/showtable", service.ShowTable).Methods("GET")

	// Assets
	imageHandler := http.FileServer(http.Dir("./images/"))
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", imageHandler))

	// Run Web Server
	fmt.Printf("Starting the server on localhost:%d...\n", 3000)
	http.ListenAndServe(":3000", r)
}


