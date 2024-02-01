package store
import(
    "context"
    "encoding/json"
    "fmt"
	"io"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "net/http"
    "go.mongodb.org/mongo-driver/bson"
)

type MongoStore struct{
	Collection *mongo.Collection
}
var m MongoStore

func(m *MongoStore) OpenConnectionWithMongoDB(){
	// MongoDB connection URI
	mongoURI := "mongodb://localhost:27017"
	// MongoDB database name and collection name
	dbName := "CF-RAS"
	collectionName := "recentActions"

	// Establish connection to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	defer client.Disconnect(context.Background())

	// Access the collection in the database
	m.Collection := client.Database(dbName).Collection(collectionName)

}

func (m *MongoStore) StoreRecentActionsInTheDatabase(actions []models.RecentAction){
	// Parse the JSON response
	var response RecentActionsResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	recentActions:= response.Result

	// Convert recentActions to interface{} slice
	var documents []interface{}
	for _, action := range recentActions {
		documents = append(documents, action)
	}

	// Insert the data into MongoDB
	_, err = collection.InsertMany(context.Background(), documents)
	if err != nil {
		fmt.Println("Error inserting data into MongoDB:", err)
		return
	}

	fmt.Println("Data inserted into MongoDB successfully!")
}






// // OpenConnectionWithMongoDB()
// // MongoDB connection URI
// mongoURI := "mongodb://localhost:27017"
// // MongoDB database name and collection name
// dbName := "CF-RAS"
// collectionName := "recentActions"

// // Establish connection to MongoDB
// client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
// if err != nil {
// 	fmt.Println("Error connecting to MongoDB:", err)
// 	return
// }
// defer client.Disconnect(context.Background())

// // Access the collection in the database
// collection := client.Database(dbName).Collection(collectionName)



// // StoreRecentActionsInTheDatabase(actions []models.RecentAction)
// // Parse the JSON response
// var response RecentActionsResponse
// body, err := io.ReadAll(resp.Body)
// if err != nil {
// 	fmt.Println("Error reading response body:", err)
// 	return
// }
// if err := json.Unmarshal(body, &response); err != nil {
// 	fmt.Println("Error unmarshalling JSON:", err)
// 	return
// }

// recentActions:= response.Result

// // Convert recentActions to interface{} slice
// var documents []interface{}
// for _, action := range recentActions {
// 	documents = append(documents, action)
// }

// // Insert the data into MongoDB
// _, err = collection.InsertMany(context.Background(), documents)
// if err != nil {
// 	fmt.Println("Error inserting data into MongoDB:", err)
// 	return
// }

// fmt.Println("Data inserted into MongoDB successfully!")
// }



// QueryRecentActions() ([]models.RecentAction, error)




// GetMaxTimeStamp() (int64, error)