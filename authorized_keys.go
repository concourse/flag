package flag

import (
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v2"
)

type AuthorizedKeysMap map[string]AuthorizedKeys

// Can be removed once flags are deprecated
func (a AuthorizedKeysMap) Set(value string) error {
	authorizedKeyPaths := make(map[string]string)
	err := yaml.Unmarshal([]byte(value), &authorizedKeyPaths)
	if err != nil {
		return err
	}

	if authorizedKeyPaths != nil {
		return a.set(authorizedKeyPaths)
	}

	return nil
}

// Can be removed once flags are deprecated
func (a AuthorizedKeysMap) String() string {
	if a == nil {
		return ""
	}

	authorizedKeysString, _ := a.convertToString()
	return authorizedKeysString
}

// Can be removed once flags are deprecated
func (a AuthorizedKeysMap) Type() string {
	return "AuthorizedKeysMap"
}

func (a AuthorizedKeysMap) convertToString() (string, error) {
	authorizedKeysPaths := make(map[string]string)
	for key, authorizedKey := range a {
		authorizedKeysPaths[key] = authorizedKey.File
	}

	authorizedKeysString, err := yaml.Marshal(authorizedKeysPaths)
	if err != nil {
		return "", err
	}

	return string(authorizedKeysString), nil
}

func (a *AuthorizedKeysMap) set(value map[string]string) error {
	authorizedKeysMap := make(AuthorizedKeysMap)
	for key, authorizedKeyPath := range value {
		var authorizedKeys AuthorizedKeys
		err := authorizedKeys.Set(authorizedKeyPath)
		if err != nil {
			return err
		}

		authorizedKeysMap[key] = authorizedKeys
	}

	a = &authorizedKeysMap

	return nil
}

type AuthorizedKeys struct {
	File string
	Keys []ssh.PublicKey
}

// Reload reloads the value of the Keys
func (a *AuthorizedKeys) Reload() error {
	return a.Set(a.File)
}

func (a AuthorizedKeys) MarshalYAML() (interface{}, error) {
	return a.File, nil
}

func (a *AuthorizedKeys) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var path string
	err := unmarshal(&path)
	if err != nil {
		return err
	}

	if path != "" {
		return a.Set(path)
	}

	return nil
}

// Can be removed once flags are deprecated
func (a *AuthorizedKeys) Set(value string) error {
	authorizedKeysBytes, err := ioutil.ReadFile(value)
	if err != nil {
		return fmt.Errorf("failed to read authorized keys: %s", err)
	}

	var authorizedKeys []ssh.PublicKey

	for {
		key, _, _, rest, err := ssh.ParseAuthorizedKey(authorizedKeysBytes)
		if err != nil {
			// there's no good error to check for here
			break
		}

		authorizedKeys = append(authorizedKeys, key)

		authorizedKeysBytes = rest
	}

	a.File = value
	a.Keys = authorizedKeys

	return nil
}

// Can be removed once flags are deprecated
func (a *AuthorizedKeys) String() string {
	if a == nil {
		return ""
	}

	return a.File
}

// Can be removed once flags are deprecated
func (a *AuthorizedKeys) Type() string {
	return "AuthorizedKeys"
}
