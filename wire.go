//go:build wireinject
// +build wireinject

package main

import (
	"contacts/controllers"
	"contacts/jobs"
	"contacts/repositories"
	"contacts/routers"
	"contacts/services"
	"github.com/google/wire"
)

func initializeDependencies() Server {

	wire.Build(
		provideServer,
		routers.ProvideRouter,
		controllers.ProvideUserContactController,
		controllers.ProvideUserController,
		services.ProvideUserContactService,
		services.ProvideUserService,
		repositories.ProvideUserContactRepository,
		repositories.ProvideUserRepository,
		repositories.ProvideGormFromEnv,
		repositories.ProvideUserContactModificationRepository,
		jobs.ProvideMigrationJob,
		jobs.ProvideOnStartUpJob,
	)

	return nil
}
