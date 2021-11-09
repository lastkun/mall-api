package goods

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mall-api/mall-goods-web/api/base"
	"mall-api/mall-goods-web/forms"
	"mall-api/mall-goods-web/global"
	"mall-api/mall-goods-web/proto"
	"net/http"
	"strconv"
)

func List(ctx *gin.Context) {
	request := &proto.GoodsFilterRequest{}

	priceMin := ctx.DefaultQuery("priceMin", "0")
	priceMinInt, _ := strconv.Atoi(priceMin)
	request.PriceMin = int32(priceMinInt)

	priceMax := ctx.DefaultQuery("priceMax", "0")
	priceMaxInt, _ := strconv.Atoi(priceMax)
	request.PriceMax = int32(priceMaxInt)

	isHot := ctx.DefaultQuery("isHot", "0")
	if isHot == "1" {
		request.IsHot = true
	}
	isNew := ctx.DefaultQuery("isNew", "0")
	if isNew == "1" {
		request.IsNew = true
	}

	isTab := ctx.DefaultQuery("isTab", "0")
	if isTab == "1" {
		request.IsTab = true
	}

	categoryId := ctx.DefaultQuery("categoryId", "0")
	categoryIdInt, _ := strconv.Atoi(categoryId)
	request.TopCategory = int32(categoryIdInt)

	pages := ctx.DefaultQuery("p", "0")
	pagesInt, _ := strconv.Atoi(pages)
	request.Pages = int32(pagesInt)

	perNums := ctx.DefaultQuery("perNums", "0")
	perNumsInt, _ := strconv.Atoi(perNums)
	request.PagePerNums = int32(perNumsInt)

	//TODO  q??
	keywords := ctx.DefaultQuery("q", "")
	request.KeyWords = keywords

	brandId := ctx.DefaultQuery("brandId", "0")
	brandIdInt, _ := strconv.Atoi(brandId)
	request.Brand = int32(brandIdInt)
	resp, e := global.GoodsServiceClient.GoodsList(context.Background(), request)
	if e != nil {
		zap.S().Errorw("[商品api服务] List 出错")
		base.HandleGRPCErrorToHttp(e, ctx)
		return
	}

	gMap := map[string]interface{}{
		"total": resp.Total,
	}

	goodsList := make([]interface{}, 0)
	for _, value := range resp.Data {
		goodsList = append(goodsList, map[string]interface{}{
			"id":          value.Id,
			"name":        value.Name,
			"goods_brief": value.GoodsBrief,
			"desc":        value.GoodsDesc,
			"ship_free":   value.ShipFree,
			"images":      value.Images,
			"desc_images": value.DescImages,
			"front_image": value.GoodsFrontImage,
			"shop_price":  value.ShopPrice,
			"category": map[string]interface{}{
				"id":   value.Category.Id,
				"name": value.Category.Name,
			},
			"brand": map[string]interface{}{
				"id":   value.Brand.Id,
				"name": value.Brand.Name,
				"logo": value.Brand.Logo,
			},
			"is_hot":  value.IsHot,
			"is_new":  value.IsNew,
			"on_sale": value.OnSale,
		})
	}
	gMap["data"] = goodsList
	ctx.JSON(http.StatusOK, gMap)
}

func New(ctx *gin.Context) {
	//表单校验
	form := forms.GoodsForm{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}

	client := global.GoodsServiceClient
	resp, err := client.CreateGoods(context.Background(), &proto.CreateGoodsInfo{
		Name:            form.Name,
		GoodsSn:         form.GoodsSn,
		Stocks:          form.Stocks,
		MarketPrice:     form.MarketPrice,
		ShopPrice:       form.ShopPrice,
		GoodsBrief:      form.GoodsBrief,
		ShipFree:        *form.ShipFree,
		Images:          form.Images,
		DescImages:      form.DescImages,
		GoodsFrontImage: form.FrontImage,
		CategoryId:      form.CategoryId,
		BrandId:         form.Brand,
	})
	if err != nil {
		base.HandleGRPCErrorToHttp(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

//获取商品详情
func Detail(ctx *gin.Context) {
	id := ctx.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		ctx.Status(http.StatusBadRequest)
	}
	client := global.GoodsServiceClient
	result, e := client.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{
		Id: int32(i),
	})
	if e != nil {
		base.HandleGRPCErrorToHttp(e, ctx)
		return
	}

	resp := map[string]interface{}{
		"id":          result.Id,
		"name":        result.Name,
		"goods_brief": result.GoodsBrief,
		"desc":        result.GoodsDesc,
		"ship_free":   result.ShipFree,
		"images":      result.Images,
		"desc_images": result.DescImages,
		"front_image": result.GoodsFrontImage,
		"shop_price":  result.ShopPrice,
		"ctegory": map[string]interface{}{
			"id":   result.Category.Id,
			"name": result.Category.Name,
		},
		"brand": map[string]interface{}{
			"id":   result.Brand.Id,
			"name": result.Brand.Name,
			"logo": result.Brand.Logo,
		},
		"is_hot":  result.IsHot,
		"is_new":  result.IsNew,
		"on_sale": result.OnSale,
	}
	ctx.JSON(http.StatusOK, resp)
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	_, err = global.GoodsServiceClient.DeleteGoods(context.Background(), &proto.DeleteGoodsInfo{Id: int32(i)})
	if err != nil {
		base.HandleGRPCErrorToHttp(err, ctx)
		return
	}

	ctx.Status(http.StatusOK)
	return
}

func Stocks(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	//TODO 库存
	return
}

func UpdateStatus(ctx *gin.Context) {
	form := forms.GoodsStatusForm{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}

	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)

	if _, err = global.GoodsServiceClient.UpdateGoods(context.Background(), &proto.CreateGoodsInfo{
		Id:     int32(i),
		IsHot:  *form.IsHot,
		IsNew:  *form.IsNew,
		OnSale: *form.OnSale,
	}); err != nil {
		base.HandleGRPCErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "修改成功",
	})
}

func Update(ctx *gin.Context) {
	form := forms.GoodsForm{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}

	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if _, err = global.GoodsServiceClient.UpdateGoods(context.Background(), &proto.CreateGoodsInfo{
		Id:              int32(i),
		Name:            form.Name,
		GoodsSn:         form.GoodsSn,
		Stocks:          form.Stocks,
		MarketPrice:     form.MarketPrice,
		ShopPrice:       form.ShopPrice,
		GoodsBrief:      form.GoodsBrief,
		ShipFree:        *form.ShipFree,
		Images:          form.Images,
		DescImages:      form.DescImages,
		GoodsFrontImage: form.FrontImage,
		CategoryId:      form.CategoryId,
		BrandId:         form.Brand,
	}); err != nil {
		base.HandleGRPCErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}
