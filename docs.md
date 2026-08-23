# binance

```go
import "github.com/youjianglong/exchanges/binance"
```

## Index

- [Constants](<#constants>)
- [Variables](<#variables>)
- [func AsProxyError\(e error\) \(error, bool\)](<#AsProxyError>)
- [func ConvertProxyError\(e error\) error](<#ConvertProxyError>)
- [func IsAPIError\(e error\) bool](<#IsAPIError>)
- [type APIError](<#APIError>)
  - [func \(e APIError\) Error\(\) string](<#APIError.Error>)
  - [func \(e APIError\) IsValid\(\) bool](<#APIError.IsValid>)
- [type CapitalDepositHistory](<#CapitalDepositHistory>)
- [type CapitalDepositHistoryService](<#CapitalDepositHistoryService>)
  - [func \(s \*CapitalDepositHistoryService\) Coin\(coin string\) \*CapitalDepositHistoryService](<#CapitalDepositHistoryService.Coin>)
  - [func \(s \*CapitalDepositHistoryService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]CapitalDepositHistory, error\)](<#CapitalDepositHistoryService.Do>)
  - [func \(s \*CapitalDepositHistoryService\) EndTime\(endTime int64\) \*CapitalDepositHistoryService](<#CapitalDepositHistoryService.EndTime>)
  - [func \(s \*CapitalDepositHistoryService\) IncludeSource\(includeSource bool\) \*CapitalDepositHistoryService](<#CapitalDepositHistoryService.IncludeSource>)
  - [func \(s \*CapitalDepositHistoryService\) Limit\(limit int\) \*CapitalDepositHistoryService](<#CapitalDepositHistoryService.Limit>)
  - [func \(s \*CapitalDepositHistoryService\) Offset\(offset int\) \*CapitalDepositHistoryService](<#CapitalDepositHistoryService.Offset>)
  - [func \(s \*CapitalDepositHistoryService\) StartTime\(startTime int64\) \*CapitalDepositHistoryService](<#CapitalDepositHistoryService.StartTime>)
  - [func \(s \*CapitalDepositHistoryService\) Status\(status int\) \*CapitalDepositHistoryService](<#CapitalDepositHistoryService.Status>)
  - [func \(s \*CapitalDepositHistoryService\) TxId\(txId string\) \*CapitalDepositHistoryService](<#CapitalDepositHistoryService.TxId>)
- [type CapitalWithdrawHistory](<#CapitalWithdrawHistory>)
- [type CapitalWithdrawHistoryService](<#CapitalWithdrawHistoryService>)
  - [func \(s \*CapitalWithdrawHistoryService\) Coin\(coin string\) \*CapitalWithdrawHistoryService](<#CapitalWithdrawHistoryService.Coin>)
  - [func \(s \*CapitalWithdrawHistoryService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]CapitalWithdrawHistory, error\)](<#CapitalWithdrawHistoryService.Do>)
  - [func \(s \*CapitalWithdrawHistoryService\) EndTime\(endTime int64\) \*CapitalWithdrawHistoryService](<#CapitalWithdrawHistoryService.EndTime>)
  - [func \(s \*CapitalWithdrawHistoryService\) IdList\(idList \[\]string\) \*CapitalWithdrawHistoryService](<#CapitalWithdrawHistoryService.IdList>)
  - [func \(s \*CapitalWithdrawHistoryService\) Limit\(limit int\) \*CapitalWithdrawHistoryService](<#CapitalWithdrawHistoryService.Limit>)
  - [func \(s \*CapitalWithdrawHistoryService\) Offset\(offset int\) \*CapitalWithdrawHistoryService](<#CapitalWithdrawHistoryService.Offset>)
  - [func \(s \*CapitalWithdrawHistoryService\) StartTime\(startTime int64\) \*CapitalWithdrawHistoryService](<#CapitalWithdrawHistoryService.StartTime>)
  - [func \(s \*CapitalWithdrawHistoryService\) Status\(status int\) \*CapitalWithdrawHistoryService](<#CapitalWithdrawHistoryService.Status>)
  - [func \(s \*CapitalWithdrawHistoryService\) WithdrawOrderId\(withdrawOrderId string\) \*CapitalWithdrawHistoryService](<#CapitalWithdrawHistoryService.WithdrawOrderId>)
- [type Client](<#Client>)
  - [func NewClient\(apiKey, secretKey string\) \*Client](<#NewClient>)
  - [func NewClientWithHttpClient\(apiKey, secretKey string, httpClient \*http.Client\) \*Client](<#NewClientWithHttpClient>)
  - [func NewTestClient\(apiKey, secretKey string\) \*Client](<#NewTestClient>)
  - [func \(c \*Client\) NewFApiCancelOrderService\(symbol string\) \*FApiCancelOrderService](<#Client.NewFApiCancelOrderService>)
  - [func \(c \*Client\) NewFApiChangeLeverageService\(symbol string, leverage int64\) \*FApiChangeLeverageService](<#Client.NewFApiChangeLeverageService>)
  - [func \(c \*Client\) NewFApiGetAccountBalanceService\(\) \*FApiGetAccountBalanceService](<#Client.NewFApiGetAccountBalanceService>)
  - [func \(c \*Client\) NewFApiGetAllOpenOrdersService\(\) \*FApiGetAllOpenOrdersService](<#Client.NewFApiGetAllOpenOrdersService>)
  - [func \(c \*Client\) NewFApiGetAllOrdersService\(symbol string\) \*FApiGetAllOrdersService](<#Client.NewFApiGetAllOrdersService>)
  - [func \(c \*Client\) NewFApiGetBookTickerService\(\) \*FApiGetBookTickerService](<#Client.NewFApiGetBookTickerService>)
  - [func \(c \*Client\) NewFApiGetDepthService\(symbol string\) \*FApiGetDepthService](<#Client.NewFApiGetDepthService>)
  - [func \(c \*Client\) NewFApiGetIncomeService\(\) \*FApiGetIncomeService](<#Client.NewFApiGetIncomeService>)
  - [func \(c \*Client\) NewFApiGetKLinesService\(symbol, interval string\) \*FApiGetKLinesService](<#Client.NewFApiGetKLinesService>)
  - [func \(c \*Client\) NewFApiGetOrderService\(symbol string\) \*FApiGetOrderService](<#Client.NewFApiGetOrderService>)
  - [func \(c \*Client\) NewFApiGetPositionsService\(\) \*FApiGetPositionsService](<#Client.NewFApiGetPositionsService>)
  - [func \(c \*Client\) NewFApiGetService\(endpoint string, params params\) \*GetService](<#Client.NewFApiGetService>)
  - [func \(c \*Client\) NewFApiGetUserTradesService\(symbol string\) \*FApiGetUserTradesService](<#Client.NewFApiGetUserTradesService>)
  - [func \(c \*Client\) NewFApiOrderService\(symbol string, side string, orderType string\) \*FApiOrderService](<#Client.NewFApiOrderService>)
  - [func \(c \*Client\) NewFApiPostService\(endpoint string, params params\) \*PostService](<#Client.NewFApiPostService>)
  - [func \(c \*Client\) NewGetCapitalDepositHistoryService\(\) \*CapitalDepositHistoryService](<#Client.NewGetCapitalDepositHistoryService>)
  - [func \(c \*Client\) NewGetCapitalWithdrawHistoryService\(\) \*CapitalWithdrawHistoryService](<#Client.NewGetCapitalWithdrawHistoryService>)
  - [func \(c \*Client\) NewGetFApiSwapOpenOrdersService\(\) \*GetFApiSwapOpenOrdersService](<#Client.NewGetFApiSwapOpenOrdersService>)
  - [func \(c \*Client\) NewGetFApiSwapOrdersService\(\) \*GetFApiSwapOrdersService](<#Client.NewGetFApiSwapOrdersService>)
  - [func \(c \*Client\) NewGetFApiSymbolInfosService\(\) \*GetFApiSymbolInfosService](<#Client.NewGetFApiSymbolInfosService>)
  - [func \(c \*Client\) NewGetFApiTicker24HService\(\) \*GetFApiTicker24HService](<#Client.NewGetFApiTicker24HService>)
  - [func \(c \*Client\) NewGetFApiTickerPriceService\(\) \*GetFApiTickerPriceService](<#Client.NewGetFApiTickerPriceService>)
  - [func \(c \*Client\) NewGetPApiAccountBalanceService\(\) \*GetPApiAccountBalanceService](<#Client.NewGetPApiAccountBalanceService>)
  - [func \(c \*Client\) NewGetPApiSwapOpenOrdersService\(\) \*GetPApiSwapOpenOrdersService](<#Client.NewGetPApiSwapOpenOrdersService>)
  - [func \(c \*Client\) NewGetPremiumIndexService\(\) \*GetPremiumIndexService](<#Client.NewGetPremiumIndexService>)
  - [func \(c \*Client\) NewGetService\(baseURL \*string, endpoint string, params params\) \*GetService](<#Client.NewGetService>)
  - [func \(c \*Client\) NewGetSpotOpenOrdersService\(\) \*GetSpotOpenOrdersService](<#Client.NewGetSpotOpenOrdersService>)
  - [func \(c \*Client\) NewGetSpotOrdersService\(\) \*GetSpotOrdersService](<#Client.NewGetSpotOrdersService>)
  - [func \(c \*Client\) NewGetSpotSymbolInfosService\(\) \*GetSpotSymbolInfosService](<#Client.NewGetSpotSymbolInfosService>)
  - [func \(c \*Client\) NewGetSpotTicker24HService\(\) \*GetSpotTicker24HService](<#Client.NewGetSpotTicker24HService>)
  - [func \(c \*Client\) NewGetSpotTickerPriceService\(\) \*GetSpotTickerPriceService](<#Client.NewGetSpotTickerPriceService>)
  - [func \(c \*Client\) NewGetUMAccountDetailService\(\) \*GetUMAccountDetailService](<#Client.NewGetUMAccountDetailService>)
  - [func \(c \*Client\) NewGetUMAccountService\(\) \*GetUMAccountService](<#Client.NewGetUMAccountService>)
  - [func \(c \*Client\) NewGetUserAssetService\(\) \*GetUserAssetService](<#Client.NewGetUserAssetService>)
  - [func \(c \*Client\) NewGetWalletBalanceService\(\) \*GetWalletBalanceService](<#Client.NewGetWalletBalanceService>)
  - [func \(c \*Client\) NewPApiGetAllOrdersService\(symbol string\) \*PApiGetAllOrdersService](<#Client.NewPApiGetAllOrdersService>)
  - [func \(c \*Client\) NewPApiGetOpenOrdersService\(\) \*PApiGetOpenOrdersService](<#Client.NewPApiGetOpenOrdersService>)
  - [func \(c \*Client\) NewPApiGetService\(endpoint string, params params\) \*GetService](<#Client.NewPApiGetService>)
  - [func \(c \*Client\) NewPApiGetUMPositionsService\(\) \*PApiGetUMPositionsService](<#Client.NewPApiGetUMPositionsService>)
  - [func \(c \*Client\) NewPApiPostService\(endpoint string, params params\) \*PostService](<#Client.NewPApiPostService>)
  - [func \(c \*Client\) NewPApiUmCancelAllOpenOrdersService\(symbol string\) \*PApiUmCancelAllOpenOrdersService](<#Client.NewPApiUmCancelAllOpenOrdersService>)
  - [func \(c \*Client\) NewPApiUmCancelOrderService\(\) \*PApiUmCancelOrderService](<#Client.NewPApiUmCancelOrderService>)
  - [func \(c \*Client\) NewPApiUmChangeLeverageService\(symbol string, leverage int64\) \*PApiUmChangeLeverageService](<#Client.NewPApiUmChangeLeverageService>)
  - [func \(c \*Client\) NewPApiUmGetIncomeService\(\) \*PApiUmGetIncomeService](<#Client.NewPApiUmGetIncomeService>)
  - [func \(c \*Client\) NewPApiUmGetOrderService\(symbol string\) \*PApiUmGetOrderService](<#Client.NewPApiUmGetOrderService>)
  - [func \(c \*Client\) NewPApiUmGetUserTradesService\(symbol string\) \*PApiUmGetUserTradesService](<#Client.NewPApiUmGetUserTradesService>)
  - [func \(c \*Client\) NewPApiUmOrderService\(symbol string, side string, orderType string\) \*PApiUmOrderService](<#Client.NewPApiUmOrderService>)
  - [func \(c \*Client\) NewPapiUMGetPositionRiskService\(\) \*PapiUMGetPositionRiskService](<#Client.NewPapiUMGetPositionRiskService>)
  - [func \(c \*Client\) NewPingService\(\) \*PingService](<#Client.NewPingService>)
  - [func \(c \*Client\) NewPostService\(baseURL \*string, endpoint string, params params\) \*PostService](<#Client.NewPostService>)
  - [func \(c \*Client\) NewSpotGetService\(endpoint string, params params\) \*GetService](<#Client.NewSpotGetService>)
  - [func \(c \*Client\) NewSpotPostService\(endpoint string, params params\) \*PostService](<#Client.NewSpotPostService>)
  - [func \(c \*Client\) SetProxyURL\(proxyURL \*url.URL\)](<#Client.SetProxyURL>)
  - [func \(c \*Client\) WithHttpClient\(httpClient \*http.Client\) \*Client](<#Client.WithHttpClient>)
- [type Depth](<#Depth>)
- [type FApiAccountBalance](<#FApiAccountBalance>)
- [type FApiBookTicker](<#FApiBookTicker>)
- [type FApiCancelOrderService](<#FApiCancelOrderService>)
  - [func \(s \*FApiCancelOrderService\) Do\(ctx context.Context, opts ...RequestOption\) \(\*SwapOrder, error\)](<#FApiCancelOrderService.Do>)
  - [func \(s \*FApiCancelOrderService\) OrderId\(orderId string\) \*FApiCancelOrderService](<#FApiCancelOrderService.OrderId>)
  - [func \(s \*FApiCancelOrderService\) OrigClientOrderId\(origClientOrderId string\) \*FApiCancelOrderService](<#FApiCancelOrderService.OrigClientOrderId>)
- [type FApiChangeLeverageService](<#FApiChangeLeverageService>)
  - [func \(s \*FApiChangeLeverageService\) Do\(ctx context.Context, opts ...RequestOption\) error](<#FApiChangeLeverageService.Do>)
- [type FApiGetAccountBalanceService](<#FApiGetAccountBalanceService>)
  - [func \(s \*FApiGetAccountBalanceService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*FApiAccountBalance, error\)](<#FApiGetAccountBalanceService.Do>)
- [type FApiGetAllOpenOrdersService](<#FApiGetAllOpenOrdersService>)
  - [func \(s \*FApiGetAllOpenOrdersService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*SwapOpenOrder, error\)](<#FApiGetAllOpenOrdersService.Do>)
  - [func \(s \*FApiGetAllOpenOrdersService\) Symbol\(symbol string\) \*FApiGetAllOpenOrdersService](<#FApiGetAllOpenOrdersService.Symbol>)
- [type FApiGetAllOrdersService](<#FApiGetAllOrdersService>)
  - [func \(s \*FApiGetAllOrdersService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*SwapOrder, error\)](<#FApiGetAllOrdersService.Do>)
  - [func \(s \*FApiGetAllOrdersService\) EndTime\(endTime int64\) \*FApiGetAllOrdersService](<#FApiGetAllOrdersService.EndTime>)
  - [func \(s \*FApiGetAllOrdersService\) Limit\(limit int\) \*FApiGetAllOrdersService](<#FApiGetAllOrdersService.Limit>)
  - [func \(s \*FApiGetAllOrdersService\) OrderId\(orderId string\) \*FApiGetAllOrdersService](<#FApiGetAllOrdersService.OrderId>)
  - [func \(s \*FApiGetAllOrdersService\) StartTime\(startTime int64\) \*FApiGetAllOrdersService](<#FApiGetAllOrdersService.StartTime>)
- [type FApiGetBookTickerService](<#FApiGetBookTickerService>)
  - [func \(s \*FApiGetBookTickerService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*FApiBookTicker, error\)](<#FApiGetBookTickerService.Do>)
  - [func \(s \*FApiGetBookTickerService\) Symbol\(symbol string\) \*FApiGetBookTickerService](<#FApiGetBookTickerService.Symbol>)
- [type FApiGetDepthService](<#FApiGetDepthService>)
  - [func \(s \*FApiGetDepthService\) Do\(ctx context.Context, opts ...RequestOption\) \(\*Depth, error\)](<#FApiGetDepthService.Do>)
  - [func \(s \*FApiGetDepthService\) Limit\(limit int\) \*FApiGetDepthService](<#FApiGetDepthService.Limit>)
- [type FApiGetIncomeService](<#FApiGetIncomeService>)
  - [func \(s \*FApiGetIncomeService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]Income, error\)](<#FApiGetIncomeService.Do>)
  - [func \(s \*FApiGetIncomeService\) EndTime\(endTime int64\) \*FApiGetIncomeService](<#FApiGetIncomeService.EndTime>)
  - [func \(s \*FApiGetIncomeService\) IncomeType\(incomeType string\) \*FApiGetIncomeService](<#FApiGetIncomeService.IncomeType>)
  - [func \(s \*FApiGetIncomeService\) Limit\(limit int\) \*FApiGetIncomeService](<#FApiGetIncomeService.Limit>)
  - [func \(s \*FApiGetIncomeService\) Page\(page int\) \*FApiGetIncomeService](<#FApiGetIncomeService.Page>)
  - [func \(s \*FApiGetIncomeService\) StartTime\(startTime int64\) \*FApiGetIncomeService](<#FApiGetIncomeService.StartTime>)
  - [func \(s \*FApiGetIncomeService\) Symbol\(symbol string\) \*FApiGetIncomeService](<#FApiGetIncomeService.Symbol>)
- [type FApiGetKLinesService](<#FApiGetKLinesService>)
  - [func \(s \*FApiGetKLinesService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\[\]Mixed, error\)](<#FApiGetKLinesService.Do>)
  - [func \(s \*FApiGetKLinesService\) EndTime\(endTime int64\) \*FApiGetKLinesService](<#FApiGetKLinesService.EndTime>)
  - [func \(s \*FApiGetKLinesService\) Limit\(limit int\) \*FApiGetKLinesService](<#FApiGetKLinesService.Limit>)
  - [func \(s \*FApiGetKLinesService\) StartTime\(startTime int64\) \*FApiGetKLinesService](<#FApiGetKLinesService.StartTime>)
- [type FApiGetOrderService](<#FApiGetOrderService>)
  - [func \(s \*FApiGetOrderService\) Do\(ctx context.Context, opts ...RequestOption\) \(\*SwapOrder, error\)](<#FApiGetOrderService.Do>)
  - [func \(s \*FApiGetOrderService\) OrderId\(orderId string\) \*FApiGetOrderService](<#FApiGetOrderService.OrderId>)
  - [func \(s \*FApiGetOrderService\) OrigClientOrderId\(origClientOrderId string\) \*FApiGetOrderService](<#FApiGetOrderService.OrigClientOrderId>)
- [type FApiGetPositionsService](<#FApiGetPositionsService>)
  - [func \(s \*FApiGetPositionsService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*Position, error\)](<#FApiGetPositionsService.Do>)
- [type FApiGetUserTradesService](<#FApiGetUserTradesService>)
  - [func \(s \*FApiGetUserTradesService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*TradeRecord, error\)](<#FApiGetUserTradesService.Do>)
  - [func \(s \*FApiGetUserTradesService\) EndTime\(endTime int64\) \*FApiGetUserTradesService](<#FApiGetUserTradesService.EndTime>)
  - [func \(s \*FApiGetUserTradesService\) FromId\(fromId int64\) \*FApiGetUserTradesService](<#FApiGetUserTradesService.FromId>)
  - [func \(s \*FApiGetUserTradesService\) Limit\(limit int\) \*FApiGetUserTradesService](<#FApiGetUserTradesService.Limit>)
  - [func \(s \*FApiGetUserTradesService\) OrderId\(orderId string\) \*FApiGetUserTradesService](<#FApiGetUserTradesService.OrderId>)
  - [func \(s \*FApiGetUserTradesService\) StartTime\(startTime int64\) \*FApiGetUserTradesService](<#FApiGetUserTradesService.StartTime>)
- [type FApiOrderService](<#FApiOrderService>)
  - [func \(s \*FApiOrderService\) ActivationPrice\(activationPrice string\) \*FApiOrderService](<#FApiOrderService.ActivationPrice>)
  - [func \(s \*FApiOrderService\) CallbackRate\(callbackRate string\) \*FApiOrderService](<#FApiOrderService.CallbackRate>)
  - [func \(s \*FApiOrderService\) ClosePosition\(closePosition string\) \*FApiOrderService](<#FApiOrderService.ClosePosition>)
  - [func \(s \*FApiOrderService\) Do\(ctx context.Context, opts ...RequestOption\) \(\*SwapOrder, error\)](<#FApiOrderService.Do>)
  - [func \(s \*FApiOrderService\) GoodTillDate\(goodTillDate int64\) \*FApiOrderService](<#FApiOrderService.GoodTillDate>)
  - [func \(s \*FApiOrderService\) NewClientOrderId\(newClientOrderId string\) \*FApiOrderService](<#FApiOrderService.NewClientOrderId>)
  - [func \(s \*FApiOrderService\) NewOrderRespType\(newOrderRespType string\) \*FApiOrderService](<#FApiOrderService.NewOrderRespType>)
  - [func \(s \*FApiOrderService\) OrderType\(orderType string\) \*FApiOrderService](<#FApiOrderService.OrderType>)
  - [func \(s \*FApiOrderService\) PositionSide\(positionSide string\) \*FApiOrderService](<#FApiOrderService.PositionSide>)
  - [func \(s \*FApiOrderService\) Price\(price string\) \*FApiOrderService](<#FApiOrderService.Price>)
  - [func \(s \*FApiOrderService\) PriceMatch\(priceMatch string\) \*FApiOrderService](<#FApiOrderService.PriceMatch>)
  - [func \(s \*FApiOrderService\) PriceProtect\(priceProtect bool\) \*FApiOrderService](<#FApiOrderService.PriceProtect>)
  - [func \(s \*FApiOrderService\) Quantity\(quantity string\) \*FApiOrderService](<#FApiOrderService.Quantity>)
  - [func \(s \*FApiOrderService\) ReduceOnly\(reduceOnly bool\) \*FApiOrderService](<#FApiOrderService.ReduceOnly>)
  - [func \(s \*FApiOrderService\) SelfTradePreventionMode\(selfTradePreventionMode string\) \*FApiOrderService](<#FApiOrderService.SelfTradePreventionMode>)
  - [func \(s \*FApiOrderService\) StopPrice\(stopPrice string\) \*FApiOrderService](<#FApiOrderService.StopPrice>)
  - [func \(s \*FApiOrderService\) TimeInForce\(timeInForce string\) \*FApiOrderService](<#FApiOrderService.TimeInForce>)
  - [func \(s \*FApiOrderService\) WorkingType\(workingType string\) \*FApiOrderService](<#FApiOrderService.WorkingType>)
- [type FApiSymbolInfos](<#FApiSymbolInfos>)
- [type FApiTicker24H](<#FApiTicker24H>)
- [type FutureSubtype](<#FutureSubtype>)
- [type GetFApiSwapOpenOrdersService](<#GetFApiSwapOpenOrdersService>)
  - [func \(s \*GetFApiSwapOpenOrdersService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*SwapOpenOrder, error\)](<#GetFApiSwapOpenOrdersService.Do>)
- [type GetFApiSwapOrdersService](<#GetFApiSwapOrdersService>)
  - [func \(s \*GetFApiSwapOrdersService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*SwapOrder, error\)](<#GetFApiSwapOrdersService.Do>)
  - [func \(s \*GetFApiSwapOrdersService\) EndTime\(endTime int64\) SwapOrdersGetter](<#GetFApiSwapOrdersService.EndTime>)
  - [func \(s \*GetFApiSwapOrdersService\) Limit\(limit int\) SwapOrdersGetter](<#GetFApiSwapOrdersService.Limit>)
  - [func \(s \*GetFApiSwapOrdersService\) OrderId\(orderId string\) SwapOrdersGetter](<#GetFApiSwapOrdersService.OrderId>)
  - [func \(s \*GetFApiSwapOrdersService\) StartTime\(startTime int64\) SwapOrdersGetter](<#GetFApiSwapOrdersService.StartTime>)
  - [func \(s \*GetFApiSwapOrdersService\) Symbol\(symbol string\) SwapOrdersGetter](<#GetFApiSwapOrdersService.Symbol>)
- [type GetFApiSymbolInfosService](<#GetFApiSymbolInfosService>)
  - [func \(s \*GetFApiSymbolInfosService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*SymbolInfo, error\)](<#GetFApiSymbolInfosService.Do>)
- [type GetFApiTicker24HService](<#GetFApiTicker24HService>)
  - [func \(s \*GetFApiTicker24HService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*FApiTicker24H, error\)](<#GetFApiTicker24HService.Do>)
  - [func \(s \*GetFApiTicker24HService\) Symbol\(symbol string\) \*GetFApiTicker24HService](<#GetFApiTicker24HService.Symbol>)
- [type GetFApiTickerPriceService](<#GetFApiTickerPriceService>)
  - [func \(s \*GetFApiTickerPriceService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*TickerPrice, error\)](<#GetFApiTickerPriceService.Do>)
  - [func \(s \*GetFApiTickerPriceService\) Symbol\(symbol string\) \*GetFApiTickerPriceService](<#GetFApiTickerPriceService.Symbol>)
- [type GetPApiAccountBalanceService](<#GetPApiAccountBalanceService>)
  - [func \(s \*GetPApiAccountBalanceService\) Asset\(asset string\) \*GetPApiAccountBalanceService](<#GetPApiAccountBalanceService.Asset>)
  - [func \(s \*GetPApiAccountBalanceService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*PApiAccountBalance, error\)](<#GetPApiAccountBalanceService.Do>)
- [type GetPApiSwapOpenOrdersService](<#GetPApiSwapOpenOrdersService>)
  - [func \(s \*GetPApiSwapOpenOrdersService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*SwapOpenOrder, error\)](<#GetPApiSwapOpenOrdersService.Do>)
  - [func \(s \*GetPApiSwapOpenOrdersService\) Symbol\(symbol string\) SwapOpenOrdersGetter](<#GetPApiSwapOpenOrdersService.Symbol>)
- [type GetPremiumIndexService](<#GetPremiumIndexService>)
  - [func \(s \*GetPremiumIndexService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*PremiumIndex, error\)](<#GetPremiumIndexService.Do>)
  - [func \(s \*GetPremiumIndexService\) Symbol\(symbol string\) \*GetPremiumIndexService](<#GetPremiumIndexService.Symbol>)
- [type GetService](<#GetService>)
  - [func \(s \*GetService\) Do\(ctx context.Context, dest any, opts ...RequestOption\) error](<#GetService.Do>)
- [type GetSpotOpenOrdersService](<#GetSpotOpenOrdersService>)
  - [func \(s \*GetSpotOpenOrdersService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*SpotOpenOrder, error\)](<#GetSpotOpenOrdersService.Do>)
  - [func \(s \*GetSpotOpenOrdersService\) Symbol\(symbol string\) \*GetSpotOpenOrdersService](<#GetSpotOpenOrdersService.Symbol>)
- [type GetSpotOrdersService](<#GetSpotOrdersService>)
  - [func \(s \*GetSpotOrdersService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*SpotOrder, error\)](<#GetSpotOrdersService.Do>)
  - [func \(s \*GetSpotOrdersService\) EndTime\(endTime int64\) \*GetSpotOrdersService](<#GetSpotOrdersService.EndTime>)
  - [func \(s \*GetSpotOrdersService\) Limit\(limit int\) \*GetSpotOrdersService](<#GetSpotOrdersService.Limit>)
  - [func \(s \*GetSpotOrdersService\) OrderId\(orderId string\) \*GetSpotOrdersService](<#GetSpotOrdersService.OrderId>)
  - [func \(s \*GetSpotOrdersService\) StartTime\(startTime int64\) \*GetSpotOrdersService](<#GetSpotOrdersService.StartTime>)
  - [func \(s \*GetSpotOrdersService\) Symbol\(symbol string\) \*GetSpotOrdersService](<#GetSpotOrdersService.Symbol>)
- [type GetSpotSymbolInfosService](<#GetSpotSymbolInfosService>)
  - [func \(s \*GetSpotSymbolInfosService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*SymbolInfo, error\)](<#GetSpotSymbolInfosService.Do>)
  - [func \(s \*GetSpotSymbolInfosService\) Status\(status string\) \*GetSpotSymbolInfosService](<#GetSpotSymbolInfosService.Status>)
- [type GetSpotTicker24HService](<#GetSpotTicker24HService>)
  - [func \(s \*GetSpotTicker24HService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*SpotTicker24H, error\)](<#GetSpotTicker24HService.Do>)
  - [func \(s \*GetSpotTicker24HService\) Symbol\(symbol string\) \*GetSpotTicker24HService](<#GetSpotTicker24HService.Symbol>)
  - [func \(s \*GetSpotTicker24HService\) Symbols\(symbols \[\]string\) \*GetSpotTicker24HService](<#GetSpotTicker24HService.Symbols>)
- [type GetSpotTickerPriceService](<#GetSpotTickerPriceService>)
  - [func \(s \*GetSpotTickerPriceService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*TickerPrice, error\)](<#GetSpotTickerPriceService.Do>)
  - [func \(s \*GetSpotTickerPriceService\) Symbol\(symbol string\) \*GetSpotTickerPriceService](<#GetSpotTickerPriceService.Symbol>)
  - [func \(s \*GetSpotTickerPriceService\) Symbols\(symbols \[\]string\) \*GetSpotTickerPriceService](<#GetSpotTickerPriceService.Symbols>)
- [type GetUMAccountDetailService](<#GetUMAccountDetailService>)
  - [func \(s \*GetUMAccountDetailService\) Do\(ctx context.Context, opts ...RequestOption\) \(\*UMAccountDetail, error\)](<#GetUMAccountDetailService.Do>)
- [type GetUMAccountService](<#GetUMAccountService>)
  - [func \(s \*GetUMAccountService\) Do\(ctx context.Context, opts ...RequestOption\) \(\*UMAccount, error\)](<#GetUMAccountService.Do>)
- [type GetUserAssetService](<#GetUserAssetService>)
  - [func \(s \*GetUserAssetService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]UserAsset, error\)](<#GetUserAssetService.Do>)
- [type GetWalletBalanceService](<#GetWalletBalanceService>)
  - [func \(s \*GetWalletBalanceService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]WalletBalance, error\)](<#GetWalletBalanceService.Do>)
  - [func \(s \*GetWalletBalanceService\) QuoteAsset\(quoteAsset string\) \*GetWalletBalanceService](<#GetWalletBalanceService.QuoteAsset>)
- [type Income](<#Income>)
- [type LimitFilter](<#LimitFilter>)
- [type LotLimit](<#LotLimit>)
- [type MarkPriceEvent](<#MarkPriceEvent>)
- [type PApiAccountBalance](<#PApiAccountBalance>)
- [type PApiGetAllOrdersService](<#PApiGetAllOrdersService>)
  - [func \(s \*PApiGetAllOrdersService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*SwapOrder, error\)](<#PApiGetAllOrdersService.Do>)
  - [func \(s \*PApiGetAllOrdersService\) EndTime\(endTime int64\) \*PApiGetAllOrdersService](<#PApiGetAllOrdersService.EndTime>)
  - [func \(s \*PApiGetAllOrdersService\) Limit\(limit int\) \*PApiGetAllOrdersService](<#PApiGetAllOrdersService.Limit>)
  - [func \(s \*PApiGetAllOrdersService\) OrderId\(orderId string\) \*PApiGetAllOrdersService](<#PApiGetAllOrdersService.OrderId>)
  - [func \(s \*PApiGetAllOrdersService\) StartTime\(startTime int64\) \*PApiGetAllOrdersService](<#PApiGetAllOrdersService.StartTime>)
- [type PApiGetOpenOrdersService](<#PApiGetOpenOrdersService>)
  - [func \(s \*PApiGetOpenOrdersService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*SwapOpenOrder, error\)](<#PApiGetOpenOrdersService.Do>)
  - [func \(s \*PApiGetOpenOrdersService\) Symbol\(symbol string\) \*PApiGetOpenOrdersService](<#PApiGetOpenOrdersService.Symbol>)
- [type PApiGetUMPositionsService](<#PApiGetUMPositionsService>)
  - [func \(s \*PApiGetUMPositionsService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*Position, error\)](<#PApiGetUMPositionsService.Do>)
- [type PApiUmCancelAllOpenOrdersService](<#PApiUmCancelAllOpenOrdersService>)
  - [func \(s \*PApiUmCancelAllOpenOrdersService\) Do\(ctx context.Context, opts ...RequestOption\) error](<#PApiUmCancelAllOpenOrdersService.Do>)
- [type PApiUmCancelOrderService](<#PApiUmCancelOrderService>)
  - [func \(s \*PApiUmCancelOrderService\) Do\(ctx context.Context, opts ...RequestOption\) \(\*SwapOrder, error\)](<#PApiUmCancelOrderService.Do>)
  - [func \(s \*PApiUmCancelOrderService\) OrderId\(orderId string\) \*PApiUmCancelOrderService](<#PApiUmCancelOrderService.OrderId>)
  - [func \(s \*PApiUmCancelOrderService\) OrigClientOrderId\(origClientOrderId string\) \*PApiUmCancelOrderService](<#PApiUmCancelOrderService.OrigClientOrderId>)
  - [func \(s \*PApiUmCancelOrderService\) Symbol\(symbol string\) \*PApiUmCancelOrderService](<#PApiUmCancelOrderService.Symbol>)
- [type PApiUmChangeLeverageService](<#PApiUmChangeLeverageService>)
  - [func \(s \*PApiUmChangeLeverageService\) Do\(ctx context.Context, opts ...RequestOption\) error](<#PApiUmChangeLeverageService.Do>)
- [type PApiUmGetIncomeService](<#PApiUmGetIncomeService>)
  - [func \(s \*PApiUmGetIncomeService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]Income, error\)](<#PApiUmGetIncomeService.Do>)
  - [func \(s \*PApiUmGetIncomeService\) EndTime\(endTime int64\) \*PApiUmGetIncomeService](<#PApiUmGetIncomeService.EndTime>)
  - [func \(s \*PApiUmGetIncomeService\) IncomeType\(incomeType string\) \*PApiUmGetIncomeService](<#PApiUmGetIncomeService.IncomeType>)
  - [func \(s \*PApiUmGetIncomeService\) Limit\(limit int\) \*PApiUmGetIncomeService](<#PApiUmGetIncomeService.Limit>)
  - [func \(s \*PApiUmGetIncomeService\) StartTime\(startTime int64\) \*PApiUmGetIncomeService](<#PApiUmGetIncomeService.StartTime>)
  - [func \(s \*PApiUmGetIncomeService\) Symbol\(symbol string\) \*PApiUmGetIncomeService](<#PApiUmGetIncomeService.Symbol>)
- [type PApiUmGetOrderService](<#PApiUmGetOrderService>)
  - [func \(s \*PApiUmGetOrderService\) Do\(ctx context.Context, opts ...RequestOption\) \(\*SwapOrder, error\)](<#PApiUmGetOrderService.Do>)
  - [func \(s \*PApiUmGetOrderService\) OrderId\(orderId string\) \*PApiUmGetOrderService](<#PApiUmGetOrderService.OrderId>)
  - [func \(s \*PApiUmGetOrderService\) OrigClientOrderId\(origClientOrderId string\) \*PApiUmGetOrderService](<#PApiUmGetOrderService.OrigClientOrderId>)
- [type PApiUmGetUserTradesService](<#PApiUmGetUserTradesService>)
  - [func \(s \*PApiUmGetUserTradesService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*TradeRecord, error\)](<#PApiUmGetUserTradesService.Do>)
  - [func \(s \*PApiUmGetUserTradesService\) EndTime\(endTime int64\) \*PApiUmGetUserTradesService](<#PApiUmGetUserTradesService.EndTime>)
  - [func \(s \*PApiUmGetUserTradesService\) FromId\(fromId int64\) \*PApiUmGetUserTradesService](<#PApiUmGetUserTradesService.FromId>)
  - [func \(s \*PApiUmGetUserTradesService\) Limit\(limit int\) \*PApiUmGetUserTradesService](<#PApiUmGetUserTradesService.Limit>)
  - [func \(s \*PApiUmGetUserTradesService\) StartTime\(startTime int64\) \*PApiUmGetUserTradesService](<#PApiUmGetUserTradesService.StartTime>)
- [type PApiUmOrderService](<#PApiUmOrderService>)
  - [func \(s \*PApiUmOrderService\) Do\(ctx context.Context, opts ...RequestOption\) \(\*SwapOrder, error\)](<#PApiUmOrderService.Do>)
  - [func \(s \*PApiUmOrderService\) GoodTillDate\(goodTillDate int64\) \*PApiUmOrderService](<#PApiUmOrderService.GoodTillDate>)
  - [func \(s \*PApiUmOrderService\) NewClientOrderId\(newClientOrderId string\) \*PApiUmOrderService](<#PApiUmOrderService.NewClientOrderId>)
  - [func \(s \*PApiUmOrderService\) NewOrderRespType\(newOrderRespType string\) \*PApiUmOrderService](<#PApiUmOrderService.NewOrderRespType>)
  - [func \(s \*PApiUmOrderService\) OrderType\(orderType string\) \*PApiUmOrderService](<#PApiUmOrderService.OrderType>)
  - [func \(s \*PApiUmOrderService\) PositionSide\(positionSide string\) \*PApiUmOrderService](<#PApiUmOrderService.PositionSide>)
  - [func \(s \*PApiUmOrderService\) Price\(price string\) \*PApiUmOrderService](<#PApiUmOrderService.Price>)
  - [func \(s \*PApiUmOrderService\) PriceMatch\(priceMatch string\) \*PApiUmOrderService](<#PApiUmOrderService.PriceMatch>)
  - [func \(s \*PApiUmOrderService\) Quantity\(quantity string\) \*PApiUmOrderService](<#PApiUmOrderService.Quantity>)
  - [func \(s \*PApiUmOrderService\) ReduceOnly\(reduceOnly bool\) \*PApiUmOrderService](<#PApiUmOrderService.ReduceOnly>)
  - [func \(s \*PApiUmOrderService\) SelfTradePreventionMode\(selfTradePreventionMode string\) \*PApiUmOrderService](<#PApiUmOrderService.SelfTradePreventionMode>)
  - [func \(s \*PApiUmOrderService\) TimeInForce\(timeInForce string\) \*PApiUmOrderService](<#PApiUmOrderService.TimeInForce>)
- [type PapiUMGetPositionRiskService](<#PapiUMGetPositionRiskService>)
  - [func \(s \*PapiUMGetPositionRiskService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]UMPosition, error\)](<#PapiUMGetPositionRiskService.Do>)
  - [func \(s \*PapiUMGetPositionRiskService\) Symbol\(symbol string\) \*PapiUMGetPositionRiskService](<#PapiUMGetPositionRiskService.Symbol>)
- [type PingService](<#PingService>)
  - [func \(s \*PingService\) Do\(ctx context.Context, opts ...RequestOption\) error](<#PingService.Do>)
- [type Position](<#Position>)
- [type PositionsGetter](<#PositionsGetter>)
- [type PostService](<#PostService>)
  - [func \(s \*PostService\) Do\(ctx context.Context, dest any, opts ...RequestOption\) error](<#PostService.Do>)
- [type PremiumIndex](<#PremiumIndex>)
- [type RequestOption](<#RequestOption>)
  - [func WithHeader\(key, value string, replace bool\) RequestOption](<#WithHeader>)
  - [func WithHeaders\(header http.Header\) RequestOption](<#WithHeaders>)
  - [func WithRecvWindow\(recvWindow int64\) RequestOption](<#WithRecvWindow>)
- [type SpotOpenOrder](<#SpotOpenOrder>)
- [type SpotOrder](<#SpotOrder>)
- [type SpotSymbolInfos](<#SpotSymbolInfos>)
- [type SpotTicker24H](<#SpotTicker24H>)
- [type SwapOpenOrder](<#SwapOpenOrder>)
- [type SwapOpenOrdersGetter](<#SwapOpenOrdersGetter>)
- [type SwapOrder](<#SwapOrder>)
- [type SwapOrdersGetter](<#SwapOrdersGetter>)
- [type SymbolInfo](<#SymbolInfo>)
- [type TickerPrice](<#TickerPrice>)
- [type TimeInForceType](<#TimeInForceType>)
- [type TradeRecord](<#TradeRecord>)
- [type UMAccount](<#UMAccount>)
- [type UMAccountDetail](<#UMAccountDetail>)
- [type UMAsset](<#UMAsset>)
- [type UMPosition](<#UMPosition>)
- [type UserAsset](<#UserAsset>)
- [type UserDataEventType](<#UserDataEventType>)
- [type WalletBalance](<#WalletBalance>)
- [type WsAllMarkPriceEvent](<#WsAllMarkPriceEvent>)
- [type WsAllMiniTickerEvent](<#WsAllMiniTickerEvent>)
- [type WsAllTickerEvent](<#WsAllTickerEvent>)
- [type WsBookTickerEvent](<#WsBookTickerEvent>)
- [type WsMiniTickerEvent](<#WsMiniTickerEvent>)
- [type WsPublicBaseService](<#WsPublicBaseService>)
  - [func \(s \*WsPublicBaseService\) IncrIdx\(\) int64](<#WsPublicBaseService.IncrIdx>)
  - [func \(s \*WsPublicBaseService\) SetHttpClient\(httpClient \*http.Client\) \*WsPublicBaseService](<#WsPublicBaseService.SetHttpClient>)
  - [func \(s \*WsPublicBaseService\) SetLogger\(logIn func\(int, \[\]byte\), logOut func\(int, \[\]byte\)\) \*WsPublicBaseService](<#WsPublicBaseService.SetLogger>)
  - [func \(s \*WsPublicBaseService\) Start\(\) error](<#WsPublicBaseService.Start>)
  - [func \(s \*WsPublicBaseService\) Stop\(\)](<#WsPublicBaseService.Stop>)
  - [func \(s \*WsPublicBaseService\) Subscribe\(channel string, handler ws.WsHandler\) \*WsPublicBaseService](<#WsPublicBaseService.Subscribe>)
  - [func \(s \*WsPublicBaseService\) Unsubscribe\(channel string\) \*WsPublicBaseService](<#WsPublicBaseService.Unsubscribe>)
- [type WsSpotPublicService](<#WsSpotPublicService>)
  - [func NewWsSpotPublicService\(baseUrl string\) WsSpotPublicService](<#NewWsSpotPublicService>)
  - [func \(s WsSpotPublicService\) SubscribeAllMiniTicker\(handler func\(event WsAllMiniTickerEvent\)\)](<#WsSpotPublicService.SubscribeAllMiniTicker>)
  - [func \(s WsSpotPublicService\) SubscribeAllTicker\(handler func\(event WsAllTickerEvent\)\)](<#WsSpotPublicService.SubscribeAllTicker>)
  - [func \(s WsSpotPublicService\) SubscribeSymbolsBookTicker\(handler func\(event WsBookTickerEvent\), symbols ...string\)](<#WsSpotPublicService.SubscribeSymbolsBookTicker>)
  - [func \(s WsSpotPublicService\) SubscribeSymbolsMiniTicker\(handler func\(event WsMiniTickerEvent\), symbols ...string\)](<#WsSpotPublicService.SubscribeSymbolsMiniTicker>)
  - [func \(s WsSpotPublicService\) SubscribeSymbolsTicker\(handler func\(event WsTickerEvent\), symbols ...string\)](<#WsSpotPublicService.SubscribeSymbolsTicker>)
- [type WsSwapPublicService](<#WsSwapPublicService>)
  - [func NewWsSwapPublicService\(baseUrl string\) WsSwapPublicService](<#NewWsSwapPublicService>)
  - [func \(s WsSwapPublicService\) SubscribeAllMarkPrice\(handler func\(event WsAllMarkPriceEvent\)\)](<#WsSwapPublicService.SubscribeAllMarkPrice>)
  - [func \(s WsSwapPublicService\) SubscribeAllMiniTicker\(handler func\(event WsAllMiniTickerEvent\)\)](<#WsSwapPublicService.SubscribeAllMiniTicker>)
  - [func \(s WsSwapPublicService\) SubscribeAllTicker\(handler func\(event WsAllTickerEvent\)\)](<#WsSwapPublicService.SubscribeAllTicker>)
  - [func \(s WsSwapPublicService\) SubscribeBookTicker\(handler func\(event WsBookTickerEvent\), symbols ...string\)](<#WsSwapPublicService.SubscribeBookTicker>)
  - [func \(s WsSwapPublicService\) SubscribeMarkPrice\(handler func\(event MarkPriceEvent\), symbols ...string\)](<#WsSwapPublicService.SubscribeMarkPrice>)
  - [func \(s WsSwapPublicService\) SubscribeMarkPrice1s\(handler func\(event MarkPriceEvent\), symbols ...string\)](<#WsSwapPublicService.SubscribeMarkPrice1s>)
  - [func \(s WsSwapPublicService\) SubscribeSymbolsMiniTicker\(handler func\(event WsMiniTickerEvent\), symbols ...string\)](<#WsSwapPublicService.SubscribeSymbolsMiniTicker>)
  - [func \(s WsSwapPublicService\) SubscribeSymbolsTicker\(handler func\(event WsTickerEvent\), symbols ...string\)](<#WsSwapPublicService.SubscribeSymbolsTicker>)
- [type WsTickerEvent](<#WsTickerEvent>)


## Constants

<a name="UserDataEventTypeFutureOrderTradeUpdate"></a>

```go
const (
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
```

## Variables

<a name="BaseSpotApiMainURL"></a>

```go
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
```

<a name="AsProxyError"></a>
## func AsProxyError

```go
func AsProxyError(e error) (error, bool)
```

AsProxyError convert APIError to ProxyError

<a name="ConvertProxyError"></a>
## func ConvertProxyError

```go
func ConvertProxyError(e error) error
```

ConvertProxyError convert error to ProxyError

<a name="IsAPIError"></a>
## func IsAPIError

```go
func IsAPIError(e error) bool
```

IsAPIError check if e is an API error

<a name="APIError"></a>
## type APIError

APIError define API error when response status is 4xx or 5xx

```go
type APIError struct {
    Code     int64  `json:"code"`
    Message  string `json:"msg"`
    Response []byte `json:"-"` // Assign the body value when the Code and Message fields are invalid.
}
```

<a name="APIError.Error"></a>
### func \(APIError\) Error

```go
func (e APIError) Error() string
```

Error return error code and message

<a name="APIError.IsValid"></a>
### func \(APIError\) IsValid

```go
func (e APIError) IsValid() bool
```



<a name="CapitalDepositHistory"></a>
## type CapitalDepositHistory



```go
type CapitalDepositHistory struct {
    ID            string `json:"id"`
    Amount        string `json:"amount"`
    Coin          string `json:"coin"`
    Network       string `json:"network"`
    Status        int    `json:"status"`
    Address       string `json:"address"`
    AddressTag    string `json:"addressTag"`
    TxID          string `json:"txId"`
    InsertTime    int64  `json:"insertTime"`
    TransferType  int    `json:"transferType"`
    ConfirmTimes  string `json:"confirmTimes"`
    UnlockConfirm int    `json:"unlockConfirm"`
    WalletType    int    `json:"walletType"`
}
```

<a name="CapitalDepositHistoryService"></a>
## type CapitalDepositHistoryService

CapitalDepositHistoryService 获取资金充值历史

```go
type CapitalDepositHistoryService struct {
    // contains filtered or unexported fields
}
```

<a name="CapitalDepositHistoryService.Coin"></a>
### func \(\*CapitalDepositHistoryService\) Coin

```go
func (s *CapitalDepositHistoryService) Coin(coin string) *CapitalDepositHistoryService
```



<a name="CapitalDepositHistoryService.Do"></a>
### func \(\*CapitalDepositHistoryService\) Do

```go
func (s *CapitalDepositHistoryService) Do(ctx context.Context, opts ...RequestOption) ([]CapitalDepositHistory, error)
```



<a name="CapitalDepositHistoryService.EndTime"></a>
### func \(\*CapitalDepositHistoryService\) EndTime

```go
func (s *CapitalDepositHistoryService) EndTime(endTime int64) *CapitalDepositHistoryService
```



<a name="CapitalDepositHistoryService.IncludeSource"></a>
### func \(\*CapitalDepositHistoryService\) IncludeSource

```go
func (s *CapitalDepositHistoryService) IncludeSource(includeSource bool) *CapitalDepositHistoryService
```



<a name="CapitalDepositHistoryService.Limit"></a>
### func \(\*CapitalDepositHistoryService\) Limit

```go
func (s *CapitalDepositHistoryService) Limit(limit int) *CapitalDepositHistoryService
```



<a name="CapitalDepositHistoryService.Offset"></a>
### func \(\*CapitalDepositHistoryService\) Offset

```go
func (s *CapitalDepositHistoryService) Offset(offset int) *CapitalDepositHistoryService
```



<a name="CapitalDepositHistoryService.StartTime"></a>
### func \(\*CapitalDepositHistoryService\) StartTime

```go
func (s *CapitalDepositHistoryService) StartTime(startTime int64) *CapitalDepositHistoryService
```



<a name="CapitalDepositHistoryService.Status"></a>
### func \(\*CapitalDepositHistoryService\) Status

```go
func (s *CapitalDepositHistoryService) Status(status int) *CapitalDepositHistoryService
```



<a name="CapitalDepositHistoryService.TxId"></a>
### func \(\*CapitalDepositHistoryService\) TxId

```go
func (s *CapitalDepositHistoryService) TxId(txId string) *CapitalDepositHistoryService
```



<a name="CapitalWithdrawHistory"></a>
## type CapitalWithdrawHistory



```go
type CapitalWithdrawHistory struct {
    ID              string `json:"id"`              // 该笔提现在币安的id
    Amount          string `json:"amount"`          // 提现转出金额
    TransactionFee  string `json:"transactionFee"`  // 手续费
    Coin            string `json:"coin"`            // 币种
    Status          int    `json:"status"`          // 状态
    Address         string `json:"address"`         // 提现地址
    TxID            string `json:"txId"`            // 提现交易id
    ApplyTime       string `json:"applyTime"`       // UTC 时间
    Network         string `json:"network"`         // 网络
    TransferType    int    `json:"transferType"`    // 1: 站内转账, 0: 站外转账
    WithdrawOrderID string `json:"withdrawOrderId"` // 自定义ID, 如果没有则不返回该字段
    Info            string `json:"info"`            // 提币失败原因
    ConfirmNo       int    `json:"confirmNo"`       // 提现确认数
    WalletType      int    `json:"walletType"`      // 1: 资金钱包 0:现货钱包
    TxKey           string `json:"txKey"`           // 交易key
    CompleteTime    string `json:"completeTime"`    // 提现完成，成功下账时间(UTC)
}
```

<a name="CapitalWithdrawHistoryService"></a>
## type CapitalWithdrawHistoryService

CapitalWithdrawHistoryService 获取资金提现历史

```go
type CapitalWithdrawHistoryService struct {
    // contains filtered or unexported fields
}
```

<a name="CapitalWithdrawHistoryService.Coin"></a>
### func \(\*CapitalWithdrawHistoryService\) Coin

```go
func (s *CapitalWithdrawHistoryService) Coin(coin string) *CapitalWithdrawHistoryService
```



<a name="CapitalWithdrawHistoryService.Do"></a>
### func \(\*CapitalWithdrawHistoryService\) Do

```go
func (s *CapitalWithdrawHistoryService) Do(ctx context.Context, opts ...RequestOption) ([]CapitalWithdrawHistory, error)
```



<a name="CapitalWithdrawHistoryService.EndTime"></a>
### func \(\*CapitalWithdrawHistoryService\) EndTime

```go
func (s *CapitalWithdrawHistoryService) EndTime(endTime int64) *CapitalWithdrawHistoryService
```



<a name="CapitalWithdrawHistoryService.IdList"></a>
### func \(\*CapitalWithdrawHistoryService\) IdList

```go
func (s *CapitalWithdrawHistoryService) IdList(idList []string) *CapitalWithdrawHistoryService
```



<a name="CapitalWithdrawHistoryService.Limit"></a>
### func \(\*CapitalWithdrawHistoryService\) Limit

```go
func (s *CapitalWithdrawHistoryService) Limit(limit int) *CapitalWithdrawHistoryService
```



<a name="CapitalWithdrawHistoryService.Offset"></a>
### func \(\*CapitalWithdrawHistoryService\) Offset

```go
func (s *CapitalWithdrawHistoryService) Offset(offset int) *CapitalWithdrawHistoryService
```



<a name="CapitalWithdrawHistoryService.StartTime"></a>
### func \(\*CapitalWithdrawHistoryService\) StartTime

```go
func (s *CapitalWithdrawHistoryService) StartTime(startTime int64) *CapitalWithdrawHistoryService
```



<a name="CapitalWithdrawHistoryService.Status"></a>
### func \(\*CapitalWithdrawHistoryService\) Status

```go
func (s *CapitalWithdrawHistoryService) Status(status int) *CapitalWithdrawHistoryService
```



<a name="CapitalWithdrawHistoryService.WithdrawOrderId"></a>
### func \(\*CapitalWithdrawHistoryService\) WithdrawOrderId

```go
func (s *CapitalWithdrawHistoryService) WithdrawOrderId(withdrawOrderId string) *CapitalWithdrawHistoryService
```



<a name="Client"></a>
## type Client

Client define API client

```go
type Client struct {
    APIKey      string
    SecretKey   string
    PApiBaseURL string
    ApiBaseURL  string
    FApiBaseURL string
    HTTPClient  *http.Client
    Debug       bool
    Logger      *slog.Logger
    TimeOffset  int64
    // contains filtered or unexported fields
}
```

<a name="NewClient"></a>
### func NewClient

```go
func NewClient(apiKey, secretKey string) *Client
```

NewClient initialize an API client instance with API key and secret key. You should always call this function before using this SDK. Services will be created by the form client.NewXXXService\(\).

<a name="NewClientWithHttpClient"></a>
### func NewClientWithHttpClient

```go
func NewClientWithHttpClient(apiKey, secretKey string, httpClient *http.Client) *Client
```



<a name="NewTestClient"></a>
### func NewTestClient

```go
func NewTestClient(apiKey, secretKey string) *Client
```



<a name="Client.NewFApiCancelOrderService"></a>
### func \(\*Client\) NewFApiCancelOrderService

```go
func (c *Client) NewFApiCancelOrderService(symbol string) *FApiCancelOrderService
```



<a name="Client.NewFApiChangeLeverageService"></a>
### func \(\*Client\) NewFApiChangeLeverageService

```go
func (c *Client) NewFApiChangeLeverageService(symbol string, leverage int64) *FApiChangeLeverageService
```



<a name="Client.NewFApiGetAccountBalanceService"></a>
### func \(\*Client\) NewFApiGetAccountBalanceService

```go
func (c *Client) NewFApiGetAccountBalanceService() *FApiGetAccountBalanceService
```



<a name="Client.NewFApiGetAllOpenOrdersService"></a>
### func \(\*Client\) NewFApiGetAllOpenOrdersService

```go
func (c *Client) NewFApiGetAllOpenOrdersService() *FApiGetAllOpenOrdersService
```



<a name="Client.NewFApiGetAllOrdersService"></a>
### func \(\*Client\) NewFApiGetAllOrdersService

```go
func (c *Client) NewFApiGetAllOrdersService(symbol string) *FApiGetAllOrdersService
```



<a name="Client.NewFApiGetBookTickerService"></a>
### func \(\*Client\) NewFApiGetBookTickerService

```go
func (c *Client) NewFApiGetBookTickerService() *FApiGetBookTickerService
```



<a name="Client.NewFApiGetDepthService"></a>
### func \(\*Client\) NewFApiGetDepthService

```go
func (c *Client) NewFApiGetDepthService(symbol string) *FApiGetDepthService
```



<a name="Client.NewFApiGetIncomeService"></a>
### func \(\*Client\) NewFApiGetIncomeService

```go
func (c *Client) NewFApiGetIncomeService() *FApiGetIncomeService
```



<a name="Client.NewFApiGetKLinesService"></a>
### func \(\*Client\) NewFApiGetKLinesService

```go
func (c *Client) NewFApiGetKLinesService(symbol, interval string) *FApiGetKLinesService
```



<a name="Client.NewFApiGetOrderService"></a>
### func \(\*Client\) NewFApiGetOrderService

```go
func (c *Client) NewFApiGetOrderService(symbol string) *FApiGetOrderService
```



<a name="Client.NewFApiGetPositionsService"></a>
### func \(\*Client\) NewFApiGetPositionsService

```go
func (c *Client) NewFApiGetPositionsService() *FApiGetPositionsService
```



<a name="Client.NewFApiGetService"></a>
### func \(\*Client\) NewFApiGetService

```go
func (c *Client) NewFApiGetService(endpoint string, params params) *GetService
```



<a name="Client.NewFApiGetUserTradesService"></a>
### func \(\*Client\) NewFApiGetUserTradesService

```go
func (c *Client) NewFApiGetUserTradesService(symbol string) *FApiGetUserTradesService
```



<a name="Client.NewFApiOrderService"></a>
### func \(\*Client\) NewFApiOrderService

```go
func (c *Client) NewFApiOrderService(symbol string, side string, orderType string) *FApiOrderService
```



<a name="Client.NewFApiPostService"></a>
### func \(\*Client\) NewFApiPostService

```go
func (c *Client) NewFApiPostService(endpoint string, params params) *PostService
```



<a name="Client.NewGetCapitalDepositHistoryService"></a>
### func \(\*Client\) NewGetCapitalDepositHistoryService

```go
func (c *Client) NewGetCapitalDepositHistoryService() *CapitalDepositHistoryService
```



<a name="Client.NewGetCapitalWithdrawHistoryService"></a>
### func \(\*Client\) NewGetCapitalWithdrawHistoryService

```go
func (c *Client) NewGetCapitalWithdrawHistoryService() *CapitalWithdrawHistoryService
```



<a name="Client.NewGetFApiSwapOpenOrdersService"></a>
### func \(\*Client\) NewGetFApiSwapOpenOrdersService

```go
func (c *Client) NewGetFApiSwapOpenOrdersService() *GetFApiSwapOpenOrdersService
```



<a name="Client.NewGetFApiSwapOrdersService"></a>
### func \(\*Client\) NewGetFApiSwapOrdersService

```go
func (c *Client) NewGetFApiSwapOrdersService() *GetFApiSwapOrdersService
```



<a name="Client.NewGetFApiSymbolInfosService"></a>
### func \(\*Client\) NewGetFApiSymbolInfosService

```go
func (c *Client) NewGetFApiSymbolInfosService() *GetFApiSymbolInfosService
```



<a name="Client.NewGetFApiTicker24HService"></a>
### func \(\*Client\) NewGetFApiTicker24HService

```go
func (c *Client) NewGetFApiTicker24HService() *GetFApiTicker24HService
```



<a name="Client.NewGetFApiTickerPriceService"></a>
### func \(\*Client\) NewGetFApiTickerPriceService

```go
func (c *Client) NewGetFApiTickerPriceService() *GetFApiTickerPriceService
```



<a name="Client.NewGetPApiAccountBalanceService"></a>
### func \(\*Client\) NewGetPApiAccountBalanceService

```go
func (c *Client) NewGetPApiAccountBalanceService() *GetPApiAccountBalanceService
```



<a name="Client.NewGetPApiSwapOpenOrdersService"></a>
### func \(\*Client\) NewGetPApiSwapOpenOrdersService

```go
func (c *Client) NewGetPApiSwapOpenOrdersService() *GetPApiSwapOpenOrdersService
```



<a name="Client.NewGetPremiumIndexService"></a>
### func \(\*Client\) NewGetPremiumIndexService

```go
func (c *Client) NewGetPremiumIndexService() *GetPremiumIndexService
```



<a name="Client.NewGetService"></a>
### func \(\*Client\) NewGetService

```go
func (c *Client) NewGetService(baseURL *string, endpoint string, params params) *GetService
```



<a name="Client.NewGetSpotOpenOrdersService"></a>
### func \(\*Client\) NewGetSpotOpenOrdersService

```go
func (c *Client) NewGetSpotOpenOrdersService() *GetSpotOpenOrdersService
```



<a name="Client.NewGetSpotOrdersService"></a>
### func \(\*Client\) NewGetSpotOrdersService

```go
func (c *Client) NewGetSpotOrdersService() *GetSpotOrdersService
```



<a name="Client.NewGetSpotSymbolInfosService"></a>
### func \(\*Client\) NewGetSpotSymbolInfosService

```go
func (c *Client) NewGetSpotSymbolInfosService() *GetSpotSymbolInfosService
```



<a name="Client.NewGetSpotTicker24HService"></a>
### func \(\*Client\) NewGetSpotTicker24HService

```go
func (c *Client) NewGetSpotTicker24HService() *GetSpotTicker24HService
```



<a name="Client.NewGetSpotTickerPriceService"></a>
### func \(\*Client\) NewGetSpotTickerPriceService

```go
func (c *Client) NewGetSpotTickerPriceService() *GetSpotTickerPriceService
```



<a name="Client.NewGetUMAccountDetailService"></a>
### func \(\*Client\) NewGetUMAccountDetailService

```go
func (c *Client) NewGetUMAccountDetailService() *GetUMAccountDetailService
```



<a name="Client.NewGetUMAccountService"></a>
### func \(\*Client\) NewGetUMAccountService

```go
func (c *Client) NewGetUMAccountService() *GetUMAccountService
```



<a name="Client.NewGetUserAssetService"></a>
### func \(\*Client\) NewGetUserAssetService

```go
func (c *Client) NewGetUserAssetService() *GetUserAssetService
```



<a name="Client.NewGetWalletBalanceService"></a>
### func \(\*Client\) NewGetWalletBalanceService

```go
func (c *Client) NewGetWalletBalanceService() *GetWalletBalanceService
```



<a name="Client.NewPApiGetAllOrdersService"></a>
### func \(\*Client\) NewPApiGetAllOrdersService

```go
func (c *Client) NewPApiGetAllOrdersService(symbol string) *PApiGetAllOrdersService
```



<a name="Client.NewPApiGetOpenOrdersService"></a>
### func \(\*Client\) NewPApiGetOpenOrdersService

```go
func (c *Client) NewPApiGetOpenOrdersService() *PApiGetOpenOrdersService
```



<a name="Client.NewPApiGetService"></a>
### func \(\*Client\) NewPApiGetService

```go
func (c *Client) NewPApiGetService(endpoint string, params params) *GetService
```



<a name="Client.NewPApiGetUMPositionsService"></a>
### func \(\*Client\) NewPApiGetUMPositionsService

```go
func (c *Client) NewPApiGetUMPositionsService() *PApiGetUMPositionsService
```



<a name="Client.NewPApiPostService"></a>
### func \(\*Client\) NewPApiPostService

```go
func (c *Client) NewPApiPostService(endpoint string, params params) *PostService
```



<a name="Client.NewPApiUmCancelAllOpenOrdersService"></a>
### func \(\*Client\) NewPApiUmCancelAllOpenOrdersService

```go
func (c *Client) NewPApiUmCancelAllOpenOrdersService(symbol string) *PApiUmCancelAllOpenOrdersService
```



<a name="Client.NewPApiUmCancelOrderService"></a>
### func \(\*Client\) NewPApiUmCancelOrderService

```go
func (c *Client) NewPApiUmCancelOrderService() *PApiUmCancelOrderService
```



<a name="Client.NewPApiUmChangeLeverageService"></a>
### func \(\*Client\) NewPApiUmChangeLeverageService

```go
func (c *Client) NewPApiUmChangeLeverageService(symbol string, leverage int64) *PApiUmChangeLeverageService
```



<a name="Client.NewPApiUmGetIncomeService"></a>
### func \(\*Client\) NewPApiUmGetIncomeService

```go
func (c *Client) NewPApiUmGetIncomeService() *PApiUmGetIncomeService
```



<a name="Client.NewPApiUmGetOrderService"></a>
### func \(\*Client\) NewPApiUmGetOrderService

```go
func (c *Client) NewPApiUmGetOrderService(symbol string) *PApiUmGetOrderService
```



<a name="Client.NewPApiUmGetUserTradesService"></a>
### func \(\*Client\) NewPApiUmGetUserTradesService

```go
func (c *Client) NewPApiUmGetUserTradesService(symbol string) *PApiUmGetUserTradesService
```



<a name="Client.NewPApiUmOrderService"></a>
### func \(\*Client\) NewPApiUmOrderService

```go
func (c *Client) NewPApiUmOrderService(symbol string, side string, orderType string) *PApiUmOrderService
```



<a name="Client.NewPapiUMGetPositionRiskService"></a>
### func \(\*Client\) NewPapiUMGetPositionRiskService

```go
func (c *Client) NewPapiUMGetPositionRiskService() *PapiUMGetPositionRiskService
```



<a name="Client.NewPingService"></a>
### func \(\*Client\) NewPingService

```go
func (c *Client) NewPingService() *PingService
```



<a name="Client.NewPostService"></a>
### func \(\*Client\) NewPostService

```go
func (c *Client) NewPostService(baseURL *string, endpoint string, params params) *PostService
```



<a name="Client.NewSpotGetService"></a>
### func \(\*Client\) NewSpotGetService

```go
func (c *Client) NewSpotGetService(endpoint string, params params) *GetService
```



<a name="Client.NewSpotPostService"></a>
### func \(\*Client\) NewSpotPostService

```go
func (c *Client) NewSpotPostService(endpoint string, params params) *PostService
```



<a name="Client.SetProxyURL"></a>
### func \(\*Client\) SetProxyURL

```go
func (c *Client) SetProxyURL(proxyURL *url.URL)
```



<a name="Client.WithHttpClient"></a>
### func \(\*Client\) WithHttpClient

```go
func (c *Client) WithHttpClient(httpClient *http.Client) *Client
```



<a name="Depth"></a>
## type Depth



```go
type Depth struct {
    LastUpdateId Int64        `json:"lastUpdateId"` // 最后更新ID
    EventTime    Int64        `json:"E"`            // 事件时间
    UpdateTime   Int64        `json:"T"`            // 撮合更新时间
    Asks         [][2]Float64 `json:"asks"`         // 卖盘
    Bids         [][2]Float64 `json:"bids"`         // 买盘
}
```

<a name="FApiAccountBalance"></a>
## type FApiAccountBalance



```go
type FApiAccountBalance struct {
    AccountAlias       string  `json:"accountAlias"`       // 账户唯一识别码
    Asset              string  `json:"asset"`              // 资产
    Balance            Float64 `json:"balance"`            // 总余额
    CrossWalletBalance Float64 `json:"crossWalletBalance"` // 全仓余额
    CrossUnPnl         Float64 `json:"crossUnPnl"`         // 全仓持仓未实现盈亏
    AvailableBalance   Float64 `json:"availableBalance"`   // 下单可用余额
    MaxWithdrawAmount  string  `json:"maxWithdrawAmount"`  // 最大可转出余额
    MarginAvailable    bool    `json:"marginAvailable"`    // 是否可用作联合保证金
    UpdateTime         int64   `json:"updateTime"`         // 更新时间
}
```

<a name="FApiBookTicker"></a>
## type FApiBookTicker



```go
type FApiBookTicker struct {
    Symbol   string  `json:"symbol"`   // 交易对
    BidPrice Float64 `json:"bidPrice"` // 最高买价
    BidQty   Float64 `json:"bidQty"`   // 最高买价挂单量
    AskPrice Float64 `json:"askPrice"` // 最低卖价
    AskQty   Float64 `json:"askQty"`   // 最低卖价挂单量
    Time     Int64   `json:"time"`     // 更新时间
}
```

<a name="FApiCancelOrderService"></a>
## type FApiCancelOrderService

取消订单 https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/rest-api

```go
type FApiCancelOrderService struct {
    // contains filtered or unexported fields
}
```

<a name="FApiCancelOrderService.Do"></a>
### func \(\*FApiCancelOrderService\) Do

```go
func (s *FApiCancelOrderService) Do(ctx context.Context, opts ...RequestOption) (*SwapOrder, error)
```



<a name="FApiCancelOrderService.OrderId"></a>
### func \(\*FApiCancelOrderService\) OrderId

```go
func (s *FApiCancelOrderService) OrderId(orderId string) *FApiCancelOrderService
```



<a name="FApiCancelOrderService.OrigClientOrderId"></a>
### func \(\*FApiCancelOrderService\) OrigClientOrderId

```go
func (s *FApiCancelOrderService) OrigClientOrderId(origClientOrderId string) *FApiCancelOrderService
```



<a name="FApiChangeLeverageService"></a>
## type FApiChangeLeverageService

调整开仓杠杆 https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/rest-api/Change-Initial-Leverage

```go
type FApiChangeLeverageService struct {
    // contains filtered or unexported fields
}
```

<a name="FApiChangeLeverageService.Do"></a>
### func \(\*FApiChangeLeverageService\) Do

```go
func (s *FApiChangeLeverageService) Do(ctx context.Context, opts ...RequestOption) error
```



<a name="FApiGetAccountBalanceService"></a>
## type FApiGetAccountBalanceService



```go
type FApiGetAccountBalanceService struct {
    // contains filtered or unexported fields
}
```

<a name="FApiGetAccountBalanceService.Do"></a>
### func \(\*FApiGetAccountBalanceService\) Do

```go
func (s *FApiGetAccountBalanceService) Do(ctx context.Context, opts ...RequestOption) ([]*FApiAccountBalance, error)
```



<a name="FApiGetAllOpenOrdersService"></a>
## type FApiGetAllOpenOrdersService

查询当前所有挂单 https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/rest-api/Current-All-Open-Orders

```go
type FApiGetAllOpenOrdersService struct {
    // contains filtered or unexported fields
}
```

<a name="FApiGetAllOpenOrdersService.Do"></a>
### func \(\*FApiGetAllOpenOrdersService\) Do

```go
func (s *FApiGetAllOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]*SwapOpenOrder, error)
```



<a name="FApiGetAllOpenOrdersService.Symbol"></a>
### func \(\*FApiGetAllOpenOrdersService\) Symbol

```go
func (s *FApiGetAllOpenOrdersService) Symbol(symbol string) *FApiGetAllOpenOrdersService
```



<a name="FApiGetAllOrdersService"></a>
## type FApiGetAllOrdersService

查询所有订单 https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/rest-api/All-Orders

```go
type FApiGetAllOrdersService struct {
    // contains filtered or unexported fields
}
```

<a name="FApiGetAllOrdersService.Do"></a>
### func \(\*FApiGetAllOrdersService\) Do

```go
func (s *FApiGetAllOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]*SwapOrder, error)
```



<a name="FApiGetAllOrdersService.EndTime"></a>
### func \(\*FApiGetAllOrdersService\) EndTime

```go
func (s *FApiGetAllOrdersService) EndTime(endTime int64) *FApiGetAllOrdersService
```



<a name="FApiGetAllOrdersService.Limit"></a>
### func \(\*FApiGetAllOrdersService\) Limit

```go
func (s *FApiGetAllOrdersService) Limit(limit int) *FApiGetAllOrdersService
```



<a name="FApiGetAllOrdersService.OrderId"></a>
### func \(\*FApiGetAllOrdersService\) OrderId

```go
func (s *FApiGetAllOrdersService) OrderId(orderId string) *FApiGetAllOrdersService
```



<a name="FApiGetAllOrdersService.StartTime"></a>
### func \(\*FApiGetAllOrdersService\) StartTime

```go
func (s *FApiGetAllOrdersService) StartTime(startTime int64) *FApiGetAllOrdersService
```



<a name="FApiGetBookTickerService"></a>
## type FApiGetBookTickerService



```go
type FApiGetBookTickerService struct {
    // contains filtered or unexported fields
}
```

<a name="FApiGetBookTickerService.Do"></a>
### func \(\*FApiGetBookTickerService\) Do

```go
func (s *FApiGetBookTickerService) Do(ctx context.Context, opts ...RequestOption) ([]*FApiBookTicker, error)
```



<a name="FApiGetBookTickerService.Symbol"></a>
### func \(\*FApiGetBookTickerService\) Symbol

```go
func (s *FApiGetBookTickerService) Symbol(symbol string) *FApiGetBookTickerService
```



<a name="FApiGetDepthService"></a>
## type FApiGetDepthService



```go
type FApiGetDepthService struct {
    // contains filtered or unexported fields
}
```

<a name="FApiGetDepthService.Do"></a>
### func \(\*FApiGetDepthService\) Do

```go
func (s *FApiGetDepthService) Do(ctx context.Context, opts ...RequestOption) (*Depth, error)
```



<a name="FApiGetDepthService.Limit"></a>
### func \(\*FApiGetDepthService\) Limit

```go
func (s *FApiGetDepthService) Limit(limit int) *FApiGetDepthService
```



<a name="FApiGetIncomeService"></a>
## type FApiGetIncomeService



```go
type FApiGetIncomeService struct {
    // contains filtered or unexported fields
}
```

<a name="FApiGetIncomeService.Do"></a>
### func \(\*FApiGetIncomeService\) Do

```go
func (s *FApiGetIncomeService) Do(ctx context.Context, opts ...RequestOption) ([]Income, error)
```



<a name="FApiGetIncomeService.EndTime"></a>
### func \(\*FApiGetIncomeService\) EndTime

```go
func (s *FApiGetIncomeService) EndTime(endTime int64) *FApiGetIncomeService
```



<a name="FApiGetIncomeService.IncomeType"></a>
### func \(\*FApiGetIncomeService\) IncomeType

```go
func (s *FApiGetIncomeService) IncomeType(incomeType string) *FApiGetIncomeService
```



<a name="FApiGetIncomeService.Limit"></a>
### func \(\*FApiGetIncomeService\) Limit

```go
func (s *FApiGetIncomeService) Limit(limit int) *FApiGetIncomeService
```



<a name="FApiGetIncomeService.Page"></a>
### func \(\*FApiGetIncomeService\) Page

```go
func (s *FApiGetIncomeService) Page(page int) *FApiGetIncomeService
```



<a name="FApiGetIncomeService.StartTime"></a>
### func \(\*FApiGetIncomeService\) StartTime

```go
func (s *FApiGetIncomeService) StartTime(startTime int64) *FApiGetIncomeService
```



<a name="FApiGetIncomeService.Symbol"></a>
### func \(\*FApiGetIncomeService\) Symbol

```go
func (s *FApiGetIncomeService) Symbol(symbol string) *FApiGetIncomeService
```



<a name="FApiGetKLinesService"></a>
## type FApiGetKLinesService



```go
type FApiGetKLinesService struct {
    // contains filtered or unexported fields
}
```

<a name="FApiGetKLinesService.Do"></a>
### func \(\*FApiGetKLinesService\) Do

```go
func (s *FApiGetKLinesService) Do(ctx context.Context, opts ...RequestOption) ([][]Mixed, error)
```



<a name="FApiGetKLinesService.EndTime"></a>
### func \(\*FApiGetKLinesService\) EndTime

```go
func (s *FApiGetKLinesService) EndTime(endTime int64) *FApiGetKLinesService
```



<a name="FApiGetKLinesService.Limit"></a>
### func \(\*FApiGetKLinesService\) Limit

```go
func (s *FApiGetKLinesService) Limit(limit int) *FApiGetKLinesService
```



<a name="FApiGetKLinesService.StartTime"></a>
### func \(\*FApiGetKLinesService\) StartTime

```go
func (s *FApiGetKLinesService) StartTime(startTime int64) *FApiGetKLinesService
```



<a name="FApiGetOrderService"></a>
## type FApiGetOrderService

查询订单 https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/rest-api/Query-Order

```go
type FApiGetOrderService struct {
    // contains filtered or unexported fields
}
```

<a name="FApiGetOrderService.Do"></a>
### func \(\*FApiGetOrderService\) Do

```go
func (s *FApiGetOrderService) Do(ctx context.Context, opts ...RequestOption) (*SwapOrder, error)
```



<a name="FApiGetOrderService.OrderId"></a>
### func \(\*FApiGetOrderService\) OrderId

```go
func (s *FApiGetOrderService) OrderId(orderId string) *FApiGetOrderService
```



<a name="FApiGetOrderService.OrigClientOrderId"></a>
### func \(\*FApiGetOrderService\) OrigClientOrderId

```go
func (s *FApiGetOrderService) OrigClientOrderId(origClientOrderId string) *FApiGetOrderService
```



<a name="FApiGetPositionsService"></a>
## type FApiGetPositionsService



```go
type FApiGetPositionsService struct {
    // contains filtered or unexported fields
}
```

<a name="FApiGetPositionsService.Do"></a>
### func \(\*FApiGetPositionsService\) Do

```go
func (s *FApiGetPositionsService) Do(ctx context.Context, opts ...RequestOption) ([]*Position, error)
```



<a name="FApiGetUserTradesService"></a>
## type FApiGetUserTradesService



```go
type FApiGetUserTradesService struct {
    // contains filtered or unexported fields
}
```

<a name="FApiGetUserTradesService.Do"></a>
### func \(\*FApiGetUserTradesService\) Do

```go
func (s *FApiGetUserTradesService) Do(ctx context.Context, opts ...RequestOption) ([]*TradeRecord, error)
```



<a name="FApiGetUserTradesService.EndTime"></a>
### func \(\*FApiGetUserTradesService\) EndTime

```go
func (s *FApiGetUserTradesService) EndTime(endTime int64) *FApiGetUserTradesService
```



<a name="FApiGetUserTradesService.FromId"></a>
### func \(\*FApiGetUserTradesService\) FromId

```go
func (s *FApiGetUserTradesService) FromId(fromId int64) *FApiGetUserTradesService
```



<a name="FApiGetUserTradesService.Limit"></a>
### func \(\*FApiGetUserTradesService\) Limit

```go
func (s *FApiGetUserTradesService) Limit(limit int) *FApiGetUserTradesService
```



<a name="FApiGetUserTradesService.OrderId"></a>
### func \(\*FApiGetUserTradesService\) OrderId

```go
func (s *FApiGetUserTradesService) OrderId(orderId string) *FApiGetUserTradesService
```



<a name="FApiGetUserTradesService.StartTime"></a>
### func \(\*FApiGetUserTradesService\) StartTime

```go
func (s *FApiGetUserTradesService) StartTime(startTime int64) *FApiGetUserTradesService
```



<a name="FApiOrderService"></a>
## type FApiOrderService

https://developers.binance.com/docs/zh-CN/derivatives/usds-margined-futures/trade/rest-api

```go
type FApiOrderService struct {
    // contains filtered or unexported fields
}
```

<a name="FApiOrderService.ActivationPrice"></a>
### func \(\*FApiOrderService\) ActivationPrice

```go
func (s *FApiOrderService) ActivationPrice(activationPrice string) *FApiOrderService
```



<a name="FApiOrderService.CallbackRate"></a>
### func \(\*FApiOrderService\) CallbackRate

```go
func (s *FApiOrderService) CallbackRate(callbackRate string) *FApiOrderService
```



<a name="FApiOrderService.ClosePosition"></a>
### func \(\*FApiOrderService\) ClosePosition

```go
func (s *FApiOrderService) ClosePosition(closePosition string) *FApiOrderService
```



<a name="FApiOrderService.Do"></a>
### func \(\*FApiOrderService\) Do

```go
func (s *FApiOrderService) Do(ctx context.Context, opts ...RequestOption) (*SwapOrder, error)
```



<a name="FApiOrderService.GoodTillDate"></a>
### func \(\*FApiOrderService\) GoodTillDate

```go
func (s *FApiOrderService) GoodTillDate(goodTillDate int64) *FApiOrderService
```



<a name="FApiOrderService.NewClientOrderId"></a>
### func \(\*FApiOrderService\) NewClientOrderId

```go
func (s *FApiOrderService) NewClientOrderId(newClientOrderId string) *FApiOrderService
```



<a name="FApiOrderService.NewOrderRespType"></a>
### func \(\*FApiOrderService\) NewOrderRespType

```go
func (s *FApiOrderService) NewOrderRespType(newOrderRespType string) *FApiOrderService
```



<a name="FApiOrderService.OrderType"></a>
### func \(\*FApiOrderService\) OrderType

```go
func (s *FApiOrderService) OrderType(orderType string) *FApiOrderService
```



<a name="FApiOrderService.PositionSide"></a>
### func \(\*FApiOrderService\) PositionSide

```go
func (s *FApiOrderService) PositionSide(positionSide string) *FApiOrderService
```



<a name="FApiOrderService.Price"></a>
### func \(\*FApiOrderService\) Price

```go
func (s *FApiOrderService) Price(price string) *FApiOrderService
```



<a name="FApiOrderService.PriceMatch"></a>
### func \(\*FApiOrderService\) PriceMatch

```go
func (s *FApiOrderService) PriceMatch(priceMatch string) *FApiOrderService
```



<a name="FApiOrderService.PriceProtect"></a>
### func \(\*FApiOrderService\) PriceProtect

```go
func (s *FApiOrderService) PriceProtect(priceProtect bool) *FApiOrderService
```



<a name="FApiOrderService.Quantity"></a>
### func \(\*FApiOrderService\) Quantity

```go
func (s *FApiOrderService) Quantity(quantity string) *FApiOrderService
```



<a name="FApiOrderService.ReduceOnly"></a>
### func \(\*FApiOrderService\) ReduceOnly

```go
func (s *FApiOrderService) ReduceOnly(reduceOnly bool) *FApiOrderService
```



<a name="FApiOrderService.SelfTradePreventionMode"></a>
### func \(\*FApiOrderService\) SelfTradePreventionMode

```go
func (s *FApiOrderService) SelfTradePreventionMode(selfTradePreventionMode string) *FApiOrderService
```



<a name="FApiOrderService.StopPrice"></a>
### func \(\*FApiOrderService\) StopPrice

```go
func (s *FApiOrderService) StopPrice(stopPrice string) *FApiOrderService
```



<a name="FApiOrderService.TimeInForce"></a>
### func \(\*FApiOrderService\) TimeInForce

```go
func (s *FApiOrderService) TimeInForce(timeInForce string) *FApiOrderService
```



<a name="FApiOrderService.WorkingType"></a>
### func \(\*FApiOrderService\) WorkingType

```go
func (s *FApiOrderService) WorkingType(workingType string) *FApiOrderService
```



<a name="FApiSymbolInfos"></a>
## type FApiSymbolInfos



```go
type FApiSymbolInfos struct {
    Symbols []struct {
        Symbol              string            `json:"symbol"`              // 交易对
        QuoteAsset          string            `json:"quoteAsset"`          // 报价币种
        ContractType        string            `json:"contractType"`        // 合约类型
        Status              string            `json:"status"`              // 状态
        PricePrecision      int               `json:"pricePrecision"`      // 价格精度
        QuantityPrecision   int               `json:"quantityPrecision"`   // 数量精度
        BaseAssetPrecision  int               `json:"baseAssetPrecision"`  // 基础货币精度
        QuoteAssetPrecision int               `json:"quoteAssetPrecision"` // 报价货币精度
        Filters             []json.RawMessage `json:"filters"`             // 过滤器
    } `json:"symbols"`
}
```

<a name="FApiTicker24H"></a>
## type FApiTicker24H



```go
type FApiTicker24H struct {
    Symbol             string  `json:"symbol"`
    PriceChange        Float64 `json:"priceChange"`        // 24小时价格变动
    PriceChangePercent Float64 `json:"priceChangePercent"` // 24小时价格变动百分比
    WeightedAvgPrice   Float64 `json:"weightedAvgPrice"`   // 加权平均价
    LastPrice          Float64 `json:"lastPrice"`          // 最近一次成交价
    LastQty            Float64 `json:"lastQty"`            // 最近一次成交额
    OpenPrice          Float64 `json:"openPrice"`          // 24小时内第一次成交的价格
    HighPrice          Float64 `json:"highPrice"`          // 24小时最高价
    LowPrice           Float64 `json:"lowPrice"`           // 24小时最低价
    Volume             Float64 `json:"volume"`             // 24小时成交量
    QuoteVolume        Float64 `json:"quoteVolume"`        // 24小时成交额
    OpenTime           Int64   `json:"openTime"`           // 24小时内，第一笔交易的发生时间
    CloseTime          Int64   `json:"closeTime"`          // 24小时内，最后一笔交易的发生时间
    FirstId            Int64   `json:"firstId"`            // 首笔成交id
    LastId             Int64   `json:"lastId"`             // 末笔成交id
    Count              Int64   `json:"count"`              // 成交笔数
}
```

<a name="FutureSubtype"></a>
## type FutureSubtype



```go
type FutureSubtype string
```

<a name="GetFApiSwapOpenOrdersService"></a>
## type GetFApiSwapOpenOrdersService



```go
type GetFApiSwapOpenOrdersService struct {
    // contains filtered or unexported fields
}
```

<a name="GetFApiSwapOpenOrdersService.Do"></a>
### func \(\*GetFApiSwapOpenOrdersService\) Do

```go
func (s *GetFApiSwapOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]*SwapOpenOrder, error)
```



<a name="GetFApiSwapOrdersService"></a>
## type GetFApiSwapOrdersService



```go
type GetFApiSwapOrdersService struct {
    // contains filtered or unexported fields
}
```

<a name="GetFApiSwapOrdersService.Do"></a>
### func \(\*GetFApiSwapOrdersService\) Do

```go
func (s *GetFApiSwapOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]*SwapOrder, error)
```



<a name="GetFApiSwapOrdersService.EndTime"></a>
### func \(\*GetFApiSwapOrdersService\) EndTime

```go
func (s *GetFApiSwapOrdersService) EndTime(endTime int64) SwapOrdersGetter
```



<a name="GetFApiSwapOrdersService.Limit"></a>
### func \(\*GetFApiSwapOrdersService\) Limit

```go
func (s *GetFApiSwapOrdersService) Limit(limit int) SwapOrdersGetter
```



<a name="GetFApiSwapOrdersService.OrderId"></a>
### func \(\*GetFApiSwapOrdersService\) OrderId

```go
func (s *GetFApiSwapOrdersService) OrderId(orderId string) SwapOrdersGetter
```



<a name="GetFApiSwapOrdersService.StartTime"></a>
### func \(\*GetFApiSwapOrdersService\) StartTime

```go
func (s *GetFApiSwapOrdersService) StartTime(startTime int64) SwapOrdersGetter
```



<a name="GetFApiSwapOrdersService.Symbol"></a>
### func \(\*GetFApiSwapOrdersService\) Symbol

```go
func (s *GetFApiSwapOrdersService) Symbol(symbol string) SwapOrdersGetter
```



<a name="GetFApiSymbolInfosService"></a>
## type GetFApiSymbolInfosService

GetFApiSymbolInfosService 获取永续合约交易对信息

```go
type GetFApiSymbolInfosService struct {
    // contains filtered or unexported fields
}
```

<a name="GetFApiSymbolInfosService.Do"></a>
### func \(\*GetFApiSymbolInfosService\) Do

```go
func (s *GetFApiSymbolInfosService) Do(ctx context.Context, opts ...RequestOption) ([]*SymbolInfo, error)
```



<a name="GetFApiTicker24HService"></a>
## type GetFApiTicker24HService



```go
type GetFApiTicker24HService struct {
    // contains filtered or unexported fields
}
```

<a name="GetFApiTicker24HService.Do"></a>
### func \(\*GetFApiTicker24HService\) Do

```go
func (s *GetFApiTicker24HService) Do(ctx context.Context, opts ...RequestOption) ([]*FApiTicker24H, error)
```



<a name="GetFApiTicker24HService.Symbol"></a>
### func \(\*GetFApiTicker24HService\) Symbol

```go
func (s *GetFApiTicker24HService) Symbol(symbol string) *GetFApiTicker24HService
```



<a name="GetFApiTickerPriceService"></a>
## type GetFApiTickerPriceService

GetFApiTickerPriceService 获取永续合约交易对价格

```go
type GetFApiTickerPriceService struct {
    // contains filtered or unexported fields
}
```

<a name="GetFApiTickerPriceService.Do"></a>
### func \(\*GetFApiTickerPriceService\) Do

```go
func (s *GetFApiTickerPriceService) Do(ctx context.Context, opts ...RequestOption) ([]*TickerPrice, error)
```



<a name="GetFApiTickerPriceService.Symbol"></a>
### func \(\*GetFApiTickerPriceService\) Symbol

```go
func (s *GetFApiTickerPriceService) Symbol(symbol string) *GetFApiTickerPriceService
```



<a name="GetPApiAccountBalanceService"></a>
## type GetPApiAccountBalanceService

GetPApiAccountBalanceService 获取账户余额 https://developers.binance.com/docs/zh-CN/derivatives/portfolio-margin/account

```go
type GetPApiAccountBalanceService struct {
    // contains filtered or unexported fields
}
```

<a name="GetPApiAccountBalanceService.Asset"></a>
### func \(\*GetPApiAccountBalanceService\) Asset

```go
func (s *GetPApiAccountBalanceService) Asset(asset string) *GetPApiAccountBalanceService
```



<a name="GetPApiAccountBalanceService.Do"></a>
### func \(\*GetPApiAccountBalanceService\) Do

```go
func (s *GetPApiAccountBalanceService) Do(ctx context.Context, opts ...RequestOption) ([]*PApiAccountBalance, error)
```



<a name="GetPApiSwapOpenOrdersService"></a>
## type GetPApiSwapOpenOrdersService



```go
type GetPApiSwapOpenOrdersService struct {
    // contains filtered or unexported fields
}
```

<a name="GetPApiSwapOpenOrdersService.Do"></a>
### func \(\*GetPApiSwapOpenOrdersService\) Do

```go
func (s *GetPApiSwapOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]*SwapOpenOrder, error)
```



<a name="GetPApiSwapOpenOrdersService.Symbol"></a>
### func \(\*GetPApiSwapOpenOrdersService\) Symbol

```go
func (s *GetPApiSwapOpenOrdersService) Symbol(symbol string) SwapOpenOrdersGetter
```



<a name="GetPremiumIndexService"></a>
## type GetPremiumIndexService

GetPremiumIndexService 获取资金费率和指数价格

```go
type GetPremiumIndexService struct {
    // contains filtered or unexported fields
}
```

<a name="GetPremiumIndexService.Do"></a>
### func \(\*GetPremiumIndexService\) Do

```go
func (s *GetPremiumIndexService) Do(ctx context.Context, opts ...RequestOption) ([]*PremiumIndex, error)
```



<a name="GetPremiumIndexService.Symbol"></a>
### func \(\*GetPremiumIndexService\) Symbol

```go
func (s *GetPremiumIndexService) Symbol(symbol string) *GetPremiumIndexService
```



<a name="GetService"></a>
## type GetService



```go
type GetService struct {
    // contains filtered or unexported fields
}
```

<a name="GetService.Do"></a>
### func \(\*GetService\) Do

```go
func (s *GetService) Do(ctx context.Context, dest any, opts ...RequestOption) error
```



<a name="GetSpotOpenOrdersService"></a>
## type GetSpotOpenOrdersService



```go
type GetSpotOpenOrdersService struct {
    // contains filtered or unexported fields
}
```

<a name="GetSpotOpenOrdersService.Do"></a>
### func \(\*GetSpotOpenOrdersService\) Do

```go
func (s *GetSpotOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]*SpotOpenOrder, error)
```



<a name="GetSpotOpenOrdersService.Symbol"></a>
### func \(\*GetSpotOpenOrdersService\) Symbol

```go
func (s *GetSpotOpenOrdersService) Symbol(symbol string) *GetSpotOpenOrdersService
```



<a name="GetSpotOrdersService"></a>
## type GetSpotOrdersService



```go
type GetSpotOrdersService struct {
    // contains filtered or unexported fields
}
```

<a name="GetSpotOrdersService.Do"></a>
### func \(\*GetSpotOrdersService\) Do

```go
func (s *GetSpotOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]*SpotOrder, error)
```



<a name="GetSpotOrdersService.EndTime"></a>
### func \(\*GetSpotOrdersService\) EndTime

```go
func (s *GetSpotOrdersService) EndTime(endTime int64) *GetSpotOrdersService
```



<a name="GetSpotOrdersService.Limit"></a>
### func \(\*GetSpotOrdersService\) Limit

```go
func (s *GetSpotOrdersService) Limit(limit int) *GetSpotOrdersService
```



<a name="GetSpotOrdersService.OrderId"></a>
### func \(\*GetSpotOrdersService\) OrderId

```go
func (s *GetSpotOrdersService) OrderId(orderId string) *GetSpotOrdersService
```



<a name="GetSpotOrdersService.StartTime"></a>
### func \(\*GetSpotOrdersService\) StartTime

```go
func (s *GetSpotOrdersService) StartTime(startTime int64) *GetSpotOrdersService
```



<a name="GetSpotOrdersService.Symbol"></a>
### func \(\*GetSpotOrdersService\) Symbol

```go
func (s *GetSpotOrdersService) Symbol(symbol string) *GetSpotOrdersService
```



<a name="GetSpotSymbolInfosService"></a>
## type GetSpotSymbolInfosService

GetSpotSymbolInfosService 获取现货交易对信息

```go
type GetSpotSymbolInfosService struct {
    // contains filtered or unexported fields
}
```

<a name="GetSpotSymbolInfosService.Do"></a>
### func \(\*GetSpotSymbolInfosService\) Do

```go
func (s *GetSpotSymbolInfosService) Do(ctx context.Context, opts ...RequestOption) ([]*SymbolInfo, error)
```



<a name="GetSpotSymbolInfosService.Status"></a>
### func \(\*GetSpotSymbolInfosService\) Status

```go
func (s *GetSpotSymbolInfosService) Status(status string) *GetSpotSymbolInfosService
```



<a name="GetSpotTicker24HService"></a>
## type GetSpotTicker24HService



```go
type GetSpotTicker24HService struct {
    // contains filtered or unexported fields
}
```

<a name="GetSpotTicker24HService.Do"></a>
### func \(\*GetSpotTicker24HService\) Do

```go
func (s *GetSpotTicker24HService) Do(ctx context.Context, opts ...RequestOption) ([]*SpotTicker24H, error)
```



<a name="GetSpotTicker24HService.Symbol"></a>
### func \(\*GetSpotTicker24HService\) Symbol

```go
func (s *GetSpotTicker24HService) Symbol(symbol string) *GetSpotTicker24HService
```



<a name="GetSpotTicker24HService.Symbols"></a>
### func \(\*GetSpotTicker24HService\) Symbols

```go
func (s *GetSpotTicker24HService) Symbols(symbols []string) *GetSpotTicker24HService
```



<a name="GetSpotTickerPriceService"></a>
## type GetSpotTickerPriceService

GetSpotTickerPriceService 获取现货交易对价格

```go
type GetSpotTickerPriceService struct {
    // contains filtered or unexported fields
}
```

<a name="GetSpotTickerPriceService.Do"></a>
### func \(\*GetSpotTickerPriceService\) Do

```go
func (s *GetSpotTickerPriceService) Do(ctx context.Context, opts ...RequestOption) ([]*TickerPrice, error)
```



<a name="GetSpotTickerPriceService.Symbol"></a>
### func \(\*GetSpotTickerPriceService\) Symbol

```go
func (s *GetSpotTickerPriceService) Symbol(symbol string) *GetSpotTickerPriceService
```



<a name="GetSpotTickerPriceService.Symbols"></a>
### func \(\*GetSpotTickerPriceService\) Symbols

```go
func (s *GetSpotTickerPriceService) Symbols(symbols []string) *GetSpotTickerPriceService
```



<a name="GetUMAccountDetailService"></a>
## type GetUMAccountDetailService

获取U本位账户详情 https://developers.binance.com/docs/zh-CN/derivatives/portfolio-margin/account/Get-UM-Account-Detail

```go
type GetUMAccountDetailService struct {
    // contains filtered or unexported fields
}
```

<a name="GetUMAccountDetailService.Do"></a>
### func \(\*GetUMAccountDetailService\) Do

```go
func (s *GetUMAccountDetailService) Do(ctx context.Context, opts ...RequestOption) (*UMAccountDetail, error)
```



<a name="GetUMAccountService"></a>
## type GetUMAccountService

获取U本位账户信息 https://developers.binance.com/docs/zh-CN/derivatives/portfolio-margin/account/Account-Information

```go
type GetUMAccountService struct {
    // contains filtered or unexported fields
}
```

<a name="GetUMAccountService.Do"></a>
### func \(\*GetUMAccountService\) Do

```go
func (s *GetUMAccountService) Do(ctx context.Context, opts ...RequestOption) (*UMAccount, error)
```



<a name="GetUserAssetService"></a>
## type GetUserAssetService



```go
type GetUserAssetService struct {
    // contains filtered or unexported fields
}
```

<a name="GetUserAssetService.Do"></a>
### func \(\*GetUserAssetService\) Do

```go
func (s *GetUserAssetService) Do(ctx context.Context, opts ...RequestOption) ([]UserAsset, error)
```



<a name="GetWalletBalanceService"></a>
## type GetWalletBalanceService

GetWalletBalanceService 获取钱包余额

```go
type GetWalletBalanceService struct {
    // contains filtered or unexported fields
}
```

<a name="GetWalletBalanceService.Do"></a>
### func \(\*GetWalletBalanceService\) Do

```go
func (s *GetWalletBalanceService) Do(ctx context.Context, opts ...RequestOption) ([]WalletBalance, error)
```



<a name="GetWalletBalanceService.QuoteAsset"></a>
### func \(\*GetWalletBalanceService\) QuoteAsset

```go
func (s *GetWalletBalanceService) QuoteAsset(quoteAsset string) *GetWalletBalanceService
```



<a name="Income"></a>
## type Income



```go
type Income struct {
    Symbol     string  `json:"symbol"`     // 交易对，仅针对涉及交易对的资金流
    IncomeType string  `json:"incomeType"` // 资金流类型
    Income     Float64 `json:"income"`     // 资金流数量，正数代表流入，负数代表流出
    Asset      string  `json:"asset"`      // 资产内容
    Info       string  `json:"info"`       // 备注信息，取决于流水类型
    Time       Int64   `json:"time"`       // 时间
    TranId     Int64   `json:"tranId"`     // 划转ID
    TradeId    Int64   `json:"tradeId"`    // 引起流水产生的原始交易ID
}
```

<a name="LimitFilter"></a>
## type LimitFilter



```go
type LimitFilter struct {
    FilterType string  `json:"filterType"`
    MinQty     Float64 `json:"minQty"`
    MaxQty     Float64 `json:"maxQty"`
    StepSize   Float64 `json:"stepSize"`
    Notional   Float64 `json:"notional"`
}
```

<a name="LotLimit"></a>
## type LotLimit



```go
type LotLimit struct {
    MinQty   Float64 `json:"minQty"`
    MaxQty   Float64 `json:"maxQty"`
    StepSize Float64 `json:"stepSize"`
}
```

<a name="MarkPriceEvent"></a>
## type MarkPriceEvent



```go
type MarkPriceEvent struct {
    Event           string  `json:"e"` // 事件类型
    Time            int64   `json:"E"` // 事件时间
    Symbol          string  `json:"s"` // 交易对
    MarkPrice       Float64 `json:"p"` // 标记价格
    IndexPrice      Float64 `json:"i"` // 现货指数价格
    EstimatedSettle Float64 `json:"P"` // 预估结算价格
    FundingRate     Float64 `json:"r"` // 资金费率
    NextFundingTime int64   `json:"T"` // 下一个资金费率时间
}
```

<a name="PApiAccountBalance"></a>
## type PApiAccountBalance



```go
type PApiAccountBalance struct {
    Asset               string  `json:"asset"`               // 资产
    TotalWalletBalance  string  `json:"totalWalletBalance"`  // 钱包余额 = 全仓杠杆未锁定 + 全仓杠杆锁定 + u本位合约钱包余额 + 币本位合约钱包余额
    CrossMarginAsset    Float64 `json:"crossMarginAsset"`    // 全仓资产 = 全仓杠杆未锁定 + 全仓杠杆锁定
    CrossMarginBorrowed string  `json:"crossMarginBorrowed"` // 全仓杠杆借贷
    CrossMarginFree     Float64 `json:"crossMarginFree"`     // 全仓杠杆未锁定
    CrossMarginInterest string  `json:"crossMarginInterest"` // 全仓杠杆利息
    CrossMarginLocked   Float64 `json:"crossMarginLocked"`   // 全仓杠杆锁定
    UmWalletBalance     Float64 `json:"umWalletBalance"`     // u本位合约钱包余额
    UmUnrealizedPNL     string  `json:"umUnrealizedPNL"`     // u本位未实现盈亏
    CmWalletBalance     Float64 `json:"cmWalletBalance"`     // 币本位合约钱包余额
    CmUnrealizedPNL     string  `json:"cmUnrealizedPNL"`     // 币本位未实现盈亏
    UpdateTime          int64   `json:"updateTime"`          // 更新时间
    NegativeBalance     string  `json:"negativeBalance"`     // 负余额
}
```

<a name="PApiGetAllOrdersService"></a>
## type PApiGetAllOrdersService



```go
type PApiGetAllOrdersService struct {
    // contains filtered or unexported fields
}
```

<a name="PApiGetAllOrdersService.Do"></a>
### func \(\*PApiGetAllOrdersService\) Do

```go
func (s *PApiGetAllOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]*SwapOrder, error)
```



<a name="PApiGetAllOrdersService.EndTime"></a>
### func \(\*PApiGetAllOrdersService\) EndTime

```go
func (s *PApiGetAllOrdersService) EndTime(endTime int64) *PApiGetAllOrdersService
```



<a name="PApiGetAllOrdersService.Limit"></a>
### func \(\*PApiGetAllOrdersService\) Limit

```go
func (s *PApiGetAllOrdersService) Limit(limit int) *PApiGetAllOrdersService
```



<a name="PApiGetAllOrdersService.OrderId"></a>
### func \(\*PApiGetAllOrdersService\) OrderId

```go
func (s *PApiGetAllOrdersService) OrderId(orderId string) *PApiGetAllOrdersService
```



<a name="PApiGetAllOrdersService.StartTime"></a>
### func \(\*PApiGetAllOrdersService\) StartTime

```go
func (s *PApiGetAllOrdersService) StartTime(startTime int64) *PApiGetAllOrdersService
```



<a name="PApiGetOpenOrdersService"></a>
## type PApiGetOpenOrdersService



```go
type PApiGetOpenOrdersService struct {
    // contains filtered or unexported fields
}
```

<a name="PApiGetOpenOrdersService.Do"></a>
### func \(\*PApiGetOpenOrdersService\) Do

```go
func (s *PApiGetOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]*SwapOpenOrder, error)
```



<a name="PApiGetOpenOrdersService.Symbol"></a>
### func \(\*PApiGetOpenOrdersService\) Symbol

```go
func (s *PApiGetOpenOrdersService) Symbol(symbol string) *PApiGetOpenOrdersService
```



<a name="PApiGetUMPositionsService"></a>
## type PApiGetUMPositionsService



```go
type PApiGetUMPositionsService struct {
    // contains filtered or unexported fields
}
```

<a name="PApiGetUMPositionsService.Do"></a>
### func \(\*PApiGetUMPositionsService\) Do

```go
func (s *PApiGetUMPositionsService) Do(ctx context.Context, opts ...RequestOption) ([]*Position, error)
```



<a name="PApiUmCancelAllOpenOrdersService"></a>
## type PApiUmCancelAllOpenOrdersService



```go
type PApiUmCancelAllOpenOrdersService struct {
    // contains filtered or unexported fields
}
```

<a name="PApiUmCancelAllOpenOrdersService.Do"></a>
### func \(\*PApiUmCancelAllOpenOrdersService\) Do

```go
func (s *PApiUmCancelAllOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) error
```



<a name="PApiUmCancelOrderService"></a>
## type PApiUmCancelOrderService



```go
type PApiUmCancelOrderService struct {
    // contains filtered or unexported fields
}
```

<a name="PApiUmCancelOrderService.Do"></a>
### func \(\*PApiUmCancelOrderService\) Do

```go
func (s *PApiUmCancelOrderService) Do(ctx context.Context, opts ...RequestOption) (*SwapOrder, error)
```



<a name="PApiUmCancelOrderService.OrderId"></a>
### func \(\*PApiUmCancelOrderService\) OrderId

```go
func (s *PApiUmCancelOrderService) OrderId(orderId string) *PApiUmCancelOrderService
```



<a name="PApiUmCancelOrderService.OrigClientOrderId"></a>
### func \(\*PApiUmCancelOrderService\) OrigClientOrderId

```go
func (s *PApiUmCancelOrderService) OrigClientOrderId(origClientOrderId string) *PApiUmCancelOrderService
```



<a name="PApiUmCancelOrderService.Symbol"></a>
### func \(\*PApiUmCancelOrderService\) Symbol

```go
func (s *PApiUmCancelOrderService) Symbol(symbol string) *PApiUmCancelOrderService
```



<a name="PApiUmChangeLeverageService"></a>
## type PApiUmChangeLeverageService



```go
type PApiUmChangeLeverageService struct {
    // contains filtered or unexported fields
}
```

<a name="PApiUmChangeLeverageService.Do"></a>
### func \(\*PApiUmChangeLeverageService\) Do

```go
func (s *PApiUmChangeLeverageService) Do(ctx context.Context, opts ...RequestOption) error
```



<a name="PApiUmGetIncomeService"></a>
## type PApiUmGetIncomeService



```go
type PApiUmGetIncomeService struct {
    // contains filtered or unexported fields
}
```

<a name="PApiUmGetIncomeService.Do"></a>
### func \(\*PApiUmGetIncomeService\) Do

```go
func (s *PApiUmGetIncomeService) Do(ctx context.Context, opts ...RequestOption) ([]Income, error)
```



<a name="PApiUmGetIncomeService.EndTime"></a>
### func \(\*PApiUmGetIncomeService\) EndTime

```go
func (s *PApiUmGetIncomeService) EndTime(endTime int64) *PApiUmGetIncomeService
```



<a name="PApiUmGetIncomeService.IncomeType"></a>
### func \(\*PApiUmGetIncomeService\) IncomeType

```go
func (s *PApiUmGetIncomeService) IncomeType(incomeType string) *PApiUmGetIncomeService
```



<a name="PApiUmGetIncomeService.Limit"></a>
### func \(\*PApiUmGetIncomeService\) Limit

```go
func (s *PApiUmGetIncomeService) Limit(limit int) *PApiUmGetIncomeService
```



<a name="PApiUmGetIncomeService.StartTime"></a>
### func \(\*PApiUmGetIncomeService\) StartTime

```go
func (s *PApiUmGetIncomeService) StartTime(startTime int64) *PApiUmGetIncomeService
```



<a name="PApiUmGetIncomeService.Symbol"></a>
### func \(\*PApiUmGetIncomeService\) Symbol

```go
func (s *PApiUmGetIncomeService) Symbol(symbol string) *PApiUmGetIncomeService
```



<a name="PApiUmGetOrderService"></a>
## type PApiUmGetOrderService



```go
type PApiUmGetOrderService struct {
    // contains filtered or unexported fields
}
```

<a name="PApiUmGetOrderService.Do"></a>
### func \(\*PApiUmGetOrderService\) Do

```go
func (s *PApiUmGetOrderService) Do(ctx context.Context, opts ...RequestOption) (*SwapOrder, error)
```



<a name="PApiUmGetOrderService.OrderId"></a>
### func \(\*PApiUmGetOrderService\) OrderId

```go
func (s *PApiUmGetOrderService) OrderId(orderId string) *PApiUmGetOrderService
```



<a name="PApiUmGetOrderService.OrigClientOrderId"></a>
### func \(\*PApiUmGetOrderService\) OrigClientOrderId

```go
func (s *PApiUmGetOrderService) OrigClientOrderId(origClientOrderId string) *PApiUmGetOrderService
```



<a name="PApiUmGetUserTradesService"></a>
## type PApiUmGetUserTradesService



```go
type PApiUmGetUserTradesService struct {
    // contains filtered or unexported fields
}
```

<a name="PApiUmGetUserTradesService.Do"></a>
### func \(\*PApiUmGetUserTradesService\) Do

```go
func (s *PApiUmGetUserTradesService) Do(ctx context.Context, opts ...RequestOption) ([]*TradeRecord, error)
```



<a name="PApiUmGetUserTradesService.EndTime"></a>
### func \(\*PApiUmGetUserTradesService\) EndTime

```go
func (s *PApiUmGetUserTradesService) EndTime(endTime int64) *PApiUmGetUserTradesService
```



<a name="PApiUmGetUserTradesService.FromId"></a>
### func \(\*PApiUmGetUserTradesService\) FromId

```go
func (s *PApiUmGetUserTradesService) FromId(fromId int64) *PApiUmGetUserTradesService
```



<a name="PApiUmGetUserTradesService.Limit"></a>
### func \(\*PApiUmGetUserTradesService\) Limit

```go
func (s *PApiUmGetUserTradesService) Limit(limit int) *PApiUmGetUserTradesService
```



<a name="PApiUmGetUserTradesService.StartTime"></a>
### func \(\*PApiUmGetUserTradesService\) StartTime

```go
func (s *PApiUmGetUserTradesService) StartTime(startTime int64) *PApiUmGetUserTradesService
```



<a name="PApiUmOrderService"></a>
## type PApiUmOrderService



```go
type PApiUmOrderService struct {
    // contains filtered or unexported fields
}
```

<a name="PApiUmOrderService.Do"></a>
### func \(\*PApiUmOrderService\) Do

```go
func (s *PApiUmOrderService) Do(ctx context.Context, opts ...RequestOption) (*SwapOrder, error)
```



<a name="PApiUmOrderService.GoodTillDate"></a>
### func \(\*PApiUmOrderService\) GoodTillDate

```go
func (s *PApiUmOrderService) GoodTillDate(goodTillDate int64) *PApiUmOrderService
```



<a name="PApiUmOrderService.NewClientOrderId"></a>
### func \(\*PApiUmOrderService\) NewClientOrderId

```go
func (s *PApiUmOrderService) NewClientOrderId(newClientOrderId string) *PApiUmOrderService
```



<a name="PApiUmOrderService.NewOrderRespType"></a>
### func \(\*PApiUmOrderService\) NewOrderRespType

```go
func (s *PApiUmOrderService) NewOrderRespType(newOrderRespType string) *PApiUmOrderService
```



<a name="PApiUmOrderService.OrderType"></a>
### func \(\*PApiUmOrderService\) OrderType

```go
func (s *PApiUmOrderService) OrderType(orderType string) *PApiUmOrderService
```



<a name="PApiUmOrderService.PositionSide"></a>
### func \(\*PApiUmOrderService\) PositionSide

```go
func (s *PApiUmOrderService) PositionSide(positionSide string) *PApiUmOrderService
```



<a name="PApiUmOrderService.Price"></a>
### func \(\*PApiUmOrderService\) Price

```go
func (s *PApiUmOrderService) Price(price string) *PApiUmOrderService
```



<a name="PApiUmOrderService.PriceMatch"></a>
### func \(\*PApiUmOrderService\) PriceMatch

```go
func (s *PApiUmOrderService) PriceMatch(priceMatch string) *PApiUmOrderService
```



<a name="PApiUmOrderService.Quantity"></a>
### func \(\*PApiUmOrderService\) Quantity

```go
func (s *PApiUmOrderService) Quantity(quantity string) *PApiUmOrderService
```



<a name="PApiUmOrderService.ReduceOnly"></a>
### func \(\*PApiUmOrderService\) ReduceOnly

```go
func (s *PApiUmOrderService) ReduceOnly(reduceOnly bool) *PApiUmOrderService
```



<a name="PApiUmOrderService.SelfTradePreventionMode"></a>
### func \(\*PApiUmOrderService\) SelfTradePreventionMode

```go
func (s *PApiUmOrderService) SelfTradePreventionMode(selfTradePreventionMode string) *PApiUmOrderService
```



<a name="PApiUmOrderService.TimeInForce"></a>
### func \(\*PApiUmOrderService\) TimeInForce

```go
func (s *PApiUmOrderService) TimeInForce(timeInForce string) *PApiUmOrderService
```



<a name="PapiUMGetPositionRiskService"></a>
## type PapiUMGetPositionRiskService



```go
type PapiUMGetPositionRiskService struct {
    // contains filtered or unexported fields
}
```

<a name="PapiUMGetPositionRiskService.Do"></a>
### func \(\*PapiUMGetPositionRiskService\) Do

```go
func (s *PapiUMGetPositionRiskService) Do(ctx context.Context, opts ...RequestOption) ([]UMPosition, error)
```



<a name="PapiUMGetPositionRiskService.Symbol"></a>
### func \(\*PapiUMGetPositionRiskService\) Symbol

```go
func (s *PapiUMGetPositionRiskService) Symbol(symbol string) *PapiUMGetPositionRiskService
```



<a name="PingService"></a>
## type PingService



```go
type PingService struct {
    // contains filtered or unexported fields
}
```

<a name="PingService.Do"></a>
### func \(\*PingService\) Do

```go
func (s *PingService) Do(ctx context.Context, opts ...RequestOption) error
```



<a name="Position"></a>
## type Position



```go
type Position struct {
    Symbol           string  `json:"symbol"`           // 交易对
    PositionAmt      Float64 `json:"positionAmt"`      // 头寸数量，符号代表多空方向, 正数为多，负数为空
    PositionSide     string  `json:"positionSide"`     // 持仓方向
    EntryPrice       Float64 `json:"entryPrice"`       // 开仓均价
    Leverage         Int64   `json:"leverage"`         // 杠杆倍数
    UnRealizedProfit Float64 `json:"unRealizedProfit"` // 持仓未实现盈亏
    InitialMargin    Float64 `json:"initialMargin"`    // 初始保证金
    MaintMargin      Float64 `json:"maintMargin"`      // 维持保证金
    MarkPrice        Float64 `json:"markPrice"`        // 当前标记价格
    Notional         Float64 `json:"notional"`         // 名义价值
    LiquidationPrice Float64 `json:"liquidationPrice"` // 预估强平价格
    UpdateTime       Int64   `json:"updateTime"`       // 更新时间
}
```

<a name="PositionsGetter"></a>
## type PositionsGetter



```go
type PositionsGetter interface {
    Do(ctx context.Context, opts ...RequestOption) ([]*Position, error)
}
```

<a name="PostService"></a>
## type PostService



```go
type PostService struct {
    // contains filtered or unexported fields
}
```

<a name="PostService.Do"></a>
### func \(\*PostService\) Do

```go
func (s *PostService) Do(ctx context.Context, dest any, opts ...RequestOption) error
```



<a name="PremiumIndex"></a>
## type PremiumIndex



```go
type PremiumIndex struct {
    Symbol          string  `json:"symbol"`               // 交易对
    MarkPrice       Float64 `json:"markPrice"`            // 标记价格
    IndexPrice      Float64 `json:"indexPrice"`           // 指数价格
    EstSettlePrice  string  `json:"estimatedSettlePrice"` // 预估结算价,仅在交割开始前最后一小时有意义
    LastFundingRate Float64 `json:"lastFundingRate"`      // 最近更新的资金费率
    InterestRate    string  `json:"interestRate"`         // 标的资产基础利率
    NextFundingTime Int64   `json:"nextFundingTime"`      // 下次资金费时间
    Time            Int64   `json:"time"`                 // 更新时间
}
```

<a name="RequestOption"></a>
## type RequestOption

RequestOption define option type for request

```go
type RequestOption func(*request)
```

<a name="WithHeader"></a>
### func WithHeader

```go
func WithHeader(key, value string, replace bool) RequestOption
```

WithHeader set or add a header value to the request

<a name="WithHeaders"></a>
### func WithHeaders

```go
func WithHeaders(header http.Header) RequestOption
```

WithHeaders set or replace the headers of the request

<a name="WithRecvWindow"></a>
### func WithRecvWindow

```go
func WithRecvWindow(recvWindow int64) RequestOption
```

WithRecvWindow set recvWindow param for the request

<a name="SpotOpenOrder"></a>
## type SpotOpenOrder



```go
type SpotOpenOrder struct {
    Symbol                  string `json:"symbol"`                  // 交易对
    OrderId                 int64  `json:"orderId"`                 // 订单ID
    OrderListId             int64  `json:"orderListId"`             // 除非此单是订单列表的一部分, 否则此值为 -1
    ClientOrderId           string `json:"clientOrderId"`           // 客户端订单ID
    Price                   string `json:"price"`                   // 订单价格
    OrigQty                 string `json:"origQty"`                 // 原始数量
    ExecutedQty             string `json:"executedQty"`             // 已执行数量
    OrigQuoteOrderQty       string `json:"origQuoteOrderQty"`       // 原始报价订单数量
    CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`     // 累计报价数量
    Status                  string `json:"status"`                  // 订单状态
    TimeInForce             string `json:"timeInForce"`             // 有效方式(GTC/IOC/FOK)
    Type                    string `json:"type"`                    // 订单类型
    Side                    string `json:"side"`                    // 订单方向(BUY/SELL)
    StopPrice               string `json:"stopPrice"`               // 止损价
    IcebergQty              string `json:"icebergQty"`              // 冰山订单数量
    Time                    int64  `json:"time"`                    // 订单时间
    UpdateTime              int64  `json:"updateTime"`              // 最后更新时间
    IsWorking               bool   `json:"isWorking"`               // 订单是否生效
    WorkingTime             int64  `json:"workingTime"`             // 工作时间
    SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 自成交预防模式
}
```

<a name="SpotOrder"></a>
## type SpotOrder



```go
type SpotOrder struct {
    Symbol                  string `json:"symbol"`                  // 交易对
    OrderId                 int64  `json:"orderId"`                 // 订单ID
    OrderListId             int64  `json:"orderListId"`             // 订单列表ID
    ClientOrderId           string `json:"clientOrderId"`           // 客户订单ID
    Price                   string `json:"price"`                   // 价格
    OrigQty                 string `json:"origQty"`                 // 原始挂单数量
    ExecutedQty             string `json:"executedQty"`             // 成交数量
    CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`     // 累积成交额
    Status                  string `json:"status"`                  // 订单状态
    TimeInForce             string `json:"timeInForce"`             // 时间类型
    Type                    string `json:"type"`                    // 订单类型
    Side                    string `json:"side"`                    // BUY or SELL
    StopPrice               string `json:"stopPrice"`               // 止损价格
    IcebergQty              string `json:"icebergQty"`              // 冰山订单数量
    Time                    int64  `json:"time"`                    // 订单创建时间
    UpdateTime              int64  `json:"updateTime"`              // 订单更新时间
    IsWorking               bool   `json:"isWorking"`               // 是否正在处理
    OrigQuoteOrderQty       string `json:"origQuoteOrderQty"`       // 原始报价订单数量
    WorkingTime             int64  `json:"workingTime"`             // 处理时间
    SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 自成交预防模式
}
```

<a name="SpotSymbolInfos"></a>
## type SpotSymbolInfos



```go
type SpotSymbolInfos struct {
    Symbols []struct {
        Symbol     string            `json:"symbol"`     // 交易对
        Status     string            `json:"status"`     // 状态
        QuoteAsset string            `json:"quoteAsset"` // 报价币种
        Filters    []json.RawMessage `json:"filters"`    // 过滤器
    } `json:"symbols"`
}
```

<a name="SpotTicker24H"></a>
## type SpotTicker24H



```go
type SpotTicker24H struct {
    Symbol      string  `json:"symbol"`      // 交易对
    OpenPrice   Float64 `json:"openPrice"`   // 间隔开盘价
    HighPrice   Float64 `json:"highPrice"`   // 间隔最高价
    LowPrice    Float64 `json:"lowPrice"`    // 间隔最低价
    LastPrice   Float64 `json:"lastPrice"`   // 间隔收盘价
    Volume      Float64 `json:"volume"`      // 总交易量
    QuoteVolume Float64 `json:"quoteVolume"` // 总交易额
    OpenTime    Int64   `json:"openTime"`    // ticker间隔的开始时间
    CloseTime   Int64   `json:"closeTime"`   // ticker间隔的结束时间
    FirstId     Int64   `json:"firstId"`     // 统计时间内的第一笔trade id
    LastId      Int64   `json:"lastId"`      // 统计时间内的最后一笔trade id
    Count       Int64   `json:"count"`       // 统计时间内交易笔数
}
```

<a name="SwapOpenOrder"></a>
## type SwapOpenOrder



```go
type SwapOpenOrder struct {
    AvgPrice             string  `json:"avgPrice"`                // 平均成交价
    ClientOrderId        string  `json:"clientOrderId"`           // 用户自定义订单ID
    CumQuote             Float64 `json:"cumQuote"`                // 成交金额
    ExecutedQty          Float64 `json:"executedQty"`             // 已成交数量
    OrderId              Int64   `json:"orderId"`                 // 订单ID
    OrigQty              Float64 `json:"origQty"`                 // 原始委托数量
    OrigType             string  `json:"origType"`                // 原始订单类型
    Price                Float64 `json:"price"`                   // 委托价格
    ReduceOnly           bool    `json:"reduceOnly"`              // 是否仅减仓
    Side                 string  `json:"side"`                    // 买卖方向
    PositionSide         string  `json:"positionSide"`            // 持仓方向
    Status               string  `json:"status"`                  // 订单状态
    Symbol               string  `json:"symbol"`                  // 交易对
    Time                 Int64   `json:"time"`                    // 订单时间
    TimeInForce          string  `json:"timeInForce"`             // 有效方式
    Type                 string  `json:"type"`                    // 订单类型
    UpdateTime           Int64   `json:"updateTime"`              // 更新时间
    SelfTradePreventMode string  `json:"selfTradePreventionMode"` // 自成交预防模式
    GoodTillDate         Int64   `json:"goodTillDate"`            // 订单到期时间
    PriceMatch           string  `json:"priceMatch"`              // 价格匹配模式

}
```

<a name="SwapOpenOrdersGetter"></a>
## type SwapOpenOrdersGetter



```go
type SwapOpenOrdersGetter interface {
    Do(ctx context.Context, opts ...RequestOption) ([]*SwapOpenOrder, error)
}
```

<a name="SwapOrder"></a>
## type SwapOrder



```go
type SwapOrder struct {
    AvgPrice             Float64 `json:"avgPrice"`                // 平均价格
    ClientOrderId        string  `json:"clientOrderId"`           // 客户订单ID
    CumQuote             Float64 `json:"cumQuote"`                // 累积成交额
    ExecutedQty          Float64 `json:"executedQty"`             // 成交数量
    OrderId              Int64   `json:"orderId"`                 // 订单ID
    OrigQty              Float64 `json:"origQty"`                 // 原始挂单数量
    OrigType             string  `json:"origType"`                // 原始订单类型
    Price                Float64 `json:"price"`                   // 价格
    ReduceOnly           bool    `json:"reduceOnly"`              // 是否只减仓
    Side                 string  `json:"side"`                    // 买卖方向 BUY/SELL
    PositionSide         string  `json:"positionSide"`            // 持仓方向
    Status               string  `json:"status"`                  // 订单状态
    ClosePosition        bool    `json:"closePosition"`           // 是否全平仓
    Symbol               string  `json:"symbol"`                  // 交易对
    Time                 Int64   `json:"time"`                    // 订单创建时间
    TimeInForce          string  `json:"timeInForce"`             // 时间类型
    Type                 string  `json:"type"`                    // 订单类型
    UpdateTime           Int64   `json:"updateTime"`              // 订单更新时间
    SelfTradePreventMode string  `json:"selfTradePreventionMode"` // 自成交预防模式
    GoodTillDate         Int64   `json:"goodTillDate"`            // 有效期
    PriceMatch           string  `json:"priceMatch"`              // 价格匹配规则
}
```

<a name="SwapOrdersGetter"></a>
## type SwapOrdersGetter



```go
type SwapOrdersGetter interface {
    Symbol(symbol string) SwapOrdersGetter
    OrderId(orderId string) SwapOrdersGetter
    StartTime(startTime int64) SwapOrdersGetter
    EndTime(endTime int64) SwapOrdersGetter
    Limit(limit int) SwapOrdersGetter
    Do(ctx context.Context, opts ...RequestOption) ([]*SwapOrder, error)
}
```

<a name="SymbolInfo"></a>
## type SymbolInfo



```go
type SymbolInfo struct {
    Symbol              string   // 交易对
    QuoteAsset          string   // 报价币种
    PricePrecision      int      // 价格精度
    QuantityPrecision   int      // 数量精度
    BaseAssetPrecision  int      // 基础货币精度
    QuoteAssetPrecision int      // 报价货币精度
    MinValue            Float64  // 最小下单价值
    LotSize             LotLimit // 限价单下单量
    MarketLotSize       LotLimit // 市价单下单量
}
```

<a name="TickerPrice"></a>
## type TickerPrice



```go
type TickerPrice struct {
    Symbol string  `json:"symbol"` // 交易对
    Price  Float64 `json:"price"`  // 价格
    Time   Int64   `json:"time"`   // 更新时间
}
```

<a name="TimeInForceType"></a>
## type TimeInForceType



```go
type TimeInForceType string
```

<a name="TradeRecord"></a>
## type TradeRecord



```go
type TradeRecord struct {
    Symbol          string  `json:"symbol"`
    ID              Int64   `json:"id"`
    OrderID         Int64   `json:"orderId"`
    Side            string  `json:"side"`
    Price           Float64 `json:"price"`
    Qty             Float64 `json:"qty"`
    RealizedPnl     Float64 `json:"realizedPnl"`
    QuoteQty        Float64 `json:"quoteQty"`
    Commission      Float64 `json:"commission"`
    CommissionAsset string  `json:"commissionAsset"`
    Time            Int64   `json:"time"`
    Buyer           bool    `json:"buyer"`
    Maker           bool    `json:"maker"`
    PositionSide    string  `json:"positionSide"`
}
```

<a name="UMAccount"></a>
## type UMAccount



```go
type UMAccount struct {
    UniMMR                   Float64 `json:"uniMMR"`                   // 统一账户维持保证金率
    AccountEquity            Float64 `json:"accountEquity"`            // 账户权益（美元）
    ActualEquity             Float64 `json:"actualEquity"`             // 实际权益（不包含抵押）
    AccountInitialMargin     Float64 `json:"accountInitialMargin"`     // 账户初始保证金
    AccountMaintMargin       Float64 `json:"accountMaintMargin"`       // 账户维持保证金（美元）
    AccountStatus            string  `json:"accountStatus"`            // 账户状态："NORMAL", "MARGIN_CALL", "SUPPLY_MARGIN", "REDUCE_ONLY", "ACTIVE_LIQUIDATION", "FORCE_LIQUIDATION", "BANKRUPTED"
    VirtualMaxWithdrawAmount Float64 `json:"virtualMaxWithdrawAmount"` // 虚拟最大提现金额
    TotalAvailableBalance    Float64 `json:"totalAvailableBalance"`    // 最高可转出金额（美元）
    TotalMarginOpenLoss      Float64 `json:"totalMarginOpenLoss"`      // 美元保证金未结订单
    UpdateTime               int64   `json:"updateTime"`               // 更新时间
}
```

<a name="UMAccountDetail"></a>
## type UMAccountDetail



```go
type UMAccountDetail struct {
    Assets    []UMAsset    `json:"assets"`
    Positions []UMPosition `json:"positions"`
}
```

<a name="UMAsset"></a>
## type UMAsset



```go
type UMAsset struct {
    Asset                  string  `json:"asset"`                  // 资产
    CrossWalletBalance     Float64 `json:"crossWalletBalance"`     // 全仓余额
    CrossUnPnl             Float64 `json:"crossUnPnl"`             // 全仓未实现盈亏
    MaintMargin            Float64 `json:"maintMargin"`            // 维持保证金
    InitialMargin          Float64 `json:"initialMargin"`          // 初始保证金
    PositionInitialMargin  Float64 `json:"positionInitialMargin"`  // 持仓初始保证金
    OpenOrderInitialMargin Float64 `json:"openOrderInitialMargin"` // 挂单初始保证金
    UpdateTime             Int64   `json:"updateTime"`             // 更新时间
}
```

<a name="UMPosition"></a>
## type UMPosition



```go
type UMPosition struct {
    Symbol                 string  `json:"symbol"`                 // 交易对
    InitialMargin          Float64 `json:"initialMargin"`          // 当前标记价格下的初始保证金
    MaintMargin            Float64 `json:"maintMargin"`            // 维持保证金
    UnrealizedProfit       Float64 `json:"unrealizedProfit"`       // 未实现盈亏
    UnRealizedProfit       Float64 `json:"unRealizedProfit"`       // 未实现盈亏
    PositionInitialMargin  string  `json:"positionInitialMargin"`  // 当前标记价格下持仓所需初始保证金
    OpenOrderInitialMargin string  `json:"openOrderInitialMargin"` // 当前标记价格下挂单所需初始保证金
    Leverage               Int64   `json:"leverage"`               // 当前杠杆倍数
    EntryPrice             Float64 `json:"entryPrice"`             // 平均入场价格
    MarkPrice              Float64 `json:"markPrice"`              // 当前标记价格
    LiquidationPrice       Float64 `json:"liquidationPrice"`       // 预估强平价格
    MaxNotional            string  `json:"maxNotional"`            // 当前杠杆下最大可用名义价值
    MaxNotionalValue       Float64 `json:"maxNotionalValue"`       // 当前杠杆下最大可用名义价值
    Notional               Float64 `json:"notional"`               // 名义价值
    BidNotional            string  `json:"bidNotional"`            // 买单名义价值(忽略)
    AskNotional            string  `json:"askNotional"`            // 卖单名义价值(忽略)
    PositionSide           string  `json:"positionSide"`           // 持仓方向
    PositionAmt            Float64 `json:"positionAmt"`            // 持仓数量
    UpdateTime             Int64   `json:"updateTime"`             // 最后更新时间
}
```

<a name="UserAsset"></a>
## type UserAsset



```go
type UserAsset struct {
    Asset       string  `json:"asset"`       // 资产名称
    Free        Float64 `json:"free"`        // 可用余额
    Locked      Float64 `json:"locked"`      // 锁定余额
    Freeze      Float64 `json:"freeze"`      // 冻结余额
    Withdrawing Float64 `json:"withdrawing"` // 提现中金额

}
```

<a name="UserDataEventType"></a>
## type UserDataEventType



```go
type UserDataEventType string
```

<a name="WalletBalance"></a>
## type WalletBalance

WalletBalance 钱包余额

```go
type WalletBalance struct {
    Activate   bool    `json:"activate"`
    Balance    Float64 `json:"balance"`
    WalletName string  `json:"walletName"`
}
```

<a name="WsAllMarkPriceEvent"></a>
## type WsAllMarkPriceEvent



```go
type WsAllMarkPriceEvent []*MarkPriceEvent
```

<a name="WsAllMiniTickerEvent"></a>
## type WsAllMiniTickerEvent

WsAllMiniTickerEvent define array of websocket market mini\-ticker statistics events

```go
type WsAllMiniTickerEvent []*WsMiniTickerEvent
```

<a name="WsAllTickerEvent"></a>
## type WsAllTickerEvent

WsAllTickerEvent define array of websocket market statistics events

```go
type WsAllTickerEvent []*WsTickerEvent
```

<a name="WsBookTickerEvent"></a>
## type WsBookTickerEvent



```go
type WsBookTickerEvent struct {
    UpdateID Int64   `json:"u"` // 更新ID
    Symbol   string  `json:"s"` // 交易对
    BidPrice Float64 `json:"b"` // 最高买价
    BidQty   Float64 `json:"B"` // 最高买价挂单量
    AskPrice Float64 `json:"a"` // 最低卖价
    AskQty   Float64 `json:"A"` // 最低卖价挂单量
    Time     Int64   `json:"E"` // 事件时间
}
```

<a name="WsMiniTickerEvent"></a>
## type WsMiniTickerEvent

WsMiniTickerEvent define websocket market mini\-ticker statistics event

```go
type WsMiniTickerEvent struct {
    Event       string  `json:"e"` // 事件类型
    Time        Int64   `json:"E"` // 事件时间
    Symbol      string  `json:"s"` // 交易对
    LastPrice   Float64 `json:"c"` // 最新价格
    OpenPrice   Float64 `json:"o"` // 开盘价格
    HighPrice   Float64 `json:"h"` // 最高价格
    LowPrice    Float64 `json:"l"` // 最低价格
    BaseVolume  Float64 `json:"v"` // 成交量
    QuoteVolume Float64 `json:"q"` // 成交额
}
```

<a name="WsPublicBaseService"></a>
## type WsPublicBaseService

WsPublicBaseService 现货公共行情WebSocket服务

```go
type WsPublicBaseService struct {
    KeepInterval time.Duration
    PongTimeout  time.Duration
    // contains filtered or unexported fields
}
```

<a name="WsPublicBaseService.IncrIdx"></a>
### func \(\*WsPublicBaseService\) IncrIdx

```go
func (s *WsPublicBaseService) IncrIdx() int64
```



<a name="WsPublicBaseService.SetHttpClient"></a>
### func \(\*WsPublicBaseService\) SetHttpClient

```go
func (s *WsPublicBaseService) SetHttpClient(httpClient *http.Client) *WsPublicBaseService
```



<a name="WsPublicBaseService.SetLogger"></a>
### func \(\*WsPublicBaseService\) SetLogger

```go
func (s *WsPublicBaseService) SetLogger(logIn func(int, []byte), logOut func(int, []byte)) *WsPublicBaseService
```



<a name="WsPublicBaseService.Start"></a>
### func \(\*WsPublicBaseService\) Start

```go
func (s *WsPublicBaseService) Start() error
```



<a name="WsPublicBaseService.Stop"></a>
### func \(\*WsPublicBaseService\) Stop

```go
func (s *WsPublicBaseService) Stop()
```



<a name="WsPublicBaseService.Subscribe"></a>
### func \(\*WsPublicBaseService\) Subscribe

```go
func (s *WsPublicBaseService) Subscribe(channel string, handler ws.WsHandler) *WsPublicBaseService
```



<a name="WsPublicBaseService.Unsubscribe"></a>
### func \(\*WsPublicBaseService\) Unsubscribe

```go
func (s *WsPublicBaseService) Unsubscribe(channel string) *WsPublicBaseService
```



<a name="WsSpotPublicService"></a>
## type WsSpotPublicService



```go
type WsSpotPublicService struct {
    *WsPublicBaseService
}
```

<a name="NewWsSpotPublicService"></a>
### func NewWsSpotPublicService

```go
func NewWsSpotPublicService(baseUrl string) WsSpotPublicService
```



<a name="WsSpotPublicService.SubscribeAllMiniTicker"></a>
### func \(WsSpotPublicService\) SubscribeAllMiniTicker

```go
func (s WsSpotPublicService) SubscribeAllMiniTicker(handler func(event WsAllMiniTickerEvent))
```



<a name="WsSpotPublicService.SubscribeAllTicker"></a>
### func \(WsSpotPublicService\) SubscribeAllTicker

```go
func (s WsSpotPublicService) SubscribeAllTicker(handler func(event WsAllTickerEvent))
```



<a name="WsSpotPublicService.SubscribeSymbolsBookTicker"></a>
### func \(WsSpotPublicService\) SubscribeSymbolsBookTicker

```go
func (s WsSpotPublicService) SubscribeSymbolsBookTicker(handler func(event WsBookTickerEvent), symbols ...string)
```

SubscribeSymbolsBookTicker 订阅指定交易对的盘口信息

<a name="WsSpotPublicService.SubscribeSymbolsMiniTicker"></a>
### func \(WsSpotPublicService\) SubscribeSymbolsMiniTicker

```go
func (s WsSpotPublicService) SubscribeSymbolsMiniTicker(handler func(event WsMiniTickerEvent), symbols ...string)
```



<a name="WsSpotPublicService.SubscribeSymbolsTicker"></a>
### func \(WsSpotPublicService\) SubscribeSymbolsTicker

```go
func (s WsSpotPublicService) SubscribeSymbolsTicker(handler func(event WsTickerEvent), symbols ...string)
```



<a name="WsSwapPublicService"></a>
## type WsSwapPublicService



```go
type WsSwapPublicService struct {
    *WsPublicBaseService
}
```

<a name="NewWsSwapPublicService"></a>
### func NewWsSwapPublicService

```go
func NewWsSwapPublicService(baseUrl string) WsSwapPublicService
```



<a name="WsSwapPublicService.SubscribeAllMarkPrice"></a>
### func \(WsSwapPublicService\) SubscribeAllMarkPrice

```go
func (s WsSwapPublicService) SubscribeAllMarkPrice(handler func(event WsAllMarkPriceEvent))
```

SubscribeAllMarkPrice 订阅全市场标记价格

<a name="WsSwapPublicService.SubscribeAllMiniTicker"></a>
### func \(WsSwapPublicService\) SubscribeAllMiniTicker

```go
func (s WsSwapPublicService) SubscribeAllMiniTicker(handler func(event WsAllMiniTickerEvent))
```

SubscribeAllMiniTicker 订阅全市场简易信息

<a name="WsSwapPublicService.SubscribeAllTicker"></a>
### func \(WsSwapPublicService\) SubscribeAllTicker

```go
func (s WsSwapPublicService) SubscribeAllTicker(handler func(event WsAllTickerEvent))
```

SubscribeAllTicker 订阅全市场信息

<a name="WsSwapPublicService.SubscribeBookTicker"></a>
### func \(WsSwapPublicService\) SubscribeBookTicker

```go
func (s WsSwapPublicService) SubscribeBookTicker(handler func(event WsBookTickerEvent), symbols ...string)
```

SubscribeBookTicker 订阅指定交易对的盘口信息

<a name="WsSwapPublicService.SubscribeMarkPrice"></a>
### func \(WsSwapPublicService\) SubscribeMarkPrice

```go
func (s WsSwapPublicService) SubscribeMarkPrice(handler func(event MarkPriceEvent), symbols ...string)
```

SubscribeMarkPrice 订阅标记价格

<a name="WsSwapPublicService.SubscribeMarkPrice1s"></a>
### func \(WsSwapPublicService\) SubscribeMarkPrice1s

```go
func (s WsSwapPublicService) SubscribeMarkPrice1s(handler func(event MarkPriceEvent), symbols ...string)
```

SubscribeMarkPrice1s 订阅标记价格

<a name="WsSwapPublicService.SubscribeSymbolsMiniTicker"></a>
### func \(WsSwapPublicService\) SubscribeSymbolsMiniTicker

```go
func (s WsSwapPublicService) SubscribeSymbolsMiniTicker(handler func(event WsMiniTickerEvent), symbols ...string)
```

SubscribeSymbolsMiniTicker 订阅指定交易对的简易信息

<a name="WsSwapPublicService.SubscribeSymbolsTicker"></a>
### func \(WsSwapPublicService\) SubscribeSymbolsTicker

```go
func (s WsSwapPublicService) SubscribeSymbolsTicker(handler func(event WsTickerEvent), symbols ...string)
```



<a name="WsTickerEvent"></a>
## type WsTickerEvent

WsTickerEvent define websocket market statistics event

```go
type WsTickerEvent struct {
    Event              string `json:"e"` // 事件类型
    Time               int64  `json:"E"` // 事件时间
    Symbol             string `json:"s"` // 交易对
    PriceChange        string `json:"p"` // 24小时价格变化
    PriceChangePercent string `json:"P"` // 24小时价格变化（百分比）
    WeightedAvgPrice   string `json:"w"` // 平均价格
    PrevClosePrice     string `json:"x"` // 整整24小时之前，向前数的最后一次成交价格
    LastPrice          string `json:"c"` // 最新成交价格
    CloseQty           string `json:"Q"` // 最新成交交易的成交量
    BidPrice           string `json:"b"` // 目前最高买单价
    BidQty             string `json:"B"` // 目前最高买单价的挂单量
    AskPrice           string `json:"a"` // 目前最低卖单价
    AskQty             string `json:"A"` // 目前最低卖单价的挂单量
    OpenPrice          string `json:"o"` // 开盘价
    HighPrice          string `json:"h"` // 24小时内最高成交价
    LowPrice           string `json:"l"` // 24小时内最低成交价
    BaseVolume         string `json:"v"` // 24小时内成交量
    QuoteVolume        string `json:"q"` // 24小时内成交额
    OpenTime           int64  `json:"O"` // 统计开始时间
    CloseTime          int64  `json:"C"` // 统计结束时间
    FirstID            int64  `json:"F"` // 24小时内第一笔成交交易ID
    LastID             int64  `json:"L"` // 24小时内最后一笔成交交易ID
    Count              int64  `json:"n"` // 24小时内成交数
}
```

# common

```go
import "github.com/youjianglong/exchanges/common"
```

## Index

- [Constants](<#constants>)
- [Variables](<#variables>)
- [func Abs\[T Float64 | Int64 | Mixed\]\(val T\) T](<#Abs>)
- [func BatchConvertToFloat64\(values ...string\) \(\[\]float64, error\)](<#BatchConvertToFloat64>)
- [func BatchConvertToFloat64Map\(values ...string\) \(map\[string\]float64, error\)](<#BatchConvertToFloat64Map>)
- [func CheckProxy\(proxyUrl \*url.URL, request \*http.Request\) \(time.Duration, error\)](<#CheckProxy>)
- [func CloneHttpTransport\(\) \*http.Transport](<#CloneHttpTransport>)
- [func ConvertToFloat64\(value string\) \(float64, error\)](<#ConvertToFloat64>)
- [func ConvertToInt64\(value string\) \(int64, error\)](<#ConvertToInt64>)
- [func FormatFloat64\(value float64\) string](<#FormatFloat64>)
- [func FormatFloat64Percent\(value float64, precision int\) string](<#FormatFloat64Percent>)
- [func IsZeroStr\(s string\) bool](<#IsZeroStr>)
- [func MustConvertToInt64\(value string, logger \*slog.Logger, name string\) int64](<#MustConvertToInt64>)
- [func NewHttpClient\(timeout time.Duration, proxyURL \*url.URL\) \*http.Client](<#NewHttpClient>)
- [func NewHttpClientWithProxy\(proxyUrl \*url.URL\) \(\*http.Client, error\)](<#NewHttpClientWithProxy>)
- [func SplitRange\(startVal int64, endVal int64, step int64\) \[\]\[\]int64](<#SplitRange>)
- [func StandardizeCoin\(coin string\) string](<#StandardizeCoin>)
- [func StandardizeSymbol\(symbol string\) string](<#StandardizeSymbol>)
- [func StrictDecode\(data \[\]byte, target any\) error](<#StrictDecode>)
- [func StrictUnmarshal\(data \[\]byte, target any\) error](<#StrictUnmarshal>)
- [func TcpPing\(proxyUrl \*url.URL\) \(time.Duration, error\)](<#TcpPing>)
- [type Float64](<#Float64>)
  - [func NewFloat64\(f float64, decimals ...int\) Float64](<#NewFloat64>)
  - [func \(f Float64\) Add\(other Float64\) Float64](<#Float64.Add>)
  - [func \(f Float64\) Ceil\(precision int\) Float64](<#Float64.Ceil>)
  - [func \(f Float64\) Float64\(\) \(float64, error\)](<#Float64.Float64>)
  - [func \(f Float64\) Floor\(precision int\) Float64](<#Float64.Floor>)
  - [func \(f Float64\) IsZero\(\) bool](<#Float64.IsZero>)
  - [func \(f Float64\) MarshalJSON\(\) \(\[\]byte, error\)](<#Float64.MarshalJSON>)
  - [func \(f Float64\) Round\(precision int\) Float64](<#Float64.Round>)
  - [func \(f Float64\) String\(\) string](<#Float64.String>)
  - [func \(f Float64\) Sub\(other Float64\) Float64](<#Float64.Sub>)
  - [func \(f \*Float64\) UnmarshalJSON\(data \[\]byte\) error](<#Float64.UnmarshalJSON>)
  - [func \(f Float64\) Value\(\) float64](<#Float64.Value>)
- [type Int64](<#Int64>)
  - [func NewInt64\(i int64\) Int64](<#NewInt64>)
  - [func \(i Int64\) Int64\(\) \(int64, error\)](<#Int64.Int64>)
  - [func \(i Int64\) IsZero\(\) bool](<#Int64.IsZero>)
  - [func \(i Int64\) MarshalJSON\(\) \(\[\]byte, error\)](<#Int64.MarshalJSON>)
  - [func \(i Int64\) String\(\) string](<#Int64.String>)
  - [func \(i \*Int64\) UnmarshalJSON\(data \[\]byte\) error](<#Int64.UnmarshalJSON>)
  - [func \(i Int64\) Value\(\) int64](<#Int64.Value>)
- [type Mixed](<#Mixed>)
  - [func \(i Mixed\) Float64\(\) Float64](<#Mixed.Float64>)
  - [func \(i Mixed\) Int64\(\) Int64](<#Mixed.Int64>)
  - [func \(i Mixed\) IsZero\(\) bool](<#Mixed.IsZero>)
  - [func \(i Mixed\) MarshalJSON\(\) \(\[\]byte, error\)](<#Mixed.MarshalJSON>)
  - [func \(i Mixed\) String\(\) string](<#Mixed.String>)
  - [func \(i \*Mixed\) UnmarshalJSON\(data \[\]byte\) error](<#Mixed.UnmarshalJSON>)
  - [func \(i Mixed\) Value\(\) string](<#Mixed.Value>)
- [type PingResult](<#PingResult>)
- [type ProxyNode](<#ProxyNode>)
- [type ProxyPool](<#ProxyPool>)
  - [func NewProxyPool\(proxyURLs \[\]string\) \(\*ProxyPool, error\)](<#NewProxyPool>)
  - [func \(p \*ProxyPool\) Check\(ping func\(\) error, checkInterval time.Duration\)](<#ProxyPool.Check>)
  - [func \(p \*ProxyPool\) GetCurrentProxy\(\) \*ProxyNode](<#ProxyPool.GetCurrentProxy>)
  - [func \(p \*ProxyPool\) GetCurrentProxyURL\(\) \*url.URL](<#ProxyPool.GetCurrentProxyURL>)
  - [func \(p \*ProxyPool\) Proxy\(\*http.Request\) \(\*url.URL, error\)](<#ProxyPool.Proxy>)
  - [func \(p \*ProxyPool\) SwitchProxy\(\) error](<#ProxyPool.SwitchProxy>)
  - [func \(p \*ProxyPool\) WithLogger\(logger \*slog.Logger\) \*ProxyPool](<#ProxyPool.WithLogger>)
  - [func \(p \*ProxyPool\) WithSetter\(setter ProxyURLSetter\) \*ProxyPool](<#ProxyPool.WithSetter>)
- [type ProxyURLSetter](<#ProxyURLSetter>)
- [type Symbol](<#Symbol>)
  - [func NewSymbol\(symbol string\) \*Symbol](<#NewSymbol>)
  - [func NewSymbolPair\(base, quote string\) \*Symbol](<#NewSymbolPair>)
  - [func \(s \*Symbol\) Equals\(other \*Symbol\) bool](<#Symbol.Equals>)
  - [func \(s \*Symbol\) Format\(spe ...string\) string](<#Symbol.Format>)
  - [func \(s \*Symbol\) String\(\) string](<#Symbol.String>)


## Constants

<a name="USDT"></a>

```go
const (
    USDT = "USDT"
    USDC = "USDC"
    BTC  = "BTC"
    ETH  = "ETH"
    SOL  = "SOL"
)
```

## Variables

<a name="ErrNoProxyAvailable"></a>

```go
var (
    ErrNoProxyAvailable = errors.New("no proxy available") // 没有可用的代理
    ErrProxy            = errors.New("<ProxyError>")       // 代理不可用
)
```

<a name="BaseSymbols"></a>

```go
var BaseSymbols = []string{"USDT", "USDC", "USD", "BTC", "BNB", "ETH"}
```

<a name="ErrEmptyString"></a>

```go
var ErrEmptyString = errors.New("empty string")
```

<a name="Abs"></a>
## func Abs

```go
func Abs[T Float64 | Int64 | Mixed](val T) T
```

Abs 取绝对值

<a name="BatchConvertToFloat64"></a>
## func BatchConvertToFloat64

```go
func BatchConvertToFloat64(values ...string) ([]float64, error)
```

BatchConvertToFloat64 批量转换字符串为float64

<a name="BatchConvertToFloat64Map"></a>
## func BatchConvertToFloat64Map

```go
func BatchConvertToFloat64Map(values ...string) (map[string]float64, error)
```

BatchConvertToFloat64Map 批量转换字符串为float64

<a name="CheckProxy"></a>
## func CheckProxy

```go
func CheckProxy(proxyUrl *url.URL, request *http.Request) (time.Duration, error)
```



<a name="CloneHttpTransport"></a>
## func CloneHttpTransport

```go
func CloneHttpTransport() *http.Transport
```

CloneHttpTransport 克隆一个http.Transport

<a name="ConvertToFloat64"></a>
## func ConvertToFloat64

```go
func ConvertToFloat64(value string) (float64, error)
```

ConvertToFloat64 转换字符串为float64

<a name="ConvertToInt64"></a>
## func ConvertToInt64

```go
func ConvertToInt64(value string) (int64, error)
```

ConvertToInt64 转换字符串为int64

<a name="FormatFloat64"></a>
## func FormatFloat64

```go
func FormatFloat64(value float64) string
```



<a name="FormatFloat64Percent"></a>
## func FormatFloat64Percent

```go
func FormatFloat64Percent(value float64, precision int) string
```



<a name="IsZeroStr"></a>
## func IsZeroStr

```go
func IsZeroStr(s string) bool
```



<a name="MustConvertToInt64"></a>
## func MustConvertToInt64

```go
func MustConvertToInt64(value string, logger *slog.Logger, name string) int64
```



<a name="NewHttpClient"></a>
## func NewHttpClient

```go
func NewHttpClient(timeout time.Duration, proxyURL *url.URL) *http.Client
```

NewHttpClient 创建一个http.Client

<a name="NewHttpClientWithProxy"></a>
## func NewHttpClientWithProxy

```go
func NewHttpClientWithProxy(proxyUrl *url.URL) (*http.Client, error)
```



<a name="SplitRange"></a>
## func SplitRange

```go
func SplitRange(startVal int64, endVal int64, step int64) [][]int64
```

SplitRange 将一个范围分成多个区间

<a name="StandardizeCoin"></a>
## func StandardizeCoin

```go
func StandardizeCoin(coin string) string
```



<a name="StandardizeSymbol"></a>
## func StandardizeSymbol

```go
func StandardizeSymbol(symbol string) string
```

标准格式化交易对

<a name="StrictDecode"></a>
## func StrictDecode

```go
func StrictDecode(data []byte, target any) error
```

StrictDecode 严格大小写敏感的 JSON 反序列化统一入口。 struct 走 StrictUnmarshal，slice 逐元素 StrictUnmarshal，其余 fallback json.Unmarshal。

<a name="StrictUnmarshal"></a>
## func StrictUnmarshal

```go
func StrictUnmarshal(data []byte, target any) error
```

StrictUnmarshal 严格大小写敏感的 JSON 反序列化 自动从结构体的 json tag 中提取字段映射，无需手动写字段映射

<a name="TcpPing"></a>
## func TcpPing

```go
func TcpPing(proxyUrl *url.URL) (time.Duration, error)
```



<a name="Float64"></a>
## type Float64



```go
type Float64 string
```

<a name="NewFloat64"></a>
### func NewFloat64

```go
func NewFloat64(f float64, decimals ...int) Float64
```



<a name="Float64.Add"></a>
### func \(Float64\) Add

```go
func (f Float64) Add(other Float64) Float64
```



<a name="Float64.Ceil"></a>
### func \(Float64\) Ceil

```go
func (f Float64) Ceil(precision int) Float64
```



<a name="Float64.Float64"></a>
### func \(Float64\) Float64

```go
func (f Float64) Float64() (float64, error)
```



<a name="Float64.Floor"></a>
### func \(Float64\) Floor

```go
func (f Float64) Floor(precision int) Float64
```



<a name="Float64.IsZero"></a>
### func \(Float64\) IsZero

```go
func (f Float64) IsZero() bool
```



<a name="Float64.MarshalJSON"></a>
### func \(Float64\) MarshalJSON

```go
func (f Float64) MarshalJSON() ([]byte, error)
```



<a name="Float64.Round"></a>
### func \(Float64\) Round

```go
func (f Float64) Round(precision int) Float64
```



<a name="Float64.String"></a>
### func \(Float64\) String

```go
func (f Float64) String() string
```



<a name="Float64.Sub"></a>
### func \(Float64\) Sub

```go
func (f Float64) Sub(other Float64) Float64
```



<a name="Float64.UnmarshalJSON"></a>
### func \(\*Float64\) UnmarshalJSON

```go
func (f *Float64) UnmarshalJSON(data []byte) error
```



<a name="Float64.Value"></a>
### func \(Float64\) Value

```go
func (f Float64) Value() float64
```



<a name="Int64"></a>
## type Int64



```go
type Int64 string
```

<a name="NewInt64"></a>
### func NewInt64

```go
func NewInt64(i int64) Int64
```



<a name="Int64.Int64"></a>
### func \(Int64\) Int64

```go
func (i Int64) Int64() (int64, error)
```



<a name="Int64.IsZero"></a>
### func \(Int64\) IsZero

```go
func (i Int64) IsZero() bool
```



<a name="Int64.MarshalJSON"></a>
### func \(Int64\) MarshalJSON

```go
func (i Int64) MarshalJSON() ([]byte, error)
```



<a name="Int64.String"></a>
### func \(Int64\) String

```go
func (i Int64) String() string
```



<a name="Int64.UnmarshalJSON"></a>
### func \(\*Int64\) UnmarshalJSON

```go
func (i *Int64) UnmarshalJSON(data []byte) error
```



<a name="Int64.Value"></a>
### func \(Int64\) Value

```go
func (i Int64) Value() int64
```



<a name="Mixed"></a>
## type Mixed



```go
type Mixed string
```

<a name="Mixed.Float64"></a>
### func \(Mixed\) Float64

```go
func (i Mixed) Float64() Float64
```



<a name="Mixed.Int64"></a>
### func \(Mixed\) Int64

```go
func (i Mixed) Int64() Int64
```



<a name="Mixed.IsZero"></a>
### func \(Mixed\) IsZero

```go
func (i Mixed) IsZero() bool
```



<a name="Mixed.MarshalJSON"></a>
### func \(Mixed\) MarshalJSON

```go
func (i Mixed) MarshalJSON() ([]byte, error)
```



<a name="Mixed.String"></a>
### func \(Mixed\) String

```go
func (i Mixed) String() string
```



<a name="Mixed.UnmarshalJSON"></a>
### func \(\*Mixed\) UnmarshalJSON

```go
func (i *Mixed) UnmarshalJSON(data []byte) error
```



<a name="Mixed.Value"></a>
### func \(Mixed\) Value

```go
func (i Mixed) Value() string
```



<a name="PingResult"></a>
## type PingResult



```go
type PingResult struct {
    // contains filtered or unexported fields
}
```

<a name="ProxyNode"></a>
## type ProxyNode



```go
type ProxyNode struct {
    URL         *url.URL
    FailCount   int
    LastChecked time.Time
}
```

<a name="ProxyPool"></a>
## type ProxyPool



```go
type ProxyPool struct {
    // contains filtered or unexported fields
}
```

<a name="NewProxyPool"></a>
### func NewProxyPool

```go
func NewProxyPool(proxyURLs []string) (*ProxyPool, error)
```

创建代理池

<a name="ProxyPool.Check"></a>
### func \(\*ProxyPool\) Check

```go
func (p *ProxyPool) Check(ping func() error, checkInterval time.Duration)
```



<a name="ProxyPool.GetCurrentProxy"></a>
### func \(\*ProxyPool\) GetCurrentProxy

```go
func (p *ProxyPool) GetCurrentProxy() *ProxyNode
```



<a name="ProxyPool.GetCurrentProxyURL"></a>
### func \(\*ProxyPool\) GetCurrentProxyURL

```go
func (p *ProxyPool) GetCurrentProxyURL() *url.URL
```



<a name="ProxyPool.Proxy"></a>
### func \(\*ProxyPool\) Proxy

```go
func (p *ProxyPool) Proxy(*http.Request) (*url.URL, error)
```



<a name="ProxyPool.SwitchProxy"></a>
### func \(\*ProxyPool\) SwitchProxy

```go
func (p *ProxyPool) SwitchProxy() error
```



<a name="ProxyPool.WithLogger"></a>
### func \(\*ProxyPool\) WithLogger

```go
func (p *ProxyPool) WithLogger(logger *slog.Logger) *ProxyPool
```



<a name="ProxyPool.WithSetter"></a>
### func \(\*ProxyPool\) WithSetter

```go
func (p *ProxyPool) WithSetter(setter ProxyURLSetter) *ProxyPool
```



<a name="ProxyURLSetter"></a>
## type ProxyURLSetter



```go
type ProxyURLSetter interface {
    SetProxyURL(proxyUrl *url.URL)
}
```

<a name="Symbol"></a>
## type Symbol



```go
type Symbol struct {
    Symbol string `json:"symbol"`
    Base   string `json:"base"`
    Quote  string `json:"quote"`
}
```

<a name="NewSymbol"></a>
### func NewSymbol

```go
func NewSymbol(symbol string) *Symbol
```



<a name="NewSymbolPair"></a>
### func NewSymbolPair

```go
func NewSymbolPair(base, quote string) *Symbol
```



<a name="Symbol.Equals"></a>
### func \(\*Symbol\) Equals

```go
func (s *Symbol) Equals(other *Symbol) bool
```



<a name="Symbol.Format"></a>
### func \(\*Symbol\) Format

```go
func (s *Symbol) Format(spe ...string) string
```



<a name="Symbol.String"></a>
### func \(\*Symbol\) String

```go
func (s *Symbol) String() string
```



# errorx

```go
import "github.com/youjianglong/exchanges/errorx"
```

## Index

- [func FormatFrames\(frames \[\]runtime.Frame, indent int\) string](<#FormatFrames>)
- [func GetFrame\(skip int\) \*runtime.Frame](<#GetFrame>)
- [func GetFrames\(skip int\) \[\]runtime.Frame](<#GetFrames>)


<a name="FormatFrames"></a>
## func FormatFrames

```go
func FormatFrames(frames []runtime.Frame, indent int) string
```



<a name="GetFrame"></a>
## func GetFrame

```go
func GetFrame(skip int) *runtime.Frame
```



<a name="GetFrames"></a>
## func GetFrames

```go
func GetFrames(skip int) []runtime.Frame
```



# io2

```go
import "github.com/youjianglong/exchanges/io2"
```

## Index

- [Constants](<#constants>)
- [Variables](<#variables>)
- [type Closer](<#Closer>)
- [type CompressFileAdapter](<#CompressFileAdapter>)
  - [func NewCompressFileAdapter\(opts CompressFileAdapterOptions\) \*CompressFileAdapter](<#NewCompressFileAdapter>)
  - [func \(a \*CompressFileAdapter\) ClearAdapter\(\) OutputClearAdapter](<#CompressFileAdapter.ClearAdapter>)
  - [func \(a \*CompressFileAdapter\) SplitAdapter\(\) OutputSplitAdapter](<#CompressFileAdapter.SplitAdapter>)
  - [func \(a \*CompressFileAdapter\) WriteAdapter\(\) OutputWriteAdapter](<#CompressFileAdapter.WriteAdapter>)
- [type CompressFileAdapterOptions](<#CompressFileAdapterOptions>)
- [type Output](<#Output>)
  - [func NewCompressFileOutput\(ctx context.Context, opts CompressFileAdapterOptions\) \*Output](<#NewCompressFileOutput>)
  - [func NewFileOutput\(ctx context.Context, path string, split, expire time.Duration\) \*Output](<#NewFileOutput>)
  - [func \(t \*Output\) Close\(\) error](<#Output.Close>)
  - [func \(t \*Output\) DoClear\(now time.Time\)](<#Output.DoClear>)
  - [func \(t \*Output\) DoSplit\(\)](<#Output.DoSplit>)
  - [func \(t \*Output\) GetWriter\(\) \(Writer, error\)](<#Output.GetWriter>)
  - [func \(t \*Output\) SetAdapter\(a OutputAdapter\)](<#Output.SetAdapter>)
  - [func \(t \*Output\) SetClearAdapter\(a OutputClearAdapter\)](<#Output.SetClearAdapter>)
  - [func \(t \*Output\) SetSplitAdapter\(a OutputSplitAdapter\)](<#Output.SetSplitAdapter>)
  - [func \(t \*Output\) SetWriteAdapter\(a OutputWriteAdapter\)](<#Output.SetWriteAdapter>)
  - [func \(t \*Output\) Write\(b \[\]byte\) \(int, error\)](<#Output.Write>)
- [type OutputAdapter](<#OutputAdapter>)
- [type OutputClearAdapter](<#OutputClearAdapter>)
  - [func OutputFileClearAdapter\(prefix string, expire time.Duration\) OutputClearAdapter](<#OutputFileClearAdapter>)
- [type OutputSplitAdapter](<#OutputSplitAdapter>)
  - [func OutputTickSplitAdapter\(split time.Duration\) OutputSplitAdapter](<#OutputTickSplitAdapter>)
- [type OutputWriteAdapter](<#OutputWriteAdapter>)
  - [func OutputFileAdapter\(prefix string, split time.Duration\) OutputWriteAdapter](<#OutputFileAdapter>)
- [type ReadCloser](<#ReadCloser>)
- [type Reader](<#Reader>)
- [type Splitter](<#Splitter>)
- [type WriteCloser](<#WriteCloser>)
- [type Writer](<#Writer>)


## Constants

<a name="Day"></a>

```go
const (
    Day = time.Hour * 24
)
```

## Variables

<a name="MultiReader"></a>MultiReader io.MultiReader

```go
var MultiReader = io.MultiReader
```

<a name="MultiWriter"></a>MultiWriter io.MultiWriter

```go
var MultiWriter = io.MultiWriter
```

<a name="Closer"></a>
## type Closer

Closer io.Closer

```go
type Closer = io.Closer
```

<a name="CompressFileAdapter"></a>
## type CompressFileAdapter

CompressFileAdapter 带压缩功能的文件适配器

```go
type CompressFileAdapter struct {
    // contains filtered or unexported fields
}
```

<a name="NewCompressFileAdapter"></a>
### func NewCompressFileAdapter

```go
func NewCompressFileAdapter(opts CompressFileAdapterOptions) *CompressFileAdapter
```

NewCompressFileAdapter 创建压缩文件适配器

<a name="CompressFileAdapter.ClearAdapter"></a>
### func \(\*CompressFileAdapter\) ClearAdapter

```go
func (a *CompressFileAdapter) ClearAdapter() OutputClearAdapter
```

ClearAdapter 实现 OutputAdapter 接口

<a name="CompressFileAdapter.SplitAdapter"></a>
### func \(\*CompressFileAdapter\) SplitAdapter

```go
func (a *CompressFileAdapter) SplitAdapter() OutputSplitAdapter
```

SplitAdapter 实现 OutputAdapter 接口

<a name="CompressFileAdapter.WriteAdapter"></a>
### func \(\*CompressFileAdapter\) WriteAdapter

```go
func (a *CompressFileAdapter) WriteAdapter() OutputWriteAdapter
```

WriteAdapter 实现 OutputAdapter 接口

<a name="CompressFileAdapterOptions"></a>
## type CompressFileAdapterOptions

CompressFileAdapterOptions 压缩文件适配器选项

```go
type CompressFileAdapterOptions struct {
    Prefix         string        // 文件前缀
    Split          time.Duration // 切割间隔
    Expire         time.Duration // 原始文件过期时间（0表示不删除原始文件）
    CompressAge    time.Duration // 文件达到此年龄后压缩（0表示不压缩）
    CompressExpire time.Duration // 压缩文件过期时间（0表示不删除压缩文件）
    KeepOriginal   bool          // 压缩后是否保留原始文件
    CompressLevel  int           // 压缩级别（-1=默认, 0=不压缩, 1=最快, 9=最优）
}
```

<a name="Output"></a>
## type Output

Output 输出到文件

```go
type Output struct {
    // contains filtered or unexported fields
}
```

<a name="NewCompressFileOutput"></a>
### func NewCompressFileOutput

```go
func NewCompressFileOutput(ctx context.Context, opts CompressFileAdapterOptions) *Output
```

NewCompressFileOutput 构造带压缩功能的文件输出结构

<a name="NewFileOutput"></a>
### func NewFileOutput

```go
func NewFileOutput(ctx context.Context, path string, split, expire time.Duration) *Output
```

NewFileOutput 构造输入文件结构

<a name="Output.Close"></a>
### func \(\*Output\) Close

```go
func (t *Output) Close() error
```

Close 关闭输出

<a name="Output.DoClear"></a>
### func \(\*Output\) DoClear

```go
func (t *Output) DoClear(now time.Time)
```

DoClear 执行清理

<a name="Output.DoSplit"></a>
### func \(\*Output\) DoSplit

```go
func (t *Output) DoSplit()
```

DoSplit 执行切割

<a name="Output.GetWriter"></a>
### func \(\*Output\) GetWriter

```go
func (t *Output) GetWriter() (Writer, error)
```

GetWriter 读取文件句柄

<a name="Output.SetAdapter"></a>
### func \(\*Output\) SetAdapter

```go
func (t *Output) SetAdapter(a OutputAdapter)
```



<a name="Output.SetClearAdapter"></a>
### func \(\*Output\) SetClearAdapter

```go
func (t *Output) SetClearAdapter(a OutputClearAdapter)
```

SetClearAdapter 设置清理适配器

<a name="Output.SetSplitAdapter"></a>
### func \(\*Output\) SetSplitAdapter

```go
func (t *Output) SetSplitAdapter(a OutputSplitAdapter)
```

SetSplitAdapter 设置切割适配器

<a name="Output.SetWriteAdapter"></a>
### func \(\*Output\) SetWriteAdapter

```go
func (t *Output) SetWriteAdapter(a OutputWriteAdapter)
```

SetWriteAdapter 设置写入适配器

<a name="Output.Write"></a>
### func \(\*Output\) Write

```go
func (t *Output) Write(b []byte) (int, error)
```



<a name="OutputAdapter"></a>
## type OutputAdapter



```go
type OutputAdapter interface {
    WriteAdapter() OutputWriteAdapter
    ClearAdapter() OutputClearAdapter
    SplitAdapter() OutputSplitAdapter
}
```

<a name="OutputClearAdapter"></a>
## type OutputClearAdapter



```go
type OutputClearAdapter func(time.Time) error
```

<a name="OutputFileClearAdapter"></a>
### func OutputFileClearAdapter

```go
func OutputFileClearAdapter(prefix string, expire time.Duration) OutputClearAdapter
```

OutputFileClearAdapter 默认文件清理适配器

<a name="OutputSplitAdapter"></a>
## type OutputSplitAdapter



```go
type OutputSplitAdapter func(context.Context, Splitter)
```

<a name="OutputTickSplitAdapter"></a>
### func OutputTickSplitAdapter

```go
func OutputTickSplitAdapter(split time.Duration) OutputSplitAdapter
```

OutputTickSplitAdapter 定时切割适配器

<a name="OutputWriteAdapter"></a>
## type OutputWriteAdapter



```go
type OutputWriteAdapter func(time.Time) (WriteCloser, error)
```

<a name="OutputFileAdapter"></a>
### func OutputFileAdapter

```go
func OutputFileAdapter(prefix string, split time.Duration) OutputWriteAdapter
```

OutputFileAdapter 默认文件适配器

<a name="ReadCloser"></a>
## type ReadCloser

ReadCloser io.ReadCloser

```go
type ReadCloser = io.ReadCloser
```

<a name="Reader"></a>
## type Reader

Reader io.Reader

```go
type Reader = io.Reader
```

<a name="Splitter"></a>
## type Splitter



```go
type Splitter interface {
    DoSplit()
    DoClear(now time.Time)
}
```

<a name="WriteCloser"></a>
## type WriteCloser

WriteCloser io.WriteCloser

```go
type WriteCloser = io.WriteCloser
```

<a name="Writer"></a>
## type Writer

Writer io.Writer

```go
type Writer = io.Writer
```

# mapx

```go
import "github.com/youjianglong/exchanges/mapx"
```

## Index

- [func KeyValues\[K comparable, V any\]\(m map\[K\]V\) \(\[\]K, \[\]V\)](<#KeyValues>)
- [func Keys\[K comparable, V any\]\(m map\[K\]V\) \[\]K](<#Keys>)
- [func Values\[K comparable, V any\]\(m map\[K\]V\) \[\]V](<#Values>)
- [type Map](<#Map>)
  - [func New\[K comparable, V any\]\(\) \*Map\[K, V\]](<#New>)
  - [func NewMap\[K comparable, V any\]\(\) \*Map\[K, V\]](<#NewMap>)
  - [func NewMapWithCapacity\[K comparable, V any\]\(capacity int\) \*Map\[K, V\]](<#NewMapWithCapacity>)
  - [func \(m \*Map\[K, V\]\) CopyFrom\(data map\[K\]V\)](<#Map[K, V].CopyFrom>)
  - [func \(m \*Map\[K, V\]\) Delete\(key K\)](<#Map[K, V].Delete>)
  - [func \(m \*Map\[K, V\]\) Get\(key K\) \(V, bool\)](<#Map[K, V].Get>)
  - [func \(m \*Map\[K, V\]\) GetOrCreate\(key K, c func\(\) V\) V](<#Map[K, V].GetOrCreate>)
  - [func \(m \*Map\[K, V\]\) Has\(key K\) bool](<#Map[K, V].Has>)
  - [func \(m \*Map\[K, V\]\) KeyValues\(\) \(\[\]K, \[\]V\)](<#Map[K, V].KeyValues>)
  - [func \(m \*Map\[K, V\]\) Keys\(\) \[\]K](<#Map[K, V].Keys>)
  - [func \(m \*Map\[K, V\]\) Len\(\) int](<#Map[K, V].Len>)
  - [func \(m \*Map\[K, V\]\) MarshalJSON\(\) \(\[\]byte, error\)](<#Map[K, V].MarshalJSON>)
  - [func \(m \*Map\[K, V\]\) Pop\(key K\) \(V, bool\)](<#Map[K, V].Pop>)
  - [func \(m \*Map\[K, V\]\) Range\(h func\(K, V\) bool\)](<#Map[K, V].Range>)
  - [func \(m \*Map\[K, V\]\) Reset\(data map\[K\]V\)](<#Map[K, V].Reset>)
  - [func \(m \*Map\[K, V\]\) Set\(key K, value V\)](<#Map[K, V].Set>)
  - [func \(m \*Map\[K, V\]\) ToMap\(\) map\[K\]V](<#Map[K, V].ToMap>)
  - [func \(m \*Map\[K, V\]\) Traversal\(h func\(K, V\) bool\)](<#Map[K, V].Traversal>)
  - [func \(m \*Map\[K, V\]\) UnmarshalJSON\(b \[\]byte\) error](<#Map[K, V].UnmarshalJSON>)
  - [func \(m \*Map\[K, V\]\) Update\(key K, c func\(V\) V\)](<#Map[K, V].Update>)
  - [func \(m \*Map\[K, V\]\) Values\(\) \[\]V](<#Map[K, V].Values>)


<a name="KeyValues"></a>
## func KeyValues

```go
func KeyValues[K comparable, V any](m map[K]V) ([]K, []V)
```



<a name="Keys"></a>
## func Keys

```go
func Keys[K comparable, V any](m map[K]V) []K
```



<a name="Values"></a>
## func Values

```go
func Values[K comparable, V any](m map[K]V) []V
```



<a name="Map"></a>
## type Map



```go
type Map[K comparable, V any] struct {
    // contains filtered or unexported fields
}
```

<a name="New"></a>
### func New

```go
func New[K comparable, V any]() *Map[K, V]
```



<a name="NewMap"></a>
### func NewMap

```go
func NewMap[K comparable, V any]() *Map[K, V]
```



<a name="NewMapWithCapacity"></a>
### func NewMapWithCapacity

```go
func NewMapWithCapacity[K comparable, V any](capacity int) *Map[K, V]
```



<a name="Map[K, V].CopyFrom"></a>
### func \(\*Map\[K, V\]\) CopyFrom

```go
func (m *Map[K, V]) CopyFrom(data map[K]V)
```



<a name="Map[K, V].Delete"></a>
### func \(\*Map\[K, V\]\) Delete

```go
func (m *Map[K, V]) Delete(key K)
```



<a name="Map[K, V].Get"></a>
### func \(\*Map\[K, V\]\) Get

```go
func (m *Map[K, V]) Get(key K) (V, bool)
```



<a name="Map[K, V].GetOrCreate"></a>
### func \(\*Map\[K, V\]\) GetOrCreate

```go
func (m *Map[K, V]) GetOrCreate(key K, c func() V) V
```



<a name="Map[K, V].Has"></a>
### func \(\*Map\[K, V\]\) Has

```go
func (m *Map[K, V]) Has(key K) bool
```



<a name="Map[K, V].KeyValues"></a>
### func \(\*Map\[K, V\]\) KeyValues

```go
func (m *Map[K, V]) KeyValues() ([]K, []V)
```



<a name="Map[K, V].Keys"></a>
### func \(\*Map\[K, V\]\) Keys

```go
func (m *Map[K, V]) Keys() []K
```



<a name="Map[K, V].Len"></a>
### func \(\*Map\[K, V\]\) Len

```go
func (m *Map[K, V]) Len() int
```



<a name="Map[K, V].MarshalJSON"></a>
### func \(\*Map\[K, V\]\) MarshalJSON

```go
func (m *Map[K, V]) MarshalJSON() ([]byte, error)
```



<a name="Map[K, V].Pop"></a>
### func \(\*Map\[K, V\]\) Pop

```go
func (m *Map[K, V]) Pop(key K) (V, bool)
```



<a name="Map[K, V].Range"></a>
### func \(\*Map\[K, V\]\) Range

```go
func (m *Map[K, V]) Range(h func(K, V) bool)
```



<a name="Map[K, V].Reset"></a>
### func \(\*Map\[K, V\]\) Reset

```go
func (m *Map[K, V]) Reset(data map[K]V)
```



<a name="Map[K, V].Set"></a>
### func \(\*Map\[K, V\]\) Set

```go
func (m *Map[K, V]) Set(key K, value V)
```



<a name="Map[K, V].ToMap"></a>
### func \(\*Map\[K, V\]\) ToMap

```go
func (m *Map[K, V]) ToMap() map[K]V
```



<a name="Map[K, V].Traversal"></a>
### func \(\*Map\[K, V\]\) Traversal

```go
func (m *Map[K, V]) Traversal(h func(K, V) bool)
```



<a name="Map[K, V].UnmarshalJSON"></a>
### func \(\*Map\[K, V\]\) UnmarshalJSON

```go
func (m *Map[K, V]) UnmarshalJSON(b []byte) error
```



<a name="Map[K, V].Update"></a>
### func \(\*Map\[K, V\]\) Update

```go
func (m *Map[K, V]) Update(key K, c func(V) V)
```



<a name="Map[K, V].Values"></a>
### func \(\*Map\[K, V\]\) Values

```go
func (m *Map[K, V]) Values() []V
```



# okx

```go
import "github.com/youjianglong/exchanges/okx"
```

## Index

- [Constants](<#constants>)
- [Variables](<#variables>)
- [func AsProxyError\(e error\) \(error, bool\)](<#AsProxyError>)
- [func ConvertProxyError\(e error\) error](<#ConvertProxyError>)
- [func IsAPIError\(e error\) bool](<#IsAPIError>)
- [func ToInstId\(symbol \*common.Symbol\) string](<#ToInstId>)
- [func ToSymbol\(instId string\) \*common.Symbol](<#ToSymbol>)
- [type APIError](<#APIError>)
  - [func \(e APIError\) Error\(\) string](<#APIError.Error>)
  - [func \(e APIError\) IsValid\(\) bool](<#APIError.IsValid>)
- [type AccountBalance](<#AccountBalance>)
- [type AccountBalanceDetail](<#AccountBalanceDetail>)
- [type AccountPosition](<#AccountPosition>)
- [type AdapterWebsocket](<#AdapterWebsocket>)
  - [func NewAdapterWebsocket\(wsUrl string, auth \*common.Auth, logger \*slog.Logger, httpClient \*http.Client\) \*AdapterWebsocket](<#NewAdapterWebsocket>)
  - [func \(a \*AdapterWebsocket\) Login\(\) error](<#AdapterWebsocket.Login>)
  - [func \(a \*AdapterWebsocket\) Operate\(op string, args ...any\) error](<#AdapterWebsocket.Operate>)
  - [func \(a \*AdapterWebsocket\) RegisterHandler\(channel string, handler func\(\*wsEvent\)\) \*AdapterWebsocket](<#AdapterWebsocket.RegisterHandler>)
  - [func \(a \*AdapterWebsocket\) RegisterHandlers\(handlers map\[string\]func\(\*wsEvent\)\) \*AdapterWebsocket](<#AdapterWebsocket.RegisterHandlers>)
  - [func \(a \*AdapterWebsocket\) Send\(op string, args ...any\) error](<#AdapterWebsocket.Send>)
  - [func \(a \*AdapterWebsocket\) SetLogger\(logIn func\(msgType int, msg \[\]byte\), logOut func\(msgType int, msg \[\]byte\)\) \*AdapterWebsocket](<#AdapterWebsocket.SetLogger>)
  - [func \(a \*AdapterWebsocket\) Start\(\) error](<#AdapterWebsocket.Start>)
  - [func \(a \*AdapterWebsocket\) Stop\(\)](<#AdapterWebsocket.Stop>)
  - [func \(a \*AdapterWebsocket\) Subscribe\(args ...any\) error](<#AdapterWebsocket.Subscribe>)
- [type ArgsLogin](<#ArgsLogin>)
  - [func NewArgsLoginFromAuth\(auth \*common.Auth\) \*ArgsLogin](<#NewArgsLoginFromAuth>)
- [type AssetBalance](<#AssetBalance>)
- [type AssetDepositHistory](<#AssetDepositHistory>)
- [type AssetWithdrawHistory](<#AssetWithdrawHistory>)
- [type CancelOrderService](<#CancelOrderService>)
  - [func \(s \*CancelOrderService\) Do\(ctx context.Context, opts ...RequestOption\) \(\*OrderResult, error\)](<#CancelOrderService.Do>)
  - [func \(s \*CancelOrderService\) OrdId\(ordId string\) \*CancelOrderService](<#CancelOrderService.OrdId>)
- [type Client](<#Client>)
  - [func NewClient\(apiKey, secretKey, passphrase string\) \*Client](<#NewClient>)
  - [func NewClientWithHttpClient\(apiKey, secretKey, passphrase string, test bool, httpClient \*http.Client\) \*Client](<#NewClientWithHttpClient>)
  - [func NewTestClient\(apiKey, secretKey, passphrase string\) \*Client](<#NewTestClient>)
  - [func \(c \*Client\) NewCancelOrderService\(instId string\) \*CancelOrderService](<#Client.NewCancelOrderService>)
  - [func \(c \*Client\) NewGetAccountBalanceService\(\) \*GetAccountBalanceService](<#Client.NewGetAccountBalanceService>)
  - [func \(c \*Client\) NewGetAccountPositionsService\(\) \*GetAccountPositionsService](<#Client.NewGetAccountPositionsService>)
  - [func \(c \*Client\) NewGetAssetBalancesService\(\) \*GetAssetBalancesService](<#Client.NewGetAssetBalancesService>)
  - [func \(c \*Client\) NewGetAssetDepositHistoryService\(\) \*GetAssetDepositHistoryService](<#Client.NewGetAssetDepositHistoryService>)
  - [func \(c \*Client\) NewGetAssetWithdrawHistoryService\(\) \*GetAssetWithdrawHistoryService](<#Client.NewGetAssetWithdrawHistoryService>)
  - [func \(c \*Client\) NewGetFundingRateService\(instId string\) \*GetFundingRateService](<#Client.NewGetFundingRateService>)
  - [func \(c \*Client\) NewGetInstrumentsService\(instType string\) \*GetInstrumentsService](<#Client.NewGetInstrumentsService>)
  - [func \(c \*Client\) NewGetMarkPriceService\(instType string\) \*GetMarkPriceService](<#Client.NewGetMarkPriceService>)
  - [func \(c \*Client\) NewGetMarketIndexTickersService\(\) \*GetMarketIndexTickersService](<#Client.NewGetMarketIndexTickersService>)
  - [func \(c \*Client\) NewGetMarketTickersService\(instType string\) \*GetMarketTickersService](<#Client.NewGetMarketTickersService>)
  - [func \(c \*Client\) NewGetOrderHistoryService\(\) \*GetOrderHistoryService](<#Client.NewGetOrderHistoryService>)
  - [func \(c \*Client\) NewGetOrderService\(instId string\) \*GetOrderService](<#Client.NewGetOrderService>)
  - [func \(c \*Client\) NewGetOrdersPendingService\(\) \*GetOrdersPendingService](<#Client.NewGetOrdersPendingService>)
  - [func \(c \*Client\) NewGetPriceLimitService\(instId string\) \*GetPriceLimitService](<#Client.NewGetPriceLimitService>)
  - [func \(c \*Client\) NewGetTradeFillService\(\) \*GetTradeFillService](<#Client.NewGetTradeFillService>)
  - [func \(c \*Client\) NewPingService\(\) \*PingService](<#Client.NewPingService>)
  - [func \(c \*Client\) NewSetAccountLeverageService\(lever string, mgnMode string\) \*SetAccountLeverageService](<#Client.NewSetAccountLeverageService>)
  - [func \(c \*Client\) NewTradeOrderService\(instId string, tdMode string, side string, ordType string, sz string\) \*TradeOrderService](<#Client.NewTradeOrderService>)
  - [func \(c \*Client\) SetProxyURL\(proxyURL \*url.URL\)](<#Client.SetProxyURL>)
  - [func \(c \*Client\) WithHttpClient\(httpClient \*http.Client\) \*Client](<#Client.WithHttpClient>)
- [type FundingRate](<#FundingRate>)
- [type GetAccountBalanceService](<#GetAccountBalanceService>)
  - [func \(s \*GetAccountBalanceService\) Do\(ctx context.Context, opts ...RequestOption\) \(\*AccountBalance, error\)](<#GetAccountBalanceService.Do>)
- [type GetAccountPositionsService](<#GetAccountPositionsService>)
  - [func \(s \*GetAccountPositionsService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]AccountPosition, error\)](<#GetAccountPositionsService.Do>)
- [type GetAssetBalancesService](<#GetAssetBalancesService>)
  - [func \(s \*GetAssetBalancesService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]AssetBalance, error\)](<#GetAssetBalancesService.Do>)
- [type GetAssetDepositHistoryService](<#GetAssetDepositHistoryService>)
  - [func \(s \*GetAssetDepositHistoryService\) After\(after string\) \*GetAssetDepositHistoryService](<#GetAssetDepositHistoryService.After>)
  - [func \(s \*GetAssetDepositHistoryService\) Before\(before string\) \*GetAssetDepositHistoryService](<#GetAssetDepositHistoryService.Before>)
  - [func \(s \*GetAssetDepositHistoryService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]AssetDepositHistory, error\)](<#GetAssetDepositHistoryService.Do>)
  - [func \(s \*GetAssetDepositHistoryService\) EndTime\(endTime int64\) \*GetAssetDepositHistoryService](<#GetAssetDepositHistoryService.EndTime>)
  - [func \(s \*GetAssetDepositHistoryService\) Limit\(limit int\) \*GetAssetDepositHistoryService](<#GetAssetDepositHistoryService.Limit>)
  - [func \(s \*GetAssetDepositHistoryService\) StartTime\(startTime int64\) \*GetAssetDepositHistoryService](<#GetAssetDepositHistoryService.StartTime>)
  - [func \(s \*GetAssetDepositHistoryService\) State\(state string\) \*GetAssetDepositHistoryService](<#GetAssetDepositHistoryService.State>)
- [type GetAssetWithdrawHistoryService](<#GetAssetWithdrawHistoryService>)
  - [func \(s \*GetAssetWithdrawHistoryService\) After\(after string\) \*GetAssetWithdrawHistoryService](<#GetAssetWithdrawHistoryService.After>)
  - [func \(s \*GetAssetWithdrawHistoryService\) Before\(before string\) \*GetAssetWithdrawHistoryService](<#GetAssetWithdrawHistoryService.Before>)
  - [func \(s \*GetAssetWithdrawHistoryService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]AssetWithdrawHistory, error\)](<#GetAssetWithdrawHistoryService.Do>)
  - [func \(s \*GetAssetWithdrawHistoryService\) EndTime\(endTime int64\) \*GetAssetWithdrawHistoryService](<#GetAssetWithdrawHistoryService.EndTime>)
  - [func \(s \*GetAssetWithdrawHistoryService\) Limit\(limit int\) \*GetAssetWithdrawHistoryService](<#GetAssetWithdrawHistoryService.Limit>)
  - [func \(s \*GetAssetWithdrawHistoryService\) StartTime\(startTime int64\) \*GetAssetWithdrawHistoryService](<#GetAssetWithdrawHistoryService.StartTime>)
  - [func \(s \*GetAssetWithdrawHistoryService\) State\(state string\) \*GetAssetWithdrawHistoryService](<#GetAssetWithdrawHistoryService.State>)
- [type GetFundingRateService](<#GetFundingRateService>)
  - [func \(s \*GetFundingRateService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*FundingRate, error\)](<#GetFundingRateService.Do>)
- [type GetInstrumentsService](<#GetInstrumentsService>)
  - [func \(s \*GetInstrumentsService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*Instrument, error\)](<#GetInstrumentsService.Do>)
  - [func \(s \*GetInstrumentsService\) InstFamily\(instFamily string\) \*GetInstrumentsService](<#GetInstrumentsService.InstFamily>)
  - [func \(s \*GetInstrumentsService\) InstId\(instId string\) \*GetInstrumentsService](<#GetInstrumentsService.InstId>)
  - [func \(s \*GetInstrumentsService\) Uly\(uly string\) \*GetInstrumentsService](<#GetInstrumentsService.Uly>)
- [type GetMarkPriceService](<#GetMarkPriceService>)
  - [func \(s \*GetMarkPriceService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*MarkPrice, error\)](<#GetMarkPriceService.Do>)
  - [func \(s \*GetMarkPriceService\) InstFamily\(instFamily string\) \*GetMarkPriceService](<#GetMarkPriceService.InstFamily>)
  - [func \(s \*GetMarkPriceService\) InstId\(instId string\) \*GetMarkPriceService](<#GetMarkPriceService.InstId>)
- [type GetMarketIndexTickersService](<#GetMarketIndexTickersService>)
  - [func \(s \*GetMarketIndexTickersService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*MarketIndexTicker, error\)](<#GetMarketIndexTickersService.Do>)
  - [func \(s \*GetMarketIndexTickersService\) InstId\(instId string\) \*GetMarketIndexTickersService](<#GetMarketIndexTickersService.InstId>)
  - [func \(s \*GetMarketIndexTickersService\) QuoteCcy\(quoteCcy string\) \*GetMarketIndexTickersService](<#GetMarketIndexTickersService.QuoteCcy>)
- [type GetMarketTickersService](<#GetMarketTickersService>)
  - [func \(s \*GetMarketTickersService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*MarketTicker, error\)](<#GetMarketTickersService.Do>)
  - [func \(s \*GetMarketTickersService\) InstFamily\(instFamily string\) \*GetMarketTickersService](<#GetMarketTickersService.InstFamily>)
  - [func \(s \*GetMarketTickersService\) Uly\(uly string\) \*GetMarketTickersService](<#GetMarketTickersService.Uly>)
- [type GetOrderHistoryService](<#GetOrderHistoryService>)
  - [func \(s \*GetOrderHistoryService\) After\(after string\) \*GetOrderHistoryService](<#GetOrderHistoryService.After>)
  - [func \(s \*GetOrderHistoryService\) Archive\(archive bool\) \*GetOrderHistoryService](<#GetOrderHistoryService.Archive>)
  - [func \(s \*GetOrderHistoryService\) Before\(before string\) \*GetOrderHistoryService](<#GetOrderHistoryService.Before>)
  - [func \(s \*GetOrderHistoryService\) Begin\(begin int64\) \*GetOrderHistoryService](<#GetOrderHistoryService.Begin>)
  - [func \(s \*GetOrderHistoryService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*Order, error\)](<#GetOrderHistoryService.Do>)
  - [func \(s \*GetOrderHistoryService\) End\(end int64\) \*GetOrderHistoryService](<#GetOrderHistoryService.End>)
  - [func \(s \*GetOrderHistoryService\) InstId\(instId string\) \*GetOrderHistoryService](<#GetOrderHistoryService.InstId>)
  - [func \(s \*GetOrderHistoryService\) InstType\(instType string\) \*GetOrderHistoryService](<#GetOrderHistoryService.InstType>)
  - [func \(s \*GetOrderHistoryService\) Limit\(limit int\) \*GetOrderHistoryService](<#GetOrderHistoryService.Limit>)
  - [func \(s \*GetOrderHistoryService\) State\(state string\) \*GetOrderHistoryService](<#GetOrderHistoryService.State>)
- [type GetOrderService](<#GetOrderService>)
  - [func \(s \*GetOrderService\) ClOrdId\(clOrdId string\) \*GetOrderService](<#GetOrderService.ClOrdId>)
  - [func \(s \*GetOrderService\) Do\(ctx context.Context, opts ...RequestOption\) \(\*Order, error\)](<#GetOrderService.Do>)
  - [func \(s \*GetOrderService\) OrdId\(ordId string\) \*GetOrderService](<#GetOrderService.OrdId>)
- [type GetOrdersPendingService](<#GetOrdersPendingService>)
  - [func \(s \*GetOrdersPendingService\) After\(after string\) \*GetOrdersPendingService](<#GetOrdersPendingService.After>)
  - [func \(s \*GetOrdersPendingService\) Before\(before string\) \*GetOrdersPendingService](<#GetOrdersPendingService.Before>)
  - [func \(s \*GetOrdersPendingService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*PendingOrder, error\)](<#GetOrdersPendingService.Do>)
  - [func \(s \*GetOrdersPendingService\) InstId\(instId string\) \*GetOrdersPendingService](<#GetOrdersPendingService.InstId>)
  - [func \(s \*GetOrdersPendingService\) InstType\(instType string\) \*GetOrdersPendingService](<#GetOrdersPendingService.InstType>)
  - [func \(s \*GetOrdersPendingService\) Limit\(limit int\) \*GetOrdersPendingService](<#GetOrdersPendingService.Limit>)
  - [func \(s \*GetOrdersPendingService\) OrdType\(ordType string\) \*GetOrdersPendingService](<#GetOrdersPendingService.OrdType>)
  - [func \(s \*GetOrdersPendingService\) State\(state string\) \*GetOrdersPendingService](<#GetOrdersPendingService.State>)
- [type GetPriceLimitService](<#GetPriceLimitService>)
  - [func \(s \*GetPriceLimitService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]\*PriceLimit, error\)](<#GetPriceLimitService.Do>)
- [type GetTradeFillService](<#GetTradeFillService>)
  - [func \(s \*GetTradeFillService\) After\(after string\) \*GetTradeFillService](<#GetTradeFillService.After>)
  - [func \(s \*GetTradeFillService\) Before\(before string\) \*GetTradeFillService](<#GetTradeFillService.Before>)
  - [func \(s \*GetTradeFillService\) Begin\(begin string\) \*GetTradeFillService](<#GetTradeFillService.Begin>)
  - [func \(s \*GetTradeFillService\) Do\(ctx context.Context, opts ...RequestOption\) \(\[\]TradeFill, error\)](<#GetTradeFillService.Do>)
  - [func \(s \*GetTradeFillService\) End\(end string\) \*GetTradeFillService](<#GetTradeFillService.End>)
  - [func \(s \*GetTradeFillService\) InstFamily\(instFamily string\) \*GetTradeFillService](<#GetTradeFillService.InstFamily>)
  - [func \(s \*GetTradeFillService\) InstId\(instId string\) \*GetTradeFillService](<#GetTradeFillService.InstId>)
  - [func \(s \*GetTradeFillService\) InstType\(instType string\) \*GetTradeFillService](<#GetTradeFillService.InstType>)
  - [func \(s \*GetTradeFillService\) Limit\(limit string\) \*GetTradeFillService](<#GetTradeFillService.Limit>)
  - [func \(s \*GetTradeFillService\) OrdId\(ordId string\) \*GetTradeFillService](<#GetTradeFillService.OrdId>)
  - [func \(s \*GetTradeFillService\) SubType\(subType string\) \*GetTradeFillService](<#GetTradeFillService.SubType>)
- [type H](<#H>)
- [type Instrument](<#Instrument>)
- [type MarkPrice](<#MarkPrice>)
- [type MarketIndexTicker](<#MarketIndexTicker>)
- [type MarketTicker](<#MarketTicker>)
- [type OkxPublicStreamAdapter](<#OkxPublicStreamAdapter>)
  - [func NewOkxPublicStreamAdapter\(proxyUrl \*url.URL\) \*OkxPublicStreamAdapter](<#NewOkxPublicStreamAdapter>)
  - [func \(s \*OkxPublicStreamAdapter\) Start\(\) error](<#OkxPublicStreamAdapter.Start>)
  - [func \(s \*OkxPublicStreamAdapter\) Stop\(\)](<#OkxPublicStreamAdapter.Stop>)
- [type OperatePublicArg](<#OperatePublicArg>)
- [type OperateRequest](<#OperateRequest>)
- [type Order](<#Order>)
- [type OrderResult](<#OrderResult>)
- [type PendingOrder](<#PendingOrder>)
- [type PingService](<#PingService>)
  - [func \(s \*PingService\) Do\(ctx context.Context\) error](<#PingService.Do>)
- [type PriceLimit](<#PriceLimit>)
- [type RequestOption](<#RequestOption>)
  - [func WithHeader\(key, value string, replace bool\) RequestOption](<#WithHeader>)
  - [func WithHeaders\(header http.Header\) RequestOption](<#WithHeaders>)
- [type Response](<#Response>)
  - [func \(resp Response\) Error\(\) error](<#Response.Error>)
  - [func \(resp Response\) IsError\(\) bool](<#Response.IsError>)
- [type SetAccountLeverageService](<#SetAccountLeverageService>)
  - [func \(s \*SetAccountLeverageService\) Ccy\(ccy string\) \*SetAccountLeverageService](<#SetAccountLeverageService.Ccy>)
  - [func \(s \*SetAccountLeverageService\) Do\(ctx context.Context, opts ...RequestOption\) error](<#SetAccountLeverageService.Do>)
  - [func \(s \*SetAccountLeverageService\) InstId\(instId string\) \*SetAccountLeverageService](<#SetAccountLeverageService.InstId>)
  - [func \(s \*SetAccountLeverageService\) PosSide\(posSide string\) \*SetAccountLeverageService](<#SetAccountLeverageService.PosSide>)
- [type TradeFill](<#TradeFill>)
- [type TradeOrderService](<#TradeOrderService>)
  - [func \(s \*TradeOrderService\) ClOrdId\(clOrdId string\) \*TradeOrderService](<#TradeOrderService.ClOrdId>)
  - [func \(s \*TradeOrderService\) Do\(ctx context.Context, opts ...RequestOption\) \(\*OrderResult, error\)](<#TradeOrderService.Do>)
  - [func \(s \*TradeOrderService\) PosSide\(posSide string\) \*TradeOrderService](<#TradeOrderService.PosSide>)
  - [func \(s \*TradeOrderService\) Px\(px string\) \*TradeOrderService](<#TradeOrderService.Px>)
  - [func \(s \*TradeOrderService\) ReduceOnly\(reduceOnly bool\) \*TradeOrderService](<#TradeOrderService.ReduceOnly>)


## Constants

<a name="PingTimeout"></a>

```go
const (
    PingTimeout  = 20 * time.Second
    PingDeadline = 10 * time.Second
)
```

<a name="OpSubscribe"></a>

```go
const (
    OpSubscribe   = "subscribe"
    OpUnsubscribe = "unsubscribe"
    OpLogin       = "login"

    EventSubscribe   = OpSubscribe
    EventUnsubscribe = OpUnsubscribe
    EventError       = "error"
)
```

## Variables

<a name="ErrNoData"></a>

```go
var (
    ErrNoData   = errors.New("no data")
    PingMessage = []byte("ping")
)
```

<a name="AsProxyError"></a>
## func AsProxyError

```go
func AsProxyError(e error) (error, bool)
```

AsProxyError convert APIError to ProxyError

<a name="ConvertProxyError"></a>
## func ConvertProxyError

```go
func ConvertProxyError(e error) error
```

ConvertProxyError convert error to ProxyError

<a name="IsAPIError"></a>
## func IsAPIError

```go
func IsAPIError(e error) bool
```

IsAPIError 检查 error 是否为 APIError 类型

<a name="ToInstId"></a>
## func ToInstId

```go
func ToInstId(symbol *common.Symbol) string
```



<a name="ToSymbol"></a>
## func ToSymbol

```go
func ToSymbol(instId string) *common.Symbol
```



<a name="APIError"></a>
## type APIError

APIError 定义 OKX API 错误，当响应状态码为 4xx 或 5xx 时返回

```go
type APIError struct {
    Code     int64  `json:"code,string"`
    Message  string `json:"msg"`
    Response []byte `json:"-"`
}
```

<a name="APIError.Error"></a>
### func \(APIError\) Error

```go
func (e APIError) Error() string
```



<a name="APIError.IsValid"></a>
### func \(APIError\) IsValid

```go
func (e APIError) IsValid() bool
```



<a name="AccountBalance"></a>
## type AccountBalance



```go
type AccountBalance struct {
    TotalEq            Float64                `json:"totalEq"`            // 美金层面权益
    AdjEq              Float64                `json:"adjEq"`              // 美金层面调整后权益（可用余额）
    Imr                Float64                `json:"imr"`                // 美金层面全仓占用保证金
    Mmr                Float64                `json:"mmr"`                // 美金层面维持保证金
    MgnRatio           Float64                `json:"mgnRatio"`           // 美金层面保证金率
    NotionalUsd        Float64                `json:"notionalUsd"`        // 美金层面持仓名义价值
    NotionalUsdForSwap Float64                `json:"notionalUsdForSwap"` // 美金层面持仓名义价值（永续）
    Upl                Float64                `json:"upl"`                // 账户的未实现盈亏
    UTime              Int64                  `json:"uTime"`              // 账户信息的更新时间
    Details            []AccountBalanceDetail `json:"details"`            // 币种维度账户信息
}
```

<a name="AccountBalanceDetail"></a>
## type AccountBalanceDetail



```go
type AccountBalanceDetail struct {
    Ccy       string `json:"ccy"`       // 币种
    Eq        string `json:"eq"`        // 币种总权益
    EqUsd     string `json:"eqUsd"`     // 币种权益美金价值
    CashBal   string `json:"cashBal"`   // 币种余额
    UTime     string `json:"uTime"`     // 更新时间
    AvailBal  string `json:"availBal"`  // 可用余额
    FrozenBal string `json:"frozenBal"` // 冻结余额
    OrdFrozen string `json:"ordFrozen"` // 挂单冻结
    Upl       string `json:"upl"`       // 未实现盈亏
    Imr       string `json:"imr"`       // 币种维度全仓占用保证金
    Mmr       string `json:"mmr"`       // 币种维度全仓维持保证金
}
```

<a name="AccountPosition"></a>
## type AccountPosition



```go
type AccountPosition struct {
    InstId   string  `json:"instId"`   // 产品ID
    InstType string  `json:"instType"` // 产品类型
    Pos      Float64 `json:"pos"`      // 持仓数量
    AvgPx    Float64 `json:"avgPx"`    // 开仓均价
    Lever    Int64   `json:"lever"`    // 杠杆倍数
    Upl      Float64 `json:"upl"`      // 未实现盈亏
    LiqPx    Float64 `json:"liqPx"`    // 强平价格
    LastPx   Float64 `json:"last"`     // 最新成交价格
    IdxPx    Float64 `json:"idxPx"`    // 指数价格
    MarkPx   Float64 `json:"markPx"`   // 最新标记价格
    BePx     Float64 `json:"bePx"`     // 盈亏平衡价
    Imr      Float64 `json:"imr"`      // 初始保证金
    Mmr      Float64 `json:"mmr"`      // 维持保证金
    MgnRatio Float64 `json:"mgnRatio"` // 保证金率
    PosId    string  `json:"posId"`    // 持仓ID
    PosSide  string  `json:"posSide"`  // 持仓方向，long：开平仓模式开多，pos为正；short：开平仓模式开空，pos为正；net：买卖模式（交割/永续/期权：pos为正代表开多，pos为负代表开空。币币杠杆时，pos均为正，posCcy为交易货币时，代表开多；posCcy为计价货币时，代表开空。）
    Ctime    Int64   `json:"ctime"`    // 持仓创建时间（Unix时间戳毫秒）
    UTime    Int64   `json:"uTime"`    // 更新时间（Unix时间戳毫秒）
}
```

<a name="AdapterWebsocket"></a>
## type AdapterWebsocket



```go
type AdapterWebsocket struct {
    // contains filtered or unexported fields
}
```

<a name="NewAdapterWebsocket"></a>
### func NewAdapterWebsocket

```go
func NewAdapterWebsocket(wsUrl string, auth *common.Auth, logger *slog.Logger, httpClient *http.Client) *AdapterWebsocket
```



<a name="AdapterWebsocket.Login"></a>
### func \(\*AdapterWebsocket\) Login

```go
func (a *AdapterWebsocket) Login() error
```



<a name="AdapterWebsocket.Operate"></a>
### func \(\*AdapterWebsocket\) Operate

```go
func (a *AdapterWebsocket) Operate(op string, args ...any) error
```



<a name="AdapterWebsocket.RegisterHandler"></a>
### func \(\*AdapterWebsocket\) RegisterHandler

```go
func (a *AdapterWebsocket) RegisterHandler(channel string, handler func(*wsEvent)) *AdapterWebsocket
```



<a name="AdapterWebsocket.RegisterHandlers"></a>
### func \(\*AdapterWebsocket\) RegisterHandlers

```go
func (a *AdapterWebsocket) RegisterHandlers(handlers map[string]func(*wsEvent)) *AdapterWebsocket
```



<a name="AdapterWebsocket.Send"></a>
### func \(\*AdapterWebsocket\) Send

```go
func (a *AdapterWebsocket) Send(op string, args ...any) error
```



<a name="AdapterWebsocket.SetLogger"></a>
### func \(\*AdapterWebsocket\) SetLogger

```go
func (a *AdapterWebsocket) SetLogger(logIn func(msgType int, msg []byte), logOut func(msgType int, msg []byte)) *AdapterWebsocket
```



<a name="AdapterWebsocket.Start"></a>
### func \(\*AdapterWebsocket\) Start

```go
func (a *AdapterWebsocket) Start() error
```



<a name="AdapterWebsocket.Stop"></a>
### func \(\*AdapterWebsocket\) Stop

```go
func (a *AdapterWebsocket) Stop()
```



<a name="AdapterWebsocket.Subscribe"></a>
### func \(\*AdapterWebsocket\) Subscribe

```go
func (a *AdapterWebsocket) Subscribe(args ...any) error
```



<a name="ArgsLogin"></a>
## type ArgsLogin



```go
type ArgsLogin struct {
    ApiKey     string `json:"apiKey"`
    Passphrase string `json:"passphrase"`
    Timestamp  string `json:"timestamp"`
    Sign       string `json:"sign"`
}
```

<a name="NewArgsLoginFromAuth"></a>
### func NewArgsLoginFromAuth

```go
func NewArgsLoginFromAuth(auth *common.Auth) *ArgsLogin
```



<a name="AssetBalance"></a>
## type AssetBalance



```go
type AssetBalance struct {
    Ccy       string `json:"ccy"`       // 币种
    Bal       string `json:"bal"`       // 余额
    FrozenBal string `json:"frozenBal"` // 冻结余额
    AvailBal  string `json:"availBal"`  // 可用余额
}
```

<a name="AssetDepositHistory"></a>
## type AssetDepositHistory



```go
type AssetDepositHistory struct {
    TxId string `json:"txId"` // 区块转账哈希记录
    Ccy  string `json:"ccy"`  // 币种
    Amt  string `json:"amt"`  // 数量
    Ts   string `json:"ts"`   // 充值时间
}
```

<a name="AssetWithdrawHistory"></a>
## type AssetWithdrawHistory



```go
type AssetWithdrawHistory struct {
    TxId string `json:"txId"` // 提币哈希记录
    Ccy  string `json:"ccy"`  // 币种
    Amt  string `json:"amt"`  // 数量
    Ts   string `json:"ts"`   // 提币时间
}
```

<a name="CancelOrderService"></a>
## type CancelOrderService

CancelOrderService 撤销订单 POST /api/v5/trade/cancel\-order

```go
type CancelOrderService struct {
    // contains filtered or unexported fields
}
```

<a name="CancelOrderService.Do"></a>
### func \(\*CancelOrderService\) Do

```go
func (s *CancelOrderService) Do(ctx context.Context, opts ...RequestOption) (*OrderResult, error)
```



<a name="CancelOrderService.OrdId"></a>
### func \(\*CancelOrderService\) OrdId

```go
func (s *CancelOrderService) OrdId(ordId string) *CancelOrderService
```



<a name="Client"></a>
## type Client



```go
type Client struct {
    APIKey     string
    SecretKey  string
    Passphrase string
    BaseURL    string
    HTTPClient *http.Client
    Debug      bool
    Test       bool
    Logger     *slog.Logger
    TimeOffset int64
    // contains filtered or unexported fields
}
```

<a name="NewClient"></a>
### func NewClient

```go
func NewClient(apiKey, secretKey, passphrase string) *Client
```



<a name="NewClientWithHttpClient"></a>
### func NewClientWithHttpClient

```go
func NewClientWithHttpClient(apiKey, secretKey, passphrase string, test bool, httpClient *http.Client) *Client
```



<a name="NewTestClient"></a>
### func NewTestClient

```go
func NewTestClient(apiKey, secretKey, passphrase string) *Client
```



<a name="Client.NewCancelOrderService"></a>
### func \(\*Client\) NewCancelOrderService

```go
func (c *Client) NewCancelOrderService(instId string) *CancelOrderService
```



<a name="Client.NewGetAccountBalanceService"></a>
### func \(\*Client\) NewGetAccountBalanceService

```go
func (c *Client) NewGetAccountBalanceService() *GetAccountBalanceService
```



<a name="Client.NewGetAccountPositionsService"></a>
### func \(\*Client\) NewGetAccountPositionsService

```go
func (c *Client) NewGetAccountPositionsService() *GetAccountPositionsService
```



<a name="Client.NewGetAssetBalancesService"></a>
### func \(\*Client\) NewGetAssetBalancesService

```go
func (c *Client) NewGetAssetBalancesService() *GetAssetBalancesService
```



<a name="Client.NewGetAssetDepositHistoryService"></a>
### func \(\*Client\) NewGetAssetDepositHistoryService

```go
func (c *Client) NewGetAssetDepositHistoryService() *GetAssetDepositHistoryService
```



<a name="Client.NewGetAssetWithdrawHistoryService"></a>
### func \(\*Client\) NewGetAssetWithdrawHistoryService

```go
func (c *Client) NewGetAssetWithdrawHistoryService() *GetAssetWithdrawHistoryService
```



<a name="Client.NewGetFundingRateService"></a>
### func \(\*Client\) NewGetFundingRateService

```go
func (c *Client) NewGetFundingRateService(instId string) *GetFundingRateService
```



<a name="Client.NewGetInstrumentsService"></a>
### func \(\*Client\) NewGetInstrumentsService

```go
func (c *Client) NewGetInstrumentsService(instType string) *GetInstrumentsService
```



<a name="Client.NewGetMarkPriceService"></a>
### func \(\*Client\) NewGetMarkPriceService

```go
func (c *Client) NewGetMarkPriceService(instType string) *GetMarkPriceService
```



<a name="Client.NewGetMarketIndexTickersService"></a>
### func \(\*Client\) NewGetMarketIndexTickersService

```go
func (c *Client) NewGetMarketIndexTickersService() *GetMarketIndexTickersService
```



<a name="Client.NewGetMarketTickersService"></a>
### func \(\*Client\) NewGetMarketTickersService

```go
func (c *Client) NewGetMarketTickersService(instType string) *GetMarketTickersService
```



<a name="Client.NewGetOrderHistoryService"></a>
### func \(\*Client\) NewGetOrderHistoryService

```go
func (c *Client) NewGetOrderHistoryService() *GetOrderHistoryService
```



<a name="Client.NewGetOrderService"></a>
### func \(\*Client\) NewGetOrderService

```go
func (c *Client) NewGetOrderService(instId string) *GetOrderService
```



<a name="Client.NewGetOrdersPendingService"></a>
### func \(\*Client\) NewGetOrdersPendingService

```go
func (c *Client) NewGetOrdersPendingService() *GetOrdersPendingService
```



<a name="Client.NewGetPriceLimitService"></a>
### func \(\*Client\) NewGetPriceLimitService

```go
func (c *Client) NewGetPriceLimitService(instId string) *GetPriceLimitService
```



<a name="Client.NewGetTradeFillService"></a>
### func \(\*Client\) NewGetTradeFillService

```go
func (c *Client) NewGetTradeFillService() *GetTradeFillService
```



<a name="Client.NewPingService"></a>
### func \(\*Client\) NewPingService

```go
func (c *Client) NewPingService() *PingService
```



<a name="Client.NewSetAccountLeverageService"></a>
### func \(\*Client\) NewSetAccountLeverageService

```go
func (c *Client) NewSetAccountLeverageService(lever string, mgnMode string) *SetAccountLeverageService
```



<a name="Client.NewTradeOrderService"></a>
### func \(\*Client\) NewTradeOrderService

```go
func (c *Client) NewTradeOrderService(instId string, tdMode string, side string, ordType string, sz string) *TradeOrderService
```



<a name="Client.SetProxyURL"></a>
### func \(\*Client\) SetProxyURL

```go
func (c *Client) SetProxyURL(proxyURL *url.URL)
```



<a name="Client.WithHttpClient"></a>
### func \(\*Client\) WithHttpClient

```go
func (c *Client) WithHttpClient(httpClient *http.Client) *Client
```



<a name="FundingRate"></a>
## type FundingRate



```go
type FundingRate struct {
    InstType        string `json:"instType"`        // 产品类型 SWAP：永续合约
    InstId          string `json:"instId"`          // 产品ID，如BTC-USD-SWAP
    Method          string `json:"method"`          // 资金费收取逻辑
    FormulaType     string `json:"formulaType"`     // 公式类型
    FundingRate     string `json:"fundingRate"`     // 资金费率
    NextFundingRate string `json:"nextFundingRate"` // 下一期预测资金费率
    FundingTime     string `json:"fundingTime"`     // 资金费时间
    NextFundingTime string `json:"nextFundingTime"` // 下一期资金费时间
    MinFundingRate  string `json:"minFundingRate"`  // 下一期的预测资金费率下限
    MaxFundingRate  string `json:"maxFundingRate"`  // 下一期的预测资金费率上限
    InterestRate    string `json:"interestRate"`    // 利率
    ImpactValue     string `json:"impactValue"`     // 深度加权金额（计价币数量）
    SettState       string `json:"settState"`       // 资金费率结算状态
    SettFundingRate string `json:"settFundingRate"` // 结算资金费率
    Premium         string `json:"premium"`         // 溢价，为合约的中间价和指数价格的差异
    Ts              string `json:"ts"`              // 数据更新时间
}
```

<a name="GetAccountBalanceService"></a>
## type GetAccountBalanceService

GetAccountBalanceService 用于查询欧易账户信息 https://www.okx.com/docs-v5/zh/#trading-account-rest-api-get-balance

```go
type GetAccountBalanceService struct {
    // contains filtered or unexported fields
}
```

<a name="GetAccountBalanceService.Do"></a>
### func \(\*GetAccountBalanceService\) Do

```go
func (s *GetAccountBalanceService) Do(ctx context.Context, opts ...RequestOption) (*AccountBalance, error)
```



<a name="GetAccountPositionsService"></a>
## type GetAccountPositionsService

GetAccountPositionsService 查看持仓信息 https://www.okx.com/docs-v5/zh/#trading-account-rest-api-get-positions

```go
type GetAccountPositionsService struct {
    // contains filtered or unexported fields
}
```

<a name="GetAccountPositionsService.Do"></a>
### func \(\*GetAccountPositionsService\) Do

```go
func (s *GetAccountPositionsService) Do(ctx context.Context, opts ...RequestOption) ([]AccountPosition, error)
```



<a name="GetAssetBalancesService"></a>
## type GetAssetBalancesService

GetAssetBalancesService 获取资金账户余额 https://www.okx.com/docs-v5/zh/#funding-account-rest-api-get-balance

```go
type GetAssetBalancesService struct {
    // contains filtered or unexported fields
}
```

<a name="GetAssetBalancesService.Do"></a>
### func \(\*GetAssetBalancesService\) Do

```go
func (s *GetAssetBalancesService) Do(ctx context.Context, opts ...RequestOption) ([]AssetBalance, error)
```



<a name="GetAssetDepositHistoryService"></a>
## type GetAssetDepositHistoryService

GetAssetDepositHistoryService 获取资金充值历史 https://www.okx.com/docs-v5/zh/#funding-account-rest-api-get-deposit-history

```go
type GetAssetDepositHistoryService struct {
    // contains filtered or unexported fields
}
```

<a name="GetAssetDepositHistoryService.After"></a>
### func \(\*GetAssetDepositHistoryService\) After

```go
func (s *GetAssetDepositHistoryService) After(after string) *GetAssetDepositHistoryService
```



<a name="GetAssetDepositHistoryService.Before"></a>
### func \(\*GetAssetDepositHistoryService\) Before

```go
func (s *GetAssetDepositHistoryService) Before(before string) *GetAssetDepositHistoryService
```



<a name="GetAssetDepositHistoryService.Do"></a>
### func \(\*GetAssetDepositHistoryService\) Do

```go
func (s *GetAssetDepositHistoryService) Do(ctx context.Context, opts ...RequestOption) ([]AssetDepositHistory, error)
```



<a name="GetAssetDepositHistoryService.EndTime"></a>
### func \(\*GetAssetDepositHistoryService\) EndTime

```go
func (s *GetAssetDepositHistoryService) EndTime(endTime int64) *GetAssetDepositHistoryService
```



<a name="GetAssetDepositHistoryService.Limit"></a>
### func \(\*GetAssetDepositHistoryService\) Limit

```go
func (s *GetAssetDepositHistoryService) Limit(limit int) *GetAssetDepositHistoryService
```



<a name="GetAssetDepositHistoryService.StartTime"></a>
### func \(\*GetAssetDepositHistoryService\) StartTime

```go
func (s *GetAssetDepositHistoryService) StartTime(startTime int64) *GetAssetDepositHistoryService
```



<a name="GetAssetDepositHistoryService.State"></a>
### func \(\*GetAssetDepositHistoryService\) State

```go
func (s *GetAssetDepositHistoryService) State(state string) *GetAssetDepositHistoryService
```



<a name="GetAssetWithdrawHistoryService"></a>
## type GetAssetWithdrawHistoryService

GetAssetWithdrawHistoryService 获取提币记录 https://www.okx.com/docs-v5/zh/#funding-account-rest-api-get-withdrawal-history

```go
type GetAssetWithdrawHistoryService struct {
    // contains filtered or unexported fields
}
```

<a name="GetAssetWithdrawHistoryService.After"></a>
### func \(\*GetAssetWithdrawHistoryService\) After

```go
func (s *GetAssetWithdrawHistoryService) After(after string) *GetAssetWithdrawHistoryService
```



<a name="GetAssetWithdrawHistoryService.Before"></a>
### func \(\*GetAssetWithdrawHistoryService\) Before

```go
func (s *GetAssetWithdrawHistoryService) Before(before string) *GetAssetWithdrawHistoryService
```



<a name="GetAssetWithdrawHistoryService.Do"></a>
### func \(\*GetAssetWithdrawHistoryService\) Do

```go
func (s *GetAssetWithdrawHistoryService) Do(ctx context.Context, opts ...RequestOption) ([]AssetWithdrawHistory, error)
```



<a name="GetAssetWithdrawHistoryService.EndTime"></a>
### func \(\*GetAssetWithdrawHistoryService\) EndTime

```go
func (s *GetAssetWithdrawHistoryService) EndTime(endTime int64) *GetAssetWithdrawHistoryService
```



<a name="GetAssetWithdrawHistoryService.Limit"></a>
### func \(\*GetAssetWithdrawHistoryService\) Limit

```go
func (s *GetAssetWithdrawHistoryService) Limit(limit int) *GetAssetWithdrawHistoryService
```



<a name="GetAssetWithdrawHistoryService.StartTime"></a>
### func \(\*GetAssetWithdrawHistoryService\) StartTime

```go
func (s *GetAssetWithdrawHistoryService) StartTime(startTime int64) *GetAssetWithdrawHistoryService
```



<a name="GetAssetWithdrawHistoryService.State"></a>
### func \(\*GetAssetWithdrawHistoryService\) State

```go
func (s *GetAssetWithdrawHistoryService) State(state string) *GetAssetWithdrawHistoryService
```



<a name="GetFundingRateService"></a>
## type GetFundingRateService



```go
type GetFundingRateService struct {
    // contains filtered or unexported fields
}
```

<a name="GetFundingRateService.Do"></a>
### func \(\*GetFundingRateService\) Do

```go
func (s *GetFundingRateService) Do(ctx context.Context, opts ...RequestOption) ([]*FundingRate, error)
```



<a name="GetInstrumentsService"></a>
## type GetInstrumentsService



```go
type GetInstrumentsService struct {
    // contains filtered or unexported fields
}
```

<a name="GetInstrumentsService.Do"></a>
### func \(\*GetInstrumentsService\) Do

```go
func (s *GetInstrumentsService) Do(ctx context.Context, opts ...RequestOption) ([]*Instrument, error)
```



<a name="GetInstrumentsService.InstFamily"></a>
### func \(\*GetInstrumentsService\) InstFamily

```go
func (s *GetInstrumentsService) InstFamily(instFamily string) *GetInstrumentsService
```



<a name="GetInstrumentsService.InstId"></a>
### func \(\*GetInstrumentsService\) InstId

```go
func (s *GetInstrumentsService) InstId(instId string) *GetInstrumentsService
```



<a name="GetInstrumentsService.Uly"></a>
### func \(\*GetInstrumentsService\) Uly

```go
func (s *GetInstrumentsService) Uly(uly string) *GetInstrumentsService
```



<a name="GetMarkPriceService"></a>
## type GetMarkPriceService



```go
type GetMarkPriceService struct {
    // contains filtered or unexported fields
}
```

<a name="GetMarkPriceService.Do"></a>
### func \(\*GetMarkPriceService\) Do

```go
func (s *GetMarkPriceService) Do(ctx context.Context, opts ...RequestOption) ([]*MarkPrice, error)
```



<a name="GetMarkPriceService.InstFamily"></a>
### func \(\*GetMarkPriceService\) InstFamily

```go
func (s *GetMarkPriceService) InstFamily(instFamily string) *GetMarkPriceService
```



<a name="GetMarkPriceService.InstId"></a>
### func \(\*GetMarkPriceService\) InstId

```go
func (s *GetMarkPriceService) InstId(instId string) *GetMarkPriceService
```



<a name="GetMarketIndexTickersService"></a>
## type GetMarketIndexTickersService



```go
type GetMarketIndexTickersService struct {
    // contains filtered or unexported fields
}
```

<a name="GetMarketIndexTickersService.Do"></a>
### func \(\*GetMarketIndexTickersService\) Do

```go
func (s *GetMarketIndexTickersService) Do(ctx context.Context, opts ...RequestOption) ([]*MarketIndexTicker, error)
```



<a name="GetMarketIndexTickersService.InstId"></a>
### func \(\*GetMarketIndexTickersService\) InstId

```go
func (s *GetMarketIndexTickersService) InstId(instId string) *GetMarketIndexTickersService
```



<a name="GetMarketIndexTickersService.QuoteCcy"></a>
### func \(\*GetMarketIndexTickersService\) QuoteCcy

```go
func (s *GetMarketIndexTickersService) QuoteCcy(quoteCcy string) *GetMarketIndexTickersService
```



<a name="GetMarketTickersService"></a>
## type GetMarketTickersService



```go
type GetMarketTickersService struct {
    // contains filtered or unexported fields
}
```

<a name="GetMarketTickersService.Do"></a>
### func \(\*GetMarketTickersService\) Do

```go
func (s *GetMarketTickersService) Do(ctx context.Context, opts ...RequestOption) ([]*MarketTicker, error)
```



<a name="GetMarketTickersService.InstFamily"></a>
### func \(\*GetMarketTickersService\) InstFamily

```go
func (s *GetMarketTickersService) InstFamily(instFamily string) *GetMarketTickersService
```



<a name="GetMarketTickersService.Uly"></a>
### func \(\*GetMarketTickersService\) Uly

```go
func (s *GetMarketTickersService) Uly(uly string) *GetMarketTickersService
```



<a name="GetOrderHistoryService"></a>
## type GetOrderHistoryService

GetOrderHistoryService 获取历史订单记录（近三个月） https://www.okx.com/docs-v5/zh/#order-book-trading-trade-get-order-list

```go
type GetOrderHistoryService struct {
    // contains filtered or unexported fields
}
```

<a name="GetOrderHistoryService.After"></a>
### func \(\*GetOrderHistoryService\) After

```go
func (s *GetOrderHistoryService) After(after string) *GetOrderHistoryService
```



<a name="GetOrderHistoryService.Archive"></a>
### func \(\*GetOrderHistoryService\) Archive

```go
func (s *GetOrderHistoryService) Archive(archive bool) *GetOrderHistoryService
```



<a name="GetOrderHistoryService.Before"></a>
### func \(\*GetOrderHistoryService\) Before

```go
func (s *GetOrderHistoryService) Before(before string) *GetOrderHistoryService
```



<a name="GetOrderHistoryService.Begin"></a>
### func \(\*GetOrderHistoryService\) Begin

```go
func (s *GetOrderHistoryService) Begin(begin int64) *GetOrderHistoryService
```



<a name="GetOrderHistoryService.Do"></a>
### func \(\*GetOrderHistoryService\) Do

```go
func (s *GetOrderHistoryService) Do(ctx context.Context, opts ...RequestOption) ([]*Order, error)
```



<a name="GetOrderHistoryService.End"></a>
### func \(\*GetOrderHistoryService\) End

```go
func (s *GetOrderHistoryService) End(end int64) *GetOrderHistoryService
```



<a name="GetOrderHistoryService.InstId"></a>
### func \(\*GetOrderHistoryService\) InstId

```go
func (s *GetOrderHistoryService) InstId(instId string) *GetOrderHistoryService
```



<a name="GetOrderHistoryService.InstType"></a>
### func \(\*GetOrderHistoryService\) InstType

```go
func (s *GetOrderHistoryService) InstType(instType string) *GetOrderHistoryService
```



<a name="GetOrderHistoryService.Limit"></a>
### func \(\*GetOrderHistoryService\) Limit

```go
func (s *GetOrderHistoryService) Limit(limit int) *GetOrderHistoryService
```



<a name="GetOrderHistoryService.State"></a>
### func \(\*GetOrderHistoryService\) State

```go
func (s *GetOrderHistoryService) State(state string) *GetOrderHistoryService
```



<a name="GetOrderService"></a>
## type GetOrderService



```go
type GetOrderService struct {
    // contains filtered or unexported fields
}
```

<a name="GetOrderService.ClOrdId"></a>
### func \(\*GetOrderService\) ClOrdId

```go
func (s *GetOrderService) ClOrdId(clOrdId string) *GetOrderService
```



<a name="GetOrderService.Do"></a>
### func \(\*GetOrderService\) Do

```go
func (s *GetOrderService) Do(ctx context.Context, opts ...RequestOption) (*Order, error)
```



<a name="GetOrderService.OrdId"></a>
### func \(\*GetOrderService\) OrdId

```go
func (s *GetOrderService) OrdId(ordId string) *GetOrderService
```



<a name="GetOrdersPendingService"></a>
## type GetOrdersPendingService



```go
type GetOrdersPendingService struct {
    // contains filtered or unexported fields
}
```

<a name="GetOrdersPendingService.After"></a>
### func \(\*GetOrdersPendingService\) After

```go
func (s *GetOrdersPendingService) After(after string) *GetOrdersPendingService
```



<a name="GetOrdersPendingService.Before"></a>
### func \(\*GetOrdersPendingService\) Before

```go
func (s *GetOrdersPendingService) Before(before string) *GetOrdersPendingService
```



<a name="GetOrdersPendingService.Do"></a>
### func \(\*GetOrdersPendingService\) Do

```go
func (s *GetOrdersPendingService) Do(ctx context.Context, opts ...RequestOption) ([]*PendingOrder, error)
```



<a name="GetOrdersPendingService.InstId"></a>
### func \(\*GetOrdersPendingService\) InstId

```go
func (s *GetOrdersPendingService) InstId(instId string) *GetOrdersPendingService
```



<a name="GetOrdersPendingService.InstType"></a>
### func \(\*GetOrdersPendingService\) InstType

```go
func (s *GetOrdersPendingService) InstType(instType string) *GetOrdersPendingService
```



<a name="GetOrdersPendingService.Limit"></a>
### func \(\*GetOrdersPendingService\) Limit

```go
func (s *GetOrdersPendingService) Limit(limit int) *GetOrdersPendingService
```



<a name="GetOrdersPendingService.OrdType"></a>
### func \(\*GetOrdersPendingService\) OrdType

```go
func (s *GetOrdersPendingService) OrdType(ordType string) *GetOrdersPendingService
```



<a name="GetOrdersPendingService.State"></a>
### func \(\*GetOrdersPendingService\) State

```go
func (s *GetOrdersPendingService) State(state string) *GetOrdersPendingService
```



<a name="GetPriceLimitService"></a>
## type GetPriceLimitService



```go
type GetPriceLimitService struct {
    // contains filtered or unexported fields
}
```

<a name="GetPriceLimitService.Do"></a>
### func \(\*GetPriceLimitService\) Do

```go
func (s *GetPriceLimitService) Do(ctx context.Context, opts ...RequestOption) ([]*PriceLimit, error)
```



<a name="GetTradeFillService"></a>
## type GetTradeFillService



```go
type GetTradeFillService struct {
    // contains filtered or unexported fields
}
```

<a name="GetTradeFillService.After"></a>
### func \(\*GetTradeFillService\) After

```go
func (s *GetTradeFillService) After(after string) *GetTradeFillService
```



<a name="GetTradeFillService.Before"></a>
### func \(\*GetTradeFillService\) Before

```go
func (s *GetTradeFillService) Before(before string) *GetTradeFillService
```



<a name="GetTradeFillService.Begin"></a>
### func \(\*GetTradeFillService\) Begin

```go
func (s *GetTradeFillService) Begin(begin string) *GetTradeFillService
```



<a name="GetTradeFillService.Do"></a>
### func \(\*GetTradeFillService\) Do

```go
func (s *GetTradeFillService) Do(ctx context.Context, opts ...RequestOption) ([]TradeFill, error)
```



<a name="GetTradeFillService.End"></a>
### func \(\*GetTradeFillService\) End

```go
func (s *GetTradeFillService) End(end string) *GetTradeFillService
```



<a name="GetTradeFillService.InstFamily"></a>
### func \(\*GetTradeFillService\) InstFamily

```go
func (s *GetTradeFillService) InstFamily(instFamily string) *GetTradeFillService
```



<a name="GetTradeFillService.InstId"></a>
### func \(\*GetTradeFillService\) InstId

```go
func (s *GetTradeFillService) InstId(instId string) *GetTradeFillService
```



<a name="GetTradeFillService.InstType"></a>
### func \(\*GetTradeFillService\) InstType

```go
func (s *GetTradeFillService) InstType(instType string) *GetTradeFillService
```



<a name="GetTradeFillService.Limit"></a>
### func \(\*GetTradeFillService\) Limit

```go
func (s *GetTradeFillService) Limit(limit string) *GetTradeFillService
```



<a name="GetTradeFillService.OrdId"></a>
### func \(\*GetTradeFillService\) OrdId

```go
func (s *GetTradeFillService) OrdId(ordId string) *GetTradeFillService
```



<a name="GetTradeFillService.SubType"></a>
### func \(\*GetTradeFillService\) SubType

```go
func (s *GetTradeFillService) SubType(subType string) *GetTradeFillService
```



<a name="H"></a>
## type H



```go
type H = map[string]any
```

<a name="Instrument"></a>
## type Instrument



```go
type Instrument struct {
    Alias            string  `json:"alias"`            // 合约日期别名，this_week：本周,next_week：次周,this_month：本月,next_month：次月,quarter：季度,next_quarter：次季度,third_quarter：第三季度,仅适用于交割,不建议使用，用户应通过 expTime 字段获取合约的交割日期
    AuctionEndTime   string  `json:"auctionEndTime"`   // 集合竞价结束时间，Unix时间戳的毫秒数格式，如 1597026383085,仅适用于通过集合竞价方式上线的币币，其余情况返回""
    BaseCcy          string  `json:"baseCcy"`          // 交易货币币种，如 BTC-USDT 中的 BTC ，仅适用于币币/币币杠杆
    Category         string  `json:"category"`         // 币种类别（已废弃）
    CtMult           string  `json:"ctMult"`           // 合约乘数，仅适用于交割/永续/期权
    CtType           string  `json:"ctType"`           // 合约类型,linear：正向合约,inverse：反向合约,仅适用于交割/永续
    CtVal            Float64 `json:"ctVal"`            // 合约面值，仅适用于交割/永续/期权
    CtValCcy         string  `json:"ctValCcy"`         // 合约面值计价币种，仅适用于交割/永续/期权
    ExpTime          string  `json:"expTime"`          // 产品下线时间,适用于币币/杠杆/交割/永续/期权，对于 交割/期权，为交割/行权日期；亦可以为产品下线时间，有变动就会推送。
    FutureSettlement bool    `json:"futureSettlement"` // 交割合约是否支持每日结算,适用于全仓交割
    InstFamily       string  `json:"instFamily"`       // 交易品种，如 BTC-USD，仅适用于杠杆/交割/永续/期权
    InstId           string  `json:"instId"`           // 产品id， 如 BTC-USDT
    InstType         string  `json:"instType"`         // 产品类型
    Lever            Int64   `json:"lever"`            // 该instId支持的最大杠杆倍数，不适用于币币、期权
    ListTime         string  `json:"listTime"`         // 上线时间，Unix时间戳的毫秒数格式，如 1597026383085
    LotSz            Float64 `json:"lotSz"`            // 下单数量精度，合约的数量单位是张，现货的数量单位是交易货币
    MaxIcebergSz     string  `json:"maxIcebergSz"`     // 冰山委托的单笔最大委托数量，合约的数量单位是张，现货的数量单位是交易货币
    MaxLmtAmt        Float64 `json:"maxLmtAmt"`        // 限价单的单笔最大美元价值
    MaxLmtSz         Float64 `json:"maxLmtSz"`         // 限价单的单笔最大委托数量，合约的数量单位是张，现货的数量单位是交易货币
    MaxMktAmt        Float64 `json:"maxMktAmt"`        // 市价单的单笔最大美元价值，仅适用于币币/币币杠杆
    MaxMktSz         Float64 `json:"maxMktSz"`         // 市价单的单笔最大委托数量，合约的数量单位是张，现货的数量单位是USDT
    MaxStopSz        Float64 `json:"maxStopSz"`        // 止盈止损市价委托的单笔最大委托数量，合约的数量单位是张，现货的数量单位是USDT
    MaxTriggerSz     Float64 `json:"maxTriggerSz"`     // 计划委托委托的单笔最大委托数量，合约的数量单位是张，现货的数量单位是交易货币
    MaxTwapSz        Float64 `json:"maxTwapSz"`        // 时间加权单的单笔最大委托数量，合约的数量单位是张，现货的数量单位是交易货币。单笔最小委托数量为 minSz*2
    MinSz            Float64 `json:"minSz"`            // 最小下单数量，合约的数量单位是张，现货的数量单位是交易货币
    OptType          string  `json:"optType"`          // 期权类型，C或P 仅适用于期权
    QuoteCcy         string  `json:"quoteCcy"`         // 计价货币币种，如 BTC-USDT 中的USDT ，仅适用于币币/币币杠杆
    SettleCcy        string  `json:"settleCcy"`        // 盈亏结算和保证金币种，如 BTC 仅适用于交割/永续/期权
    State            string  `json:"state"`            // 产品状态，live：交易中,suspend：暂停中,preopen：预上线，交割和期权合约轮转生成到开始交易；部分交易产品上线前,test：测试中（测试产品，不可交易）
    RuleType         string  `json:"ruleType"`         // 交易规则类型，normal：普通交易,pre_market：盘前交易
    Stk              string  `json:"stk"`              // 行权价格，仅适用于期权
    TickSz           Float64 `json:"tickSz"`           // 下单价格精度，如 0.0001，对于期权来说，是梯度中的最小下单价格精度，如果想要获取期权价格梯度，请使用"获取期权价格梯度"接口
    Uly              string  `json:"uly"`              // 标的指数，如 BTC-USD，仅适用于杠杆/交割/永续/期权
}
```

<a name="MarkPrice"></a>
## type MarkPrice



```go
type MarkPrice struct {
    InstType string  `json:"instType"` // 产品类型 SPOT：币币 MARGIN：杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
    InstId   string  `json:"instId"`   // 产品ID，如 BTC-USDT-SWAP
    MarkPx   Float64 `json:"markPx"`   // 标记价格
    Ts       Int64   `json:"ts"`       // 标记价格更新时间，Unix时间戳的毫秒数格式，如 1597026383085
}
```

<a name="MarketIndexTicker"></a>
## type MarketIndexTicker



```go
type MarketIndexTicker struct {
    InstId  string `json:"instId"`  // 指数
    IdxPx   string `json:"idxPx"`   // 最新指数价格
    High24h string `json:"high24h"` // 24小时最高价格
    SodUtc0 string `json:"sodUtc0"` // UTC 0 时开盘价
    Open24h string `json:"open24h"` // 24小时开盘价格
    Low24h  string `json:"low24h"`  // 24小时最低价格
    SodUtc8 string `json:"sodUtc8"` // UTC+8 时开盘价
    Ts      string `json:"ts"`      // 数据产生时间，Unix时间戳的毫秒数格式，如 1597026383085
}
```

<a name="MarketTicker"></a>
## type MarketTicker



```go
type MarketTicker struct {
    InstType  string `json:"instType"`  // 产品类型
    InstId    string `json:"instId"`    // 产品ID
    Last      string `json:"last"`      // 最新成交价
    LastSz    string `json:"lastSz"`    // 最新成交的数量，0 代表没有成交量
    AskPx     string `json:"askPx"`     // 卖一价
    AskSz     string `json:"askSz"`     // 卖一价的挂单数数量
    BidPx     string `json:"bidPx"`     // 买一价
    BidSz     string `json:"bidSz"`     // 买一价的挂单数量
    Open24h   string `json:"open24h"`   // 24小时开盘价
    High24h   string `json:"high24h"`   // 24小时最高价
    Low24h    string `json:"low24h"`    // 24小时最低价
    VolCcy24h string `json:"volCcy24h"` // 24小时成交量，以币为单位
    Vol24h    string `json:"vol24h"`    // 24小时成交量，以张为单位
    SodUtc0   string `json:"sodUtc0"`   // UTC 0 时开盘价
    SodUtc8   string `json:"sodUtc8"`   // UTC+8 时开盘价
    Ts        string `json:"ts"`        // ticker数据产生时间，Unix时间戳的毫秒数格式
}
```

<a name="OkxPublicStreamAdapter"></a>
## type OkxPublicStreamAdapter



```go
type OkxPublicStreamAdapter struct {
    *AdapterWebsocket
    // contains filtered or unexported fields
}
```

<a name="NewOkxPublicStreamAdapter"></a>
### func NewOkxPublicStreamAdapter

```go
func NewOkxPublicStreamAdapter(proxyUrl *url.URL) *OkxPublicStreamAdapter
```



<a name="OkxPublicStreamAdapter.Start"></a>
### func \(\*OkxPublicStreamAdapter\) Start

```go
func (s *OkxPublicStreamAdapter) Start() error
```



<a name="OkxPublicStreamAdapter.Stop"></a>
### func \(\*OkxPublicStreamAdapter\) Stop

```go
func (s *OkxPublicStreamAdapter) Stop()
```



<a name="OperatePublicArg"></a>
## type OperatePublicArg



```go
type OperatePublicArg struct {
    Channel string `json:"channel"`
    InstId  string `json:"instId"`
}
```

<a name="OperateRequest"></a>
## type OperateRequest



```go
type OperateRequest struct {
    Op   string `json:"op"`
    Args []any  `json:"args"`
}
```

<a name="Order"></a>
## type Order



```go
type Order struct {
    InstType  string `json:"instType"`  // 产品类型
    InstId    string `json:"instId"`    // 产品ID
    OrdId     string `json:"ordId"`     // 订单ID
    OrdType   string `json:"ordType"`   // 订单类型
    Side      string `json:"side"`      // 交易方向
    PosSide   string `json:"posSide"`   // 持仓方向
    Sz        string `json:"sz"`        // 委托数量
    AccFillSz string `json:"accFillSz"` // 累计成交数量
    AvgPx     string `json:"avgPx"`     // 成交均价
    State     string `json:"state"`     // 订单状态
    Lever     string `json:"lever"`     // 杠杆倍数
    Pnl       string `json:"pnl"`       // 盈亏
    FeeCcy    string `json:"feeCcy"`    // 手续费币种
    Fee       string `json:"fee"`       // 手续费
    CTime     string `json:"cTime"`     // 创建时间
    FillTime  string `json:"fillTime"`  // 最新成交时间
}
```

<a name="OrderResult"></a>
## type OrderResult



```go
type OrderResult struct {
    OrdId   string `json:"ordId"`   // 订单ID
    ClOrdId string `json:"clOrdId"` // 客户自定义订单ID
    Tag     string `json:"tag"`     // 订单标签
    Ts      Int64  `json:"ts"`      // 订单创建时间
    SCode   string `json:"sCode"`   // 状态码
    SMsg    string `json:"sMsg"`    // 状态信息
}
```

<a name="PendingOrder"></a>
## type PendingOrder



```go
type PendingOrder struct {
    InstType string `json:"instType"` // 产品类型
    InstId   string `json:"instId"`   // 产品ID
    OrdId    string `json:"ordId"`    // 订单ID
    OrdType  string `json:"ordType"`  // 订单类型
    Side     string `json:"side"`     // 交易方向
    PosSide  string `json:"posSide"`  // 持仓方向
    Px       string `json:"px"`       // 委托价格
    Sz       string `json:"sz"`       // 委托数量
    State    string `json:"state"`    // 订单状态
    CTime    string `json:"cTime"`    // 创建时间
    UTime    string `json:"uTime"`    // 更新时间
}
```

<a name="PingService"></a>
## type PingService

PingService 用于检测欧易服务器连通性

```go
type PingService struct {
    // contains filtered or unexported fields
}
```

<a name="PingService.Do"></a>
### func \(\*PingService\) Do

```go
func (s *PingService) Do(ctx context.Context) error
```



<a name="PriceLimit"></a>
## type PriceLimit



```go
type PriceLimit struct {
    InstType string `json:"instType"` // 产品类型 SPOT：币币 MARGIN：杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
    InstId   string `json:"instId"`   // 产品ID，如 BTC-USDT-SWAP
    BuyLmt   string `json:"buyLmt"`   // 最高买价，当enabled为false时，返回""
    SellLmt  string `json:"sellLmt"`  // 最低卖价，当enabled为false时，返回""
    Ts       string `json:"ts"`       // 限价数据更新时间，Unix时间戳的毫秒数格式，如 1597026383085
    Enabled  bool   `json:"enabled"`  // 限价是否生效 true：限价生效 false：限价不生效
}
```

<a name="RequestOption"></a>
## type RequestOption



```go
type RequestOption func(*request)
```

<a name="WithHeader"></a>
### func WithHeader

```go
func WithHeader(key, value string, replace bool) RequestOption
```



<a name="WithHeaders"></a>
### func WithHeaders

```go
func WithHeaders(header http.Header) RequestOption
```



<a name="Response"></a>
## type Response



```go
type Response struct {
    Event string `json:"event"`
    Arg   any    `json:"arg"`
    Code  string `json:"code"`
    Msg   string `json:"msg"`
}
```

<a name="Response.Error"></a>
### func \(Response\) Error

```go
func (resp Response) Error() error
```



<a name="Response.IsError"></a>
### func \(Response\) IsError

```go
func (resp Response) IsError() bool
```



<a name="SetAccountLeverageService"></a>
## type SetAccountLeverageService

SetAccountLeverageService 设置杠杆倍数 https://www.okx.com/docs-v5/zh/#trading-account-rest-api-set-leverage

```go
type SetAccountLeverageService struct {
    // contains filtered or unexported fields
}
```

<a name="SetAccountLeverageService.Ccy"></a>
### func \(\*SetAccountLeverageService\) Ccy

```go
func (s *SetAccountLeverageService) Ccy(ccy string) *SetAccountLeverageService
```



<a name="SetAccountLeverageService.Do"></a>
### func \(\*SetAccountLeverageService\) Do

```go
func (s *SetAccountLeverageService) Do(ctx context.Context, opts ...RequestOption) error
```



<a name="SetAccountLeverageService.InstId"></a>
### func \(\*SetAccountLeverageService\) InstId

```go
func (s *SetAccountLeverageService) InstId(instId string) *SetAccountLeverageService
```



<a name="SetAccountLeverageService.PosSide"></a>
### func \(\*SetAccountLeverageService\) PosSide

```go
func (s *SetAccountLeverageService) PosSide(posSide string) *SetAccountLeverageService
```



<a name="TradeFill"></a>
## type TradeFill



```go
type TradeFill struct {
    Side          string  `json:"side"`
    FillSz        Float64 `json:"fillSz"`
    FillPx        Float64 `json:"fillPx"`
    FillPxVol     string  `json:"fillPxVol"`
    FillFwdPx     string  `json:"fillFwdPx"`
    Fee           Float64 `json:"fee"`
    FillPnl       Float64 `json:"fillPnl"`
    OrdId         string  `json:"ordId"`
    FeeRate       string  `json:"feeRate"`
    InstType      string  `json:"instType"`
    FillPxUsd     string  `json:"fillPxUsd"`
    InstId        string  `json:"instId"`
    ClOrdId       string  `json:"clOrdId"`
    PosSide       string  `json:"posSide"`
    BillId        string  `json:"billId"`
    SubType       string  `json:"subType"`
    FillMarkVol   string  `json:"fillMarkVol"`
    Tag           string  `json:"tag"`
    FillTime      Int64   `json:"fillTime"`
    ExecType      string  `json:"execType"`
    FillIdxPx     string  `json:"fillIdxPx"`
    TradeId       string  `json:"tradeId"`
    FillMarkPx    string  `json:"fillMarkPx"`
    FeeCcy        string  `json:"feeCcy"`
    Ts            Int64   `json:"ts"`
    TradeQuoteCcy string  `json:"tradeQuoteCcy"`
}
```

<a name="TradeOrderService"></a>
## type TradeOrderService



```go
type TradeOrderService struct {
    // contains filtered or unexported fields
}
```

<a name="TradeOrderService.ClOrdId"></a>
### func \(\*TradeOrderService\) ClOrdId

```go
func (s *TradeOrderService) ClOrdId(clOrdId string) *TradeOrderService
```



<a name="TradeOrderService.Do"></a>
### func \(\*TradeOrderService\) Do

```go
func (s *TradeOrderService) Do(ctx context.Context, opts ...RequestOption) (*OrderResult, error)
```



<a name="TradeOrderService.PosSide"></a>
### func \(\*TradeOrderService\) PosSide

```go
func (s *TradeOrderService) PosSide(posSide string) *TradeOrderService
```



<a name="TradeOrderService.Px"></a>
### func \(\*TradeOrderService\) Px

```go
func (s *TradeOrderService) Px(px string) *TradeOrderService
```



<a name="TradeOrderService.ReduceOnly"></a>
### func \(\*TradeOrderService\) ReduceOnly

```go
func (s *TradeOrderService) ReduceOnly(reduceOnly bool) *TradeOrderService
```



# types

```go
import "github.com/youjianglong/exchanges/types"
```

## Index

- [Constants](<#constants>)
- [Variables](<#variables>)
- [type OrderSide](<#OrderSide>)
  - [func \(o OrderSide\) Equal\(other string\) bool](<#OrderSide.Equal>)
- [type OrderStatus](<#OrderStatus>)
- [type OrderType](<#OrderType>)
  - [func \(o OrderType\) Equal\(other string\) bool](<#OrderType.Equal>)
- [type PosSide](<#PosSide>)
- [type Zero](<#Zero>)


## Constants

<a name="IncomeTypeRealizedPnl"></a>

```go
const (
    IncomeTypeRealizedPnl = "REALIZED_PNL"
    IncomeTypeCommission  = "COMMISSION"
)
```

## Variables

<a name="ClosedChan"></a>

```go
var ClosedChan = makeClosedChan()
```

<a name="OrderSide"></a>
## type OrderSide

OrderSide 订单方向

```go
type OrderSide string
```

<a name="Buy"></a>

```go
const (
    Buy  OrderSide = "buy"
    Sell OrderSide = "sell"
)
```

<a name="OrderSide.Equal"></a>
### func \(OrderSide\) Equal

```go
func (o OrderSide) Equal(other string) bool
```



<a name="OrderStatus"></a>
## type OrderStatus

OrderStatus 订单状态

```go
type OrderStatus string
```

<a name="Created"></a>

```go
const (
    Created         OrderStatus = "created"          // 已创建
    Filled          OrderStatus = "filled"           // 已成交
    Cancelled       OrderStatus = "cancelled"        // 已取消
    Failed          OrderStatus = "failed"           // 失败
    PartiallyFilled OrderStatus = "partially_filled" // 部分成交
    Expired         OrderStatus = "expired"          // 已过期
    Unknown         OrderStatus = "unknown"          // 未知
)
```

<a name="OrderType"></a>
## type OrderType



```go
type OrderType string
```

<a name="OrderTypeLimit"></a>

```go
const (
    OrderTypeLimit  OrderType = "limit"
    OrderTypeMarket OrderType = "market"
)
```

<a name="OrderType.Equal"></a>
### func \(OrderType\) Equal

```go
func (o OrderType) Equal(other string) bool
```



<a name="PosSide"></a>
## type PosSide



```go
type PosSide string
```

<a name="PosSideLong"></a>

```go
const (
    PosSideLong  PosSide = "LONG"
    PosSideShort PosSide = "SHORT"
)
```

<a name="Zero"></a>
## type Zero



```go
type Zero = struct{}
```

# ws

```go
import "github.com/youjianglong/exchanges/ws"
```

## Index

- [Variables](<#variables>)
- [func GetFileLogger\(dir string, name string, split time.Duration, expire time.Duration\) \(in func\(int, \[\]byte\), out func\(int, \[\]byte\)\)](<#GetFileLogger>)
- [func GetLogger\(w io.Writer\) \(in func\(int, \[\]byte\), out func\(int, \[\]byte\)\)](<#GetLogger>)
- [type Conn](<#Conn>)
- [type ErrHandler](<#ErrHandler>)
- [type MessageType](<#MessageType>)
- [type Websocket](<#Websocket>)
  - [func NewWebsocket\(endpoint string, handler WsHandler, errHandler ErrHandler, keepAlive func\(\*Websocket\)\) \*Websocket](<#NewWebsocket>)
  - [func \(w \*Websocket\) Conn\(\) \*websocket.Conn](<#Websocket.Conn>)
  - [func \(w \*Websocket\) Done\(\) \<\-chan struct\{\}](<#Websocket.Done>)
  - [func \(w \*Websocket\) LogIn\(msgType int, message \[\]byte\)](<#Websocket.LogIn>)
  - [func \(w \*Websocket\) LogOut\(msgType int, message \[\]byte\)](<#Websocket.LogOut>)
  - [func \(w \*Websocket\) Read\(\) \(int, \[\]byte, error\)](<#Websocket.Read>)
  - [func \(w \*Websocket\) ReadJSON\(v any\) error](<#Websocket.ReadJSON>)
  - [func \(w \*Websocket\) Restart\(\)](<#Websocket.Restart>)
  - [func \(w \*Websocket\) SetEndpoint\(endpoint string\) \*Websocket](<#Websocket.SetEndpoint>)
  - [func \(w \*Websocket\) SetErrHandler\(handler ErrHandler\) \*Websocket](<#Websocket.SetErrHandler>)
  - [func \(w \*Websocket\) SetHttpClient\(httpClient \*http.Client\) \*Websocket](<#Websocket.SetHttpClient>)
  - [func \(w \*Websocket\) SetLogger\(logIn func\(int, \[\]byte\), logOut func\(int, \[\]byte\)\) \*Websocket](<#Websocket.SetLogger>)
  - [func \(w \*Websocket\) SetOnPingReceived\(handler func\(context.Context, \[\]byte\) bool\) \*Websocket](<#Websocket.SetOnPingReceived>)
  - [func \(w \*Websocket\) SetOnPongReceived\(handler func\(context.Context, \[\]byte\)\) \*Websocket](<#Websocket.SetOnPongReceived>)
  - [func \(w \*Websocket\) SetPrevConnect\(handler func\(\*Websocket\) error\) \*Websocket](<#Websocket.SetPrevConnect>)
  - [func \(w \*Websocket\) SetStartConnect\(handler func\(\*Websocket, \*websocket.Conn\) error\) \*Websocket](<#Websocket.SetStartConnect>)
  - [func \(w \*Websocket\) SetWsHandler\(handler WsHandler\) \*Websocket](<#Websocket.SetWsHandler>)
  - [func \(w \*Websocket\) Start\(\) error](<#Websocket.Start>)
  - [func \(w \*Websocket\) Stop\(\)](<#Websocket.Stop>)
  - [func \(w \*Websocket\) WaitReady\(\) \<\-chan struct\{\}](<#Websocket.WaitReady>)
  - [func \(w \*Websocket\) Write\(msgType MessageType, message \[\]byte\) error](<#Websocket.Write>)
  - [func \(w \*Websocket\) WriteJSON\(v any\) error](<#Websocket.WriteJSON>)
- [type WsHandler](<#WsHandler>)


## Variables

<a name="MessageText"></a>

```go
var (
    MessageText   = websocket.MessageText
    MessageBinary = websocket.MessageBinary
    MessagePing   = websocket.MessageBinary + 1
    MessagePong   = websocket.MessageBinary + 2
)
```

<a name="EOL"></a>

```go
var (
    EOL = []byte("\n")
)
```

<a name="ErrServiceStopped"></a>ErrServiceStopped 服务停止

```go
var ErrServiceStopped = errors.New("service stopped")
```

<a name="ErrWebsocketNotConnected"></a>

```go
var ErrWebsocketNotConnected = errors.New("websocket connection is nil")
```

<a name="GetFileLogger"></a>
## func GetFileLogger

```go
func GetFileLogger(dir string, name string, split time.Duration, expire time.Duration) (in func(int, []byte), out func(int, []byte))
```



<a name="GetLogger"></a>
## func GetLogger

```go
func GetLogger(w io.Writer) (in func(int, []byte), out func(int, []byte))
```



<a name="Conn"></a>
## type Conn



```go
type Conn = websocket.Conn
```

<a name="ErrHandler"></a>
## type ErrHandler

ErrHandler handles errors

```go
type ErrHandler func(err error)
```

<a name="MessageType"></a>
## type MessageType



```go
type MessageType = websocket.MessageType
```

<a name="Websocket"></a>
## type Websocket



```go
type Websocket struct {
    HandshakeTimeout time.Duration
    AutoReconnect    bool
    // contains filtered or unexported fields
}
```

<a name="NewWebsocket"></a>
### func NewWebsocket

```go
func NewWebsocket(endpoint string, handler WsHandler, errHandler ErrHandler, keepAlive func(*Websocket)) *Websocket
```



<a name="Websocket.Conn"></a>
### func \(\*Websocket\) Conn

```go
func (w *Websocket) Conn() *websocket.Conn
```



<a name="Websocket.Done"></a>
### func \(\*Websocket\) Done

```go
func (w *Websocket) Done() <-chan struct{}
```



<a name="Websocket.LogIn"></a>
### func \(\*Websocket\) LogIn

```go
func (w *Websocket) LogIn(msgType int, message []byte)
```



<a name="Websocket.LogOut"></a>
### func \(\*Websocket\) LogOut

```go
func (w *Websocket) LogOut(msgType int, message []byte)
```



<a name="Websocket.Read"></a>
### func \(\*Websocket\) Read

```go
func (w *Websocket) Read() (int, []byte, error)
```



<a name="Websocket.ReadJSON"></a>
### func \(\*Websocket\) ReadJSON

```go
func (w *Websocket) ReadJSON(v any) error
```



<a name="Websocket.Restart"></a>
### func \(\*Websocket\) Restart

```go
func (w *Websocket) Restart()
```



<a name="Websocket.SetEndpoint"></a>
### func \(\*Websocket\) SetEndpoint

```go
func (w *Websocket) SetEndpoint(endpoint string) *Websocket
```



<a name="Websocket.SetErrHandler"></a>
### func \(\*Websocket\) SetErrHandler

```go
func (w *Websocket) SetErrHandler(handler ErrHandler) *Websocket
```



<a name="Websocket.SetHttpClient"></a>
### func \(\*Websocket\) SetHttpClient

```go
func (w *Websocket) SetHttpClient(httpClient *http.Client) *Websocket
```



<a name="Websocket.SetLogger"></a>
### func \(\*Websocket\) SetLogger

```go
func (w *Websocket) SetLogger(logIn func(int, []byte), logOut func(int, []byte)) *Websocket
```



<a name="Websocket.SetOnPingReceived"></a>
### func \(\*Websocket\) SetOnPingReceived

```go
func (w *Websocket) SetOnPingReceived(handler func(context.Context, []byte) bool) *Websocket
```



<a name="Websocket.SetOnPongReceived"></a>
### func \(\*Websocket\) SetOnPongReceived

```go
func (w *Websocket) SetOnPongReceived(handler func(context.Context, []byte)) *Websocket
```



<a name="Websocket.SetPrevConnect"></a>
### func \(\*Websocket\) SetPrevConnect

```go
func (w *Websocket) SetPrevConnect(handler func(*Websocket) error) *Websocket
```



<a name="Websocket.SetStartConnect"></a>
### func \(\*Websocket\) SetStartConnect

```go
func (w *Websocket) SetStartConnect(handler func(*Websocket, *websocket.Conn) error) *Websocket
```



<a name="Websocket.SetWsHandler"></a>
### func \(\*Websocket\) SetWsHandler

```go
func (w *Websocket) SetWsHandler(handler WsHandler) *Websocket
```



<a name="Websocket.Start"></a>
### func \(\*Websocket\) Start

```go
func (w *Websocket) Start() error
```



<a name="Websocket.Stop"></a>
### func \(\*Websocket\) Stop

```go
func (w *Websocket) Stop()
```



<a name="Websocket.WaitReady"></a>
### func \(\*Websocket\) WaitReady

```go
func (w *Websocket) WaitReady() <-chan struct{}
```



<a name="Websocket.Write"></a>
### func \(\*Websocket\) Write

```go
func (w *Websocket) Write(msgType MessageType, message []byte) error
```



<a name="Websocket.WriteJSON"></a>
### func \(\*Websocket\) WriteJSON

```go
func (w *Websocket) WriteJSON(v any) error
```



<a name="WsHandler"></a>
## type WsHandler

WsHandler handle raw websocket message

```go
type WsHandler func(message []byte)
```

# common

```go
import "github.com/youjianglong/exchanges/okx/common"
```

## Index

- [func HmacSHA256\(data, key \[\]byte\) string](<#HmacSHA256>)
- [type Auth](<#Auth>)
  - [func NewAuth\(objectID, apiKey, secretKey, passphrase string\) Auth](<#NewAuth>)
  - [func \(a Auth\) Signature\(method, path, body string, isUnix bool\) \*Signature](<#Auth.Signature>)
- [type Signature](<#Signature>)
  - [func \(s \*Signature\) Build\(\) string](<#Signature.Build>)


<a name="HmacSHA256"></a>
## func HmacSHA256

```go
func HmacSHA256(data, key []byte) string
```

hmac sha256

<a name="Auth"></a>
## type Auth



```go
type Auth struct {
    ObjectID   string
    ApiKey     string
    SecretKey  string
    Passphrase string
}
```

<a name="NewAuth"></a>
### func NewAuth

```go
func NewAuth(objectID, apiKey, secretKey, passphrase string) Auth
```



<a name="Auth.Signature"></a>
### func \(Auth\) Signature

```go
func (a Auth) Signature(method, path, body string, isUnix bool) *Signature
```



<a name="Signature"></a>
## type Signature



```go
type Signature struct {
    Key, Timestamp, Method, Path, Body string
    IsUnix                             bool
}
```

<a name="Signature.Build"></a>
### func \(\*Signature\) Build

```go
func (s *Signature) Build() string
```

The Base64\-encoded signature \(see Signing Messages subsection for details\).
