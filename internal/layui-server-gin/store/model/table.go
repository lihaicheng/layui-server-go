package model

import "fmt"

const (
	TableNamePrefix = "gin"
)

func tableName(s string) string {
	return fmt.Sprintf("%s_%s", TableNamePrefix, s)
}

// Table is used for auto migrate
type Table interface {
	TableName() string
}
