package main

import (
	"bufio"
	"context"
	"fmt"
	"go.uber.org/zap"
	"inmem-db-go/internal/database"
	"inmem-db-go/internal/database/compute"
	"inmem-db-go/internal/database/storage"
	"inmem-db-go/internal/database/storage/engine/in_memory"
	"os"
)

func main() {
	logger := zap.NewNop()

	parser, err := compute.NewParser(logger)
	if err != nil {
		logger.Error(err.Error())
	}

	analyzer, err := compute.NewAnalyzer(logger)
	if err != nil {
		logger.Error(err.Error())
	}

	comp, err := compute.NewCompute(parser, analyzer, logger)
	if err != nil {
		logger.Error(err.Error())
	}

	engine, err := in_memory.NewEngine(in_memory.HashTableBuilder, logger)
	if err != nil {
		logger.Error(err.Error())
	}

	store, err := storage.NewStorage(engine, logger)
	if err != nil {
		logger.Error(err.Error())
	}

	database, err := database.NewDatabase(comp, store, logger)
	if err != nil {
		logger.Error(err.Error())
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		request, err := reader.ReadString('\n')
		if err != nil {
			logger.Error(err.Error())
			continue
		}

		result := database.HandleQuery(context.Background(), request)
		fmt.Printf("%s\n", result)
	}
}
