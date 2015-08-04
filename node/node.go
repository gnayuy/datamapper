package node

import (
	"image"
)

type Data struct {
	
	key int64
	val image.Image
	
}

type Node struct {

	coordx int64	// tile coordinates
	coordy int64
	coordz int64
	
	width int64	// tile dimensions
	height int64
	depth int64
	
	resx float64	// resolution
	resy float64
	resz float64
	
	xmin int64	// physical coordinates
	ymin int64
	zmin int64
	
	// xmax = xmin + width - 1
	// ymax = ymin + height - 1
	// zmax = zmin + depth - 1
	
	buf *Data
}

func NewNode(cx,cy,cz int64, w,h,d int64, rx,ry,rz float64, x,y,z int64, p *Data) *Node {
	
	var n *Node
	
	n.coordx = cx
	n.coordy = cy
	n.coordz = cz
	
	n.width = w
	n.height = h
	n.depth = d
	
	n.resx = rx
	n.resy = ry
	n.resz = rz
	
	n.xmin = x
	n.ymin = y
	n.zmin = z
	
	n.buf = p
	
	return n
	
}

func GetNode(n *Node)(int64, int64, int64, int64, int64, int64, float64, float64, float64, int64, int64, int64, *Data){
	return n.coordx,n.coordx,n.coordz,n.width,n.height,n.depth,n.resx,n.resy,n.resz,n.xmin,n.ymin,n.zmin,n.buf
}



