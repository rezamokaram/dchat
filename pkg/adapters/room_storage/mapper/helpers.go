package mapper

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

func ToNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: !t.IsZero(),
	}
}

func ToDeletedAt(t time.Time) gorm.DeletedAt {
	return gorm.DeletedAt(ToNullTime(t))
}
