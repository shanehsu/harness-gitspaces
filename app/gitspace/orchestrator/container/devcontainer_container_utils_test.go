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
