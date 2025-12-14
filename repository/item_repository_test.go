package repository

import (
    "testing"
    "time"

    "github.com/DATA-DOG/go-sqlmock"
	"mini_project3/models"
)

func TestItemRepository_GetAll(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    repo := NewItemRepository(db)

    rows := sqlmock.NewRows([]string{"id", "name", "category_id", "category_name", "price", "purchase_date", "created_at", "updated_at"}).
        AddRow(1, "Laptop", 1, "Elektronik", 15000000.00, time.Now(), time.Now(), time.Now()).
        AddRow(2, "Meja", 2, "Furniture", 1500000.00, time.Now(), time.Now(), time.Now())

    mock.ExpectQuery("SELECT i.id, i.name, i.category_id, c.name, i.price, i.purchase_date, i.created_at, i.updated_at FROM items i JOIN categories c").
        WillReturnRows(rows)

    items, err := repo.GetAll()
    if err != nil {
        t.Errorf("error was not expected: %s", err)
    }

    if len(items) != 2 {
        t.Errorf("expected 2 items, got %d", len(items))
    }

    if items[0].Name != "Laptop" {
        t.Errorf("expected name 'Laptop', got '%s'", items[0].Name)
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}

func TestItemRepository_GetByID(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    repo := NewItemRepository(db)

    rows := sqlmock.NewRows([]string{"id", "name", "category_id", "category_name", "price", "purchase_date", "created_at", "updated_at"}).
        AddRow(1, "Laptop", 1, "Elektronik", 15000000.00, time.Now(), time.Now(), time.Now())

    mock.ExpectQuery("SELECT i.id, i.name, i.category_id, c.name, i.price, i.purchase_date, i.created_at, i.updated_at FROM items i JOIN categories c").
        WithArgs(1).
        WillReturnRows(rows)

    item, err := repo.GetByID(1)
    if err != nil {
        t.Errorf("error was not expected: %s", err)
    }

    if item.Name != "Laptop" {
        t.Errorf("expected name 'Laptop', got '%s'", item.Name)
    }

    if item.CategoryName != "Elektronik" {
        t.Errorf("expected category 'Elektronik', got '%s'", item.CategoryName)
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}

func TestItemRepository_Create(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    repo := NewItemRepository(db)

    purchaseDate := time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
    rows := sqlmock.NewRows([]string{"id", "created_at"}).
        AddRow(1, time.Now())

    mock.ExpectQuery("INSERT INTO items \\(name, category_id, price, purchase_date, updated_at\\) VALUES").
        WithArgs("Laptop", 1, 15000000.00, purchaseDate, sqlmock.AnyArg()).
        WillReturnRows(rows)

    item := &models.Item{
        Name:         "Laptop",
        CategoryID:   1,
        Price:        15000000.00,
        PurchaseDate: purchaseDate,
    }

    err = repo.Create(item)
    if err != nil {
        t.Errorf("error was not expected: %s", err)
    }

    if item.ID != 1 {
        t.Errorf("expected ID 1, got %d", item.ID)
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}

func TestItemRepository_Search(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    repo := NewItemRepository(db)

    rows := sqlmock.NewRows([]string{"id", "name", "category_id", "category_name", "price", "purchase_date", "created_at", "updated_at"}).
        AddRow(1, "Laptop Dell", 1, "Elektronik", 15000000.00, time.Now(), time.Now(), time.Now()).
        AddRow(2, "Laptop HP", 1, "Elektronik", 12000000.00, time.Now(), time.Now(), time.Now())

    mock.ExpectQuery("SELECT i.id, i.name, i.category_id, c.name, i.price, i.purchase_date, i.created_at, i.updated_at FROM items i JOIN categories c").
        WithArgs("%laptop%").
        WillReturnRows(rows)

    items, err := repo.Search("laptop")
    if err != nil {
        t.Errorf("error was not expected: %s", err)
    }

    if len(items) != 2 {
        t.Errorf("expected 2 items, got %d", len(items))
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}

func TestItemRepository_GetItemsNeedReplacement(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    repo := NewItemRepository(db)

    oldDate := time.Now().AddDate(0, 0, -150)
    rows := sqlmock.NewRows([]string{"id", "name", "category_id", "category_name", "price", "purchase_date", "created_at", "updated_at"}).
        AddRow(1, "Old Laptop", 1, "Elektronik", 15000000.00, oldDate, time.Now(), time.Now())

    mock.ExpectQuery("SELECT i.id, i.name, i.category_id, c.name, i.price, i.purchase_date, i.created_at, i.updated_at FROM items i JOIN categories c").
        WithArgs(100).
        WillReturnRows(rows)

    items, err := repo.GetItemsNeedReplacement(100)
    if err != nil {
        t.Errorf("error was not expected: %s", err)
    }

    if len(items) != 1 {
        t.Errorf("expected 1 item, got %d", len(items))
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}