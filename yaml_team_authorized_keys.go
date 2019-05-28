package flag

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type TeamAuthorizedKey struct {
	Team string   `yaml:"team"`
	Keys []string `yaml:"ssh_keys,flow"`
}

type YamlTeamAuthorizedKeys struct {
	File               string
	TeamAuthorizedKeys []TeamAuthorizedKey
}

func (f *YamlTeamAuthorizedKeys) UnmarshalFlag(value string) error {
	authorizedKeysBytes, err := ioutil.ReadFile(value)
	if err != nil {
		return fmt.Errorf("failed to read yaml authorized keys: %s", err)
	}

	f.File = value

	err = yaml.Unmarshal([]byte(authorizedKeysBytes), &f.TeamAuthorizedKeys)
	if err != nil {
		return fmt.Errorf("failed to parse yaml authorized keys: %s", err)
	}

	return nil
}

func (f *YamlTeamAuthorizedKeys) Reload() error {
	return f.UnmarshalFlag(f.File)
}
