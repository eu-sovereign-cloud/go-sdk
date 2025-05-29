package client_test

import (
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/client"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	_, err := client.NewClient("http://localhost:13772/providers/seca.regions")
	require.NoError(t, err)
}
