package types

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database = *mongo.Database
type Table = *mongo.Collection

type APIServer struct {
	APIRouter *gin.RouterGroup
	DB        Database
}
