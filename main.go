package main

import (
	"log"
	"os"

	"github.com/krasin/stl"
	"github.com/krasin/voxel/nptl"
	"github.com/krasin/voxel/raster"
	"github.com/krasin/voxel/timing"
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
	mesh := raster.STLToMesh(VoxelSide*MeshMultiplier, triangles)
	timing.StopTiming("STLToMesh")

	timing.StartTiming("Rasterize")
	vol := raster.Rasterize(mesh, VoxelSide)
	timing.StopTiming("Rasterize")

	timing.StartTiming("WriteNptl")
	if err = nptl.WriteNptl(vol, mesh.Grid, os.Stdout); err != nil {
		log.Fatalf("WriteNptl: %v", err)
	}
	timing.StopTiming("WriteNptl")
	timing.StopTiming("total")
	timing.PrintTimings(os.Stderr)
}
