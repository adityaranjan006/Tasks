package book

type Book struct {
	Name     string
	Quantity int
}

func NewBook(name string, quantity int) *Book {
	return &Book{
		Name:     name,
		Quantity: quantity,
	}
}
