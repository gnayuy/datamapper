package quadtree

import (
	"math"
	"log"
	"runtime"
	"sync"
	"image"
)

type QuadTree struct {

	depth int
	level int
	leaf  bool
	
	parent *QuadTree
	
	topLeft 		*QuadTree // 00
	topRight 		*QuadTree // 01
	bottomLeft 		*QuadTree // 10
	bottomRight 	*QuadTree // 11
	
	node *Node
	
}

func init(dimx, dimy, dimz int64) *QuadTree {
		
	var qt *QuadTree
	
	depthx := int64(math.log2(dimx)+0.5)-math.log2(minW)+1
	depthy := int64(math.log2(dimy)+0.5)-math.log2(minH)+1
	
	qt.depth = math.Max(depthx, depthy)
}

func construct(parent,root *QuadTree, depth,level int, xmin,ymin,zmin,resx,resy,resz float64, cx,cy,cz,w,h,d int64, ch chan bool) {
	
	depth = depth - 1
	level = level + 1
	
	if(depth < 0){
		ch <- true
		return
		
	}else{
		
		root = &QuadTree{depth,level,false,parent,nil,nil,nil,nil,nil,nil,nil,nil,nil}
		
		root->node = &Node{cx,cy,cz,w,h,d,resx,resy,resz,xmin,ymin,zmin}
		root->depth = depth + 1
		
		root->parent = parent
		
		resx = resx / 2.0
		resy = resy / 2.0
		resz = resz / 2.0
		
		go construct(root,root->topLeft,     depth,level,xmin,ymin,zmin,resx,resy,resz,cx,cy,cz,w,h,d,ch)
		go construct(root,root->topRight,    depth,level,xmin,ymin,zmin,resx,resy,resz,cx,cy,cz,w,h,d,ch)
		go construct(root,root->bottomLeft,  depth,level,xmin,ymin,zmin,resx,resy,resz,cx,cy,cz,w,h,d,ch)
		go construct(root,root->bottomRight, depth,level,xmin,ymin,zmin,resx,resy,resz,cx,cy,cz,w,h,d,ch)
			
		
	}
	
}



