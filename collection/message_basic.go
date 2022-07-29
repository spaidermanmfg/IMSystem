package collection

type MessageBasic struct {
	Identity      string `bson:identity`
	User_Identity string `bson:user_identity`
	Room_Identity string `bson:room_identity`
	Data          string `bson:data`
	Created_At    int32  `bson:created_at`
	Updated_At    int32  `bson:updated_at`
}

func (MessageBasic) CollectionName() string {
	return "message_basic"
}
