package lib

// import "fmt"

const PAYOUT_OFFSET = 1.2;
const SIX_EIGHT_MAX_ODDS = 3;
const FIVE_NINE_MAX_ODDS = 4;
const FOUR_TEN_MAX_ODDS = 5;

type Board struct {
	PlaceFour int
	PlaceFive int
	PlaceSix int
	PlaceEight int
	PlaceNine int
	PlaceTen int

	Come int
	ComeFour int
	ComeFourOdds int
	ComeFive int
	ComeFiveOdds int
	ComeSix int
	ComeSixOdds int
	ComeEight int
	ComeEightOdds int
	ComeNine int
	ComeNineOdds int
	ComeTen int
	ComeTenOdds int

	DontCome int
	DontComeFour int
	DontComeFourOdds int
	DontComeFive int
	DontComeFiveOdds int
	DontComeSix int
	DontComeSixOdds int
	DontComeEight int
	DontComeEightOdds int
	DontComeNine int
	DontComeNineOdds int
	DontComeTen int
	DontComeTenOdds int

	HardSix int
	HardEight int
	HardFour int
	HardTen int

	HornTwo int
	HornThree int
	HornEleven int
	HornTwelve int

	PassLine int
	PassOdds int
	DontPass int
	DontOdds int

	Crap int
	Eleven int
	BigRed int

	Field int
}

func (b *Board) PlaceBets(s *Strategy, g Game) int {

	if g.Working {

		// place bets
		b.PlaceFour = s.PlaceFourAmt * g.Unit;
		b.PlaceFive = s.PlaceFiveAmt * g.Unit;
		b.PlaceSix = int(float64(s.PlaceSixAmt * g.Unit) * PAYOUT_OFFSET);
		b.PlaceEight = int(float64(s.PlaceEightAmt * g.Unit) * PAYOUT_OFFSET);
		b.PlaceNine = s.PlaceNineAmt * g.Unit;
		b.PlaceTen = s.PlaceTenAmt * g.Unit;

		// main come bet
		if s.Come {
			b.Come = s.ComeAmt * g.Unit;
		}

		// dont come bets
		switch g.DontCome {
		case 4:
			b.DontComeFour = s.DontComeAmt * g.Unit;
			b.DontComeFourOdds = s.DontComeFourOdds * g.Unit * 6;
		case 5:
			b.DontComeFive = s.DontComeAmt * g.Unit;
			b.DontComeFiveOdds = s.DontComeFiveOdds * g.Unit * 3;
		case 6:
			b.DontComeSix = s.DontComeAmt * g.Unit;
			b.DontComeSixOdds = int(float64(s.DontComeSixOdds * g.Unit * 3) * PAYOUT_OFFSET);
		case 8:
			b.DontComeEight = s.DontComeAmt * g.Unit;
			b.DontComeEightOdds = int(float64(s.DontComeEightOdds * g.Unit * 3) * PAYOUT_OFFSET);
		case 9:
			b.DontComeNine = s.DontComeAmt * g.Unit;
			b.DontComeNineOdds = s.DontComeNineOdds * g.Unit * 3;
		case 10:
			b.DontComeTen = s.DontComeAmt * g.Unit;
			b.DontComeTenOdds = s.DontComeTenOdds * g.Unit * 6;
		}

		// hard ways
		b.HardSix = s.HardSix;
		b.HardEight = s.HardEight;
		b.HardFour = s.HardFour;
		b.HardTen = s.HardTen;
	}

	// come bets
	if (g.ComeFour) {
		b.ComeFour = s.ComeAmt * g.Unit;
		b.ComeFourOdds = s.ComeFourOdds * g.Unit * FOUR_TEN_MAX_ODDS;
	}
	if (g.ComeFive) {
		b.ComeFive = s.ComeAmt * g.Unit;
		b.ComeFiveOdds = int(float64(s.ComeFiveOdds * g.Unit * FIVE_NINE_MAX_ODDS) * PAYOUT_OFFSET);
	}
	if (g.ComeSix) {
		b.ComeSix = s.ComeAmt * g.Unit;
		b.ComeSixOdds = s.ComeSixOdds * g.Unit * SIX_EIGHT_MAX_ODDS;
	}
	if (g.ComeEight) {
		b.ComeEight = s.ComeAmt * g.Unit;
		b.ComeEightOdds = s.ComeEightOdds * g.Unit * SIX_EIGHT_MAX_ODDS;
	}
	if (g.ComeNine) {
		b.ComeNine = s.ComeAmt * g.Unit;
		b.ComeNineOdds = int(float64(s.ComeNineOdds * g.Unit * FIVE_NINE_MAX_ODDS) * PAYOUT_OFFSET);
	}
	if (g.ComeTen) {
		b.ComeTen = s.ComeAmt * g.Unit;
		b.ComeTenOdds = s.ComeTenOdds * g.Unit * FOUR_TEN_MAX_ODDS;
	}

	// field bet
	if (s.Field) {
		b.Field = g.Unit;
	}

	// horn bets
	if (g.HornOn == 1) {
		b.HornTwo = s.HornTwo;
		b.HornThree = s.HornThree;
		b.HornEleven = s.HornEleven;
		b.HornTwelve = s.HornTwelve;
	}

	// line bets
	switch s.Line {
		case 1:
			b.PassLine = g.Unit;
			if (g.Working) {
				switch g.Point {
				case 4:
					b.PassOdds = s.LineOdds * g.Unit * 100;
				case 5:
					b.PassOdds = int(float64(s.LineOdds * g.Unit * 100) * PAYOUT_OFFSET);
				case 6:
					b.PassOdds = s.LineOdds * g.Unit * 100;
				case 8:
					b.PassOdds = s.LineOdds * g.Unit * 100;
				case 9:
					b.PassOdds = int(float64(s.LineOdds * g.Unit * 100) * PAYOUT_OFFSET);
				case 10:
					b.PassOdds = s.LineOdds * g.Unit * 100;
				}
			} else if (s.CrapCheck) {
				b.Crap = 1;
			}
		case 2:
			b.DontPass = g.Unit;
			if (g.Working) {
				switch g.Point {
				case 4:
					b.DontOdds = s.LineOdds * g.Unit * 6;
				case 5:
					b.DontOdds = s.LineOdds * g.Unit * 3;
				case 6:
					b.DontOdds = int(float64(s.LineOdds * g.Unit * 3) * PAYOUT_OFFSET);
				case 8:
					b.DontOdds = int(float64(s.LineOdds * g.Unit * 3) * PAYOUT_OFFSET);
				case 9:
					b.DontOdds = s.LineOdds * g.Unit * 3;
				case 10:
					b.DontOdds = s.LineOdds * g.Unit * 6;
				}
			} else if (s.CrapCheck) {
				b.Eleven = 1;
				b.BigRed = 2;
			}
		case 3:
	}

	valid, wager := b.validateWager(s.Amount);

	if !valid {
		s.Amount = -1;
	}

	return wager;	
}

