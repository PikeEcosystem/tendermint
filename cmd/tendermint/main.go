package main

import (
	"os"
	"path/filepath"

	cmd "github.com/PikeEcosystem/tendermint/cmd/tendermint/commands"
	"github.com/PikeEcosystem/tendermint/cmd/tendermint/commands/debug"
	cfg "github.com/PikeEcosystem/tendermint/config"
	"github.com/PikeEcosystem/tendermint/libs/cli"
	nm "github.com/PikeEcosystem/tendermint/node"
)

func main() {
	rootCmd := cmd.RootCmd
	rootCmd.AddCommand(
		cmd.GenValidatorCmd,
		cmd.ProbeUpnpCmd,
		cmd.LightCmd,
		cmd.ReplayCmd,
		cmd.ReplayConsoleCmd,
		cmd.ResetAllCmd,
		cmd.ResetPrivValidatorCmd,
		cmd.ResetStateCmd,
		cmd.ShowValidatorCmd,
		cmd.TestnetFilesCmd,
		cmd.ShowNodeIDCmd,
		cmd.GenNodeKeyCmd,
		cmd.VersionCmd,
		cmd.RollbackStateCmd,
		debug.DebugCmd,
		cli.NewCompletionCmd(rootCmd, true),
	)

	// NOTE:
	// Users wishing to:
	//	* Use an external signer for their validators
	//	* Supply an in-proc abci app
	//	* Supply a genesis doc file from another source
	//	* Provide their own DB implementation
	// can copy this file and use something other than the
	// DefaultNewNode function
	nodeFunc := nm.NewtendermintNode

	// Create & start node
	rootCmd.AddCommand(cmd.NewInitCmd())
	rootCmd.AddCommand(cmd.NewRunNodeCmd(nodeFunc))

	userHome, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	cmd := cli.PrepareBaseCmd(rootCmd, "OC", filepath.Join(userHome, cfg.DefaulttendermintDir))
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
