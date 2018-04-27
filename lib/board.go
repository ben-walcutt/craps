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

		if valid, _ := b.validateWager(s.Amount, g.Unit); valid && s.PlaceFour {
			b.PlaceFour = s.PlaceFourAmt * g.Unit;
		}

		if valid, _ := b.validateWager(s.Amount, g.Unit); valid && s.PlaceFive {
			b.PlaceFive = s.PlaceFiveAmt * g.Unit;
		}

		if valid, _ := b.validateWager(s.Amount, g.Unit); valid && s.PlaceSix {
			b.PlaceSix = int(float64(s.PlaceSixAmt * g.Unit) * PAYOUT_OFFSET);
		}

		if valid, _ := b.validateWager(s.Amount, g.Unit); valid && s.PlaceEight {
			b.PlaceEight = int(float64(s.PlaceEightAmt * g.Unit) * PAYOUT_OFFSET);
		}

		if valid, _ := b.validateWager(s.Amount, g.Unit); valid && s.PlaceNine {
			b.PlaceNine = s.PlaceNineAmt * g.Unit;
		}

		if valid, _ := b.validateWager(s.Amount, g.Unit); valid && s.PlaceTen {
			b.PlaceTen = s.PlaceTenAmt * g.Unit;
		}

		if valid, _ := b.validateWager(s.Amount, g.Unit); valid && g.ComeFour {
			b.ComeFour = s.ComeAmt * g.Unit;
			b.ComeFourOdds = s.ComeFourOdds * g.Unit;
		}

		if valid, _ := b.validateWager(s.Amount, g.Unit); valid && g.ComeFive{
			b.ComeFive = s.ComeAmt * g.Unit;
			b.ComeFiveOdds = int(float64(s.ComeFiveOdds * g.Unit) * PAYOUT_OFFSET);
		}

		if valid, _ := b.validateWager(s.Amount, g.Unit); valid && g.ComeSix {
			b.ComeSix = int(float64(s.ComeAmt * g.Unit) * PAYOUT_OFFSET);
			b.ComeSixOdds = s.ComeAmt * g.Unit;
		}
		
		if valid, _ := b.validateWager(s.Amount, g.Unit); valid && g.ComeEight {
			b.ComeEight = int(float64(s.ComeAmt * g.Unit) * PAYOUT_OFFSET);
			b.ComeEightOdds = s.ComeEightOdds * g.Unit;
		}

		if valid, _ := b.validateWager(s.Amount, g.Unit); valid && g.ComeNine {
			b.ComeNine = s.ComeAmt * g.Unit;
			b.ComeNineOdds = int(float64(s.ComeNineOdds * g.Unit) * PAYOUT_OFFSET);
		}

		if valid, _ := b.validateWager(s.Amount, g.Unit); valid && g.ComeTen {
			b.ComeTen = s.ComeAmt * g.Unit;
			b.ComeTenOdds = s.ComeTenOdds * g.Unit;
		}

		switch g.DontCome {
		case 4:
			if valid, _ := b.validateWager(s.Amount, g.Unit); valid && s.DontComeFour {
				b.DontComeFour = s.DontComeAmt * g.Unit;
				b.DontComeFourOdds = s.DontComeFourOdds * g.Unit;
			}
		case 5:
			if valid, _ := b.validateWager(s.Amount, g.Unit); valid && s.DontComeFive {
				b.DontComeFive = s.DontComeAmt * g.Unit;
				b.DontComeFiveOdds = s.DontComeFiveOdds * g.Unit;
			}
		case 6:
			if valid, _ := b.validateWager(s.Amount, g.Unit); valid && s.DontComeSix {
				b.DontComeSix = s.DontComeAmt * g.Unit;
				b.DontComeSixOdds = int(float64(s.DontComeSixOdds * g.Unit) * PAYOUT_OFFSET);
			}
		case 8:
			if valid, _ := b.validateWager(s.Amount, g.Unit); valid && s.DontComeEight {
				b.DontComeEight = s.DontComeAmt * g.Unit;
				b.DontComeEightOdds = int(float64(s.DontComeEightOdds * g.Unit) * PAYOUT_OFFSET);
			}
		case 9:
			if valid, _ := b.validateWager(s.Amount, g.Unit); valid && s.DontComeNine {
				b.DontComeNine = s.DontComeAmt * g.Unit;
				b.DontComeNineOdds = s.DontComeNineOdds * g.Unit;
			}
		case 10:
			if valid, _ := b.validateWager(s.Amount, g.Unit); valid && s.DontComeTen {
				b.DontComeTen = s.DontComeAmt * g.Unit;
				b.DontComeTenOdds = s.DontComeTenOdds * g.Unit;
			}
		}

		if valid, _ := b.validateWager(s.Amount, g.Unit); valid && s.Field {
			b.Field = g.Unit;
		}

		if valid, _ := b.validateWager(s.Amount, g.Unit); valid && s.HardSix {
			b.HardSix = g.Unit;
		}

		if valid, _ := b.validateWager(s.Amount, g.Unit); valid && s.HardEight {
			b.HardEight = g.Unit;
		}

		if valid, _ := b.validateWager(s.Amount, g.Unit); valid && s.HardFour {
			b.HardFour = g.Unit;
		}

		if valid, _ := b.validateWager(s.Amount, g.Unit); valid && s.HardTen {
			b.HardTen = g.Unit;
		}
	} else {

	}

	switch s.Line {
	case 1:
		if valid, _ := b.validateWager(s.Amount, g.Unit); valid {
			b.PassLine = s.Line * g.Unit;
			b.PassOdds = s.LineOdds * g.Unit;
		}
	case 2:
		if valid, _ := b.validateWager(s.Amount, g.Unit); valid {
			b.DontPass = s.Line * g.Unit;
			b.DontOdds = s.LineOdds * g.Unit;
		}
	case 3:
	}

	_, wager := b.validateWager(s.Amount, g.Unit);

	return wager;	
}

