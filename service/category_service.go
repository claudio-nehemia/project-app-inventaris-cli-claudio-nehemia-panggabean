package service

import (
	"fmt"
	"strings"

	"mini_project3/models"
	"mini_project3/repository"
	"mini_project3/utils"
)

// CategoryRepositoryInterface defines the contract for category repository
type CategoryRepositoryInterface interface {
	GetAll() ([]models.Category, error)
	GetByID(id int) (*models.Category, error)
	Create(cat *models.Category) error
	Update(cat *models.Category) error
	Delete(id int) error
	CheckNameExists(name string, excludeID int) (bool, error)
}

type CategoryService struct {
	repo CategoryRepositoryInterface
}

func NewCategoryService(repo CategoryRepositoryInterface) *CategoryService {
	return &CategoryService{repo: repo}
}

// NewCategoryServiceWithRepo creates CategoryService with concrete repository (for production)
func NewCategoryServiceWithRepo(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) GetByID(id int) (*models.Category, error) {
	if err := utils.ValidateID(id); err != nil {
		return nil, err
	}
	return s.repo.GetByID(id)
}

func (s *CategoryService) Create(name, description string) (*models.Category, error) {
	name = strings.TrimSpace(name)
	if err := utils.ValidateNotEmpty(name, "Category name"); err != nil {
		return nil, err
	}

	// Check for duplicate
	exists, err := s.repo.CheckNameExists(name, 0)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("category with name '%s' already exists", name)
	}

	cat := &models.Category{
		Name:        name,
		Description: strings.TrimSpace(description),
	}

	if err := s.repo.Create(cat); err != nil {
		return nil, err
	}

	return cat, nil
}

func (s *CategoryService) Update(id int, name, description string) error {
	if err := utils.ValidateID(id); err != nil {
		return err
	}

	name = strings.TrimSpace(name)
	if err := utils.ValidateNotEmpty(name, "Category name"); err != nil {
		return err
	}

	// Check for duplicate (excluding current ID)
	exists, err := s.repo.CheckNameExists(name, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("category with name '%s' already exists", name)
	}

	cat := &models.Category{
		ID:          id,
		Name:        name,
		Description: strings.TrimSpace(description),
	}

	return s.repo.Update(cat)
}

func (s *CategoryService) Delete(id int) error {
	if err := utils.ValidateID(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
