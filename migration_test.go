package migration

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMigration(t *testing.T) {
	m := New(&Config{
		AwsAccessKeyID:     "key",
		AwsSecretAccessKey: "secret",
		AwsRegion:          "ap-southeast-1",
		DynamoDBEndpoint:   "http://localhost:8000",
		Filename:           "./example/schema.yml",
	})

	err := m.Up()
	require.NoError(t, err)
}
