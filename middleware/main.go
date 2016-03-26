package middleware

import(

    "os"
    "strings"
    "github.com/gin-gonic/gin"
    "gopkg.in/mgo.v2"

    "github.com/Zombispormedio/smartdb/response"
     "github.com/Zombispormedio/smartdb/models"
    //"fmt"

)

func Sensor() gin.HandlerFunc{
    return func(c *gin.Context){ 
        c.Next()
    }
}

func Secret() gin.HandlerFunc{
    return func(c *gin.Context){

        equals := strings.Compare(c.Request.Header.Get("Authorization"), os.Getenv("SMARTDBSECRET"))
        if equals !=0{
            response.ErrorByString(c, 403, "No Authorization: SecretError");
            return;
        }
        c.Next();

    }
}

func Body() gin.HandlerFunc{
    return func(c *gin.Context){

        var body map[string]string


        BindingJSONError:=c.BindJSON(&body)
        if BindingJSONError!=nil{
            response.ErrorByString(c, 400, "No body in HttpRequest");
            return;
        }
        
        c.Set("body", body);
        
        c.Next();
    }
}


func Admin(session *mgo.Session) gin.HandlerFunc{
    return func(c *gin.Context){
        
        token:=c.Request.Header.Get("Authorization")
        
        if token == ""{
            response.ErrorByString(c, 403, "No Authorization: Empty Token");
            return;
        }
        
        user, err:=models.SessionToken(token, session)
       
        if err != nil{
            response.Error(c, err);
            return;
        }
        
        c.Set("user", user)
        
        
        c.Next();
        
    }
}