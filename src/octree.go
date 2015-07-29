package octree

import (
	"math"
	"log"
	"runtime"
	"sync"
	"image"
)

type Octree struct {

	depth int
	level int
	leaf  bool
	
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

func init(){
	
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