package core

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TopSecret function
// topsecret godoc
// @Summary      Top Secret
// @Description  Decrypts the messages intercepted by the satellites
// @Tags         Resistance
// @Accept       json
// @Produce      json
// @Param   	 topSecret    body  TopSecret   true  "top secret"
// @Success      200  {map}     map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /topsecret [post]
func TopSecret(c *gin.Context) {
	var secret Secret
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&secret)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	distances, messages := GetData(secret.Satellites)

	x, y := GetLocation(distances...)
	message := GetMessage(messages...)

	c.JSON(http.StatusOK, gin.H{
		"position": gin.H{
			"x": x,
			"y": y,
		},
		"message": message,
	})
}

// TopSecretSplitPOST function
// topsecret_split godoc
// @Summary      Top Secret Split
// @Description  Decrypts the message intercepted by the satelite
// @Tags         Resistance
// @Accept       json
// @Produce      json
// @Param   	 satelite_name    path  string         true  "satelite name"
// @Param   	 transmitter   	  body  Transmitter    true  "transmitter"
// @Success      200  {map}     map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /topsecret_split [post]
func TopSecretSplitPOST(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"error": "No hay suficiente información",
	})
}

// TopSecretSplitGET function
// topsecret_split godoc
// @Summary      Top Secret Split
// @Description  Decrypts the message intercepted by the satelite
// @Tags         Resistance
// @Accept       json
// @Produce      json
// @Param   	 satelite_name    path  string         true  "satelite name"
// @Success      200  {map}     map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /topsecret_split [get]
func TopSecretSplitGET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "top secret split get",
	})
}
