package service

import (
    "errors"
    "testing"
    "time"

    "mini_project3/models"
)

// Mock Item Repository
type MockItemRepository struct {
    items       []models.Item
    shouldError bool
}

func (m *MockItemRepository) GetAll() ([]models.Item, error) {
    if m.shouldError {
        return nil, errors.New("mock error")
    }
    return m.items, nil
}

func (m *MockItemRepository) GetByID(id int) (*models.Item, error) {
    if m.shouldError {
        return nil, errors.New("mock error")
    }
    for _, item := range m.items {
        if item.ID == id {
            return &item, nil
        }
    }
    return nil, errors.New("item not found")
}

func (m *MockItemRepository) Create(item *models.Item) error {
    if m.shouldError {
        return errors.New("mock error")
    }
    item.ID = len(m.items) + 1
    item.CreatedAt = time.Now()
    m.items = append(m.items, *item)
    return nil
}

func (m *MockItemRepository) Update(item *models.Item) error {
    if m.shouldError {
        return errors.New("mock error")
    }
    return nil
}

func (m *MockItemRepository) Delete(id int) error {
    if m.shouldError {
        return errors.New("mock error")
    }
    return nil
}

func (m *MockItemRepository) Search(keyword string) ([]models.Item, error) {
    if m.shouldError {
        return nil, errors.New("mock error")
    }
    return m.items, nil
}

func (m *MockItemRepository) GetItemsNeedReplacement(days int) ([]models.Item, error) {
    if m.shouldError {
        return nil, errors.New("mock error")
    }
    return m.items, nil
}

func TestItemService_Create(t *testing.T) {
    mockItemRepo := &MockItemRepository{}
    mockCatRepo := &MockCategoryRepository{
        categories: []models.Category{
            {ID: 1, Name: "Elektronik"},
        },
    }

    service := NewItemService(mockItemRepo, mockCatRepo)
    item, err := service.Create("Laptop", 1, 15000000, time.Now())

    if err != nil {
        t.Errorf("unexpected error: %s", err)
    }

    if item.Name != "Laptop" {
        t.Errorf("expected name 'Laptop', got '%s'", item.Name)
    }
}

func TestItemService_Create_EmptyName(t *testing.T) {
    mockItemRepo := &MockItemRepository{}
    mockCatRepo := &MockCategoryRepository{}

    service := NewItemService(mockItemRepo, mockCatRepo)
    _, err := service.Create("", 1, 15000000, time.Now())

    if err == nil {
        t.Error("expected error for empty name")
    }
}

func TestItemService_Create_InvalidPrice(t *testing.T) {
    mockItemRepo := &MockItemRepository{}
    mockCatRepo := &MockCategoryRepository{
        categories: []models.Category{{ID: 1}},
    }

    service := NewItemService(mockItemRepo, mockCatRepo)
    _, err := service.Create("Laptop", 1, 0, time.Now())

    if err == nil {
        t.Error("expected error for invalid price")
    }

    _, err = service.Create("Laptop", 1, -100, time.Now())
    if err == nil {
        t.Error("expected error for negative price")
    }
}

func TestItemService_CalculateDepreciation(t *testing.T) {
    mockItemRepo := &MockItemRepository{}
    mockCatRepo := &MockCategoryRepository{}

    service := NewItemService(mockItemRepo, mockCatRepo)

    // Item berusia 1 tahun
    purchaseDate := time.Now().AddDate(-1, 0, 0)
    item := models.Item{
        ID:           1,
        Name:         "Laptop",
        Price:        10000000,
        PurchaseDate: purchaseDate,
    }

    dep := service.CalculateDepreciation(item)

    // Setelah 1 tahun dengan depresiasi 20%, nilai = 10000000 * 0.8 = 8000000
    expectedValue := 8000000.0
    tolerance := 100000.0 // Toleransi untuk floating point

    if dep.CurrentValue < expectedValue-tolerance || dep.CurrentValue > expectedValue+tolerance {
        t.Errorf("expected current value around %.2f, got %.2f", expectedValue, dep.CurrentValue)
    }

    if dep.DepreciationValue < 0 {
        t.Error("depreciation value should not be negative")
    }
}

func TestItemService_Search(t *testing.T) {
    mockItemRepo := &MockItemRepository{
        items: []models.Item{
            {ID: 1, Name: "Laptop Dell"},
            {ID: 2, Name: "Laptop HP"},
        },
    }
    mockCatRepo := &MockCategoryRepository{}

    service := NewItemService(mockItemRepo, mockCatRepo)
    items, err := service.Search("laptop")

    if err != nil {
        t.Errorf("unexpected error: %s", err)
    }

    if len(items) != 2 {
        t.Errorf("expected 2 items, got %d", len(items))
    }
}

func TestItemService_Search_EmptyKeyword(t *testing.T) {
    mockItemRepo := &MockItemRepository{}
    mockCatRepo := &MockCategoryRepository{}

    service := NewItemService(mockItemRepo, mockCatRepo)
    _, err := service.Search("")

    if err == nil {
        t.Error("expected error for empty keyword")
    }
}

func TestItemService_GetTotalInvestment(t *testing.T) {
    purchaseDate := time.Now().AddDate(-1, 0, 0)
    mockItemRepo := &MockItemRepository{
        items: []models.Item{
            {ID: 1, Name: "Laptop", Price: 10000000, PurchaseDate: purchaseDate},
            {ID: 2, Name: "Monitor", Price: 5000000, PurchaseDate: purchaseDate},
        },
    }
    mockCatRepo := &MockCategoryRepository{}

    service := NewItemService(mockItemRepo, mockCatRepo)
    totalOriginal, totalCurrent, err := service.GetTotalInvestment()

    if err != nil {
        t.Errorf("unexpected error: %s", err)
    }

    if totalOriginal != 15000000 {
        t.Errorf("expected total original 15000000, got %.2f", totalOriginal)
    }

    if totalCurrent <= 0 {
        t.Error("total current should be greater than 0")
    }

    if totalCurrent >= totalOriginal {
        t.Error("total current should be less than total original due to depreciation")
    }
}