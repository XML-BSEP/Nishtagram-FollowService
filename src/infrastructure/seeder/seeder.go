package seeder

import (
	"FollowService/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

func SeedData(db string, mongoCli *mongo.Client, ctx *context.Context){

	if cnt,_ := mongoCli.Database(db).Collection("profiles").EstimatedDocumentCount(*ctx, nil); cnt == 0{
		profileCollection := mongoCli.Database(db).Collection("profiles")
		seedProfiles(profileCollection, ctx)
	}

	if cnt,_ := mongoCli.Database(db).Collection("profile_following").EstimatedDocumentCount(*ctx, nil); cnt == 0{
		followingCollection := mongoCli.Database(db).Collection("profile_following")
		seedProfileFollowing(followingCollection, ctx)
	}

	if cnt,_ := mongoCli.Database(db).Collection("profile_followers").EstimatedDocumentCount(*ctx, nil); cnt == 0{
		followersCollection := mongoCli.Database(db).Collection("profile_followers")
		seedProfileFollowers(followersCollection, ctx)
	}

	if cnt,_ := mongoCli.Database(db).Collection("follow_request").EstimatedDocumentCount(*ctx, nil); cnt == 0{
		followRequestCollection := mongoCli.Database(db).Collection("follow_request")
		seedFollowRequest(followRequestCollection, ctx)
	}

}

func seedProfiles(profileCollection *mongo.Collection, ctx *context.Context){
	//for i := 0; i < 10; i++ {
	//	var profile domain.Profile
	//	insertResult, err := profileCollection.InsertOne(context.TODO(), profile)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	profile.ID = insertResult.InsertedID.(primitive.ObjectID).Hex()
	//
	//	_, err = profileCollection.UpdateOne(
	//		context.TODO(),
	//		bson.M{"_id": insertResult.InsertedID.(primitive.ObjectID)},
	//		bson.D{
	//			{"$set", profile},
	//		},
	//	)
	//}


	_, err := profileCollection.InsertMany(*ctx, []interface{}{
		bson.D{
			{"_id", "123451"},
		},
		bson.D{
			{"_id", "123452"},
		},
		bson.D{
			{"_id", "123453"},
		},
		bson.D{
			{"_id", "123454"},
		},
		bson.D{
			{"_id", "123455"},
		},
		bson.D{
			{"_id", "123456"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func seedFollowRequest(followRequestCollection *mongo.Collection, ctx *context.Context){
	profil2 := domain.Profile{ID: "123452"}
	profil3 := domain.Profile{ID: "123453"}
	profil4 := domain.Profile{ID: "123454"}
	profil5 := domain.Profile{ID: "123455"}
	profil6 := domain.Profile{ID: "123456"}

	_, err := followRequestCollection.InsertMany(*ctx, []interface{}{
		bson.D{
			{"_id", "1231"},
			{"Timestamp" , time.Now()},
			{"UserRequested", profil2},
			{"FollowedAccount", profil5},
		},

		bson.D{
			{"_id", "1232"},
			{"Timestamp" , time.Now()},
			{"UserRequested", profil3},
			{"FollowedAccount", profil4},
		},
		bson.D{
			{"_id", "1233"},
			{"Timestamp" , time.Now()},
			{"UserRequested", profil4},
			{"FollowedAccount", profil3},
		},
		bson.D{
			{"_id", "1234"},
			{"Timestamp" , time.Now()},
			{"UserRequested", profil6},
			{"FollowedAccount", profil2},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func seedProfileFollowing(profileFollowingCollection *mongo.Collection, ctx *context.Context){
	profil1 := domain.Profile{ID: "123451"}
	profil2 := domain.Profile{ID: "123452"}
	profil3 := domain.Profile{ID: "123453"}
	profil4 := domain.Profile{ID: "123454"}
	profil5 := domain.Profile{ID: "123455"}
	profil6 := domain.Profile{ID: "123456"}

	_, err := profileFollowingCollection.InsertMany(*ctx, []interface{}{
		bson.D{
			{"_id", "12341"},
			{"CloseFriend",true},
			{"Timestamp" , time.Now()},
			{"User", profil1},
			{"Following", profil2},
		},
		bson.D{
			{"_id", "12342"},
			{"CloseFriend",true},
			{"Timestamp" , time.Now()},
			{"User", profil1},
			{"Following", profil3},
		},
		bson.D{
			{"_id", "12343"},
			{"CloseFriend",true},

			{"Timestamp" , time.Now()},
			{"User", profil1},
			{"Following", profil4},
		},
		bson.D{
			{"_id", "12344"},
			{"CloseFriend",false},
			{"Timestamp" , time.Now()},
			{"User", profil1},
			{"Following", profil5},
		},
		bson.D{
			{"_id", "12345"},
			{"CloseFriend",false},
			{"Timestamp" , time.Now()},
			{"User", profil1},
			{"Following", profil6},
		},
		bson.D{
			{"_id", "12346"},
			{"CloseFriend",false},
			{"Timestamp" , time.Now()},
			{"User", profil2},
			{"Following", profil1},
		},
		bson.D{
			{"_id", "12347"},
			{"CloseFriend",true},
			{"Timestamp" , time.Now()},
			{"User", profil2},
			{"Following", profil3},
		},
		bson.D{
			{"_id", "12348"},
			{"CloseFriend",true},
			{"Timestamp" , time.Now()},
			{"User", profil2},
			{"Following", profil4},
		},
		bson.D{
			{"_id", "12349"},
			{"CloseFriend",true},
			{"Timestamp" , time.Now()},
			{"User", profil3},
			{"Following", profil1},
		},
		bson.D{
			{"_id", "12350"},
			{"CloseFriend",true},
			{"Timestamp" , time.Now()},
			{"User", profil3},
			{"Following", profil5},
		},
		bson.D{
			{"_id", "12351"},
			{"CloseFriend",false},
			{"Timestamp" , time.Now()},
			{"User", profil3},
			{"Following", profil6},
		},
		bson.D{
			{"_id", "12352"},
			{"CloseFriend",false},
			{"Timestamp" , time.Now()},
			{"User", profil4},
			{"Following", profil1},
		},
		bson.D{
			{"_id", "12353"},
			{"CloseFriend",false},
			{"Timestamp" , time.Now()},
			{"User", profil4},
			{"Following", profil2},
		},
		bson.D{
			{"_id", "12354"},
			{"CloseFriend",false},
			{"Timestamp" , time.Now()},
			{"User", profil4},
			{"Following", profil6},
		},

		bson.D{
			{"_id", "12355"},
			{"CloseFriend",true},
			{"Timestamp" , time.Now()},
			{"User", profil5},
			{"Following", profil1},
		},
		bson.D{
			{"_id", "12356"},
			{"CloseFriend",false},
			{"Timestamp" , time.Now()},
			{"User", profil5},
			{"Following", profil2},
		},
		bson.D{
			{"_id", "12357"},
			{"CloseFriend",true},
			{"Timestamp" , time.Now()},
			{"User", profil5},
			{"Following", profil3},
		},
		bson.D{
			{"_id", "12358"},
			{"CloseFriend",false},
			{"Timestamp" , time.Now()},
			{"User", profil5},
			{"Following", profil4},
		},
		bson.D{
			{"_id", "12359"},
			{"CloseFriend",true},
			{"Timestamp" , time.Now()},
			{"User", profil5},
			{"Following", profil6},
		},

		bson.D{
			{"_id", "12360"},
			{"CloseFriend",false},
			{"Timestamp" , time.Now()},
			{"User", profil6},
			{"Following", profil1},
		},
		bson.D{
			{"_id", "12361"},
			{"CloseFriend",true},
			{"Timestamp" , time.Now()},
			{"User", profil6},
			{"Following", profil3},
		},
		bson.D{
			{"_id", "12362"},
			{"CloseFriend",true},
			{"Timestamp" , time.Now()},
			{"User", profil6},
			{"Following", profil4},
		},

	})
	if err != nil {
		log.Fatal(err)
	}
}

func seedProfileFollowers(followers_collection *mongo.Collection, ctx *context.Context) {
	profil1 := domain.Profile{ID: "123451"}
	profil2 := domain.Profile{ID: "123452"}
	profil3 := domain.Profile{ID: "123453"}
	profil4 := domain.Profile{ID: "123454"}
	profil5 := domain.Profile{ID: "123455"}
	profil6 := domain.Profile{ID: "123456"}

	_, err := followers_collection.InsertMany(*ctx, []interface{}{
		bson.D{
			{"_id", "1234561"},
			{"Timestamp" , time.Now()},
			{"User", profil1},
			{"Follower", profil2},
		},bson.D{
			{"_id", "1234562"},
			{"Timestamp" , time.Now()},
			{"User", profil1},
			{"Follower", profil3},
		},bson.D{
			{"_id", "1234563"},
			{"Timestamp" , time.Now()},
			{"User", profil1},
			{"Follower", profil4},
		},bson.D{
			{"_id", "1234564"},
			{"Timestamp" , time.Now()},
			{"User", profil1},
			{"Follower", profil5},
		},bson.D{
			{"_id", "1234565"},
			{"Timestamp" , time.Now()},
			{"User", profil1},
			{"Follower", profil6},
		},

		bson.D{
			{"_id", "1234566"},
			{"Timestamp" , time.Now()},
			{"User", profil2},
			{"Follower", profil1},
		},bson.D{
			{"_id", "1234568"},
			{"Timestamp" , time.Now()},
			{"User", profil2},
			{"Follower", profil4},
		},bson.D{
			{"_id", "1234569"},
			{"Timestamp" , time.Now()},
			{"User", profil2},
			{"Follower", profil5},
		},

		bson.D{
			{"_id", "1234570"},
			{"Timestamp" , time.Now()},
			{"User", profil3},
			{"Follower", profil1},
		},bson.D{
			{"_id", "1234571"},
			{"Timestamp" , time.Now()},
			{"User", profil3},
			{"Follower", profil2},
		},bson.D{
			{"_id", "1234572"},
			{"Timestamp" , time.Now()},
			{"User", profil3},
			{"Follower", profil5},
		},bson.D{
			{"_id", "1234573"},
			{"Timestamp" , time.Now()},
			{"User", profil3},
			{"Follower", profil6},
		},

		bson.D{
			{"_id", "1234574"},
			{"Timestamp" , time.Now()},
			{"User", profil4},
			{"Follower", profil1},
		},bson.D{
			{"_id", "1234575"},
			{"Timestamp" , time.Now()},
			{"User", profil4},
			{"Follower", profil2},
		},bson.D{
			{"_id", "1234576"},
			{"Timestamp" , time.Now()},
			{"User", profil4},
			{"Follower", profil5},
		},bson.D{
			{"_id", "1234577"},
			{"Timestamp" , time.Now()},
			{"User", profil4},
			{"Follower", profil6},
		},

		bson.D{
			{"_id", "1234578"},
			{"Timestamp" , time.Now()},
			{"User", profil5},
			{"Follower", profil1},
		},bson.D{
			{"_id", "1234580"},
			{"Timestamp" , time.Now()},
			{"User", profil5},
			{"Follower", profil3},
		},

		bson.D{
			{"_id", "1234583"},
			{"Timestamp" , time.Now()},
			{"User", profil6},
			{"Follower", profil1},
		},bson.D{
			{"_id", "1234584"},
			{"Timestamp" , time.Now()},
			{"User", profil6},
			{"Follower", profil3},
		},bson.D{
			{"_id", "1234585"},
			{"Timestamp" , time.Now()},
			{"User", profil6},
			{"Follower", profil4},
		},bson.D{
			{"_id", "1234586"},
			{"Timestamp" , time.Now()},
			{"User", profil6},
			{"Follower", profil5},
		},
		})

	if err != nil {
		log.Fatal(err)
	}
}