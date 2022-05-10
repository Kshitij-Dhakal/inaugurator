package create

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/Kshitij-Dhakal/inaugurator/common"
	"github.com/Kshitij-Dhakal/inaugurator/pkg/config"
	"google.golang.org/protobuf/proto"
)

type BoilderplaterCreatorService interface {
	CreateBoilerplate(file string, command *CommandList) error
}

type BoilderplaterCreatorServiceImpl struct {
}

func (*BoilderplaterCreatorServiceImpl) CreateBoilerplate(file string, commands *CommandList) error {
	//read config file
	//parse template path and read file in template path
	//replace template with file data
	//save bytes file in $HOME/.inaugurator/{{command}}.bytes
	bytes, err := proto.Marshal(commands)
	if err != nil {
		return err
	}
	_, found := os.LookupEnv("HOME")
	if !found {
		return errors.New("HOME not found")
	}
	if _, err := os.Stat(os.ExpandEnv("$HOME/.config/inaugurator")); os.IsNotExist(err) {
		err = os.Mkdir(os.ExpandEnv("$HOME/.config/inaugurator"), 0700)
		if err != nil {
			return err
		}
	}
	err = updateCommandList(commands)
	if err != nil {
		return err
	}
	b, _ := json.Marshal(commands)
	if config.IsDevEnv() {
		err = common.StoreString(os.ExpandEnv("$HOME/.config/inaugurator/"+commands.GetName()+".json"), string(b))
		if err != nil {
			return err
		}
	}
	err = common.StoreBytes(os.ExpandEnv("$HOME/.config/inaugurator/"+commands.GetName()+".pb"), bytes)
	return err
}

func updateCommandList(command *CommandList) error {
	for _, v := range command.GetCommands() {
		err := updateCommand(v, 0)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateCommand(command *Command, i int32) error {
	i++
	if i > 2 {
		return errors.New("max number of subcommands reached")
	}
	//TODO: template validation
	if command.GetTemplates() != nil {
		// template, err := common.ReadFile(command.GetTemplate())
		// if err != nil {
		// 	return err
		// }
		// command.Template = template
	}
	if command.GetSubcommands() != nil {
		for _, v := range command.GetSubcommands() {
			err := updateCommand(v, i)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
