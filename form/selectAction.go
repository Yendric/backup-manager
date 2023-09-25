package form

import (
	"log"

	"github.com/manifoldco/promptui"
	"github.com/yendric/backup-manager/config"
	"github.com/yendric/backup-manager/utils"
)

func SelectAction(question string, backup config.Backup, def string) (config.Action, error) {
	// Check default value
	if def != "" {
		action, err := utils.GetActionByName(backup, def)
		if err != nil {
			log.Fatalln(err)
		}
		return action, err

	}

	// Prompt value if default value is not supplied

	var items []string
	var action config.Action

	for _, action := range backup.Actions {
		items = append(items, action.Name)
	}

	prompt := promptui.Select{
		Label: question,
		Items: items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		return action, err
	}

	action, err = utils.GetActionByName(backup, result)
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		return action, err
	}

	return action, nil
}
