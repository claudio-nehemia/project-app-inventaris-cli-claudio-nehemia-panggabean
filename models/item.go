package models

import "time"

type Item struct {
    ID           int       `json:"id"`
    Name         string    `json:"name"`
    CategoryID   int       `json:"category_id"`
    CategoryName string    `json:"category_name"`
    Price        float64   `json:"price"`
    PurchaseDate time.Time `json:"purchase_date"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}

type ItemDepreciation struct {
    Item
    DaysUsed          int     `json:"days_used"`
    DepreciationRate  float64 `json:"depreciation_rate"`
    CurrentValue      float64 `json:"current_value"`
    DepreciationValue float64 `json:"depreciation_value"`
}