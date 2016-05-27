package cmd

import (
	"github.com/nii236/jmdict-toolkit/parse"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse JMDICT file into SQLite",
	Long:  "Parse JMDICT file into SQLite",
	Run: func(cmd *cobra.Command, args []string) {
		parse.Dictionary(viper.GetString("input"))
	},
}

func init() {
	RootCmd.AddCommand(parseCmd)
	parseCmd.Flags().StringP("input", "i", "data/JMdict_e.gz", "Path to the input JMDICT file")
	viper.BindPFlags(parseCmd.Flags())
}
