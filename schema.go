//go:build ignore

package main

import (
	"context"
	"database/sql"
	"embed"
	_ "embed"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//go:embed schemas/user.sql
// var ddlUser string

//go:embed schemas/store.sql
// var ddlStore string

//go:embed schemas/*.sql
var ddlFolder embed.FS

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	db, err := sql.Open("mysql", "root:root@/demo")
	if err != nil {
		return err
	}

	filenames := []string{
		"schemas/user.sql",
		"schemas/store.sql",
	}
	for _, fileName := range filenames {
		content, _ := ddlFolder.ReadFile(fileName)
		if _, err := db.ExecContext(ctx, string(content)); err != nil {
			return err
		}
	}

	return nil
}
