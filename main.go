package main

import (
	"flag"
	"fmt"
	"imaging/filter"
	"imaging/task"
	"time"
)

func main() {

	srcdir := flag.String("src", "imgs", "Input directory")
	dstdir := flag.String("dst", "output", "Output directory")
	filterType := flag.String("filter", "grayscale", "grayscale/blur")
	taskType := flag.String("task", "waitgrp", "waitgrp/channel")
	poolsize := flag.Int("poolsize", 4, "Workers pool size for the channel task")
	flag.Parse()

	var f filter.Filter

	switch *filterType {
	case "grayscale":
		f = &filter.GrayscaleFilter{}
	case "blur":
		f = &filter.BlurFilter{}
	default:
		fmt.Println("Invalid filter type")
		return
	}

	var t task.Tasker

	switch *taskType {
	case "waitgrp":
		t = task.NewWaitGrpTask(*srcdir, *dstdir, f)
	case "channel":
		t = task.NewChanTask(*srcdir, *dstdir, f, *poolsize)
	default:
		fmt.Println("Invalid task type")
		return
	}

	start := time.Now()

	err := t.Process()
	if err != nil {
		fmt.Printf("Error during image processing: %s\n", err)
		return
	}

	elapsed := time.Since(start)
	fmt.Printf("Image processing took %s\n", elapsed)
}
