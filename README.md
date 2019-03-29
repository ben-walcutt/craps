# craps

Genetic algorithm for finding optimal craps strategy

The program comes with several arguments:

- -c: how many children you want to generate
  - default: 20
- -i: number of iterations
  - default: 1000
- -n: run a named strategy
  - will default children and iterations to one
- -r: number of rolls per iteration
  - default: 20
- -t: use test strategy
  - will default children and iterations to one
- -v: verbose output
- -a: starting amount
  - default: 300

Named Strategies available:

- Iron Cross
  - Place 5, 6, 8 one unit
  - Field one unit
  - Pass line one unit
  - 1x odds
- Come
  - Come bet one unit
  - No odds
- Field
  - Field bet one unit
- 22 Inside
  - Place 5, 6, 8, 9 one unit
- 20 Outside
  - Place 4, 5, 9, 10 one unit
- 32 Across
  - Place 4, 5, 6, 8, 9, 10 one unit

You can create a strategy for nearly any available bet on a craps table; however, there is not currently a way to turn off bets. Some restrictions are baked in.

- Pass Line
- Pass Line Odds
  - Currently only up to 3x odds
  - Todo to enhance
- Don't Pass Line
- Don't Pass Odds
  - Currently only up to 3x odds
- Place Bets
  - Limited to 2 units
- Come Bets
  - If enabled, every roll that a come bet is eligible
  - Limited to 2 units
- Come Bet odds
  - 3x, 4x, 5x odds
- Don't Come Bet
  - Currently only 1 allowed
- Don't Come odds
  - Up to 6x
- Field
  - 2 and 12 are double
- Hard Ways
- Horn
