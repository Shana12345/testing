package domain

import (
	"encoding/json"
	"fmt"
)

type msi = map[string]interface{}

type Config struct {
	Data map[string]interface{}
}

func (c *Config) SetFromBytes(data []byte) error {
	var rawConfig interface{}
	if err := json.Unmarshal(data, &rawConfig); err != nil {
		return err
	}
	fmt.Printf("%#v", rawConfig)

	untypedConfig, ok := rawConfig.(map[interface{}]interface{})
	if !ok {
		return fmt.Errorf("config is not a map")
	}

	config, err := convertKeysToStrings(untypedConfig)
	if err != nil {
		return err
	}

	c.Data = config
	return nil
}

func (c *Config) Get(serviceName string) (map[string]interface{}, error) {
	return nil, nil

}

func convertKeysToStrings(m map[interface{}]interface{}) (map[string]interface{}, error) {
	n := make(map[string]interface{})

	for k, v := range m {
		str, ok := k.(string)
		if !ok {
			return nil, fmt.Errorf("config key is not a string")
		}

		if vMap, ok := v.(map[interface{}]interface{}); ok {
			var err error
			v, err = convertKeysToStrings(vMap)
			if err != nil {
				return nil, err
			}
		}

		n[str] = v
	}
	return n, nil
}
