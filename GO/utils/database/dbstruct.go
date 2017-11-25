package database

import "database/sql"

type NullStr struct {
	Nullstr sql.NullString
}

type NullInt struct {
	Nullint64 sql.NullInt64
}

type NullBool struct {
	Nullbool sql.NullBool
}

