package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "env-utils",
	Short: "The Pathpoint eXchange Sonic Screwdriver",
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
	// rootCmd.AddCommand(
	// 	lambdoos.LambdoosCmd,
	// 	reviewapps.ReviewAppCmd,
	// 	util.UtilAppCmd,
	// 	serve.ServeCmd,
	// 	coverletter.CoverLetterCmd,
	// 	eavcmd.EAVCmdRoot,
	// 	hitbox.HitboxCmdRoot,
	// 	docgen.GdocCmd,
	// 	docgen.QuoteDocCmd,
	// )

}
