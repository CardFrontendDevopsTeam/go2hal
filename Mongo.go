package main

import (
	"gopkg.in/mgo.v2"
 )

func StartDB(){
	mgo.Dial("localhost")
}


