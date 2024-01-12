package form

import (
	"log"

	"github.com/manifoldco/promptui"
	"github.com/yendric/backup-manager/backups"
	"github.com/yendric/backup-manager/config"
)

func SelectBackup(question string, def string) (backups.Backup, error) {
	// Check default value
	if def != "" {
		bcup, err := backups.GetByName(def)
		if err != nil {
			return backups.Backup{}, err
		}
		return bcup, err

	}

	// Prompt value if default value is not supplied
	var items []string

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
		return backups.Backup{}, err
	}

	backup, err := backups.GetByName(result)
	if err != nil {
		log.Fatalln(err)
		return backup, err
	}

	return backup, nil
}
