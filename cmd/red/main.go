package main

import (
	"os"

	"cosmossdk.io/log"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/jim380/Re/app"
)

func main() {
	rootCmd, _ := NewRootCmd()
	if err := svrcmd.Execute(rootCmd, EnvironmentPrefix, app.DefaultNodeHome); err != nil {
		log.NewLogger(rootCmd.OutOrStderr()).Error("failure when running app", "err", err)
		os.Exit(1)
	}
}
