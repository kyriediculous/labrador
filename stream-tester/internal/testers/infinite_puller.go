package testers

import (
	"fmt"

	"github.com/livepeer/stream-tester/internal/messenger"
	"github.com/livepeer/stream-tester/internal/model"
)

type infinitePuller struct {
	url  string
	save bool
}

// NewInfinitePuller ...
func NewInfinitePuller(url string, save bool) model.InfinitePuller {
	return &infinitePuller{
		url:  url,
		save: save,
	}
}

func (ip *infinitePuller) Start() {
	/*
		file, err := avutil.Open("media_w1620423594_b1487091_59.ts-00059.ts")
		fmt.Printf("file type %T\n", file)
		// hd := file.(*avutil.HandlerDemuxer)
		// fmt.Printf("%T\n", hd.Demuxer)
		if err != nil {
			panic(err)
		}
	*/

	done := make(chan struct{})
	// var sentTimesMap *utils.SyncedTimesMap
	down := newM3UTester(done, nil, true, true, ip.save)
	// go findSkippedSegmentsNumber(up, down)
	// sr.downloaders = append(sr.downloaders, down)
	msg := fmt.Sprintf("Starting to pull infinite stream from %s", ip.url)
	messenger.SendMessage(msg)
	down.Start(ip.url)

}
