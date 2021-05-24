package cmd

import (
	nonsecret "github.com/outline-insurance/env-utils/cmd/non-secret"
	secret "github.com/outline-insurance/env-utils/cmd/secret"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "env-utils",
	Short: "Stuff for Pathpoint Environments",
	Long:  `Utility code for the pathpoint environment repository.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func init() {
	// Add sub-command pacakges
	rootCmd.AddCommand(
		secret.SecretsCmd,
		nonsecret.NonSecretsCmd,
	)

	var Persist bool
	var LocalDevFormat bool
	rootCmd.PersistentFlags().BoolVarP(&Persist, "persist", "p", true, "persist secrets or variables to an env file?")
	rootCmd.PersistentFlags().BoolVarP(&LocalDevFormat, "local-dev-format", "l", false, "format env file to work for local dev?")
}
