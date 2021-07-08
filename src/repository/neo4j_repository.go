package repository

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

const (
	queryFollow = "MERGE (u1:User{id:$userId, name:$userId}) MERGE (u2:User{id:$followingId, name:$followingId}) MERGE (u1)-[f:following]->(u2)"
	queryUnfollow = "MERGE (u1:User{id:$userId}) MERGE (u2:User{id:$unfollowId}) MERGE (u1)-[f:following]->(u2) DELETE f"
	queryGetFollowings = "MATCH (u:User{id:$id}) MATCH (u)-[f:following]->(x) WHERE x.id <> $selfId RETURN x.id"
	queryGetFollowings2 = "MATCH (u:User{id:$id}) MATCH (u)-[f:following]->(x) WITH x MATCH (u1:User{id:$selfId}) WHERE NOT EXISTS((u1)-[:following]->(x)) AND x.id <> $selfId return x.id"
)
type neo4jRepository struct {
	Driver neo4j.Driver
}


type Neo4jRepository interface {
	Follow(userId, followingId string) error
	Unfollow(userId, unfollowId string) error
	Recommend(userId string)([]string, error)
}

func NewNeo4jRepository(driver neo4j.Driver) Neo4jRepository {
	return &neo4jRepository{Driver: driver}
}

func (n *neo4jRepository) Follow(userId, followingId string) error {

	session := n.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.Run(queryFollow, map[string]interface{}{"userId": userId, "followingId": followingId})

	return err
}

func (n *neo4jRepository) Unfollow(userId, unfollowId string) error {
	session := n.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.Run(queryUnfollow, map[string]interface{}{"userId": userId, "unfollowId": unfollowId})

	return err
}

func (n *neo4jRepository) Recommend(userId string) ([]string, error) {
	session := n.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	records, err :=session.Run(queryGetFollowings, map[string]interface{}{"id" : userId, "selfId" : userId})

	if err != nil {
		return nil, err
	}

	recordsCollection, err := records.Collect()

	if err != nil {
		return nil, err
	}

	var results []string

	for _, it := range recordsCollection {


		records, err := session.Run(queryGetFollowings2, map[string]interface{}{"id" : it.Values[0], "selfId" : userId})
		if err != nil {
			return nil, err
		}

		recordsCollection, err := records.Collect()
		if err != nil {
			return nil, err
		}




		if len(results) >= 4 {
			return results, nil
		}

		for _, it := range recordsCollection {
			results = append(results, it.Values[0].(string))
			if len(results) >= 4 {
				return results, nil
			}
		}


	}

	return results, err

}



