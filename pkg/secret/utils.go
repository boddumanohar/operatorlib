package secret

import (
	"github.com/imdario/mergo"
)

// mergeData utility function merges data and stringData. It does not
// override the alreadyexisting keys in data and only append the keys
// from stringData
func mergeData(data map[string][]byte, stringData map[string]string) error {
	newData := make(map[string][]byte, len(stringData))
	for key, value := range stringData {
		newData[key] = []byte(value)
	}

	err := mergo.Merge(data, newData)
	if err != nil {
		return nil
	}

	return nil
}
