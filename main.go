package main

import (
	"fmt"
	"io"
//	"sync"

//	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
	//"github.com/AsynkronIT/protoactor-go/router"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
  "marketplace/exchange"
  "marketplace/participant"
)

func main() {
  e:= exchange.InitExchange()
  participant.NewParticipant(e)
//	runIterations()
}
func getRootContext() *actor.RootContext {
	return actor.NewRootContext(nil)
	// return actor.NewRootContext(nil).WithSpawnMiddleware(opentracing.TracingMiddleware())
}
//func runIterations() {
	//Setting up context with middleware support
//	jaegerCloser := initJaeger()
//	defer jaegerCloser.Close()
//	rootContext := getRootContext()
	// Setting up Exchange
//	ExchangeProps := actor.PropsFromProducer(func() actor.Actor {
//		return exchange.InitExchangeActor()
//	})
//	_, err := rootContext.SpawnNamed(ExchangeProps, "Exchange")
//	if err != nil {
//		panic("Process Exchange already initialized")
//	}
	// Initialize MatketMaker and beging IPO

//	MMProps := actor.PropsFromProducer(func() actor.Actor {
//		return &MarketMakerTrader{}
//	})
//	_, err = rootContext.SpawnNamed(MMProps, "MarketMaker")
//	if err != nil {
//		panic("Process MarketMaker already initialized")
//	}

	//Setting broadcast group
//	grp := rootContext.Spawn(router.NewBroadcastGroup())
//	if err != nil {
//		panic(err)
//	}
	// Spawn Traders and add them one by one to broadcast group
//	TraderProps := actor.PropsFromProducer(func() actor.Actor {
//		return NewParticipant()
//	})
	//pid := rootContext.Spawn(TraderProps)

//	rootContext.Send(grp, &router.AddRoutee{PID: pid})
	// Setting up WaitGroup for syncrhonization.
	// Initalize ticker for clock
//	count := sync.WaitGroup{}
//	ticker := actor.PropsFromFunc(func(context actor.Context) {
	//	switch context.Message().(type) {
	//	case TICK:
	//		context.Request(grp, &router.BroadcastMessage{Message: TICK{}})
	//	case DONE:
	//		count.Done()
	//	}
	//})
//	t := rootContext.Spawn(ticker)
	// Begin Rounds. Wait for sync -- all traders repond, before begin next round
//	for rounds := 0; rounds < 5; rounds++ {
//		count.Add(1)
//		rootContext.Send(t, TICK{})
//		count.Wait()
//		fmt.Printf("Round %d done\n", rounds)

	//}

//	console.ReadLine()
	// t.Observe()
	// stock, order := t.Trade()
  // e.SubmitOrder(stock, order)

//}
func initJaeger() io.Closer {
	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	closer, err := cfg.InitGlobalTracer(
		"jaeger-test",
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		//log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		panic(fmt.Sprintf("Could not initialize jaeger tracer: %s", err.Error()))
	}
	return closer
}
