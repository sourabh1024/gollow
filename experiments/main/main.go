package main

import (
	"github.com/golang/protobuf/proto"
	"gollow/experiments"
	"gollow/logging"
)

func main() {

	t := experiments.TestDto{}
	bag := t.LoadAll()

	bytes, err := proto.Marshal(bag)

	if err != nil {
		logging.GetLogger().Error("err in marshal , err : %+v", err)
	}

	newbag := experiments.TestBag{}
	err = proto.Unmarshal(bytes, &newbag)

	if err != nil {
		logging.GetLogger().Error("err in unmarshal , err : %+v", err)
	}

	for _, item := range newbag.Entries {
		logging.GetLogger().Info("id : %d", item.Id)
	}
}
