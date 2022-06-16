package newsfeed

import (
	"database/sql"
	"log"
)

type Feed struct {
	DB *sql.DB
}


//CONSTRUCTOR FUNCTION
func NewFeed(db *sql.DB) *Feed {
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXIST "newsfeeds" (
	"ID"	INTEGER NOT NULL,
	"Content"	TEXT NOT NULL,
	PRIMARY KEY("ID" AUTOINCREMENT);
`)
if err != nil {
	log.Println("Error creating newsfeed table.%v", err)
}

//excecute the statement
stmt.Exec()
	return &Feed{
		DB: db,
	}
}


//get method
func (feed *Feed) Get() []Item{
	items := []Item{}

	rows, err := feed.DB.Query(`
	SELECT * FROM newsFeed
	`)
	if err != nil {
		log.Printf("Error getting feeds %v", err)
		return nil
	}

	var id int
	var content string
	for rows.Next(){
		
		rows.Scan(&id, &content)

		newItem := Item{
			ID:id,
			Content:content,
		}

		items = append(items, newItem)
	}

	return items
}

//an add method
func (feed *Feed) Add(item Item){

	stmt2, err :=feed.DB.Prepare(`INSERT INTO newsfeed (content) values (?)
	`)

	if err != nil {
		log.Printf("Erorr insering newsfeeds content %v", err)
	}

	stmt2.Exec(item.Content)

}