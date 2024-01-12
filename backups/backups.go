package backups

import (
	"errors"

	"github.com/yendric/backup-manager/config"
)

var backups []Backup = []Backup{}

func New(name string, createScript string, backupDir string) Backup {
	return Backup{
		name,
		createScript,
		backupDir,
		[]Action{},
	}
}

func GetByName(name string) (Backup, error) {
	var backup Backup

	for _, backup := range backups {
		if backup.name == name {
			return backup, nil
		}
	}

	return backup, errors.New("backup '" + name + "' not found")
}

func init() {
	for _, backup := range config.Configuration.Backups {
		bcup := New(backup.Name, backup.CreateScript, backup.BackupDir)

		for _, action := range backup.Actions {
			bcup.AddAction(action.Name, action.Script)
		}
		backups = append(backups, bcup)
	}
}
