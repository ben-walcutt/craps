package main

import (
	"craps/lib"
	"fmt"
	"flag"
	"time"
)

var verboseOutput bool = false;

const MAX_BET  = 2;
const UNIT_AMT = 5;
const PAYOUT_OFFSET = 1.2;

func main() {
	namedStrategy := flag.String("n", "", "Name of Strategy (children and iterations will be 1)");
	numOfChildren := flag.Int("c", 20, "Number of children");
	numOfRolls := flag.Int("r", 20, "Number of rolls");
	numOfIterations := flag.Int("i", 1000, "Number of iterations");
	isTest := flag.Bool("t", false, "Run Test Strategy");
	verbose := flag.Bool("v", false, "Verbose output");
	amount := flag.Int("a", 300, "Starting amount");

	flag.Parse();

	if *verbose {
		if *namedStrategy != "" {
			fmt.Println("name of strategy:     ", *namedStrategy);
			*numOfChildren = 1;
			*numOfIterations = 1;
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
	}

	var STARTING_AMT = *amount;

	strategies := make([]*lib.Strategy, *numOfChildren);

	if *isTest {
		fmt.Println("Using test strategy");
		testCase := [46]int {0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,1,1,1,1};
		strategies[0] = lib.BuildStrategy(testCase);
		strategies[0].Amount = STARTING_AMT;
		strategies[0].Name = "Test Strategy";
		fmt.Println("");
	} else if *namedStrategy == "Field" {
		fmt.Println("Using Field strategy");
		testCase := [46]int {0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,1, 0,0,0,0,0, 0,0,0,0,0};
		strategies[0] = lib.BuildStrategy(testCase);
		strategies[0].Amount = STARTING_AMT;
		strategies[0].Name = "Field Only";
		fmt.Println("");
	} else if *namedStrategy == "Iron Cross" {
		fmt.Println("Using Iron Cross");
		testCase := [46]int {0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,1, 1,1,1,1,1, 0,0,0,0,1, 1,3,0,0,0, 0,0,0,0,0};
		strategies[0] = lib.BuildStrategy(testCase);
		strategies[0].Amount = STARTING_AMT;
		strategies[0].Name = "Iron Cross";
		fmt.Println("");
	} else if *namedStrategy == "Come" {
		fmt.Println("Using Come strategy");
		testCase := [46]int {1,1,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0, 0,0,0,0,0};
		strategies[0] = lib.BuildStrategy(testCase);
		strategies[0].Amount = STARTING_AMT;
		strategies[0].Name = "Come Only";
		fmt.Println("");
	} else {
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
			runStrategy(s, *numOfRolls);
		}

		fmt.Println("");
		fmt.Println("Results: ");

		for i:=0; i < *numOfChildren; i++ {
			fmt.Println(strategies[i].Encode());
		}

		fmt.Println("");
	}
}

func runStrategy(s *lib.Strategy, numOfRolls int) {
	game := lib.NewGame(UNIT_AMT);

	minBalance := s.Amount;
	maxBalance := s.Amount;

	for i:=0; i < numOfRolls; i++ {
		board := lib.Board{};
		wager := board.PlaceBets(s, game);

		if s.Amount <= 0 {
			fmt.Println("Bankrupt after ", i, " rolls");
			break;
		}

		d1, d2 := game.Roll();
		game.Die1 = d1;
		game.Die2 = d2;

		var payout = game.DeterminePayout(board);
		s.Amount += payout;
		
		if verboseOutput {
			fmt.Println("payout: ", payout);
			fmt.Println("wager: ", wager);
			fmt.Println("current balance: ", s.Amount);
			fmt.Println("current game: ", game);
		}

		game = game.UpdateGame(*s, verboseOutput);

		fmt.Println("");

		if s.Amount < minBalance {
			minBalance = s.Amount;
		}

		if s.Amount > maxBalance {
			maxBalance = s.Amount;
		}
	}

	fmt.Println("Strategy minimum balance: ", minBalance);
	fmt.Println("Strategy maximum balance: ", maxBalance);
}

