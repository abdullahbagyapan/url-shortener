package app

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"url-shortener/auth-service/pkg/grpc/pb"
	"url-shortener/gateway-service/internal/ports"
	"url-shortener/shortener-service/pkg/grpc/urlpb"
	"url-shortener/token-service/pkg/grpc/tokenpb"
)

type Adapter struct {
	authService  pb.AuthServiceClient
	tokenService tokenpb.TokenServiceClient
	redis        ports.RedisPort
	urlService   urlpb.URLServiceClient
}

type registerRequest struct {
	Name     string
	Username string
	Password string
	Email    string
}

type urlshortRequest struct {
	Longurl string
}

type loginRequest struct {
	Username string
	Password string
}

func NewAdapter(authService pb.AuthServiceClient, tokenService tokenpb.TokenServiceClient, redis ports.RedisPort, urlService urlpb.URLServiceClient) *Adapter {
	return &Adapter{
		authService:  authService,
		tokenService: tokenService,
		redis:        redis,
		urlService:   urlService,
	}
}

func (A Adapter) Run() {
	app := fiber.New()

	// whitelist
	app.Get("urls/:shortURL", A.GetLongURL)
	app.Use("/login", A.Login)
	app.Use("/register", A.Register)
	app.Use(A.ValidateToken)

	app.Get("/home", func(ctx *fiber.Ctx) error {
		return ctx.JSON("hello world")
	})

	app.Post("urls", A.CreateShortURL)

	log.Fatal(app.Listen(":8080"))
}

func (A Adapter) Login(c *fiber.Ctx) error {

	var req = loginRequest{}

	err := c.BodyParser(&req)

	if err != nil {
		c.Status(500)

		return c.JSON("parse error")
	}

	grpcCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp, err := A.authService.Login(grpcCtx, &pb.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		c.Status(500)

		return c.JSON("Login error")
	}

	// cache token
	go func() {
		err := A.redis.SetToken(resp.Token)
		if err != nil {
			log.Printf("Error setting token to Redis %v", err)
		}

	}()

	c.Status(200)
	rtnMap := map[string]string{}

	rtnMap["token"] = resp.Token

	return c.JSON(rtnMap)

}

func (A Adapter) Register(c *fiber.Ctx) error {
	var req = registerRequest{}

	err := c.BodyParser(&req)

	if err != nil {
		c.Status(500)

		return c.JSON("parse error")
	}

	grpcCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp, err := A.authService.Register(grpcCtx, &pb.RegisterRequest{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		c.Status(500)

		return c.JSON("Register error")
	}
	// cache token
	go func() {
		err := A.redis.SetToken(resp.Token)
		if err != nil {
			log.Printf("Error setting token to Redis %v", err)
		}
	}()

	token := resp.Token

	c.Status(200)
	rtnMap := map[string]string{}

	rtnMap["token"] = token

	return c.JSON(rtnMap)
}

func (A Adapter) CreateShortURL(c *fiber.Ctx) error {

	var req = urlshortRequest{}

	err := c.BodyParser(&req)

	grpcCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp, err := A.urlService.GenerateShortURL(grpcCtx, &urlpb.NewShortURLRequest{
		LongURL: req.Longurl,
	})

	if err != nil {
		c.Status(500)

		return c.JSON("Url server error")
	}

	rtnMap := map[string]string{}

	rtnMap["shortURL"] = resp.ShortURL

	return c.JSON(rtnMap)

}

func (A Adapter) GetLongURL(c *fiber.Ctx) error {

	param := c.Params("shortURL")

	grpcCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp, err := A.urlService.GetLongURL(grpcCtx, &urlpb.GetLongURLRequest{
		ShortURL: param,
	})

	if err != nil {
		c.Status(500)

		return c.JSON("Url server error")
	}

	rtnMap := map[string]string{}

	rtnMap["longURL"] = resp.LongURL

	return c.JSON(rtnMap)
}

func (A Adapter) ValidateToken(ctx *fiber.Ctx) error {

	authValue := string(ctx.Request().Header.Peek("Authorization"))

	if authValue == "" {
		ctx.Status(403)

		return ctx.JSON("Undefined token")
	}

	token := parseToken(authValue)

	// first look cached tokens
	hasToken := A.redis.IsHasToken(token)

	if hasToken {
		return ctx.Next()
	}

	grpcCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp, err := A.tokenService.ValidateToken(grpcCtx, &tokenpb.ValidateTokenRequest{Token: token})

	if err != nil {
		ctx.Status(403)

		return ctx.JSON("Invalid token")
	}

	if !resp.IsValid {
		ctx.Status(403)

		return ctx.JSON("Token is not valid")
	}
	return ctx.Next()

}

func parseToken(value string) string {

	token := value[7:]

	return token

}
