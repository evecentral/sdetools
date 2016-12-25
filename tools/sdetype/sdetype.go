package main

import (
	"flag"
	"fmt"

	"github.com/evecentral/sdetools"
)

var sdepath = flag.String("sde", "sde", "Path to the SDE root")

// Build a BoltDB of all of the relevant SDE items
func main() {
	flag.Parse()

	sde := sdetools.SDE{
		BaseDir: *sdepath,
	}
	sde.Init()
	sde.BuildBoltDB()

	names, ok := sde.GetSystemNameById(30000142)
	if ok != true {
		fmt.Printf("can't load types %v\n", ok)
		return
	}
	fmt.Println(names)
	groups, ok := sde.GetGroupById(8)
	if ok != true {
		fmt.Printf("can't load groups\n")
		return
	}
	fmt.Println(groups)
}
