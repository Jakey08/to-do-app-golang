
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Todo struct {
	Content string
	Priority string
	Status string
	CreatedAt time.Time


}


func main() {

	todo := make([]Todo, 0, 5)
	t := Todo{
		Todo{
			Content: "掃除する",
			Priority: "High",
			Status: "Haven't Started",
			CreatedAt: time.Now(),
		} //なにこのエラー　解決したい
	}
	todo = append(todo, t)

	Todo {

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