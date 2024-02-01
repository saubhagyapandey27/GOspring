package main
import (
    "context"
    "encoding/json"
    "fmt"
	"io"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "net/http"
    "go.mongodb.org/mongo-driver/bson"
)

// RecentActions represents recent actions data model.
type RecentActions struct {
    TimeSeconds int       `json:"timeSeconds" bson:"timeSeconds"`
    BlogEntry   BlogEntry `json:"blogEntry" bson:"blogEntry"`
    Comment     Comment   `json:"comment" bson:"comment"`
}

// BlogEntry represents blog entry data model.
type BlogEntry struct {
    ID                     int      `json:"id" bson:"_id"`
    OriginalLocale         string   `json:"originalLocale" bson:"originalLocale"`
    CreationTimeSeconds    int      `json:"creationTimeSeconds" bson:"creationTimeSeconds"`
    AuthorHandle           string   `json:"authorHandle" bson:"authorHandle"`
    Title                  string   `json:"title" bson:"title"`
    Content                string   `json:"content" bson:"content"`
    Locale                 string   `json:"locale" bson:"locale"`
    ModificationTimeSeconds int      `json:"modificationTimeSeconds" bson:"modificationTimeSeconds"`
    AllowViewHistory       bool     `json:"allowViewHistory" bson:"allowViewHistory"`
    Tags                   []string `json:"tags" bson:"tags"`
    Rating                 int      `json:"rating" bson:"rating"`
}

// Comment represents comment data model.
type Comment struct {
    ID                  int    `json:"id" bson:"_id"`
    CreationTimeSeconds int    `json:"creationTimeSeconds" bson:"creationTimeSeconds"`
    CommentatorHandle   string `json:"commentatorHandle" bson:"commentatorHandle"`
    Locale              string `json:"locale" bson:"locale"`
    Text                string `json:"text" bson:"text"`
    ParentCommentID     int    `json:"parentCommentId,omitempty" bson:"parentCommentId,omitempty"`
    Rating              int    `json:"rating" bson:"rating"`
}

type RecentActionsResponse struct {
    Status string          `json:"status"`
    Result []RecentActions `json:"result"`
}

// User represents user data model.
type User struct {
    Handle                  string `json:"handle" bson:"_id"`
    Email                   string `json:"email,omitempty" bson:"email,omitempty"`
    VkID                    string `json:"vkId,omitempty" bson:"vkId,omitempty"`
    OpenID                  string `json:"openId,omitempty" bson:"openId,omitempty"`
    FirstName               string `json:"firstName,omitempty" bson:"firstName,omitempty"`
    LastName                string `json:"lastName,omitempty" bson:"lastName,omitempty"`
    Country                 string `json:"country,omitempty" bson:"country,omitempty"`
    City                    string `json:"city,omitempty" bson:"city,omitempty"`
    Organization            string `json:"organization,omitempty" bson:"organization,omitempty"`
    Contribution            int    `json:"contribution,omitempty" bson:"contribution,omitempty"`
    Rank                    string `json:"rank,omitempty" bson:"rank,omitempty"`
    Rating                  int    `json:"rating,omitempty" bson:"rating,omitempty"`
    MaxRank                 string `json:"maxRank,omitempty" bson:"maxRank,omitempty"`
    MaxRating               int    `json:"maxRating,omitempty" bson:"maxRating,omitempty"`
    LastOnlineTimeSeconds   int    `json:"lastOnlineTimeSeconds,omitempty" bson:"lastOnlineTimeSeconds,omitempty"`
    RegistrationTimeSeconds int    `json:"registrationTimeSeconds,omitempty" bson:"registrationTimeSeconds,omitempty"`
    FriendOfCount           int    `json:"friendOfCount,omitempty" bson:"friendOfCount,omitempty"`
    Avatar                  string `json:"avatar,omitempty" bson:"avatar,omitempty"`
    TitlePhoto              string `json:"titlePhoto,omitempty" bson:"titlePhoto,omitempty"`
}

func OpenConnectionWithMongoDB(Collection **mongo.Collection){
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
    // defer client.Disconnect(context.Background())

    // Access the collection in the database
    *Collection = client.Database(dbName).Collection(collectionName)
}

func StoreRecentActionsInTheDatabase(Collection *mongo.Collection ,recentActions []RecentActions, resp *http.Response,maxtime int){
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

	recentActions= response.Result

	// Convert recentActions to interface{} slice
	var documents []interface{}
	for _, action := range recentActions {
        if (action.TimeSeconds >maxtime){
		    documents = append(documents, action)
        }
	}

	// Insert the data into MongoDB
	_, err = Collection.InsertMany(context.Background(), documents)
	if err != nil {
		fmt.Println("Error inserting data into MongoDB:", err)
		return
	}

    fmt.Println("Data inserted into MongoDB successfully!")
}

func QueryRecentActions(Collection *mongo.Collection) ([]RecentActions, error){
    // Find all documents in the collection
    cursor, err := Collection.Find(context.TODO(), bson.D{})
    if err != nil {
        fmt.Println(err)
    }
    var results []RecentActions

    // Iterate over the cursor and append each document to the array
    defer cursor.Close(context.TODO())
    for cursor.Next(context.TODO()) {
        var result RecentActions

        // Directly decode the BSON document into the RecentActions struct
        if err := cursor.Decode(&result); err != nil {
            fmt.Println(err)
            continue
        }

        // Append the decoded document to the results array
        results = append(results, result)
    }

    // Now you have all RecentActions objects in the results array
    
    // You can iterate over the array and process the objects as needed

    return results,err
}

func GetMaxTimeStamp(Collection *mongo.Collection) (int) {
    // Construct the aggregation pipeline
    pipeline := mongo.Pipeline{
        bson.D{
            {"$group", bson.D{
                {"_id", nil},
                {"maxTimeSeconds", bson.D{
                    {"$max", "$timeSeconds"},
                }},
            }},
        },
    }
    // Create a MongoDB context with a timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Perform the aggregation query
    cursor, err := Collection.Aggregate(ctx, pipeline)
    if err != nil {
        return 0
    }
    defer cursor.Close(ctx)

    // Iterate over the aggregation result
    var result struct {
        MaxTimeSeconds int `bson:"maxTimeSeconds"`
    }
    if cursor.Next(ctx) {
        if err := cursor.Decode(&result); err != nil {
            return 0
        }
    }

    return result.MaxTimeSeconds
}
func cfapi(resp **http.Response){
    url := "https://codeforces.com/api/recentActions?maxCount=30"
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    // Make the HTTP request
    client := http.Client{}
    resp1, err := client.Do(req)
    if err != nil {
        fmt.Println("Error making request:", err)
        return
    }
    *resp=resp1
}

func PerformWork(Collection *mongo.Collection){
    for{
        var resp *http.Response
        cfapi(&resp)
        maxtime:=GetMaxTimeStamp(Collection)
        var recentActions []RecentActions
        StoreRecentActionsInTheDatabase(Collection ,recentActions, resp, maxtime)
        // defer resp.Body.Close()
        time.Sleep(5*time.Minute)
    }
}


func main() {
    var Collection *mongo.Collection
    OpenConnectionWithMongoDB(&Collection)
    PerformWork(Collection)
    
    // QueryRecentActions(Collection)    
}