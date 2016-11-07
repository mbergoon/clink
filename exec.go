package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/mbergoon/clink/icmpecho"
	"github.com/robfig/cron"
)

type Executor struct {
}

func NewExecutor() *Executor {
	return &Executor{}
}

func (e *Executor) Exec(m MonitorConfig) {

	c := cron.New()

	switch {
	case m.Stype == "PSCNMON":
		LogM(TraceLevel, "EXEC set to port scan mode")

	default:
		LogM(TraceLevel, "EXEC set to echo monitor mode")

		resQueue := make(chan icmpecho.EchoResult)
		resList := make([]icmpecho.EchoResult, 0)

		//Clink Report ▶ Hosts Scanned: name1:address1, name2:address2
		//Clink Report ▶ Requests: 145 Success: 140 Fail: 5
		//Clink Report ▶ Success Percent: 97% Failure Percent: 3%
		//Clink Report ▶ Min Response: 12.2ms Max Response: 33.3ms
		//Clink Report ▶ name1:address1 > Requests: 100 Success: 90 Fail: 10
		//Clink Report ▶ name2:address2 > Requests: 100 Success: 90 Fail: 10

		report := func() {
			var hosts string

			for i, p := range m.Probes {
				hosts += p.Name + ":" + p.Host
				if i != len(m.Probes)-1 {
					hosts += ", "
				}
			}

			reqs, succ, fail := 0, 0, 0
			succp, failp := 0.0, 0.0
			var min, max time.Duration

			reqs = len(resList)
			for i, v := range resList {
				if i == 0 {
					min = v.Elapsed
					max = v.Elapsed
				} else {
					if min > v.Elapsed {
						min = v.Elapsed
					}
					if max < v.Elapsed {
						max = v.Elapsed
					}
				}
				if v.Status {
					succ += 1
				} else {
					fail += 1
				}

			}

			if reqs > 0 {
				succp = float64(succ) / float64(reqs)
				failp = float64(fail) / float64(reqs)
			}

			rpt := ColorClear(COLOR_B_BLACK, "\n")
			rpt += ColorClear(COLOR_B_BLACK, " Clink Report ▶ Hosts Scanned: ") + hosts + "\n"
			rpt += ColorClear(COLOR_B_BLACK, " Clink Report ▶ Requests: ") + fmt.Sprint(reqs) + ColorClear(COLOR_B_BLACK, " Success: ") + fmt.Sprint(succ) + ColorClear(COLOR_B_BLACK, " Fail: ") + fmt.Sprint(fail) + "\n"
			rpt += ColorClear(COLOR_B_BLACK, " Clink Report ▶ Success Percent: ") + fmt.Sprint(succp) + "%" + ColorClear(COLOR_B_BLACK, " Fail Percent: ") + fmt.Sprint(failp) + "%" + "\n"
			rpt += ColorClear(COLOR_B_BLACK, " Clink Report ▶ Min: ") + fmt.Sprint(min) + ColorClear(COLOR_B_BLACK, " Max: ") + fmt.Sprint(max) + "\n"

			mReqs := make(map[string]int)
			mSucc := make(map[string]int)
			mFail := make(map[string]int)

			for _, v := range resList {
				if _, ok := mReqs[v.Address]; ok {
					mReqs[v.Address] += 1
				} else {
					mReqs[v.Address] = 0
				}

				if _, ok := mSucc[v.Address]; ok && v.Status {
					mSucc[v.Address] += 1
				} else {
					mSucc[v.Address] = 0
				}

				if _, ok := mFail[v.Address]; ok && !v.Status {
					mFail[v.Address] += 1
				} else {
					mFail[v.Address] = 0
				}
			}

			for _, pl := range m.Probes {
				rpt += ColorClear(COLOR_B_BLACK, " Clink Report ▶ "+pl.Name+":"+pl.Host) + ColorClear(COLOR_B_BLACK, " Requests: ") + fmt.Sprint(mReqs[pl.Host]) + ColorClear(COLOR_B_BLACK, " Success: ") + fmt.Sprint(mSucc[pl.Host]) + ColorClear(COLOR_B_BLACK, " Fail: ") + fmt.Sprint(mFail[pl.Host]) + "\n"
			}

			rpt += ColorClear(COLOR_B_BLACK, "\n")

			fmt.Print(rpt)
		}

		//Handle SIGINT
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		go func() {
			for _ = range sigint {
				c.Stop()
				report()
				os.Exit(0)
			}
		}()

		c.AddFunc("@every 10s", report)

		for _, probe := range m.Probes {

			LogM(TraceLevel, fmt.Sprint(probe))

			interval := m.Interval / 1000
			cronExp := "*/" + fmt.Sprint(interval) + " * * * * *"

			pn := probe.Name
			p := probe.Host

			c.AddFunc(cronExp, func() {
				res := icmpecho.Echo(pn, p, m.Timeout)
				resQueue <- res
			})

			c.Start()

		}

		for elem := range resQueue {
			resList = append(resList, elem)
			if elem.Status {
				s := ColorClear(COLOR_B_BLACK, "Response ▶ "+elem.Name+":"+elem.Address+" ▶ SUCCESS at "+fmt.Sprint(elem.Start)+" taking "+fmt.Sprint(elem.Elapsed)+" S: "+fmt.Sprint(elem.Sent)+"b R: "+fmt.Sprint(elem.Received)+"b ")
				LogM(TraceLevel, s)
				fmt.Println(s)
			} else {
				s := ColorClear(COLOR_B_RED, "Response ▶ "+elem.Name+":"+elem.Address+" ▶ FAIL at "+fmt.Sprint(elem.Start)+" taking "+fmt.Sprint(elem.Elapsed)+" S: "+fmt.Sprint(elem.Sent)+"b R: "+fmt.Sprint(elem.Received)+"b "+fmt.Sprint(elem.Err))
				LogM(TraceLevel, s)
				fmt.Println(s)
			}
		}
	}
}