func (b *Board) validateWager(amount int) (valid bool, wager int) {
	wager = 0;

	wager += b.PlaceFour;
	wager += b.PlaceFive;
	wager += b.PlaceSix;
	wager += b.PlaceEight;
	wager += b.PlaceNine;
	wager += b.PlaceTen;

	wager += b.Come;
	wager += b.ComeFour;
	wager += b.ComeFourOdds;
	wager += b.ComeFive;
	wager += b.ComeFiveOdds;
	wager += b.ComeSix;
	wager += b.ComeSixOdds;
	wager += b.ComeEight;
	wager += b.ComeEightOdds;
	wager += b.ComeNine;
	wager += b.ComeNineOdds;
	wager += b.ComeTen;
	wager += b.ComeTenOdds;

	wager += b.DontComeFour;
	wager += b.DontComeFourOdds;
	wager += b.DontComeFive;
	wager += b.DontComeFiveOdds;
	wager += b.DontComeSix;
	wager += b.DontComeSixOdds;
	wager += b.DontComeEight;
	wager += b.DontComeEightOdds;
	wager += b.DontComeNine;
	wager += b.DontComeNineOdds;
	wager += b.DontComeTen;
	wager += b.DontComeTenOdds;

	wager += b.HardSix;
	wager += b.HardEight;
	wager += b.HardFour;
	wager += b.HardTen;

	wager += b.HornTwo;
	wager += b.HornThree;
	wager += b.HornEleven;
	wager += b.HornTwelve;

	wager += b.PassLine;
	wager += b.PassOdds;
	wager += b.DontPass;
	wager += b.DontOdds;

	wager += b.Field;

	wager += b.Crap;
	wager += b.Eleven;
	wager += b.BigRed;

	return wager <= amount, wager;
}