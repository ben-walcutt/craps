package main

import (
	"craps/lib"
	"fmt"
	"flag"
	"math/rand"
	"time"
)

const MAX_BET  = 2;
const UNIT_AMT = 5;
const STARTING_AMT = 300;

func main() {
	numOfChildren := flag.Int("c", 20, "Number of children");
	numOfRolls := flag.Int("g", 20, "Number of rolls");
	verbose := flag.Bool("v", false, "Verbose output");

	flag.Parse();

	if *verbose {
		fmt.Println("number of children: ", *numOfChildren);
		fmt.Println("number of rolls: ", *numOfRolls);
		fmt.Println("Generating children: ");
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

	s := strategies[0];
	s.Come = true;
	g := lib.NewGame(UNIT_AMT);

	if g.ComeFour {
		fmt.Println("come four");
	}

	b := lib.Board{};
	b.PlaceBets(s, *g);

	fmt.Println(b);

}

func roll() (d1, d2 int) {
	seed := rand.NewSource(time.Now().UnixNano());
	randgen := rand.New(seed);

	d1 = randgen.Intn(6) + 1;
	time.Sleep(1);
	d2 = randgen.Intn(6) + 1;

	return d1, d2;
}

func determinePayout(s lib.Strategy, g lib.Game) {

}

func updateGame(g lib.Game) {

}

