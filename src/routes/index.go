package routes

import (
	"github.com/galaco/yamigo"
)

type index struct {
	yamigo.Route
}

func (page index) Execute(request yamigo.Request, response yamigo.Response) {
	template := yamigo.NewTemplate("index.html.tmpl")

	template.AddParam("title", "Ika")

	response.SetTemplate(template)

	response.Execute()
}

func Index() (route index) {
	return route
}