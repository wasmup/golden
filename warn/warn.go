package warn

import (
	"log"
	"os"
)

var (
	logger  = log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime)
	Println = logger.Println
	Printf  = logger.Printf
	Print   = logger.Print
)
