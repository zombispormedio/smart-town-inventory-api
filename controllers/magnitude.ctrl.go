package controllers

import (
	"github.com/Zombispormedio/smartdb/models"
	"github.com/Zombispormedio/smartdb/response"
	"github.com/Zombispormedio/smartdb/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func CreateMagnitude(c *gin.Context, session *mgo.Session) {
	defer session.Close()

	preUser, _ := c.Get("user")
	user := preUser.(string)

	bodyInterface, _ := c.Get("body")
	body := utils.InterfaceToMap(bodyInterface)

	magnitude := models.Magnitude{}

	NewMagnitudeError := magnitude.New(body, user, session)

	if NewMagnitudeError == nil {
		response.SuccessMessage(c, "Magnitude Created")
	} else {
		response.Error(c, NewMagnitudeError)
	}

}

func GetMagnitudes(c *gin.Context, session *mgo.Session) {
	defer session.Close()

	var result []models.ListMagnitudeItem

	GetAllError := models.GetMagnitudes(&result, session)
	if GetAllError == nil {
		response.Success(c, result)
	} else {
		response.Error(c, GetAllError)
	}

}

func DeleteMagnitude(c *gin.Context, session *mgo.Session) {
	id := c.Param("id")

	RemoveError := models.DeleteMagnitude(id, session)

	if RemoveError == nil {
		GetMagnitudes(c, session)
	} else {
		response.Error(c, RemoveError)
		session.Close()
	}

}

func GetMagnitudeByID(c *gin.Context, session *mgo.Session) {
	defer session.Close()
	id := c.Param("id")

	magnitude := models.Magnitude{}

	ByIdError := magnitude.ByID(id, session)

	if ByIdError == nil {
		response.Success(c, magnitude)
	} else {
		response.Error(c, ByIdError)

	}

}

func SetMagnitudeDisplayName(c *gin.Context, session *mgo.Session) {
	defer session.Close()
	id := c.Param("id")

	bodyInterface, _ := c.Get("body")
	body := utils.InterfaceToMap(bodyInterface)

	magnitude := models.Magnitude{}

	SettingError := magnitude.SetDisplayName(id, body["display_name"].(string), session)

	if SettingError == nil {
		response.Success(c, magnitude)
	} else {
		response.Error(c, SettingError)

	}
}

func SetMagnitudeType(c *gin.Context, session *mgo.Session) {
	defer session.Close()
	id := c.Param("id")

	bodyInterface, _ := c.Get("body")

	body := utils.InterfaceToMap(bodyInterface)

	magnitude := models.Magnitude{}

	SettingError := magnitude.SetType(id, body["type"].(string), session)

	if SettingError == nil {
		response.Success(c, magnitude)
	} else {
		response.Error(c, SettingError)

	}
}

func SetMagnitudeDigitalUnits(c *gin.Context, session *mgo.Session) {
	defer session.Close()
	id := c.Param("id")

	bodyInterface, _ := c.Get("body")
	body := utils.InterfaceToMap(bodyInterface)

	magnitude := models.Magnitude{}

	SettingError := magnitude.SetDigitalUnits(id, body["digital_units"], session)

	if SettingError == nil {
		response.Success(c, magnitude)
	} else {
		response.Error(c, SettingError)

	}
}

func AddMagnitudeAnalogUnit(c *gin.Context, session *mgo.Session) {
	defer session.Close()
	id := c.Param("id")

	bodyInterface, _ := c.Get("body")
	body := utils.InterfaceToMap(bodyInterface)

	magnitude := models.Magnitude{}

	SettingError := magnitude.AddAnalogUnit(id, body["analog_unit"], session)

	if SettingError == nil {
		response.Success(c, magnitude)
	} else {
		response.Error(c, SettingError)

	}
}

func UpdateMagnitudeAnalogUnit(c *gin.Context, session *mgo.Session) {
	defer session.Close()
	id := c.Param("id")
	bodyInterface, _ := c.Get("body")
	body := utils.InterfaceToMap(bodyInterface)
	magnitude := models.Magnitude{}

	SettingError := magnitude.UpdateAnalogUnit(id, body["analog_unit"], session)

	if SettingError == nil {
		response.Success(c, magnitude)
	} else {
		response.Error(c, SettingError)

	}
}
