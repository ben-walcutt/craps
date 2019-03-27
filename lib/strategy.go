package lib

import (
	"math/rand"
	"time"
)

const NUM_OF_PARAMS = 42;

type Strategy struct {
	Come bool
	ComeAmt int
	ComeFourOdds int
	ComeFiveOdds int
	ComeSixOdds int

	ComeEightOdds int
	ComeNineOdds int
	ComeTenOdds int
	DontCome bool
	DontComeAmt int

	DontComeFour bool
	DontComeFourOdds int
	DontComeFive bool
	DontComeFiveOdds int
	DontComeSix bool

	DontComeSixOdds int
	DontComeEight bool
	DontComeEightOdds int
	DontComeNine bool
	DontComeNineOdds int

	DontComeTen bool
	DontComeTenOdds int
	PlaceFour bool
	PlaceFourAmt int
	PlaceFive bool

	PlaceFiveAmt int
	PlaceSix bool
	PlaceSixAmt int
	PlaceEight bool
	PlaceEightAmt int

	PlaceNine bool
	PlaceNineAmt int
	PlaceTen bool
	PlaceTenAmt int
	Field bool

	Line int
	LineOdds int
	// TODO: add line odds for each point number
	HardSix bool
	HardEight bool
	HardFour bool

	HardTen bool
	Amount int
	Name string
}

func BuildStrategy(code [NUM_OF_PARAMS]int) *Strategy {
	s := &Strategy{
		Come: ItoB(code[0]),
		ComeAmt: code[1],
		ComeFourOdds: code[2],
		ComeFiveOdds: code[3],
		ComeSixOdds: code[4],
		ComeEightOdds: code[5],
		ComeNineOdds: code[6],
		ComeTenOdds: code[7],
		DontCome: ItoB(code[8]),
		DontComeAmt: code[9],
		DontComeFour: ItoB(code[10]),
		DontComeFourOdds: code[11],
		DontComeFive: ItoB(code[12]),
		DontComeFiveOdds: code[13],
		DontComeSix: ItoB(code[14]),
		DontComeSixOdds: code[15],
		DontComeEight: ItoB(code[16]),
		DontComeEightOdds: code[17],
		DontComeNine: ItoB(code[18]),
		DontComeNineOdds: code[19],
		DontComeTen: ItoB(code[20]),
		DontComeTenOdds: code[21],
		PlaceFour: ItoB(code[22]),
		PlaceFourAmt: code[23],
		PlaceFive: ItoB(code[24]),
		PlaceFiveAmt: code[25],
		PlaceSix: ItoB(code[26]),
		PlaceSixAmt: code[27],
		PlaceEight: ItoB(code[28]),
		PlaceEightAmt: code[29],
		PlaceNine: ItoB(code[30]),
		PlaceNineAmt: code[31],
		PlaceTen: ItoB(code[32]),
		PlaceTenAmt: code[33],
		Field: ItoB(code[34]),
		Line: code[35],
		LineOdds: code[36],
		HardSix: ItoB(code[37]),
		HardEight: ItoB(code[38]),
		HardFour: ItoB(code[39]),
		HardTen: ItoB(code[40]),
	}

	s.Name = "Stanley Hudson";

	return s;
}

func CombineStrategies(s1 Strategy, s2 Strategy) *Strategy {
	s1code := s1.Encode();
	s2code := s2.Encode();

	i := [NUM_OF_PARAMS]int {
		s1code[0],
		s2code[1],
		s1code[2],
		s2code[3],
		s1code[4],
		s2code[5],
		s1code[6],
		s2code[7],
		s1code[8],
		s2code[9],
		s1code[10],
		s2code[11],
		s1code[12],
		s2code[13],
		s1code[14],
		s2code[15],
		s1code[16],
		s2code[17],
		s1code[18],
		s2code[19],
		s1code[20],
		s2code[21],
		s1code[22],
		s2code[23],
		s1code[24],
		s2code[25],
		s1code[26],
		s2code[27],
		s1code[28],
		s2code[29],
		s1code[30],
		s2code[31],
		s1code[32],
		s2code[33],
		s1code[34],
		s2code[35],
		s1code[36],
		s2code[37],
		s1code[38],
		s2code[39],
		s1code[40],
	}

	return BuildStrategy(i);
}

func (s Strategy) Encode() [NUM_OF_PARAMS]int {
	i := [NUM_OF_PARAMS]int{
		BtoI(s.Come),
		s.ComeAmt,
		s.ComeFourOdds,
		s.ComeFiveOdds,
		s.ComeSixOdds,
		s.ComeEightOdds,
		s.ComeNineOdds,
		s.ComeTenOdds,
		BtoI(s.DontCome),
		s.DontComeAmt,
		BtoI(s.DontComeFour),
		s.DontComeFourOdds,
		BtoI(s.DontComeFive),
		s.DontComeFiveOdds,
		BtoI(s.DontComeSix),
		s.DontComeSixOdds,
		BtoI(s.DontComeEight),
		s.DontComeEightOdds,
		BtoI(s.DontComeNine),
		s.DontComeNineOdds,
		BtoI(s.DontComeTen),
		s.DontComeTenOdds,
		BtoI(s.PlaceFour),
		s.PlaceFourAmt,
		BtoI(s.PlaceFive),
		s.PlaceFiveAmt,
		BtoI(s.PlaceSix),
		s.PlaceSixAmt,
		BtoI(s.PlaceEight),
		s.PlaceEightAmt,
		BtoI(s.PlaceNine),
		s.PlaceNineAmt,
		BtoI(s.PlaceTen),
		s.PlaceTenAmt,
		BtoI(s.Field),
		s.Line,
		s.LineOdds,
		BtoI(s.HardSix),
		BtoI(s.HardEight),
		BtoI(s.HardFour),
		BtoI(s.HardTen),
		s.Amount,
	}

	return i;
}

func GenerateStrategyCode(MAX_BET int) [NUM_OF_PARAMS]int {
	seed := rand.NewSource(time.Now().UnixNano());
	randgen := rand.New(seed);

	i := [NUM_OF_PARAMS]int{
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
		0,
	}

	return i;
}

func BtoI(b bool) int {
	if (b) {
		return 1;
	} else {
		return 0;
	}
}

func ItoB(i int) bool {
	if (i == 0) {
		return false;
	} else {
		return true;
	}
}