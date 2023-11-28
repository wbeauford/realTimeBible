package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func renderIndex(c *gin.Context) {
	bibles, err := GetBibles()
	if err != nil {
		log.Println(err.Error())
	}
	render(c, gin.H{"payload": bibles}, "index.html")

}

func renderBooks(c *gin.Context) {
	r := c.Request
	bibleID := r.FormValue("bibleID")
	log.Println(bibleID)
}
