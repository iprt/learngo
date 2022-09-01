// Package citylist implements citylist generator.
package citylist

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"iproute.org/learngo/tutorial/mockserver/config"
)

// Generator represents the citylist generator.
type Generator struct {
	Tmpl *template.Template
}

// HandleRequest is the gin request filelisting for citylist generation.
func (g *Generator) HandleRequest(c *gin.Context) {
	err := g.generate(c.Writer)

	if err != nil {
		log.Printf("Cannot generate page for citylist: %v.", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (g *Generator) generate(w io.Writer) error {
	return g.Tmpl.Execute(w, struct {
		ServerAddress string
	}{
		ServerAddress: config.ServerAddress,
	})
}
