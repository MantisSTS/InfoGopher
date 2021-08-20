package memory

import (
	"fmt"

	"github.com/mozilla/masche/memaccess"
	"github.com/mozilla/masche/memsearch"
	"github.com/mozilla/masche/process"
)

type MemoryInfo struct {
	Address uintptr
}

func GetMemory() {
	p, _, _ := process.OpenFromPid(63295)
	memaccess.WalkMemory(p, uintptr(0x0), 100, func(address uintptr, buf []byte) bool {
		fmt.Println(address, string(buf))
		return true
	})
}

func FindStringInMemory(searchString string) {
	p, _, _ := process.OpenFromPid(63295)
	fmt.Println(memsearch.FindBytesSequence(p, uintptr(0x0), []byte(searchString)))

}
