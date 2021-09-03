package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizeController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizeController"],
        beego.ControllerComments{
            Method: "Add",
            Router: "/prize/add",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizeController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/prize/delete",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizeController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizeController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/prize/get",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizeController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizeController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/prize/getAll",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizeController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizeController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/prize/update",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"],
        beego.ControllerComments{
            Method: "Add",
            Router: "/prizePool/add",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"],
        beego.ControllerComments{
            Method: "AddPrize",
            Router: "/prizePool/addPrize",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"],
        beego.ControllerComments{
            Method: "DelPrize4Pool",
            Router: "/prizePool/delPrize",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/prizePool/delete",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"],
        beego.ControllerComments{
            Method: "Draw",
            Router: "/prizePool/draw",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/prizePool/get",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/prizePool/getAll",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"],
        beego.ControllerComments{
            Method: "Info",
            Router: "/prizePool/info",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/prizePool/update",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"],
        beego.ControllerComments{
            Method: "UpdatePrize",
            Router: "/prizePool/updatePrize",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
