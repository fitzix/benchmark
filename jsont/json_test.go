package jsont

import (
	"encoding/json"
	"testing"
)

// var json = jsoniter.ConfigCompatibleWithStandardLibrary

func BenchmarkDecodeSmall(b *testing.B) {
	b.ReportAllocs()
	var d SmallPayload
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(SmallFixture, &d)
	}
}

func BenchmarkEncodeSmall(b *testing.B) {
	var d SmallPayload
	_ = json.Unmarshal(SmallFixture, &d)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(d)
	}
}

func BenchmarkDecodeMedium(b *testing.B) {
	b.ReportAllocs()
	var d MediumPayload
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(MediumFixture, &d)
	}
}

func BenchmarkEncodeMedium(b *testing.B) {
	var d MediumPayload
	_ = json.Unmarshal(MediumFixture, &d)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(d)
	}
}

func BenchmarkDecodeLarge(b *testing.B) {
	b.ReportAllocs()
	var d LargePayload
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(LargeFixture, &d)
	}
}

func BenchmarkEncodeLarge(b *testing.B) {
	var d LargePayload
	_ = json.Unmarshal(LargeFixture, &d)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(d)
	}
}
