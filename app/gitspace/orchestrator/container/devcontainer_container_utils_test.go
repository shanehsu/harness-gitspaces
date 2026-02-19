package container

import (
	"testing"

	"github.com/harness/gitness/types"

	dockertypes "github.com/docker/docker/api/types/container"
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

	networks := getAdditionalDockerNetworks(dockertypes.NetworkMode("primary"), devcontainerConfig)
	require.Equal(t, []string{"traefik", "shared"}, networks)
}

func TestGetAdditionalDockerNetworksNoCustomization(t *testing.T) {
	t.Parallel()

	devcontainerConfig := types.DevcontainerConfig{}

	networks := getAdditionalDockerNetworks(dockertypes.NetworkMode("primary"), devcontainerConfig)
	require.Nil(t, networks)
}

func TestGetAdditionalDockerNetworksAllFiltered(t *testing.T) {
	t.Parallel()

	devcontainerConfig := types.DevcontainerConfig{
		Customizations: types.DevContainerConfigCustomizations{
			"shanehsu": map[string]any{
				"dockerNetworks": []string{"", " ", "primary", "primary"},
			},
		},
	}

	networks := getAdditionalDockerNetworks(dockertypes.NetworkMode("primary"), devcontainerConfig)
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

	networks := getAdditionalDockerNetworks(dockertypes.NetworkMode("primary"), devcontainerConfig)
	require.Nil(t, networks)
}
