package todo

import (
	"github.com/danielAang/todo_list/internal"
	"github.com/go-chi/chi/v5"
)

func New(configuration *internal.Config) func() *chi.Mux {
	todoRepo := NewTodoRepo(configuration.Database)
	baseHandler := newBaseHandler(todoRepo)
	return func() *chi.Mux {
		return routes(baseHandler, configuration)
	}
}
