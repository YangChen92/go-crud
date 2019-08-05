package api

import(
      _ "giligili/serializer"
		"github.com/gin-gonic/gin"
		"giligili/service"
)

//CreateVideo 成功
func CreateVideo(c *gin.Context){
	service := service.CreateVideoService{}
	if err := c.ShouldBind(&service); err == nil{
		res := service.Create()
		c.JSON(200,res)
	}else{
		c.JSON(200,ErrorResponse(err))
	}
	
}