func (b *Board) validateWager(amount int, unit int) (valid bool, wager int) {
	wager = 0;

	wager += b.PlaceFour * unit;
	wager += b.PlaceFive * unit;
	wager += b.PlaceSix;
	wager += b.PlaceEight;
	wager += b.PlaceNine * unit;
	wager += b.PlaceTen * unit;

	wager += b.Come * unit;
	wager += b.ComeFour * unit;
	wager += b.ComeFourOdds * unit;
	wager += b.ComeFive * unit;
	wager += b.ComeFiveOdds;
	wager += b.ComeSix * unit;
	wager += b.ComeSixOdds * unit;
	wager += b.ComeEight * unit;
	wager += b.ComeEightOdds * unit;
	wager += b.ComeNine * unit;
	wager += b.ComeNineOdds;
	wager += b.ComeTen * unit;
	wager += b.ComeTenOdds * unit;

	wager += b.DontComeFour * unit;
	wager += b.DontComeFourOdds * unit;
	wager += b.DontComeFive * unit;
	wager += b.DontComeFiveOdds * unit;
	wager += b.DontComeSix * unit;
	wager += b.DontComeSixOdds;
	wager += b.DontComeEight * unit;
	wager += b.DontComeEightOdds;
	wager += b.DontComeNine * unit;
	wager += b.DontComeNineOdds * unit;
	wager += b.DontComeTen * unit;
	wager += b.DontComeTenOdds * unit;

	wager += b.HardSix * unit;
	wager += b.HardEight * unit;
	wager += b.HardFour * unit;
	wager += b.HardTen * unit;

	wager += b.PassLine * unit;
	wager += b.PassOdds * unit;
	wager += b.DontPass * unit;
	wager += b.DontOdds * unit;

	wager += b.Field * unit;

	return wager <= amount, wager;
}