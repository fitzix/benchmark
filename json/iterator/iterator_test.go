package iterator

import (
	"testing"

	data "github.com/fitzix/benchmark/json"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func BenchmarkDecodeSmall(b *testing.B) {
	b.ReportAllocs()
	var d data.SmallPayload
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(data.SmallFixture, &d)
	}
}

func BenchmarkEncodeSmall(b *testing.B) {
	var d data.SmallPayload
	_ = json.Unmarshal(data.SmallFixture, &d)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(d)
	}
}

func BenchmarkStdDecodeMedium(b *testing.B) {
	b.ReportAllocs()
	var d data.MediumPayload
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(data.MediumFixture, &d)
	}
}

func BenchmarkEncodeMedium(b *testing.B) {
	var d data.MediumPayload
	_ = json.Unmarshal(data.MediumFixture, &d)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(d)
	}
}

func BenchmarkDecodeLarge(b *testing.B) {
	b.ReportAllocs()
	var d data.LargePayload
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(data.LargeFixture, &d)
	}
}

func BenchmarkEncodeLarge(b *testing.B) {
	var d data.LargePayload
	_ = json.Unmarshal(data.LargeFixture, &d)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(d)
	}
}
