package node

import "neural_network/bookForm"

type Book struct {
    Form
}

func CreateBook() (book FormInterface) {
    book.Init(bookForm.CollectionProperties)
    return
}

func (book Book) SetData()  {

}

func (book *Book) AutoSetProperties() {
    book.setDefault(bookForm.AppointmentBookProperties,bookForm.FreeTimeSpending)
    book.setDefault(bookForm.LanguageBookProperties,bookForm.InRussianLanguage)
    book.setDefault(bookForm.WayOfSpeechOrganizationBookProperties,bookForm.Prose)
    book.setDefault(bookForm.ExistanceOfLoveLineBookProperties,bookForm.LoveLine)
}
