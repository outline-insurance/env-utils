package cmd

import (
	"fmt"
	"os"

	"github.com/outline-insurance/env-utils/utils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// SecretsCmd provides commands related to populating env secrets
var SecretsCmd = &cobra.Command{
	Use:   "secret",
	Short: "commands related filling secrets",
	Long:  `commands related filling secrets`,
}

func init() {
	SecretsCmd.AddCommand(
		PopulateSecretsCMD,
	)
	var Region string
	var OutputFile string
	SecretsCmd.PersistentFlags().StringVarP(&Region, "region", "r", "us-east-1", "AWS region (required)")
	SecretsCmd.MarkFlagRequired("region")
	SecretsCmd.PersistentFlags().StringVarP(
		&OutputFile,
		"output-file",
		"o",
		".secret_env",
		"name and path to output file, defaults to current working directory",
	)
}

// PopulateSecretsCMD populates env secrets.
var PopulateSecretsCMD = &cobra.Command{

	Use:   "populate <filePath>",
	Short: "populates secrets stored in filepath from secrets manager",
	Long: `
	populates secrets stored in filepath from secrets manager
	`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {

		filePath := args[0]
		envMap, err := utils.ParseEnvFile(filePath)
		if err != nil {
			logrus.Fatal(errors.Wrapf(err, "while parsing env file %s", filePath))
		}
		region, err := cmd.Flags().GetString("region")
		if err != nil {
			logrus.Fatal(errors.Wrap(err, "while getting region flag value"))
		}
		persist, err := cmd.Flags().GetBool("persist")
		if err != nil {
			logrus.Fatal(errors.Wrap(err, "while getting persist flag value"))
		}
		localDevFormat, err := cmd.Flags().GetBool("local-dev-format")
		if err != nil {
			logrus.Fatal(errors.Wrap(err, "while getting local-dev-format flag value"))
		}
		persistString := ""
		for name, value := range *envMap {
			secret, err := utils.GetSecret(value, region)
			if err != nil {
				logrus.Fatal(errors.Wrapf(err, "while setting getting aws secrets manager value of secret %s", value))
			}
			err = os.Setenv(name, *secret)
			if err != nil {
				logrus.Fatal(errors.Wrapf(err, "while setting envirionment varable with name %s", name))
			}
			if persist {
				if localDevFormat {
					persistString = fmt.Sprintf("%s%s='%s'\n", persistString, name, *secret)
				} else {
					persistString = fmt.Sprintf("%sexport %s='%s'\n", persistString, name, *secret)
				}
			}
		}
		if persist {
			output, err := cmd.Flags().GetString("output-file")
			if err != nil {
				logrus.Fatal(errors.Wrap(err, "while getting output file flag value"))
			}
			secretFile, err := os.Create(output)
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
