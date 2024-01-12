package backups

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"strings"

	"github.com/yendric/backup-manager/notifications"
)

type Backup struct {
	name         string
	createScript string
	backupDir    string
	actions      []Action
}

func (backup *Backup) Name() string {
	return backup.name
}

func (backup *Backup) Run() error {
	scriptArgs := strings.Fields(backup.createScript)
	out, commandErr := exec.Command(scriptArgs[0], scriptArgs[1:]...).Output()

	notifErr := notifications.SendNotifications(backup.name, "create", commandErr == nil, string(out))

	if commandErr != nil {
		return commandErr
	}
	if notifErr != nil {
		return notifErr
	}

	return nil
}

func (backup *Backup) Action(name string) (Action, error) {
	for _, action := range backup.actions {
		if action.Name() == name {
			return action, nil
		}
	}

	return Action{}, errors.New("action '" + name + "' not found")
}

func (backup *Backup) Actions() []Action {
	return backup.actions
}

func (backup *Backup) RunAction(name string, dir fs.DirEntry) error {
	action, err := backup.Action(name)
	if err != nil {
		return err
	}

	return action.Run(dir)
}

func (backup *Backup) Files() []fs.DirEntry {
	files, err := os.ReadDir(backup.backupDir)
	if err != nil {
		fmt.Println(err)
	}

	return files
}

func (backup *Backup) AddAction(name string, script string) {
	backup.actions = append(backup.actions, Action{name, script, backup})
}
