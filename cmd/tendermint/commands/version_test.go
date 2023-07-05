package commands

import (
	"testing"

	"github.com/PikeEcosystem/tendermint/version"
)

func TestVersionCmd(t *testing.T) {
	version.OCCoreSemVer = "test version"
	VersionCmd.Run(VersionCmd, nil)
}
