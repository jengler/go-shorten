package storage_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thomaso-mirodin/shawty/storage"
)

func setupInmemStorage(t testing.TB) storage.Storage {
	s, err := storage.NewInmem()
	require.Nil(t, err)

	return s
}
