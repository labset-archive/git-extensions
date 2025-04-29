package workspace_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
	"github.com/viqueen/git-extensions/internal/workspace"
)

func TestRepoFromUrlConnection(t *testing.T) {
	t.Parallel()

	var sshUrl = "ssh://git@github.com:viqueen/codenav.git"

	var httpsUrl = "https://github.com/viqueen/devbox.git"

	tests := map[string]struct {
		inputUrlConnection string
		inputWorkspace     string
		expectedOutput     *workspace.Repo
		expectedError      error
	}{
		"ssh": {
			inputUrlConnection: sshUrl,
			inputWorkspace:     "test",
			expectedOutput: &workspace.Repo{
				UrlConnection: sshUrl,
				Workspace:     "test",
				Host:          "github.com",
				Owner:         "viqueen",
				Name:          "codenav",
			},
		},
		"https": {
			inputUrlConnection: httpsUrl,
			inputWorkspace:     "test",
			expectedOutput: &workspace.Repo{
				UrlConnection: httpsUrl,
				Workspace:     "test",
				Host:          "github.com",
				Owner:         "viqueen",
				Name:          "devbox",
			},
		},
		"invalid": {
			inputUrlConnection: "invalid-url",
			inputWorkspace:     "test",
			expectedOutput:     nil,
			expectedError:      workspace.ErrInvalidUrlConnection,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			repo, err := workspace.RepoFromUrlConnection(
				test.inputUrlConnection,
				test.inputWorkspace,
			)
			if test.expectedError != nil {
				assert.Nil(t, repo)
				assert.Equal(t, test.expectedError.Error(), err.Error())
			} else {
				assert.Equal(t, test.expectedOutput, repo)
				require.NoError(t, err)
			}
		})
	}
}
