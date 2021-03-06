// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package mall_goods

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"time"
)

// Fill with you ideas below.
type ResCateItem struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}
type ResCateList struct {
	ResGdataBase
	Data []ResCateItem `json:"data"`
}

type ResGdataList struct {
	ResGdataBase
	Data []GoodsId `json:"data"`
}
type ResGdataBase struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Page    int    `json:"page"`
}

type GoodsId struct {
	Id int64 `json:"id"`
}

//获取分类下的商品
func GetGoodsList(r *ghttp.Request) {
	glog.Println("getGoodsList")

	//var catMap = make(map[string]int,0)
	var catselice = make([]int64, 0)
	//获取分类
	var catPath = fmt.Sprintf("https://api.jiafuminkang.com/api/Category/getCatelist?type=1")
	response1, err := ghttp.Get(catPath) // GET请求
	if err != nil {
		panic(err)
	}
	defer response1.Close()
	dataCateStr := response1.ReadAll()
	var resCatep ResCateList
	err = json.Unmarshal(dataCateStr, &resCatep)

	for _, v := range resCatep.Data {
		//catMap[v.Name] = v.Id
		catselice = append(catselice, v.Id)
	}

	//爬回来的商品ids
	ids := make([]int64, 0)

	for _, cateId := range catselice {
		//catMap[v.Name] = v.Id

		//var cateId uint64 = 10
		//请分类下的商品列表
		var path = fmt.Sprintf("https://api.jiafuminkang.com/api/Category/getCategoryName?pid=%d&page=1&sessionId=&goods_type=1", cateId)
		response, err := ghttp.Get(path) // GET请求
		if err != nil {
			panic(err)
		}
		defer response.Close()
		//g.Log().Line().Info(response)
		//g.Log().Line().Info(response.ReadAllString())
		dataStr := response.ReadAll()
		var resp ResGdataList
		err = json.Unmarshal(dataStr, &resp)
		if err != nil {
			//g.Log().Line().Error(err)
			r.Response.WriteJson(g.Map{
				"code": 400,
				"msg":  "json.Unmarshal GoodsIdList error :" + err.Error(),
			})
			r.Exit()
		}
		//g.Log().Line().Info(resp.Data)
		//jobs := make(chan int64, len(ids)) //要做的任务
		//results := make(chan int64, len(ids)) //结果
		if len(resp.Data) <= 0 {
			continue
		}
		for _, v := range resp.Data {
			ids = append(ids, v.Id)
		}

		// 存入任务
		for _, goodsId := range ids {
			go DoGetGoodsDetail(cateId, goodsId)
		}
	}
	r.Response.WriteJson(g.Map{
		"code": 0,
		"msg":  "ok",
		"data": ids,
	})
	r.Exit()
}

type ResGdataDetail struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Page int           `json:"page"`
	Data RespGoodsData `json:"data"`
}

type IsActArr struct {
	Act     uint   `json:"act"`
	ActDesc string `json:"actDesc"`
}
type GoodsSlidsUrlList struct {
	Path string `json:"path"`
}
type RespGoodsData struct {
	ProductName        string              `json:"product_name"`
	ProductImage       string              `json:"product_image"`
	ProductPrice       string              `json:"product_price"`
	ProductMarketPrice string              `json:"product_market_price"`
	Promotion          string              `json:"promotion"`
	ProductDesc        string              `json:"product_desc"`
	ProductContent1    string              `json:"product_content1"`
	SlidsUrlList       []GoodsSlidsUrlList `json:"slids_url_list"`
	IsAct              []IsActArr          `json:"isAct"`
	Video              string              `json:"video"`
}

//api请求商品的详情
func GetGoodsDetail(r *ghttp.Request) {
	glog.Println("GetGoodsDetail")
	var cateid int64 = 7
	var goodsId int64 = 14267
	go DoGetGoodsDetail(cateid, goodsId)

	r.Response.WriteJson(g.Map{
		"code": 0,
		"msg":  "ok",
	})
	r.Exit()
}

