package cmd

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/yendric/backup-manager/form"
)

var manageCmd = &cobra.Command{
	Use:     "manage",
	Aliases: []string{"list", "action"},
	Short:   "Shows a list of all backups and allows you to manage them",
	Run: func(cmd *cobra.Command, args []string) {
		backupFlag, _ := cmd.Flags().GetString("backup")
		fileFlag, _ := cmd.Flags().GetString("index")
		actionFlag, _ := cmd.Flags().GetString("action")

		// Select backup type
		backup, err := form.SelectBackup("Which backups would you like to list?", backupFlag)
		if err != nil {
			log.Fatalln(err)
		}

		// Select backup file
		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("Index", "Name", "Date", "Size")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		files := backup.Files()
		for index, file := range files {
			fileInfo, err := file.Info()
			if err != nil {
				fmt.Println(err)
			}

			tbl.AddRow(index, file.Name(), fileInfo.ModTime().Format("2006-01-02 15:04:05"), fileInfo.Size())
		}

		tbl.Print()

		fileIndex, err := form.SelectFileIndex("Which file would you like to manage?", len(files), fileFlag)
		if err != nil {
			log.Fatalln(err)
		}

		// Select backup action
		action, err := form.SelectAction("Which action would you like to run on this file?", backup, actionFlag)
		if err != nil {
			log.Fatalln(err)
		}

		// Run specified action
		err = action.Run(files[fileIndex])
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(backup.Name() + " " + action.Name() + " action on '" + files[fileIndex].Name() + "' ran successfully")

	},
}

func init() {
	rootCmd.AddCommand(manageCmd)

	manageCmd.Flags().StringP("backup", "b", "", "The backup type to list")
	manageCmd.Flags().StringP("index", "i", "", "The backup file index to select")
	manageCmd.Flags().StringP("action", "a", "", "The action to run on the selected file")
}
