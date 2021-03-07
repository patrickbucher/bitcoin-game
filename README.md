# Bitcoin Game

Trade CHF and BTC and see how well you're doing.

## Usage

First, build the binary:

    $ go build

Second, start the game, indicating a file to write the exchange rates to:

    $ ./bitcoin-game rates.txt
    use tail -f rates.txt to see exchange rate updates
    Balance: 100000.00000 CHF
    Balance:      0.00000 BTC


Third, open a second terminal to `tail` the rates file to get updates (every 60
seconds):

    $ tail -f rates.txt
    BTC in CHF:  47438.33017 (Δ -Inf)
    BTC in CHF:  47438.33017 (Δ 0.00000)
    BTC in CHF:  47573.73930 (Δ 135.40913)
    BTC in CHF:  47573.73930 (Δ 0.00000)

Then have fun trading:

    enter [q]: quit, [b]: buy, [s]: sell
    b
    buy for CHF: 50000
    Balance:  50000.00000 CHF
    Balance:      1.05400 BTC
    enter [q]: quit, [b]: buy, [s]: sell
    s
    sell BTC: 1
    Balance:  97573.73930 CHF
    Balance:      0.05400 BTC
    enter [q]: quit, [b]: buy, [s]: sell
