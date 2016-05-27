package cmd

import (
	"github.com/nii236/jmdict-toolkit/fetch"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetches the JMDICT file from the internet",
	Long:  "Fetches the JMDICT file from the internet",
	Run: func(cmd *cobra.Command, args []string) {
		f := &fetch.Fetcher{}
		fc := &fetch.FileCreator{}
		fetch.Dictionary(viper.GetString("url"), viper.GetString("outfile"), f, fc)
	},
}

func init() {
	RootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().StringP("url", "u", "ftp://ftp.monash.edu.au/pub/nihongo/JMdict_e.gz", "HTTP path to the JMDICT file")
	fetchCmd.Flags().StringP("outfile", "o", "data/JMdict_e.gz", "Location to save the dictionary")
	viper.BindPFlags(fetchCmd.Flags())
}
