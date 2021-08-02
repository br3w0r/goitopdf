package itopdf

import (
	"fmt"
	"os"
	"testing"
	"time"
)

type testIToPDF struct {
	iToPDF
}

func (pdf *testIToPDF) AddImage(path string) error {
	time.Sleep(300)
	return nil
}

type testFileInfo struct {
}

func newTestFileInfo() os.FileInfo {
	return &testFileInfo{}
}

func (fi *testFileInfo) Name() string {
	return "test"
}

func (fi *testFileInfo) Size() int64 {
	return 1
}

func (fi *testFileInfo) Mode() os.FileMode {
	return os.ModeAppend
}

func (fi *testFileInfo) ModTime() time.Time {
	return time.Now()
}

func (fi *testFileInfo) IsDir() bool {
	return false
}

func (fi *testFileInfo) Sys() interface{} {
	return nil
}

func BenchmarkItoPDF(b *testing.B) {
	var err error
	fileInfo := newTestFileInfo()

	for i := 0; i < b.N; i++ {
		pdf := &testIToPDF{}
		walkFunc := pdf.walkFunc(nil)

		walkFunc(fmt.Sprintf("test%d", i), fileInfo, err)
		if err != nil {
			b.Fatalf("Error: %v", err)
		}
	}
}
