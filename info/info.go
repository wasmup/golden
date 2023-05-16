package info

import (
	"log"
	"os"
)

var (
	logger  = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
	Println = logger.Println
	Printf  = logger.Printf
	Print   = logger.Print
)
