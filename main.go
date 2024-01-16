package main

import (
	"app/crawler/crawler"
	"app/crawler/dbutil"
	"app/crawler/model"
	"app/crawler/repo/mysql"
	"context"
	"fmt"
)

// func main() {
// 	db := dbutil.ConnectDB()
// 	fmt.Println(db)

// 	// AutoMigrate will create tables if they don't exist
// 	db.AutoMigrate(&model.WebPage{})

// 	paragraphsP, paragraphsA, urls, titles, html, err := crawler.Crawler()
// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 		return
// 	}

// 	// Create a new WebPage instance with the collected data
// 	webPage := model.WebPage{
// 		URL:   urls[0],
// 		HTML:  html[0],
// 		Title: titles[0],
// 		P_Tag: paragraphsP[0],
// 		A_Tag: paragraphsA[0],
// 	}

// 	// Create a new repository instance
// 	webPageRepo := mysql.NewWebPageRepository(db)

// 	// Insert the webPage into the database
// 	if err := webPageRepo.InsertOne(context.Background(), &webPage); err != nil {
// 		fmt.Println("Error inserting into database: ", err)
// 		return
// 	}

// 	fmt.Println("WebPage inserted successfully.")
// }

func main() {

	db := dbutil.ConnectDB()

	paragraphsP, paragraphsA, urls, titles, html, err := crawler.Crawler()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(urls)
	fmt.Println(titles)
	fmt.Println(html)

	pChan := make(chan string, 100)
	aChan := make(chan string, 100)
	// urlChan := make(chan string, 100)
	// titleChan := make(chan string, 100)
	// htmlChan := make(chan string, 100)

	go func() {
		for _, p := range paragraphsP {
			pChan <- p
		}
		close(pChan)
	}()

	go func() {
		for _, a := range paragraphsA {
			aChan <- a
		}
		close(aChan)
	}()

	// Tương tự cho urlChan, titleChan, htmlChan

	repo := mysql.NewWebPageRepository(db)

	for {
		select {
		case p := <-pChan:
			page := model.WebPage{P_Tag: p}
			repo.InsertOne(context.Background(), &page)

		case a := <-aChan:
			page := model.WebPage{A_Tag: a}
			repo.InsertOne(context.Background(), &page)

			// ...
		}
	}
}
