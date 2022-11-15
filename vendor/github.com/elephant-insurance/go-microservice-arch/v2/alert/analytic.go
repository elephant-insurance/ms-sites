package alert

import "github.com/elephant-insurance/go-microservice-arch/v2/uf"

type AnalyticFunction func() []*uf.Event

var analytics = []AnalyticFunction{}

func AddAnalytic(f AnalyticFunction) {
	analytics = append(analytics, f)
}

func ClearAnalytics() {
	analytics = []AnalyticFunction{}
}

func RunAllAnalytics() []*uf.Event {
	rtn := []*uf.Event{}
	for i := 0; i < len(analytics); i++ {
		theseEvents := analytics[i]()
		for j := 0; j < len(theseEvents); j++ {
			if theseEvents[j] != nil {
				rtn = append(rtn, theseEvents[j])
			}
		}
	}

	if len(rtn) > 0 {
		return rtn
	}
	return nil
}
