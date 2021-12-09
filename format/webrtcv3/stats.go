package webrtc

import "time"

// StatsTimestamp is a timestamp represented by the floating point number of
// milliseconds since the epoch.
type StatsTimestamp float64

// VideoReceiverStats contains video metrics related to a specific receiver.
type VideoReceiverStats struct {
	// Timestamp is the timestamp associated with this object.
	Timestamp StatsTimestamp `json:"timestamp"`

	// FramesPerSecond represents the nominal FPS value before the degradation preference
	// is applied. It is the number of complete frames in the last second. For sending
	// tracks it is the current captured FPS and for the receiving tracks it is the
	// current decoding framerate.
	FramesPerSecond float64 `json:"framesPerSecond"`

	// FramesReceived Represents the total number of complete frames received for
	// this receiver. This metric is incremented when the complete frame is received.
	FramesReceived uint32 `json:"framesReceived"`

	FramesDelay uint32 `json:"framesDelay"`

	VideoDelay uint32 `json:"videoDelay"`

	FramesDelayTotal uint32 `json:"framesDelayTotal"`

	FramesDelayAvg uint32 `json:"framesDelayAvg"`
}

// Time returns the time.Time represented by this timestamp.
func (s StatsTimestamp) Time() time.Time {
	millis := float64(s)
	nanos := int64(millis * float64(time.Millisecond))

	return time.Unix(0, nanos).UTC()
}

func statsTimestampFrom(t time.Time) StatsTimestamp {
	return StatsTimestamp(t.UnixNano() / int64(time.Millisecond))
}

func statsTimestampNow() StatsTimestamp {
	return statsTimestampFrom(time.Now())
}
