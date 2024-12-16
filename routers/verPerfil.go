package routers

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/twitterGo/bd"
	"github.com/twitterGo/models"
)

func VerPerfil(request events.APIGatewayProxyRequest) models.RespApi {
	var r models.RespApi
	r.Status = 400

	fmt.Println("Entre a ver perfil")
	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "EL parametro de ID es obligatorio"
		return r
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		r.Message = "Ocurrio un error al intentar buscar el registro " + err.Error()
		return r
	}

	respJson, err := json.Marshal(perfil)
	if err != nil {
		r.Status = 500
		r.Message = "Ocurrio un error al formatear los datos de los usuarios como json " + err.Error()
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}
