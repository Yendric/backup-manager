package form

import (
	"errors"
	"log"
	"strconv"

	"github.com/manifoldco/promptui"
)

func SelectFileIndex(question string, maxIndex int, def string) (int, error) {
	// Check default value
	if def != "" {
		defInt, err := strconv.Atoi(def)
		if err != nil {
			log.Fatalln(err)
		}
		if defInt > maxIndex || defInt < 0 {
			log.Fatalln("invalid file index supplied")
		}
		return defInt, err

	}

	// Prompt value if default value is not supplied
	prompt := promptui.Prompt{
		Label: question,
		Validate: func(input string) error {
			inputInt, err := strconv.Atoi(input)
			if err != nil {
				return errors.New("invalid int")
			}
			if inputInt > maxIndex || inputInt < 0 {
				return errors.New("invalid index")
			}
			return nil
		},
	}

	result, err := prompt.Run()

	if err != nil {
		log.Fatalln(err)
		return 0, err
	}

	resultInt, err := strconv.Atoi(result)
	if err != nil {
		log.Fatalln(err)
	}

	return resultInt, nil
}
