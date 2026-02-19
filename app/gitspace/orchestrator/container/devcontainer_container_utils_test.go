package container

import (
	"testing"

	"github.com/harness/gitness/types"

	"github.com/stretchr/testify/require"
)

func TestGetAdditionalDockerNetworks(t *testing.T) {
	t.Parallel()

	devcontainerConfig := types.DevcontainerConfig{
		Customizations: types.DevContainerConfigCustomizations{
			"shanehsu": map[string]any{
				"dockerNetworks": []string{" traefik ", "", "traefik", "primary", "shared"},
			},
		},
	}

	networks := getAdditionalDockerNetworks(devcontainerConfig)
	require.Equal(t, []string{" traefik ", "", "traefik", "primary", "shared"}, networks)
}

func TestGetAdditionalDockerNetworksNoCustomization(t *testing.T) {
	t.Parallel()

	devcontainerConfig := types.DevcontainerConfig{}

	networks := getAdditionalDockerNetworks(devcontainerConfig)
	require.Empty(t, networks)
}

func TestGetAdditionalDockerNetworksEmptyArray(t *testing.T) {
	t.Parallel()

	devcontainerConfig := types.DevcontainerConfig{
		Customizations: types.DevContainerConfigCustomizations{
			"shanehsu": map[string]any{
				"dockerNetworks": []string{},
			},
		},
	}

	networks := getAdditionalDockerNetworks(devcontainerConfig)
	require.Empty(t, networks)
}
