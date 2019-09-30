package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/PuerkitoBio/goquery"
)

type Movies []Movie
func main() {
    // Make HTTP request
    response, err := http.Get("https://vivacinemas.com")
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    // Create a goquery document from the HTTP response
    document, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        log.Fatal("Error loading HTTP response body. ", err)
    }
	// getHomeSliders( document)
	getMovieForDay( "Mon", document)


	
}

func getHomeSliders(document *goquery.Document){
	var movies Movies
	document.Find(".movie-slide").Each(func(index int, element *goquery.Selection) {
		imgSrc, exists := element.Find("img").Attr("data-lazy-src")
		title := element.Find("h4").Text()

		if exists {
			movies = append(movies, Movie{Title : title, Image : imgSrc, Genre : "",Duration : ""})

		}else{
			fmt.Println("imgSrc")
		}
		fmt.Println(movies)

	})
}

func getMovieForDay(day string ,document *goquery.Document){
	var movies Movies

	valueToSearch := fmt.Sprintf("#%s .row.movie-tabs" ,day)
	fmt.Println(valueToSearch)
	document.Find(valueToSearch).Each(func(index int, element *goquery.Selection) {
		imgSrc, exists := element.Find("img").Attr("data-lazy-src")
		genre := element.Find(".title").Text()
		title := element.Find(".no-underline").Text()
		duration := element.Find(".running-time ").Text()

		if exists {
			movies = append(movies, Movie{Title : title, Image : imgSrc, Genre : genre, Duration: duration})
		}else{
			fmt.Println("Not found")
		}
	})
	fmt.Println(movies)

}


type Movie struct{
	Title string `json:"title"`
	Image string `json:"image"`
	Genre string `json:"genre"`
	Duration string `json:"duration"`
}