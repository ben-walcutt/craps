package lib

// import "fmt"

const PAYOUT_OFFSET = 1.2;

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

	PassLine int
	PassOdds int
	DontPass int
	DontOdds int

	Field int
}

func (b *Board) PlaceBets(s *Strategy, g Game) int {

	if g.Working {

		if valid, _ := b.validateWager(s.Amount); valid && s.PlaceFour {
			b.PlaceFour = s.PlaceFourAmt * g.Unit;
		}

		if valid, _ := b.validateWager(s.Amount); valid && s.PlaceFive {
			b.PlaceFive = s.PlaceFiveAmt * g.Unit;
		}

		if valid, _ := b.validateWager(s.Amount); valid && s.PlaceSix {
			b.PlaceSix = int(float64(s.PlaceSixAmt * g.Unit) * PAYOUT_OFFSET);
		}

		if valid, _ := b.validateWager(s.Amount); valid && s.PlaceEight {
			b.PlaceEight = int(float64(s.PlaceEightAmt * g.Unit) * PAYOUT_OFFSET);
		}

		if valid, _ := b.validateWager(s.Amount); valid && s.PlaceNine {
			b.PlaceNine = s.PlaceNineAmt * g.Unit;
		}

		if valid, _ := b.validateWager(s.Amount); valid && s.PlaceTen {
			b.PlaceTen = s.PlaceTenAmt * g.Unit;
		}

		if valid, _ := b.validateWager(s.Amount); valid && g.ComeFour && s.Come {
			b.ComeFour = s.ComeAmt * g.Unit;
			b.ComeFourOdds = s.ComeFourOdds * g.Unit;
		}

		if valid, _ := b.validateWager(s.Amount); valid && g.ComeFive && s.Come {
			b.ComeFive = s.ComeAmt * g.Unit;
			b.ComeFiveOdds = int(float64(s.ComeFiveOdds * g.Unit) * PAYOUT_OFFSET);
		}

		if valid, _ := b.validateWager(s.Amount); valid && g.ComeSix && s.Come {
			b.ComeSix = s.ComeAmt * g.Unit;
			b.ComeSixOdds = s.ComeSixOdds * g.Unit;
		}
		
		if valid, _ := b.validateWager(s.Amount); valid && g.ComeEight && s.Come {
			b.ComeEight = s.ComeAmt * g.Unit;
			b.ComeEightOdds = s.ComeEightOdds * g.Unit;
		}

		if valid, _ := b.validateWager(s.Amount); valid && g.ComeNine && s.Come {
			b.ComeNine = s.ComeAmt * g.Unit;
			b.ComeNineOdds = int(float64(s.ComeNineOdds * g.Unit) * PAYOUT_OFFSET);
		}

		if valid, _ := b.validateWager(s.Amount); valid && g.ComeTen && s.Come {
			b.ComeTen = s.ComeAmt * g.Unit;
			b.ComeTenOdds = s.ComeTenOdds * g.Unit;
		}

		switch g.DontCome {
		case 4:
			if valid, _ := b.validateWager(s.Amount); valid && s.DontComeFour {
				b.DontComeFour = s.DontComeAmt * g.Unit;
				b.DontComeFourOdds = s.DontComeFourOdds * g.Unit;
			}
		case 5:
			if valid, _ := b.validateWager(s.Amount); valid && s.DontComeFive {
				b.DontComeFive = s.DontComeAmt * g.Unit;
				b.DontComeFiveOdds = s.DontComeFiveOdds * g.Unit;
			}
		case 6:
			if valid, _ := b.validateWager(s.Amount); valid && s.DontComeSix {
				b.DontComeSix = s.DontComeAmt * g.Unit;
				b.DontComeSixOdds = int(float64(s.DontComeSixOdds * g.Unit) * PAYOUT_OFFSET);
			}
		case 8:
			if valid, _ := b.validateWager(s.Amount); valid && s.DontComeEight {
				b.DontComeEight = s.DontComeAmt * g.Unit;
				b.DontComeEightOdds = int(float64(s.DontComeEightOdds * g.Unit) * PAYOUT_OFFSET);
			}
		case 9:
			if valid, _ := b.validateWager(s.Amount); valid && s.DontComeNine {
				b.DontComeNine = s.DontComeAmt * g.Unit;
				b.DontComeNineOdds = s.DontComeNineOdds * g.Unit;
			}
		case 10:
			if valid, _ := b.validateWager(s.Amount); valid && s.DontComeTen {
				b.DontComeTen = s.DontComeAmt * g.Unit;
				b.DontComeTenOdds = s.DontComeTenOdds * g.Unit;
			}
		}

		if valid, _ := b.validateWager(s.Amount); valid && s.Field {
			b.Field = g.Unit;
		}

		if valid, _ := b.validateWager(s.Amount); valid && s.HardSix {
			b.HardSix = g.Unit;
		}

		if valid, _ := b.validateWager(s.Amount); valid && s.HardEight {
			b.HardEight = g.Unit;
		}

		if valid, _ := b.validateWager(s.Amount); valid && s.HardFour {
			b.HardFour = g.Unit;
		}

		if valid, _ := b.validateWager(s.Amount); valid && s.HardTen {
			b.HardTen = g.Unit;
		}
	} else {

	}

	switch s.Line {
	case 1:
		if valid, _ := b.validateWager(s.Amount); valid {
			b.PassLine = s.Line * g.Unit;
			b.PassOdds = s.LineOdds * g.Unit;
		}
	case 2:
		if valid, _ := b.validateWager(s.Amount); valid {
			b.DontPass = s.Line * g.Unit;
			b.DontOdds = s.LineOdds * g.Unit;
		}
	case 3:
	}

	_, wager := b.validateWager(s.Amount);

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

	wager += b.PassLine;
	wager += b.PassOdds;
	wager += b.DontPass;
	wager += b.DontOdds;

	wager += b.Field;

	return wager <= amount, wager;
}