package cmd

import (
	"envinject/pkg"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

type flagEnvKey struct {
	flag, envKey string
}

var (
	rootPath  = flagEnvKey{envKey: "ROOT_PATH", flag: "rootPath"}

	injectCmd = &cobra.Command{
		Use:   "inject",
		Short: "inject envvars into files",
		Run:   envInject}
)

func init() {
	viper.AutomaticEnv()

	rootCmd.AddCommand(injectCmd)

	var rp string
	injectCmd.Flags().StringVarP(&rp, rootPath.flag, "r", "", "File containing service validation specification")

	viper.BindPFlag(rootPath.envKey, injectCmd.Flags().Lookup(rootPath.flag))
}

func envInject(cmd *cobra.Command, args []string) {
	err := pkg.Inject(viper.GetString(rootPath.envKey))

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

