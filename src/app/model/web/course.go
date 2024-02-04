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
	Id          int    `json:"id" redis:"courseId"`
	AuthorId    int    `json:"author_id" redis:"authorId"`
	Title       string `json:"title" redis:"title"`
	Slug        string `json:"slug" redis:"slug"`
	Description string `json:"description" redis:"description"`
	Perks       string `json:"perks" redis:"perks"`
	Price       int    `json:"price" redis:"price"`
	Banner      string `json:"banner" redis:"banner"`
}
