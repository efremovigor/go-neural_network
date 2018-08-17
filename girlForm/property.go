package girlForm

const BeautifulFace = "beautiful_face"
const VeryBigTits = "very_big_tits"
const BigTits = "big_tits"
const SmallTits = "small_tits"
const NoTits = "no_tits"
const LongHair = "long_hair"
const ShortHair = "short_hair"
const CrazyHair = "crazy_hair"
const Hairless = "hairless"
const ShapelyBody = "shapely_body"
const SportBody = "sport_body"
const FatBody = "fat_body"
const LeanBody = "lean_body"
const Blonde = "blonde"
const Brunette = "brunette"
const Redhead = "redhead"
const MyHeight = "my_height"
const HighHeight = "high_height"
const LowHeight = "low_height"
const VeryLowHeight = "very_low_height"
const Slavic = "slavic"
const Afro = "afro"
const Asian = "asian"
const FewTeeth = "few_teeth"
const DirtyTeeth = "dirty_teeth"
const CrookedTeeth = "crooked_teeth"
const NormalTeeth = "normal_teeth"
const VulgarlyClothes = "vulgarly_clothes"
const HumbleClothes = "humble_clothes"
const ExpensiveClothes = "expensive_clothes"
const StylishClothes = "stylish_clothes"
const Neat = "neat"
const Money = "money"
const House = "house"
const Car = "car"
const RichParents = "rich_parents"
const Smoke = "smoke"
const Drink = "drink"
const Fitness = "fitness"
const Snowboard = "snowboard"
const Bouldering = "bouldering"
const Student = "student"
const Working = "working"
const Developer = "developer"
const VideoGame = "video_game"
const Swearing = "swearing"
const Dance = "dance"
const LookAtMe = "look_at_me"
const PassesBy = "passes_by"
const InNetwork = "in_network"
const DaughterOfMothersGirlfriend = "daughter_of_mother's_girlfriend"

var simplyProperties = []string{BeautifulFace, Neat, Money, House, Car, RichParents, Smoke, Drink, Fitness, Snowboard, Bouldering, Student, Working, Developer, VideoGame, Swearing,LookAtMe}

var TitsProperties = []string{VeryBigTits, BigTits, SmallTits, NoTits,}

var hairProperties = []string{LongHair, ShortHair, CrazyHair, Hairless,}

var bodyProperties = []string{ShapelyBody, SportBody, FatBody, LeanBody,}

var colorHairProperties = []string{Blonde, Brunette, Redhead,}

var heightProperties = []string{MyHeight, HighHeight, LowHeight, VeryLowHeight,}

var nationProperties = []string{Slavic, Afro, Asian}

var teethProperties = []string{FewTeeth, DirtyTeeth, CrookedTeeth, NormalTeeth}

var clothesProperties = []string{VulgarlyClothes, HumbleClothes, ExpensiveClothes, StylishClothes}

var actionProperties = []string{Dance, PassesBy, InNetwork, DaughterOfMothersGirlfriend}

var CollectionProperties = append(append(append(append(append(append(append(append(append(
    simplyProperties,
    TitsProperties...),
    hairProperties...),
    bodyProperties...),
    colorHairProperties...),
    heightProperties...),
    nationProperties...),
    teethProperties...),
    clothesProperties...),
    actionProperties...)
