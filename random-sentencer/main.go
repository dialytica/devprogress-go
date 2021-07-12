package main

import (
    "io"
    "fmt"
    "math/rand"
    "time"
    "os"
    "json"
    "bufio"
)

var subjects []string = []string{
    "The World",
    "I",
    "You",
    "They",
    "He",
    "She",
    "Something",
    "Someone",
    "Dog",
    "Cat",
    "Flower",
    "System",
    "This",
    "That",
    "Those",
    "These",
}

var predicates []string = []string{
    "eat",
    "is",
    "destroy",
    "blow",
    "make",
    "clean",
    "exude",
    "killed by",
    "estinguished by",
    "loved by",
    "castrated for",
    "subject to",
    "animate",
    "accelerate",
    "made from",
    "understand",
}

var objects []string = []string{
    "it",
    "her",
    "him",
    "them",
    "us",
    "cat",
    "dog",
    "flower",
    "you",
    "silk",
    "terror",
}

var adjectives []string = []string{
    "big",
    "round",
    "ugly",
    "handsome",
    "lost",
    "annoying",
    "red",
    "bouncy",
    "lovely",
    "complicated",
    "stupid",
    "suspicious",
    "confusing",
    "happy",
    "sad",
    "angry",
    "elated",
    "elongated",
    "short",
}

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

func LoadLanguageDef(reader *bufio.Reader) (LanguageDefinition, error) {
    decoder := json.NewDecoder(reader)

    var langDef LanguageDefinition
    err = decoder.Decode(&langDef)
    CheckError(err)

    return langDef, nil
}

func GetWord(index int, reader *bufio.Reader) (string, error) {
    scanner := bufio.NewScanner(reader)
    c := 0

    for scanner.Scan() {
        c++
        if c == index {
            return scanner.Text(), nil
        }
    }

    return nil, fmt.Errorf("index: %g out of range", index)
}

func CountWordStream(reader *bufio.Reader) (int, error) {
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

func (lang *LanguageDefinition) Generate(seed int) string {
}

func main() {
    s := rand.NewSource(time.Now().UnixNano())
    r := rand.New(s)

    subjectWeight := r.Intn(len(subjects))
    predicateWeight := r.Intn(len(predicates))
    objectWeight := r.Intn(len(objects))
    adjectiveWeight := r.Intn(len(adjectives))

    fmt.Println(subjects[subjectWeight], predicates[predicateWeight], adjectives[adjectiveWeight], objects[objectWeight])
}
