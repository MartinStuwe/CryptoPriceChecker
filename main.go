package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Sci-Fi-AI/Go/tree/main/client"
)

func main() {
	fiatCurrency := flag.String(
		"fiat", "USD", "Name of Fiat",
	)

	nameOfCrypto := flag.String(
		"crypto", "BTC", "Name of Crypto",
	)

	flag.Parse()

	crypto, err := client.FetchCrytpo(*fiatCurrency, *nameOfCrypto)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(crypto)
}
