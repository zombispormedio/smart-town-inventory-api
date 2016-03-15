package controllers
import(
    "github.com/gin-gonic/gin"
    "gopkg.in/mgo.v2"
    "github.com/Zombispormedio/smartdb/response"
    "github.com/Zombispormedio/smartdb/models"
    "github.com/Zombispormedio/smartdb/utils"
    // "fmt"
)



func Register(c *gin.Context, session *mgo.Session ){

    defer session.Close()

    body_interface,_:=c.Get("body")
    body:=utils.InterfaceToMapString(body_interface);

    oauth:=models.OAuth{}    

    NewOauthError:=oauth.Register(body, session)

    if  NewOauthError == nil{
        response.SuccessMessage(c, "User Registered")
    }else{
        response.Error(c, NewOauthError);
    }

}


func Login(c *gin.Context, session *mgo.Session ){
    defer session.Close()

    body_interface,_:=c.Get("body")
    body:=utils.InterfaceToMapString(body_interface);
    
    token, LoginError:=models.Login(body, session)
    
    if LoginError == nil{
        response.Success(c, token)
    }else{
           response.Error(c, LoginError);
    }
    

}

func Logout(c *gin.Context, session *mgo.Session){
     defer session.Close()
    token:=c.Request.Header.Get("Authorization")
    pre_user, _:=c.Get("user")
    user:=pre_user.(string)
     
    LogoutError:=models.Logout(token, user, session)
    
    if LogoutError == nil{
        response.SuccessMessage(c, "Congratulations, You have logged out")
    }else{
        response.Error(c, LogoutError)
    }
}