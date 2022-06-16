package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Durotimicodes/sqlite-learning/platform/newsfeed"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qkgo/yin"
)

func main() {


	//database connection
	db, err := sql.Open("sqlite3", "./phonenumbers.db")
	if err != nil {
		log.Printf("Error connecting to database is: %v", err)
		return 
	}

	feed := newsfeed.NewFeed(db)

	// feed.Add(newsfeed.Item{
	// 	Content: "Hello",
	// })

	// items := feed.Get()
	// fmt.Println(items)

	//create a router
	router := chi.NewRouter()

	//middleware
	router.Use(yin.SimpleLogger)

	router.Get("/posts", func(rw http.ResponseWriter, r *http.Request) {
		//get will return all the data that is in our database
		
		resp,_ := yin.Event(rw, r)
		//get items from database and send it back as a JSON
		items := feed.Get()
		resp.SendJSON(items)
	})

	router.Post("/posts", func(rw http.ResponseWriter, r *http.Request) {
		resp, req := yin.Event(rw, r)
		body := map[string]string{
			req.BindBody(&body)
			item := newsfeed.Item{
				Content: body["content"],
			}
			feed.Add(item)
			res.SendStatus(204)
		}

	})

	log.Println("Server listening on port 3000...")
	http.ListenAndServe(":3000", router)
}