package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getItunesMovie(c *gin.Context) {
	search := c.Param("search")

	url := "https://itunes.apple.com/search?term=" + search + "&media=movie"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	itunes, err := UnmarshalItunes(body)

	c.IndentedJSON(http.StatusOK, itunes)

}

func getItunesMusic(c *gin.Context) {
	search := c.Param("search")

	url := "https://itunes.apple.com/search?term=" + search + "&media=music"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	itunes, err := UnmarshalItunes(body)

	c.IndentedJSON(http.StatusOK, itunes)

}

func getTvMazeShow(c *gin.Context) {
	search := c.Param("search")

	url := "https://api.tvmaze.com/search/shows?q=" + search

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	tvmaze, err := UnmarshalTvmaze(body)

	c.IndentedJSON(http.StatusOK, tvmaze)

}

func main() {
	router := gin.Default()

	router.GET("/itunes-movie/:search", getItunesMovie)
	router.GET("/itunes-music/:search", getItunesMusic)
	router.GET("/tv-maze-show/:search", getTvMazeShow)

	router.Run("localhost:8080")
}
