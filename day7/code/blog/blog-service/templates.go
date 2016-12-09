package blogservice

import (
	"log"

	"github.com/cbroglie/mustache"
)

type Templates struct {
	Index *mustache.Template
}

func NewTemplates(indexTemplateFile string) *Templates {
	index, err := mustache.ParseFile(indexTemplateFile)
	if err != nil {
		log.Panicf("Error when parsing template file %q: %v", indexTemplateFile, err)
	}
	return &Templates{
		Index: index,
	}
}
