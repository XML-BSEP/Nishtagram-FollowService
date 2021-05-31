package usecase

import (
	"FollowService/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileUseCase interface {
	GetByID(id string) *mongo.SingleResult
}

type profileUseCase struct {
	ProfileRepo repository.ProfileRepo
}

func (p profileUseCase) GetByID(id string) *mongo.SingleResult {
	return p.ProfileRepo.GetByID(id)
}

func NewProfileUseCase(repo repository.ProfileRepo) ProfileUseCase {
	return &profileUseCase{
		ProfileRepo: repo,
	}
}
