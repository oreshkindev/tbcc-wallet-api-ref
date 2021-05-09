package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v4/pgxpool"
	h "github.com/oresdev/tbcc-wallet-api-v3/internal/server/controller"
	"github.com/oresdev/tbcc-wallet-api-v3/internal/server/middleware/rsa"
)

// CreateHTTPHandler ...
func CreateHTTPHandler(conn *pgxpool.Pool) (http.Handler, error) {
	mux := chi.NewMux()

	mux.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		if err := conn.Ping(r.Context()); err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(http.StatusText(200)))
	})

	mux.Route("/", func(r chi.Router) {
		r.Use(middleware.SetHeader("Content-Type", "application/json"))
		r.Use(cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           30,
		}).Handler)

		r.Use(rsa.CheckRSASignature) // CheckRSASignature

		r.Mount("/users", UserHandler(conn))
		r.Mount("/app", AppHandler(conn))

	})

	return mux, nil
}

// UserHandler ...
func UserHandler(conn *pgxpool.Pool) http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		//r.Get("/", h.GetUsersHandler(db))
		r.Get("/{uuid}", h.GetUserHandler(conn))
		r.Get("/ext/{uuid}", h.GetExtendedUserHandler(conn))
		r.Post("/{uuid}/update", h.UpdateUserHandler(conn))
		//r.Post("/", h.CreateUserHandler(db)) // TODO remove development routes

		// Migrate user data from depricated database (public scheme)
		// Returns extended user data
		r.Post("/migrate", h.MigrateUserHandler(conn))
		r.Post("/{uuid}/buy-vpn", h.PurchaseVpnKeyHandler(conn))
	})

	return r
}

// AppHandler ...
func AppHandler(conn *pgxpool.Pool) http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Post("/config", h.CreateConfigHandler(conn)) // TODO remove development routes
		r.Get("/config", h.GetConfigHandler(conn))
		//r.Post("/update", h.CreateUpdateHandler(db)) // TODO remove development routes
		r.Get("/update", h.GetUpdateHandler(conn))
		r.Post("/counter", h.CountVersionHandler(conn))
	})

	return r
}
