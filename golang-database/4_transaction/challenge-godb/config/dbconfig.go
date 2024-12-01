package config

import "fmt"

const (
	host     = "unthenamed-unthenamed.h.aivencloud.com"
	port     = 13689
	user     = "avnadmin"
	password = "AVNS_DilIFOqbMOlCyx0DADG"
	dbname   = "enigma_laundry"
	sslmode  = "require"
)

var SqlInfo string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
