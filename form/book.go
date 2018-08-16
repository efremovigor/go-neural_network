package form

type Book struct {
    Form
}

func CreateBook() (book Book) {
    book.init(bookCollectionProperties)
    return
}

func (book Book) SetData()  {

}