package backend

import (
	"math/rand"
	"time"
	"os"
	"fmt"
	"crypto/md5"
	"encoding/binary"
)

func Prase() string {
	var buf [12]byte
	//time
	binary.BigEndian.PutUint32(buf[:], uint32(time.Now().Unix()))
	//machine
	b := getMachineBytes()
	buf[4] = b[0]
	buf[5] = b[1]
	buf[6] = b[2]
	//pid
	pid := getPid()
	buf[7] = byte(pid >> 8)
	buf[8] = byte(pid)
	//inc
	i := genRandom()
	buf[9] = byte(i >> 16)
	buf[10] = byte(i >> 8)
	buf[11] = byte(i)
	return fmt.Sprintf("%x", buf)
}

func genRandom() int32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int31()
}

func getPid() int32 {
	return int32(os.Getpid())
}

func getTime() int64 {
	return time.Now().Unix()
}

func getMachineBytes() []byte {
	var id [3]byte
	ret := id[:]
	h, err := os.Hostname()
	if nil != err {
		//if cannot get hostname return a random
		binary.BigEndian.PutUint32(ret, uint32(genRandom()))
	}
	m := md5.New()
	m.Write([]byte(h))
	copy(ret, m.Sum(nil))
	return ret
}
