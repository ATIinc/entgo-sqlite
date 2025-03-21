package main

import (
	"atiinc/entgo-sqlite/ent"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current directory: %v", err)
	}
	dbPath := filepath.Join(dir, "example.db")

	client, err := ent.Open("sqlite3", fmt.Sprintf("file:%s?mode=rwc&_fk=1", dbPath))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	err = client.Schema.Create(context.Background())

	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	examples, err := client.Example.Query().All(context.Background())
	if err != nil {
		log.Fatalf("failed querying examples: %v", err)
	}

	if len(examples) == 0 {
		exampleJSONMap := map[string]any{
			"stringKey": "stringValue",
			"intKey":    123,
			"boolKey":   true,
		}

		// create an example record
		_, err := client.Example.Create().
			SetExampleString("Hello!").
			SetExampleJSON(exampleJSONMap).
			Save(context.Background())

		if err != nil {
			log.Fatalf("failed creating example: %v", err)
		}

		log.Println("example record created")
	} else {
		log.Println("example record already exists")
	}

}
