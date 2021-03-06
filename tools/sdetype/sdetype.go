package main

import (
	"flag"
	"fmt"

	"github.com/evecentral/sdetools"
)

var sdepath = flag.String("sde", "sde", "Path to the SDE root")
var conv = flag.Bool("convert", false, "Convert to BoltDB")

// Build a BoltDB of all of the relevant SDE items
func main() {
	flag.Parse()

	sde := sdetools.SDE{
		BaseDir: *sdepath,
	}
	sde.Init()
	if *conv {
		sde.BuildBoltDB()
	}

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

	mt, ok := sde.GetTypeById(34)
	if ok != true {
		fmt.Printf("can't load typeids\n")
		return
	}
	fmt.Println(mt)
	mt, ok = sde.GetTypeByExactName("Zealot")
	if ok != true {
		fmt.Printf("can't load typeids\n")
		return
	}
	fmt.Println(mt)

}
