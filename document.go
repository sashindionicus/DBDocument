package DBDocument

type Document struct {
	ID     int     `json:"id" db:"id"`
	Title  string  `json:"title" db:"title"`
	Author *Author `json:"author,omitempty"`
}

type Author struct {
	Firstname string `json:"firstname" db:"first_name"`
	Lastname  string `json:"lastname" db:"last_name"`
}
