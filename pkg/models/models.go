// package models
// // RecentActions represents recent actions data model.
// type RecentActions struct {
//     TimeSeconds int       `json:"timeSeconds" bson:"timeSeconds"`
//     BlogEntry   BlogEntry `json:"blogEntry" bson:"blogEntry"`
//     Comment     Comment   `json:"comment" bson:"comment"`
// }

// // BlogEntry represents blog entry data model.
// type BlogEntry struct {
//     ID                     int      `json:"id" bson:"_id"`
//     OriginalLocale         string   `json:"originalLocale" bson:"originalLocale"`
//     CreationTimeSeconds    int      `json:"creationTimeSeconds" bson:"creationTimeSeconds"`
//     AuthorHandle           string   `json:"authorHandle" bson:"authorHandle"`
//     Title                  string   `json:"title" bson:"title"`
//     Content                string   `json:"content" bson:"content"`
//     Locale                 string   `json:"locale" bson:"locale"`
//     ModificationTimeSeconds int      `json:"modificationTimeSeconds" bson:"modificationTimeSeconds"`
//     AllowViewHistory       bool     `json:"allowViewHistory" bson:"allowViewHistory"`
//     Tags                   []string `json:"tags" bson:"tags"`
//     Rating                 int      `json:"rating" bson:"rating"`
// }

// // Comment represents comment data model.
// type Comment struct {
//     ID                  int    `json:"id" bson:"_id"`
//     CreationTimeSeconds int    `json:"creationTimeSeconds" bson:"creationTimeSeconds"`
//     CommentatorHandle   string `json:"commentatorHandle" bson:"commentatorHandle"`
//     Locale              string `json:"locale" bson:"locale"`
//     Text                string `json:"text" bson:"text"`
//     ParentCommentID     int    `json:"parentCommentId,omitempty" bson:"parentCommentId,omitempty"`
//     Rating              int    `json:"rating" bson:"rating"`
// }

// type RecentActionsResponse struct {
//     Status string          `json:"status"`
//     Result []RecentActions `json:"result"`
// }

// // User represents user data model.
// type User struct {
//     Handle                  string `json:"handle" bson:"_id"`
//     Email                   string `json:"email,omitempty" bson:"email,omitempty"`
//     VkID                    string `json:"vkId,omitempty" bson:"vkId,omitempty"`
//     OpenID                  string `json:"openId,omitempty" bson:"openId,omitempty"`
//     FirstName               string `json:"firstName,omitempty" bson:"firstName,omitempty"`
//     LastName                string `json:"lastName,omitempty" bson:"lastName,omitempty"`
//     Country                 string `json:"country,omitempty" bson:"country,omitempty"`
//     City                    string `json:"city,omitempty" bson:"city,omitempty"`
//     Organization            string `json:"organization,omitempty" bson:"organization,omitempty"`
//     Contribution            int    `json:"contribution,omitempty" bson:"contribution,omitempty"`
//     Rank                    string `json:"rank,omitempty" bson:"rank,omitempty"`
//     Rating                  int    `json:"rating,omitempty" bson:"rating,omitempty"`
//     MaxRank                 string `json:"maxRank,omitempty" bson:"maxRank,omitempty"`
//     MaxRating               int    `json:"maxRating,omitempty" bson:"maxRating,omitempty"`
//     LastOnlineTimeSeconds   int    `json:"lastOnlineTimeSeconds,omitempty" bson:"lastOnlineTimeSeconds,omitempty"`
//     RegistrationTimeSeconds int    `json:"registrationTimeSeconds,omitempty" bson:"registrationTimeSeconds,omitempty"`
//     FriendOfCount           int    `json:"friendOfCount,omitempty" bson:"friendOfCount,omitempty"`
//     Avatar                  string `json:"avatar,omitempty" bson:"avatar,omitempty"`
//     TitlePhoto              string `json:"titlePhoto,omitempty" bson:"titlePhoto,omitempty"`
// }