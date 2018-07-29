package experiments

import (
	"encoding/json"
	"github.com/streams/apis/pb"
	"gollow/logging"
)

type sourceDTO interface {
	ToPB()
}

type random interface {
	pb.Message
	GetPrimaryID() int64
	LoadAll() []random
	Marshal() ([]byte, error)
}

type Results struct {
	data []random
}

func DoSomething(val random) {

	res := val.LoadAll()
	logging.GetLogger().Info("res", len(res))
	logging.GetLogger().Info("res id", res[0].GetPrimaryID())

	//r := Results{data: res}
	bytes, err := json.Marshal(res)

	logging.GetLogger().Info("bytes", len(bytes))

	if err != nil {
		logging.GetLogger().Error("error in marshalloing  bro , %+v", err)
		return
	}

	newr := make([]random, 0)
	json.Unmarshal(bytes, &newr)

	logging.GetLogger().Info("", len(newr))

}
