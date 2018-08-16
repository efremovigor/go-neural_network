package form

//назначение
var Business = "Business"
var Study = "Study"
var FreeTimeSpending = "FreeTimeSpending"
//популярность
var Famous = "Famous"
//отзывы
var NoReview = "NoReview"
var BadReview = "BadReview"
var GoodReview = "GoodReview"
//формат
var SeparateBook = "SeparateBook"
var SeriesOfBooks = "SeriesOfBooks"
var SeriesOfSmallStories = "SeriesOfSmallStories"
//размер
var Long = "Long"
var Medium = "Medium"
var Short = "Short"
//способ организации речи
var Poem = "Poem"
var Prose = "Prose"
//жанр по-форме
var Biography = "Biography"
var ShortStory = "ShortStory"
var Story = "Story"
var Novel = "Novel"
var Play = "Play"
//жанр по-содержанию
var Tragedy = "Tragedy"
var TragicComedy = "TragicComedy"
var Fantasy = "Fantasy"
var Mythology = "Mythology"
var Mystery = "Mystery"
var ScienceFiction = "ScienceFiction"
var Drama = "Drama"
var Romance = "Romance"
var Adventure = "Adventure"
var Satire = "Satire"
var Horror = "Horror"
var Erotic = "Erotic"
var NotFixed = "NotFixed"
//язык
var InEnglish = "InEnglish"
var InRussian = "InRussian"
var InOtherLanguage = "InOtherLanguage"
//пол автора
var MaleAuthor = "MaleAuthor"
var FemaleAuthor = "FemaleAuthor"
var Another = "Another"
//национальная принадлежность автора
var RussianAuthor = "RussianAuthor"
var BritishAuthor = "BritishAuthor"
var OtherEuropeanAuthor = "OtherEuropeanAuthor"
var NorthAmericanAuthor = "NorthAmericanAuthor"
var AsianAuthor = "AsianAuthor"
var OtherAuthor = "OtherAuthor"
//Период написания
var WrittenBefore16thCentury = "WrittenBefore16thCentury"
var WrittenIn16thCentury = "WrittenIn16thCentury"
var WrittenIn17thCentury = "WrittenIn17thCentury"
var WrittenIn18thCentury = "WrittenIn18thCentury"
var WrittenIn19thCentury = "WrittenIn19thCentury"
var WrittenIn20thCentury = "WrittenIn20thCentury"
var WrittenIn21stCentury = "WrittenIn21stCentury"
//эпоха, о которой идет речь
var BC = "BC"
var Before18thCentury = "Before18thCentury"
var Century18th = "Century18th"
var Century19th = "Century19th"
var Century20th = "Century20th"
var Present = "Present"
var Future = "Future"
//пол главного героя
var MaleMainCharacter = "MaleMainCharacter"
var FemaleMainCharacter = "FemaleMainCharacter"
var MaleAndFemaleMainCharacters = "MaleAndFemaleMainCharacters"
//возраст
var Kids = "Kids"
var Teenagers = "Teenagers"
var Youth = "Youth"
var MiddleAged = "MiddleAged"
var Old = "Old"
//существование любовной линии
var LoveLine = "LoveLine"
var NoLoveLine = "NoLoveLine"
var NotFixedLoveLine = "NotFixedLoveLine"

//базовые
var simplyBookProperties = []string{Famous}

//назначение
var appointmentBookProperties = []string{Business, Study, FreeTimeSpending}

//отзывы
var reviewsBookProperties = []string{NoReview, BadReview, GoodReview}

//формат
var formatBookProperties = []string{SeparateBook, SeriesOfBooks, SeriesOfSmallStories}

//размер
var sizeBookProperties = []string{Long, Medium, Short}

//способ организации речи
var wayOfSpeechOrganizationBookProperties = []string{Poem, Prose}

//жанр по-форме
var genreByFormBookProperties = []string{Biography, ShortStory, Story, Novel, Play}

//жанр по-содержанию
var genreByContentBookProperties = []string{Tragedy, TragicComedy, Fantasy, Mythology, Mystery, ScienceFiction, Drama, Romance, Adventure, Satire, Horror, Erotic, NotFixed}

//язык
var languageBookProperties = []string{InEnglish, InRussian, InOtherLanguage}

//пол автора
var sexOfAuthorBookProperties = []string{MaleAuthor, FemaleAuthor, Another}

//национальная принадлежность автора
var authorsNationalityBookProperties = []string{RussianAuthor, BritishAuthor, OtherEuropeanAuthor, NorthAmericanAuthor, AsianAuthor, OtherAuthor}

//Период написания
var theTimeOfWritingBookProperties = []string{WrittenBefore16thCentury, WrittenIn16thCentury, WrittenIn17thCentury, WrittenIn18thCentury, WrittenIn19thCentury, WrittenIn20thCentury, WrittenIn21stCentury}

//эпоха, о которой идет речь
var theEpochOfStoryBookProperties = []string{BC, Before18thCentury, Century18th, Century19th, Century20th, Present, Future}

//пол главного героя
var sexOfMainCharacterBookProperties = []string{MaleMainCharacter, FemaleMainCharacter, MaleAndFemaleMainCharacters}

//возраст
var ageBookProperties = []string{Kids, Teenagers, Youth, MiddleAged, Old}

//существование любовной линии
var existanceOfLoveLineBookProperties = []string{LoveLine, NoLoveLine, NotFixedLoveLine}

var bookCollectionProperties = append(append(append(append(append(append(append(append(append(append(append(append(append(append(append(
    simplyBookProperties,
    appointmentBookProperties...),
    reviewsBookProperties...),
    formatBookProperties...),
    sizeBookProperties...),
    wayOfSpeechOrganizationBookProperties...),
    genreByFormBookProperties...),
    genreByContentBookProperties...),
    languageBookProperties...),
    sexOfAuthorBookProperties...),
    authorsNationalityBookProperties...),
    theTimeOfWritingBookProperties...),
    theEpochOfStoryBookProperties...),
    sexOfMainCharacterBookProperties...),
    ageBookProperties...),
    existanceOfLoveLineBookProperties...)
