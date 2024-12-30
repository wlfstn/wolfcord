package wc

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool
var DbConn string

func initDatabase(databaseUrl string) {
	ctx := context.Background()
	var err error
	Pool, err = pgxpool.New(ctx, databaseUrl)
	if err != nil {
		fmt.Printf("Unable to create connection pool: %v\n", err)
	}
}

func EndPGX() {
	Pool.Close()
}

// variable input with variable output
func RunQuery(file string, args ...interface{}) [][]interface{} {
	query := ResourceLoadSQL(file)
	fmt.Printf("args: %v\n", args...)

	var results [][]interface{} = make([][]interface{}, 0)
	rows, err := Pool.Query(context.Background(), query, args...)
	if err != nil {
		fmt.Println("Error querying database: ", err)
		return [][]interface{}{}
	}
	defer rows.Close()

	colCount := len(rows.FieldDescriptions())
	if colCount == 0 {
		fmt.Println("No columns returned")
		return [][]interface{}{}
	}
	fmt.Println("Columns: ", colCount)

	for rows.Next() {
		rowData := make([]interface{}, colCount)
		columnPointers := make([]interface{}, colCount)

		for i := range rowData {
			columnPointers[i] = &rowData[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			fmt.Println("Error scanning rows: ", err)
			return [][]interface{}{}
		}

		results = append(results, rowData)
	}

	return results
}

func ResourceLoadSQL(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}
	defer file.Close()

	query, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}
	return string(query)
}
