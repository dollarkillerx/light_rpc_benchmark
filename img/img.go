package main

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	lightRPC := plotter.Values{96562, 111433}
	grpc := plotter.Values{64935, 109063}

	p := plot.New()

	p.Title.Text = "LightRPC vs GRPC"
	p.Y.Label.Text = ""

	w := vg.Points(20)
	stdBar, err := plotter.NewBarChart(lightRPC, w)
	if err != nil {
		log.Fatal(err)
	}
	stdBar.LineStyle.Width = vg.Length(0)
	stdBar.Color = plotutil.Color(0)
	stdBar.Offset = -w

	easyjsonBar, err := plotter.NewBarChart(grpc, w)
	if err != nil {
		log.Fatal(err)
	}
	easyjsonBar.LineStyle.Width = vg.Length(0)
	easyjsonBar.Color = plotutil.Color(1)

	p.Add(stdBar, easyjsonBar)
	p.Legend.Add("LightRPC", stdBar)
	p.Legend.Add("GRPC", easyjsonBar)
	p.Legend.Top = true
	//p.Legend.XOffs = 1
	p.NominalX("Single", "100 Thread Pool")

	if err = p.Save(5*vg.Inch, 5*vg.Inch, "./img/barchart.png"); err != nil {
		log.Fatal(err)
	}
}
