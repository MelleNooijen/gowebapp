package structs

type PageData struct {
	Title string
	User  string
	Error IntError
}

// A User class provides a basic structure for a user.
type User struct {
	Id   int
	Name string
}

type IntError struct {
	Page   string
	Code   int
	Output string
}
