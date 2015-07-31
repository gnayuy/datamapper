package octree

import(
	"testing"
	"fmt"
	"os"
	"runtime/debug"
	
)

func test(t *testing.T){
	
	version := 1.0
	t.Error("Expected 1.0, got ", version)
	
}