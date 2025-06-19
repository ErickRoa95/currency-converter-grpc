package model

import "log"

type Currency struct {
	CountryCode string
	CurrencyName string
	Exchange float32
  Base string
}

func(c *Currency) Describe() {
	log.Printf("$1 %s equals %.2f %s", c.Base, c.Exchange, c.CurrencyName)
}
