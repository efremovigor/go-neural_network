package node

import "neural_network/bookForm"

type Book struct {
    Form
}

func CreateBook() (book FormInterface) {
    book = &Book{}
    book.Init(bookForm.CollectionProperties)
    return
}

func (book *Book) SetProperty(name string, value bool) {
    book.Properties[name].State = value
}

func (book *Book) AutoSetProperties() {
    book.setDefault(bookForm.AppointmentBookProperties, bookForm.FreeTimeSpending)
    book.setDefault(bookForm.LanguageBookProperties, bookForm.InRussianLanguage)
    book.setDefault(bookForm.WayOfSpeechOrganizationBookProperties, bookForm.Prose)
    book.setDefault(bookForm.ExistanceOfLoveLineBookProperties, bookForm.LoveLine)
}
