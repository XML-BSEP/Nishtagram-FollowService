package usecase

import "FollowService/repository"

type neo4jUsecase struct {
	Neo4jRepository repository.Neo4jRepository
}

type Neo4jUsecase interface {
	Follow(userId, followingId string) error
	Unfollow(userId, unfollowId string) error
	Recommend(userId string)([]string, error)
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

func (n *neo4jUsecase) Recommend(userId string) ([]string, error) {
	return n.Neo4jRepository.Recommend(userId)
}
