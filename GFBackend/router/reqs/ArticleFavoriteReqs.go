package reqs

import "github.com/gin-gonic/gin"

func InitArticleFavoriteReqs(baseGroup *gin.RouterGroup) *gin.RouterGroup {

	// articleFavoriteController, _ := InitializeArticleFavoriteController()

	articleFavoriteReqsGroup := baseGroup.Group("/articlefavorite")
	{

	}

	return articleFavoriteReqsGroup
}
