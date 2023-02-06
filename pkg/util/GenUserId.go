package util

import (
	"fmt"
	sf "github.com/bwmarrin/snowflake"
	"time"
)

var node *sf.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}
func GenID() int64 {
	return node.Generate().Int64()
}

func GetID() int64 {
	if err := Init("2021-12-03", 1); err != nil {
		fmt.Println("Init() failed, err = ", err)
	}
	id := GenID()
	return id
}
