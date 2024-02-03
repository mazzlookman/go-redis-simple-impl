package web

type CourseCreateInput struct {
	AuthorId    int
	Title       string
	Slug        string
	Description string
	Perks       string
	Price       int
	Banner      string
}

type CourseResponse struct {
	Id          int
	AuthorId    int
	Title       string
	Slug        string
	Description string
	Perks       string
	Price       int
	Banner      string
}
