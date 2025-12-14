package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"mini_project3/config"
	"mini_project3/handler"
	"mini_project3/repository"
	"mini_project3/service"

	"github.com/spf13/cobra"
)

var (
	categoryHandler *handler.CategoryHandler
	itemHandler     *handler.ItemHandler
)

func main() {
	// Initialize database
	cfg := config.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "clau",
		DBName:   "inventory_office",
	}

	db, err := config.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	categoryRepo := repository.NewCategoryRepository(db)
	itemRepo := repository.NewItemRepository(db)

	// Initialize services
	categoryService := service.NewCategoryServiceWithRepo(categoryRepo)
	itemService := service.NewItemServiceWithRepo(itemRepo, categoryRepo)

	// Initialize handlers
	categoryHandler = handler.NewCategoryHandler(categoryService)
	itemHandler = handler.NewItemHandler(itemService)

	// Execute root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "inventory",
	Short: "Sistem Inventaris Barang Kantor",
	Long:  `Aplikasi CLI untuk mengelola inventaris barang kantor dengan fitur kategori, barang, dan laporan depresiasi.`,
}

func init() {
	rootCmd.AddCommand(categoryCmd)
	rootCmd.AddCommand(itemCmd)
	rootCmd.AddCommand(reportCmd)
}

// ==================== CATEGORY COMMANDS ====================

var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Kelola kategori barang",
}

