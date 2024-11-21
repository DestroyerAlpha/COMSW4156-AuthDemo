package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DestroyerAlpha/COMSW4156-AuthDemo/auth"
	database "github.com/DestroyerAlpha/COMSW4156-AuthDemo/db"
	"github.com/DestroyerAlpha/COMSW4156-AuthDemo/user/dao"
	"github.com/DestroyerAlpha/COMSW4156-AuthDemo/user/model"
)

type CreateUserRequest struct {
	Username string      `json:"username"`
	Password string      `json:"password"`
	Name     *model.Name `json:"name"`
	Friends  []string    `json:"friends"`
}

type GetUserRequest struct {
	UserId string `json:"userId"`
}

type AddFriendRequest struct {
	FriendUserId string `json:"friendUserId"`
}

type GetFriendsRequest struct {
	UserId string `json:"userId"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}
	defer func() {
		_ = r.Body.Close()
	}()
	if err = auth.CreateAuthCredentials(req.Username, req.Password); err != nil {
		http.Error(w, fmt.Sprintf("Error in creating auth credentials: %v", err), http.StatusBadRequest)
		return
	}
	if err = dao.CreateUser(&model.User{
		Id:      req.Username,
		Name:    req.Name,
		Friends: req.Friends,
	}); err != nil {
		http.Error(w, "Error in creating user", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var req GetUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}
	db := database.GetDatabase()
	for _, usr := range db.Users {
		if usr.Id == req.UserId {
			if err = json.NewEncoder(w).Encode(usr); err != nil {
				http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)
}

func AddFriend(w http.ResponseWriter, r *http.Request) {
	var req AddFriendRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}
	db := database.GetDatabase()
	isFriendAUser := false
	for _, usr := range db.Users {
		if usr.Id == req.FriendUserId {
			isFriendAUser = true
		}
	}
	if !isFriendAUser {
		http.Error(w, "friend not found", http.StatusNotFound)
		return
	}
	for _, usr := range db.Users {
		if usr.Id == r.Header.Get("user_id") {
			usr.Friends = append(usr.Friends, req.FriendUserId)
		}
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, "user not found", http.StatusNotFound)
}

func GetFriends(w http.ResponseWriter, r *http.Request) {
	var req GetFriendsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}
	db := database.GetDatabase()
	for _, usr := range db.Users {
		if usr.Id == req.UserId {
			if err = json.NewEncoder(w).Encode(usr.Friends); err != nil {
				http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)
}
