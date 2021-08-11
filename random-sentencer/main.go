package main

import (
	"fmt"
	"io"
	"path/filepath"
	// "math/rand"
	// "time"
	"bufio"
	"encoding/json"
	"os"
)

type LanguageDefinition struct {
	Structures []string
	Data       DataMap
}

type DataMap map[string]WordDefinition

type WordDefinition struct {
	Name    string
	Sources []SourceDefinition
	Rules   []WordRuleDefinition
}

type SourceDefinition struct {
	Name   string
	Type   string
	Data   []string
	Tags   []string
	Weight float64
}

type WordRuleDefinition struct {
	Tags []string
	Rule []string
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadLanguageDef(reader io.Reader) (*LanguageDefinition, error) {
	decoder := json.NewDecoder(reader)

	var langDef LanguageDefinition
	err := decoder.Decode(&langDef)
	if err != nil {
		return nil, err
	}

	return &langDef, nil
}

func GetWord(index int, reader io.Reader) (string, error) {
	scanner := bufio.NewScanner(reader)
	c := 0

	for scanner.Scan() {
		if c == index {
			return scanner.Text(), nil
		}
		c++
	}

	return "", fmt.Errorf("index: %g out of range", index)
}

func CountWordStream(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	c := 0

	for scanner.Scan() {
		c++
	}

	if err := scanner.Err(); err != nil && err != io.EOF {
		return -1, err
	}

	return c, nil
}

// func (lang *LanguageDefinition) Generate(seed int) string {
// }

func main() {
	langDefDirectoryPath := filepath.Join(".", "data", "en")
	structureFile, err := os.Open(filepath.Join(langDefDirectoryPath, "structure.json"))
	CheckError(err)

	var langDef *LanguageDefinition

	langDef, err = LoadLanguageDef(structureFile)
	defer structureFile.Close()
	CheckError(err)

	fmt.Println("Data file:")
	fmt.Println(langDef.Data["S"].Sources[4].Data[0])

	dataFile, err := os.Open(filepath.Join(langDefDirectoryPath, langDef.Data["S"].Sources[4].Data[0]))

	CheckError(err)

	word, err := GetWord(0, dataFile)
	CheckError(err)
	fmt.Println("Get word:")

	fmt.Println(word)

}
