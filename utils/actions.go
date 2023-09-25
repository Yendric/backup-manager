package utils

import (
	"errors"
	"io/fs"
	"os/exec"
	"strings"

	"github.com/yendric/backup-manager/config"
	"github.com/yendric/backup-manager/notifications"
)

func GetActionByName(backup config.Backup, name string) (config.Action, error) {
	var action config.Action

	for _, action := range backup.Actions {
		if action.Name == name {
			return action, nil
		}
	}

	return action, errors.New("action '" + name + "' not found")
}

func RunActionWithFile(backup config.Backup, action config.Action, file fs.DirEntry) error {
	scriptArgs := strings.Fields(action.Script)
	scriptArgs = append(scriptArgs, file.Name())
	out, commandErr := exec.Command(scriptArgs[0], scriptArgs[1:]...).Output()

	notifErr := notifications.SendNotifications(backup.Name+" "+file.Name(), action.Name, commandErr == nil, string(out))

	if commandErr != nil {
		return commandErr
	}
	if notifErr != nil {
		return notifErr
	}

	return nil
}
