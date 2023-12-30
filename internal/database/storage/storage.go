package storage

import (
	"context"
	"errors"
	"go.uber.org/zap"
)

type Engine interface {
	Set(context.Context, string, string)
	Get(context.Context, string) (string, bool)
	Del(context.Context, string)
}

type Storage struct {
	engine Engine
	logger *zap.Logger
}

func NewStorage(engine Engine, logger *zap.Logger) (*Storage, error) {
	if engine == nil {
		return nil, errors.New("engine is invalid")
	}

	if logger == nil {
		return nil, errors.New("logger is invalid")
	}

	return &Storage{
		engine: engine,
		logger: logger,
	}, nil
}

func (s *Storage) Set(ctx context.Context, key, value string) error {
	if ctx.Err() != nil {
		txID := ctx.Value("tx").(int64)
		s.logger.Debug("query canceled", zap.Int64("tx", txID))
		return ctx.Err()
	}

	s.engine.Set(ctx, key, value)
	return nil
}

func (s *Storage) Get(ctx context.Context, key string) (string, error) {
	if ctx.Err() != nil {
		txID := ctx.Value("tx").(int64)
		s.logger.Debug("query canceled", zap.Int64("tx", txID))
		return "", ctx.Err()
	}

	value, _ := s.engine.Get(ctx, key)
	return value, nil
}

func (s *Storage) Del(ctx context.Context, key string) error {
	if ctx.Err() != nil {
		txID := ctx.Value("tx").(int64)
		s.logger.Debug("query canceled", zap.Int64("tx", txID))
		return ctx.Err()
	}

	s.engine.Del(ctx, key)
	return nil
}
