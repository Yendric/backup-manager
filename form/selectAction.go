package form

import (
	"log"

	"github.com/manifoldco/promptui"
	"github.com/yendric/backup-manager/backups"
)

func SelectAction(question string, bcup backups.Backup, def string) (backups.Action, error) {
	// Check default value
	if def != "" {
		action, err := bcup.Action(def)
		if err != nil {
			log.Fatalln(err)
		}
		return action, err

	}

	// Prompt value if default value is not supplied
	var items []string

	for _, action := range bcup.Actions() {
		items = append(items, action.Name())
	}

	prompt := promptui.Select{
		Label: question,
		Items: items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		return backups.Action{}, err
	}

	action, err := bcup.Action(result)
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		return backups.Action{}, err
	}

	return action, nil
}
