package boilerplate

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type BoilderplateListerService interface {
	ListBoilerplates() error
}

type BoilerplateListerServiceImpl struct {
}

func (s BoilerplateListerServiceImpl) ListBoilerplates() error {
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
	entry, err := os.ReadDir(os.ExpandEnv("$HOME/.config/inaugurator"))
	if err != nil {
		return nil
	}
	fmt.Println("Listing available boilerplates : ")
	var boilerplates map[string]string = make(map[string]string)
	for _, v := range entry {
		boilerplates[strings.Replace(strings.Replace(v.Name(), ".pb", "", -1), ".json", "", -1)] = ""
	}
	for k, _ := range boilerplates {
		fmt.Println(k)
	}
	return nil
}
