package okx

import (
	"context"
	"testing"
)

func TestGetAccountBalance(t *testing.T) {
	client := newTestClient()
	account, err := client.NewGetAccountBalanceService().Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", account)
}

func TestGetAccountPositions(t *testing.T) {
	client := newTestClient()
	account, err := client.NewGetAccountPositionsService().Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", account)
}

func TestGetAssetBalances(t *testing.T) {
	client := newTestClient()
	account, err := client.NewGetAssetBalancesService().Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", account)
}
