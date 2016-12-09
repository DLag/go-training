package blogservice

import "gopkg.in/gin-gonic/gin.v1"

type Service struct {
	templates    *Templates
	blog         *Blog
	listen       string
	indexHandler func(c *gin.Context)
}

func NewService(indexTemplateFile, dbFile, listen string) *Service {
	s := new(Service)
	s.blog = NewModelBlog(dbFile)
	s.templates = NewTemplates(indexTemplateFile)
	s.indexHandler = NewIndexHandler(s.templates, s.blog)
	s.listen = listen
	return s
}

func (s *Service) Run() {
	r := gin.Default()
	r.GET("/", s.indexHandler)
	r.Static("/css", "./assets/css")
	r.Static("/js", "./assets/js")
	r.Static("/fonts", "./assets/fonts")
	r.Run(s.listen)
}
