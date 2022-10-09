// loading the template src from local storage or from git
package cli

import (
	"errors"
	"io/ioutil"

	"github.com/bitcubix/tego/config"
	"github.com/manifoldco/promptui"
)

func Load() (string, error) {
	prompt := promptui.Prompt{
		Label: "Set Template",
	}

	folder_path, err := prompt.Run()
	if err != nil {
		return "", err
	}

	// search for template conig in dir
	files, err := ioutil.ReadDir(folder_path)
	if err != nil {
		return "", err
	}

	var config_file_path string

	for _, file := range files {
		for _, conf_file := range config.ALLOWED_PROJECT_JSON {
			if conf_file == file.Name() {
				config_file_path = folder_path + "/" + file.Name()
				break
			}
		}

		if config_file_path != "" {
			break
		}
	}

	if config_file_path != "" {
		return config_file_path, nil
	}

	return "", errors.New("template not found")
}
