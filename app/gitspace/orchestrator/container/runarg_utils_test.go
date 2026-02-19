package container

import (
	"testing"

	"github.com/harness/gitness/types"

	"github.com/stretchr/testify/require"
)

func TestGetNetworkMode(t *testing.T) {
	t.Parallel()

	runArgsMap := map[types.RunArg]*types.RunArgValue{
		types.RunArgNetwork: &types.RunArgValue{
			Name:   types.RunArgNetwork,
			Values: []string{"primary"},
		},
	}

	require.Equal(t, "primary", string(getNetworkMode(runArgsMap)))
}
