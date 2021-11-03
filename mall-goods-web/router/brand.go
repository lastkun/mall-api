package router

import (
	"github.com/gin-gonic/gin"
	"mall-api/mall-goods-web/api/brand"
)


func InitBrandRouter(Router *gin.RouterGroup) {
	BrandRouter := Router.Group("brands")
	{
		BrandRouter.GET("", brand.BrandList)
		BrandRouter.DELETE("/:id", brand.DeleteBrand)
		BrandRouter.POST("", brand.NewBrand)
		BrandRouter.PUT("/:id", brand.UpdateBrand)
	}

	CategoryBrandRouter := Router.Group("categorybrands")
	{
		CategoryBrandRouter.GET("", brand.CategoryBrandList)
		CategoryBrandRouter.DELETE("/:id", brand.DeleteCategoryBrand)
		CategoryBrandRouter.POST("", brand.NewCategoryBrand)
		CategoryBrandRouter.PUT("/:id", brand.UpdateCategoryBrand)
		CategoryBrandRouter.GET("/:id", brand.GetCategoryBrandList) //获取分类的品牌
	}
}