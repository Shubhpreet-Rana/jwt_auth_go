package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Shubhpreet-Rana/jwt_auth_go/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandler(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	t.Run("should fail if the user payloas is invalid", func(t *testing.T) {

		payload := types.RegisterUserPayload{
			FirstName: "Shubhpreet",
			LastName:  "Rana",
			Email:     "invalid",
			Password:  "1234",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, git %d", http.StatusBadRequest, rr.Code)
		}

	})

	t.Run("should correctrly register user", func(t *testing.T) {

		payload := types.RegisterUserPayload{
			FirstName: "Shubhpreet",
			LastName:  "Rana",
			Email:     "shubhpreet@yopmail.com",
			Password:  "1234",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, git %d", http.StatusCreated, rr.Code)
		}

	})

}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("User not found")
}
func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
func (m *mockUserStore) Createuser(types.User) error {
	return nil
}
