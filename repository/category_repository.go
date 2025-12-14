package repository

import (
    "database/sql"
    "fmt"
    "time"

    "mini_project3/models"
)

type CategoryRepository struct {
    db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
    return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll() ([]models.Category, error) {
    query := `SELECT id, name, description, created_at, updated_at FROM categories ORDER BY id`
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, fmt.Errorf("error querying categories: %w", err)
    }
    defer rows.Close()

    var categories []models.Category
    for rows.Next() {
        var cat models.Category
        if err := rows.Scan(&cat.ID, &cat.Name, &cat.Description, &cat.CreatedAt, &cat.UpdatedAt); err != nil {
            return nil, fmt.Errorf("error scanning category: %w", err)
        }
        categories = append(categories, cat)
    }

    return categories, nil
}

func (r *CategoryRepository) GetByID(id int) (*models.Category, error) {
    query := `SELECT id, name, description, created_at, updated_at FROM categories WHERE id = $1`
    var cat models.Category
    err := r.db.QueryRow(query, id).Scan(&cat.ID, &cat.Name, &cat.Description, &cat.CreatedAt, &cat.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("category with ID %d not found", id)
        }
        return nil, fmt.Errorf("error querying category: %w", err)
    }
    return &cat, nil
}

func (r *CategoryRepository) Create(cat *models.Category) error {
    query := `INSERT INTO categories (name, description, updated_at) VALUES ($1, $2, $3) RETURNING id, created_at`
    err := r.db.QueryRow(query, cat.Name, cat.Description, time.Now()).Scan(&cat.ID, &cat.CreatedAt)
    if err != nil {
        return fmt.Errorf("error creating category: %w", err)
    }
    return nil
}

func (r *CategoryRepository) Update(cat *models.Category) error {
    query := `UPDATE categories SET name = $1, description = $2, updated_at = $3 WHERE id = $4`
    result, err := r.db.Exec(query, cat.Name, cat.Description, time.Now(), cat.ID)
    if err != nil {
        return fmt.Errorf("error updating category: %w", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error getting rows affected: %w", err)
    }
    if rows == 0 {
        return fmt.Errorf("category with ID %d not found", cat.ID)
    }

    return nil
}

func (r *CategoryRepository) Delete(id int) error {
    query := `DELETE FROM categories WHERE id = $1`
    result, err := r.db.Exec(query, id)
    if err != nil {
        return fmt.Errorf("error deleting category: %w", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error getting rows affected: %w", err)
    }
    if rows == 0 {
        return fmt.Errorf("category with ID %d not found", id)
    }

    return nil
}

func (r *CategoryRepository) CheckNameExists(name string, excludeID int) (bool, error) {
    query := `SELECT COUNT(*) FROM categories WHERE name = $1 AND id != $2`
    var count int
    err := r.db.QueryRow(query, name, excludeID).Scan(&count)
    if err != nil {
        return false, fmt.Errorf("error checking category name: %w", err)
    }
    return count > 0, nil
}