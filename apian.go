package apian

import (
	"time"

	"github.com/gotomicro/ego/core/elog"

	"github.com/kl7sn/apian/module/forward/webhook"
	"github.com/kl7sn/apian/module/monitor"
	"github.com/kl7sn/apian/pkg/dto"
)

type Apian struct {
	opts           *options
	memAvg         uint64
	memCoolingTime time.Time

	Forwarder dto.Forwarder
}

// New is the entry of apian
// 1. start monitor
// 2. storage data
// 3. calculate data
// 4. generate pprof files
// 5. upload to oss
func New(opts ...Option) (*Apian, error) {
	apian := &Apian{
		opts: newOptions(),
	}
	for _, o := range opts {
		if err := o.apply(apian.opts); err != nil {
			return nil, err
		}
	}
	return apian, nil
}

// EnableMem enables the mem dump.
func (a *Apian) EnableMem() *Apian {
	a.opts.memOpts.Enable = true
	return a
}

func (a *Apian) Start() error {
	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				a.calculateData()
			}
		}
	}()
	return nil
}

func (a *Apian) calculateData() {
	current, usedPercent := monitor.ReadMemStats()

	if a.memAvg == 0 {
		a.memAvg = current
		return
	}
	diff := (float64(current) - float64(a.memAvg)) * 100 / float64(a.memAvg)

	elog.Info("calculate", elog.Int("avgMiB", int(a.memAvg)), elog.Int("currentMiB", int(current)), elog.Int("diffPercent", int(diff)), elog.Int("usedPercent", int(usedPercent)))

	if uint64(usedPercent) >= a.opts.memOpts.TriggerPercent &&
		uint64(diff) >= a.opts.memOpts.TriggerDiff {
		a.pprof()
	}
	a.memAvg = (a.memAvg + current) / 2
}

func (a *Apian) pprof() {
	if a.memCoolingTime.After(time.Now()) {
		elog.Info("cooling time")
		return
	}
	a.memCoolingTime = time.Now().Add(a.opts.memOpts.CoolingTime)
	webhook.Webhook(a.Forwarder.WebHook)
}
