```mermaid
erDiagram
    CATEGORIES ||--o{ ITEMS : "one-to-many"
    
    CATEGORIES {
        serial id PK "Unique identifier for category"
        varchar(100) name UK "Category name - must be unique"
        text description "Detailed category description"
        timestamp created_at "Record creation timestamp"
        timestamp updated_at "Last update timestamp"
    }
    
    ITEMS {
        serial id PK "Unique identifier for item"
        varchar(200) name "Item name"
        integer category_id FK "Reference to categories table"
        decimal(15-2) price "Purchase price in Rupiah"
        date purchase_date "Date when item was purchased"
        timestamp created_at "Record creation timestamp"
        timestamp updated_at "Last update timestamp"
    }
    }
```