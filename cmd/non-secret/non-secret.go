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
	var OutputFile string
	NonSecretsCmd.PersistentFlags().StringVarP(
		&OutputFile,
		"output-file",
		"o",
		".non_secret_env",
		"name and path to output file, defaults to current working directory",
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
		if err != nil {
			logrus.Fatal(errors.Wrap(err, "while getting persist flag value"))
		}
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
			output, err := cmd.Flags().GetString("output-file")
			if err != nil {
				logrus.Fatal(errors.Wrap(err, "while getting output file flag value"))
			}
			nonSecretFile, err := os.Create(output)
			if err != nil {
				logrus.Fatal(errors.Wrap(err, "while creating env persist file"))
			}
			_, err = nonSecretFile.WriteString(persistString)

			if err != nil {
				logrus.Fatal(errors.Wrapf(err, "while writing to env persist file"))
			}
		}

	},
	Args: cobra.ExactArgs(1),
}
