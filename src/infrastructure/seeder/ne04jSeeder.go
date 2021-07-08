package seeder

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

const (
	deleteAll = "MATCH (n) DETACH DELETE n"
	follow = "MERGE (u1:User{id:$userId, name:$userId}) MERGE (u2:User{id:$followingId, name:$followingId}) MERGE (u1)-[f:following]->(u2)"
)

func Seed(driver neo4j.Driver) error {
	session:= driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.Run(deleteAll, nil)
	if err != nil {
		return err
	}
	return seedFollowings(driver)
}

func seedFollowings(driver neo4j.Driver) error{
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	user1 := "e2b5f92e-c31b-11eb-8529-0242ac130003"
	user2 := "424935b1-766c-4f99-b306-9263731518bc"
	user3 := "a2c2f993-dc32-4a82-82ed-a5f6866f7d03"
	user4 := "43420055-3174-4c2a-9823-a8f060d644c3"
	user5 := "ead67925-e71c-43f4-8739-c3b823fe21bb"
	user6 := "23ddb1dd-4303-428b-b506-ff313071d5d7"

	if err := seedFollowing(driver, user2, user1); err != nil {
		return err
	}

	if err := seedFollowing(driver, user3, user1); err != nil {
		return err
	}

	if err := seedFollowing(driver, user4, user1); err != nil {
		return err
	}

	if err := seedFollowing(driver, user5, user1); err != nil {
		return err
	}

	if err := seedFollowing(driver, user6, user1); err != nil {
		return err
	}

	if err := seedFollowing(driver, user1, user2); err != nil {
		return err
	}

	if err := seedFollowing(driver, user1, user3); err != nil {
		return err
	}

	if err := seedFollowing(driver, user1, user4); err != nil {
		return err
	}

	if err := seedFollowing(driver, user1, user5); err != nil {
		return err
	}

	if err := seedFollowing(driver, user1, user6); err != nil {
		return err
	}
	return nil
}

func seedFollowing(driver neo4j.Driver, userId, followingId string) error {
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.Run(follow, map[string]interface{}{"userId": userId, "followingId": followingId})

	return err


}