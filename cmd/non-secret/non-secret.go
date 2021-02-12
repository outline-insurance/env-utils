package cmd

import (
	"fmt"
	"os"

	"github.com/outline-insurance/env-utils/utils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NonSecretsCmd provides commands related to populating env vars
var NonSecretsCmd = &cobra.Command{
	Use:   "non-secret",
	Short: "commands related filling regular env vars",
	Long:  `commands related filling regular env vars`,
}

func init() {
	NonSecretsCmd.AddCommand(
		PopulateNonSecretsCMD,
	)
}

// PopulateNonSecretsCMD populates env secrets.
var PopulateNonSecretsCMD = &cobra.Command{

	Use:   "populate <filePath>",
	Short: "populates env vars stored in filepath",
	Long: `
	populates env vars stored in filepath
	`,
	Run: func(cmd *cobra.Command, args []string) {

		filePath := args[0]
		envMap, err := utils.ParseEnvFile(filePath)
		if err != nil {
			logrus.Fatal(errors.Wrapf(err, "while parsing env file %s", filePath))
		}

		persist, err := cmd.Flags().GetBool("persist")
		persistString := ""
		for name, value := range *envMap {
			err = os.Setenv(name, value)
			if err != nil {
				logrus.Fatal(errors.Wrapf(err, "while setting envirionment varable with name %s", name))
			}
			if persist {
				persistString = fmt.Sprintf("%sexport %s=%s\n", persistString, name, value)
			}
		}
		if persist {
			secretFile, err := os.Create(".non_secret_env")
			if err != nil {
				logrus.Fatal(errors.Wrap(err, "while creating secret env persist file"))
			}
			_, err = secretFile.WriteString(persistString)

			if err != nil {
				logrus.Fatal(errors.Wrapf(err, "while writing to secret env persist file"))
			}
		}

	},
	Args: cobra.ExactArgs(1),
}
