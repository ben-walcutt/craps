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
	numOfRolls := flag.Int("g", 20, "Number of rolls");
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

	strategies := make([]lib.Strategy, *numOfChildren);
	
	for i:=0; i < *numOfChildren; i++ {
		time.Sleep(1);
		code := lib.GenerateStrategyCode(MAX_BET);
		strategies[i] = *lib.BuildStrategy(code);
		if *verbose {
			fmt.Println(strategies[i].Encode());
		}
	}

	for i:=0; i < *numOfIterations; i++ {

		for j:=0; j < *numOfChildren; j++ {

			s := strategies[j];
			runStrategy(&s, *numOfRolls);
		}
	}

}

func runStrategy(s *lib.Strategy, numOfRolls int) {
	game := lib.NewGame(UNIT_AMT);

	for i:=0; i < numOfRolls; i++ {
		d1, d2 := roll();
		game.Die1 = d1;
		game.Die2 = d2;

		board := lib.Board{};
		board.PlaceBets(s, *game);

		var payout = determinePayout(game, board);
		s.Amount += payout;

		updateGame(game);
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

func determinePayout(g *lib.Game, b lib.Board) int {
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
			payout -= b.Come * g.Unit;
			payout += b.DontCome * g.Unit;
		case 3:
			payout -= b.Come * g.Unit;
			payout += b.DontCome * g.Unit;
		case 4:
			payout += b.ComeFour * g.Unit;
			payout += b.ComeFourOdds * g.Unit * 2;
			payout += b.PlaceFour * g.Unit / 5 * 9
			payout -= b.DontComeFour * g.Unit;
			payout -= b.DontComeFourOdds * g.Unit;
		case 5:
			payout += b.ComeFive * g.Unit;
			payout += int(float64(b.ComeFiveOdds * g.Unit) * PAYOUT_OFFSET) / 2 * 3;
			payout += b.PlaceFive * g.Unit / 5 * 7;
			payout -= b.DontComeFive * g.Unit;
			payout -= b.DontComeFiveOdds * g.Unit;
		case 6:
			payout += b.ComeSix * g.Unit;
			payout += b.ComeSixOdds * g.Unit / 5 * 6;
			payout += int(float64(b.PlaceSix * g.Unit) * PAYOUT_OFFSET) / 6 * 7;
			payout -= b.DontComeSix * g.Unit;
			payout -= b.DontComeSixOdds * g.Unit;
		case 8:
			payout += b.ComeEight * g.Unit;
			payout += b.ComeEightOdds * g.Unit / 5 * 6;
			payout += int(float64(b.PlaceEight * g.Unit) * PAYOUT_OFFSET) / 6 * 7;
			payout -= b.DontComeEight * g.Unit;
			payout -= b.DontComeEightOdds * g.Unit;
		case 9:
			payout += b.ComeNine * g.Unit;
			payout += int(float64(b.ComeNineOdds * g.Unit) * PAYOUT_OFFSET) / 2 * 3;
			payout += b.PlaceNine * g.Unit / 5 * 7;
			payout -= b.DontComeNine * g.Unit;
			payout -= b.DontComeNineOdds * g.Unit;
		case 10:
			payout += b.ComeTen * g.Unit;
			payout += b.ComeTenOdds * g.Unit * 2;
			payout += b.PlaceTen * g.Unit / 5 * 9;
			payout -= b.DontComeTen * g.Unit;
			payout -= b.DontComeTenOdds * g.Unit;
		case 11:
			payout += b.Come * g.Unit;
			payout -= b.DontCome * g.Unit;
		case 12:
			payout -= b.Come * g.Unit;
		case 7:
			payout += b.Come * g.Unit;
			payout -= b.DontCome * g.Unit;
			payout -= b.ComeFour * g.Unit;
			payout -= b.ComeFourOdds * g.Unit;
			payout -= b.PlaceFour * g.Unit;
			payout -= b.ComeFive * g.Unit;
			payout -= int(float64(b.ComeFiveOdds * g.Unit) * PAYOUT_OFFSET);
			payout -= b.PlaceFive * g.Unit;
			payout -= b.ComeSix * g.Unit;
			payout -= b.ComeSixOdds * g.Unit;
			payout -= int(float64(b.PlaceSix * g.Unit) * PAYOUT_OFFSET);
			payout -= b.ComeEight * g.Unit;
			payout -= b.ComeEightOdds * g.Unit;
			payout -= int(float64(b.PlaceEight * g.Unit) * PAYOUT_OFFSET);
			payout -= b.ComeNine * g.Unit;
			payout -= int(float64(b.ComeNineOdds * g.Unit) * PAYOUT_OFFSET);
			payout -= b.PlaceNine * g.Unit;
			payout -= b.ComeTen * g.Unit;
			payout -= b.ComeTenOdds * g.Unit;
			payout -= b.PlaceTen * g.Unit;

			payout += b.DontComeFour * g.Unit;
			payout += b.DontComeFourOdds * g.Unit / 2;
			payout += b.DontComeFive * g.Unit;
			payout += b.DontComeFiveOdds * g.Unit / 3 * 2;
			payout += b.DontComeSix * g.Unit;
			payout += int(float64(b.DontComeSixOdds * g.Unit) * PAYOUT_OFFSET) / 6 * 5;
			payout += b.DontComeEight * g.Unit;
			payout += int(float64(b.DontComeEightOdds * g.Unit) * PAYOUT_OFFSET) / 6 * 5;
			// TODO finish
		}
	}


}

func updateGame(g *lib.Game) {

}

