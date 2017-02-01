package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"os"
)

type Shoe struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Style string
	Color string
	Size  string
	Brand string
}

func main() {

	r := gin.Default()
	lab := os.Getenv("MONGOLAB_URI")
	db := os.Getenv("Shoe")

	r.LoadHTMLGlob("*.html")

	r.Static("/public", "public")
	//
	session, err := mgo.Dial(lab)
	col := session.DB(db).C("shoes")
	if err != nil {
		panic(err)

	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"

	}

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"HelloMessage": "Sup World!",
	// 	})
	// })

	r.GET("/", func(c *gin.Context) {
		var shoes []Shoe
		col.Find(nil).All(&shoes)
		c.JSON(http.StatusOK, gin.H{
			"shoes": shoes,
		})

	})

	r.POST("/", func(c *gin.Context) {
		style := c.PostForm("style")
		color := c.PostForm("color")
		size := c.PostForm("size")
		brand := c.PostForm("brand")
		// query := bson.M{"current": true}
		// change := bson.M{"$set": bson.M{"current": false}}
		// col.UpdateAll(query, change)
		err = col.Insert(&Shoe{Style: style, Color: color, Size: size, Brand: brand})
		if err != nil {
			panic(err)
		}
		c.Redirect(http.StatusMovedPermanently, "/")
	})


	r.GET("/shoe/:id", func(c *gin.Context) {
		shoes := Shoe{}
		style := c.Param("id")
		query := bson.M{"style": style}
		col.Find(query).One(&shoes)
		c.JSON(http.StatusOK, gin.H{
			"shoe": shoes,
		})
	})

 r.POST("/delete/:id", func(c *gin.Context) {
    style := c.Param("id")
 		err := col.Remove(bson.M{"style": style})
 		if err != nil {
 			panic(err)
 		}

    c.JSON(http.StatusOK, gin.H{
      "deleted": "deleted",
    })
 	   	c.Redirect(http.StatusMovedPermanently, "/")
 	})

  r.POST("/edit/:id", func(c *gin.Context) {
     style := c.Param("id")
     newstyle := c.PostForm("style")
     color := c.PostForm("color")
     size := c.PostForm("size")
     brand := c.PostForm("brand")

     update := bson.M{
			"style":       newstyle,
			"color":       color,
			"size":        size,
			"Brand":       brand,
		}
     change := bson.M{"$set": update}
     col.Update(bson.M{ "style": style}, change)
     c.JSON(http.StatusOK, gin.H{
       "updated": "I guess so",
     })
       c.Redirect(http.StatusMovedPermanently, "/")
   })








	r.Run(":" + port)
}
