package web

type CourseCreateInput struct {
	AuthorId    int    `json:"author_id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Perks       string `json:"perks"`
	Price       int    `json:"price"`
	Banner      string `json:"banner"`
}

type CourseResponse struct {
	Id          int    `json:"id"`
	AuthorId    int    `json:"author_id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Perks       string `json:"perks"`
	Price       int    `json:"price"`
	Banner      string `json:"banner"`
}
