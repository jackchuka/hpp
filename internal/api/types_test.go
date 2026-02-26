package api

import (
	"encoding/json"
	"testing"
)

func TestGourmetResponseUnmarshal(t *testing.T) {
	data := `{"results":{"api_version":"1.20","results_available":1,"results_returned":"1","results_start":1,"shop":[{"id":"J001","name":"Test Shop","genre":{"code":"G001","name":"居酒屋"},"budget":{"code":"B001","name":"2001～3000円","average":"2500円"},"middle_area":{"code":"Y005","name":"銀座"},"urls":{"pc":"https://example.com"}}]}}`
	var resp GourmetResponse
	if err := json.Unmarshal([]byte(data), &resp); err != nil {
		t.Fatal(err)
	}
	if len(resp.Results.Shops) != 1 {
		t.Fatalf("expected 1 shop, got %d", len(resp.Results.Shops))
	}
	if resp.Results.Shops[0].Name != "Test Shop" {
		t.Fatalf("expected Test Shop, got %s", resp.Results.Shops[0].Name)
	}
}

func TestGenreResponseUnmarshal(t *testing.T) {
	data := `{"results":{"api_version":"1.20","results_available":2,"results_returned":"2","results_start":1,"genre":[{"code":"G001","name":"居酒屋"},{"code":"G002","name":"ダイニングバー"}]}}`
	var resp GenreResponse
	if err := json.Unmarshal([]byte(data), &resp); err != nil {
		t.Fatal(err)
	}
	if len(resp.Results.Genres) != 2 {
		t.Fatalf("expected 2 genres, got %d", len(resp.Results.Genres))
	}
}

func TestBudgetResponseUnmarshal(t *testing.T) {
	data := `{"results":{"api_version":"1.20","results_available":1,"results_returned":"1","results_start":1,"budget":[{"code":"B001","name":"～2000円"}]}}`
	var resp BudgetResponse
	if err := json.Unmarshal([]byte(data), &resp); err != nil {
		t.Fatal(err)
	}
	if len(resp.Results.Budgets) != 1 {
		t.Fatalf("expected 1 budget, got %d", len(resp.Results.Budgets))
	}
}
