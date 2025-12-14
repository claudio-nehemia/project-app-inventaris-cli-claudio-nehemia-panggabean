package utils

import (
    "fmt"
    "os"
    "text/tabwriter"
    "time"
)

func PrintCategoriesTable(categories interface{}) {
    w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
    fmt.Fprintln(w, "ID\tNama\tDeskripsi\tDibuat")
    fmt.Fprintln(w, "---\t---\t---\t---")

    switch v := categories.(type) {
    case []interface{}:
        for _, cat := range v {
            if catMap, ok := cat.(map[string]interface{}); ok {
                fmt.Fprintf(w, "%v\t%v\t%v\t%v\n",
                    catMap["id"],
                    catMap["name"],
                    catMap["description"],
                    catMap["created_at"])
            }
        }
    }

    w.Flush()
}

func PrintItemsTable(items interface{}) {
    w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
    fmt.Fprintln(w, "ID\tNama\tKategori\tHarga\tTgl Beli\tHari Digunakan")
    fmt.Fprintln(w, "---\t---\t---\t---\t---\t---")

    switch v := items.(type) {
    case []interface{}:
        for _, item := range v {
            if itemMap, ok := item.(map[string]interface{}); ok {
                purchaseDate, _ := time.Parse(time.RFC3339, itemMap["purchase_date"].(string))
                daysUsed := int(time.Since(purchaseDate).Hours() / 24)

                fmt.Fprintf(w, "%v\t%v\t%v\tRp %.2f\t%v\t%d hari\n",
                    itemMap["id"],
                    itemMap["name"],
                    itemMap["category_name"],
                    itemMap["price"],
                    purchaseDate.Format("2006-01-02"),
                    daysUsed)
            }
        }
    }

    w.Flush()
}

func PrintDepreciationTable(items interface{}) {
    w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
    fmt.Fprintln(w, "ID\tNama\tHarga Awal\tNilai Sekarang\tDepresiasi\tHari Digunakan")
    fmt.Fprintln(w, "---\t---\t---\t---\t---\t---")

    switch v := items.(type) {
    case []interface{}:
        for _, item := range v {
            if itemMap, ok := item.(map[string]interface{}); ok {
                fmt.Fprintf(w, "%v\t%v\tRp %.2f\tRp %.2f\tRp %.2f\t%v hari\n",
                    itemMap["id"],
                    itemMap["name"],
                    itemMap["price"],
                    itemMap["current_value"],
                    itemMap["depreciation_value"],
                    itemMap["days_used"])
            }
        }
    }

    w.Flush()
}