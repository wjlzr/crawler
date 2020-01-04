package douBan

import (
	"crawler/models/douBan"
	"fmt"
	"strings"

	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"crawler/services/redis"
	"time"
)

type CrawlMovieController struct {
	beego.Controller
}

func (c *CrawlMovieController) CrawlMovie() {
	//关闭模板渲染
	//c.EnableRender = false

	var movieInfo douBan.MovieInfo

	//连接到redis
	redis.ConnectRedis("127.0.0.1:6379","redisp")
	sUrl := "https://movie.douban.com/subject/25827935/"

	//先添加到队列中
	redis.PutinQueue(sUrl)

	for {
		length := redis.GetQueueLength()
		fmt.Println("%i",length)
		if length == 0 {
			break //如果url队列为空 则退出当前循环
		}

		sUrl = redis.PopfromQueue()
		//我们应当判断sUrl是否应该被访问过
		if redis.IsVisit(sUrl) {
			continue
		}
		rsp := httplib.Get(sUrl)
		//设置User-agent以及cookie是为了防止  豆瓣网的 403
		rsp.Header("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64; rv:50.0) Gecko/20100101 Firefox/50.0")
		rsp.Header("Cookie", `bid=gFP9qSgGTfA; __utma=30149280.1124851270.1482153600.1483055851.1483064193.8; __utmz=30149280.1482971588.4.2.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; ll="118221"; _pk_ref.100001.4cf6=%5B%22%22%2C%22%22%2C1483064193%2C%22https%3A%2F%2Fwww.douban.com%2F%22%5D; _pk_id.100001.4cf6=5afcf5e5496eab22.1482413017.7.1483066280.1483057909.; __utma=223695111.1636117731.1482413017.1483055857.1483064193.7; __utmz=223695111.1483055857.6.5.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; _vwo_uuid_v2=BDC2DBEDF8958EC838F9D9394CC5D9A0|2cc6ef7952be8c2d5408cb7c8cce2684; ap=1; viewed="1006073"; gr_user_id=e5c932fc-2af6-4861-8a4f-5d696f34570b; __utmc=30149280; __utmc=223695111; _pk_ses.100001.4cf6=*; __utmb=30149280.0.10.1483064193; __utmb=223695111.0.10.1483064193`)

		sMovieHtml, err := rsp.String()
		if err != nil {
			logs.Error(err)
		}

		movieInfo.Movie_name = douBan.GetMovieName(sMovieHtml)
		if movieInfo.Movie_name != "" {

			movieInfo.Movie_name = douBan.GetMovieName(sMovieHtml)
			movieInfo.Movie_director = douBan.GetMovieDirector(sMovieHtml)
			movieInfo.Movie_main_character = douBan.GetMovieMainCharacters(sMovieHtml)
			movieInfo.Movie_type = douBan.GetMovieGenre(sMovieHtml)
			movieInfo.Movie_on_time = douBan.GetMovieOnTime(sMovieHtml) // 上映时间：2016-09-14(中国大陆)
			movieInfo.Movie_grade = douBan.GetMovieGrade(sMovieHtml)
			movieInfo.Movie_span = douBan.GetMovieRunningTime(sMovieHtml)
			movieInfo.Movie_pic = douBan.GetMovieImg(sMovieHtml)

			movieInfo.Movie_on_time = movieInfo.Movie_on_time[0:strings.Index(douBan.GetMovieOnTime(sMovieHtml), "(")] //上映时间：2016-09-14
			movieInfo.Create_time = time.Now().Format("2006-1-2 15:04:05")

			douBan.AddMovie(&movieInfo)

		}

		//提取该页面的所有连接
		urls := douBan.GetMovieUrls(sMovieHtml)

		for _, url := range urls {
			redis.PutinQueue(url)
			c.Ctx.WriteString("<br>" + url + "</br>")
		}

		//sUrl 应当记录到 访问set中
		redis.AddToSet(sUrl)

		time.Sleep(1000000)
	}
	c.Ctx.WriteString("end of crawl!")
}
