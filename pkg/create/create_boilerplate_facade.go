package create

import (
	"encoding/json"
	"errors"

	"github.com/Kshitij-Dhakal/inaugurator/common"
)

type BoilerplateCreator interface {
	CreateBoilerplate(file string, args ...string) (string, error)
}

type BoilerplateCreatorImpl struct {
	Service BoilderplaterCreatorService
}

func (c BoilerplateCreatorImpl) CreateBoilerplate(file string, args ...string) (string, error) {
	if file == "" {
		return "", errors.New("file name is required")
	}
	data, err := common.ReadFile(file)
	if err != nil {
		return "", err
	}
	var commandList CommandList
	x := []byte(data)
	err = json.Unmarshal(x, &commandList)
	if err != nil {
		return "", err
	}
	if commandList.GetName() == "" {
		return "", errors.New("command name is required")
	}
	for _, v := range commandList.Commands {
		err = validateCommand(v)
		if err != nil {
			return "", err
		}
		//only one layer of subcommands
		for _, w := range v.Subcommands {
			err = validateCommand(w)
			if err != nil {
				return "", err
			}
		}
	}

	return c.Service.CreateBoilerplate(file, &commandList)
}

func validateCommand(command *Command) error {
	if command.Command == "" {
		return errors.New("command is required")
	}
	if command.Template == "" && command.Subcommands == nil {
		return errors.New("template or subcommands is required")
	}
	return nil
}
