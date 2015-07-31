package octree

import (
	"math"
	"log"
	"runtime"
	"sync"
	"image"
	
	"github.com/gnayuy/datamapper/node"
)

type Octree struct {

	depth int
	level int
	
	isLeaf		bool
	dataAvail	bool
	
	parent *Octree
	
	topLeftFront 		*Octree // 000
	topLeftBack			*Octree // 001
	bottomLeftFront		*Octree // 010
	bottomLeftBack		*Octree // 011
	topRightFront		*Octree // 100
	topRightBack		*Octree // 101
	bottomRightFront	*Octree // 110
	bottomRightBack		*Octree // 111

	node *Node
	
}

func init(xmin, xmax, ymin, ymax, zmin, zmax int64, resx, resy, resz float64) *QuadTree {
		
	var qt *QuadTree
	
	dimx := xmax - xmin + 1
	dimy := ymax - ymin + 1
	dimz := zmax - zmin + 1
	
	depthx := int64(math.log2(dimx)+0.5)-math.log2(otW)+1
	depthy := int64(math.log2(dimy)+0.5)-math.log2(otH)+1
	depthz := int64(math.log2(dimz)+0.5)-math.log2(otD)+1
	
	depth := math.Max(depthx, depthy)
	
	construct(nil,qt,depth,0,0,0,0,resx,resy,resz,xmin,ymin,zmin,otW,otH,otD,ch)
}

func construct(parent,root *Octree, depth,level int, xmin,ymin,zmin,resx,resy,resz float64, cx,cy,cz,w,h,d int64, ch chan bool) {
	
	depth = depth - 1
	level = level + 1
	
	if(depth < 0){
		ch <- true
		return
		
	}else{
		
		root = &Octree{depth,level,false,parent,nil,nil,nil,nil,nil,nil,nil,nil,nil}
		
		root->node = &Node{cx,cy,cz,w,h,d,resx,resy,resz,xmin,ymin,zmin}
		root->depth = depth + 1
		
		root->parent = parent
		
		resx = resx / 2.0
		resy = resy / 2.0
		resz = resz / 2.0
		
		go construct(root,root->topLeftFront,     depth,level,xmin,ymin,zmin,resx,resy resz,cx,cy,cz,w,h,d,ch)
		go construct(root,root->topLeftBack,      depth,level,xmin,ymin,zmin,resx,resy resz,cx,cy,cz,w,h,d,ch)
		go construct(root,root->bottomLeftFront,  depth,level,xmin,ymin,zmin,resx,resy resz,cx,cy,cz,w,h,d,ch)
		go construct(root,root->bottomLeftBack,   depth,level,xmin,ymin,zmin,resx,resy resz,cx,cy,cz,w,h,d,ch)
		go construct(root,root->topRightFront,    depth,level,xmin,ymin,zmin,resx,resy resz,cx,cy,cz,w,h,d,ch)
		go construct(root,root->topRightBack,     depth,level,xmin,ymin,zmin,resx,resy resz,cx,cy,cz,w,h,d,ch)
		go construct(root,root->bottomRightFront, depth,level,xmin,ymin,zmin,resx,resy resz,cx,cy,cz,w,h,d,ch)
		go construct(root,root->bottomRightBack,  depth,level,xmin,ymin,zmin,resx,resy resz,cx,cy,cz,w,h,d,ch)
			
		
	}
	
}