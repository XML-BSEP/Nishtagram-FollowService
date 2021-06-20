package usecase

import (
	"FollowService/repository"
	logger "github.com/jelena-vlajkov/logger/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileUseCase interface {
	GetByID(id string) *mongo.SingleResult

}

type profileUseCase struct {
	ProfileRepo repository.ProfileRepo
	logger *logger.Logger
}

func (p profileUseCase) GetByID(id string) *mongo.SingleResult {
	p.logger.Logger.Infof("getting profile by id %v\n", id)
	return p.ProfileRepo.GetByID(id)
}

func NewProfileUseCase(repo repository.ProfileRepo, logger *logger.Logger) ProfileUseCase {
	return &profileUseCase{
		ProfileRepo: repo,
		logger: logger,
	}
}
