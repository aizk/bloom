package bloom

import "math"

type Bloom struct {
	Storage Storager
	Config *Config
}

type Config struct {
    ErrorRate float64
    Elements uint
    BitNumbers uint
    HashNumbers uint
}

func (l *Bloom) New(config *Config) *Bloom {
	if config.ErrorRate == 0 {
		config.ErrorRate = 0.001
	}
	if config.Elements == 0 {
		config.Elements = 1000
	}

	if config.BitNumbers == 0 || config.HashNumbers == 0 {
		config.BitNumbers, config.HashNumbers = l.Estimate(config.Elements, config.ErrorRate)
	}

	return &Bloom{
		Config: config,
	}
}

func (l *Bloom) UseStorage(s Storager) *Bloom {
	if s == nil {
		panic("storage is nil.")
	}
	l.Storage = s
	return l
}

func (l *Bloom) Exist(key string, value interface{}) (ex bool, err error) {
	// get
	return l.Storage.Exist(key, value)
}

// get need bit numbers and hash func numbers
func (l *Bloom) Estimate(n uint, p float64) (uint, uint) {
	m := math.Ceil(float64(n) * math.Log(p) / math.Log(1.0/math.Pow(2.0, math.Ln2)))
	k := math.Ln2*m/float64(n) + 0.5
	return uint(m), uint(k)
}