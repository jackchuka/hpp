package api

import (
	"testing"

	"github.com/google/go-querystring/query"
)

func TestGourmetSearchParams_Encode(t *testing.T) {
	p := GourmetSearchParams{
		Keyword: strPtr("ramen"),
		WiFi:    true,
		Lunch:   true,
		Count:   intPtr(5),
	}
	vals, err := query.Values(p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if vals.Get("keyword") != "ramen" {
		t.Fatalf("expected keyword=ramen, got %s", vals.Get("keyword"))
	}
	if vals.Get("wifi") != "1" {
		t.Fatalf("expected wifi=1, got %s", vals.Get("wifi"))
	}
	if vals.Get("lunch") != "1" {
		t.Fatalf("expected lunch=1, got %s", vals.Get("lunch"))
	}
	if vals.Get("count") != "5" {
		t.Fatalf("expected count=5, got %s", vals.Get("count"))
	}
	if vals.Get("pet") != "" {
		t.Fatal("expected pet to be omitted")
	}
}

func TestGourmetSearchParams_BoolOmitted(t *testing.T) {
	p := GourmetSearchParams{}
	vals, err := query.Values(p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if vals.Get("wifi") != "" {
		t.Fatal("expected wifi to be omitted when false")
	}
	if vals.Get("keyword") != "" {
		t.Fatal("expected keyword to be omitted when nil")
	}
}

func TestShopSearchParams_Encode(t *testing.T) {
	kw := "sushi"
	p := ShopSearchParams{Keyword: &kw}
	vals, err := query.Values(p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if vals.Get("keyword") != "sushi" {
		t.Fatalf("expected keyword=sushi, got %s", vals.Get("keyword"))
	}
}

func strPtr(s string) *string { return &s }
func intPtr(i int) *int       { return &i }
