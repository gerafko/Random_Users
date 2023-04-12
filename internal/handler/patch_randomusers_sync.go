package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"test/internal/model"
	"test/pkg/randomuser"
)

type PatchRandomUsersSync struct {
	ua userAccessor
	us userStorage
}

func NewPatchRandomUsersSync(ua userAccessor, us userStorage) *PatchRandomUsersSync {
	return &PatchRandomUsersSync{
		ua: ua,
		us: us,
	}
}

func (p *PatchRandomUsersSync) Handle(c *gin.Context) {
	res, err := p.ua.RandomUserGET(c.Request.Context())
	if err != nil {
		if err == randomuser.ErrNoUsers {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		log.Errorf("random users sync failed with such error: %v", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	users := model.MapGetRandomUserResponseToUsers(res)

	//
	// Здесь могла быть ваша бизнес логика
	//

	err = p.us.InsertList(users)
	if err != nil {
		log.Errorf("random users sync failed with storage error: %v", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "sync successful",
	})
}
