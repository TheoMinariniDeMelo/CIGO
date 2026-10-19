package protocol


type RequestType int

const (
	Register RequestType = iota
	Test
	TestSuccess
	TestFail
	Ping
	Pong
)

var typeName = map[RequestType]string{
    Register: "register",
    Test: "test",
	TestSuccess: "test success",
	TestFail: "test fail",
	Ping:     "ping",
    Pong:  "pong",
}

func (rt RequestType) String() string {
    return typeName[rt]
}

type Message struct {
	Type string `json: "type"`
	Message string `json: "message"`
	Commit string `json: "commit"`
}
