package interactor

import (
	"FollowService/http/handler"
	"FollowService/repository"
	"FollowService/usecase"
	logger "github.com/jelena-vlajkov/logger/logger"
	"go.mongodb.org/mongo-driver/mongo"
)
type interactor struct {
	db *mongo.Client
	logger *logger.Logger
}

func (i *interactor) NewRequestRepository() repository.FollowRequestRepo {
	return repository.NewFollowRequestRepo(i.db, i.logger)
}

func (i *interactor) NewRequestUseCase() usecase.FollowRequestUseCase {
	return usecase.NewFollowRequestUseCase(i.NewRequestRepository(), i.NewFollowerRepository(), i.NewFollowingRepository(), i.logger)
}

func (i *interactor) NewRequestHandler() handler.FollowRequestHandler {
	return handler.NewFollowRequestHandler(i.NewRequestUseCase(), i.logger)
}





func (i *interactor) NewFollowingRepository() repository.FollowingRepo {
	return repository.NewFollowingRepo(i.db, i.logger)
}

func (i *interactor) NewFollowingUseCase() usecase.FollowingUseCase {
	return usecase.NewFollowingUseCase(i.NewFollowingRepository(), i.NewRequestUseCase(), i.NewFollowerUseCase(), i.NewFollowerRepository(), i.logger)
}

func (i *interactor) NewFollowingHandler() handler.FollowingHandler {
	return handler.NewFollowingHandler(i.NewFollowingUseCase(), i.NewRequestUseCase(), i.logger)
}





func (i *interactor) NewFollowerRepository() repository.FollowerRepo {
	return repository.NewFollowerRepo(i.db, i.logger)
}

func (i *interactor) NewFollowerUseCase() usecase.FollowerUseCase {
	return usecase.NewFollowerUseCase(i.NewFollowerRepository(), i.logger)
}

func (i *interactor) NewFollowerHandler() handler.FollowerHandler {
	return handler.NewFollowerHandler(i.NewFollowerUseCase(), i.logger)
}




func (i *interactor) NewAppHandler() AppHandler {
	appHandler := &appHandler{}
	appHandler.FollowingHandler = i.NewFollowingHandler()
	appHandler.FollowerHandler = i.NewFollowerHandler()
	appHandler.FollowRequestHandler = i.NewRequestHandler()

	return appHandler

}

type Interactor interface {

	NewFollowingRepository() repository.FollowingRepo
	NewFollowerRepository() repository.FollowerRepo
	NewRequestRepository() repository.FollowRequestRepo

	NewFollowingUseCase() usecase.FollowingUseCase
	NewFollowerUseCase() usecase.FollowerUseCase
	NewRequestUseCase() usecase.FollowRequestUseCase

	NewFollowingHandler() handler.FollowingHandler
	NewFollowerHandler() handler.FollowerHandler
	NewRequestHandler() handler.FollowRequestHandler

	NewAppHandler() AppHandler

}

type appHandler struct {
	handler.FollowingHandler
	handler.FollowerHandler
	handler.FollowRequestHandler
}

type AppHandler interface {
	handler.FollowingHandler
	handler.FollowerHandler
	handler.FollowRequestHandler


}

func NewInteractor(db *mongo.Client, logger *logger.Logger) Interactor {
	return &interactor{db: db, logger: logger}
}