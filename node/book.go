package node

import "neural_network/bookForm"

type Book struct {
    Form
}

func CreateBook() (book Book) {
    book.Init(bookForm.CollectionProperties)
    return
}

func (book Book) SetData()  {

}