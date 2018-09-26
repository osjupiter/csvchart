package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"io"
	"strconv"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)
type Chart interface{
	Render(rp chart.RendererProvider, w io.Writer) error
}

func pieChart(bars []chart.Value) Chart {
	return chart.PieChart{
		Title:      "Test Bar Chart",
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height: 512,
		Values: bars,
	}
}

func barChart(bars []chart.Value) Chart {
	return chart.BarChart{
		Title:      "Test Bar Chart",
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		XAxis: chart.Style{
			Show: true,
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		Bars: bars,
	}
}
func readRecords(name string) [][]string {
	in, e1 := os.Open(name)
	if e1 != nil {
		panic(e1)
	}
	r := csv.NewReader(in)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records

}
func constructValue(records [][]string) []chart.Value {
	xkey := "Name"
	ykey := "N"

	// invest shema
	columnNames := (records[0])
	columnMap := make(map[string]int)
	for i, v := range columnNames {
		columnMap[v] = i
	}
	// construct value
	values := []chart.Value{}
	for i, row := range records {
		if i == 0 {
			continue
		}
		xind := columnMap[xkey]
		yind := columnMap[ykey]
		f, e1 := strconv.ParseFloat(row[yind], 32)
		if e1 != nil {
			panic(e1)
		}
		values = append(values, chart.Value{Value: f, Label: row[xind]})
	}
	return values
}

func main1() {
	records := readRecords("groupName")
	// todo: from option
	values := constructValue(records)
	//barChart(values)
	pieChart(values)
	//}

}

func constructSeris(records [][]string) ([]float64, []float64) {
	xkey := "SepalLength"
	ykey := "SepalWidth"
	// invest shema
	columnNames := (records[0])
	columnMap := make(map[string]int)
	for i, v := range columnNames {
		columnMap[v] = i
	}
	// construct value
	xvalues := []float64{}
	yvalues := []float64{}
	//series := []chart.Series{
	//	chart.ContinuousSeries{
	//		XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
	//		YValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
	//	},
	for i, row := range records {
		if i == 0 {
			continue
		}
		xind := columnMap[xkey]
		yind := columnMap[ykey]
		xx, e1 := strconv.ParseFloat(row[xind], 32)
		if e1 != nil {
			panic(e1)
		}
		yy, e2 := strconv.ParseFloat(row[yind], 32)
		if e2 != nil {
			panic(e1)
		}
		xvalues = append(xvalues, xx)
		yvalues = append(yvalues, yy)
	}
	return xvalues, yvalues

}
func scatter(xvalues []float64, yvalues []float64) Chart {
	series := []chart.Series{
		chart.ContinuousSeries{
			Style: chart.Style{
				Show:        true,
				StrokeWidth: chart.Disabled,
				DotWidth:    5,
				DotColor:    drawing.Color{255, 0, 0, 255},
				Padding:     chart.NewBox(10, 10, 10, 10),
			},
			XValues: xvalues,
			YValues: yvalues,
		},
	}
	return chart.Chart{
		Series: series,
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
	}

}

func main() {
	records := readRecords("len_wid")
	// todo: from option
	xv, yv := constructSeris(records)

	graph := scatter(xv, yv)

	file, err := os.Create("hoge.png")
	if err != nil {
		panic(err)
	}

	err2 := graph.Render(chart.PNG, file)
	if err2 != nil {
		fmt.Printf("Error rendering chart: %v\n", err2)
	}
	//}

}
