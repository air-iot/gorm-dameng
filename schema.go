package dameng

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CurrentSchema(stmt *gorm.Statement, table string) (interface{}, interface{}) {
	if strings.Contains(table, ".") {
		if tables := strings.Split(table, `.`); len(tables) == 2 {
			return tables[0], tables[1]
		}
	}

	if stmt.TableExpr != nil {
		if tables := strings.Split(stmt.TableExpr.SQL, `"."`); len(tables) == 2 {
			return strings.TrimPrefix(tables[0], `"`), table
		}
	}
	return clause.Expr{SQL: "SF_GET_SCHEMA_NAME_BY_ID(CURRENT_SCHID)"}, table
}

func GetTableName(schema, tableName interface{}) string {
	return fmt.Sprintf("%v.%v", schema, tableName)
}

func GetIndexName(table interface{}, name string) string {
	return fmt.Sprintf("%v_%s", table, name)
}
