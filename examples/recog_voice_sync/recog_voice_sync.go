package main

import (
	"fmt"
	"io"
	"log"

	vad "github.com/henryleu/go-vad"
	examples "github.com/henryleu/go-vad/examples"
	wav "github.com/henryleu/go-wav"
)

func main() {
	// fn := "../data/8ef79f2695c811ea.wav"
	// fn := "../data/tts-01.wav"
	fn := "../data/haichao_test_01.wav"

	r, err := wav.NewReaderFromFile(fn)
	if err != nil {
		log.Fatalf("wav.NewReader() error = %v", err)
	}
	log.Println("SamplesPerSec", r.FmtChunk.Data.SamplesPerSec)
	examples.InitSpeaker(int(r.FmtChunk.Data.SamplesPerSec), 100)

	c := vad.NewDefaultConfig()
	c.SampleRate = int(r.FmtChunk.Data.SamplesPerSec)
	c.BytesPerSample = int(r.FmtChunk.Data.BitsPerSamples / 8)
	log.Println("SampleRate", c.SampleRate)
	log.Println("BytesPerSample", c.BytesPerSample)
	// 设置一下参数效果最佳
	c.SilenceTimeout = 400
	c.SpeechTimeout = 500
	c.NoinputTimeout = 20000
	c.RecognitionTimeout = 10000
	c.VADLevel = 1

	err = c.Validate()
	if err != nil {
		log.Fatalf("Config.Validate() error = %v", err)
	}
	d := c.NewDetector()
	err = d.Init()
	if err != nil {
		log.Fatalf("Detector.Init() error = %v", err)
	}
	log.Printf("frame length %v\n", d.BytesPerFrame())
	frame := make([]byte, d.BytesPerFrame())
	for {
		_, err := io.ReadFull(r, frame)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			log.Println("file is EOF")
			d.Finalize()
			break
		}
		if err != nil {
			log.Fatalf("io.ReadFull() error = %v", err)
		}
		err = d.Process(frame)
		if err != nil {
			log.Fatalf("Detector.Process() error = %v", err)
			d.Finalize()
			break
		}
		if !d.Working() {
			log.Println("detector is stopped")
			d.Finalize()
			break
		}
	}

events_loop:
	for e := range d.Events {
		switch e.Type {
		case vad.EventVoiceBegin:
			log.Println("voice begin")
		case vad.EventVoiceEnd:
			log.Println("voice end")
			f, err := examples.NewFile()
			e.Clip.SaveToWriter(f)
			e.Clip.PrintDetail()
			log.Println(len(e.Clip.Data))
			log.Println("clip digest", e.Clip.GenerateDigest())

			tc := d.GetTotalClip()
			tc.PrintDetail()
			log.Println("total clip degest", tc.GenerateDigest())

			wn := f.Name()
			rf, err := examples.OpenFile(wn)
			if err != nil {
				log.Fatalf("fs.Open() error = %v", err)
			}
			examples.PlayWaveFile(rf)
			break events_loop
		case vad.EventNoinput:
			fmt.Println("no input")
			f, err := examples.NewFile()
			e.Clip.SaveToWriter(f)
			wn := f.Name()
			rf, err := examples.OpenFile(wn)
			if err != nil {
				log.Fatalf("fs.Open() error = %v", err)
			}
			examples.PlayWaveFile(rf)
			break events_loop
		default:
			log.Printf("illegal event type %v\n", e.Type)
		}
	}
}
