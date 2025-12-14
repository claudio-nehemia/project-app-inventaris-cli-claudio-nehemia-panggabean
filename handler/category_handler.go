package handler

import (
    "fmt"
    "os"
    "text/tabwriter"

    "mini_project3/service"
)

type CategoryHandler struct {
    service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
    return &CategoryHandler{service: service}
}

func (h *CategoryHandler) ListCategories() error {
    categories, err := h.service.GetAll()
    if err != nil {
        return fmt.Errorf("failed to get categories: %w", err)
    }

    if len(categories) == 0 {
        fmt.Println("No categories found.")
        return nil
    }

    w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
    fmt.Fprintln(w, "ID\tNama\tDeskripsi\tDibuat")
    fmt.Fprintln(w, "---\t---\t---\t---")

    for _, cat := range categories {
        fmt.Fprintf(w, "%d\t%s\t%s\t%s\n",
            cat.ID,
            cat.Name,
            cat.Description,
            cat.CreatedAt.Format("2006-01-02 15:04"))
    }

    w.Flush()
    return nil
}

func (h *CategoryHandler) GetCategory(id int) error {
    cat, err := h.service.GetByID(id)
    if err != nil {
        return fmt.Errorf("failed to get category: %w", err)
    }

    fmt.Printf("\n=== Detail Kategori ===\n")
    fmt.Printf("ID          : %d\n", cat.ID)
    fmt.Printf("Nama        : %s\n", cat.Name)
    fmt.Printf("Deskripsi   : %s\n", cat.Description)
    fmt.Printf("Dibuat      : %s\n", cat.CreatedAt.Format("2006-01-02 15:04:05"))
    fmt.Printf("Diperbarui  : %s\n", cat.UpdatedAt.Format("2006-01-02 15:04:05"))

    return nil
}

func (h *CategoryHandler) CreateCategory(name, description string) error {
    cat, err := h.service.Create(name, description)
    if err != nil {
        return fmt.Errorf("failed to create category: %w", err)
    }

    fmt.Printf("\n✓ Kategori berhasil ditambahkan dengan ID: %d\n", cat.ID)
    return nil
}

func (h *CategoryHandler) UpdateCategory(id int, name, description string) error {
    if err := h.service.Update(id, name, description); err != nil {
        return fmt.Errorf("failed to update category: %w", err)
    }

    fmt.Printf("\n✓ Kategori dengan ID %d berhasil diperbarui\n", id)
    return nil
}

func (h *CategoryHandler) DeleteCategory(id int) error {
    if err := h.service.Delete(id); err != nil {
        return fmt.Errorf("failed to delete category: %w", err)
    }

    fmt.Printf("\n✓ Kategori dengan ID %d berhasil dihapus\n", id)
    return nil
}