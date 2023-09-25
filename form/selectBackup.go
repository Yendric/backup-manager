package form

import (
	"log"

	"github.com/manifoldco/promptui"
	"github.com/yendric/backup-manager/config"
	"github.com/yendric/backup-manager/utils"
)

func SelectBackup(question string, def string) (config.Backup, error) {
	// Check default value
	if def != "" {
		backup, err := utils.GetBackupByName(def)
		if err != nil {
			log.Fatalln(err)
		}
		return backup, err

	}

	// Prompt value if default value is not supplied
	var items []string
	var backup config.Backup

	for _, backup := range config.Configuration.Backups {
		items = append(items, backup.Name)
	}

	prompt := promptui.Select{
		Label: question,
		Items: items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		log.Fatalln(err)
		return backup, err
	}

	backup, err = utils.GetBackupByName(result)
	if err != nil {
		log.Fatalln(err)
		return backup, err
	}

	return backup, nil
}
