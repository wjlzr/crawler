package routers

import (
	"crawler/controllers"
	"crawler/controllers/douBan"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //豆瓣网
	beego.Router("/crawlmovie", &douBan.CrawlMovieController{},"*:CrawlMovie")
}
