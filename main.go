
package main

import (
  // import in standard libraries
	"net/http"
	"os"

  // import the web framework
	"github.com/gin-gonic/gin"
)


func main() {

	r := gin.Default()


	r.LoadHTMLGlob("*.html")

	r.Static("/public", "public")


	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}


	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"HelloMessage": "Sup World!",
		})
	})


	r.Run(":" + port)
}
