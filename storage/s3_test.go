package storage_test

import (
	"testing"

	"github.com/mitchellh/goamz/aws"
	"github.com/stretchr/testify/require"
	"github.com/thomaso-mirodin/shawty/storage"
)

func setupS3Storage(t testing.TB) storage.Storage {
	auth, err := aws.SharedAuth()
	require.Nil(t, err)

	s, err := storage.NewS3(auth, aws.USWest2, "databricks-go-shortener-test")
	require.Nil(t, err)

	return s
}

func cleanupS3Storage() error {
	auth, err := aws.SharedAuth()
	if err != nil {
		return err
	}

	s, err := storage.NewS3(auth, aws.USWest2, "databricks-go-shortener-test")
	if err != nil {
		return err
	}

	bc, err := s.Bucket.GetBucketContents()
	for k, _ := range *bc {
		s.Bucket.Del(k)
	}

	// if err := s.Bucket.DelBucket(); err != nil {
	// 	return err
	// }

	return nil
}
