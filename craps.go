package main

import (
	"craps/lib"
	"fmt"
	"flag"
	"time"
	"strings"
	"strconv"
	"math"
)

var verboseOutput bool = false;

const NUM_OF_PARAMS = 47;
const MAX_BET  = 2;
const PAYOUT_OFFSET = 1.2;
// blank strategy -> [46]int {0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0};


func main() {
	namedStrategy := flag.String("n", "", "Name of Strategy (children and iterations will be 1)");
	numOfChildren := flag.Int("c", 20, "Number of children");
	numOfRolls := flag.Int("r", 20, "Number of rolls");
	numOfIterations := flag.Int("i", 1, "Number of iterations");
	isTest := flag.Bool("t", false, "Run Test Strategy");
	verbose := flag.Bool("v", false, "Verbose output");
	amount := flag.Int("a", 300, "Starting amount");
	unit_amt := flag.Int("u", 5, "Table Minimum");
	manual := flag.Bool("m", false, "Manual Roll");
	triplefield := flag.Bool("f", false, "Triple 12 on field");

	flag.Parse();

	if *verbose {
		if *namedStrategy != "" {
			fmt.Println("name of strategy:     ", *namedStrategy);
			*numOfChildren = 1;
		}
		if *isTest {
			fmt.Println("is test strategy:     ", *isTest);
			*numOfChildren = 1;
			*numOfIterations = 1;
		}
		fmt.Println("starting amount:      ", *amount);
		fmt.Println("number of children:   ", *numOfChildren);
		fmt.Println("number of rolls:      ", *numOfRolls);
		fmt.Println("number of iterations: ", *numOfIterations);
		fmt.Println("");
		fmt.Println("generating children: ");
		verboseOutput = true;
	} else {
		verboseOutput = false;

		if *namedStrategy != "" {
			*numOfChildren = 1;
		}
		if *isTest {
			*numOfChildren = 1;
			*numOfIterations = 1;
			verboseOutput = true;
		}
	}

	var STARTING_AMT = *amount;

	strategies := make([]*lib.Strategy, *numOfChildren);

	if *isTest {
		fmt.Println("Using test strategy");
		testCase := [NUM_OF_PARAMS]int {0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 2,1,0,0,0, 0,0,0,0,0, 0};
		strategies[0] = lib.BuildStrategy(testCase);
		strategies[0].Amount = STARTING_AMT;
		strategies[0].Name = "Test Strategy";
		fmt.Println("");
	} else if *namedStrategy == "Field" {
		fmt.Println("Using Field strategy");
		testCase := [NUM_OF_PARAMS]int {0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,1, 0,0,0,0,0, 0,0,0,0,0, 0};
		strategies[0] = lib.BuildStrategy(testCase);
		strategies[0].Amount = STARTING_AMT;
		strategies[0].Name = "Field Only";
		fmt.Println("");
	} else if *namedStrategy == "Iron" {
		fmt.Println("Using Iron Cross");
		testCase := [NUM_OF_PARAMS]int {0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,1, 1,1,1,1,1, 0,0,0,0,1, 1,3,0,0,0, 0,0,0,0,0, 0};
		strategies[0] = lib.BuildStrategy(testCase);
		strategies[0].Amount = STARTING_AMT;
		strategies[0].Name = "Iron Cross";
		fmt.Println("");
	} else if *namedStrategy == "Come" {
		fmt.Println("Using Come strategy");
		testCase := [NUM_OF_PARAMS]int {1,1,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0};
		strategies[0] = lib.BuildStrategy(testCase);
		strategies[0].Amount = STARTING_AMT;
		strategies[0].Name = "Come Only";
		fmt.Println("");
	} else if *namedStrategy == "22" {
		fmt.Println("Using 22 Inside strategy");
		testCase := [NUM_OF_PARAMS]int {0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,1, 1,1,1,1,1, 1,1,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0};
		strategies[0] = lib.BuildStrategy(testCase);
		strategies[0].Amount = STARTING_AMT;
		strategies[0].Name = "22 Inside";
		fmt.Println("");
	} else if *namedStrategy == "20" {
		fmt.Println("Using 20 Outside strategy");
		testCase := [NUM_OF_PARAMS]int {0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,1,1,1, 1,0,0,0,0, 1,1,1,1,0, 0,0,0,0,0, 0,0,0,0,0, 0};
		strategies[0] = lib.BuildStrategy(testCase);
		strategies[0].Amount = STARTING_AMT;
		strategies[0].Name = "20 Outside";
		fmt.Println("");
	} else if *namedStrategy == "32" {
		fmt.Println("Using 32 Across strategy");
		testCase := [NUM_OF_PARAMS]int {0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,1,1,1, 1,1,1,1,1, 1,1,1,1,0, 0,0,0,0,0, 0,0,0,0,0, 0};
		strategies[0] = lib.BuildStrategy(testCase);
		strategies[0].Amount = STARTING_AMT;
		strategies[0].Name = "32 Across";
		fmt.Println("");
	} else if *namedStrategy == "Pass" {
		fmt.Println("Using Pass Only");
		testCase := [NUM_OF_PARAMS]int {0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 1,1,0,0,0, 0,0,0,0,0, 0};
		strategies[0] = lib.BuildStrategy(testCase);
		strategies[0].Amount = STARTING_AMT;
		strategies[0].Name = "Pass Only";
	} else if *namedStrategy == "SixEight" {
		fmt.Println("Using Six Eight");
		testCase := [NUM_OF_PARAMS]int {0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,1,1,1,1, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0};
		strategies[0] = lib.BuildStrategy(testCase);
		strategies[0].Amount = STARTING_AMT;
		strategies[0].Name = "Six Eight";
	} else if *namedStrategy == "SixEightHard" {
		fmt.Println("Using Six Eight With Hardways");
		testCase := [NUM_OF_PARAMS]int {0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,1,1,1,1, 0,0,0,0,0, 0,0,1,1,0, 0,0,0,0,0, 0};
		strategies[0] = lib.BuildStrategy(testCase);
		strategies[0].Amount = STARTING_AMT;
		strategies[0].Name = "Six Eight With Hardways";
	} else if *namedStrategy == "Dark" {
		fmt.Println("Using Dark Side");
		testCase := [NUM_OF_PARAMS]int {0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 2,1,0,0,0, 0,0,0,0,0, 0};
		strategies[0] = lib.BuildStrategy(testCase);
		strategies[0].Amount = STARTING_AMT;
		strategies[0].Name = "Dark Side";
	} else if *namedStrategy == "DarkPlace" {
		fmt.Println("Using Dark Side With Place Bets");
		testCase := [NUM_OF_PARAMS]int {0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,1,1,1,1, 0,0,0,0,0, 2,1,0,0,0, 0,0,0,0,0, 0};
		strategies[0] = lib.BuildStrategy(testCase);
		strategies[0].Amount = STARTING_AMT;
		strategies[0].Name = "Dark Side";
	} else {
		fmt.Println("Algorithmically building best strategy");
		for i:=0; i < *numOfChildren; i++ {
			time.Sleep(1);
			code := lib.GenerateStrategyCode(MAX_BET);
			strategies[i] = lib.BuildStrategy(code);
			strategies[i].Amount = STARTING_AMT;
			if *verbose {
				fmt.Println(strategies[i].Encode());
			}
		}
	}

	fmt.Println("");

	for i:=0; i < *numOfIterations; i++ {

		fmt.Println("Running iteration: ", i);

		for j:=0; j < *numOfChildren; j++ {
			fmt.Println("Using strategy: ", j, " ", strategies[j].Name);

			s := strategies[j];
			fmt.Println(strategies[j].Encode());
			runStrategy(s, *numOfRolls, *unit_amt, *manual, *triplefield);
		}

		fmt.Println("\nResults: ");

		for j:=0; j < *numOfChildren; j++ {
			fmt.Println(strategies[j].Encode());
		}

		fmt.Println("");
	}
}

