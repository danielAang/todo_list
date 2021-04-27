package todo

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/danielAang/todo_list/internal"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type baseHandler struct {
	todoRepository TodoRepository
}

func newBaseHandler(todoRepository TodoRepository) *baseHandler {
	return &baseHandler{
		todoRepository: todoRepository,
	}
}

func routes(b *baseHandler, config *internal.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{id}", b.GetById())
	router.Delete("/{id}", b.DeleteById())
	router.Post("/", b.Create())
	router.Get("/", b.GetAll())
	return router
}

func (h *baseHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		todo, err := h.todoRepository.FindById(id)
		if err != nil {
			render.JSON(w, r, err)
			return
		}
		render.JSON(w, r, todo)
	}
}

func (h *baseHandler) DeleteById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		err := h.todoRepository.Delete(id)
		if err != nil {
			render.JSON(w, r, err)
			return
		}
		render.Status(r, http.StatusAccepted)
	}
}
func (h *baseHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t Todo
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&t); err != nil {
			render.JSON(w, r, err)
			return
		}
		if err := h.todoRepository.Save(&t); err != nil {
			render.PlainText(w, r, err.Error())
			return
		}
		render.JSON(w, r, t)
	}
}

func (h *baseHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		skip, err := strconv.ParseInt(r.URL.Query().Get("skip"), 10, 64)
		if err != nil {
			render.JSON(w, r, err)
			return
		}
		limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
		if err != nil {
			render.JSON(w, r, err)
			return
		}
		todos, err := h.todoRepository.FindAll(skip, limit)
		if err != nil {
			render.JSON(w, r, err)
			return
		}
		render.JSON(w, r, todos)
	}
}
