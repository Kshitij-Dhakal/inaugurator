package create

import (
	"errors"
	"os"

	"github.com/Kshitij-Dhakal/inaugurator/common"
	"google.golang.org/protobuf/proto"
)

type BoilderplaterCreatorService interface {
	CreateBoilerplate(file string, command *CommandList) (string, error)
}

type BoilderplaterCreatorServiceImpl struct {
}

func (*BoilderplaterCreatorServiceImpl) CreateBoilerplate(file string, command *CommandList) (string, error) {
	//read config file
	//parse template path and read file in template path
	//replace template with file data
	//save bytes file in $HOME/.inaugurator/{{command}}.bytes
	bytes, err := proto.Marshal(command)
	if err != nil {
		return "", err
	}
	_, found := os.LookupEnv("HOME")
	if !found {
		return "", errors.New("HOME not found")
	}
	if _, err := os.Stat(os.ExpandEnv("$HOME/.config/inaugurator")); os.IsNotExist(err) {
		err = os.Mkdir(os.ExpandEnv("$HOME/.config/inaugurator"), 0700)
		if err != nil {
			return "", err
		}
	}
	
	err = common.StoreBytes(os.ExpandEnv("$HOME/.config/inaugurator/"+command.GetName()+".pb"), bytes)
	return "", err
}
