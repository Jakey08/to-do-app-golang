
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	Content string
	Priority int //0: Low, 1: Medium, 2: High, other: Critical
	Status int //0: Not yet, 1: In progress, 2: Stuck, 3:On hold, other: Done
	//CreatedAt time.Time



}

func main() {


	todo := make([]Todo, 0, 5)
	todo =  []Todo{
		Todo{
			Content: "原宿で友達と合う",
			Priority: 1,
			Status: 0,
			//CreatedAt: time.Now(),
			},
		Todo{
			Content: "アイス買いに行く",
			Priority: 1,
			Status: 1,
			//CreatedAt: time.Now(),
		},
		Todo{
			Content: "横浜で焼き肉食べる",
			Priority: 1,
			Status: 2,
			//CreatedAt: time.Now(),
		},
		Todo{
			Content: "柔軟剤買う",
			Priority: 1,
			Status: 100,
			//CreatedAt: time.Now(),
		},

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