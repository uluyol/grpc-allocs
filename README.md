# GRPC benchmarks with Go

Here are some benchmarks for using GRPC with Go and different codecs. Currently there are tests for proto3, gogoprotobuf with proto2, and flatbuffers.

```
BenchmarkProto3-4            	   10000	    7100 B/op	     147 allocs/op
BenchmarkProto2Gogo-4        	   10000	    7167 B/op	     142 allocs/op
BenchmarkFlatbufferTable-4   	   10000	    7596 B/op	     137 allocs/op
```

While flatbuffers appear to have the fewest allocations, it's unclear why there are over 100 allocs per request to begin with. Is this GRPC overhead? More measurements are TODO.
