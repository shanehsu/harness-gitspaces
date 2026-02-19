package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractShanehsuDockerNetworks(t *testing.T) {
	t.Parallel()

	customizations := DevContainerConfigCustomizations{
		"shanehsu": map[string]any{
			"dockerNetworks": []string{"traefik", "shared"},
		},
	}

	require.Equal(t, []string{"traefik", "shared"}, customizations.ExtractShanehsuDockerNetworks())
}

func TestExtractShanehsuDockerNetworksMissingCustomization(t *testing.T) {
	t.Parallel()

	customizations := DevContainerConfigCustomizations{}

	require.Nil(t, customizations.ExtractShanehsuDockerNetworks())
}
