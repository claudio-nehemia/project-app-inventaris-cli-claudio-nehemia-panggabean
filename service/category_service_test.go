package service

import (
    "errors"
    "testing"
    "time"

    "mini_project3/models"
)

// Mock Repository
type MockCategoryRepository struct {
    categories     []models.Category
    shouldError    bool
    checkNameError bool
    nameExists     bool
}

func (m *MockCategoryRepository) GetAll() ([]models.Category, error) {
    if m.shouldError {
        return nil, errors.New("mock error")
    }
    return m.categories, nil
}

func (m *MockCategoryRepository) GetByID(id int) (*models.Category, error) {
    if m.shouldError {
        return nil, errors.New("mock error")
    }
    for _, cat := range m.categories {
        if cat.ID == id {
            return &cat, nil
        }
    }
    return nil, errors.New("category not found")
}

func (m *MockCategoryRepository) Create(cat *models.Category) error {
    if m.shouldError {
        return errors.New("mock error")
    }
    cat.ID = len(m.categories) + 1
    cat.CreatedAt = time.Now()
    m.categories = append(m.categories, *cat)
    return nil
}

func (m *MockCategoryRepository) Update(cat *models.Category) error {
    if m.shouldError {
        return errors.New("mock error")
    }
    return nil
}

func (m *MockCategoryRepository) Delete(id int) error {
    if m.shouldError {
        return errors.New("mock error")
    }
    return nil
}

func (m *MockCategoryRepository) CheckNameExists(name string, excludeID int) (bool, error) {
    if m.checkNameError {
        return false, errors.New("mock error")
    }
    return m.nameExists, nil
}

func TestCategoryService_GetAll(t *testing.T) {
    mockRepo := &MockCategoryRepository{
        categories: []models.Category{
            {ID: 1, Name: "Elektronik", Description: "Peralatan elektronik"},
            {ID: 2, Name: "Furniture", Description: "Mebel kantor"},
        },
    }

    service := NewCategoryService(mockRepo)
    categories, err := service.GetAll()

    if err != nil {
        t.Errorf("unexpected error: %s", err)
    }

    if len(categories) != 2 {
        t.Errorf("expected 2 categories, got %d", len(categories))
    }
}

func TestCategoryService_GetByID(t *testing.T) {
    mockRepo := &MockCategoryRepository{
        categories: []models.Category{
            {ID: 1, Name: "Elektronik", Description: "Peralatan elektronik"},
        },
    }

    service := NewCategoryService(mockRepo)
    category, err := service.GetByID(1)

    if err != nil {
        t.Errorf("unexpected error: %s", err)
    }

    if category.Name != "Elektronik" {
        t.Errorf("expected name 'Elektronik', got '%s'", category.Name)
    }
}

func TestCategoryService_GetByID_InvalidID(t *testing.T) {
    mockRepo := &MockCategoryRepository{}
    service := NewCategoryService(mockRepo)

    _, err := service.GetByID(0)
    if err == nil {
        t.Error("expected error for invalid ID")
    }

    _, err = service.GetByID(-1)
    if err == nil {
        t.Error("expected error for negative ID")
    }
}

func TestCategoryService_Create(t *testing.T) {
    mockRepo := &MockCategoryRepository{
        nameExists: false,
    }

    service := NewCategoryService(mockRepo)
    cat, err := service.Create("Test Category", "Test Description")

    if err != nil {
        t.Errorf("unexpected error: %s", err)
    }

    if cat.Name != "Test Category" {
        t.Errorf("expected name 'Test Category', got '%s'", cat.Name)
    }
}

func TestCategoryService_Create_EmptyName(t *testing.T) {
    mockRepo := &MockCategoryRepository{}
    service := NewCategoryService(mockRepo)

    _, err := service.Create("", "Description")
    if err == nil {
        t.Error("expected error for empty name")
    }

    _, err = service.Create("   ", "Description")
    if err == nil {
        t.Error("expected error for whitespace name")
    }
}

func TestCategoryService_Create_DuplicateName(t *testing.T) {
    mockRepo := &MockCategoryRepository{
        nameExists: true,
    }

    service := NewCategoryService(mockRepo)
    _, err := service.Create("Existing Category", "Description")

    if err == nil {
        t.Error("expected error for duplicate name")
    }
}

func TestCategoryService_Update(t *testing.T) {
    mockRepo := &MockCategoryRepository{
        categories: []models.Category{
            {ID: 1, Name: "Old Name", Description: "Old Description"},
        },
        nameExists: false,
    }

    service := NewCategoryService(mockRepo)
    err := service.Update(1, "New Name", "New Description")

    if err != nil {
        t.Errorf("unexpected error: %s", err)
    }
}

func TestCategoryService_Delete(t *testing.T) {
    mockRepo := &MockCategoryRepository{
        categories: []models.Category{
            {ID: 1, Name: "Category", Description: "Description"},
        },
    }

    service := NewCategoryService(mockRepo)
    err := service.Delete(1)

    if err != nil {
        t.Errorf("unexpected error: %s", err)
    }
}