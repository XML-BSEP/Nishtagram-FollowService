package usecase

import (
	"FollowService/dto"
	"FollowService/gateway"
	"FollowService/repository"
	"context"
)

type neo4jUsecase struct {
	Neo4jRepository repository.Neo4jRepository
}

type Neo4jUsecase interface {
	Follow(userId, followingId string) error
	Unfollow(userId, unfollowId string) error
	Recommend(userId string)([]dto.SearchUserDTO, error)
}

func NewNe04jUsecase(neo4jRepository repository.Neo4jRepository) Neo4jUsecase {
	return &neo4jUsecase{Neo4jRepository: neo4jRepository}
}

func (n *neo4jUsecase) Follow(userId, followingId string) error {
	return n.Neo4jRepository.Follow(userId, followingId)
}

func (n *neo4jUsecase) Unfollow(userId, unfollowId string) error {
	return n.Neo4jRepository.Unfollow(userId, unfollowId)
}

func (n *neo4jUsecase) Recommend(userId string) ([]dto.SearchUserDTO, error) {
	recommendations, err := n.Neo4jRepository.Recommend(userId)
	if err != nil {
		return nil, err
	}
	userIds := dto.UserIdsDto{Ids: recommendations}
	userInfos, err := gateway.GetSearchResults(context.Background(), userIds)
	if err != nil {
		return nil, err
	}
	return userInfos, err


}