var categoryListCmd = &cobra.Command{
	Use:   "list",
	Short: "Tampilkan semua kategori",
	Run: func(cmd *cobra.Command, args []string) {
		if err := categoryHandler.ListCategories(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

var categoryGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Tampilkan detail kategori berdasarkan ID",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		if err := categoryHandler.GetCategory(id); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

var categoryCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Tambah kategori baru",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		desc, _ := cmd.Flags().GetString("description")
		if err := categoryHandler.CreateCategory(name, desc); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

var categoryUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update kategori",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		desc, _ := cmd.Flags().GetString("description")
		if err := categoryHandler.UpdateCategory(id, name, desc); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

var categoryDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Hapus kategori",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		if err := categoryHandler.DeleteCategory(id); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	categoryCmd.AddCommand(categoryListCmd)
	categoryCmd.AddCommand(categoryGetCmd)
	categoryCmd.AddCommand(categoryCreateCmd)
	categoryCmd.AddCommand(categoryUpdateCmd)
	categoryCmd.AddCommand(categoryDeleteCmd)

	// Flags for category commands
	categoryGetCmd.Flags().IntP("id", "i", 0, "Category ID")
	categoryGetCmd.MarkFlagRequired("id")

	categoryCreateCmd.Flags().StringP("name", "n", "", "Category name")
	categoryCreateCmd.Flags().StringP("description", "d", "", "Category description")
	categoryCreateCmd.MarkFlagRequired("name")

	categoryUpdateCmd.Flags().IntP("id", "i", 0, "Category ID")
	categoryUpdateCmd.Flags().StringP("name", "n", "", "Category name")
	categoryUpdateCmd.Flags().StringP("description", "d", "", "Category description")
	categoryUpdateCmd.MarkFlagRequired("id")
	categoryUpdateCmd.MarkFlagRequired("name")

	categoryDeleteCmd.Flags().IntP("id", "i", 0, "Category ID")
	categoryDeleteCmd.MarkFlagRequired("id")
}

// ==================== ITEM COMMANDS ====================

var itemCmd = &cobra.Command{
	Use:   "item",
	Short: "Kelola barang inventaris",
}

var itemListCmd = &cobra.Command{
	Use:   "list",
	Short: "Tampilkan semua barang",
	Run: func(cmd *cobra.Command, args []string) {
		if err := itemHandler.ListItems(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

var itemGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Tampilkan detail barang berdasarkan ID",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		if err := itemHandler.GetItem(id); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

var itemCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Tambah barang baru",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		categoryID, _ := cmd.Flags().GetInt("category")
		price, _ := cmd.Flags().GetFloat64("price")
		dateStr, _ := cmd.Flags().GetString("date")

		purchaseDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: invalid date format (use YYYY-MM-DD): %v\n", err)
			os.Exit(1)
		}

		if err := itemHandler.CreateItem(name, categoryID, price, purchaseDate); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

var itemUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update barang",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		categoryID, _ := cmd.Flags().GetInt("category")
		price, _ := cmd.Flags().GetFloat64("price")
		dateStr, _ := cmd.Flags().GetString("date")

		purchaseDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: invalid date format (use YYYY-MM-DD): %v\n", err)
			os.Exit(1)
		}

		if err := itemHandler.UpdateItem(id, name, categoryID, price, purchaseDate); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

var itemDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Hapus barang",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		if err := itemHandler.DeleteItem(id); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

var itemSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Cari barang berdasarkan nama",
	Run: func(cmd *cobra.Command, args []string) {
		keyword, _ := cmd.Flags().GetString("keyword")
		if err := itemHandler.SearchItems(keyword); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

var itemReplacementCmd = &cobra.Command{
	Use:   "replacement",
	Short: "Tampilkan barang yang perlu diganti (> 100 hari)",
	Run: func(cmd *cobra.Command, args []string) {
		if err := itemHandler.ListItemsNeedReplacement(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	itemCmd.AddCommand(itemListCmd)
	itemCmd.AddCommand(itemGetCmd)
	itemCmd.AddCommand(itemCreateCmd)
	itemCmd.AddCommand(itemUpdateCmd)
	itemCmd.AddCommand(itemDeleteCmd)
	itemCmd.AddCommand(itemSearchCmd)
	itemCmd.AddCommand(itemReplacementCmd)

	// Flags for item commands
	itemGetCmd.Flags().IntP("id", "i", 0, "Item ID")
	itemGetCmd.MarkFlagRequired("id")

	itemCreateCmd.Flags().StringP("name", "n", "", "Item name")
	itemCreateCmd.Flags().IntP("category", "c", 0, "Category ID")
	itemCreateCmd.Flags().Float64P("price", "p", 0, "Item price")
	itemCreateCmd.Flags().StringP("date", "d", "", "Purchase date (YYYY-MM-DD)")
	itemCreateCmd.MarkFlagRequired("name")
	itemCreateCmd.MarkFlagRequired("category")
	itemCreateCmd.MarkFlagRequired("price")
	itemCreateCmd.MarkFlagRequired("date")

	itemUpdateCmd.Flags().IntP("id", "i", 0, "Item ID")
	itemUpdateCmd.Flags().StringP("name", "n", "", "Item name")
	itemUpdateCmd.Flags().IntP("category", "c", 0, "Category ID")
	itemUpdateCmd.Flags().Float64P("price", "p", 0, "Item price")
	itemUpdateCmd.Flags().StringP("date", "d", "", "Purchase date (YYYY-MM-DD)")
	itemUpdateCmd.MarkFlagRequired("id")
	itemUpdateCmd.MarkFlagRequired("name")
	itemUpdateCmd.MarkFlagRequired("category")
	itemUpdateCmd.MarkFlagRequired("price")
	itemUpdateCmd.MarkFlagRequired("date")

	itemDeleteCmd.Flags().IntP("id", "i", 0, "Item ID")
	itemDeleteCmd.MarkFlagRequired("id")

	itemSearchCmd.Flags().StringP("keyword", "k", "", "Search keyword")
	itemSearchCmd.MarkFlagRequired("keyword")
}

// ==================== REPORT COMMANDS ====================

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Laporan investasi dan depresiasi",
}

var reportTotalCmd = &cobra.Command{
	Use:   "total",
	Short: "Tampilkan total investasi dan depresiasi",
	Run: func(cmd *cobra.Command, args []string) {
		if err := itemHandler.ShowTotalInvestment(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

var reportItemCmd = &cobra.Command{
	Use:   "item",
	Short: "Tampilkan laporan depresiasi barang tertentu",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		if err := itemHandler.ShowItemDepreciation(id); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	reportCmd.AddCommand(reportTotalCmd)
	reportCmd.AddCommand(reportItemCmd)

	reportItemCmd.Flags().IntP("id", "i", 0, "Item ID")
	reportItemCmd.MarkFlagRequired("id")
}
