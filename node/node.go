package node

import (
	"image"
)

const (
	otW=64
	otH=64
	otD=64
	
	qtW=512
	qtH=512
)

type Data struct {
	
	key int64
	val image.Image
	
}

type Node struct {

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
	
	buf *Data
}



