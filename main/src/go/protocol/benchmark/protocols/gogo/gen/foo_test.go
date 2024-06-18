package gen

import (
	"testing"
)

func BenchmarkStruct_Marshal(b *testing.B) {
	ins := &TestStruct{
		Name: "test",
		Age:  12345,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ins.Marshal()
	}
}

func BenchmarkArray_Marshal(b *testing.B) {
	ins := &TestArray{
		Numbers: []int32{1, 2, 3, 4, 5},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ins.Marshal()
	}
}

func BenchmarkMap_Marshal(b *testing.B) {
	ins := &TestMap{
		Entries: map[string]int32{
			"a": 2,
			"b": 4,
			"c": 6,
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ins.Marshal()
	}
}

func BenchmarkStruct_Unmarshal(b *testing.B) {
	ins := &TestStruct{
		Name: "test",
		Age:  12345,
	}
	bs, err := ins.Marshal()
	if err != nil {
		panic(err)
	}
	ins2 := &TestStruct{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ins2.Unmarshal(bs)
	}
}
func BenchmarkArray_Unmarshal(b *testing.B) {
	ins := &TestArray{
		Numbers: []int32{1, 2, 3, 4, 5},
	}
	bs, err := ins.Marshal()
	if err != nil {
		panic(err)
	}
	ins2 := &TestArray{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ins2.Unmarshal(bs)
	}
}

func BenchmarkMap_Unmarshal(b *testing.B) {
	ins := &TestMap{
		Entries: map[string]int32{
			"a": 2,
			"b": 4,
			"c": 6,
		},
	}
	bs, err := ins.Marshal()
	if err != nil {
		panic(err)
	}
	ins2 := &TestMap{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ins2.Unmarshal(bs)
	}
}
