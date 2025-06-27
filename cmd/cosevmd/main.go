package main

import (
	"fmt"
	"os"

	clienthelpers "cosmossdk.io/client/v2/helpers"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	sdk "github.com/cosmos/cosmos-sdk/types"
	evmconfig "github.com/cosmos/evm/cmd/config"

	"github.com/green901612/cosevm/app"
	"github.com/green901612/cosevm/cmd/cosevmd/cmd"
)

func main() {
	// configuration
	config := sdk.GetConfig()
	cmd.SetBech32Prefixes(config)
	evmconfig.SetBip44CoinType(config)
	config.Seal()

	// cmd
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, clienthelpers.EnvPrefix, app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
