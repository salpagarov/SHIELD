package mainframe

import "log"

func Init_API() {
	log.Println("Started API")
}

func init() {
	Init_DB()
}
