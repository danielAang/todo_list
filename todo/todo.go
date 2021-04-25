package todo

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Todo struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{id}", GetById)
	router.Delete("/{id}", DeleteById)
	router.Post("/", Create)
	router.Get("/", GetAll)
	return router
}

func GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	todo := Todo{
		Id:    id,
		Title: "Hello World",
		Body:  "Hello from Go!",
	}
	render.JSON(w, r, todo)
}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted successfully"
	render.JSON(w, r, response)
}
func Create(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created successfully"
	render.JSON(w, r, response)
}
func GetAll(w http.ResponseWriter, r *http.Request) {
	todo := []Todo{
		{
			Id:    "some_id",
			Title: "Hello World",
			Body:  "Hello from Go!",
		},
	}
	render.JSON(w, r, todo)
}
