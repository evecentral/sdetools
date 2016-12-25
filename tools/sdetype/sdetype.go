package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/evecentral/sdetools"
)

var sdepath = flag.String("sde", "sde", "Path to the SDE root")

func main() {
	flag.Parse()

	sde := sdetools.SDE{
		BaseDir: *sdepath,
	}

	names, ok := sde.GetSystemNameById(30000142)
	if ok != true {
		fmt.Printf("can't load types %v\n", ok)
		return
	}
	fmt.Println(names)
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println(m)
	f, err := os.Create("memp")
	if err != nil {
		return
	}
	pprof.WriteHeapProfile(f)
	f.Close()
	time.Sleep(60 * time.Second)
}
