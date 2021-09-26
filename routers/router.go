package routers

import (
	"lucky-draw/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Include(&controllers.PrizeController{})
	beego.Include(&controllers.PrizePoolController{})
}