func runStrategy(s *lib.Strategy, numOfRolls int, unit_amt int, manual bool, triplefield bool) {
	histogram := lib.Histogram{};
	game := lib.NewGame(unit_amt);

	minBalance := s.Amount;
	maxBalance := s.Amount;

	i := 0;

	for ; i < numOfRolls; i++ {
		board := lib.Board{};
		wager := board.PlaceBets(s, game);

		if s.Amount <= 0 {
			fmt.Println("Bankrupt after ", i, " rolls");
			break;
		}

		if verboseOutput {
			fmt.Println("");
			fmt.Println("working?			", game.Working);
			if game.Working {
				fmt.Println("point: 				", game.Point);
			}
			fmt.Println("wager: 				", wager);
		}

		var d1 int;
		var d2 int;

		if manual {
			for {
				fmt.Println("\nRoll: ");
				var input string;
				fmt.Scanln(&input);
				dice := strings.Split(input, ",");
				d1, _ = strconv.Atoi(dice[0]);
				d2, _ = strconv.Atoi(dice[1]);

				if d1 > 6 || d2 > 6 {
					fmt.Println("bad input");
				} else {
					break;
				}
			}
		} else {
			d1, d2 = game.Roll();
		}

		game.Die1 = d1;
		game.Die2 = d2;

		diceTotal := d1 + d2;
		switch diceTotal {
		case 2:
			histogram.Two++;
		case 3:
			histogram.Three++;
		case 4:
			histogram.Four++;
		case 5:
			histogram.Five++;
		case 6:
			histogram.Six++;
		case 7:
			histogram.Seven++;
		case 8:
			histogram.Eight++;
		case 9:
			histogram.Nine++;
		case 10:
			histogram.Ten++;
		case 11:
			histogram.Eleven++;
		case 12:
			histogram.Twelve++;
		}

		var payout = game.DeterminePayout(board, triplefield);
		s.Amount += payout;
		
		if verboseOutput {
			fmt.Println("");
			fmt.Println("d1:     			", d1);
			fmt.Println("d2:     			", d2);
			fmt.Println("payout: 			", payout);
			fmt.Println("current balance: 	", s.Amount);
			fmt.Println("current game: 		", game);
		}

		game = game.UpdateGame(*s, verboseOutput);

		if verboseOutput {
			fmt.Println("------------");
		}

		if s.Amount < minBalance {
			minBalance = s.Amount;
		}

		if s.Amount > maxBalance {
			maxBalance = s.Amount;
		}
	}

	fmt.Println("Strategy minimum balance: ", minBalance);
	fmt.Println("Strategy maximum balance: ", maxBalance);
	if verboseOutput {
		fmt.Println("Two occurance:		", histogram.Two, "->", math.Floor(float64(histogram.Two) / float64(i) * 10000) / 100, "%");
		fmt.Println("Three occurance:	", histogram.Three, "->", math.Floor(float64(histogram.Three) / float64(i) * 10000) / 100, "%");
		fmt.Println("Four occurance:		", histogram.Four, "->", math.Floor(float64(histogram.Four) / float64(i) * 10000) / 100, "%");
		fmt.Println("Five occurance:		", histogram.Five, "->", math.Floor(float64(histogram.Five) / float64(i) * 10000) / 100, "%");
		fmt.Println("Six occurance:		", histogram.Six, "->", math.Floor(float64(histogram.Six) / float64(i) * 10000) / 100, "%");
		fmt.Println("Seven occurance:	", histogram.Seven, "->", math.Floor(float64(histogram.Seven) / float64(i) * 10000) / 100, "%");
		fmt.Println("Eight occurance:	", histogram.Eight, "->", math.Floor(float64(histogram.Eight) / float64(i) * 10000) / 100, "%");
		fmt.Println("Nine occurance:		", histogram.Nine, "->", math.Floor(float64(histogram.Nine) / float64(i) * 10000) / 100, "%");
		fmt.Println("Ten occurance:		", histogram.Ten, "->", math.Floor(float64(histogram.Ten) / float64(i) * 10000) / 100, "%");
		fmt.Println("Eleven occurance:	", histogram.Eleven, "->", math.Floor(float64(histogram.Eleven) / float64(i) * 10000) / 100, "%");
		fmt.Println("Twelve occurance:	", histogram.Twelve, "->", math.Floor(float64(histogram.Twelve) / float64(i) * 10000) / 100, "%");
	}
}

