package brand

import (
	"context"
	"mall-api/mall-goods-web/api/base"
	"mall-api/mall-goods-web/forms"
	"mall-api/mall-goods-web/global"
	"mall-api/mall-goods-web/proto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BrandList(ctx *gin.Context) {
	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)

	resp, err := global.GoodsServiceClient.BrandList(context.Background(), &proto.BrandFilterRequest{
		Pages:       int32(pnInt),
		PagePerNums: int32(pSizeInt),
	})

	if err != nil {
		base.HandleGRPCErrorToHttp(err, ctx)
		return
	}

	result := make([]interface{}, 0)
	reMap := make(map[string]interface{})
	reMap["total"] = resp.Total
	for _, value := range resp.Data[pnInt : pnInt*pSizeInt+pSizeInt] {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["name"] = value.Name
		reMap["logo"] = value.Logo

		result = append(result, reMap)
	}

	reMap["data"] = result

	ctx.JSON(http.StatusOK, reMap)
}

func NewBrand(ctx *gin.Context) {
	brandForm := forms.BrandForm{}
	if err := ctx.ShouldBindJSON(&brandForm); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}

	resp, err := global.GoodsServiceClient.CreateBrand(context.Background(), &proto.BrandRequest{
		Name: brandForm.Name,
		Logo: brandForm.Logo,
	})
	if err != nil {
		base.HandleGRPCErrorToHttp(err, ctx)
		return
	}

	request := make(map[string]interface{})
	request["id"] = resp.Id
	request["name"] = resp.Name
	request["logo"] = resp.Logo

	ctx.JSON(http.StatusOK, request)
}

func DeleteBrand(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	_, err = global.GoodsServiceClient.DeleteBrand(context.Background(), &proto.BrandRequest{Id: int32(i)})
	if err != nil {
		base.HandleGRPCErrorToHttp(err, ctx)
		return
	}

	ctx.Status(http.StatusOK)
}

func UpdateBrand(ctx *gin.Context) {
	brandForm := forms.BrandForm{}
	if err := ctx.ShouldBindJSON(&brandForm); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}

	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	_, err = global.GoodsServiceClient.UpdateBrand(context.Background(), &proto.BrandRequest{
		Id:   int32(i),
		Name: brandForm.Name,
		Logo: brandForm.Logo,
	})
	if err != nil {
		base.HandleGRPCErrorToHttp(err, ctx)
		return
	}
	ctx.Status(http.StatusOK)
}

func GetCategoryBrandList(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	resp, err := global.GoodsServiceClient.GetCategoryBrandList(context.Background(), &proto.CategoryInfoRequest{
		Id: int32(i),
	})
	if err != nil {
		base.HandleGRPCErrorToHttp(err, ctx)
		return
	}

	result := make([]interface{}, 0)
	for _, value := range resp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["name"] = value.Name
		reMap["logo"] = value.Logo

		result = append(result, reMap)
	}

	ctx.JSON(http.StatusOK, result)
}

func CategoryBrandList(ctx *gin.Context) {
	resp, err := global.GoodsServiceClient.CategoryBrandList(context.Background(), &proto.CategoryBrandFilterRequest{})
	if err != nil {
		base.HandleGRPCErrorToHttp(err, ctx)
		return
	}
	reMap := map[string]interface{}{
		"total": resp.Total,
	}

	result := make([]interface{}, 0)
	for _, value := range resp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["category"] = map[string]interface{}{
			"id":   value.Category.Id,
			"name": value.Category.Name,
		}
		reMap["brand"] = map[string]interface{}{
			"id":   value.Brand.Id,
			"name": value.Brand.Name,
			"logo": value.Brand.Logo,
		}

		result = append(result, reMap)
	}

	reMap["data"] = result
	ctx.JSON(http.StatusOK, reMap)
}

func NewCategoryBrand(ctx *gin.Context) {
	categoryBrandForm := forms.CategoryBrandForm{}
	if err := ctx.ShouldBindJSON(&categoryBrandForm); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}

	resp, err := global.GoodsServiceClient.CreateCategoryBrand(context.Background(), &proto.CategoryBrandRequest{
		CategoryId: int32(categoryBrandForm.CategoryId),
		BrandId:    int32(categoryBrandForm.BrandId),
	})
	if err != nil {
		base.HandleGRPCErrorToHttp(err, ctx)
		return
	}

	response := make(map[string]interface{})
	response["id"] = resp.Id

	ctx.JSON(http.StatusOK, response)
}

func UpdateCategoryBrand(ctx *gin.Context) {
	categoryBrandForm := forms.CategoryBrandForm{}
	if err := ctx.ShouldBindJSON(&categoryBrandForm); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}

	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	_, err = global.GoodsServiceClient.UpdateCategoryBrand(context.Background(), &proto.CategoryBrandRequest{
		Id:         int32(i),
		CategoryId: int32(categoryBrandForm.CategoryId),
		BrandId:    int32(categoryBrandForm.BrandId),
	})
	if err != nil {
		base.HandleGRPCErrorToHttp(err, ctx)
		return
	}
	ctx.Status(http.StatusOK)
}

func DeleteCategoryBrand(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	_, err = global.GoodsServiceClient.DeleteCategoryBrand(context.Background(), &proto.CategoryBrandRequest{Id: int32(i)})
	if err != nil {
		base.HandleGRPCErrorToHttp(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, "")
}
