package controllers

import (
	"fmt"
	"go-myGram/database"
	"go-myGram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllComment(ctx *gin.Context) {
	db := database.GetDB()

	var comments []models.Comment

	err := db.Model(&models.Comment{}).Find(&comments).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"comments": comments,
		"code":     200,
	})
}

func GetOneComment(ctx *gin.Context) {
	db := database.GetDB()

	var (
		commentId string
		comment   models.Comment
	)

	commentId = ctx.Param("commentId")

	if err := db.First(&models.Comment{}, "id = ?", commentId).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("gagal mengambil data comment dengan id %v karna data tidak di temukan", commentId),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"comment": comment,
		"code":    200,
	})
}

func CreateComment(ctx *gin.Context) {
	db := database.GetDB()
	var comment models.Comment

	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	err := db.Debug().Create(&comment).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"comment": comment,
		"code":    200,
	})
}

func UpdateComment(ctx *gin.Context) {
	db := database.GetDB()
	var comment models.Comment
	id := ctx.Query("id")

	if err := db.First(&models.Comment{}, "id = ?", id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("gagal Update order Id %v tidak di temukan", id),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	err := db.Model(&models.Comment{}).Where("id = ?", id).Updates(comment).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"comment": comment,
		"code":    200,
	})
}

func DeleteComment(ctx *gin.Context) {
	db := database.GetDB()
	comment := models.Comment{}

	id := ctx.Query("id")

	if err := db.First(&comment, "order_id = ?", id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Gagal menghapus comment id %v tidak di temukan", id),
		})
		return
	}

	if err := db.Delete(&comment).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("comment dengan id %v berhasil dihapus", id),
		"code":    200,
	})
}