//获取商品详情
func DoGetGoodsDetail(cateid int64, goodsId int64) {
	glog.Println("GetGoodsDetail")
	var path = fmt.Sprintf("https://api.jiafuminkang.com/app/Product/bulkDetailXiao?goods_id=%d&mer_id=&sessionId=", goodsId)
	response, err := ghttp.Get(path) // GET请求
	if err != nil {
		panic(err)
	}

	defer response.Close()
	//g.Log().Line().Info(response)
	dataStr := response.ReadAll()
	//g.Log().Line().Info(dataStr)

	//var resp ResGdata
	//var resp interface{}

	var respGD ResGdataDetail
	err = json.Unmarshal(dataStr, &respGD)
	if err != nil {
		return
		//g.Log().Line().Error(err)
		//r.Response.WriteJson(g.Map{
		//	"code": 400,
		//	"msg":  "json.Unmarshal GoodsIdList error :" + err.Error(),
		//})
		//r.Exit()
	}
	g.Log().Line().Info(respGD)
	var tags = ""
	for _, v := range respGD.Data.IsAct {
		if tags == "" {
			tags = v.ActDesc
		} else {
			tags += "," + v.ActDesc
		}
	}
	imageList := ""
	for _, v := range respGD.Data.SlidsUrlList {
		if imageList == "" {
			imageList = v.Path
		} else {
			imageList += "|" + v.Path
		}
	}
	t := time.Now().Unix()
	var saveData Entity
	saveData.ShopId = 1 // 店铺id
	//saveData.Type = 0                               // 商品类型:0=国内商品;1=海外商品
	//saveData.TypeTags = 0                           // 商品类型标签:0=无;1=保税;2=直邮;3=一般贸易
	//saveData.BuyType = 0                            // 购买类型:0=全部;1=微信;2=点卡
	saveData.BrandId = 236       // 品牌ID
	saveData.CateId = cateid + 1 // 商品分类id
	//saveData.UnitId = 0                             // 商品单位ID
	saveData.SpecId = 4                                                 // 规格ID
	saveData.TagsId = "11"                                              // 商品标签ID
	saveData.Tags = tags                                                // 商品展示小标签,输入请使用英文逗号隔开
	saveData.GoodsTitle = respGD.Data.ProductName                       // 商品名
	saveData.GoodsStock = 10000                                         // 商品总库存
	saveData.GoodsMinPrice = gconv.Float64(respGD.Data.ProductPrice)    // 最低售价
	saveData.GoodsPrice = gconv.Float64(respGD.Data.ProductMarketPrice) // 展示价
	saveData.GoodsSale = 100                                            // 销售数量
	saveData.GoodsLook = 1000                                           // 浏览量
	saveData.GoodsContent = respGD.Data.ProductContent1                 // 商品内容
	saveData.GoodsLogo = respGD.Data.ProductImage                       // 商品LOGO
	saveData.GoodsImage = imageList                                     // 商品图片地址
	saveData.GoodsVideo = respGD.Data.Video                             // 商品视频URL
	saveData.GoodsDesc = respGD.Data.Promotion                          // 商品描述
	saveData.IsAudit = 1                                                // 是否审核:0=待审核;1=已审核;2=不通过
	//saveData.AuditText     =   // 审核内容
	saveData.AuditAdmin = "10000"   // 审核人
	saveData.AuditTime = 1594200710 // 审核时间
	//saveData.IsShare = 0            // 是否共享:0=不共享;1=共享
	//saveData.ShareTj = 0            // 是否共享推荐
	//saveData.TcShare = 0            // 同城共享0不限1是
	//saveData.IsSpecial = 0          // 是否特价:0=非特价;1=特价
	//saveData.Integral = 0           // 特价消耗积分
	//saveData.OffsetPrice = 0        // 积分抵扣的金额
	//saveData.Status = 0             // 商品状态:0=销售中;1=已下架
	//saveData.IsTj = 1               // 是否推荐:0=推荐;1=不推荐
	//saveData.IsDeleted = 0          // 删除状态:0=未删除;1=删除
	saveData.ITime = t // 创建时间
	//saveData.IAdmin        =   // 添加用户
	saveData.UTime = t // 更新时间
	//saveData.UAdmin        =   // 更新用户

	goodsInsertRes, err := saveData.Save()
	if err != nil {
		return
	}
	newGoodsId, _ := goodsInsertRes.LastInsertId()

	//插入扩展信息
	extendData := g.Map{
		`shop_id`:       1,                      //'店铺ID',
		`goods_id`:      newGoodsId,             //'商品ID',
		`spec_id`:       4,                      // '规格ID',
		`specs`:         ",17,",                 // '规格属性信息',
		`show_price`:    saveData.GoodsPrice,    //'商品展示价格',
		`selling_price`: saveData.GoodsMinPrice, // '商品销售价格',
		`stock`:         100,                    // '商品库存',
		`integral`:      0,                      // '积分',
		`shop_rebate`:   0,                      // '店铺分销返利',
		`user_rebate`:   0,                      // '用户分销返利',
		`status`:        1,                      //'是否销售:0=禁用,1=启用',
		`is_deleted`:    0,                      //  '删除状态:0=未删除,1=删除',
	}

	g.DB().Table("mall_goods_extend").Data(extendData).Save()

	successMsg := fmt.Sprintf("商品%d已经存入完成,新goodsId%d", goodsId, newGoodsId)
	g.Log().Line().Info(successMsg)

}
