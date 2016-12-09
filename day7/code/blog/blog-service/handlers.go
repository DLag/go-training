package blogservice

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

func NewIndexHandler(view *Templates, model *Blog) func(c *gin.Context) {
	return func(c *gin.Context) {
		posts := model.GetAllPosts()
		log.Printf("Trying to render %d posts", len(posts))
		rendered, _ := view.Index.Render(map[string]interface{}{"Posts": posts})
		c.Data(200, "text/html", []byte(rendered))
	}
}
