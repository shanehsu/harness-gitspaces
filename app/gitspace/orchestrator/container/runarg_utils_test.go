package container

import (
	"testing"

	"github.com/harness/gitness/types"

	"github.com/stretchr/testify/require"
)

func TestGetNetworks(t *testing.T) {
	t.Parallel()

	runArgsMap := map[types.RunArg]*types.RunArgValue{
		types.RunArgNetwork: &types.RunArgValue{
			Name:   types.RunArgNetwork,
			Values: []string{"primary", "", "secondary"},
		},
	}

	require.Equal(t, []string{"primary", "secondary"}, getNetworks(runArgsMap))
}

func TestGetNetworkModeUsesFirstNetwork(t *testing.T) {
	t.Parallel()

	runArgsMap := map[types.RunArg]*types.RunArgValue{
		types.RunArgNetwork: &types.RunArgValue{
			Name:   types.RunArgNetwork,
			Values: []string{"primary", "secondary"},
		},
	}

	require.Equal(t, "primary", string(getNetworkMode(runArgsMap)))
}
