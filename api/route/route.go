package route

import (
	"time"

	"database/sql"
	"github.com/gin-gonic/gin"
)

func Setup(timeout time.Duration, db *sql.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	NewCategoryRouter(timeout, db, publicRouter)
}
