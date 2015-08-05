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
	
	TL		*QuadTree // 00  0
	TR		*QuadTree // 01  1
	BL		*QuadTree // 10  2
	BR		*QuadTree // 11  3
	
	node 		*node.Node
	
}

type empty struct {}
type semaphore chan empty

// VoxelSize, MinPoint, MaxPoint
func (qt *QuadTree) Init(xmin, xmax, ymin, ymax, zmin, zmax int64, resx, resy, resz float64) []*QuadTree {
		
	dimz := zmax - zmin + 1
	
	qtlist := make([]*QuadTree, dimz)	
	for i := range qtlist {
    	qtlist[i] = new(QuadTree)
	}

	dimx := xmax - xmin + 1
	dimy := ymax - ymin + 1
	
	depthx := int(math.Log2(float64(dimx))+0.5)-int(math.Log2(float64(qtW)))+1
	depthy := int(math.Log2(float64(dimy))+0.5)-int(math.Log2(float64(qtH)))+1
	
	depth := int(math.Max(float64(depthx), float64(depthy)))
	
	fmt.Println("The depth of this quadtree is ", depth)
	
	sem := make (semaphore, dimz);
	for z := zmin; z < zmax; z++ {
		ch := make(chan bool)
        go func (z int64) {
            qtlist[z].Construct(nil,0,depth,-1,xmin,ymin,z,xmax,ymax,resx,resy,resz,0,0,0,qtW,qtH,1,ch) 
            sem <- empty{};
			
			<-ch
			
        } (z);
    }
    for z := zmin; z < zmax; z++ {
		<- sem // release dimz resources
	}
	
	fmt.Printf("~~~parent's children %v %v %v %v \n", qtlist[0].TL, qtlist[0].TR, qtlist[0].BL, qtlist[0].BR)
	
	return qtlist
}

func (qt *QuadTree) Construct(parent *QuadTree, child,depth,level int, xmin,ymin,zmin,xmax,ymax int64, resx,resy,resz float64, cx,cy,cz,w,h,d int64, ch chan bool) {
	
	depth = depth - 1
	level = level + 1
	
	fmt.Printf("Construct depth %v level %v cx %v cy %v\n",depth, level, cx, cy)
	
	if(depth < 0 || xmin>xmax || ymin>ymax){
		ch <- true
		return
		
	}else{
		
		qt = &QuadTree{depth,level,false,false,parent,nil,nil,nil,nil,nil}
		
		fmt.Printf("current tile %v \n",qt)
		
		qt.node = new(node.Node)
		qt.node.NewNode(cx,cy,cz,w,h,d,resx,resy,resz,xmin,ymin,zmin,nil)
		
		qt.depth = depth
		
		qt.dataAvail = false
		
		if(depth==0){
			qt.isLeaf = true
		}else{
			qt.isLeaf = false
		}
		
		qt.parent = parent
		
		if(parent!=nil){
						
			if(child==0){
				fmt.Println("case 0")
				qt.parent.TL = qt
			}else if(child==1){
				fmt.Println("case 1")
				qt.parent.TR = qt
			}else if(child==2){
				fmt.Println("case 2")
				qt.parent.BL = qt
			}else if(child==3){
				fmt.Println("case 3")
				qt.parent.BR = qt
			}else
			{
				fmt.Println("Invalid child", child)
				ch <- false
				return
			}
			
			fmt.Printf("parent's children %v %v %v %v \n", parent.TL, parent.TR, parent.BL, parent.BR)
			
		}
		
		resx = resx / 2.0
		resy = resy / 2.0
		resz = resz / 2.0
		
		cx = cx * 2
		cy = cy * 2
		cz = cz * 2

		go func() {
			qt.TL.Construct(qt,0,depth,level,xmin,  ymin,  zmin,xmax,ymax,resx,resy,resz,cx,  cy,  cz,w,h,d,ch)
			qt.TR.Construct(qt,1,depth,level,xmin+w,ymin,  zmin,xmax,ymax,resx,resy,resz,cx+1,cy,  cz,w,h,d,ch)
			qt.BL.Construct(qt,2,depth,level,xmin,  ymin+h,zmin,xmax,ymax,resx,resy,resz,cx,  cy+1,cz,w,h,d,ch)
			qt.BR.Construct(qt,3,depth,level,xmin+w,ymin+h,zmin,xmax,ymax,resx,resy,resz,cx+1,cy+1,cz,w,h,d,ch)
		}()
		
	}
	
}

func (qt *QuadTree) GetData(ch chan bool) {
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
			
			go qt.TL.GetData(ch)
			go qt.TR.GetData(ch)
			go qt.BL.GetData(ch)
			go qt.BR.GetData(ch)
			
			// resize
			
			// save data
		}
		
		qt.dataAvail = true;
	}
	
}


