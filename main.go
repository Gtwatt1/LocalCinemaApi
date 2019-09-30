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
	getMovieForDay("Mon", document)

	
}

func getHomeSliders(document *goquery.Document){
	document.Find("img .lazyloaded").Each(func(index int, element *goquery.Selection) {
		imgSrc, exists := element.Attr("data-lazy-src")

		if exists {
			fmt.Println(imgSrc)
		}else{
			fmt.Println("Not Found")
		}
	})
}

func getMovieForDay(day string ,document *goquery.Document){
	var movies Movies

	valueToSearch := fmt.Sprintf("#%s .row.movie-tabs" ,day)
	fmt.Println(valueToSearch)
	document.Find(valueToSearch).Each(func(index int, element *goquery.Selection) {
		imgSrc, exists := element.Find("img").Attr("data-lazy-src")
		genre := element.Find(".title").Text()
		title := element.Find(" .no-underline").Text()

		if exists {
			movies = append(movies, Movie{Title : title, Image : imgSrc, Genre : genre})
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

}