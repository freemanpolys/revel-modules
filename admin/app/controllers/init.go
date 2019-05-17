package controllers

import (
	"github.com/revel/revel"
	"log"
)

func init() {
	//revel.InterceptMethod(App.AddUser, revel.)

	revel.RegisterModuleInit(func(module *revel.Module){
		DoveaAdminApp{}.DbAutoMigrate()
		log.Print("=== Module Admin migration ========")
	})

}
