package dtos

type DebugLoginDto struct {
	Token string `json:"token"`
}

func NewDebugLoginDto(token string) (*DebugLoginDto) {
	return &DebugLoginDto{
		Token: token,
	}
}