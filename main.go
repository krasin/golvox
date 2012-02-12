package main

import (
	"log"
	"os"

	"github.com/krasin/stl"
	//	"github.com/krasin/voxel/set"
	"github.com/krasin/voxel/timing"
	//	"github.com/krasin/voxel/triangle"
	//	"github.com/krasin/voxel/volume"
)

const (
	VoxelSide      = 512
	MeshMultiplier = 2048
)

func main() {
	timing.StartTiming("total")
	timing.StartTiming("Read STL from Stdin")
	triangles, err := stl.ReadSTL(os.Stdin)
	if err != nil {
		log.Fatalf("ReadSTL: %v", err)
	}
	timing.StopTiming("Read STL from Stdin")

	timing.StartTiming("STLToMesh")
	mesh := STLToMesh(VoxelSide*MeshMultiplier, triangles)
	timing.StopTiming("STLToMesh")

	timing.StartTiming("Rasterize")
	vol := Rasterize(mesh, VoxelSide)
	timing.StopTiming("Rasterize")

	timing.StartTiming("WriteNptl")
	if err = WriteNptl(vol, mesh.Grid, os.Stdout); err != nil {
		log.Fatalf("WriteNptl: %v", err)
	}
	timing.StopTiming("WriteNptl")
	timing.StopTiming("total")
	timing.PrintTimings(os.Stderr)
}
