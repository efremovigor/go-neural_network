package bookForm

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
var InEnglishLanguage = "InEnglishLanguage"
var InRussianLanguage = "InRussianLanguage"
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
var SimplyBookProperties = []string{Famous}

//назначение
var AppointmentBookProperties = []string{Business, Study, FreeTimeSpending}

//отзывы
var ReviewsBookProperties = []string{NoReview, BadReview, GoodReview}

//формат
var FormatBookProperties = []string{SeparateBook, SeriesOfBooks, SeriesOfSmallStories}

//размер
var SizeBookProperties = []string{Long, Medium, Short}

//способ организации речи
var WayOfSpeechOrganizationBookProperties = []string{Poem, Prose}

//жанр по-форме
var GenreByFormBookProperties = []string{Biography, ShortStory, Story, Novel, Play}

//жанр по-содержанию
var GenreByContentBookProperties = []string{Tragedy, TragicComedy, Fantasy, Mythology, Mystery, ScienceFiction, Drama, Romance, Adventure, Satire, Horror, Erotic, NotFixed}

//язык
var LanguageBookProperties = []string{InEnglishLanguage, InRussianLanguage, InOtherLanguage}

//пол автора
var SexOfAuthorBookProperties = []string{MaleAuthor, FemaleAuthor, Another}

//национальная принадлежность автора
var AuthorsNationalityBookProperties = []string{RussianAuthor, BritishAuthor, OtherEuropeanAuthor, NorthAmericanAuthor, AsianAuthor, OtherAuthor}

//Период написания
var TheTimeOfWritingBookProperties = []string{WrittenBefore16thCentury, WrittenIn16thCentury, WrittenIn17thCentury, WrittenIn18thCentury, WrittenIn19thCentury, WrittenIn20thCentury, WrittenIn21stCentury}

//эпоха, о которой идет речь
var TheEpochOfStoryBookProperties = []string{BC, Before18thCentury, Century18th, Century19th, Century20th, Present, Future}

//пол главного героя
var SexOfMainCharacterBookProperties = []string{MaleMainCharacter, FemaleMainCharacter, MaleAndFemaleMainCharacters}

//возраст
var AgeBookProperties = []string{Kids, Teenagers, Youth, MiddleAged, Old}

//существование любовной линии
var ExistanceOfLoveLineBookProperties = []string{LoveLine, NoLoveLine, NotFixedLoveLine}

var CollectionProperties = append(append(append(append(append(append(append(append(append(append(append(append(append(append(append(
    SimplyBookProperties,
    AppointmentBookProperties...),
    ReviewsBookProperties...),
    FormatBookProperties...),
    SizeBookProperties...),
    WayOfSpeechOrganizationBookProperties...),
    GenreByFormBookProperties...),
    GenreByContentBookProperties...),
    LanguageBookProperties...),
    SexOfAuthorBookProperties...),
    AuthorsNationalityBookProperties...),
    TheTimeOfWritingBookProperties...),
    TheEpochOfStoryBookProperties...),
    SexOfMainCharacterBookProperties...),
    AgeBookProperties...),
    ExistanceOfLoveLineBookProperties...)
