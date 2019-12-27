package main

import (
	"fmt"
	"io"
	"sync"

	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/actor/middleware/opentracing"
	"github.com/AsynkronIT/protoactor-go/router"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

func main() {
	runIterations()
}
func runIterations() {
	// InitExchange()
	jaegerCloser := initJaeger()
	defer jaegerCloser.Close()
	rootContext := actor.NewRootContext(nil).WithSpawnMiddleware(opentracing.TracingMiddleware())
	ExchangeProps := actor.PropsFromProducer(func() actor.Actor {
		return InitExchange()
	})
	grp := rootContext.Spawn(router.NewBroadcastGroup())

	pid, err := rootContext.SpawnNamed(ExchangeProps, "Exchange")
	rootContext.Send(grp, &router.AddRoutee{PID: pid})
	if err != nil {
		panic(err)
	}

	TraderProps := actor.PropsFromProducer(func() actor.Actor {
		return NewPariticipant()
	})

	pid = rootContext.Spawn(TraderProps)

	rootContext.Send(grp, &router.AddRoutee{PID: pid})
	count := sync.WaitGroup{}
	ticker := actor.PropsFromFunc(func(context actor.Context) {
		switch context.Message().(type) {
		case TICK:
			context.Request(grp, &router.BroadcastMessage{Message: TICK{}})
		case DONE:
			count.Done()
		}
	})
	t := rootContext.Spawn(ticker)
	for rounds := 0; rounds < 1; rounds++ {
		count.Add(1)
		rootContext.Send(t, TICK{})
		count.Wait()
		fmt.Printf("Round %d done", rounds)

	}

	console.ReadLine()
	// t.Observe()
	// stock, order := t.Trade()
	// e.SubmitOrder(stock, order)

}
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
