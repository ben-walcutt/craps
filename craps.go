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

func main() {
	numOfChildren := flag.Int("c", 20, "Number of children");
	numOfRolls := flag.Int("g", 20, "Number of rolls");

	flag.Parse();

	fmt.Println("number of children: ", *numOfChildren);
	fmt.Println("number of rolls: ", *numOfRolls);

	strategies := make([]lib.Strategy, *numOfChildren);
	
	for i:=0; i < *numOfChildren; i++ {
		time.Sleep(1);
		code := generateStrategyCode();
		strategies[i] = *lib.BuildStrategy(code);
		fmt.Println(strategies[i].Encode());
	}

	

}

func generateStrategyCode() [41]int {
	seed := rand.NewSource(time.Now().UnixNano());
	randgen := rand.New(seed);

	i := [41]int{
		randgen.Intn(2),
		randgen.Intn(MAX_BET),
		randgen.Intn(4),
		randgen.Intn(5),
		randgen.Intn(6),
		randgen.Intn(6),
		randgen.Intn(5),
		randgen.Intn(4),

		randgen.Intn(2),
		randgen.Intn(MAX_BET),
		randgen.Intn(2),
		randgen.Intn(6),
		randgen.Intn(2),
		randgen.Intn(6),
		randgen.Intn(2),
		randgen.Intn(6),
		randgen.Intn(2),
		randgen.Intn(6),
		randgen.Intn(2),
		randgen.Intn(6),
		randgen.Intn(2),
		randgen.Intn(6),

		randgen.Intn(2),
		randgen.Intn(MAX_BET),
		randgen.Intn(2),
		randgen.Intn(MAX_BET),
		randgen.Intn(2),
		randgen.Intn(MAX_BET),
		randgen.Intn(2),
		randgen.Intn(MAX_BET),
		randgen.Intn(2),
		randgen.Intn(MAX_BET),
		randgen.Intn(2),
		randgen.Intn(MAX_BET),

		randgen.Intn(2),

		randgen.Intn(MAX_BET),
		randgen.Intn(4),

		randgen.Intn(2),
		randgen.Intn(2),
		randgen.Intn(2),
		randgen.Intn(2),
	}

	return i;
}

func roll() int {
	seed := rand.NewSource(time.Now().UnixNano());
	randgen := rand.New(seed);

	d1 := randgen.Intn(6) + 1;
	d2 := randgen.Intn(6) + 1;

	return d1 + d2;
}

