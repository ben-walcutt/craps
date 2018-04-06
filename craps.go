package main

import "fmt"
import "flag"
import "craps/lib"
import "time"

func main() {

	numOfIterations := flag.Int("i", 1000, "Number of iterations");
	numOfGames := flag.Int("g", 10, "Number of games to play");
	numOfChildren := flag.Int("c", 20, "Number of children");
	tableMin := flag.Int("m", 5, "Table minimum");
	flag.Parse();

	if (*numOfIterations == 1000) {
		msg := fmt.Sprintf("Default number of iterations being used (1000)");
		fmt.Println(msg);
	}

	if (*numOfGames == 10) {
		msg := fmt.Sprintf("Default number of games being used (10)");
		fmt.Println(msg);
	}

	if (*numOfChildren == 20) {
		msg := fmt.Sprintf("Default number of children being used (20)");
		fmt.Println(msg);
	}

	if (*tableMin == 5) {
		msg := fmt.Sprintf("Default table minimum being used ($5)");
		fmt.Println(msg);
	}

	fmt.Println("");
	fmt.Println("Generating initial challengers...");

	challengers := make([]lib.Challenger, *numOfChildren);

	for i:=0;i<*numOfChildren;i++ {
		challengers[i] = lib.NewChallenger();
		fmt.Println(lib.EncodeGenes(challengers[i]));
		time.Sleep(1);
	}

	fmt.Println("");
	fmt.Println("Generating board...");

	board := lib.Board {
		Point: 0,
		Working: false,
		ComeFour: false,
		ComeFive: false,
		ComeSix: false,
		ComeEight: false,
		ComeNine: false,
		ComeTen: false,
		DontCome: 0,
	}

	fmt.Println(lib.ShowBoard(board));

}