package quadtree

import (
	"math"
	"log"
	"runtime"
	"sync"
	"image"
	"github.com/nfnt/resize"
)

type QuadTree struct {

	depth		int
	level		int
	leaf		bool
	
	dataAvail	bool
	
	parent		*QuadTree
	
	topLeft 	*QuadTree // 00
	topRight 	*QuadTree // 01
	bottomLeft 	*QuadTree // 10
	bottomRight *QuadTree // 11
	
	node 		*Node
	
}

// VoxelSize, MinPoint, MaxPoint
func init(xmin, xmax, ymin, ymax, zmin, zmax int64, resx, resy, resz float64) *QuadTree {
		
	var qt *QuadTree
	
	dimx := xmax - xmin + 1
	dimy := ymax - ymin + 1
	
	depthx := int64(math.log2(dimx)+0.5)-math.log2(qtW)+1
	depthy := int64(math.log2(dimy)+0.5)-math.log2(qtH)+1
	
	depth := math.Max(depthx, depthy)
	
	construct(nil,qt,depth,0,0,0,0,resx,resy,resz,xmin,ymin,zmin,minW,minH,minD,ch)
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
		
		root->dataAvail = false
		
		if(depth==1){
			root->leaf = true
		}else{
			root->leaf = false
		}	
		
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

func getData(qt *QuadTree, ch chan bool) {
	// if it is leaf, get the data from database
	// else get the data from its childtren and then resize
	
	if(qt->dataAvail==true){
		ch <- true
		return;
	}else{
		
		if(qt->leaf==true){
			// get data
			
			// save data
			
		}else{
			// get data from its children's data
			
			// resize
			
			// save data
		}
		
		qt->dataAvail = true;
	}
	
}

