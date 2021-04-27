package todo

type Todo struct {
	Id    string `json:"id" bson:"_id,omitempty"`
	Title string `json:"title" bson:"title,omitempty"`
	Body  string `json:"body" bson:"body,omitempty"`
}

type TodoRepository interface {
	FindById(ID string) (*Todo, error)
	FindAll(skip, limit int64) ([]Todo, error)
	Save(todo *Todo) error
	Delete(ID string) error
}
