package blogservice

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

func NewIndexHandler(templates *Templates, blog *Blog) func(c *gin.Context) {
	return func(c *gin.Context) {
		posts := blog.GetAllPosts()
		log.Printf("Trying to render %d posts", len(posts))
		rendered, _ := templates.Index.Render(map[string]interface{}{"Posts": posts})
		c.Data(200, "text/html", []byte(rendered))
	}
}
