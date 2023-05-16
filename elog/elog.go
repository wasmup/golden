package elog

import (
	"log"
	"os"
)

var (
	logger  = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime)
	Println = logger.Println
	Printf  = logger.Printf
	Print   = logger.Print
)
