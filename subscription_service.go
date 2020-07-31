package muelle

import (
	"fmt"
	"time"

	"github.com/TheSmallBoat/cabinet"
	sr "github.com/TheSmallBoat/carlo/streaming_rpc"
	"github.com/TheSmallBoat/marina"
)

func subscriptionHandler(ctx *sr.Context) {
	latest := time.Now()
	switch ctx.Headers[ActionHeader] {
	case ActionSubscribe:
		topic := ctx.Headers[TopicHeader]
		qos := ctx.Headers[QosHeader]
		provider := MuNode.StreamNode.Providers().FindProvider(ctx.Conn())
		if provider == nil {
			fmt.Printf("conn is not a provider,%s:%d", ctx.KadId.Host.String(), ctx.KadId.Port)
			ctx.WriteHeader(ResponseSubscribeHeader, ServiceSubscribe)
			ctx.WriteHeader(ResponseActionHeader, ActionSubscribe)
			ctx.WriteHeader(ResponseStatusHeader, ResponseFailure)
			ctx.WriteHeader(BodyTitleHeader, ServiceError)
			ctx.Write([]byte("this conn is not a twin's provider."))
			return
		}

		var tsp marina.TwinServiceProvider = newTwinProvider(provider)
		MuNode.subSrv.sw.PeerNodeSubscribe(&tsp, qos[0], []byte(topic))

		ctx.WriteHeader(ResponseSubscribeHeader, ServiceSubscribe)
		ctx.WriteHeader(ResponseActionHeader, ActionSubscribe)
		ctx.WriteHeader(ResponseStatusHeader, ResponseSuccess)
		ctx.WriteHeader(BodyTitleHeader, ProcessTime)
		ctx.Write([]byte(time.Since(latest).String()))
	case ActionUnSubscribe:
		topic := ctx.Headers[TopicHeader]
		qos := ctx.Headers[QosHeader]
		MuNode.subSrv.sw.PeerNodeUnSubscribe(ctx.KadId.Pub, qos[0], []byte(topic))
		ctx.WriteHeader(ResponseSubscribeHeader, ServiceSubscribe)
		ctx.WriteHeader(ResponseActionHeader, ActionUnSubscribe)
		ctx.WriteHeader(ResponseStatusHeader, ResponseSuccess)
		ctx.WriteHeader(BodyTitleHeader, ProcessTime)
		ctx.Write([]byte(time.Since(latest).String()))
	default:
		ctx.WriteHeader(ResponseSubscribeHeader, ServiceSubscribe)
		ctx.WriteHeader(ResponseStatusHeader, ResponseFailure)
		ctx.WriteHeader(BodyTitleHeader, ActionError)
		ctx.Write([]byte("not a subscription operation request."))
	}
}

type subscriptionService struct {
	sw     *marina.SubscribeWorker
	handle sr.Handler
}

func NewSubscriptionService(twp *marina.TwinsPool, tt *cabinet.TTree, h sr.Handler) *subscriptionService {
	return &subscriptionService{
		sw:     marina.NewSubscribeWorker(twp, tt),
		handle: h,
	}
}
