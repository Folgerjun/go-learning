package main

import (
	"fmt"
	"sort"
	"strings"
)

var original = []string{
	"Nonmetals",
	"    Hydrogen",
	"    Carbon",
	"    Nitrogen",
	"    Oxygen",
	"Inner Transitionals",
	"    Lanthanides",
	"        Europium",
	"        Cerium",
	"    Actinides",
	"        Uranium",
	"        Plutonium",
	"        Curium",
	"Alkali Metals",
	"    Lithium",
	"    Sodium",
	"    Potassium",
}

func main() {
	fmt.Println("|     Original      |       Sorted      |")
	fmt.Println("|-------------------|-------------------|")
	sorted := SortedIndentedStrings(original) // 最初是 []string
	for i := range original {                 // 在全局变量中设置
		fmt.Printf("|%-19s|%-19s|\n", original[i], sorted[i])
	}
	/*
		|     Original      |       Sorted      |
		|-------------------|-------------------|
		|Nonmetals          |Alkali Metals      |
		|    Hydrogen       |    Lithium        |
		|    Carbon         |    Potassium      |
		|    Nitrogen       |    Sodium         |
		|    Oxygen         |Inner Transitionals|
		|Inner Transitionals|    Actinides      |
		|    Lanthanides    |        Curium     |
		|        Europium   |        Plutonium  |
		|        Cerium     |        Uranium    |
		|    Actinides      |    Lanthanides    |
		|        Uranium    |        Cerium     |
		|        Plutonium  |        Europium   |
		|        Curium     |Nonmetals          |
		|Alkali Metals      |    Carbon         |
		|    Lithium        |    Hydrogen       |
		|    Sodium         |    Nitrogen       |
		|    Potassium      |    Oxygen         |
	*/
}

// 排序
func SortedIndentedStrings(slice []string) []string {
	entries := populateEntries(slice)
	return sortedEntries(entries)
}
func populateEntries(slice []string) Entries {
	indent, indentSize := computeIndent(slice)
	entries := make(Entries, 0)
	for _, item := range slice {
		i, level := 0, 0
		for strings.HasPrefix(item[i:], indent) {
			i += indentSize
			level++
		}
		// strings.TrimSpace(s) 返回将 s 前后端所有空白（unicode.IsSpace指定）都去掉的字符串
		key := strings.ToLower(strings.TrimSpace(item))
		addEntry(level, key, item, &entries)
	}
	return entries
}

// 返回空格字符串
func computeIndent(slice []string) (string, int) {
	for _, item := range slice {
		if len(item) > 0 && (item[0] == ' ' || item[0] == '\t') {
			whitespace := rune(item[0])
			for i, char := range item[1:] {
				if char != whitespace {
					i++
					// 返回 i 个空格字符串串联的字符串
					return strings.Repeat(string(whitespace), i), i
				}
			}
		}
	}
	return "", 0
}
func addEntry(level int, key, value string, entries *Entries) {
	if level == 0 {
		*entries = append(*entries, Entry{key, value, make(Entries, 0)})
	} else {
		addEntry(level-1, key, value,
			&((*entries)[entries.Len()-1].children))
	}
}
func sortedEntries(entries Entries) []string {
	var indentedSlice []string
	sort.Sort(entries)
	for _, entry := range entries {
		populateIndentedStrings(entry, &indentedSlice)
	}
	return indentedSlice
}
func populateIndentedStrings(entry Entry, indentedSlice *[]string) {
	*indentedSlice = append(*indentedSlice, entry.value)
	sort.Sort(entry.children)
	for _, child := range entry.children {
		populateIndentedStrings(child, indentedSlice)
	}
}

type Entry struct {
	key      string
	value    string
	children Entries
}
type Entries []Entry

func (entries Entries) Len() int { return len(entries) }
func (entries Entries) Less(i, j int) bool {
	return entries[i].key < entries[j].key
}
func (entries Entries) Swap(i, j int) {
	entries[i], entries[j] = entries[j], entries[i]
}
