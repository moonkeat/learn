package main

// Usage:
// GUI, Data structure travesal, Aggregate of objects, Database objects, File processing, Social Media feeds

// Pros:
// Seperate of concerns
// Hide complex data structure from clients
// Flexibility to traverse collection (forward, backward, etc)

// Cons:
// Complexity, overkill for simple use case
// Iterators become invalid if collection modified during interation
// Not safe for concurrent modification and iteration

func main() {
	bookCollection := &BookCollection{}
	bookCollection.AddBook(Book{Title: "Book 1"})
	bookCollection.AddBook(Book{Title: "Book 2"})

	iterator := bookCollection.CreateIterator()
	for iterator.HasNext() {
		book := iterator.Next()
		println(book.Title)
	}
}

type Book struct {
	Title string
}

type Iterator interface {
	HasNext() bool
	Next() *Book
}

type BookIterator struct {
	collection BookCollection
	index      int
}

func (i *BookIterator) HasNext() bool {
	return i.index < len(i.collection.items)
}

func (i *BookIterator) Next() *Book {
	if !i.HasNext() {
		return nil
	}
	item := i.collection.items[i.index]
	i.index++
	return &item
}

type Aggregate interface {
	CreateIterator() Iterator
}

type BookCollection struct {
	items []Book
}

func (c BookCollection) CreateIterator() Iterator {
	return &BookIterator{collection: c}
}

func (c *BookCollection) AddBook(book Book) {
	c.items = append(c.items, book)
}
