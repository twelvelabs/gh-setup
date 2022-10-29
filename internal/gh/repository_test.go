package gh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_RemoteURL(t *testing.T) {
	tests := []struct {
		desc      string
		repo      *Repository
		responses map[Protocol]string
	}{
		{
			desc: "returns the correct remote for the given protocol",
			repo: &Repository{
				CloneURL: "https://github.com",
				GitURL:   "git://github.com",
				SSHURL:   "git@github.com",
			},
			responses: map[Protocol]string{
				ProtocolHTTPS:  "https://github.com",
				ProtocolGit:    "git://github.com",
				ProtocolSSH:    "git@github.com",
				"non-protocol": "https://github.com",
				"":             "https://github.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			for protocol, expected := range tt.responses {
				actual := tt.repo.RemoteURL(protocol)
				assert.Equal(t, expected, actual)
			}
		})
	}
}
