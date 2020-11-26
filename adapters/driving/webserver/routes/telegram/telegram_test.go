package telegram

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
	"sentaly.com/telegram-bot/ports"
	"sentaly.com/telegram-bot/test/mocks"
)

func setupTestRouter(s ports.BotService) *gin.Engine {
	r := gin.Default()
	api := r.Group("/")

	SetRoutes(api, s)

	return r
}

func Test_TelegramShouldReturnErrorIfParseRequestHaveError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := mocks.NewMockBotService(ctrl)

	s.EXPECT().ParseRequest(nil).Return(nil, errors.New("Parse error"))

	router := setupTestRouter(s)
	w := httptest.NewRecorder()
	
	req, _ := http.NewRequest("POST", "/telegram", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_TelegramShouldReturnErrorIfProcessRequestHaveError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := mocks.NewMockBotService(ctrl)

	update := &ports.Update{}

	s.EXPECT().ParseRequest(nil).Return(update, nil)
	s.EXPECT().ProcessRequest(update).Return(errors.New("Request error"))

	router := setupTestRouter(s)
	w := httptest.NewRecorder()
	
	req, _ := http.NewRequest("POST", "/telegram", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func Test_TelegramShouldReturnOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := mocks.NewMockBotService(ctrl)

	update := &ports.Update{}

	s.EXPECT().ParseRequest(nil).Return(update, nil)
	s.EXPECT().ProcessRequest(update).Return(nil)

	router := setupTestRouter(s)
	w := httptest.NewRecorder()
	
	req, _ := http.NewRequest("POST", "/telegram", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
