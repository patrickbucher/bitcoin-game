package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	startCHF      = 100000.0
	startBTCRatio = 0.5
	interval      = 60 * time.Second
)

var (
	mu       sync.Mutex
	chfInBTC float64
)

func main() {
	var initialTotal, currentTotal float64
	balanceCHF := startCHF

	if len(os.Args) < 2 {
		log.Fatalf("usage: %s [rates file]", os.Args[0])
	}
	ratesPath := os.Args[1]

	ratesFile, err := os.OpenFile(ratesPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatalf("open %s: %v", ratesPath, err)
	}
	fmt.Printf("use tail -f %s to see exchange rate updates\n", ratesPath)

	ready := make(chan struct{})
	go updateRate(interval, ratesFile, ready)
	<-ready
	balanceBTC := (startCHF * startBTCRatio) * chfInBTC // buy bitcoin for half of the funds
	balanceCHF *= startBTCRatio
	initialTotal = balanceCHF + balanceBTC*(1.0/chfInBTC)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		currentTotal = balanceCHF + balanceBTC*(1.0/chfInBTC)
		fmt.Printf("Balance: %12.5f CHF\n", balanceCHF)
		fmt.Printf("Balance: %12.5f BTC\n", balanceBTC)
		fmt.Printf("Total:   %12.5f CHF (Δ %.2f)\n", currentTotal, currentTotal-initialTotal)
		fmt.Println("enter [b]: buy, [s]: sell, [q]: quit, [c]: calculate")
		input := strings.ToLower(readline(scanner))
		if len(input) == 0 {
			continue
		}
		if strings.HasPrefix(input, "q") {
			break
		}
		if strings.HasPrefix(input, "c") {
			continue
		}
		if strings.HasPrefix(input, "b") {
			fmt.Print("buy for CHF: ")
			buy := readnum(scanner)
			if buy <= 0.0 || buy > balanceCHF {
				fmt.Printf("can't buy for %f CHF\n", buy)
				continue
			}
			mu.Lock()
			balanceCHF -= buy
			balanceBTC += buy * chfInBTC
			mu.Unlock()
		}
		if strings.HasPrefix(input, "s") {
			fmt.Print("sell BTC: ")
			sell := readnum(scanner)
			if sell <= 0.0 || sell > balanceBTC {
				fmt.Printf("can't sell %f BTC\n", sell)
				continue
			}
			mu.Lock()
			balanceBTC -= sell
			balanceCHF += sell * (1.0 / chfInBTC)
			mu.Unlock()
		}
	}
}

func readnum(scanner *bufio.Scanner) float64 {
	line := readline(scanner)
	if line == "" {
		return 0.0
	}
	if val, err := strconv.ParseFloat(line, 64); err != nil {
		return 0.0
	} else {
		return val
	}
}

func readline(scanner *bufio.Scanner) string {
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text())
	}
	return ""
}

func updateRate(interval time.Duration, output io.Writer, ready chan<- struct{}) {
	wasReady := false
	for {
		mu.Lock()
		oldBTCinCHF := 1.0 / chfInBTC
		rate, err := getCHFInBTC()
		if err != nil {
			log.Printf("get CHF in BTC: %v", err)
		} else {
			chfInBTC = rate
		}
		newBTCinCHF := 1.0 / chfInBTC
		mu.Unlock()
		if !wasReady {
			ready <- struct{}{}
		}
		wasReady = true
		btcDiff := newBTCinCHF - oldBTCinCHF
		fmt.Fprintf(output, "BTC in CHF: %12.5f (Δ %.5f)\n", 1.0/chfInBTC, btcDiff)
		time.Sleep(interval)
	}
}

func getCHFInBTC() (float64, error) {
	url := "https://blockchain.info/tobtc?currency=CHF&value=1"
	resp, err := http.Get(url)
	if err != nil {
		return 0.0, fmt.Errorf("retrieve url %s: %v", url, err)
	}
	defer resp.Body.Close()
	r := bufio.NewReader(resp.Body)
	line, err := r.ReadString(byte('\n'))
	if err != nil && err != io.EOF {
		return 0.0, fmt.Errorf("read body: %v", err)
	}
	rate, err := strconv.ParseFloat(line, 64)
	if err != nil {
		return 0.0, fmt.Errorf("convert '%s' to float64: %v", line, err)
	}
	return rate, nil
}
