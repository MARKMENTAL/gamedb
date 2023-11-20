package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mysql", "root:pwhere@tcp(localhost:3306)/gamedb")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func getVideoGames() ([]VideoGame, error) {
	rows, err := db.Query("SELECT * FROM video_games")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []VideoGame
	for rows.Next() {
		var g VideoGame
		if err := rows.Scan(&g.ID, &g.Title, &g.Genre, &g.Developer, &g.Publisher, &g.ReleaseDate, &g.Platform, &g.Rating, &g.BoxArtURL); err != nil {
			return nil, err
		}
		games = append(games, g)
	}
	return games, nil
}

type VideoGame struct {
	ID          int
	Title       string
	Genre       string
	Developer   string
	Publisher   string
	ReleaseDate string
	Platform    string
	Rating      float64
	BoxArtURL   string
}

// StarRating returns a string of star characters based on the rating.
func (g VideoGame) StarRating() string {
	starCount := int(g.Rating) // Convert to int for simplicity; adjust as needed
	return strings.Repeat("★", starCount)
}

// Logger middleware to log IP address
func logger(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		log.Printf("Access from IP: %s", ip)
		handler(w, r)
	}
}

func main() {
	initDB()
	http.HandleFunc("/", logger(handleIndex))
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	games, err := getVideoGames()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/video_games.html"))
	tmpl.Execute(w, games)
}
