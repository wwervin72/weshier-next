package article

import (
	"weshierNext/handler"
	"weshierNext/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Create create article
func Create(c *gin.Context) {
	body := &CreateArticleReqStruct{}
	if err := c.ShouldBindJSON(&body); err != nil {
		handler.SendResponse(c, errno.ErrBadRequest, nil)
		return
	}

}
