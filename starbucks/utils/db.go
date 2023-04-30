package utils

import "database/sql"

func ToSqlNullString(s string) sql.NullString {
	return sql.NullString{
		String: s, Valid: true,
	}
}
