package lib

type Game struct {
	ComeFour bool
	ComeFive bool
	ComeSix bool
	ComeEight bool
	ComeNine bool
	ComeTen bool

	DontCome int
	Working bool

	Unit int

	HornOn int

	Point int

	Die1 int
	Die2 int
}

func NewGame (unit int) Game {
	g := Game {
		Unit: unit,
	}

	return g;
}