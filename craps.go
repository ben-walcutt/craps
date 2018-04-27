package main

import (
	"craps/lib"
	"fmt"
	"flag"
	"time"
	"math/rand"
)

var verboseOutput bool = false;

const MAX_BET  = 2;
const UNIT_AMT = 5;
const STARTING_AMT = 300;
const PAYOUT_OFFSET = 1.2;

func main() {
	numOfChildren := flag.Int("c", 20, "Number of children");
	numOfRolls := flag.Int("r", 20, "Number of rolls");
	numOfIterations := flag.Int("i", 1000, "Number of iterations");
	verbose := flag.Bool("v", false, "Verbose output");

	flag.Parse();

	if *verbose {
		fmt.Println("number of children: ", *numOfChildren);
		fmt.Println("number of rolls: ", *numOfRolls);
		fmt.Println("number of iterations: ", *numOfIterations);
		fmt.Println("");
		fmt.Println("generating children: ");
		verboseOutput = true;
	}

	strategies := make([]*lib.Strategy, *numOfChildren);
	
	for i:=0; i < *numOfChildren; i++ {
		time.Sleep(1);
		code := lib.GenerateStrategyCode(MAX_BET);
		strategies[i] = lib.BuildStrategy(code);
		strategies[i].Amount = STARTING_AMT;
		if *verbose {
			fmt.Println(strategies[i].Encode());
		}
	}

	testCase := [42]int {0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,1, 1,1,1,1,1, 1,1,0,0,0, 0,0,0,0,0, 0};
	strategies[0] = lib.BuildStrategy(testCase);
	strategies[0].Amount = STARTING_AMT;

	for i:=0; i < *numOfIterations; i++ {

		for j:=0; j < *numOfChildren; j++ {

			s := strategies[j];
			fmt.Println(strategies[i].Encode());
			runStrategy(s, *numOfRolls);
		}
	}

	for i:=0; i < *numOfChildren; i++ {
		fmt.Println(strategies[i].Encode());
	}
}

func runStrategy(s *lib.Strategy, numOfRolls int) {
	game := lib.NewGame(UNIT_AMT);

	for i:=0; i < numOfRolls; i++ {
		board := lib.Board{};
		wager := board.PlaceBets(s, game);

		d1, d2 := roll();
		game.Die1 = d1;
		game.Die2 = d2;

		if verboseOutput {
			fmt.Println("current game: ", game);
			fmt.Println("wager: ", wager);
		}

		var payout = determinePayout(game, board);
		s.Amount += payout;
		if verboseOutput {
			fmt.Println("current balance: ", s.Amount);
		}

		game = updateGame(game, *s);
	}
}

func roll() (d1, d2 int) {
	seed := rand.NewSource(time.Now().UnixNano());
	randgen := rand.New(seed);

	d1 = randgen.Intn(6) + 1;
	time.Sleep(1);
	d2 = randgen.Intn(6) + 1;

	return d1, d2;
}

func determinePayout(g lib.Game, b lib.Board) int {
	payout := 0;

	if g.Die1 == 3 && g.Die2 == 3 {
		payout += b.HardSix * 9;
	}

	if g.Die1 == 4 && g.Die2 == 4 {
		payout += b.HardEight * 9;
	}

	if g.Die1 == 2 && g.Die2 == 2 {
		payout += b.HardFour * 7;
	}

	if g.Die1 == 5 && g.Die2 == 5 {
		payout += b.HardTen * 7;
	}

	diceTotal := g.Die1 + g.Die2;

	if g.Working {
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
		case 11:
			payout += b.Come;
			payout -= b.DontCome;
		case 12:
			payout -= b.Come;
		}
	} else {
		switch diceTotal {
		case 4:
		case 5:
		case 6:
		case 8:
		case 9:
		case 10:
			payout += 0;
		case 2:
		case 3:
			payout -= b.PassLine;
			payout += b.DontPass;
		case 7:
		case 11:
			payout += b.PassLine;
			payout -= b.DontPass;
		case 12:
			payout -= b.PassLine;
		}
	}

	switch diceTotal {
	case 2:
	case 12:
		payout += b.Field  * 2;
	case 3:
	case 4:
	case 9:
	case 10:
	case 11:
		payout += b.Field;
	}

	if verboseOutput {
		fmt.Println("payout: ", payout);
	}
	return payout;
}

func updateGame(g lib.Game, s lib.Strategy) lib.Game {
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
				g = lib.NewGame(UNIT_AMT);
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
				g = lib.NewGame(UNIT_AMT);
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
				g = lib.NewGame(UNIT_AMT);
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
			g = lib.NewGame(UNIT_AMT);
			if verboseOutput {
				fmt.Println("seven out");
			}
		case 8:
			if g.Point == 8 {
				g = lib.NewGame(UNIT_AMT);
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
				g = lib.NewGame(UNIT_AMT);
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
				g = lib.NewGame(UNIT_AMT);
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

	return g;
}

