package collection

type RoomBasic struct {
	Identity      string `bson:identity`
	Number        string `bson:number`
	Name          string `bson:name`
	Info          string `bson:info`
	User_Identity string `bson:info`
	Created_At    int32  `bson:created_at`
	Updated_At    int32  `bson:updated_at`
}

func (RoomBasic) CollectionName() string {
	return "room_basic"
}
