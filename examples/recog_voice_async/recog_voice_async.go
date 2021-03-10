package main

import (
	"io"
	"log"

	vad "github.com/henryleu/go-vad"
	examples "github.com/henryleu/go-vad/examples"
	wav "github.com/henryleu/go-wav"
)

func main() {
	fn := "../data/8ef79f2695c811ea.wav"

	r, err := wav.NewReaderFromFile(fn)
	if err != nil {
		log.Fatalf("wav.NewReader() error = %v", err)
	}
	examples.InitSpeaker(int(r.FmtChunk.Data.SamplesPerSec), 100)

	c := vad.NewDefaultConfig()
	c.SampleRate = int(r.FmtChunk.Data.SamplesPerSec)
	c.BytesPerSample = int(r.FmtChunk.Data.BitsPerSamples / 8)
	// 设置一下参数效果最佳
	c.SilenceTimeout = 800
	c.SpeechTimeout = 800
	c.NoinputTimeout = 20000
	c.VADLevel = 3

	err = c.Validate()
	if err != nil {
		log.Fatalf("Config.Validate() error = %v", err)
	}
	d := c.NewDetector()
	err = d.Init()
	if err != nil {
		log.Fatalf("Detector.Init() error = %v", err)
	}

	frame := make([]byte, d.BytesPerFrame())
	done := make(chan bool)
	go func() {
		for e := range d.Events {
			switch e.Type {
			case vad.EventVoiceBegin:
				log.Println("voice begin")
				break
			case vad.EventVoiceEnd:
				log.Println("voice end")
				f, err := examples.NewFile()
				e.Clip.SaveToWriter(f)
				wn := f.Name()
				rf, err := examples.OpenFile(wn)
				if err != nil {
					log.Fatalf("fs.Open() error = %v", err)
				}
				examples.PlayWaveFile(rf)
				done <- true
				break
			case vad.EventNoinput:
				log.Println("no input")
				f, err := examples.NewFile()
				e.Clip.SaveToWriter(f)
				wn := f.Name()
				rf, err := examples.OpenFile(wn)
				if err != nil {
					log.Fatalf("fs.Open() error = %v", err)
				}
				examples.PlayWaveFile(rf)
				done <- true
				break
			default:
				log.Printf("illegal event type %v\n", e.Type)
				done <- true
			}
		}
	}()

	for {
		_, err := io.ReadFull(r, frame)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			log.Println("file is EOF")
			d.Finalize()
			break
		} else if err != nil {
			log.Fatalf("io.ReadFull() error = %v", err)
		}
		err = d.Process(frame)
		if err != nil {
			log.Fatalf("Detector.Process() error = %v", err)
		}
		if !d.Working() {
			log.Println("detector is stopped")
			break
		}
	}
	<-done
}
