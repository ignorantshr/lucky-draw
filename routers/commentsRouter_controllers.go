package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizeController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizeController"],
        beego.ControllerComments{
            Method: "Add",
            Router: "/lucky-draw/prize/add",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizeController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/lucky-draw/prize/delete",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizeController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizeController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/lucky-draw/prize/get",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizeController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizeController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/lucky-draw/prize/getAll",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizeController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizeController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/lucky-draw/prize/update",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"],
        beego.ControllerComments{
            Method: "Add",
            Router: "/lucky-draw/prizePool/add",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"],
        beego.ControllerComments{
            Method: "AddPrize",
            Router: "/lucky-draw/prizePool/addPrize",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"],
        beego.ControllerComments{
            Method: "DelPrize4Pool",
            Router: "/lucky-draw/prizePool/delPrize",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/lucky-draw/prizePool/delete",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"],
        beego.ControllerComments{
            Method: "Draw",
            Router: "/lucky-draw/prizePool/draw",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/lucky-draw/prizePool/get",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/lucky-draw/prizePool/getAll",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"],
        beego.ControllerComments{
            Method: "GetUnpoolPrizes",
            Router: "/lucky-draw/prizePool/getUnpoolPrizes",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"],
        beego.ControllerComments{
            Method: "Info",
            Router: "/lucky-draw/prizePool/info",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/lucky-draw/prizePool/update",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"] = append(beego.GlobalControllerRouter["lucky-draw/controllers:PrizePoolController"],
        beego.ControllerComments{
            Method: "UpdatePrize",
            Router: "/lucky-draw/prizePool/updatePrize",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
