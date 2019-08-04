package api

import(
        "giligili/serializer"
    	"github.com/gin-gonic/gin"
)

//CreateVideo 成功
func CreateVideo(c *gin.Context){
	c.JSON(200, serializer.Response{
		Status: 0,
		Msg:    "createVideo success",
	})
}
