package butler

type Butler struct {
	Status		int
	RoomId		string
	RoomName	string
	PlayerId	string		//唯一Id
}

const(
	IN_ROOM = iota
	OUT_ROOM
)

func NewButler(playerId string) *Butler {
	return &Butler{
		PlayerId:playerId,
		Status:OUT_ROOM,
	}
}

func(b *Butler) EnterRoom(roomId , roomName string) {
	b.RoomId = roomId
	b.RoomName = roomName
	b.Status = IN_ROOM
}

func(b *Butler) LeaveRoom() {
	b.RoomId = ""
	b.RoomName = ""
	b.Status = OUT_ROOM
}