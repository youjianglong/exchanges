package binance

import (
	"time"

	. "github.com/youjianglong/exchanges/common"
)

type (
	UserDataEventType string
	FutureSubtype     string
	TimeInForceType   string
)

const (
	timestampKey  = "timestamp"
	signatureKey  = "signature"
	recvWindowKey = "recvWindow"

	UserDataEventTypeFutureOrderTradeUpdate UserDataEventType = "ORDER_TRADE_UPDATE"
	UserDataEventTypeMarginOrderTradeUpdate UserDataEventType = "executionReport"

	FutureSubtypeUM FutureSubtype = "UM"
	FutureSubtypeCM FutureSubtype = "CM"

	UserDataEventTypeOutboundAccountPosition UserDataEventType = "outboundAccountPosition" // 帐户资产变更
	UserDataEventTypeBalanceUpdate           UserDataEventType = "balanceUpdate"           // 余额更新事件
	UserDataEventTypeExecutionReport         UserDataEventType = "executionReport"         // 订单更新事件
	UserDataEventTypeListStatus              UserDataEventType = "ListStatus"              // 挂单列表事件
	UserDataEventTypeListenKeyExpired        UserDataEventType = "listenKeyExpired"        // 监听key过期事件
)

var (
	// Endpoints
	BaseApiURL                 = "https://api.binance.com"
	BasePApiURL                = "https://papi.binance.com"
	BaseFApiURL                = "https://fapi.binance.com"
	BaseWsSpotMainURL          = "wss://stream.binance.com:9443/ws/"
	BaseWsTestnetURL           = "wss://testnet.binance.vision/ws/"
	BaseSpotCombinedMainURL    = "wss://stream.binance.com:9443/stream?streams="
	BaseSpotCombinedTestnetURL = "wss://testnet.binance.vision/stream?streams="
	BaseWsApiMainURL           = "wss://ws-api.binance.com:443/ws-api/v3"
	BaseWsApiTestnetURL        = "wss://testnet.binance.vision/ws-api/v3"

	BaseWsSwapMainURL            = "wss://fstream.binance.com/ws/"
	BaseWsSwapTestnetURL         = "wss://testnet.binance.vision/ws/"
	BaseWsSwapCombinedMainURL    = "wss://fstream.binance.com/stream?streams="
	BaseWsSwapCombinedTestnetURL = "wss://testnet.binance.vision/stream?streams="

	ListenKeyLifetime = time.Hour // listenkey默认有效期1小时

	// https://developers.binance.com/docs/zh-CN/derivatives/portfolio-margin/user-data-streams
	PortfolioMarginUserWsURL = "wss://fstream.binance.com/pm/ws/"

	// WebsocketPingInterval is interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketPingInterval = time.Second * 60
	// WebsocketPingPongTimeout is timeout for sending a ping/pong message
	WebsocketPingPongTimeout = time.Second * 10
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive = true
)

var lotSizeKey = []byte("LOT_SIZE")

type LotSize struct {
	FilterType string  `json:"filterType"`
	MinQty     Float64 `json:"minQty"`
	MaxQty     Float64 `json:"maxQty"`
	StepSize   Float64 `json:"stepSize"`
}

type SymbolInfo struct {
	Symbol              string  // 交易对
	QuoteAsset          string  // 报价币种
	PricePrecision      int     // 价格精度
	QuantityPrecision   int     // 数量精度
	BaseAssetPrecision  int     // 基础货币精度
	QuoteAssetPrecision int     // 报价货币精度
	LotSize             LotSize // 限价单下单量
	MarketLotSize       LotSize // 市价单下单量
}
