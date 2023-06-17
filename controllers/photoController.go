package controllers

import (
	"fmt"
	"go-myGram/database"
	"go-myGram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPhoto(ctx *gin.Context) {
	db := database.GetDB()

	var photos []models.Photo

	err := db.Model(&models.Photo{}).Find(&photos).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"photos": photos,
		"code":   200,
	})
}

func GetOnePhoto(ctx *gin.Context) {
	db := database.GetDB()

	var (
		photoId string
		photo   models.Photo
	)

	photoId = ctx.Param("photoId")

	if err := db.First(&models.Photo{}, "id = ?", photoId).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("gagal mengambil data photo dengan id %v karna data tidak di temukan", photoId),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"photo": photo,
		"code":  200,
	})
}

func CreatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	var photo models.Photo

	if err := ctx.ShouldBindJSON(&photo); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	err := db.Debug().Create(&photo).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"photo": photo,
		"code":  200,
	})
}

func UpdatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	var photo models.Photo
	id := ctx.Query("id")

	if err := db.First(&models.Photo{}, "id = ?", id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("gagal Update order Id %v tidak di temukan", id),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&photo); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	err := db.Model(&models.Photo{}).Where("id = ?", id).Updates(photo).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"photo": photo,
		"code":  200,
	})
}

func DeletePhoto(ctx *gin.Context) {
	db := database.GetDB()
	photo := models.Photo{}

	id := ctx.Query("id")

	if err := db.First(&photo, "order_id = ?", id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Gagal menghapus photo id %v tidak di temukan", id),
		})
		return
	}

	if err := db.Delete(&photo).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("photo dengan id %v berhasil dihapus", id),
		"code":    200,
	})
}
