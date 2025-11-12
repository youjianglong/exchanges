package okx

import (
	"strings"

	"github.com/youjianglong/exchanges/common"
)

func ToInstId(symbol *common.Symbol) string {
	return symbol.Format("-") + "-SWAP"
}

func ToSymbol(instId string) *common.Symbol {
	return common.NewSymbol(strings.TrimSuffix(instId, "-SWAP"))
}
