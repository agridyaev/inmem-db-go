package database

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"inmem-db-go/internal/database/compute"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	ctrl := gomock.NewController(t)
	computeLayer := NewMockcomputeLayer(ctrl)
	storageLayer := NewMockstorageLayer(ctrl)

	database, err := NewDatabase(nil, nil, nil)
	require.Error(t, err, "compute is invalid")
	require.Nil(t, database)

	database, err = NewDatabase(computeLayer, nil, nil)
	require.Error(t, err, "storage is invalid")
	require.Nil(t, database)

	database, err = NewDatabase(computeLayer, storageLayer, nil)
	require.Error(t, err, "logger is invalid")
	require.Nil(t, database)

	database, err = NewDatabase(computeLayer, storageLayer, zap.NewNop())
	require.NoError(t, err)
	require.NotNil(t, database)
}

func TestHandleSetQuery(t *testing.T) {
	ctx := context.WithValue(context.Background(), "tx", int64(555))

	ctrl := gomock.NewController(t)
	computeLayer := NewMockcomputeLayer(ctrl)
	computeLayer.EXPECT().
		HandleQuery(ctx, "SET one 1").
		Return(compute.NewQuery(compute.SetCommandID, []string{"one", "1"}), nil)

	storageLayer := NewMockstorageLayer(ctrl)
	storageLayer.EXPECT().
		Set(ctx, "one", "1").
		Return(nil)

	database, err := NewDatabase(computeLayer, storageLayer, zap.NewNop())
	require.NoError(t, err)
	require.NotNil(t, database)

	res := database.HandleQuery(ctx, "SET one 1")
	assert.Equal(t, res, "[ok]")
}

func TestHandleGetQuery(t *testing.T) {
	ctx := context.WithValue(context.Background(), "tx", int64(555))

	ctrl := gomock.NewController(t)
	computeLayer := NewMockcomputeLayer(ctrl)
	computeLayer.EXPECT().
		HandleQuery(ctx, "GET one").
		Return(compute.NewQuery(compute.GetCommandID, []string{"one"}), nil)

	storageLayer := NewMockstorageLayer(ctrl)
	storageLayer.EXPECT().
		Get(ctx, "one").
		Return("1", nil)

	database, err := NewDatabase(computeLayer, storageLayer, zap.NewNop())
	require.NoError(t, err)
	require.NotNil(t, database)

	res := database.HandleQuery(ctx, "GET one")
	assert.Equal(t, res, "[ok] 1")
}

func TestHandleDelQuery(t *testing.T) {
	ctx := context.WithValue(context.Background(), "tx", int64(555))

	ctrl := gomock.NewController(t)
	computeLayer := NewMockcomputeLayer(ctrl)
	computeLayer.EXPECT().
		HandleQuery(ctx, "DEL one").
		Return(compute.NewQuery(compute.DelCommandID, []string{"one"}), nil)

	storageLayer := NewMockstorageLayer(ctrl)
	storageLayer.EXPECT().
		Del(ctx, "one").
		Return(nil)

	database, err := NewDatabase(computeLayer, storageLayer, zap.NewNop())
	require.NoError(t, err)
	require.NotNil(t, database)

	res := database.HandleQuery(ctx, "DEL one")
	assert.Equal(t, res, "[ok]")
}
