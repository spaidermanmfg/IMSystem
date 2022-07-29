package collection

type UserRoom struct {
	User_Identity    string `bson:user_identity`
	Room_Identity    string `bson:room_identity`
	Message_Identity string `bson:message_identity`
	Created_At       int32  `bson:created_at`
	Updated_At       int32  `bson:updated_at`
}

func (UserRoom) CollectionName() string {
	return "user_room"
}
