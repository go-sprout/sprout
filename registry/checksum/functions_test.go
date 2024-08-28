package checksum_test

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash/adler32"
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/checksum"
)

func TestSha1Sum(t *testing.T) {
	noHash := sha1.Sum([]byte(""))
	soloHash := sha1.Sum([]byte("a"))
	multiHash := sha1.Sum([]byte("hello world"))

	var tc = []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{sha1Sum ""}}`, ExpectedOutput: hex.EncodeToString(noHash[:])},
		{Name: "TestSingleByteInput", Input: `{{sha1Sum "a"}}`, ExpectedOutput: hex.EncodeToString(soloHash[:])},
		{Name: "TestMultiByteInput", Input: `{{sha1Sum "hello world"}}`, ExpectedOutput: hex.EncodeToString(multiHash[:])},
	}
	pesticide.RunTestCases(t, checksum.NewRegistry(), tc)
}

func TestSha256Sum(t *testing.T) {
	noHash := sha256.Sum256([]byte(""))
	soloHash := sha256.Sum256([]byte("a"))
	multiHash := sha256.Sum256([]byte("hello world"))

	var tc = []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{sha256Sum ""}}`, ExpectedOutput: hex.EncodeToString(noHash[:])},
		{Name: "TestSingleByteInput", Input: `{{sha256Sum "a"}}`, ExpectedOutput: hex.EncodeToString(soloHash[:])},
		{Name: "TestMultiByteInput", Input: `{{sha256Sum "hello world"}}`, ExpectedOutput: hex.EncodeToString(multiHash[:])},
	}
	pesticide.RunTestCases(t, checksum.NewRegistry(), tc)
}

func TestSha512sum(t *testing.T) {
	noHash := sha512.Sum512([]byte(""))
	soloHash := sha512.Sum512([]byte("a"))
	multiHash := sha512.Sum512([]byte("hello world"))

	var tc = []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{sha512Sum ""}}`, ExpectedOutput: hex.EncodeToString(noHash[:])},
		{Name: "TestSingleByteInput", Input: `{{sha512Sum "a"}}`, ExpectedOutput: hex.EncodeToString(soloHash[:])},
		{Name: "TestMultiByteInput", Input: `{{sha512Sum "hello world"}}`, ExpectedOutput: hex.EncodeToString(multiHash[:])},
	}
	pesticide.RunTestCases(t, checksum.NewRegistry(), tc)
}

func TestAdler32Sum(t *testing.T) {
	noHash := adler32.Checksum([]byte(""))
	soloHash := adler32.Checksum([]byte("a"))
	multiHash := adler32.Checksum([]byte("hello world"))

	var tc = []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{adler32Sum ""}}`, ExpectedOutput: fmt.Sprint(noHash)},
		{Name: "TestSingleByteInput", Input: `{{adler32Sum "a"}}`, ExpectedOutput: fmt.Sprint(soloHash)},
		{Name: "TestMultiByteInput", Input: `{{adler32Sum "hello world"}}`, ExpectedOutput: fmt.Sprint(multiHash)},
	}
	pesticide.RunTestCases(t, checksum.NewRegistry(), tc)
}

func TestMD5Sum(t *testing.T) {
	noHash := md5.Sum([]byte(""))
	soloHash := md5.Sum([]byte("a"))
	multiHash := md5.Sum([]byte("hello world"))

	var tc = []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{md5Sum ""}}`, ExpectedOutput: hex.EncodeToString(noHash[:])},
		{Name: "TestSingleByteInput", Input: `{{md5Sum "a"}}`, ExpectedOutput: hex.EncodeToString(soloHash[:])},
		{Name: "TestMultiByteInput", Input: `{{md5Sum "hello world"}}`, ExpectedOutput: hex.EncodeToString(multiHash[:])},
	}
	pesticide.RunTestCases(t, checksum.NewRegistry(), tc)
}
