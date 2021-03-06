package uploader_test

import (
	"fmt"
	"github.com/tickstep/library-go/cachepool"
	"github.com/tickstep/library-go/requester/rio"
	"github.com/tickstep/cloudpan189-go/library/requester/transfer"
	"github.com/tickstep/cloudpan189-go/internal/file/uploader"
	"io"
	"testing"
)

var (
	blockList = uploader.SplitBlock(10000, 999)
)

func TestSplitBlock(t *testing.T) {
	for k, e := range blockList {
		fmt.Printf("%d %#v\n", k, e)
	}
}

func TestSplitUnitRead(t *testing.T) {
	var size int64 = 65536*2+3432
	buffer := rio.NewBuffer(cachepool.RawMallocByteSlice(int(size)))
	unit := uploader.NewBufioSplitUnit(buffer, transfer.Range{Begin: 2, End: size}, nil, nil)

	buf := cachepool.RawMallocByteSlice(1022)
	for {
		n, err := unit.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatalf("read error: %s\n", err)
		}
		fmt.Printf("n: %d, left: %d\n", n, unit.Left())
	}
}
