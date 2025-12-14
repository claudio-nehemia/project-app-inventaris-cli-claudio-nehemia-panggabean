package handler

import (
    "fmt"
    "os"
    "text/tabwriter"
    "time"

    "mini_project3/service"
)

type ItemHandler struct {
    service *service.ItemService
}

func NewItemHandler(service *service.ItemService) *ItemHandler {
    return &ItemHandler{service: service}
}

func (h *ItemHandler) ListItems() error {
    items, err := h.service.GetAll()
    if err != nil {
        return fmt.Errorf("failed to get items: %w", err)
    }

    if len(items) == 0 {
        fmt.Println("No items found.")
        return nil
    }

    w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
    fmt.Fprintln(w, "ID\tNama\tKategori\tHarga\tTgl Beli\tHari Digunakan")
    fmt.Fprintln(w, "---\t---\t---\t---\t---\t---")

    for _, item := range items {
        daysUsed := int(time.Since(item.PurchaseDate).Hours() / 24)
        fmt.Fprintf(w, "%d\t%s\t%s\tRp %.2f\t%s\t%d hari\n",
            item.ID,
            item.Name,
            item.CategoryName,
            item.Price,
            item.PurchaseDate.Format("2006-01-02"),
            daysUsed)
    }

    w.Flush()
    return nil
}

func (h *ItemHandler) GetItem(id int) error {
    item, err := h.service.GetByID(id)
    if err != nil {
        return fmt.Errorf("failed to get item: %w", err)
    }

    daysUsed := int(time.Since(item.PurchaseDate).Hours() / 24)

    fmt.Printf("\n=== Detail Barang ===\n")
    fmt.Printf("ID              : %d\n", item.ID)
    fmt.Printf("Nama            : %s\n", item.Name)
    fmt.Printf("Kategori        : %s (ID: %d)\n", item.CategoryName, item.CategoryID)
    fmt.Printf("Harga           : Rp %.2f\n", item.Price)
    fmt.Printf("Tgl Beli        : %s\n", item.PurchaseDate.Format("2006-01-02"))
    fmt.Printf("Hari Digunakan  : %d hari\n", daysUsed)
    fmt.Printf("Dibuat          : %s\n", item.CreatedAt.Format("2006-01-02 15:04:05"))
    fmt.Printf("Diperbarui      : %s\n", item.UpdatedAt.Format("2006-01-02 15:04:05"))

    return nil
}

func (h *ItemHandler) CreateItem(name string, categoryID int, price float64, purchaseDate time.Time) error {
    item, err := h.service.Create(name, categoryID, price, purchaseDate)
    if err != nil {
        return fmt.Errorf("failed to create item: %w", err)
    }

    fmt.Printf("\n✓ Barang berhasil ditambahkan dengan ID: %d\n", item.ID)
    return nil
}

func (h *ItemHandler) UpdateItem(id int, name string, categoryID int, price float64, purchaseDate time.Time) error {
    if err := h.service.Update(id, name, categoryID, price, purchaseDate); err != nil {
        return fmt.Errorf("failed to update item: %w", err)
    }

    fmt.Printf("\n✓ Barang dengan ID %d berhasil diperbarui\n", id)
    return nil
}

func (h *ItemHandler) DeleteItem(id int) error {
    if err := h.service.Delete(id); err != nil {
        return fmt.Errorf("failed to delete item: %w", err)
    }

    fmt.Printf("\n✓ Barang dengan ID %d berhasil dihapus\n", id)
    return nil
}

func (h *ItemHandler) SearchItems(keyword string) error {
    items, err := h.service.Search(keyword)
    if err != nil {
        return fmt.Errorf("failed to search items: %w", err)
    }

    if len(items) == 0 {
        fmt.Printf("Tidak ada barang ditemukan dengan kata kunci '%s'\n", keyword)
        return nil
    }

    fmt.Printf("\nHasil pencarian untuk '%s':\n\n", keyword)

    w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
    fmt.Fprintln(w, "ID\tNama\tKategori\tHarga\tTgl Beli\tHari Digunakan")
    fmt.Fprintln(w, "---\t---\t---\t---\t---\t---")

    for _, item := range items {
        daysUsed := int(time.Since(item.PurchaseDate).Hours() / 24)
        fmt.Fprintf(w, "%d\t%s\t%s\tRp %.2f\t%s\t%d hari\n",
            item.ID,
            item.Name,
            item.CategoryName,
            item.Price,
            item.PurchaseDate.Format("2006-01-02"),
            daysUsed)
    }

    w.Flush()
    return nil
}

