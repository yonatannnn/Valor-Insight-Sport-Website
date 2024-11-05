package domain

type Player struct {
	ID             string `bson:"_id" json:"_id"`
	Name           string `bson:"name" json:"name"`
	Age            int    `bson:"age" json:"age"`
	Position       string `bson:"position" json:"position"`
	TeamID         string `bson:"team_id" json:"team_id"` // reference to Teams collection
	Nationality    string `bson:"nationality" json:"nationality"`
	Stats          Stats  `bson:"stats" json:"stats"`
	Bio            string `bson:"bio" json:"bio"`
	ProfilePicture string `bson:"profile_picture" json:"profile_picture"` // URL or path to profile image
}

type Stats struct {
	Goals   int `bson:"goals" json:"goals"`
	Assists int `bson:"assists" json:"assists"`
}
