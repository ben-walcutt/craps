# craps

Genetic algorithm for finding optimal craps strategy

The program comes with several arguments:

- -c: how many children you want to generate
  - default: 20
- -i: number of iterations
  - default: 1000
- -n: run a named strategy
  - will default childer and iterations to one
- -r: number of rolls per iteration
  - default: 20
- -t: use test strategy
  - will default childer and iterations to one
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
