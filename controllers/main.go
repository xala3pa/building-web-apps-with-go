package main

import (
	"net/http"

	"gopkg.in/unrolled/render.v1"
)

type Action func(rw http.ResponseWriter, r *http.Request) error

type AppController struct{}

func (c *AppController) Action(a Action) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := a(rw, r); err != nil {
			http.Error(rw, err.Error(), 500)
		}
	})
}

type MyController struct {
	AppController
	*render.Render
}

func (c *MyController) getJSON(rw http.ResponseWriter, r *http.Request) error {
	c.JSON(rw, 200, map[string]string{"Hello": "JSON"})
	return nil
}

func (c *MyController) getData(rw http.ResponseWriter, r *http.Request) error {
	c.Data(rw, http.StatusOK, []byte("Some binary data here."))
	return nil
}

func (c *MyController) Action(a Action) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := a(rw, r); err != nil {
			c.HTML(w, http.StatusOK, "error", nil)
		}
	})
}

func main() {
	c := &MyController{Render: render.New(render.Options{})}

	http.Handle("/json", c.Action(c.getJSON))
	http.Handle("/data", c.Action(c.getData))

	http.ListenAndServe(":8080", nil)
}
