package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "backup-manager",
	Short: "Backup manager allows you to manage backups with custom scripts.",
	Long: `Backup manager allows you to create backups with custom scripts
and manage them with an easy to use CLI.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
