package chap3

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func RunSurface() {
	http.HandleFunc("/", surfacePlot)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func surfacePlot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	// GPT says string builder may be better optimized for strings since any conversion is needed
	var buf bytes.Buffer

	// This may be inefficient, but it is simple to understand for now.
	minZ, maxZ := math.Inf(1), math.Inf(-1)
	for i := range cells {
		for j := range cells {
			_, _, z := corner(i, j)
			if math.IsNaN(z) || math.IsInf(z, 0) {
				continue
			}
			if z < minZ {
				minZ = z
			}
			if z > maxZ {
				maxZ = z
			}
		}
	}

	fmt.Fprintf(&buf, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke:grey; fill:white; stroke-width:0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := range cells {
		for j := range cells {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			// I can also includo isInf
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(az) ||
				math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(bz) ||
				math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(cz) ||
				math.IsNaN(dx) || math.IsNaN(dy) || math.IsNaN(dz) {
				continue
			}

			avgZ := (az + bz + cz + dz) / 4
			normZ := (avgZ - minZ) / (maxZ - minZ)

			r := int(255 * normZ)
			b := int(255 * (1 - normZ))
			fill := fmt.Sprintf("#%02x00%02x", r, b)

			fmt.Fprintf(&buf, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, fill)
		}
	}

	buf.WriteString("</svg>")

	fmt.Fprint(w, buf.String())
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
