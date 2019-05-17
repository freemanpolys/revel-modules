package controllers

import (
	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
	"my-revel-modules/admin/app/models"
)

type DoveaAdminApp struct {
	gormc.TxnController
}

func (c DoveaAdminApp) DbAutoMigrate()  {

}
func (c DoveaAdminApp) Index() revel.Result {
	//c.Txn.LogMode(true)
	c.Txn.AutoMigrate(&models.User{})
	return c.RenderTemplate("DoveaAdminApp/Login/index.html")
}
