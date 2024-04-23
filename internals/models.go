package internals

type Todo struct {
	Id          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}

type TodoRequest struct {
	Description string `json:"description,omitempty"`
}

type TodoResponse struct {
	Id string `json:"id,omitempty"`
}
