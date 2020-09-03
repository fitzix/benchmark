# json std/jsonIterator benchmark

### std

```
goversion: 1.14.4
goos: darwin
goarch: amd64
pkg: github.com/fitzix/benchmark/json/std
BenchmarkDecodeSmall-4            378150              2890 ns/op             272 B/op          6 allocs/op
BenchmarkDecodeSmall-4            402830              2887 ns/op             272 B/op          6 allocs/op
BenchmarkDecodeSmall-4            405274              2892 ns/op             272 B/op          6 allocs/op
BenchmarkDecodeSmall-4            414421              2895 ns/op             272 B/op          6 allocs/op
BenchmarkEncodeSmall-4           2285047               515 ns/op             192 B/op          2 allocs/op
BenchmarkEncodeSmall-4           2308908               514 ns/op             192 B/op          2 allocs/op
BenchmarkEncodeSmall-4           2302418               519 ns/op             192 B/op          2 allocs/op
BenchmarkEncodeSmall-4           2345596               512 ns/op             192 B/op          2 allocs/op
BenchmarkStdDecodeMedium-4         49272             23508 ns/op             496 B/op         11 allocs/op
BenchmarkStdDecodeMedium-4         49128             23352 ns/op             496 B/op         11 allocs/op
BenchmarkStdDecodeMedium-4         50677             23692 ns/op             496 B/op         11 allocs/op
BenchmarkStdDecodeMedium-4         49669             29693 ns/op             496 B/op         11 allocs/op
BenchmarkEncodeMedium-4          1219962              1073 ns/op             224 B/op          2 allocs/op
BenchmarkEncodeMedium-4          1206912               946 ns/op             224 B/op          2 allocs/op
BenchmarkEncodeMedium-4          1276762              1016 ns/op             224 B/op          2 allocs/op
BenchmarkEncodeMedium-4          1000000              1169 ns/op             224 B/op          2 allocs/op
BenchmarkDecodeLarge-4              4393            287284 ns/op             344 B/op          7 allocs/op
BenchmarkDecodeLarge-4              4164            263470 ns/op             344 B/op          7 allocs/op
BenchmarkDecodeLarge-4              4620            257807 ns/op             344 B/op          7 allocs/op
BenchmarkDecodeLarge-4              4250            256071 ns/op             344 B/op          7 allocs/op
BenchmarkEncodeLarge-4           1116111              1096 ns/op             160 B/op          2 allocs/op
BenchmarkEncodeLarge-4           1030826              1216 ns/op             160 B/op          2 allocs/op
BenchmarkEncodeLarge-4            952455              1212 ns/op             160 B/op          2 allocs/op
BenchmarkEncodeLarge-4            996213              1303 ns/op             160 B/op          2 allocs/op
```

### jsonIterator
```
goversion: 1.14.4
goos: darwin
goarch: amd64
pkg: github.com/fitzix/benchmark/json/iterator
BenchmarkDecodeSmall-4           1435369               860 ns/op              64 B/op          2 allocs/op
BenchmarkDecodeSmall-4           1464547               900 ns/op              64 B/op          2 allocs/op
BenchmarkDecodeSmall-4           1510258               793 ns/op              64 B/op          2 allocs/op
BenchmarkDecodeSmall-4           1655820               729 ns/op              64 B/op          2 allocs/op
BenchmarkEncodeSmall-4           2179290               554 ns/op             192 B/op          2 allocs/op
BenchmarkEncodeSmall-4           2187823               548 ns/op             192 B/op          2 allocs/op
BenchmarkEncodeSmall-4           2137189               581 ns/op             192 B/op          2 allocs/op
BenchmarkEncodeSmall-4           2020676               733 ns/op             192 B/op          2 allocs/op
BenchmarkStdDecodeMedium-4        164416              7974 ns/op             384 B/op         41 allocs/op
BenchmarkStdDecodeMedium-4        129982             10463 ns/op             384 B/op         41 allocs/op
BenchmarkStdDecodeMedium-4        186229              6461 ns/op             384 B/op         41 allocs/op
BenchmarkStdDecodeMedium-4        190923              8741 ns/op             384 B/op         41 allocs/op
BenchmarkEncodeMedium-4          1000000              1131 ns/op             224 B/op          2 allocs/op
BenchmarkEncodeMedium-4          1000000              1036 ns/op             224 B/op          2 allocs/op
BenchmarkEncodeMedium-4          1000000              1100 ns/op             224 B/op          2 allocs/op
BenchmarkEncodeMedium-4          1000000              1069 ns/op             224 B/op          2 allocs/op
BenchmarkDecodeLarge-4             10000            116528 ns/op           12881 B/op       1124 allocs/op
BenchmarkDecodeLarge-4             10000            107951 ns/op           12881 B/op       1124 allocs/op
BenchmarkDecodeLarge-4             10000            101734 ns/op           12881 B/op       1124 allocs/op
BenchmarkDecodeLarge-4             10000            135991 ns/op           12881 B/op       1124 allocs/op
BenchmarkEncodeLarge-4           1888773               629 ns/op             160 B/op          2 allocs/op
BenchmarkEncodeLarge-4           1907782               620 ns/op             160 B/op          2 allocs/op
BenchmarkEncodeLarge-4           1934618               622 ns/op             160 B/op          2 allocs/op
BenchmarkEncodeLarge-4           1898805               683 ns/op             160 B/op          2 allocs/op
```

### benchstat
```
name               old time/op    new time/op    delta
DecodeSmall-4        2.93µs ± 1%    0.77µs ±19%     -73.57%  (p=0.029 n=4+4)
EncodeSmall-4         598ns ±20%     588ns ±10%        ~     (p=0.886 n=4+4)
StdDecodeMedium-4    40.2µs ±39%     6.4µs ± 0%     -83.99%  (p=0.029 n=4+4)
EncodeMedium-4       1.32µs ±48%    1.24µs ±17%        ~     (p=0.886 n=4+4)
DecodeLarge-4         269µs ± 5%     106µs ± 2%     -60.72%  (p=0.029 n=4+4)
EncodeLarge-4        1.19µs ± 9%    0.64µs ± 2%     -46.70%  (p=0.029 n=4+4)

name               old alloc/op   new alloc/op   delta
DecodeSmall-4          272B ± 0%       64B ± 0%     -76.47%  (p=0.029 n=4+4)
EncodeSmall-4          192B ± 0%      192B ± 0%        ~     (all equal)
StdDecodeMedium-4      496B ± 0%      384B ± 0%     -22.58%  (p=0.029 n=4+4)
EncodeMedium-4         224B ± 0%      224B ± 0%        ~     (all equal)
DecodeLarge-4          344B ± 0%    12881B ± 0%   +3644.48%  (p=0.029 n=4+4)
EncodeLarge-4          160B ± 0%      160B ± 0%        ~     (all equal)

name               old allocs/op  new allocs/op  delta
DecodeSmall-4          6.00 ± 0%      2.00 ± 0%     -66.67%  (p=0.029 n=4+4)
EncodeSmall-4          2.00 ± 0%      2.00 ± 0%        ~     (all equal)
StdDecodeMedium-4      11.0 ± 0%      41.0 ± 0%    +272.73%  (p=0.029 n=4+4)
EncodeMedium-4         2.00 ± 0%      2.00 ± 0%        ~     (all equal)
DecodeLarge-4          7.00 ± 0%   1124.00 ± 0%  +15957.14%  (p=0.029 n=4+4)
EncodeLarge-4          2.00 ± 0%      2.00 ± 0%        ~     (all equal)
```