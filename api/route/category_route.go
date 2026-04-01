package route

import (
	"database/sql"
	"time"

	"github.com/fauzimhub/stokku/api/controller"
	"github.com/fauzimhub/stokku/repository"
	"github.com/fauzimhub/stokku/usecase"
	"github.com/gin-gonic/gin"
)

func NewCategoryRouter(timeout time.Duration, db *sql.DB, group *gin.RouterGroup) {
	cr := repository.NewCategoryRepository(db)
	cc := &controller.CategoryController{
		CategoryUsecase: usecase.NewCategoryUsecase(cr, timeout),
	}
	group.POST("/category", cc.Create)
	group.GET("/category", cc.Fetch)
	group.GET("/category/id/:id", cc.GetByID)
	group.GET("/category/name/:name", cc.GetByName)

}
