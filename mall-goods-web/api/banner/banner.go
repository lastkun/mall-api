package banner

import (
	"context"
	"mall-api/mall-goods-web/api/base"
	"mall-api/mall-goods-web/forms"
	"mall-api/mall-goods-web/global"
	"mall-api/mall-goods-web/proto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
)

func List(ctx *gin.Context) {
	resp, err := global.GoodsServiceClient.BannerList(context.Background(), &empty.Empty{})
	if err != nil {
		base.HandleGRPCErrorToHttp(err, ctx)
		return
	}

	result := make([]interface{}, 0)
	for _, value := range resp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["index"] = value.Index
		reMap["image"] = value.Image
		reMap["url"] = value.Url

		result = append(result, reMap)
	}

	ctx.JSON(http.StatusOK, result)
}

func New(ctx *gin.Context) {
	bannerForm := forms.BannerForm{}
	if err := ctx.ShouldBindJSON(&bannerForm); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}

	resp, err := global.GoodsServiceClient.CreateBanner(context.Background(), &proto.BannerRequest{
		Index: int32(bannerForm.Index),
		Url:   bannerForm.Url,
		Image: bannerForm.Image,
	})
	if err != nil {
		base.HandleGRPCErrorToHttp(err, ctx)
		return
	}

	response := make(map[string]interface{})
	response["id"] = resp.Id
	response["index"] = resp.Index
	response["url"] = resp.Url
	response["image"] = resp.Image

	ctx.JSON(http.StatusOK, response)
}

func Update(ctx *gin.Context) {
	bannerForm := forms.BannerForm{}
	if err := ctx.ShouldBindJSON(&bannerForm); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}

	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	_, err = global.GoodsServiceClient.UpdateBanner(context.Background(), &proto.BannerRequest{
		Id:    int32(i),
		Index: int32(bannerForm.Index),
		Url:   bannerForm.Url,
	})
	if err != nil {
		base.HandleGRPCErrorToHttp(err, ctx)
		return
	}

	ctx.Status(http.StatusOK)
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	_, err = global.GoodsServiceClient.DeleteBanner(context.Background(), &proto.BannerRequest{Id: int32(i)})
	if err != nil {
		base.HandleGRPCErrorToHttp(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, "")
}
