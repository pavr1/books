package renderer

import (
	"net/http"

	"github.com/go-chi/render"
)

type CustomRenderer struct {
	StatusCode int
	Status     string
	Data       interface{}
}

func (c CustomRenderer) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, c.StatusCode)
	render.Respond(w, r, c)

	return nil
}

func NewRenderer(statusCode int, err error, data interface{}) render.Renderer {
	status := ""
	if err != nil {
		status = err.Error()
	} else {
		status = "Ok"
	}

	return CustomRenderer{
		StatusCode: statusCode,
		Status:     status,
		Data:       data,
	}
}
