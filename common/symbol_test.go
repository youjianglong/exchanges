package common

import (
	"testing"
)

func TestSymbol_Format(t *testing.T) {
	symbol := NewSymbol("BTCUSDT")
	t.Logf("symbol: %v", symbol.Format(""))
}
