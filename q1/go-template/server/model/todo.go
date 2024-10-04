package model

type TodoData struct {
	/* id/unique not update able */
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

type TodoUnique struct {
	Id *int `json:"id"`
}

type TodoFilter struct {
	Id   *int    `json:"id"`
	Name *string `json:"name"`
}
