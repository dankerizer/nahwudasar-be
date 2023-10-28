package presenter

type Meta struct {
	Rows        int `json:"rows,omitempty"`
	CurrentPage int `json:"currentPage,omitempty"`
	Totals      int `json:"totals,omitempty"`
}

type Result[T any] struct {
	Meta *Meta `json:"meta,omitempty"`
	Data []T   `json:"data"`
}
