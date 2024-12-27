package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/twitterGo/bd"
	"github.com/twitterGo/models"
)

func LeoTweets(request events.APIGatewayProxyRequest) models.RespApi {
	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	pagina := request.QueryStringParameters["pagina"]

	if len(ID) < 1 {
		r.Message = "El paramatro ID es obligatorio"
		return r
	}
	if len(pagina) < 1 {
		pagina = "1"
	}
	pag, err := strconv.Atoi(pagina)
	if err != nil {
		r.Message = "Debe enviar el parametro pagina comoo un valor mayor a 0"
		return r
	}

	tweets, correcto := bd.LeoTweets(ID, int64(pag))
	if !correcto {
		r.Message = "Error al leer los tweets"
		return r
	}

	respJson, err := json.Marshal(tweets)
	if err != nil {
		r.Status = 500
		r.Message = "error al formatear los datos de usuarios como JSON"
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}
