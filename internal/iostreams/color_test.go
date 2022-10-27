package iostreams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvColorFuncs(t *testing.T) {
	tests := []struct {
		desc          string
		NoColor       string
		CliColor      string
		CliColorForce string
		disabled      bool
		forced        bool
	}{
		{
			desc:          "colors are enabled when all env vars are unset",
			NoColor:       "",
			CliColor:      "",
			CliColorForce: "",
			disabled:      false,
			forced:        false,
		},

		{
			desc:          "colors are disabled whenever NO_COLOR has a value",
			NoColor:       "something",
			CliColor:      "",
			CliColorForce: "",
			disabled:      true,
			forced:        false,
		},

		{
			desc:          "colors are disabled whenever CLICOLOR is set to 0",
			NoColor:       "",
			CliColor:      "0",
			CliColorForce: "",
			disabled:      true,
			forced:        false,
		},

		{
			desc:          "colors are not disabled if CLICOLOR is set to anything other than 0",
			NoColor:       "",
			CliColor:      "something",
			CliColorForce: "",
			disabled:      false,
			forced:        false,
		},

		{
			desc:          "colors are forced when CLICOLOR_FORCE is non-zero",
			NoColor:       "",
			CliColor:      "",
			CliColorForce: "1",
			disabled:      false,
			forced:        true,
		},
		{
			desc:          "forcing colors takes priority over disabling",
			NoColor:       "1",
			CliColor:      "",
			CliColorForce: "1",
			disabled:      true,
			forced:        true,
		},
		{
			desc:          "colors are not forced when CLICOLOR_FORCE is set to 0",
			NoColor:       "",
			CliColor:      "",
			CliColorForce: "0",
			disabled:      false,
			forced:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			t.Setenv("NO_COLOR", tt.NoColor)
			t.Setenv("CLICOLOR", tt.CliColor)
			t.Setenv("CLICOLOR_FORCE", tt.CliColorForce)
			assert.Equal(t, tt.disabled, EnvColorDisabled())
			assert.Equal(t, tt.forced, EnvColorForced())
		})
	}
}
