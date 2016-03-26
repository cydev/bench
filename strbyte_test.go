package bench

import (
	"math/rand"
	"testing"
	"time"
	"reflect"
	"unsafe"
)

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

type HasName interface {
	Name() string
}

type HasByteName interface {
	Name() []byte
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type WithNameS struct {
	name string
}

func (s WithNameS) Name() string {
	return s.name
}

type WithNameB struct {
	name []byte
}

func (s WithNameB) Name() []byte {
	return s.name
}

func (s WithNameB) NameS() string {
	return string(s.name)
}

func (s WithNameB) NameSU() string {
	return BytesToString(s.name)
}

func nBenchStringInterface(b *testing.B, length int) {
	testWithNameN := WithNameS{RandStringRunes(length)}
	testWithNameInt := HasName(testWithNameN)
	b.ResetTimer()
	var s string
	for i := 0; i < b.N; i++ {
		s = testWithNameInt.Name()
	}
	if len(s) == 0 {
		panic("zero")
	}
}

func nBenchByteInterface(b *testing.B, length int) {
	testWithNameN := WithNameB{[]byte(RandStringRunes(length))}
	testWithNameInt := HasByteName(testWithNameN)
	b.ResetTimer()
	var s []byte
	for i := 0; i < b.N; i++ {
		s = testWithNameInt.Name()
	}
	if len(s) == 0 {
		panic("zero")
	}
}

func nBenchString(b *testing.B, length int) {
	testWithNameN := WithNameS{RandStringRunes(length)}
	b.ResetTimer()
	var s string
	for i := 0; i < b.N; i++ {
		s = testWithNameN.Name()
	}
	if len(s) == 0 {
		panic("zero")
	}
}

func nBenchByte(b *testing.B, length int) {
	testWithNameN := WithNameB{[]byte(RandStringRunes(length))}
	b.ResetTimer()
	var s []byte
	for i := 0; i < b.N; i++ {
		s = testWithNameN.Name()
	}
	if len(s) == 0 {
		panic("zero")
	}
}

func nBenchByteToS(b *testing.B, length int) {
	testWithNameN := WithNameB{[]byte(RandStringRunes(length))}
	b.ResetTimer()
	var s string
	for i := 0; i < b.N; i++ {
		s = testWithNameN.NameS()
	}
	if len(s) == 0 {
		panic("zero")
	}
}

func nBenchByteToSU(b *testing.B, length int) {
	testWithNameN := WithNameB{[]byte(RandStringRunes(length))}
	b.ResetTimer()
	var s string
	for i := 0; i < b.N; i++ {
		s = testWithNameN.NameSU()
	}
	if len(s) == 0 {
		panic("zero")
	}
}

func BenchmarkNameString1024(b *testing.B) { nBenchString(b, 1024) }

func BenchmarkNameByte1024(b *testing.B) { nBenchByte(b, 1024) }

func BenchmarkNameByteToS100(b *testing.B) { nBenchByteToS(b, 100) }

func BenchmarkNameByteToSU100(b *testing.B) { nBenchByteToSU(b, 100) }

func BenchmarkNameString100(b *testing.B) { nBenchString(b, 100) }

func BenchmarkNameStringInterface100(b *testing.B) { nBenchStringInterface(b, 100) }

func BenchmarkNameByteInterface100(b *testing.B) { nBenchByteInterface(b, 100) }

func BenchmarkNameByte100(b *testing.B) { nBenchByte(b, 100) }

func BenchmarkNameString10(b *testing.B) { nBenchString(b, 10) }

func BenchmarkNameByte10(b *testing.B) {
	nBenchByte(b, 10)
}
