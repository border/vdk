package format

import (
	"github.com/border/vdk/av/avutil"
	"github.com/border/vdk/format/aac"
	"github.com/border/vdk/format/flv"
	"github.com/border/vdk/format/mp4"
	"github.com/border/vdk/format/rtmp"
	"github.com/border/vdk/format/rtsp"
	"github.com/border/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
