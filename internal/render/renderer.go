package render

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

const DPI = 96.0

func Render(query string, data map[string][][]float64, width int, height int, filename string) error {
	p := plot.New()

	p.Title.Text = query
	p.X.Label.Text = "Time UTC"
	p.Y.Label.Text = "Value"

	p.X.Tick.Marker = plot.TimeTicks{
		Format: "15:04",
	}

	counter := 0

	for label, series := range data {
		points := make(plotter.XYs, len(series))

		for i, point := range series {
			points[i].X = point[0]
			points[i].Y = point[1]
		}

		line, err := plotter.NewLine(points)
		if err != nil {
			return err
		}

		line.Color = plotutil.Color(counter)
		line.Width = vg.Points(1.5)

		p.Add(line)
		p.Legend.Add(label, line)

		counter++
	}

	return p.Save(vg.Length(float64(width)/DPI)*vg.Inch, vg.Length(float64(height)/DPI)*vg.Inch, filename)
}
