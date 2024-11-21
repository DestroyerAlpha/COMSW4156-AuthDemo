package rebac

import (
	"encoding/json"
	"log"
	"net/http"

	database "github.com/DestroyerAlpha/COMSW4156-AuthDemo/db"
	pkgResource "github.com/DestroyerAlpha/COMSW4156-AuthDemo/pkg/resource"
	"github.com/DestroyerAlpha/COMSW4156-AuthDemo/user"
)

func RebacMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("rebac middleware", r.URL)
		resource := r.Header.Get("resource")
		if isSkipMiddleware(resource) {
			log.Println("skip rebac middleware", resource)
			h.ServeHTTP(w, r)
		}
		userId := r.Header.Get("user_id")
		if !isPermitted(r, userId, resource) {
			http.Error(w, "user not allowed access to this endpoint", http.StatusForbidden)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func isSkipMiddleware(resource string) bool {
	if resource == pkgResource.CreateUser || resource == pkgResource.Authenticate {
		return true
	}
	return false
}

func isPermitted(r *http.Request, userId, resource string) bool {
	switch resource {
	case pkgResource.GetUser:
		var req user.GetUserRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return false
		}
		defer func() {
			_ = r.Body.Close()
		}()
		if userId == req.UserId || isFriend(userId, req.UserId) {
			return true
		}
	case pkgResource.AddFriend:
		//var req user.AddFriendRequest
		//err := json.NewDecoder(r.Body).Decode(&req)
		//if err != nil {
		//	return false
		//}
		//defer func() {
		//	_ = r.Body.Close()
		//}()
		//if userId == req.UserId {
		//	return true
		//}
	case pkgResource.GetFriends:
		var req user.GetFriendsRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return false
		}
		defer func() {
			_ = r.Body.Close()
		}()
		if userId == req.UserId || isFriend(userId, req.UserId) {
			return true
		}
	default:
		return false
	}
	return false
}

func isFriend(userId, reqUserId string) bool {
	db := database.GetDatabase()
	for _, usr := range db.Users {
		if usr.Id == reqUserId {
			for _, friend := range usr.Friends {
				if friend == userId {
					return true
				}
			}
		}
	}
	return false
}
