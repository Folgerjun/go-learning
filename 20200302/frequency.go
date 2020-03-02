package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	// 没有文件路径参数 异常退出
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <file1> [<file2> [... <fileN>]]\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	frequencyForWord := map[string]int{} // 与:make(map[string]int)相同
	for _, filename := range commandLineFiles(os.Args[1:]) {
		updateFrequencies(filename, frequencyForWord)
	}
	reportByWords(frequencyForWord)
	wordsForFrequency := invertStringIntMap(frequencyForWord)
	reportByFrequency(wordsForFrequency)
}

// 实现跨平台的处理，当程序运行在 Windows 平台时，实现文件名通配功能。
func commandLineFiles(files []string) []string {
	if runtime.GOOS == "windows" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name) // 无效模式
			} else if matches != nil {
				args = append(args, matches...)
			}
		}
		return args
	}
	return files
}

// 处理文件的，它打开给定的文件，并使用 defer 在函数返回时关闭文件
func updateFrequencies(filename string, frequencyForWord map[string]int) {
	var file *os.File
	var err error
	if file, err = os.Open(filename); err != nil {
		log.Println("failed to open the file: ", err)
		return
	}
	defer file.Close()
	readAndUpdateFrequencies(bufio.NewReader(file), frequencyForWord)
}
func readAndUpdateFrequencies(reader *bufio.Reader,
	frequencyForWord map[string]int) {
	for { // 以字符串的形式一行一行地读取数据
		line, err := reader.ReadString('\n')
		// 忽略掉非单词的字符，并且过滤掉字符串开头和结尾的空白。
		for _, word := range SplitOnNonLetters(strings.TrimSpace(line)) {
			// 只需要记录含有两个以上（包括两个）字母的单词
			if len(word) > utf8.UTFMax ||
				utf8.RuneCountInString(word) > 1 {
				frequencyForWord[strings.ToLower(word)] += 1
			}
		}
		if err != nil {
			if err != io.EOF {
				log.Println("failed to finish reading the file: ", err)
			}
			break
		}
	}
}

// 非单词字符上对一个字符串进行切分
func SplitOnNonLetters(s string) []string {
	notALetter := func(char rune) bool { return !unicode.IsLetter(char) }
	return strings.FieldsFunc(s, notALetter)
}
func invertStringIntMap(intForString map[string]int) map[int][]string {
	stringsForInt := make(map[int][]string, len(intForString))
	for key, value := range intForString {
		stringsForInt[value] = append(stringsForInt[value], key)
	}
	return stringsForInt
}

// 输出结果按照字母顺序排序好打印
func reportByWords(frequencyForWord map[string]int) {
	words := make([]string, 0, len(frequencyForWord))
	wordWidth, frequencyWidth := 0, 0
	for word, frequency := range frequencyForWord {
		words = append(words, word)
		if width := utf8.RuneCountInString(word); width > wordWidth {
			wordWidth = width
		}
		if width := len(fmt.Sprint(frequency)); width > frequencyWidth {
			frequencyWidth = width
		}
	}
	sort.Strings(words)
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")
	for _, word := range words {
		fmt.Printf("%-*s %*d\n", wordWidth, word, frequencyWidth,
			frequencyForWord[word])
	}
}

// 按照频率升序排列打印
func reportByFrequency(wordsForFrequency map[int][]string) {
	frequencies := make([]int, 0, len(wordsForFrequency))
	for frequency := range wordsForFrequency {
		frequencies = append(frequencies, frequency)
	}
	sort.Ints(frequencies)
	width := len(fmt.Sprint(frequencies[len(frequencies)-1]))
	fmt.Println("Frequency → Words")
	for _, frequency := range frequencies {
		words := wordsForFrequency[frequency]
		sort.Strings(words)
		fmt.Printf("%*d %s\n", width, frequency, strings.Join(words, ", "))
	}
}
