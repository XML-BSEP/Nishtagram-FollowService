package repository

import (
	"FollowService/domain"
	"FollowService/dto"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type FollowerRepo interface {
	CreateFollower(follower *domain.ProfileFollower) (*domain.ProfileFollower, error)
	GetByID(id string) *mongo.SingleResult
	Delete(id string) *mongo.DeleteResult
	GetAllUsersFollowers(user dto.ProfileDTO) ([]bson.M, error)
	RemoveFollower(ctx context.Context, unfollow dto.Unfollow) error
	AlreadyFollowing(ctx context.Context, following *domain.ProfileFollowing) (bool, error)

}

type followerRepo struct {
	collection *mongo.Collection
	db *mongo.Client
}

func (f *followerRepo) AlreadyFollowing(ctx context.Context, following *domain.ProfileFollowing) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	/*userBson := bson.M{"_id" : following.User}
	followingBson := bson.M{"_id" : following.Following}*/

	var val *domain.ProfileFollowing
	err := f.collection.FindOne(ctx, bson.M{"user._id" : following.User.ID, "following._id" : following.Following.ID}).Decode(&val)

	if err != nil {
		return false, err
	}

	return true, err
}

func (f *followerRepo) RemoveFollower(ctx context.Context, unfollow dto.Unfollow) error {
	var follower bson.M
	//fmt.Println(unfollow)
	unfollowerBson := bson.M{"_id" :unfollow.UserUnfollowing}
	// ja pratim peru i hocu peru da otpratim
	// u ovom slucaju ja sam userUnfollowing, a pera je UserToUnfollow

	//sto znaci da ja trebam da budem obrisan iz tabele kao follower
	//a pera kao user

	userBson := bson.M{"_id" : unfollow.UserToUnfollow}

	err := f.collection.FindOne(ctx, bson.M{"user": unfollowerBson, "follower":userBson}).Decode(&follower)

	if err !=nil{
		log.Fatal(err)
		return err
	}

	var fusrodah dto.ProfileFollowerDTO

	bsonBytes, _ := bson.Marshal(follower)
	err = bson.Unmarshal(bsonBytes, &fusrodah)
	if err != nil {
		return err
	}


	if f.Delete(fusrodah.ID).DeletedCount==1{
		return nil
	}else{
		err1 := errors.New("deleting error: no followers deleted")
		return err1
	}

}
func (f *followerRepo) CreateFollower(follower *domain.ProfileFollower) (*domain.ProfileFollower, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := f.collection.InsertOne(ctx, *follower)

	if err != nil {
		panic(err)
	}

	return follower, nil
}
func (f followerRepo) GetByID(id string) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result := f.collection.FindOne(ctx, bson.M{"_id": id})
	return result
}

func (f followerRepo) Delete(id string) *mongo.DeleteResult {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := f.collection.DeleteOne(ctx, bson.M{"_id": id})
	if result.DeletedCount==0{
		objID, err := primitive.ObjectIDFromHex(id)
		if err!=nil{
			return result
		}
		result, err = f.collection.DeleteOne(ctx, bson.M{"_id" :objID})

	}

	if err != nil {
		return &mongo.DeleteResult{DeletedCount: 0}

		//log.Fatal("DeleteOne() ERROR:", err)
	}

	return result
}


func (f followerRepo) GetAllUsersFollowers(user dto.ProfileDTO) ([]bson.M, error){
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filterCursor, err := f.collection.Find(ctx, bson.M{"user": user})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var usersFollowersBson []bson.M
	if err = filterCursor.All(ctx, &usersFollowersBson); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return usersFollowersBson, nil
}

func NewFollowerRepo(db *mongo.Client) FollowerRepo {
	return &followerRepo{
		db: db,
		collection : db.Database("follow_db").Collection("profile_followers"),

	}
}