func (h *ItemHandler) ListItemsNeedReplacement() error {
    items, err := h.service.GetItemsNeedReplacement()
    if err != nil {
        return fmt.Errorf("failed to get items need replacement: %w", err)
    }

    if len(items) == 0 {
        fmt.Println("Tidak ada barang yang perlu diganti (> 100 hari)")
        return nil
    }

    fmt.Printf("\n=== Barang yang Perlu Diganti (> 100 hari) ===\n\n")

    w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
    fmt.Fprintln(w, "ID\tNama\tKategori\tHarga\tTgl Beli\tHari Digunakan")
    fmt.Fprintln(w, "---\t---\t---\t---\t---\t---")

    for _, item := range items {
        daysUsed := int(time.Since(item.PurchaseDate).Hours() / 24)
        fmt.Fprintf(w, "%d\t%s\t%s\tRp %.2f\t%s\t%d hari\n",
            item.ID,
            item.Name,
            item.CategoryName,
            item.Price,
            item.PurchaseDate.Format("2006-01-02"),
            daysUsed)
    }

    w.Flush()
    
    fmt.Printf("\nTotal: %d barang perlu diganti\n", len(items))
    return nil
}

func (h *ItemHandler) ShowTotalInvestment() error {
    totalOriginal, totalCurrent, err := h.service.GetTotalInvestment()
    if err != nil {
        return fmt.Errorf("failed to calculate total investment: %w", err)
    }

    totalDepreciation := totalOriginal - totalCurrent
    percentageDepreciation := 0.0
    if totalOriginal > 0 {
        percentageDepreciation = (totalDepreciation / totalOriginal) * 100
    }

    fmt.Printf("\n=== Laporan Total Investasi ===\n")
    fmt.Printf("Total Investasi Awal    : Rp %s\n", formatCurrency(totalOriginal))
    fmt.Printf("Total Nilai Sekarang    : Rp %s\n", formatCurrency(totalCurrent))
    fmt.Printf("Total Depresiasi        : Rp %s\n", formatCurrency(totalDepreciation))
    fmt.Printf("Persentase Depresiasi   : %.2f%%\n", percentageDepreciation)
    fmt.Printf("\nMetode Depresiasi: Saldo Menurun 20%% per tahun\n")

    return nil
}

func (h *ItemHandler) ShowItemDepreciation(id int) error {
    dep, err := h.service.GetItemDepreciation(id)
    if err != nil {
        return fmt.Errorf("failed to calculate item depreciation: %w", err)
    }

    yearsUsed := float64(dep.DaysUsed) / 365.0
    percentageDepreciation := 0.0
    if dep.Price > 0 {
        percentageDepreciation = (dep.DepreciationValue / dep.Price) * 100
    }

    fmt.Printf("\n=== Laporan Depresiasi Barang ===\n")
    fmt.Printf("ID                  : %d\n", dep.ID)
    fmt.Printf("Nama                : %s\n", dep.Name)
    fmt.Printf("Kategori            : %s\n", dep.CategoryName)
    fmt.Printf("Harga Awal          : Rp %s\n", formatCurrency(dep.Price))
    fmt.Printf("Tanggal Beli        : %s\n", dep.PurchaseDate.Format("2006-01-02"))
    fmt.Printf("Hari Digunakan      : %d hari (%.2f tahun)\n", dep.DaysUsed, yearsUsed)
    fmt.Printf("Rate Depresiasi     : %.0f%% per tahun\n", dep.DepreciationRate*100)
    fmt.Printf("Nilai Sekarang      : Rp %s\n", formatCurrency(dep.CurrentValue))
    fmt.Printf("Total Depresiasi    : Rp %s\n", formatCurrency(dep.DepreciationValue))
    fmt.Printf("Persentase Depresiasi: %.2f%%\n", percentageDepreciation)
    fmt.Printf("\nMetode: Saldo Menurun (Declining Balance Method)\n")
    fmt.Printf("Formula: Nilai Sekarang = Harga Awal × (1 - 0.20)^tahun\n")

    return nil
}

// formatCurrency memformat angka menjadi format mata uang Indonesia
func formatCurrency(amount float64) string {
    // Format dengan pemisah ribuan
    intPart := int64(amount)
    decPart := int((amount - float64(intPart)) * 100)
    
    // Konversi ke string dengan pemisah
    result := ""
    intStr := fmt.Sprintf("%d", intPart)
    
    for i, digit := range intStr {
        if i > 0 && (len(intStr)-i)%3 == 0 {
            result += "."
        }
        result += string(digit)
    }
    
    return fmt.Sprintf("%s,%02d", result, decPart)
}