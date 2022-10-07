package upload

import (
	v1 "betxin/api/v1"
	"betxin/service/upload"
	"betxin/utils/errmsg"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		v1.SendResponse(c, errmsg.ERROR, nil)
	}
	fileSize := fileHeader.Size
	url, code := upload.UpLoadFile(file, fileSize)
	if code != errmsg.SUCCSE {
		fmt.Println("上传出错")
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, url)
}
