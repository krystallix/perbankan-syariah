package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func LoadConf() {
	ConfServerIp = "192.168.1.84"
	ConfServerPort = "8787"
	ConfDbServer = "localhost"
	ConfDbName = "jur_ps"
	ConfDbUser = "admin"
	ConfDbPass = "admin"
}

var ConfServerIp string
var ConfServerPort string
var ConfPortServer string
var ConfDbServer string
var ConfDbName string
var ConfDbUser string
var ConfDbPass string
