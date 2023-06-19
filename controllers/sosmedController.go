package controllers

import (
	"fmt"
	"go-myGram/database"
	"go-myGram/helpers"
	"go-myGram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetAllSosmed(ctx *gin.Context) {
	db := database.GetDB()

	var sosmeds []models.Sosmed

	err := db.Model(&models.Sosmed{}).Find(&sosmeds).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"sosmeds": sosmeds,
		"code":    200,
	})
}

func GetOneSosmed(ctx *gin.Context) {
	db := database.GetDB()

	var (
		sosmedId string
		sosmed   models.Sosmed
	)

	sosmedId = ctx.Param("sosmedId")

	if err := db.First(&models.Sosmed{}, "id = ?", sosmedId).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("gagal mengambil data sosmed dengan id %v karna data tidak di temukan", sosmedId),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"sosmed": sosmed,
		"code":   200,
	})
}

func CreateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)

	Sosmed := models.Sosmed{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Sosmed)
	} else {
		ctx.ShouldBind(&Sosmed)
	}

	Sosmed.UserID = userID

	err := db.Debug().Create(&Sosmed).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"sosmed": Sosmed,
		"code":   200,
	})
}

func UpdateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	var sosmed models.Sosmed
	id, _ := strconv.Atoi(ctx.Param("socmed_id"))

	if err := db.First(&models.Sosmed{}, "id = ?", id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Gagal Update Socmed Id %v tidak di temukan", id),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&sosmed); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	err := db.Model(&models.Sosmed{}).Where("id = ?", id).Updates(sosmed).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"sosmed": sosmed,
		"code":   200,
	})
}

func DeleteSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	sosmed := models.Sosmed{}

	id, _ := strconv.Atoi(ctx.Param("socmed_id"))

	if err := db.First(&sosmed, "id = ?", id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Gagal menghapus socmed id %v tidak di temukan", id),
		})
		return
	}

	if err := db.Delete(&sosmed).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("socmed dengan id %v berhasil dihapus", id),
		"code":    200,
	})
}
