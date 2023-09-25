/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/yendric/backup-manager/config"
	"github.com/yendric/backup-manager/form"
	"github.com/yendric/backup-manager/utils"
)

var backupCmd = &cobra.Command{
	Use:     "backup",
	Aliases: []string{"create", "new", "make", "run"},
	Short:   "Creates a new backup",
	Run: func(cmd *cobra.Command, args []string) {
		var backup config.Backup

		backupFlag, _ := cmd.Flags().GetString("backup")

		backup, err := form.SelectBackup("What would you like to backup?", backupFlag)
		if err != nil {
			log.Fatalln(err)
		}

		// Run action
		err = utils.CreateBackup(backup)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(backup.Name + " backup created successfully")
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)

	backupCmd.Flags().StringP("backup", "b", "", "The backup type to create")
}
