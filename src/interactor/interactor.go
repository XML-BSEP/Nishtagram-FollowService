package interactor

import (
	"FollowService/http/handler"
	"FollowService/infrastructure/grpc/follow_service/implementation"
	"FollowService/repository"
	"FollowService/usecase"
	logger "github.com/jelena-vlajkov/logger/logger"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"go.mongodb.org/mongo-driver/mongo"
)
type interactor struct {
	db *mongo.Client
	logger *logger.Logger
	neo4JDriver neo4j.Driver
}

func (i *interactor) NewNeo4jRepository() repository.Neo4jRepository {
	return repository.NewNeo4jRepository(i.neo4JDriver)
}

func (i *interactor) NewNeo4jUsecase() repository.Neo4jRepository {
	return usecase.NewNe04jUsecase(i.NewNeo4jRepository())
}

func (i *interactor) NewFollowServiceImpl() *implementation.FollowServiceImpl {
	return implementation.NewFollowServiceImpl(i.NewFollowingUseCase())
}

func (i *interactor) NewRequestRepository() repository.FollowRequestRepo {
	return repository.NewFollowRequestRepo(i.db, i.logger)
}

func (i *interactor) NewRequestUseCase() usecase.FollowRequestUseCase {
	return usecase.NewFollowRequestUseCase(i.NewRequestRepository(), i.NewFollowerRepository(), i.NewFollowingRepository(), i.logger)
}

func (i *interactor) NewRequestHandler() handler.FollowRequestHandler {
	return handler.NewFollowRequestHandler(i.NewRequestUseCase(), i.logger, i.NewNeo4jUsecase())
}





func (i *interactor) NewFollowingRepository() repository.FollowingRepo {
	return repository.NewFollowingRepo(i.db, i.logger)
}

func (i *interactor) NewFollowingUseCase() usecase.FollowingUseCase {
	return usecase.NewFollowingUseCase(i.NewFollowingRepository(), i.NewRequestUseCase(), i.NewFollowerUseCase(), i.NewFollowerRepository(), i.logger)
}

func (i *interactor) NewFollowingHandler() handler.FollowingHandler {
	return handler.NewFollowingHandler(i.NewFollowingUseCase(), i.NewRequestUseCase(), i.logger, i.NewNeo4jUsecase())
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
	NewNeo4jRepository() repository.Neo4jRepository

	NewFollowingUseCase() usecase.FollowingUseCase
	NewFollowerUseCase() usecase.FollowerUseCase
	NewRequestUseCase() usecase.FollowRequestUseCase
	NewNeo4jUsecase() repository.Neo4jRepository

	NewFollowingHandler() handler.FollowingHandler
	NewFollowerHandler() handler.FollowerHandler
	NewRequestHandler() handler.FollowRequestHandler

	NewAppHandler() AppHandler



	NewFollowServiceImpl() *implementation.FollowServiceImpl
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

func NewInteractor(db *mongo.Client, logger *logger.Logger, neo4jDriver neo4j.Driver) Interactor {
	return &interactor{db: db, logger: logger, neo4JDriver: neo4jDriver}
}