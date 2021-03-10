package main

import (
	"io"
	"log"
	"time"

	vad "github.com/henryleu/go-vad"
	examples "github.com/henryleu/go-vad/examples"
	wav "github.com/henryleu/go-wav"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in f", r)
		}
	}()

	fn := "../data/8ef79f2695c811ea.wav"
	// fn := "../data/tts-01.wav"

	r, err := wav.NewReaderFromFile(fn)
	if err != nil {
		log.Fatalf("wav.NewReader() error = %v", err)
	}

	c := vad.NewDefaultConfig()
	c.SampleRate = int(r.FmtChunk.Data.SamplesPerSec)
	c.BytesPerSample = int(r.FmtChunk.Data.BitsPerSamples / 8)
	// 设置一下参数效果最佳
	c.SilenceTimeout = 800
	c.SpeechTimeout = 800
	c.VADLevel = 3
	log.Printf("vad level: %v\n", c.VADLevel)
	c.Multiple = true
	err = c.Validate()
	if err != nil {
		log.Fatalf("Config.Validate() error = %v", err)
	}
	d := c.NewDetector()
	err = d.Init()
	if err != nil {
		log.Fatalf("Detector.Init() error = %v", err)
	}

	// 使用语音探测器切割一个音频文件
	frame := make([]byte, d.BytesPerFrame())
	for {
		_, err := io.ReadFull(r, frame)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			log.Println(err)
			d.Finalize()
			break
		}
		if err != nil {
			log.Fatalf("io.ReadFull() error = %v", err)
		}
		err = d.Process(frame)
		if err != nil {
			log.Fatalf("Detector.Process() error = %v", err)
		}
	}

	log.Printf("clip number: %v\n", len(d.Clips))
	examples.InitSpeaker(c.SampleRate, 100)

	// play all the speech clips detected
	// 播放“分段录音”（已识别的讲话片段）
	for i, c := range d.Clips {
		f, err := examples.NewFile()
		wn := f.Name()
		log.Printf("clip %v : %v\n", i+1, wn)
		c.SaveToWriter(f)
		log.Println()
		rf, err := examples.OpenFile(wn)
		if err != nil {
			log.Fatalf("fs.Open() error = %v", err)
		}
		examples.PlayWaveFile(rf)
		time.Sleep(time.Millisecond * 100)
	}

	// play the total record of all the processed voice data
	// 播放“全程录音”（所有处理过的音频数据的）
	tc := d.GetTotalClip()
	f, err := examples.NewFile()
	wn := f.Name()
	log.Printf("original clip: %v\n", wn)
	tc.SaveToWriter(f)
	// rf, err := examples.OpenFile(wn)
	// if err != nil {
	// 	log.Fatalf("fs.Open() error = %v", err)
	// }
	// examples.PlayWaveFile(rf)
}
