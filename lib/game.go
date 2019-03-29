package lib

import (
	"time"
	"math/rand"
	"fmt"
)

const FIELD_TWO_MULTIPLIER = 2;
const FIELD_TWELVE_MULTIPLIER = 2;
const UNIT_AMT = 5;

type Game struct {
	ComeFour bool
	ComeFive bool
	ComeSix bool
	ComeEight bool
	ComeNine bool
	ComeTen bool

	DontCome int
	Working bool

	StartingAmount int
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

func (g Game) Roll() (d1, d2 int) {
	seed := rand.NewSource(time.Now().UnixNano());
	randgen := rand.New(seed);

	d1 = randgen.Intn(6) + 1;
	time.Sleep(1);
	d2 = randgen.Intn(6) + 1;

	return d1, d2;
}

func (g Game) DeterminePayout(b Board) int {
	payout := 0;

	diceTotal := g.Die1 + g.Die2;

	if g.Working {
		// handling working payout
		switch diceTotal {
		case 2:
			payout -= b.Come;
			payout += b.DontCome;
		case 3:
			payout -= b.Come;
			payout += b.DontCome;
		case 4:
			payout += b.ComeFour;
			payout += b.ComeFourOdds * 2;
			payout += b.PlaceFour / 5 * 9
			payout -= b.DontComeFour;
			payout -= b.DontComeFourOdds;

			if g.Point == 4 {
				payout += b.PassLine;
				payout += b.PassOdds  * 2;
				payout -= b.DontPass;
				payout -= b.DontOdds;
			}

			if (b.HardFour > 0) {
				if g.Die1 == 2 && g.Die2 == 2 {
					payout += b.HardFour * 7;
				} else {
					payout -= b.HardFour;
				}
			}
		case 5:
			payout += b.ComeFive;
			payout += b.ComeFiveOdds / 2 * 3;
			payout += b.PlaceFive / 5 * 7;
			payout -= b.DontComeFive;
			payout -= b.DontComeFiveOdds;

			if g.Point == 5 {
				payout += b.PassLine;
				payout += b.PassOdds / 2 * 3;
				payout -= b.DontPass;
				payout -= b.DontOdds;
			}
		case 6:
			payout += b.ComeSix;
			payout += b.ComeSixOdds / 5 * 6;
			payout += b.PlaceSix / 6 * 7;
			payout -= b.DontComeSix;
			payout -= b.DontComeSixOdds;

			if g.Point == 6 {
				payout += b.PassLine;
				payout += b.PassOdds / 5 * 6;
				payout -= b.DontPass;
				payout -= b.DontOdds;
			}
			if (b.HardSix > 0) {
				if g.Die1 == 3 && g.Die2 == 3 {
					payout += b.HardSix * 9;
				} else {
					payout -= b.HardSix;
				}
			}
		case 7:
			payout += b.Come;
			payout -= b.DontCome;
			payout -= b.ComeFour;
			payout -= b.ComeFourOdds;
			payout -= b.PlaceFour;
			payout -= b.ComeFive;
			payout -= b.ComeFiveOdds;
			payout -= b.PlaceFive;
			payout -= b.ComeSix;
			payout -= b.ComeSixOdds;
			payout -= b.PlaceSix;
			payout -= b.ComeEight;
			payout -= b.ComeEightOdds;
			payout -= b.PlaceEight;
			payout -= b.ComeNine;
			payout -= b.ComeNineOdds;
			payout -= b.PlaceNine;
			payout -= b.ComeTen;
			payout -= b.ComeTenOdds;
			payout -= b.PlaceTen;

			payout += b.DontComeFour;
			payout += b.DontComeFourOdds / 2;
			payout += b.DontComeFive;
			payout += b.DontComeFiveOdds / 3 * 2;
			payout += b.DontComeSix;
			payout += b.DontComeSixOdds / 6 * 5;
			payout += b.DontComeEight;
			payout += b.DontComeEightOdds / 6 * 5;
			payout += b.DontComeNine;
			payout += b.DontComeNineOdds / 3 * 2;
			payout += b.DontComeTen;
			payout += b.DontComeTenOdds / 2;

			payout -= b.HardSix;
			payout -= b.HardEight;
			payout -= b.HardFour;
			payout -= b.HardTen;

			payout -= b.PassLine;
			payout -= b.PassOdds;
			payout += b.DontPass;

			if (g.Point == 4 || g.Point == 10) {
				payout += b.DontPass / 2;
			}
			if (g.Point == 5 || g.Point == 9) {
				payout += b.DontPass / 3 * 2;
			}
			if (g.Point == 6 || g.Point == 8) {
				payout += b.DontPass / 6 * 5;
			}
		case 8:
			payout += b.ComeEight;
			payout += b.ComeEightOdds / 5 * 6;
			payout += b.PlaceEight / 6 * 7;
			payout -= b.DontComeEight;
			payout -= b.DontComeEightOdds;

			if g.Point == 8 {
				payout += b.PassLine;
				payout += b.PassOdds / 5 * 6;
				payout -= b.DontPass;
				payout -= b.DontOdds;
			}
			if (b.HardEight > 0) {
				if g.Die1 == 4 && g.Die2 == 4 {
					payout += b.HardEight * 9;
				} else {
					payout -= b.HardEight;
				}
			}
		case 9:
			payout += b.ComeNine;
			payout += b.ComeNineOdds / 2 * 3;
			payout += b.PlaceNine  / 5 * 7;
			payout -= b.DontComeNine;
			payout -= b.DontComeNineOdds;

			if g.Point == 9 {
				payout += b.PassLine;
				payout += b.PassOdds / 2 * 3;
				payout -= b.DontPass;
				payout -= b.DontOdds;
			}
		case 10:
			payout += b.ComeTen;
			payout += b.ComeTenOdds  * 2;
			payout += b.PlaceTen / 5 * 9;
			payout -= b.DontComeTen;
			payout -= b.DontComeTenOdds;

			if g.Point == 10 {
				payout += b.PassLine;
				payout += b.PassOdds  * 2;
				payout -= b.DontPass;
				payout -= b.DontOdds;
			}
			if (b.HardTen > 0) {
				if g.Die1 == 5 && g.Die2 == 5 {
					payout += b.HardTen * 7;
				} else {
					payout -= b.HardTen;
				}
			}
		case 11:
			payout += b.Come;
			payout -= b.DontCome;
		case 12:
			payout -= b.Come;
		}
	} else {
		// handling non working payout
		switch diceTotal {
		case 4:
			payout += 0;
		case 5:
			payout += 0;
		case 6:
			payout += 0;
		case 8:
			payout += 0;
		case 9:
			payout += 0;
		case 10:
			payout += 0;
		case 2:
			payout -= b.PassLine;
			payout += b.DontPass;
		case 3:
			payout -= b.PassLine;
			payout += b.DontPass;
		case 7:
			payout += b.PassLine;
			payout -= b.DontPass;
		case 11:
			payout += b.PassLine;
			payout -= b.DontPass;
		case 12:
			payout -= b.PassLine;
		}
	}


	// handling field payout
	switch diceTotal {
	case 2:
		payout += b.Field  * FIELD_TWO_MULTIPLIER;
	case 12:
		payout += b.Field  * FIELD_TWELVE_MULTIPLIER;
	case 3:
		payout += b.Field;
	case 4:
		payout += b.Field;
	case 5:
		payout -= b.Field;
	case 6:
		payout -= b.Field;
	case 7:
		payout -= b.Field;
	case 8:
		payout -= b.Field;
	case 9:
		payout += b.Field;
	case 10:
		payout += b.Field;
	case 11:
		payout += b.Field;
	}

	//handling horn payout
	switch diceTotal {
	case 2:
		payout += b.HornTwo * 30;
		payout -= b.HornThree;
		payout -= b.HornEleven;
		payout -= b.HornTwelve;
	case 3:
		payout -= b.HornTwo;
		payout += b.HornThree * 15;
		payout -= b.HornEleven;
		payout -= b.HornTwelve;
	case 11:
		payout -= b.HornTwo;
		payout -= b.HornThree;
		payout += b.HornEleven * 15;
		payout -= b.HornTwelve;
	case 12:
		payout -= b.HornTwo;
		payout -= b.HornThree;
		payout -= b.HornEleven;
		payout += b.HornTwelve * 30;
	default:
		payout -= b.HornTwo;
		payout -= b.HornThree;
		payout -= b.HornEleven;
		payout -= b.HornTwelve;
	}
	return payout;
}

func (g Game) UpdateGame(s Strategy, verboseOutput bool) Game {
	diceTotal := g.Die1 + g.Die2;

	var dontComeEstablished bool;

	if g.DontCome == 0 {
		dontComeEstablished = false;
	} else {
		dontComeEstablished = true;
	}

	if g.Working {
		switch diceTotal {
		case 4:
			if g.Point == 4 {
				g = NewGame(UNIT_AMT);
				if verboseOutput {
					fmt.Println("win 4");
				}
			} else {
				g.ComeFour = true;
				if !dontComeEstablished && s.DontComeFour {
					g.DontCome = 4;
				}
			}
		case 5:
			if g.Point == 5 {
				g = NewGame(UNIT_AMT);
				if verboseOutput {
					fmt.Println("win 5");
				}
			} else {
				g.ComeFive = true;
				if !dontComeEstablished && s.DontComeFive {
					g.DontCome = 5;
				}
			}
		case 6:
			if g.Point == 6 {
				g = NewGame(UNIT_AMT);
				if verboseOutput {
					fmt.Println("win 6");
				}
			} else {
				g.ComeSix = true;
				if !dontComeEstablished && s.DontComeSix {
					g.DontCome = 6;
				}
			}
		case 7:
			g = NewGame(UNIT_AMT);
			if verboseOutput {
				fmt.Println("seven out");
			}
		case 8:
			if g.Point == 8 {
				g = NewGame(UNIT_AMT);
				if verboseOutput {
					fmt.Println("win 8");
				}
			} else {
				g.ComeEight = true;
				if !dontComeEstablished && s.DontComeEight {
					g.DontCome = 8;
				}
			}
		case 9:
			if g.Point == 9 {
				g = NewGame(UNIT_AMT);
				if verboseOutput {
					fmt.Println("win 9");
				}
			} else {
				g.ComeNine = true;
				if !dontComeEstablished && s.DontComeNine {
					g.DontCome = 9;
				}
			}
		case 10:
			if g.Point == 10 {
				g = NewGame(UNIT_AMT);
				if verboseOutput {
					fmt.Println("win 10");
				}
			} else {
				g.ComeTen = true;
				if !dontComeEstablished && s.DontComeTen {
					g.DontCome = 10;
				}
			}
		}
	} else {
		switch diceTotal {
		case 4:
			g.Point = 4;
			g.Working = true;
		case 5:
			g.Point = 5;
			g.Working = true;
		case 6:
			g.Point = 6;
			g.Working = true;
		case 8:
			g.Point = 8;
			g.Working = true;
		case 9:
			g.Point = 9;
			g.Working = true;
		case 10:
			g.Point = 10;
			g.Working = true;
		}
	}

	if (diceTotal == 2 || diceTotal == 3 || diceTotal == 11 || diceTotal == 12) {
		g.HornOn = 1;
	} else {
		g.HornOn = 0;
	}

	return g;
}