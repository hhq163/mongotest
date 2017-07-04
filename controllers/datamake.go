package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

type DataMakeController struct {
	BaseController
}

type GameRoundInfo struct {
	StartTime      time.Time
	CalculatedFlag int16
	Valid_flag     int16
	ServerId       int32
	GameId         int32
	GameCatId      int32
	GamePeriod     string
	TableNo        string
	RoundNo        string
	Remark         string
	GameResult     string
	State          int16
}

type UserChartInfo struct {
	GameRoundId        string
	UserId             int32
	UserName           string
	UserRank           int16
	CharId             int16
	TotalBetScore      float64
	TotalWinScore      float64
	ValidWinScore      float64
	OperatorWinScore   float64
	HallId             int32
	AgentId            int32
	GameId             int32
	CatId              int32
	TableNo            int32
	ServerId           int32
	StartTime          time.Time
	OccupancyRate      float64
	OccupancyRateHall  float64
	OccupancyRateAgent float64
	ServerName         string
	IsCancel           int16
	RoundNo            string
	GamePeriod         string
	DwRound            int32
	Remark             string
	Account            string
	Encry              string
	IsMark             int16
	GameHallId         int32
	GameHallTitle      string
	LoginType          int16
	GameCatCode        string
	GameHallCode       string
	HallName           string
	AgentName          string
	GameName           string
	GameResult         string
}

/**
* 往game_round_info集合中造测试数据
**/
func (this *DataMakeController) Get() {

	db_name := beego.AppConfig.String("mongodb_dbname")

	dialInfo := &mgo.DialInfo{
		Addrs:     []string{beego.AppConfig.String("mongodb_url")},
		Timeout:   time.Second * 2,
		Database:  beego.AppConfig.String("mongodb_dbname"),
		Username:  beego.AppConfig.String("mongodb_user"),
		Password:  beego.AppConfig.String("mongodb_pass"),
		Mechanism: beego.AppConfig.String("mechanism_type"),
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if nil != err {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	db := session.DB(db_name)             //数据库名称
	collection := db.C("game_round_info") //如果该集合已经存在的话，则直接返回

	for i := 1; i <= 50; i++ {
		fmt.Println("insert:" + strconv.Itoa(i))
		stime := "06/0" + strconv.Itoa(i/20+1) + "/2017"
		fmt.Println("stime:" + stime)
		tm2, _ := time.Parse("01/02/2006", stime)

		round_no := this.GenerateMixed(16)
		ground_id := this.GenerateMixed(20)

		//往game_round_info集合中造测试数据
		err = collection.Insert(&GameRoundInfo{
			tm2,
			1,
			1,
			24,
			97,
			1,
			"23-41",
			"V2-T02",
			round_no,
			"12,57,0;28,35,0",
			"9;3",
			0})
		if err != nil {
			panic(err)
		}

		collection := db.C("user_chart_info") //如果该集合已经存在的话，则直接返回
		//往user_chart_info集合中造测试数据
		err = collection.Insert(&UserChartInfo{
			ground_id,
			int32(1601 + i),
			"TEST504384" + strconv.Itoa(1601+i),
			1,
			2,
			17500.0,
			0.0,
			0.0,
			-0.0,
			1,
			2,
			91,
			1,
			0,
			9,
			tm2,
			0.0,
			0.0,
			0.0,
			"10",
			0,
			round_no,
			"259-61",
			61,
			"53,4,0;41,43,0",
			"a9TEST504384",
			"Mzc2ODA4LDE3NTAwLDA=",
			1,
			0,
			"10",
			1,
			"GC0001",
			"GH0001",
			"csj",
			"agent_test",
			"百家乐",
			"9;9",
		})

		if err != nil {
			panic(err)
		}

	}

	this.Data["result"] = "正在进行数据插入操作！"
	this.TplName = "datamake.tpl"

}

/**
* 默认提示内容
**/
type WelcomeController struct {
	BaseController
}

func (this *WelcomeController) Get() {

	fmt.Println("欢迎来到测试数据插入页面！")
	this.Data["welcome"] = "测试数据插入默认页面"

	//	beego.Debug("欢迎来到测试数据插入页面！")
	this.TplName = "welcome.tpl"
}
