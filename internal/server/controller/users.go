package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/oresdev/tbcc-wallet-api-v3/internal/server/model"
	"github.com/oresdev/tbcc-wallet-api-v3/internal/server/service"
	"github.com/oresdev/tbcc-wallet-api-v3/internal/server/util"
	"github.com/sirupsen/logrus"
)

// GetUsersHandler ...
func GetUsersHandler(conn *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := service.DbGetUsers(conn, r.Context())
		if err != nil {
			logrus.Errorf("GetUsersHandler db: %v", err)
			http.Error(w, "GetUsersHandler err", http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(users); err != nil {
			logrus.Errorf("GetUsersHandler write: %v", err)
			http.Error(w, "GetUsersHandler write err", http.StatusInternalServerError)
			return
		}
	}
}

// GetUserHandler ...
func GetUserHandler(conn *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "uuid")
		uuid, _ := util.FromString(id)

		user, err := service.DbGetUserByID(uuid, conn, r.Context())

		if err != nil {
			logrus.Errorf("getUserHandler db: %v", err)
			http.Error(w, "getUserHandler err", http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(user); err != nil {
			logrus.Errorf("GetUserHandler write: %v", err)
			http.Error(w, "GetUserHandler write err", http.StatusInternalServerError)
			return
		}
	}
}

// GetExtendedUserHandler ...
func GetExtendedUserHandler(conn *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "uuid")

		user, err := service.DbGetUserExt(id, conn, r.Context())

		if err != nil {
			logrus.Errorf("GetExtendedUserHandler db: %v", err)
			http.Error(w, "GetExtendedUserHandler err", http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(user); err != nil {
			logrus.Errorf("GetExtendedUserHandler write: %v", err)
			http.Error(w, "GetExtendedUserHandler write err", http.StatusInternalServerError)
			return
		}
	}
}

// UpdateUserHandler ...
func UpdateUserHandler(conn *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := chi.URLParam(r, "uuid")

		u := model.UserMigrate{}

		user, err := service.DbUpdateUser(uuid, u.Address, conn, r.Context())
		if err != nil {
			logrus.Errorf("DbUpdateUser: %v", err)
			http.Error(w, "DbUpdateUser err", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(user); err != nil {
			logrus.Errorf("UpdateUserHandler write uuid: %v", err)
			http.Error(w, "UpdateUserHandler write uuid", http.StatusInternalServerError)
			return
		}
	}
}

// CreateUserHandler ...
func CreateUserHandler(conn *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := model.User{}

		user, err := service.DbCreateUser(u.Useraddress, u.Accounttype, u.Smartcard, conn, r.Context())
		if err != nil {
			logrus.Errorf("createUserHandler db: %v", err)
			http.Error(w, "createUserHandler err", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(user); err != nil {
			logrus.Errorf("CreateUserHandler write id: %v", err)
			http.Error(w, "CreateUserHandler write id", http.StatusInternalServerError)
			return
		}
	}
}

// MigrateUserHandler ...
func MigrateUserHandler(conn *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := model.User{}

		user, err := service.DbMigrateUser(u.Useraddress, conn, r.Context())

		if err != nil {
			logrus.Errorf("GetExtendedUserHandler db: %v", err)
			http.Error(w, "GetExtendedUserHandler err", http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(user); err != nil {
			logrus.Errorf("GetExtendedUserHandler write: %v", err)
			http.Error(w, "GetExtendedUserHandler write err", http.StatusInternalServerError)
			return
		}
	}
}

// PurchaseVpnKeyHandler ...
func PurchaseVpnKeyHandler(conn *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := chi.URLParam(r, "uuid")

		v := model.VpnKeyBuyBody{}

		if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
			http.Error(w, "PurchaseVpnKeyHandler read invalid params", http.StatusBadRequest)
			return
		}

		key, err := service.DbUpdateVpnKey(uuid, v.TxHash, conn, r.Context())
		if err != nil {
			logrus.Errorf("DbUpdateVpnKey db: %v", err)
			http.Error(w, "DbUpdateVpnKey err", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(&key); err != nil {
			logrus.Errorf("PurchaseVpnKeyHandler write id: %v", err)
			http.Error(w, "PurchaseVpnKeyHandler write id", http.StatusInternalServerError)
			return
		}
	}
}
