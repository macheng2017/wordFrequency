package chart

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"os"
	"wordFrequency/sort"
)

const wordCount = 30

// generate random data for bar chart
func generateBarItems(list sort.PairList) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < wordCount; i++ {
		items = append(items, opts.BarData{Value: list[i].Value})
	}
	return items
}

func CreateChart(list sort.PairList) {
	// create a new bar instance
	bar := charts.NewBar()

	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "格林童话中的词频",
			Subtitle: "This is the subtitle.",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "1200px",
			Height: "600px",
		}),
	)

	// iowriter
	f, err := os.Create("bar.html")
	if err != nil {
		panic(err)
	}
	var keys []string
	for i := 0; i < wordCount; i++ {
		keys = append(keys, list[i].Key)
	}
	fmt.Println(keys)
	// Put some data in instance
	bar.SetXAxis(keys).
		AddSeries("Category A", generateBarItems(list))

	// Where the magic happens
	bar.Render(f)

}
