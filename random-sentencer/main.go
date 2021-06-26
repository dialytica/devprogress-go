package main

import (
    "fmt"
    "math/rand"
    "time"
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

func main() {
    s := rand.NewSource(time.Now().UnixNano())
    r := rand.New(s)

    subjectWeight := r.Intn(len(subjects))
    predicateWeight := r.Intn(len(predicates))
    objectWeight := r.Intn(len(objects))
    adjectiveWeight := r.Intn(len(adjectives))

    fmt.Println(subjects[subjectWeight], predicates[predicateWeight], adjectives[adjectiveWeight], objects[objectWeight])
}
