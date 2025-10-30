package common

import (
	"fmt"
	"strings"
)

type Symbol struct {
	Symbol string `json:"symbol"`
	Base   string `json:"base"`
	Quote  string `json:"quote"`
}

func NewSymbol(symbol string) *Symbol {
	s := &Symbol{
		Symbol: symbol,
	}
	s.parse()
	return s
}

func (s *Symbol) String() string {
	if s.Symbol == "" {
		return fmt.Sprintf("%s-%s", s.Base, s.Quote)
	}
	return s.Symbol
}

func (s *Symbol) parse() {
	if s.Base != "" || s.Symbol == "" {
		return
	}
	s.Symbol = strings.ToUpper(s.Symbol)
	if strings.Contains(s.Symbol, "-") {
		sp := strings.SplitN(s.Symbol, "-", 2)
		if len(sp) != 2 {
			return
		}
		base := sp[0]
		quote := sp[1]
		if base == "" || quote == "" {
			return
		}
		if s.Quote != "" && quote != s.Quote {
			return
		}
		s.Base = base
		s.Quote = quote
	} else if s.Quote == "" {
		if strings.HasSuffix(s.Symbol, "USDT") {
			s.Quote = "USDT"
			s.Base = strings.TrimSuffix(s.Symbol, "USDT")
		} else if strings.HasSuffix(s.Symbol, "USDC") {
			s.Quote = "USDC"
			s.Base = strings.TrimSuffix(s.Symbol, "USDC")
		} else if strings.HasSuffix(s.Symbol, "BTC") {
			s.Quote = "BTC"
			s.Base = strings.TrimSuffix(s.Symbol, "BTC")
		} else if strings.HasSuffix(s.Symbol, "ETH") {
			s.Quote = "ETH"
			s.Base = strings.TrimSuffix(s.Symbol, "ETH")
		} else if strings.HasSuffix(s.Symbol, "SOL") {
			s.Quote = "SOL"
			s.Base = strings.TrimSuffix(s.Symbol, "SOL")
		} else {
			return
		}
	}
	s.Symbol = s.Format()
}

func (s *Symbol) Format(spe ...string) string {
	s.parse()
	if s.Base == "" {
		return ""
	}
	sp := "-"
	if len(spe) > 0 {
		sp = spe[0]
	}
	return fmt.Sprintf("%s%s%s", s.Base, sp, s.Quote)
}

func (s *Symbol) Equals(other *Symbol) bool {
	return s.Base == other.Base && s.Quote == other.Quote
}
