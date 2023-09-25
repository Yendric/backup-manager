package utils

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/yendric/backup-manager/config"
	"github.com/yendric/backup-manager/notifications"
)

func GetBackupByName(name string) (config.Backup, error) {
	var backup config.Backup

	for _, backup := range config.Configuration.Backups {
		if backup.Name == name {
			return backup, nil
		}
	}

	return backup, errors.New("backup '" + name + "' not found")
}

func CreateBackup(backup config.Backup) error {
	scriptArgs := strings.Fields(backup.CreateScript)
	out, commandErr := exec.Command(scriptArgs[0], scriptArgs[1:]...).Output()

	notifErr := notifications.SendNotifications(backup.Name, "create", commandErr == nil, string(out))

	if commandErr != nil {
		return commandErr
	}
	if notifErr != nil {
		return notifErr
	}

	return nil
}
