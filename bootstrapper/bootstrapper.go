package bootstrapper

import (
	"context"
	"log"
	"log/slog"
	"main/api/restapi"
	"main/api/restapi/operations"
	"main/config"
	"main/internal/controller"
	"main/internal/database"
	"main/internal/handlers"
	logger "main/internal/lib/logger"
	repo "main/internal/repositories"
	"main/internal/service"

	"github.com/go-openapi/loads"
	"github.com/go-playground/validator/v10"
)

var Secret string = ""

type RootBootstrapper struct {
	Infrastructure struct {
		Logger *slog.Logger
		Server *restapi.Server
		DB     database.DB
	}
	Controller controller.Controller
	Config     *config.Config
	Handlers   handlers.Handlers
	Repository repo.Repository
	Service    service.Service

	Validator *validator.Validate
}

type RootBoot interface {
	registerRepositoriesAndServices(ctx context.Context, db database.DB)
	registerAPIServer(cfg config.Config) error
	RunAPI() error
}

func New() RootBoot {
	return &RootBootstrapper{
		Config: config.NewConfig(),
	}
}

func (r *RootBootstrapper) RunAPI() error {
	ctx := context.Background()
	r.Infrastructure.Logger = logger.NewLogger(r.Config.LogLevel)

	r.registerRepositoriesAndServices(ctx, r.Infrastructure.DB)
	err := r.registerAPIServer(*r.Config)
	if err != nil {
		return err
	}

	return err
}

func (r *RootBootstrapper) registerRepositoriesAndServices(ctx context.Context, db database.DB) {
	logger := r.Infrastructure.Logger
	r.Infrastructure.DB = database.NewDB().NewConn(ctx, *r.Config, logger)
	r.Repository = repo.NewUserRepo(r.Infrastructure.DB, logger)
	r.Service = service.New(r.Repository, logger)
}

func (r *RootBootstrapper) registerAPIServer(cfg config.Config) error {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return err
	}

	api := operations.NewMusicLibraryAPI(swaggerSpec)

	logger := r.Infrastructure.Logger

	r.Controller = controller.New(r.Service, logger)

	r.Validator = validator.New(validator.WithRequiredStructEnabled())

	r.Handlers = handlers.New(r.Controller, r.Validator, logger)
	r.Handlers.Link(api)
	if r.Handlers == nil {
		return err
	}

	r.Infrastructure.Server = restapi.NewServer(api)
	r.Infrastructure.Server.Port = cfg.ServerPort
	r.Infrastructure.Server.ConfigureAPI()
	if err := r.Infrastructure.Server.Serve(); err != nil {
		log.Fatalln(err)
	}

	return err
}
