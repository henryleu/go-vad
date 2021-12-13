package vad

import (
	"strings"
	"testing"
)

func TestConfig_Validate(t *testing.T) {
	SpeechTimeout := defaultConfig
	SpeechTimeout.SpeechTimeout = 0
	SilenceTimeout := defaultConfig
	SilenceTimeout.SilenceTimeout = 0
	NoinputTimeout := defaultConfig
	NoinputTimeout.NoinputTimers = true
	NoinputTimeout.NoinputTimeout = 0
	NoinputTimers := defaultConfig
	NoinputTimers.NoinputTimers = false
	RecognitionTimeout := defaultConfig
	RecognitionTimeout.RecognitionTimers = true
	RecognitionTimeout.RecognitionTimeout = 0
	RecognitionTimers := defaultConfig
	RecognitionTimers.RecognitionTimers = false
	VADLevel := defaultConfig
	VADLevel.VADLevel = 5
	SampleRate := defaultConfig
	SampleRate.SampleRate = 44100
	BytesPerSample := defaultConfig
	BytesPerSample.BytesPerSample = 3
	FrameDuration := defaultConfig
	FrameDuration.FrameDuration = 40
	Multiple := defaultConfig
	Multiple.Multiple = true

	tests := []struct {
		name         string
		c            *Config
		wantErr      bool
		invalidField string
	}{
		{
			name:         "Config.Validate() - invalid SpeechTimeout",
			c:            &SpeechTimeout,
			wantErr:      true,
			invalidField: "SpeechTimeout",
		},
		{
			name:         "Config.Validate() - invalid SilenceTimeout",
			c:            &SilenceTimeout,
			wantErr:      true,
			invalidField: "SilenceTimeout",
		},
		{
			name:         "Config.Validate() - invalid NoinputTimeout",
			c:            &NoinputTimeout,
			wantErr:      true,
			invalidField: "NoinputTimeout",
		},
		{
			name:         "Config.Validate() - invalid NoinputTimers",
			c:            &NoinputTimers,
			wantErr:      false,
			invalidField: "",
		},
		{
			name:         "Config.Validate() - invalid RecognitionTimeout",
			c:            &RecognitionTimeout,
			wantErr:      true,
			invalidField: "RecognitionTimeout",
		},
		{
			name:         "Config.Validate() - invalid RecognitionTimers",
			c:            &RecognitionTimers,
			wantErr:      false,
			invalidField: "",
		},
		{
			name:         "Config.Validate() - invalid VADLevel",
			c:            &VADLevel,
			wantErr:      true,
			invalidField: "VADLevel",
		},
		{
			name:         "Config.Validate() - invalid SampleRate",
			c:            &SampleRate,
			wantErr:      true,
			invalidField: "SampleRate",
		},
		{
			name:         "Config.Validate() - invalid BytesPerSample",
			c:            &BytesPerSample,
			wantErr:      true,
			invalidField: "BytesPerSample",
		},
		{
			name:         "Config.Validate() - invalid FrameDuration",
			c:            &FrameDuration,
			wantErr:      true,
			invalidField: "FrameDuration",
		},
		{
			name:         "Config.Validate() - invalid Multiple",
			c:            &Multiple,
			wantErr:      false,
			invalidField: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.c.Validate()
			hasErr := err != nil
			if hasErr != tt.wantErr {
				t.Errorf("Config.Validate() error = %v, wantErr %v", err, tt.wantErr)
			} else if hasErr && tt.invalidField != "" {
				msg := err.Error()
				if !strings.Contains(msg, tt.invalidField) {
					t.Errorf("Config.Validate() error = %v, wantErr %v, invalidField %v", err, tt.wantErr, tt.invalidField)
				}
			}
		})
	}
}

func TestConfig_NewDetector(t *testing.T) {
	c := defaultConfig
	type fields struct {
		SpeechTimeout      int
		SilenceTimeout     int
		NoinputTimeout     int
		NoinputTimers      bool
		RecognitionTimeout int
		RecognitionTimers  bool
		VADLevel           Level
		SampleRate         int
		BytesPerSample     int
		FrameDuration      int
		Multiple           bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *Detector
	}{
		{
			name:   "Config.NewDetector - use default config",
			fields: fields(defaultConfig),
			want: &Detector{
				Config:              c,
				state:               StateInactivity,
				duration:            0,
				recognitionDuration: 0,
				speechStart:         0,
				noinputDuration:     0,
				noinputStart:        0,
				sampleCount:         0,
				vadSampleCount:      0,
				bytesPerMillisecond: 0,
				bytesPerFrame:       0,
				vad:                 nil,
				work:                true,
				Events:              make(chan *Event),
				cache:               make([]byte, 0, cacheCap),
				Clips:               make([]*Clip, 0, 1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				SpeechTimeout:      tt.fields.SpeechTimeout,
				SilenceTimeout:     tt.fields.SilenceTimeout,
				NoinputTimeout:     tt.fields.NoinputTimeout,
				NoinputTimers:      tt.fields.NoinputTimers,
				RecognitionTimeout: tt.fields.RecognitionTimeout,
				RecognitionTimers:  tt.fields.RecognitionTimers,
				VADLevel:           tt.fields.VADLevel,
				SampleRate:         tt.fields.SampleRate,
				BytesPerSample:     tt.fields.BytesPerSample,
				FrameDuration:      tt.fields.FrameDuration,
				Multiple:           tt.fields.Multiple,
			}

			got := c.NewDetector()
			if got.SpeechTimeout != tt.want.SpeechTimeout {
				t.Errorf("Config.NewDetector() = %+v, want %+v", got.SpeechTimeout, tt.want.SpeechTimeout)
			}
			if got.SilenceTimeout != tt.want.SilenceTimeout {
				t.Errorf("Config.NewDetector() = %+v, want %+v", got.SilenceTimeout, tt.want.SilenceTimeout)
			}
			if got.NoinputTimeout != tt.want.NoinputTimeout {
				t.Errorf("Config.NewDetector() = %+v, want %+v", got.NoinputTimeout, tt.want.NoinputTimeout)
			}
			if got.NoinputTimers != tt.want.NoinputTimers {
				t.Errorf("Config.NewDetector() = %+v, want %+v", got.NoinputTimers, tt.want.NoinputTimers)
			}
			if got.RecognitionTimeout != tt.want.RecognitionTimeout {
				t.Errorf("Config.NewDetector() = %+v, want %+v", got.RecognitionTimeout, tt.want.RecognitionTimeout)
			}
			if got.RecognitionTimers != tt.want.RecognitionTimers {
				t.Errorf("Config.NewDetector() = %+v, want %+v", got.RecognitionTimers, tt.want.RecognitionTimers)
			}
			if got.VADLevel != tt.want.VADLevel {
				t.Errorf("Config.NewDetector() = %+v, want %+v", got.VADLevel, tt.want.VADLevel)
			}
			if got.SampleRate != tt.want.SampleRate {
				t.Errorf("Config.NewDetector() = %+v, want %+v", got.SampleRate, tt.want.SampleRate)
			}
			if got.BytesPerSample != tt.want.BytesPerSample {
				t.Errorf("Config.NewDetector() = %+v, want %+v", got.BytesPerSample, tt.want.BytesPerSample)
			}
			if got.FrameDuration != tt.want.FrameDuration {
				t.Errorf("Config.NewDetector() = %+v, want %+v", got.FrameDuration, tt.want.FrameDuration)
			}
			if got.Multiple != tt.want.Multiple {
				t.Errorf("Config.NewDetector() = %+v, want %+v", got.Multiple, tt.want.Multiple)
			}
		})
	}
}
