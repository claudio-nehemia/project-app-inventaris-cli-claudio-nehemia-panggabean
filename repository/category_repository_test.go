package repository

import (
    "database/sql"
    "testing"
    "time"

    "github.com/DATA-DOG/go-sqlmock"
    "mini_project3/models"
)

func TestCategoryRepository_GetAll(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    repo := NewCategoryRepository(db)

    rows := sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at"}).
        AddRow(1, "Elektronik", "Peralatan elektronik", time.Now(), time.Now()).
        AddRow(2, "Furniture", "Mebel kantor", time.Now(), time.Now())

    mock.ExpectQuery("SELECT id, name, description, created_at, updated_at FROM categories ORDER BY id").
        WillReturnRows(rows)

    categories, err := repo.GetAll()
    if err != nil {
        t.Errorf("error was not expected: %s", err)
    }

    if len(categories) != 2 {
        t.Errorf("expected 2 categories, got %d", len(categories))
    }

    if categories[0].Name != "Elektronik" {
        t.Errorf("expected name 'Elektronik', got '%s'", categories[0].Name)
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}

func TestCategoryRepository_GetByID(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    repo := NewCategoryRepository(db)

    rows := sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at"}).
        AddRow(1, "Elektronik", "Peralatan elektronik", time.Now(), time.Now())

    mock.ExpectQuery("SELECT id, name, description, created_at, updated_at FROM categories WHERE id = \\$1").
        WithArgs(1).
        WillReturnRows(rows)

    category, err := repo.GetByID(1)
    if err != nil {
        t.Errorf("error was not expected: %s", err)
    }

    if category.Name != "Elektronik" {
        t.Errorf("expected name 'Elektronik', got '%s'", category.Name)
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}

func TestCategoryRepository_GetByID_NotFound(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    repo := NewCategoryRepository(db)

    mock.ExpectQuery("SELECT id, name, description, created_at, updated_at FROM categories WHERE id = \\$1").
        WithArgs(999).
        WillReturnError(sql.ErrNoRows)

    _, err = repo.GetByID(999)
    if err == nil {
        t.Error("expected error, got nil")
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}

func TestCategoryRepository_Create(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    repo := NewCategoryRepository(db)

    cat := &models.Category{
        Name:        "Test Category",
        Description: "Test Description",
    }

    rows := sqlmock.NewRows([]string{"id", "created_at"}).
        AddRow(1, time.Now())

    mock.ExpectQuery("INSERT INTO categories \\(name, description, updated_at\\) VALUES \\(\\$1, \\$2, \\$3\\) RETURNING id, created_at").
        WithArgs(cat.Name, cat.Description, sqlmock.AnyArg()).
        WillReturnRows(rows)

    err = repo.Create(cat)
    if err != nil {
        t.Errorf("error was not expected: %s", err)
    }

    if cat.ID != 1 {
        t.Errorf("expected ID 1, got %d", cat.ID)
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}

func TestCategoryRepository_Update(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    repo := NewCategoryRepository(db)

    cat := &models.Category{
        ID:          1,
        Name:        "Updated Category",
        Description: "Updated Description",
    }

    mock.ExpectExec("UPDATE categories SET name = \\$1, description = \\$2, updated_at = \\$3 WHERE id = \\$4").
        WithArgs(cat.Name, cat.Description, sqlmock.AnyArg(), cat.ID).
        WillReturnResult(sqlmock.NewResult(0, 1))

    err = repo.Update(cat)
    if err != nil {
        t.Errorf("error was not expected: %s", err)
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}

func TestCategoryRepository_Delete(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    repo := NewCategoryRepository(db)

    mock.ExpectExec("DELETE FROM categories WHERE id = \\$1").
        WithArgs(1).
        WillReturnResult(sqlmock.NewResult(0, 1))

    err = repo.Delete(1)
    if err != nil {
        t.Errorf("error was not expected: %s", err)
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}

func TestCategoryRepository_CheckNameExists(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    repo := NewCategoryRepository(db)

    rows := sqlmock.NewRows([]string{"count"}).AddRow(1)

    mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM categories WHERE name = \\$1 AND id != \\$2").
        WithArgs("Elektronik", 0).
        WillReturnRows(rows)

    exists, err := repo.CheckNameExists("Elektronik", 0)
    if err != nil {
        t.Errorf("error was not expected: %s", err)
    }

    if !exists {
        t.Error("expected exists to be true")
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}