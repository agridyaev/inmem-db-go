package in_memory

import (
	"context"
	"errors"
	"go.uber.org/zap"
)

type hashTable interface {
	Set(string, string)
	Get(string) (string, bool)
	Del(string)
}

type Engine struct {
	hashTable hashTable
	logger    *zap.Logger
}

func NewEngine(tableBuilder func() hashTable, logger *zap.Logger) (*Engine, error) {
	if tableBuilder == nil {
		return nil, errors.New("hash table builder is invalid")
	}

	if logger == nil {
		return nil, errors.New("logger is invalid")
	}

	return &Engine{
		hashTable: tableBuilder(),
		logger:    logger,
	}, nil
}

func (e *Engine) Set(ctx context.Context, key, value string) {
	e.hashTable.Set(key, value)

	txID := ctx.Value("tx").(int64)
	e.logger.Debug("success set query", zap.Int64("tx", txID))
}

func (e *Engine) Get(ctx context.Context, key string) (string, bool) {
	value, found := e.hashTable.Get(key)

	txID := ctx.Value("tx").(int64)
	e.logger.Debug("success get query", zap.Int64("tx", txID))
	return value, found
}

func (e *Engine) Del(ctx context.Context, key string) {
	e.hashTable.Del(key)

	txID := ctx.Value("tx").(int64)
	e.logger.Debug("success del query", zap.Int64("tx", txID))
}
