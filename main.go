package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Joke struct {
	ID    int    `json:"id" binding:"required"`
	Likes int    `json:"likes"`
	Joke  string `json:"joke" binding:"required"`
}

var jokes = []Joke{
	{1, 0, "Ano tawag sa cellphone ng mga lola? A: Edi MotoLOLA"},
	{2, 0, "Ano tawag sa maliit na kambing? A: Edi KapirangGOAT"},
	{3, 0, "Ano ang similarity ang UTOT at TULA? A: Pareho silang nagmula sa POET."},
	{4, 0, "Bakit maswerte ang kalendaryo? A: Dahil marami siyang date."},
	{5, 0, "Bakit malungkot ang kalendaryo? A:  Kasi bilang na ang araw niya."},
	{6, 0, "Anong puno ang hindi pwedeng akyatin? A: eh di yung nakatumba!"},
	{7, 0, "Ano ang pwede mong gawin sa GABI na hindi mo pwedeng gawin sa UMAGA? A: eh di MAGPUYAT."},
}

func GetAllJokes(c *gin.Context) {

	c.JSON(http.StatusOK, jokes)
}

func LikeJoke(c *gin.Context) {

	if jokeID, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		for i := 0; i < len(jokes); i++ {
			if jokes[i].ID == jokeID {
				jokes[i].Likes += 1
			}
		}
		c.JSON(http.StatusOK, &jokes)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()

	router.Use(cors.Default())

	api := router.Group("/api")
	api.GET("/jokes", GetAllJokes)
	api.POST("/jokes/like/:jokeID", LikeJoke)

	router.Run(":" + port)
}
