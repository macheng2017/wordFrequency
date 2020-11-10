package main

import (
	"bufio"
	"fmt"
	"github.com/yanyiwu/gojieba"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"unicode"
	chart "wordFrequency/chart"
	sort2 "wordFrequency/sort"
)

func ReadFile(path string) (bytes []byte, err error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(file)
	bytes, err = ioutil.ReadAll(reader)

	//[Golang rune []byte string 的相互转换_次元代码-CSDN博客](https://blog.csdn.net/dengming0922/article/details/80883574)
	//[golang - 如何判断字符是不是中文？ - SegmentFault 思否](https://segmentfault.com/q/1010000000595663)
	r := []rune(string(bytes))
	var bytesIsHan []rune
	for _, ru := range r {
		// 判断字符是否是汉字
		if unicode.Is(unicode.Han, ru) {
			bytesIsHan = append(bytesIsHan, ru)
		}
	}

	return []byte(string(bytesIsHan)), err
}

func SplitWord(bytes []byte) []string {
	x := gojieba.NewJieba()
	defer x.Free()
	var words []string
	//words = x.CutAll(string(bytes))
	words = x.Cut(string(bytes), true)
	//fmt.Println("全模式:", strings.Join(words, "/"))
	//fmt.Printf("font size %d", len(words))
	return words
}

// 使用struct 灌入map的数据对value进行排序，必须实现sort接口
func rankByWordCount(wordFrequencies map[string]interface{}) sort2.PairList {
	pl := make(sort2.PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = sort2.Pair{k, v.(int)}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

func main() {
	bytes, err := ReadFile("./doc/gelin.txt")
	if err != nil {
		log.Fatal(err)
	}
	words := SplitWord(bytes)

	wordFrequencyMap := make(map[string]interface{})
	for _, word := range words {

		if value, ok := wordFrequencyMap[word]; ok {
			wordFrequencyMap[word] = value.(int) + 1
		} else {
			wordFrequencyMap[word] = 1
		}
	}
	pairList := rankByWordCount(wordFrequencyMap)
	fmt.Print(pairList)
	examples := chart.WordcloudExamples{}
	examples.Examples(wordFrequencyMap)
	chart.CreateChart(pairList)

}
