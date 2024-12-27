package routers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/twitterGo/bd"
	"github.com/twitterGo/models"
)

func EliminarTweet(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {

	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El paramatro ID es obligatorio"
		return r
	}

	err := bd.BorroTweet(ID, claim.ID.Hex())
	if err != nil {
		r.Message = "Ocurrio un error al eliminar el tweet"
		return r
	}

	r.Message = "Tweet eliminado"
	r.Status = 200
	return r

}
