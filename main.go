package main

import (
	//"bytes"
	//"context"
	"fmt"

	// "time"
	// "log"

	"net/http"

	"github.com/RanggaPrathama/message_broker_fcm/domain/models"
	"github.com/RanggaPrathama/message_broker_fcm/domain/repository"
	"github.com/RanggaPrathama/message_broker_fcm/handler"
	"github.com/RanggaPrathama/message_broker_fcm/lib"
	"github.com/RanggaPrathama/message_broker_fcm/routes"
	"github.com/RanggaPrathama/message_broker_fcm/service"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/sessions"

	//"net/http/httptest"

	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func main() {
	app := fiber.New()
	lib.ConnectionPostgree()

	
	var store = sessions.NewCookieStore([]byte(lib.LoadEnv("SESSION_SECRET")))

	gothic.Store = store


	//lib.Database.Migrator().DropTable(&models.User{}, &models.DeviceUser{}, &models.RoomChat{}, &models.RoomMember{}, &models.Message{})
	lib.Database.AutoMigrate(&models.User{}, &models.DeviceUser{}, &models.RoomChat{}, &models.RoomMember{}, &models.Message{})

	goth.UseProviders(
		google.New(
			lib.LoadEnv("GOOGLE_CLIENT_ID"),
			lib.LoadEnv("GOOGLE_CLIENT_SECRET"),
			lib.LoadEnv("GOOGLE_CLIENT_CALLBACK"),
		),
	)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	hub := lib.NewHub()
	go hub.Run()

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		handler.WebSocketHandler(hub)

	}))

	userRepo := repository.NewUserRepository(lib.Database)
	deviceRepo := repository.NewDeviceRepository(lib.Database)

	userService := service.NewUserService(userRepo)
	userhandler := handler.NewUserHandler(userService)

	authService := service.NewAuthService(userRepo, deviceRepo)
	authHandler := handler.NewAuthHandler(authService)

	deviceService := service.NewDeviceService(deviceRepo, userRepo)
	deviceHandler := handler.NewDeviceHandler(deviceService)

	routes.UserRoute(app, userhandler)
	routes.AuthRoute(app, authHandler)
	routes.DeviceRoute(app, deviceHandler)

	// Tes Message
	routes.MessageRoute(app)

	type contextKey string

	const providerKey contextKey = "provider"


	app.Get("/auth/login/google", adaptor.HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//  HandlerLogin(w, r, "google")
		handler.HandlerLogin(w, r, "google")
	 }))


	// app.Get("/auth/login/google", adaptor.HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	// provider := chi.URLParam(r, "provider")
	// 	// set context provider
	// 	r = r.WithContext(context.WithValue(context.Background(), "provider", "google"))
	// 	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
	// 	 fmt.Println("error sini: ", gothUser)
	// 	} else {
	// 	 gothic.BeginAuthHandler(w, r)
	// 	}
	// }))


	app.Get("/auth/google/callback", authHandler.HandlerLoginCallback)


	// app.Get("/auth/google/callback", func(c *fiber.Ctx) error {
	// 	//provider := c.Params("provider")

	// 	ctx := context.WithValue(c.Context(), providerKey, "google")

	// 	// Membuat http.Request dari Fiber request
	// 	r, err := http.NewRequest(c.Method(), c.OriginalURL(), bytes.NewReader(c.Body()))
	// 	if err != nil {
	// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 			"error":   true,
	// 			"message": "Gagal membuat request",
	// 		})
	// 	}

	// 	// Menyalin header dari Fiber request ke http.Request
	// 	for key, values := range c.GetReqHeaders() {
	// 		for _, value := range values {
	// 			r.Header.Add(key, value)
	// 		}
	// 	}

	// 	r = r.WithContext(ctx)      // Menambahkan context provider
	// 	w := httptest.NewRecorder() // Response writer untuk menangkap respons

	// 	user, err := gothic.CompleteUserAuth(w, r)
	// 	if err != nil {
	// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 			"error":   true,
	// 			"message": err.Error(),
	// 		})
	// 	}

	// 	// Mengembalikan data pengguna yang berhasil login
	// 	return c.JSON(fiber.Map{
	// 		"success":  true,
	// 		"provider": user.Provider,
	// 		"user": fiber.Map{
	// 			"user_id":    user.UserID,
	// 			"name":       user.Name,
	// 			"email":      user.Email,
	// 			"avatar_url": user.AvatarURL,
	// 		},
	// 	})
	// })

	// Login dengan Google
	// http.HandleFunc("/auth/google", func(w http.ResponseWriter, r *http.Request) {
	// 	url, err := gothic.GetAuthURL(w, r)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	// })

	// http.HandleFunc("/auth/google/callback", func(w http.ResponseWriter, r *http.Request) {
	// 	user, err := gothic.CompleteUserAuth(w, r)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	fmt.Fprintf(w, "Hello, %s!", user.Name)
	// })

	// app.Listen(fmt.Sprintf(":%s", lib.LoadEnv("APP_PORT")))
	go func() {
		if err := app.Listen(fmt.Sprintf(":%s", lib.LoadEnv("APP_PORT"))); err != nil {
			fmt.Println("Failed to start Fiber:", err)
		}
	}()

	// conn , _ := lib.ConnectionRabbitMQ()
	// channel, q := lib.ChannelRabbitMQ(conn)

	// exchangeName := "notif_exchange"
	// lib.DeclareExchange(channel, exchangeName)

	//  lib.ConsumeRabbitMQ(channel, q, exchangeName)

	// go func() {
	// 	lib.ConsumeRabbitMQ(channel, q, exchangeName)
	// }()

	// Tunggu agar main tidak langsung selesai
	select {}

}



