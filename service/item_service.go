package service

import (
	"fmt"
	"math"
	"strings"
	"time"

	"mini_project3/models"
	"mini_project3/repository"
	"mini_project3/utils"
)

// ItemRepositoryInterface defines the contract for item repository
type ItemRepositoryInterface interface {
	GetAll() ([]models.Item, error)
	GetByID(id int) (*models.Item, error)
	Create(item *models.Item) error
	Update(item *models.Item) error
	Delete(id int) error
	Search(keyword string) ([]models.Item, error)
	GetItemsNeedReplacement(days int) ([]models.Item, error)
}

type ItemService struct {
	itemRepo     ItemRepositoryInterface
	categoryRepo CategoryRepositoryInterface
}

func NewItemService(itemRepo ItemRepositoryInterface, categoryRepo CategoryRepositoryInterface) *ItemService {
	return &ItemService{
		itemRepo:     itemRepo,
		categoryRepo: categoryRepo,
	}
}

// NewItemServiceWithRepo creates ItemService with concrete repositories (for production)
func NewItemServiceWithRepo(itemRepo *repository.ItemRepository, categoryRepo *repository.CategoryRepository) *ItemService {
	return &ItemService{
		itemRepo:     itemRepo,
		categoryRepo: categoryRepo,
	}
}

func (s *ItemService) GetAll() ([]models.Item, error) {
	return s.itemRepo.GetAll()
}

func (s *ItemService) GetByID(id int) (*models.Item, error) {
	if err := utils.ValidateID(id); err != nil {
		return nil, err
	}
	return s.itemRepo.GetByID(id)
}

func (s *ItemService) Create(name string, categoryID int, price float64, purchaseDate time.Time) (*models.Item, error) {
	name = strings.TrimSpace(name)
	if err := utils.ValidateNotEmpty(name, "Item name"); err != nil {
		return nil, err
	}

	if err := utils.ValidateID(categoryID); err != nil {
		return nil, fmt.Errorf("invalid category ID: %w", err)
	}

	if price <= 0 {
		return nil, fmt.Errorf("price must be greater than 0")
	}

	// Check if category exists
	_, err := s.categoryRepo.GetByID(categoryID)
	if err != nil {
		return nil, fmt.Errorf("category not found: %w", err)
	}

	item := &models.Item{
		Name:         name,
		CategoryID:   categoryID,
		Price:        price,
		PurchaseDate: purchaseDate,
	}

	if err := s.itemRepo.Create(item); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *ItemService) Update(id int, name string, categoryID int, price float64, purchaseDate time.Time) error {
	if err := utils.ValidateID(id); err != nil {
		return err
	}

	name = strings.TrimSpace(name)
	if err := utils.ValidateNotEmpty(name, "Item name"); err != nil {
		return err
	}

	if err := utils.ValidateID(categoryID); err != nil {
		return fmt.Errorf("invalid category ID: %w", err)
	}

	if price <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}

	// Check if category exists
	_, err := s.categoryRepo.GetByID(categoryID)
	if err != nil {
		return fmt.Errorf("category not found: %w", err)
	}

	item := &models.Item{
		ID:           id,
		Name:         name,
		CategoryID:   categoryID,
		Price:        price,
		PurchaseDate: purchaseDate,
	}

	return s.itemRepo.Update(item)
}

func (s *ItemService) Delete(id int) error {
	if err := utils.ValidateID(id); err != nil {
		return err
	}
	return s.itemRepo.Delete(id)
}

func (s *ItemService) Search(keyword string) ([]models.Item, error) {
	keyword = strings.TrimSpace(keyword)
	if keyword == "" {
		return nil, fmt.Errorf("search keyword cannot be empty")
	}
	return s.itemRepo.Search(keyword)
}

func (s *ItemService) GetItemsNeedReplacement() ([]models.Item, error) {
	return s.itemRepo.GetItemsNeedReplacement(100)
}

// CalculateDepreciation menggunakan metode saldo menurun 20% per tahun
func (s *ItemService) CalculateDepreciation(item models.Item) models.ItemDepreciation {
	now := time.Now()
	daysUsed := int(now.Sub(item.PurchaseDate).Hours() / 24)
	yearsUsed := float64(daysUsed) / 365.0

	// Metode saldo menurun: Nilai Sekarang = Nilai Awal × (1 - Tingkat Depresiasi)^Tahun
	depreciationRate := 0.20
	currentValue := item.Price * math.Pow(1-depreciationRate, yearsUsed)
	depreciationValue := item.Price - currentValue

	return models.ItemDepreciation{
		Item:              item,
		DaysUsed:          daysUsed,
		DepreciationRate:  depreciationRate,
		CurrentValue:      currentValue,
		DepreciationValue: depreciationValue,
	}
}

func (s *ItemService) GetTotalInvestment() (float64, float64, error) {
	items, err := s.itemRepo.GetAll()
	if err != nil {
		return 0, 0, err
	}

	var totalOriginal, totalCurrent float64
	for _, item := range items {
		dep := s.CalculateDepreciation(item)
		totalOriginal += item.Price
		totalCurrent += dep.CurrentValue
	}

	return totalOriginal, totalCurrent, nil
}

func (s *ItemService) GetItemDepreciation(id int) (*models.ItemDepreciation, error) {
	if err := utils.ValidateID(id); err != nil {
		return nil, err
	}

	item, err := s.itemRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	dep := s.CalculateDepreciation(*item)
	return &dep, nil
}
