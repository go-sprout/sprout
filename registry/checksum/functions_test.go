package checksum_test

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash/adler32"
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/checksum"
)

func TestSha1sum(t *testing.T) {
	noHash := sha1.Sum([]byte(""))
	soloHash := sha1.Sum([]byte("a"))
	multiHash := sha1.Sum([]byte("hello world"))

	var tc = []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{sha1sum ""}}`, Expected: hex.EncodeToString(noHash[:])},
		{Name: "TestSingleByteInput", Input: `{{sha1sum "a"}}`, Expected: hex.EncodeToString(soloHash[:])},
		{Name: "TestMultiByteInput", Input: `{{sha1sum "hello world"}}`, Expected: hex.EncodeToString(multiHash[:])},
	}
	pesticide.RunTestCases(t, checksum.NewRegistry(), tc)
}

func TestSha256sum(t *testing.T) {
	noHash := sha256.Sum256([]byte(""))
	soloHash := sha256.Sum256([]byte("a"))
	multiHash := sha256.Sum256([]byte("hello world"))

	var tc = []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{sha256sum ""}}`, Expected: hex.EncodeToString(noHash[:])},
		{Name: "TestSingleByteInput", Input: `{{sha256sum "a"}}`, Expected: hex.EncodeToString(soloHash[:])},
		{Name: "TestMultiByteInput", Input: `{{sha256sum "hello world"}}`, Expected: hex.EncodeToString(multiHash[:])},
	}
	pesticide.RunTestCases(t, checksum.NewRegistry(), tc)
}

func TestAdler32sum(t *testing.T) {
	noHash := adler32.Checksum([]byte(""))
	soloHash := adler32.Checksum([]byte("a"))
	multiHash := adler32.Checksum([]byte("hello world"))

	var tc = []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{adler32sum ""}}`, Expected: fmt.Sprintf("%d", noHash)},
		{Name: "TestSingleByteInput", Input: `{{adler32sum "a"}}`, Expected: fmt.Sprintf("%d", soloHash)},
		{Name: "TestMultiByteInput", Input: `{{adler32sum "hello world"}}`, Expected: fmt.Sprintf("%d", multiHash)},
	}
	pesticide.RunTestCases(t, checksum.NewRegistry(), tc)
}

func TestMD5sum(t *testing.T) {
	noHash := md5.Sum([]byte(""))
	soloHash := md5.Sum([]byte("a"))
	multiHash := md5.Sum([]byte("hello world"))

	var tc = []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{md5sum ""}}`, Expected: hex.EncodeToString(noHash[:])},
		{Name: "TestSingleByteInput", Input: `{{md5sum "a"}}`, Expected: hex.EncodeToString(soloHash[:])},
		{Name: "TestMultiByteInput", Input: `{{md5sum "hello world"}}`, Expected: hex.EncodeToString(multiHash[:])},
	}
	pesticide.RunTestCases(t, checksum.NewRegistry(), tc)
}
