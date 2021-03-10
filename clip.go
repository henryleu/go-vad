package vad

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	wav "github.com/henryleu/go-wav"
)

// Clip defines voice clip for processing and persisting
type Clip struct {
	// SampleRate defines the number of samples per second, aka. sample rate.
	SampleRate int

	// BytesPerSample defines bytes per sample (sample depth) for linear pcm
	BytesPerSample int

	// Time defines the starting time of the voice clip in the whole voice
	// stream in milliseconds.
	Start int

	// Duration defines the time span of the voice clip in milliseconds.
	Duration time.Duration

	// Data is the chunk data of the voice clip as the specific sample rate and depth
	Data []byte
}

// SaveToFile creates a file and write the clip to it
func (c *Clip) SaveToFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		log.Printf("Fail to save clip to %v, error: %v", path, err)
		return err
	}
	return c.SaveToWriter(f)
}

// SaveToWriter creates a file and write the clip to a wave file.
func (c *Clip) SaveToWriter(wc io.WriteCloser) error {
	param := wav.WriterParam{
		Out:           wc,
		Channel:       int(1),
		SampleRate:    int(c.SampleRate),
		BitsPerSample: int(c.BytesPerSample * 8),
	}

	w, err := wav.NewWriter(param)
	defer w.Close()
	if err != nil {
		log.Printf("Fail to create a new wave clip writer, error: %v", err)
		return err
	}

	_, err = w.Write(c.Data)
	if err != nil {
		log.Printf("Fail to write clip data, error: %v", err)
		return err
	}

	return nil
}

// GenerateDigest generates a sha256 hash of voice data as hex string
func (c *Clip) GenerateDigest() string {
	h := sha256.New()
	h.Write(c.Data)
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

// PrintDetail prints detailed properties of the clip
func (c *Clip) PrintDetail() {
	log.Printf("Clip SampleRate:\t%v\n", c.SampleRate)
	log.Printf("Clip BytesPerSample:\t%v\n", c.BytesPerSample)
	log.Printf("Clip Start:\t%v\n", c.Start)
	log.Printf("Clip Duration:\t%v\n", c.Duration)
	log.Printf("Clip Data Size:\t%v\n", len(c.Data))
}
