package main

import (
	"context"
	"fmt"

	"github.com/waku-org/go-waku/waku/v2/node"
	"github.com/waku-org/go-waku/waku/v2/protocol"
	"github.com/waku-org/go-waku/waku/v2/protocol/pb"
	"github.com/waku-org/go-waku/waku/v2/utils"
)

func InitWaku() {
	wakuNode, err := node.New(node.WithLightPush())
	CheckError("Error initializing new Waku node", err)

	err = wakuNode.Start(context.Background())
	CheckError("Error starting new Waku node", err)

	var value uint32 = 0
	ptr := &value
	contentTopic, err := protocol.NewContentTopic("basic2", "1", "test", "proto")
	CheckError("Error creating protocol NewContentTopic", err)

	msg := &pb.WakuMessage{
		Payload:      []byte("Hello World"),
		Version:      ptr,
		ContentTopic: contentTopic.String(),
		Timestamp:    utils.GetUnixEpoch(),
	}

	msgId, err := wakuNode.Lightpush().Publish(context.Background(), msg)
	CheckError("Error sending a message", err)

	println(msgId)
	fmt.Printf("%s", msgId)

}
