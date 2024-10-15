package routes

import (
	"github.com/evanpalsson/basesite-main/controllers"
	"github.com/gin-gonic/gin"
)

type PostRouteController struct {
	postController controllers.PostController
}

func NewPostControllerRoute(postController controllers.PostController) PostRouteController {
	return PostRouteController{postController}
}

func (r *PostRouteController) PostRoute(rg *gin.RouterGroup) {
	router := rg.Group("/posts")

	router.GET("/", r.postController.FindPosts)
	router.GET("/:postId", r.postController.FindPostById)
	router.POST("/", r.postController.CreatePost)
	router.PATCH("/:postId", r.postController.UpdatePost)
	router.DELETE("/:postId", r.postController.DeletePost)
}
