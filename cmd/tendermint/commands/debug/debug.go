package debug

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/PikeEcosystem/tendermint/libs/log"
)

var (
	nodeRPCAddr string
	profAddr    string
	frequency   uint

	flagNodeRPCAddr = "rpc-laddr"
	flagProfAddr    = "pprof-laddr"
	flagFrequency   = "frequency"

	logger = log.NewOCLogger(log.NewSyncWriter(os.Stdout))
)

// DebugCmd defines the root command containing subcommands that assist in
// debugging running tendermint processes.
var DebugCmd = &cobra.Command{
	Use:   "debug",
	Short: "A utility to kill or watch an tendermint process while aggregating debugging data",
}

func init() {
	DebugCmd.PersistentFlags().SortFlags = true
	DebugCmd.PersistentFlags().StringVar(
		&nodeRPCAddr,
		flagNodeRPCAddr,
		"tcp://localhost:26657",
		"the tendermint node's RPC address (<host>:<port>)",
	)

	DebugCmd.AddCommand(killCmd)
	DebugCmd.AddCommand(dumpCmd)
}
