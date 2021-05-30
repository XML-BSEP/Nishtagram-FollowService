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

	dropDatabase(db,mongoCli, ctx)

	if cnt,_ := mongoCli.Database(db).Collection("profiles").EstimatedDocumentCount(*ctx, nil); cnt == 0{
		profileCollection := mongoCli.Database(db).Collection("profiles")
		seedProfiles(profileCollection, ctx)
	}

	if cnt,_ := mongoCli.Database(db).Collection("profile_following").EstimatedDocumentCount(*ctx, nil); cnt == 0{
		followingCollection := mongoCli.Database(db).Collection("profile_followings")
		seedProfileFollowing(followingCollection, ctx)
	}

	if cnt,_ := mongoCli.Database(db).Collection("profile_followers").EstimatedDocumentCount(*ctx, nil); cnt == 0{
		followersCollection := mongoCli.Database(db).Collection("profile_followers")
		seedProfileFollowers(followersCollection, ctx)
	}

	if cnt,_ := mongoCli.Database(db).Collection("follow_request").EstimatedDocumentCount(*ctx, nil); cnt == 0{
		followRequestCollection := mongoCli.Database(db).Collection("follow_requests")
		seedFollowRequest(followRequestCollection, ctx)
	}

}
func dropDatabase(db string, mongoCli *mongo.Client, ctx *context.Context){
	err := mongoCli.Database(db).Drop(*ctx)
	if err != nil {
		return
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
			{"private", true},
		},
		bson.D{
			{"_id", "123452"},
			{"private", true},
		},
		bson.D{
			{"_id", "123453"},
			{"private", true},

		},
		bson.D{
			{"_id", "123454"},
			{"private", false},

		},
		bson.D{
			{"_id", "123455"},
			{"private", false},
		},
		bson.D{
			{"_id", "123456"},
			{"private", false},

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
			{"user_requested", profil2},
			{"followed_account", profil5},
		},

		bson.D{
			{"_id", "1232"},
			{"Timestamp" , time.Now()},
			{"user_requested", profil3},
			{"followed_account", profil4},
		},
		bson.D{
			{"_id", "1233"},
			{"Timestamp" , time.Now()},
			{"user_requested", profil4},
			{"followed_account", profil3},
		},
		bson.D{
			{"_id", "1234"},
			{"Timestamp" , time.Now()},
			{"user_requested", profil6},
			{"followed_account", profil2},
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
			{"timestamp" , time.Now()},
			{"user", profil1},
			{"following", profil2},
		},
		bson.D{
			{"_id", "12342"},
			{"timestamp" , time.Now()},
			{"user", profil1},
			{"following", profil3},
		},
		bson.D{
			{"_id", "12343"},
			{"timestamp" , time.Now()},
			{"user", profil1},
			{"following", profil4},
		},
		bson.D{
			{"_id", "12344"},
			{"timestamp" , time.Now()},
			{"user", profil1},
			{"following", profil5},
		},
		bson.D{
			{"_id", "12345"},
			{"timestamp" , time.Now()},
			{"user", profil1},
			{"following", profil6},
		},
		bson.D{
			{"_id", "12346"},
			{"timestamp" , time.Now()},
			{"user", profil2},
			{"following", profil1},
		},
		bson.D{
			{"_id", "12347"},
			{"timestamp" , time.Now()},
			{"user", profil2},
			{"following", profil3},
		},
		bson.D{
			{"_id", "12348"},
			{"timestamp" , time.Now()},
			{"user", profil2},
			{"following", profil4},
		},
		bson.D{
			{"_id", "12349"},
			{"timestamp" , time.Now()},
			{"user", profil3},
			{"following", profil1},
		},
		bson.D{
			{"_id", "12350"},
			{"timestamp" , time.Now()},
			{"user", profil3},
			{"following", profil5},
		},
		bson.D{
			{"_id", "12351"},
			{"timestamp" , time.Now()},
			{"user", profil3},
			{"following", profil6},
		},
		bson.D{
			{"_id", "12352"},
			{"timestamp" , time.Now()},
			{"user", profil4},
			{"following", profil1},
		},
		bson.D{
			{"_id", "12353"},
			{"timestamp" , time.Now()},
			{"user", profil4},
			{"following", profil2},
		},
		bson.D{
			{"_id", "12354"},
			{"timestamp" , time.Now()},
			{"user", profil4},
			{"following", profil6},
		},

		bson.D{
			{"_id", "12355"},
			{"timestamp" , time.Now()},
			{"user", profil5},
			{"following", profil1},
		},
		bson.D{
			{"_id", "12356"},
			{"timestamp" , time.Now()},
			{"user", profil5},
			{"following", profil2},
		},
		bson.D{
			{"_id", "12357"},
			{"timestamp" , time.Now()},
			{"user", profil5},
			{"following", profil3},
		},
		bson.D{
			{"_id", "12358"},
			{"timestamp" , time.Now()},
			{"user", profil5},
			{"following", profil4},
		},
		bson.D{
			{"_id", "12359"},
			{"timestamp" , time.Now()},
			{"user", profil5},
			{"following", profil6},
		},

		bson.D{
			{"_id", "12360"},
			{"timestamp" , time.Now()},
			{"user", profil6},
			{"following", profil1},
		},
		bson.D{
			{"_id", "12361"},
			{"timestamp" , time.Now()},
			{"user", profil6},
			{"following", profil3},
		},
		bson.D{
			{"_id", "12362"},
			{"timestamp" , time.Now()},
			{"user", profil6},
			{"following", profil4},
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
			{"close_friend",true},
			{"timestamp" , time.Now()},
			{"user", profil1},
			{"follower", profil2},
		},bson.D{
			{"_id", "1234562"},
			{"close_friend",true},
			{"timestamp" , time.Now()},
			{"user", profil1},
			{"follower", profil3},
		},bson.D{
			{"_id", "1234563"},
			{"close_friend",false},
			{"timestamp" , time.Now()},
			{"user", profil1},
			{"follower", profil4},
		},bson.D{
			{"_id", "1234564"},
			{"close_friend",true},
			{"timestamp" , time.Now()},
			{"user", profil1},
			{"follower", profil5},
		},bson.D{
			{"_id", "1234565"},
			{"close_friend",true},
			{"timestamp" , time.Now()},
			{"user", profil1},
			{"follower", profil6},
		},

		bson.D{
			{"_id", "1234566"},
			{"close_friend",false},
			{"timestamp" , time.Now()},
			{"user", profil2},
			{"follower", profil1},
		},bson.D{
			{"_id", "1234568"},
			{"close_friend",false},
			{"timestamp" , time.Now()},
			{"user", profil2},
			{"follower", profil4},
		},bson.D{
			{"_id", "1234569"},
			{"close_friend",true},
			{"timestamp" , time.Now()},
			{"user", profil2},
			{"follower", profil5},
		},

		bson.D{
			{"_id", "1234570"},
			{"close_friend",true},
			{"timestamp" , time.Now()},
			{"user", profil3},
			{"follower", profil1},
		},bson.D{
			{"_id", "1234571"},
			{"close_friend",true},
			{"timestamp" , time.Now()},
			{"user", profil3},
			{"follower", profil2},
		},bson.D{
			{"_id", "1234572"},
			{"close_friend",false},
			{"timestamp" , time.Now()},
			{"user", profil3},
			{"follower", profil5},
		},bson.D{
			{"_id", "1234573"},
			{"close_friend",true},
			{"timestamp" , time.Now()},
			{"user", profil3},
			{"follower", profil6},
		},

		bson.D{
			{"_id", "1234574"},
			{"close_friend",false},
			{"timestamp" , time.Now()},
			{"user", profil4},
			{"follower", profil1},
		},bson.D{
			{"_id", "1234575"},
			{"close_friend",true},
			{"timestamp" , time.Now()},
			{"user", profil4},
			{"follower", profil2},
		},bson.D{
			{"_id", "1234576"},
			{"close_friend",true},
			{"timestamp" , time.Now()},
			{"user", profil4},
			{"follower", profil5},
		},bson.D{
			{"_id", "1234577"},
			{"close_friend",false},
			{"timestamp" , time.Now()},
			{"user", profil4},
			{"follower", profil6},
		},

		bson.D{
			{"_id", "1234578"},
			{"timestamp" , time.Now()},
			{"user", profil5},
			{"follower", profil1},
		},bson.D{
			{"_id", "1234580"},
			{"timestamp" , time.Now()},
			{"user", profil5},
			{"follower", profil3},
		},

		bson.D{
			{"_id", "1234583"},
			{"close_friend",false},
			{"timestamp" , time.Now()},
			{"user", profil6},
			{"follower", profil1},
		},bson.D{
			{"_id", "1234584"},
			{"close_friend",true},
			{"timestamp" , time.Now()},
			{"user", profil6},
			{"follower", profil3},
		},bson.D{
			{"_id", "1234585"},
			{"close_friend",true},
			{"timestamp" , time.Now()},
			{"user", profil6},
			{"follower", profil4},
		},bson.D{
			{"_id", "1234586"},
			{"close_friend",false},
			{"timestamp" , time.Now()},
			{"user", profil6},
			{"follower", profil5},
		},
		})

	if err != nil {
		log.Fatal(err)
	}
}