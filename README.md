# Bitcoin Game

Trade CHF and BTC and see how well you're doing.

## Usage

First, build the binary:

    $ go build

Second, start the game, indicating a file to write the exchange rates to:

    $ ./bitcoin-game rates.txt
    use tail -f rates.txt to see exchange rate updates
    Balance:  50000.00000 CHF
    Balance:      1.06300 BTC
    Total:   100000.00000 CHF (Δ 0.00)


Third, open a second terminal to `tail` the rates file to get updates (every 60
seconds):

    $ tail -f rates.txt
    BTC in CHF:  47036.68862 (Δ -Inf)
    BTC in CHF:  47014.57452 (Δ -22.11410)
    BTC in CHF:  46992.48120 (Δ -22.09332)
    BTC in CHF:  47036.68862 (Δ 44.20741)
    
For Windows Powershell use:
    Get-Content -Path "rates.txt" -Wait
    BTC in CHF:  47258.97921 (Î” -Inf)
    BTC in CHF:  47125.35344 (Î” -133.62577)
    BTC in CHF:  47214.35316 (Î” 88.99972)
    BTC in CHF:  47147.57190 (Î” -66.78126)
    BTC in CHF:  47125.35344 (Î” -22.21846)
    BTC in CHF:  47103.15591 (Î” -22.19753)

Then have fun trading:

    enter [b]: buy, [s]: sell, [q]: quit, [c]: calculate
    c
    Balance:  50000.00000 CHF
    Balance:      1.06300 BTC
    Total:    99976.49271 CHF (Δ -23.51)
    enter [b]: buy, [s]: sell, [q]: quit, [c]: calculate
    c
    Balance:  50000.00000 CHF
    Balance:      1.06300 BTC
    Total:    99953.00752 CHF (Δ -46.99)
    enter [b]: buy, [s]: sell, [q]: quit, [c]: calculate
    b
    buy for CHF: 25000
    Balance:  25000.00000 CHF
    Balance:      1.59500 BTC
    Total:    99953.00752 CHF (Δ -46.99)
    enter [b]: buy, [s]: sell, [q]: quit, [c]: calculate
    s
    sell BTC: 1.0
    Balance:  72036.68862 CHF
    Balance:      0.59500 BTC
    Total:   100023.51834 CHF (Δ 23.52)
    enter [b]: buy, [s]: sell, [q]: quit, [c]: calculate
    q
