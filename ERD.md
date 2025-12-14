## Entity Relationship Diagram
erDiagram
    CATEGORIES ||--o{ ITEMS : contains
    
    CATEGORIES {
        serial id PK
        varchar name UK
        text description
        timestamp created_at
        timestamp updated_at
    }
    
    ITEMS {
        serial id PK
        varchar name
        integer category_id FK
        decimal price
        date purchase_date
        timestamp created_at
        timestamp updated_at
    }