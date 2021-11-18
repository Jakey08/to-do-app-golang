
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



func main() {

	todo := []string{
		"映画見に行く",
		"原宿に5時に行く",
		"帰りに柔軟剤を買う",

	}


	r := gin.Default()
	r.LoadHTMLFiles("./template/index.html")
	r.GET("/nullpo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message": "ga!",
		})
	})
	r.GET("/todo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]interface{}{
			"todo": todo,
		})
	})


	r.Run() //localhost:8080
}