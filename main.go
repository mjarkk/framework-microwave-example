package main

import (
	microw "github.com/mjarkk/framework-microwave"
	T "github.com/mjarkk/framework-microwave/pkg/types"
)

func main() {
	microw.Init(&T.Init{
		MigrationPath: "./database/migration/",
	})
}
