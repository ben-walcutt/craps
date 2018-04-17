package lib

type Board struct {
	PlaceFour int
	PlaceFive int
	PlaceSix int
	PlaceEight int
	PlaceNine int
	PlaceTen int

	Come int
	ComeFour int
	ComeFourOdds int
	ComeFive int
	ComeFiveOdds int
	ComeSix int
	ComeSixOdds int
	ComeEight int
	ComeEightOdds int
	ComeNine int
	ComeNineOdds int
	ComeTen int
	ComeTenOdds int

	DontCome int
	DontComeFour int
	DontComeFourOdds int
	DontComeFive int
	DontComeFiveOdds int
	DontComeSix int
	DontComeSixOdds int
	DontComeEight int
	DontComeEightOdds int
	DontComeNine int
	DontComeNineOdds int
	DontComeTen int
	DontComeTenOdds int

	HardSix int
	HardEight int
	HardFour int
	HardTen int

	PassLine int
	PassOdds int
	DontPass int
	DontOdds int
}