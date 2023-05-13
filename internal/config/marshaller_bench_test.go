package config

import (
	"module07/internal/convertor"
	"testing"
)

var benchConfig *Config = NewConfig()

func BenchmarkStructToMapReflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		convertor.StructToMap(benchConfig)
	}
}

func BenchmarkStructToMapTemplCodegen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchConfig.StructToMap()
	}
}
