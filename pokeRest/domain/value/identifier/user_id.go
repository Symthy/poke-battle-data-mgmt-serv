package identifier

type UserId struct {
	ValueId
}

func NewUserId(value uint) UserId {
	return UserId{ValueId{value}}
}

func NewEmptyUserId() UserId {
	return UserId{ValueId{0}}
}
