package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configuration struct {
	Host string `json:"host,omitempty"`
	Port int    `json:"port"`
}

func ParseConfig(filepath string) (*Configuration, error) {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("error reading file " + filepath)
		return nil, err
	}
	defer f.Close()
	result, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("error reading file " + err.Error())
		return nil, err
	}
	c := new(Configuration)
	err = json.Unmarshal(result, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
