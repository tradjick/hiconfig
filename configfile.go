package hiconfig

import (
	"encoding/json"
	"os"
)

func ConfigJsonFile(p string, i interface{}) {
	fc, err := os.ReadFile(p)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(fc, i)
	if err != nil {
		panic(err)
	}
}
