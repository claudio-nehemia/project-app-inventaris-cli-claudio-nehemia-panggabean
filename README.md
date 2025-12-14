# Sistem Inventaris Barang Kantor - CLI

Aplikasi Command Line Interface (CLI) untuk mengelola inventaris barang kantor dengan fitur kategori, manajemen barang, dan laporan depresiasi.

## Fitur

### 1. Manajemen Kategori
- ✅ Menampilkan daftar kategori
- ✅ Menambahkan kategori baru
- ✅ Melihat detail kategori
- ✅ Mengedit kategori
- ✅ Menghapus kategori

### 2. Manajemen Barang Inventaris
- ✅ Menampilkan daftar barang dengan informasi lengkap
- ✅ Menambahkan barang baru
- ✅ Melihat detail barang
- ✅ Mengedit data barang
- ✅ Menghapus barang
- ✅ Pencarian barang berdasarkan nama

### 3. Barang yang Perlu Diganti
- ✅ Menampilkan barang yang sudah digunakan > 100 hari

### 4. Laporan Investasi dan Depresiasi
- ✅ Laporan total investasi dengan depresiasi
- ✅ Laporan depresiasi per barang
- ✅ Menggunakan metode saldo menurun 20% per tahun

## Requirements

- Go 1.21 atau lebih baru
- PostgreSQL 12 atau lebih baru

## Instalasi

### 1. Clone Repository
```bash
git clone https://github.com/username/project-app-inventaris-cli-nama.git
cd project-app-inventaris-cli-nama
```

### 2. Setup Database
```bash
# Login ke PostgreSQL
psql -U postgres

# Jalankan script SQL
\i database/schema.sql
```

### 3. Install Dependencies
```bash
go mod download
```

### 4. Build Aplikasi
```bash
go build -o inventory cmd/main.go
```

## Penggunaan

### Kategori

#### Lihat Semua Kategori
```bash
./inventory category list
```

#### Tambah Kategori
```bash
./inventory category create --name "Elektronik" --description "Peralatan elektronik kantor"
```

#### Lihat Detail Kategori
```bash
./inventory category get --id 1
```

#### Update Kategori
```bash
./inventory category update --id 1 --name "Elektronik" --description "Updated description"
```

#### Hapus Kategori
```bash
./inventory category delete --id 1
```

### Barang Inventaris

#### Lihat Semua Barang
```bash
./inventory item list
```

#### Tambah Barang
```bash
./inventory item create --name "Laptop Dell XPS 13" --category 1 --price 15000000 --date "2024-06-01"
```

#### Lihat Detail Barang
```bash
./inventory item get --id 1
```

#### Update Barang
```bash
./inventory item update --id 1 --name "Laptop Dell XPS 15" --category 1 --price 18000000 --date "2024-06-01"
```

#### Hapus Barang
```bash
./inventory item delete --id 1
```

#### Cari Barang
```bash
./inventory item search --keyword "laptop"
```

#### Barang yang Perlu Diganti
```bash
./inventory item replacement
```

### Laporan

#### Laporan Total Investasi
```bash
./inventory report total
```

#### Laporan Depresiasi Per Barang
```bash
./inventory report item --id 1
```

## Testing
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Struktur Project
```
project-app-inventaris-cli-nama/
├── cmd/
│   └── main.go              # Entry point aplikasi
├── config/
│   └── database.go          # Konfigurasi database
├── models/
│   ├── category.go          # Model kategori
│   └── item.go              # Model barang
├── repository/
│   ├── category_repository.go  # Repository kategori
│   └── item_repository.go      # Repository barang
├── service/
│   ├── category_service.go  # Business logic kategori
│   └── item_service.go      # Business logic barang
├── handler/
│   ├── category_handler.go  # Handler CLI kategori
│   └── item_handler.go      # Handler CLI barang
├── utils/
│   ├── table.go             # Utility untuk tampilan tabel
│   └── validation.go        # Utility validasi
├── database/
│   └── schema.sql           # Database schema
├── go.mod
├── go.sum
└── README.md
```

## Metode Depresiasi

Aplikasi menggunakan **Metode Saldo Menurun** dengan rate 20% per tahun:
```
Nilai Sekarang = Nilai Awal × (1 - Rate Depresiasi)^Tahun
Nilai Sekarang = Nilai Awal × (1 - 0.20)^Tahun
Nilai Sekarang = Nilai Awal × 0.80^Tahun
```

Contoh:
- Laptop seharga Rp 15.000.000
- Digunakan selama 1 tahun
- Nilai sekarang = 15.000.000 × 0.80^1 = Rp 12.000.000
- Depresiasi = Rp 3.000.000

## Fitur Tambahan

- ✅ Menggunakan Cobra untuk CLI framework
- ✅ Tampilan tabel rapi dengan tabwriter
- ✅ Validasi input data
- ✅ Error handling konsisten
- ✅ Pencarian barang
- ✅ Unit testing dengan coverage 40%+

## Link Youtube Explanation

bit.ly/explanation_mp_inventory