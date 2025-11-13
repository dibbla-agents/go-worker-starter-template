# How to Create Database Tables with GORM

Define Go structs with GORM tags - tables auto-create on startup. No SQL migrations needed.

## Step 1: Define Model

Create `internal/models/your_table.go`:

```go
package models

import (
    "gorm.io/datatypes"
    "github.com/pgvector/pgvector-go"
)

type YourTable struct {
    ID          uint             `gorm:"primaryKey;autoIncrement" json:"id"`
    Name        string           `gorm:"type:varchar(255);not null" json:"name"`
    Description string           `gorm:"type:text" json:"description"`
    Count       int              `gorm:"type:integer;default:0" json:"count"`
    IsActive    bool             `gorm:"type:boolean;default:true" json:"is_active"`
    Metadata    datatypes.JSON   `gorm:"type:jsonb" json:"metadata,omitempty"`
    Embedding   *pgvector.Vector `gorm:"type:vector(1536)" json:"embedding,omitempty"`
    CreatedAt   *string          `gorm:"type:timestamp" json:"created_at,omitempty"`
}

// Optional: Custom table name (defaults to lowercase plural: "your_tables")
func (YourTable) TableName() string {
    return "your_custom_name"
}
```

## Step 2: Register in `internal/models/models.go`

```go
func AllModels() []interface{} {
    return []interface{}{
        &Status{},
        &YourTable{},  // Add your model here
    }
}
```

## Step 3: Done ✅

GORM creates/updates tables automatically when AsyncGlobalState initializes.

---

## Common GORM Tags

```go
`gorm:"primaryKey"`                      // Primary key
`gorm:"type:varchar(255);not null"`      // Required text field
`gorm:"type:text"`                       // Long text
`gorm:"type:integer;default:0"`          // Integer with default
`gorm:"type:boolean;default:true"`       // Boolean with default
`gorm:"type:jsonb"`                      // JSON data
`gorm:"type:vector(1536)"`               // pgvector embedding
`gorm:"unique"`                          // Unique constraint
`gorm:"index"`                           // Create index
```

**Combine tags:** `gorm:"type:varchar(255);not null;unique;index"`

---

## Type Reference

| Go Type | PostgreSQL | Example |
|---------|------------|---------|
| `string` | VARCHAR/TEXT | `type:varchar(255)` or `type:text` |
| `int` | INTEGER | `type:integer` |
| `bool` | BOOLEAN | `type:boolean` |
| `datatypes.JSON` | JSONB | `type:jsonb` |
| `*pgvector.Vector` | VECTOR | `type:vector(1536)` |
| `*string` | (nullable) | Use pointer for nullable fields |

---

## Dynamic Table Routing

When table name depends on **runtime context** (user input, channel, etc.):

**❌ Don't use `TableName()` method** - it's evaluated once and can cause wrong table access

**✅ Use explicit routing:**

1. Create routing function in `internal/models/table_routing.go`:

```go
package models

func GetYourTableName(context string) string {
    if context == "special" {
        return "special_your_table"
    }
    return "your_table"
}
```

2. Use `.Table()` explicitly in queries:

```go
tableName := models.GetYourTableName(context)
db.Table(tableName).Where("id = ?", id).First(&item)
```

---

## Key Points

- ✅ Add models to `AllModels()` - required for auto-migration
- ✅ Use pointer types (`*string`, `*int`) for nullable columns
- ✅ Use `TableName()` for static table names only
- ❌ Don't write SQL migrations - GORM handles it
- ⚠️ GORM adds columns but won't drop them - manual ALTER needed

## Troubleshooting

| Problem | Solution |
|---------|----------|
| Table not created | Check if model is in `AllModels()` |
| Wrong table name | Add `TableName()` method |
| Column type not updating | GORM won't modify existing types - manual ALTER needed |
| Default name wrong | Default is lowercase plural: `YourTable` → `your_tables` |

