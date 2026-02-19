package orchestrator

import (
	"testing"

	"github.com/harness/gitness/app/gitspace/orchestrator/container/response"
	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/enum"

	"github.com/stretchr/testify/require"
)

func TestRenderExternalIDEURLVSCodeWebAddsFolder(t *testing.T) {
	t.Parallel()

	ideURL, err := renderExternalIDEURL(
		"https://{{.ContainerName}}.dev.example.com",
		externalIDEURLTemplateData{ContainerName: "gitspace-test"},
		enum.IDETypeVSCodeWeb,
		"/home/harness/repo",
	)

	require.NoError(t, err)
	require.Equal(t, "https://gitspace-test.dev.example.com?folder=home%2Fharness%2Frepo", ideURL)
}

func TestRenderExternalIDEURLNonBrowserIDEDoesNotAddFolder(t *testing.T) {
	t.Parallel()

	ideURL, err := renderExternalIDEURL(
		"https://{{.ContainerName}}.dev.example.com",
		externalIDEURLTemplateData{ContainerName: "gitspace-test"},
		enum.IDETypeVSCode,
		"/home/harness/repo",
	)

	require.NoError(t, err)
	require.Equal(t, "https://gitspace-test.dev.example.com", ideURL)
}

func TestBuildExternalIDEURLTemplateDataUsesGeneratedContainerName(t *testing.T) {
	t.Parallel()

	gitspaceConfig := types.GitspaceConfig{
		Identifier: "sample",
		GitspaceUser: types.GitspaceUser{
			Identifier: "john",
		},
	}

	templateData := buildExternalIDEURLTemplateData(
		gitspaceConfig,
		types.Infrastructure{},
		&response.StartResponse{},
	)

	require.Equal(t, "gitspace-john-sample", templateData.ContainerName)
}
