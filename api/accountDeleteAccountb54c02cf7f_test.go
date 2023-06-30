package api

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/wil-ckaew/gofinance-backend/db"
)

type mockStore struct{
	db.SQLStore
}

func (m *mockStore) DeleteAccount(ctx *gin.Context, id int64) error {
	if id == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func TestDeleteAccountb54c02cf7f(t *testing.T) {
	server := NewServer(&mockStore{})

	t.Run("delete an existing account", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodDelete, "/accounts/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		response := httptest.NewRecorder()
		server.router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "true", response.Body.String())
	})

	t.Run("delete a non-existing account", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodDelete, "/accounts/0", nil)
		if err != nil {
			t.Fatal(err)
		}

		response := httptest.NewRecorder()
		server.router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Contains(t, response.Body.String(), sql.ErrNoRows.Error())
	})
}
