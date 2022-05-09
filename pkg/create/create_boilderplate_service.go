package create

import (
	"encoding/json"
	"os"

	"github.com/Kshitij-Dhakal/inaugurator/common"
)

type BoilderplaterCreatorService interface {
	CreateBoilerplate(file string, command *Command) (string, error)
}


type BoilderplaterCreatorServiceImpl struct {
	
}

func (*BoilderplaterCreatorServiceImpl) CreateBoilerplate(file string, command *Command) (string, error) {
	//read config file
	//parse template path and read file in template path
	//replace template with file data
	//save json file in $HOME/.inaugurator/{{command}}.json
	json, err := json.Marshal(command)
	if err != nil {
		return "", err
	}
	common.WriteFile(os.ExpandEnv("$HOME/.inaugurator/" + command.Command + ".json"), string(json))
	return "", nil
}