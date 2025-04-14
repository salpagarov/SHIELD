package mainframe

import "log"

func db_init() {
	log.Println("DB connected")
}

func db_stop() {
	log.Println("DB disconnected")
}

func Update() {
	db_init()
	log.Println("DB updated")
	db_stop()
}
