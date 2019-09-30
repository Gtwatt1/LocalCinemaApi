package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
	"github.com/PuerkitoBio/goquery"
	"encoding/json"
	"strings"
)

type Movies []Movie

func main() {
   
	// getHomeSliders( document)
	// getMovieForDay( "Mon", document)
	r := mux.NewRouter()
	r.HandleFunc("/", handleHomeSlidersReq)
	r.HandleFunc("/movies/{day}", handleMovieForDay)
	log.Fatal(http.ListenAndServe(":5000", r))
}

func handleMovieForDay(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	day := params["day"]
	json.NewEncoder(w).Encode(getMovieForDay(day))

} 

func handleHomeSlidersReq(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(getHomeSliders())
}

func getDom() (*goquery.Document){
	response, err := http.Get("https://vivacinemas.com")
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()
	document, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        log.Fatal("Error loading HTTP response body. ", err)
	}
	return document
}


func getHomeSliders() ( Movies){
	document := getDom()
	var movies Movies
	document.Find(".movie-slide").Each(func(index int, element *goquery.Selection) {
		imgSrc, exists := element.Find("img").Attr("data-lazy-src")
		title := element.Find("h4").Text()

		if exists {
			movies = append(movies, Movie{Title : title, Image : imgSrc, Genre : "",Duration : "", Link : ""})
		}else{
			fmt.Println("imgSrc")
		}
		fmt.Println(movies)

	})
	return movies
}

func getMovieForDay(day string) (Movies){
	document := getDom()
	var movies Movies

	valueToSearch := fmt.Sprintf("#%s .row.movie-tabs" ,day)
	fmt.Println(valueToSearch)
	document.Find(valueToSearch).Each(func(index int, element *goquery.Selection) {
		imgSrc, exists := element.Find("img").Attr("data-lazy-src")
		genre := element.Find(".title").Text()
		title := element.Find(".no-underline").Text()
		duration := element.Find(".running-time ").Text()
		link, _ := element.Find("a").Attr("href")

		if exists {
			movies = append(movies, Movie{Title : title, Image : imgSrc, Genre : strings.TrimSpace(genre), Duration: strings.TrimSpace(duration), Link: link})
		}else{
			fmt.Println("Not found")
		}
	})
	return movies
}


type Movie struct{
	Title string `json:"title"`
	Image string `json:"image"`
	Genre string `json:"genre"`
	Duration string `json:"duration"`
	Link string `json:"link"`

}