package vad

// EventType defines events of activity detector
type EventType int

const (
	// EventVoiceBegin means voice is detected at the time (transition state from inactivity to activity)
	EventVoiceBegin EventType = iota

	// EventVoiceEnd means voice is over at the time (transition state from activity to inactivity)
	EventVoiceEnd

	// EventNoinput means no input event occurred
	EventNoinput
)

// Event is emitted in the process of speech detection
type Event struct {

	// Type is the event type which can be emitted by the detector
	Type EventType

	// Clip is the detected and clipped audio during voice frame processing.
	// It is populated only on the EventVoiceEnd event, or nil
	Clip *Clip
}
