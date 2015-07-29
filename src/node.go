package Node

import (
	"math"
	"log"
	"runtime"
	"sync"
	"image"
)

const (
	minW=64
	minH=64
	minD=64
)

type NodeOT struct {

	coordx int64	// tile coordinates
	coordy int64
	coordz int64
	
	width int64		// tile dimensions
	height int64
	depth int64
	
	resx float64	// resolution
	resy float64
	resz float64
	
	xmin float64	// physical coordinates
	ymin float64
	zmin float64
	
	// xmax = xmin + width - 1
	// ymax = ymin + height - 1
	// zmax = zmin + depth - 1
	
	data interface{}
}
