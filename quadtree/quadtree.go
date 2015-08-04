package quadtree

import (
	"math"
	//"log"
	//"runtime"
	//"sync"
	//"image"
	//"github.com/nfnt/resize"
	
	"fmt"
	
	"github.com/gnayuy/datamapper/node"
)

 // Quadrant ordering (Z-order curve):
 //    +––––––+––––––+
 //    |      |      |
 //    |TL(00)|TR(01)|
 //    |      |      |
 //    +–––––––––––––+
 //    |      |      |
 //    |BL(10)|BR(11)|
 //    |      |      |
 //    +––––––+––––––+
 
const (
	qtW=512
	qtH=512
)

type QuadTree struct {

	depth		int
	level		int
	
	isLeaf		bool
	dataAvail	bool // in immutable storage ?
	
	parent		*QuadTree
	
	TL		*QuadTree // 00
	TR		*QuadTree // 01
	BL		*QuadTree // 10
	BR		*QuadTree // 11
	
	node 		*node.Node
	
}

type empty struct {}
type semaphore chan empty

// VoxelSize, MinPoint, MaxPoint
func Init(xmin, xmax, ymin, ymax, zmin, zmax int64, resx, resy, resz float64) *[]QuadTree {
		
	dimz := zmax - zmin + 1
	
	qt := make([]QuadTree, dimz)
	
	dimx := xmax - xmin + 1
	dimy := ymax - ymin + 1
	
	depthx := int(math.Log2(float64(dimx))+0.5)-int(math.Log2(float64(qtW)))+1
	depthy := int(math.Log2(float64(dimy))+0.5)-int(math.Log2(float64(qtH)))+1
	
	depth := int(math.Max(float64(depthx), float64(depthy)))
	
	fmt.Println("The depth of this quadtree is ", depth)

	//go func() {
	//    Construct(nil,qt,depth,0,xmin,ymin,zmin,resx,resy,resz,0,0,0,qtW,qtH,1,ch)
	//    ch <- false
	//}()
	//<- ch
	
	sem := make (semaphore, dimz);
	for z := zmin; z < zmax; z++ {
		ch := make(chan bool)
        go func (z int64) {
            Construct(nil,&qt[z],depth,0,xmin,ymin,z,resx,resy,resz,0,0,0,qtW,qtH,1,ch) 
            sem <- empty{};
			
			<-ch
			
        } (z);
    }
    for z := zmin; z < zmax; z++ {
		<- sem // release dimz resources
	}
	
	return &qt
}

func Construct(parent,root *QuadTree, depth,level int, xmin,ymin,zmin int64, resx,resy,resz float64, cx,cy,cz,w,h,d int64, ch chan bool) {
	
	depth = depth - 1
	level = level + 1
	
	fmt.Printf("Construct depth %v level %v \n",depth, level)
	
	if(depth < 0){
		ch <- true
		return
		
	}else{
		
		root = &QuadTree{depth,level,false,false,parent,nil,nil,nil,nil,nil}
		
		root.node = new(node.Node)
		root.node.NewNode(cx,cy,cz,w,h,d,resx,resy,resz,xmin,ymin,zmin,nil)
		
		root.depth = depth + 1
		
		root.dataAvail = false
		
		if(depth==0){
			root.isLeaf = true
		}else{
			root.isLeaf = false
		}
		
		root.parent = parent
		
		resx = resx / 2.0
		resy = resy / 2.0
		resz = resz / 2.0
		
		cx = cx * 2
		cy = cy * 2
		cz = cz * 2
		
		go func() {
			Construct(root,root.TL,depth,level,xmin,ymin,zmin,resx,resy,resz,cx,  cy,  cz,w,h,d,ch)
			Construct(root,root.TR,depth,level,xmin,ymin,zmin,resx,resy,resz,cx+1,cy,  cz,w,h,d,ch)
			Construct(root,root.BL,depth,level,xmin,ymin,zmin,resx,resy,resz,cx,  cy+1,cz,w,h,d,ch)
			Construct(root,root.BR,depth,level,xmin,ymin,zmin,resx,resy,resz,cx+1,cy+1,cz,w,h,d,ch)
		}()
		
	}
	
}

func GetData(qt *QuadTree, ch chan bool) {
	// if it is leaf, get the data from database
	// else get the data from its childtren and then resize
	
	if(qt.dataAvail==true){
		ch <- true
		return;
	}else{
		
		if(qt.isLeaf==true){
			// get data
			
			// save data if not empty
			
		}else{
			// get data from its children's data
			
			go GetData(qt.TL,ch)
			go GetData(qt.TR,ch)
			go GetData(qt.BL,ch)
			go GetData(qt.BR,ch)
			
			// resize
			
			// save data
		}
		
		qt.dataAvail = true;
	}
	
}


