package utils

import "database/sql"

func ToSqlNullString(s string) sql.NullString {
	return sql.NullString{
		String: s, Valid: true,
	}
}

func ToSqlNullInt64(s int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: s, Valid: true,
	}
}
