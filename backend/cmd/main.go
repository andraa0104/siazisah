package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/siazisah/config"
	"github.com/yourusername/siazisah/internal/handlers"
	"github.com/yourusername/siazisah/internal/middleware"
	"github.com/yourusername/siazisah/internal/repository"
)

func main() {
	config.LoadConfig()

	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	masjidRepo := repository.NewMasjidRepository(db)
	pengurusRepo := repository.NewPengurusRepository(db)
	muzakkiRepo := repository.NewMuzakkiRepository(db)
	mustahiqRepo := repository.NewMustahiqRepository(db)
	transaksiRepo := repository.NewTransaksiZakatRepository(db)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(userRepo)
	userHandler := handlers.NewUserHandler(userRepo)
	masjidHandler := handlers.NewMasjidHandler(masjidRepo)
	pengurusHandler := handlers.NewPengurusHandler(pengurusRepo)
	muzakkiHandler := handlers.NewMuzakkiHandler(muzakkiRepo)
	mustahiqHandler := handlers.NewMustahiqHandler(mustahiqRepo)
	transaksiHandler := handlers.NewTransaksiHandler(transaksiRepo, muzakkiRepo)
	publicHandler := handlers.NewPublicHandler(db)

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	// Public routes (no authentication)
	public := router.Group("/api/public")
	{
		public.GET("/dashboard", publicHandler.GetPublicDashboard)
		public.GET("/masjid", publicHandler.GetAllMasjid)
		public.GET("/masjid/:id/stats", publicHandler.GetMasjidStats)
		public.GET("/masjid/:id/pengurus-masjid", pengurusHandler.GetPengurusMasjidPublic)
		public.GET("/masjid/:id/pengurus-zakat", pengurusHandler.GetPengurusZakatPublic)
		public.GET("/stats/fitrah", publicHandler.GetZakatFitrahStats)
		public.GET("/stats/mal", publicHandler.GetZakatMalStats)
		public.GET("/stats/fidyah", publicHandler.GetFidyahStats)
	}

	// Auth routes
	auth := router.Group("/api/auth")
	{
		auth.POST("/login", authHandler.Login)
	}

	// Protected routes
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// Superadmin only routes
		superadmin := api.Group("/superadmin")
		superadmin.Use(middleware.SuperAdminOnly())
		{
			superadmin.GET("/masjid", masjidHandler.GetAll)
			superadmin.POST("/masjid", masjidHandler.Create)
			superadmin.GET("/masjid/:id", masjidHandler.GetByID)
			superadmin.PUT("/masjid/:id", masjidHandler.Update)
			superadmin.DELETE("/masjid/:id", masjidHandler.Delete)

			superadmin.GET("/users", userHandler.GetAll)
			superadmin.POST("/users", userHandler.Create)
			superadmin.GET("/users/:id", userHandler.GetByID)
			superadmin.PUT("/users/:id", userHandler.Update)
			superadmin.DELETE("/users/:id", userHandler.Delete)
		}

		// Petugas routes
		petugas := api.Group("/petugas")
		petugas.Use(middleware.PetugasOnly())
		{
			petugas.GET("/masjid", masjidHandler.GetMyMasjid)
			petugas.PUT("/masjid", masjidHandler.UpdateMyMasjid)

			petugas.GET("/pengurus-masjid", pengurusHandler.GetPengurusMasjid)
			petugas.POST("/pengurus-masjid", pengurusHandler.SavePengurusMasjid)
			
			petugas.GET("/pengurus-zakat", pengurusHandler.GetPengurusZakat)
			petugas.POST("/pengurus-zakat", pengurusHandler.SavePengurusZakat)
			petugas.DELETE("/pengurus-zakat/:id", pengurusHandler.DeletePengurusZakat)

			petugas.GET("/muzakki", muzakkiHandler.GetAll)
			petugas.POST("/muzakki", muzakkiHandler.Create)
			petugas.GET("/muzakki/:id", muzakkiHandler.GetByID)
			petugas.PUT("/muzakki/:id", muzakkiHandler.Update)
			petugas.DELETE("/muzakki/:id", muzakkiHandler.Delete)

			petugas.GET("/mustahiq", mustahiqHandler.GetAll)
			petugas.POST("/mustahiq/check", mustahiqHandler.CheckDuplicate)
			petugas.POST("/mustahiq", mustahiqHandler.Create)
			petugas.GET("/mustahiq/:id", mustahiqHandler.GetByID)
			petugas.PUT("/mustahiq/:id", mustahiqHandler.Update)
			petugas.DELETE("/mustahiq/:id", mustahiqHandler.Delete)

			petugas.GET("/transaksi", transaksiHandler.GetAll)
			petugas.POST("/transaksi", transaksiHandler.Create)
			petugas.GET("/transaksi/:id", transaksiHandler.GetByID)
			petugas.GET("/transaksi/:id/print", transaksiHandler.PrintReceipt)
			petugas.DELETE("/transaksi/:id", transaksiHandler.Delete)

			petugas.GET("/dashboard", transaksiHandler.GetDashboard)
		}
	}

	log.Printf("Server running on port %s", config.AppConfig.ServerPort)
	router.Run(":" + config.AppConfig.ServerPort)
}
