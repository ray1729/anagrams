package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "os"
    "sort"
    "strings"
)

type runeSlice []rune

func (r runeSlice) Len() int {
    return len(r)
}

func (r runeSlice) Less(i,j int) bool {
    return r[i] < r[j]
}

func (r runeSlice) Swap(i,j int) {
    r[i], r[j] = r[j], r[i]
}

func SortString(s string) string {
    r := []rune(s)
    sort.Sort(runeSlice(r))
    return string(r)
}

type StrSet map[string]bool

func (s StrSet) Elements() []string {
    xs := make([]string, 0)
    for x := range s {
        xs = append(xs, x)
    }
    return xs
}

func (s StrSet) Add(x string) {
    s[x] = true
}

func StrToKey(s string) string {
    return SortString(strings.ToLower(s))
}

type AnagramDictionary map[string]StrSet

func (dict AnagramDictionary) Add(s string) {
    key := StrToKey(s)
    xs, ok := dict[key]
    if !ok {
        xs = StrSet{}
    }
    xs.Add(s)
    dict[key] = xs
}

func (dict AnagramDictionary) Get(s string) []string {
    key := StrToKey(s)
    xs, ok := dict[key]
    if !ok {
        return nil
    }
    return xs.Elements()
}

func (dict AnagramDictionary) Has(s string) bool {
    key := StrToKey(s)
    _, ok := dict[key]
    return ok
}

func BuildDictionary(dictPath string) AnagramDictionary {
    file, err := os.Open(dictPath)
    if err != nil {
        log.Fatalf("Open %s: %v", dictPath, err)
    }
    defer file.Close()
    var dict = AnagramDictionary{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        word := scanner.Text()
        if strings.HasPrefix(word, "#") {
            continue
        }
        dict.Add(word)
    }
    return dict
}

func main() {
    dictPath := flag.String("dict", "/usr/share/dict/words", "Dictionary Path")
    flag.Parse()
    dict := BuildDictionary(*dictPath)
    for _, w := range flag.Args() {
        anagrams := dict.Get(w)
        if len(anagrams) == 0 {
            continue
        }
        fmt.Printf("%s: %s\n", w, strings.Join(anagrams, ", "))
    }
}
