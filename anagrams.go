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

func StrToKey(s string) string {
    return SortString(strings.ToLower(s))
}

type AnagramDictionary map[string]map[string]bool

func (dict AnagramDictionary) Add(s string) {
    key := StrToKey(s)
    xs, ok := dict[key]
    if !ok {
        xs = make(map[string]bool)
    }
    xs[s] = true
    dict[key] = xs
}

func (dict AnagramDictionary) Get(s string) []string {
    key := StrToKey(s)
    xs, ok := dict[key]
    if !ok {
        return nil
    }
    elements := make([]string, len(xs))
    i := 0
    for x := range xs {
        elements[i] = x
        i++
    }
    return elements
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
