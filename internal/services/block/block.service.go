package block

import (
	"golang.org/x/exp/slices"
)

var SCHEMA_PATH = "./schema/"

type FileData struct {
	Blocks []BlockData `yaml:"blocks"`
}

type BlockData struct {
	Name       string       `yaml:"name"`
	Sql        string       `yaml:"sql"`
	Measures   []Measures   `yaml:"measures"`
	Dimensions []Dimensions `yaml:"dimensions"`
}

type Measures struct {
	Name string `yaml:"name"`
	Sql  string `yaml:"sql"`
	Type string `yaml:"type"`
}

type Dimensions struct {
	Name string `yaml:"name"`
	Sql  string `yaml:"sql"`
	Type string `yaml:"type"`
}

func GetBlockFromName(name string) *BlockData {
	blockInstance := GetInstance()
	for _, blockData := range blockInstance.Blocks {
		b := slices.IndexFunc(blockData.Blocks, func(data BlockData) bool { return data.Name == name })
		if b != -1 {
			return &blockData.Blocks[b]
		}
	}
	return nil
}
