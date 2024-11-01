package extensions

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserID(ctx *gin.Context) (uuid.UUID, error) {
	userId, exists := ctx.Get("userId")
	if !exists {
		return uuid.Nil, fmt.Errorf("Unauthorized: user ID not found in context")
	}

	userUUID, err := uuid.Parse(userId.(string))
	if err != nil {
		return uuid.Nil, err
	}

	return userUUID, nil
}
