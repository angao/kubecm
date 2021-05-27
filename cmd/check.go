package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
)

// CheckCommand check cmd struct
type CheckCommand struct {
	BaseCommand
}

// Init CheckCommand
func (lc *CheckCommand) Init() {
	lc.command = &cobra.Command{
		Use:     "check",
		Short:   "Check ClusterStatus",
		Long:    "Check ClusterStatus",
		Aliases: []string{"ck"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return lc.runCheck(cmd, args)
		},
		Example: checkExample(),
	}
	lc.command.DisableFlagsInUseLine = true
}

func (lc *CheckCommand) runCheck(command *cobra.Command, args []string) error {
	config, err := clientcmd.LoadFromFile(cfgFile)
	if err != nil {
		return err
	}
	config = CheckValidContext(false, config)

	if err = ClusterStatus(); err != nil {
		printWarning(os.Stdout, "Cluster check failure!\n")
		return err
	}
	return nil
}

func checkExample() string {
	return `
# List all the contexts in your KubeConfig file
kubecm check
# Alias
kubecm c
# Filter out keywords(Multi-keyword support)
kubecm check kind
`
}
