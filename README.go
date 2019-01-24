# vlog
golang  log reconsitution--重构log日志

example::

package main

import (
	"vlog"
)

func main() {
	vlog.SetFlags(vlog.Ldate | vlog.Ltime | vlog.Lshortfile)
	vlog.SetFilepath("lo_test.log")   //save file
	vlog.SetFileMax(20 * 1024 * 1024) //20M
	vlog.Println("test log1........")
	vlog.Println("test log2........")
	vlog.Println("test log3........")
	vlog.Println("test log4........")
	vlog.Println("test log5........")
}


