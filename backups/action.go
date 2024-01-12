package backups

import (
	"io/fs"
	"os/exec"
	"strings"

	"github.com/yendric/backup-manager/notifications"
)

type Action struct {
	name   string
	script string
	backup *Backup
}

func (action *Action) Name() string {
	return action.name
}

func (action *Action) Run(file fs.DirEntry) error {
	scriptArgs := strings.Fields(action.script)
	scriptArgs = append(scriptArgs, file.Name())
	out, commandErr := exec.Command(scriptArgs[0], scriptArgs[1:]...).Output()

	notifErr := notifications.SendNotifications(action.backup.Name()+" "+file.Name(), action.name, commandErr == nil, string(out))

	if commandErr != nil {
		return commandErr
	}
	if notifErr != nil {
		return notifErr
	}

	return nil
}
