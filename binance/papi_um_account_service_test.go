package binance

import (
	"context"
	"testing"
)

func TestPApiGetPositions(t *testing.T) {
	client := newTestClient()
	positions, err := client.NewPApiGetUMPositionsService().Do(context.Background())
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

func TestPapiUMGetPositionRiskService(t *testing.T) {
	client := newTestClient()
	positions, err := client.NewPapiUMGetPositionRiskService().Symbol(*symbol).Do(context.Background())
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

func TestPApiUmGetIncomeService(t *testing.T) {
	client := newTestClient()
	incomes, err := client.NewPApiUmGetIncomeService().Symbol(*symbol).IncomeType(*incomeType).StartTime(*startTime).EndTime(*endTime).Limit(*limit).Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, income := range incomes {
		t.Logf("%+v", income)
	}
}

func TestExportIncome(t *testing.T) {
	client := newTestClient()
	svc := client.NewPApiUmGetIncomeService().IncomeType(*incomeType).
		StartTime(*startTime).
		EndTime(*endTime).
		Limit(*limit)
	if *symbol != "" {
		svc = svc.Symbol(*symbol)
	}
	incomes, err := svc.Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, income := range incomes {
		t.Logf("%+v", income)
	}
}
