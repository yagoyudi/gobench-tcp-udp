package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	p := plot.New()

	p.Title.Text = "TCP vs UDP"
	p.Y.Label.Text = "Quantidade de bytes (MB)"
	p.X.Label.Text = "Tempo total de envio (s)"

	err := plotutil.AddLinePoints(p,
		"TCP", tcpPoints(),
		"UDP", udpPoints(),
	)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(5*vg.Inch, 5*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

// Sobre esses dados:
// - Em ambos os casos, todos os bytes chegaram.
// - Feito por wifi.

func udpPoints() plotter.XYs {
	pts := make(plotter.XYs, 5)
	yValues := []float64{
		100,   // 100mb
		500,   // 500mb
		1000,  // 1gb
		5000,  // 5gb
		10000, // 10gb
	}
	xValues := []float64{
		2.556249124,
		12.51920333,
		25.515900918,
		129.323465052,
		260.609461813,
	}

	for i := range pts {
		pts[i].X = xValues[i]
		pts[i].Y = yValues[i]
	}
	return pts
}

func tcpPoints() plotter.XYs {
	pts := make(plotter.XYs, 5)
	yValues := []float64{
		100,
		500,
		1000,
		5000,
		10000,
	}
	xValues := []float64{
		1.478977309,
		6.938055906,
		14.18291412,
		69.708409905,
		141.318771262,
	}

	for i := range pts {
		pts[i].X = xValues[i]
		pts[i].Y = yValues[i]
	}
	return pts
}
