package lib

import "time"
import "math/rand"

type Challenger struct {
	Come bool
	ComeFour bool
	ComeOddsFour bool
	ComeOddsFourValue int
	ComeFive bool
	ComeOddsFive bool
	ComeOddsFiveValue int
	ComeSix bool
	ComeOddsSix bool
	ComeOddsSixValue int
	ComeEight bool
	ComeOddsEight bool
	ComeOddsEightValue int
	ComeNine bool
	ComeOddsNine bool
	ComeOddsNineValue int
	ComeTen bool
	ComeOddsTen bool
	ComeOddsTenValue int
	PlaceFour bool
	PlaceFive bool
	PlaceSix bool
	PlaceEight bool
	PlaceNine bool
	PlaceTen bool
	Pass bool
	DontPass bool
}

func EncodeGenes(c Challenger) [27]int {
	var retVal [27]int;

	retVal[0] = BtoI(c.Come);
	retVal[1] = BtoI(c.ComeFour);
	retVal[2] = BtoI(c.ComeOddsFour);
	retVal[3] = c.ComeOddsFourValue;
	retVal[4] = BtoI(c.ComeFive);
	retVal[5] = BtoI(c.ComeOddsFive);
	retVal[6] = c.ComeOddsFiveValue;
	retVal[7] = BtoI(c.ComeSix);
	retVal[8] = BtoI(c.ComeOddsSix);
	retVal[9] = c.ComeOddsSixValue;
	retVal[10] = BtoI(c.ComeEight);
	retVal[11] = BtoI(c.ComeOddsEight);
	retVal[12] = c.ComeOddsEightValue;
	retVal[13] = BtoI(c.ComeNine);
	retVal[14] = BtoI(c.ComeOddsNine);
	retVal[15] = c.ComeOddsNineValue;
	retVal[16] = BtoI(c.ComeTen);
	retVal[17] = BtoI(c.ComeOddsTen);
	retVal[18] = c.ComeOddsTenValue;
	retVal[19] = BtoI(c.PlaceFour);
	retVal[20] = BtoI(c.PlaceFive);
	retVal[21] = BtoI(c.PlaceSix);
	retVal[22] = BtoI(c.PlaceEight);
	retVal[23] = BtoI(c.PlaceNine);
	retVal[24] = BtoI(c.PlaceTen);
	retVal[25] = BtoI(c.Pass);
	retVal[26] = BtoI(c.DontPass);

	return retVal;
}

func DecodeGenes(genes [27]int) Challenger {
	c := Challenger {
		Come: (genes[0] != 0),
		ComeFour: (genes[1] != 0),
		ComeOddsFour: (genes[2] != 0),
		ComeOddsFourValue: genes[3],
		ComeFive: (genes[4] != 0),
		ComeOddsFive: (genes[5] != 0),
		ComeOddsFiveValue: genes[6],
		ComeSix: (genes[7] != 0),
		ComeOddsSix: (genes[8] != 0),
		ComeOddsSixValue: genes[9],
		ComeEight: (genes[10] != 0),
		ComeOddsEight: (genes[11] != 0),
		ComeOddsEightValue: genes[12],
		ComeNine: (genes[13] != 0),
		ComeOddsNine: (genes[14] != 0),
		ComeOddsNineValue: genes[15],
		ComeTen: (genes[16] != 0),
		ComeOddsTen: (genes[17] != 0),
		ComeOddsTenValue: genes[18],
		PlaceFour: (genes[19] != 0),
		PlaceFive: (genes[20] != 0),
		PlaceSix: (genes[21] != 0),
		PlaceEight: (genes[22] != 0),
		PlaceNine: (genes[23] != 0),
		PlaceTen: (genes[24] != 0),
		Pass: (genes[25] != 0),
		DontPass: (genes[26] != 0),
	}

	return c;
}

func NewChallenger() Challenger {
	randSource := rand.NewSource(time.Now().UnixNano());
	randGen := rand.New(randSource);

	var values [27]int;
	values[0] = randGen.Intn(2);
	values[1] = randGen.Intn(2);
	values[2] = randGen.Intn(2);
	values[3] = randGen.Intn(4);
	values[4] = randGen.Intn(2);
	values[5] = randGen.Intn(2);
	values[6] = randGen.Intn(5);
	values[7] = randGen.Intn(2);
	values[8] = randGen.Intn(2);
	values[9] = randGen.Intn(6);
	values[10] = randGen.Intn(2);
	values[11] = randGen.Intn(2);
	values[12] = randGen.Intn(6);
	values[13] = randGen.Intn(2);
	values[14] = randGen.Intn(2);
	values[15] = randGen.Intn(5);
	values[16] = randGen.Intn(2);
	values[17] = randGen.Intn(2);
	values[18] = randGen.Intn(4);
	values[19] = randGen.Intn(2);
	values[20] = randGen.Intn(2);
	values[21] = randGen.Intn(2);
	values[22] = randGen.Intn(2);
	values[23] = randGen.Intn(2);
	values[24] = randGen.Intn(2);
	values[25] = randGen.Intn(2);
	values[26] = randGen.Intn(2);

	return DecodeGenes(values);
}

type Board struct {
	Point int
	Working bool
	ComeFour bool
	ComeFive bool
	ComeSix bool
	ComeEight bool
	ComeNine bool
	ComeTen bool
	DontCome int
}

func ShowBoard(b Board) [9]int {
	var retVal [9]int;

	retVal[0] = b.Point;
	retVal[1] = BtoI(b.Working);
	retVal[2] = BtoI(b.ComeFour);
	retVal[3] = BtoI(b.ComeFive);
	retVal[4] = BtoI(b.ComeSix);
	retVal[5] = BtoI(b.ComeEight);
	retVal[6] = BtoI(b.ComeNine);
	retVal[7] = BtoI(b.ComeTen);
	retVal[8] = b.DontCome;

	return retVal;
}

func BtoI(b bool) int {
	if b {
		return 1
	}
	return 0
}