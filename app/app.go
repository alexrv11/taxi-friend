package app

import (
	"alex/taxi-server/config"
	"alex/taxi-server/handlers"
	"alex/taxi-server/repository"
	"alex/taxi-server/services"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/labstack/echo"
)

func Start(e *echo.Echo){
	configuration := config.LoadConfig()
	mapper := BuildMapper(configuration)

	initRouter(e, mapper)

	e.Logger.Fatal(e.Start(":1323"))
}

func initRouter(router *echo.Echo, mapper *Mapper) {
	routingAPI(router, mapper)
}

func BuildMapper(configuration *config.ApiConfig) *Mapper {

	db := ConnectDatabase(configuration)

	repoBuilder := BuildRepositoryFactory(db)

	driverService := services.NewDriver(repoBuilder)
	driverHandler := handlers.NewDriver(driverService)
	driverMapper := NewDriverMapper(driverHandler)

	qrService := &services.Qr{QrRepository: repoBuilder.CreateQrRepository()}
	qrHandler := &handlers.Qr{ QrService: qrService}
	qrMapper := NewQrMapper(qrHandler)

	orderService := &services.Order{ OrderRepository: repoBuilder.CreateOrderRepository()}
	orderHandler := &handlers.Order{ OrderService: orderService }
	orderMapper := NewOrderMapper(orderHandler)

	mapper := NewMapper(driverMapper, qrMapper, orderMapper)

	return mapper
}

func ConnectDatabase(configuration *config.ApiConfig) *dynamo.DB {
	var cfg aws.Config
	cfg.WithEndpoint(configuration.DB.Endpoint).
		WithLogLevel(aws.LogDebugWithHTTPBody)
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:     cfg,
		//Profile: configuration.DB.Profile, /*run locally*/
		SharedConfigState: session.SharedConfigEnable,
	}))

	db := dynamo.New(sess, cfg.WithEndpoint(configuration.DB.Endpoint).WithLogLevel(aws.LogDebugWithHTTPBody))

	return db
}

func BuildRepositoryFactory(db *dynamo.DB) repository.IFactory {
	return &repository.Factory{DB: db}
}