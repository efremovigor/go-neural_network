package node

import "go-neural_network/bookForm"

type Book struct {
	Form
}

func CreateBook(entity DataEntity) (book FormInterface) {
	book = &Book{}
	book.Init(bookForm.CollectionProperties)
	for key, value := range entity.Properties {
		book.SetProperty(key, value)
	}
	book.SetResult(entity.Result)
	return
}

func (book *Book) AutoSetProperties() {
	book.setDefault(bookForm.AppointmentBookProperties, bookForm.FreeTimeSpending)
	book.setDefault(bookForm.LanguageBookProperties, bookForm.InRussianLanguage)
	book.setDefault(bookForm.WayOfSpeechOrganizationBookProperties, bookForm.Prose)
	book.setDefault(bookForm.ExistanceOfLoveLineBookProperties, bookForm.LoveLine)
}
