package commands

import (
	"fmt"

	"github.com/PikeEcosystem/tendermint/node"
	"github.com/PikeEcosystem/tendermint/types"
	"github.com/spf13/cobra"

	cfg "github.com/PikeEcosystem/tendermint/config"
	tmjson "github.com/PikeEcosystem/tendermint/libs/json"
	tmos "github.com/PikeEcosystem/tendermint/libs/os"
	"github.com/PikeEcosystem/tendermint/privval"
)

// ShowValidatorCmd adds capabilities for showing the validator info.
var ShowValidatorCmd = &cobra.Command{
	Use:     "show-validator",
	Aliases: []string{"show_validator"},
	Short:   "Show this node's validator info",
	RunE: func(cmd *cobra.Command, args []string) error {
		return showValidator(cmd, args, config)
	},
	PreRun: deprecateSnakeCase,
}

func showValidator(cmd *cobra.Command, args []string, config *cfg.Config) error {
	var pv types.PrivValidator
	if config.PrivValidatorListenAddr != "" {
		chainID, err := loadChainID(config)
		if err != nil {
			return err
		}
		pv, err = node.CreateAndStartPrivValidatorSocketClient(config.PrivValidatorListenAddr, chainID, logger)
		if err != nil {
			return err
		}
	} else {
		keyFilePath := config.PrivValidatorKeyFile()
		if !tmos.FileExists(keyFilePath) {
			return fmt.Errorf("private validator file %s does not exist", keyFilePath)
		}
		pv = privval.LoadFilePV(keyFilePath, config.PrivValidatorStateFile())
	}

	pubKey, err := pv.GetPubKey()
	if err != nil {
		return fmt.Errorf("can't get pubkey: %w", err)
	}

	bz, err := tmjson.Marshal(pubKey)
	if err != nil {
		return fmt.Errorf("failed to marshal private validator pubkey: %w", err)
	}

	fmt.Println(string(bz))
	return nil
}

func loadChainID(config *cfg.Config) (string, error) {
	stateDB, err := node.DefaultDBProvider(&node.DBContext{ID: "state", Config: config})
	if err != nil {
		return "", err
	}
	defer func() {
		var _ = stateDB.Close()
	}()
	genesisDocProvider := node.DefaultGenesisDocProviderFunc(config)
	_, genDoc, err := node.LoadStateFromDBOrGenesisDocProvider(stateDB, genesisDocProvider)
	if err != nil {
		return "", err
	}
	return genDoc.ChainID, nil
}
