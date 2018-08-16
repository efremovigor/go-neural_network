package main

import "neural_network/form"

var sourcePath = "data.json"
var dataSource Source

func main() {

    getDataSource()
    book := form.CreateBook()
    book.SetData()
}

func getDataSource()  {
    if dataSource.entity.Data == nil {
        dataSource.initDataSource()
    }
}





//{
//"properties": {
//"FreeTimeSpending": true,
//"Famous": true,
//"GoodReview": true,
//"SeparateBook": true,
//"Medium": true,
//"Prose": true,
//"Novel": true,
//"Satire": true,
//"InRussian": true,
//"MaleAuthor": true,
//"OtherEuropeanAuthor": true,
//"WrittenIn18thCentury": true,
//"Century18th": true,
//"MaleAndFemaleMainCharacters": true,
//"Youth": true,
//"LoveLine": true
//},
//"result": true
//}
