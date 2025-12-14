package repository

import (
    "database/sql"
    "fmt"
    "strings"
    "time"

    "mini_project3/models"
)

type ItemRepository struct {
    db *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
    return &ItemRepository{db: db}
}

func (r *ItemRepository) GetAll() ([]models.Item, error) {
    query := `
        SELECT i.id, i.name, i.category_id, c.name, i.price, i.purchase_date, i.created_at, i.updated_at
        FROM items i
        JOIN categories c ON i.category_id = c.id
        ORDER BY i.id
    `
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, fmt.Errorf("error querying items: %w", err)
    }
    defer rows.Close()

    var items []models.Item
    for rows.Next() {
        var item models.Item
        if err := rows.Scan(&item.ID, &item.Name, &item.CategoryID, &item.CategoryName, &item.Price, &item.PurchaseDate, &item.CreatedAt, &item.UpdatedAt); err != nil {
            return nil, fmt.Errorf("error scanning item: %w", err)
        }
        items = append(items, item)
    }

    return items, nil
}

func (r *ItemRepository) GetByID(id int) (*models.Item, error) {
    query := `
        SELECT i.id, i.name, i.category_id, c.name, i.price, i.purchase_date, i.created_at, i.updated_at
        FROM items i
        JOIN categories c ON i.category_id = c.id
        WHERE i.id = $1
    `
    var item models.Item
    err := r.db.QueryRow(query, id).Scan(&item.ID, &item.Name, &item.CategoryID, &item.CategoryName, &item.Price, &item.PurchaseDate, &item.CreatedAt, &item.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("item with ID %d not found", id)
        }
        return nil, fmt.Errorf("error querying item: %w", err)
    }
    return &item, nil
}

func (r *ItemRepository) Create(item *models.Item) error {
    query := `INSERT INTO items (name, category_id, price, purchase_date, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`
    err := r.db.QueryRow(query, item.Name, item.CategoryID, item.Price, item.PurchaseDate, time.Now()).Scan(&item.ID, &item.CreatedAt)
    if err != nil {
        return fmt.Errorf("error creating item: %w", err)
    }
    return nil
}

func (r *ItemRepository) Update(item *models.Item) error {
    query := `UPDATE items SET name = $1, category_id = $2, price = $3, purchase_date = $4, updated_at = $5 WHERE id = $6`
    result, err := r.db.Exec(query, item.Name, item.CategoryID, item.Price, item.PurchaseDate, time.Now(), item.ID)
    if err != nil {
        return fmt.Errorf("error updating item: %w", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error getting rows affected: %w", err)
    }
    if rows == 0 {
        return fmt.Errorf("item with ID %d not found", item.ID)
    }

    return nil
}

func (r *ItemRepository) Delete(id int) error {
    query := `DELETE FROM items WHERE id = $1`
    result, err := r.db.Exec(query, id)
    if err != nil {
        return fmt.Errorf("error deleting item: %w", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error getting rows affected: %w", err)
    }
    if rows == 0 {
        return fmt.Errorf("item with ID %d not found", id)
    }

    return nil
}

func (r *ItemRepository) Search(keyword string) ([]models.Item, error) {
    query := `
        SELECT i.id, i.name, i.category_id, c.name, i.price, i.purchase_date, i.created_at, i.updated_at
        FROM items i
        JOIN categories c ON i.category_id = c.id
        WHERE LOWER(i.name) LIKE LOWER($1)
        ORDER BY i.id
    `
    keyword = "%" + strings.ToLower(keyword) + "%"
    rows, err := r.db.Query(query, keyword)
    if err != nil {
        return nil, fmt.Errorf("error searching items: %w", err)
    }
    defer rows.Close()

    var items []models.Item
    for rows.Next() {
        var item models.Item
        if err := rows.Scan(&item.ID, &item.Name, &item.CategoryID, &item.CategoryName, &item.Price, &item.PurchaseDate, &item.CreatedAt, &item.UpdatedAt); err != nil {
            return nil, fmt.Errorf("error scanning item: %w", err)
        }
        items = append(items, item)
    }

    return items, nil
}

func (r *ItemRepository) GetItemsNeedReplacement(days int) ([]models.Item, error) {
    query := `
        SELECT i.id, i.name, i.category_id, c.name, i.price, i.purchase_date, i.created_at, i.updated_at
        FROM items i
        JOIN categories c ON i.category_id = c.id
        WHERE CURRENT_DATE - i.purchase_date > $1
        ORDER BY i.purchase_date ASC
    `
    rows, err := r.db.Query(query, days)
    if err != nil {
        return nil, fmt.Errorf("error querying items need replacement: %w", err)
    }
    defer rows.Close()

    var items []models.Item
    for rows.Next() {
        var item models.Item
        if err := rows.Scan(&item.ID, &item.Name, &item.CategoryID, &item.CategoryName, &item.Price, &item.PurchaseDate, &item.CreatedAt, &item.UpdatedAt); err != nil {
            return nil, fmt.Errorf("error scanning item: %w", err)
        }
        items = append(items, item)
    }

    return items, nil
}