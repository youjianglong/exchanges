package binance

import (
	"context"
	"testing"
)

func TestGetUMAccount(t *testing.T) {
	client := newTestClient()
	account, err := client.NewGetUMAccountService().Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", account)
}

// func TestGetUMAccountDetail(t *testing.T) {
// 	client := newTestClient()
// 	account, err := client.NewGetUMAccountDetailService().Do(context.Background())
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log("assets:\n")
// 	for _, asset := range account.Assets {
// 		if IsZeroStr(asset.CrossWalletBalance) {
// 			continue
// 		}
// 		t.Logf("%+v\n", asset)
// 	}
// 	t.Log("positions:\n")
// 	for _, p := range account.Positions {
// 		if p.PositionAmt == 0 {
// 			continue
// 		}
// 		t.Logf("%+v\n", p)
// 	}
// }

func TestPApiGetAccountBalance(t *testing.T) {
	client := newTestClient()
	balances, err := client.NewGetPApiAccountBalanceService().Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, balance := range balances {
		t.Logf("%+v", balance)
	}
}

func TestFApiGetAccountBalance(t *testing.T) {
	client := newTestClient()
	balances, err := client.NewGetFApiAccountBalanceService().Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, balance := range balances {
		t.Logf("%+v", balance)
	}
}

func TestGetPApiPositions(t *testing.T) {
	client := newTestClient()
	positions, err := client.NewGetPApiPositionsService().Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, position := range positions {
		if position.PositionAmt.Value() == 0 {
			continue
		}
		t.Logf("%+v", position)
	}
}

func TestGetFApiPositions(t *testing.T) {
	client := newTestClient()
	positions, err := client.NewGetFApiPositionsService().Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, position := range positions {
		if position.PositionAmt.Value() == 0 {
			continue
		}
		t.Logf("%+v", position)
	}
}
