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
	BaseSpotApiMainURL      = "https://api.binance.com"                       // 现货API主网
	BaseSpotApiTestURL      = "https://testnet.binance.vision"                // 现货API测试网
	BaseSpotWsApiMainURL    = "wss://ws-api.binance.com:443/ws-api/v3"        // 现货WebSocket API主网
	BaseSpotWsApiTestURL    = "wss://ws-api.testnet.binance.vision/ws-api/v3" // 现货WebSocket API测试网
	BaseSpotWsMainURL       = "wss://stream.binance.com:9443/ws/"             // 现货主网行情推送WebSocket
	BaseSpotWsTestURL       = "wss://stream.testnet.binance.vision/ws/"       // 现货测试网行情推送WebSocket
	BaseSpotCombinedMainURL = "wss://stream.binance.com:9443/stream?streams=" // 现货主网行情推送WebSocket合并
	BaseSpotCombinedTestURL = "wss://testnet.binance.vision/stream?streams="  // 现货测试网行情推送WebSocket合并

	BaseFApiMainURL           = "https://fapi.binance.com"                     // U本位合约主网
	BaseFApiTestURL           = "https://demo-fapi.binance.com"                // U本位合约测试网
	BaseSwapWsMainURL         = "wss://fstream.binance.com/ws/"                // U本位合约WebSocket API
	BaseSwapWsTestURL         = "wss://testnet.binance.vision/ws/"             // U本位合约测试网WebSocket API
	BaseSwapCombinedWsMainURL = "wss://fstream.binance.com/stream?streams="    // U本位合约主网行情推送WebSocket合并
	BaseSwapCombinedWsTestURL = "wss://testnet.binance.vision/stream?streams=" // U本位合约测试网行情推送WebSocket合并

	// https://developers.binance.com/docs/zh-CN/derivatives/portfolio-margin/user-data-streams
	BasePApiURL              = "https://papi.binance.com"         // 统一账户API主网
	PortfolioMarginUserWsURL = "wss://fstream.binance.com/pm/ws/" // 统一账户WebSocket API

	FutureApiTestURL   = "https://testnet.binancefuture.com" // 币本位合约测试网
	FutureWsApiTestURL = "wss://dstream.binancefuture.com"   // 币本位合约测试网WebSocket API

	ListenKeyLifetime = time.Hour // listenkey默认有效期1小时
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
