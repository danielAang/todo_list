package todo

import (
	"net/http"

	"github.com/danielAang/todo_list/internal"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Todo struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func Routes(config *internal.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{id}", GetById(config))
	router.Delete("/{id}", DeleteById(config))
	router.Post("/", Create(config))
	router.Get("/", GetAll(config))
	return router
}

func GetById(config *internal.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		todo := Todo{
			Id:    id,
			Title: "Hello World",
			Body:  "Hello from Go!",
		}
		render.JSON(w, r, todo)
	}
}

func DeleteById(config *internal.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := make(map[string]string)
		response["message"] = "Deleted successfully"
		render.JSON(w, r, response)
	}
}
func Create(config *internal.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := make(map[string]string)
		response["message"] = "Created successfully"
		render.JSON(w, r, response)
	}
}

func GetAll(config *internal.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todo := []Todo{
			{
				Id:    "some_id",
				Title: "Hello World",
				Body:  "Hello from Go!",
			},
		}
		render.JSON(w, r, todo)
	}
}
