package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DestroyerAlpha/COMSW4156-AuthDemo/auth"
	"github.com/DestroyerAlpha/COMSW4156-AuthDemo/rebac"
	"github.com/DestroyerAlpha/COMSW4156-AuthDemo/router"
	"github.com/DestroyerAlpha/COMSW4156-AuthDemo/user"
	"github.com/DestroyerAlpha/COMSW4156-AuthDemo/user/model"
	"github.com/gorilla/mux"
)

var url = "http://localhost:8080"

func startServer(middlewareFunc ...mux.MiddlewareFunc) *httptest.Server {
	ts := httptest.NewServer(router.SetupRoutes(middlewareFunc...))
	return ts
}

func initialSetup(t *testing.T, middlewareFunc ...mux.MiddlewareFunc) *httptest.Server {
	ts := startServer(middlewareFunc...)
	url = ts.URL
	//go http.ListenAndServe(url, router.SetupRoutes())
	req1, _ := json.Marshal(user.CreateUserRequest{
		Username: "user-1",
		Password: "password",
		Name: &model.Name{
			FirstName: "user",
			LastName:  "1",
		},
	})
	resp, err := http.Post(url+"/api/user/create", "application/json", bytes.NewBuffer(req1))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("got %d, want %d", resp.StatusCode, http.StatusCreated)
	}
	req2, _ := json.Marshal(user.CreateUserRequest{
		Username: "user-2",
		Password: "password",
		Name: &model.Name{
			FirstName: "user",
			LastName:  "2",
		},
	})
	resp, err = http.Post(url+"/api/user/create", "application/json", bytes.NewBuffer(req2))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("got %d, want %d", resp.StatusCode, http.StatusCreated)
	}
	req3, _ := json.Marshal(user.CreateUserRequest{
		Username: "user-3",
		Password: "password",
		Name: &model.Name{
			FirstName: "user",
			LastName:  "3",
		},
	})
	resp, err = http.Post(url+"/api/user/create", "application/json", bytes.NewBuffer(req3))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("got %d, want %d", resp.StatusCode, http.StatusCreated)
	}
	return ts
}

func TestSample(t *testing.T) {
	t.Run("Non Friend Cannot Access User Data", func(t *testing.T) {
		ts := initialSetup(t, rebac.RebacMiddleware, auth.AuthMiddleware)
		defer ts.Close()
		req, _ := json.Marshal(auth.AuthenticateRequest{
			Username: "user-1",
			Password: "password",
		})
		resp, err := http.Post(url+"/api/auth/authenticate", "", bytes.NewBuffer(req))
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("got %d, want %d", resp.StatusCode, http.StatusOK)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		var authResp auth.AuthenticateResponse
		if err = json.Unmarshal(body, &authResp); err != nil {
			t.Fatal(err)
		}
	})
}
