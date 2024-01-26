package raid

import (
	"fmt"
	"graidtechHW/util"
	"strings"
)

type Raid1 struct {
	disks []*Disk
}

func NewRaid1() IRaid {

	r := new(Raid1)

	// 產生硬碟
	r.disks = make([]*Disk, 2)

	// len := util.RandRange(util.DISK_LEN_MIN, util.DISK_LEN_MAX)
	len := int64(util.DISK_LEN_MIN)
	// 硬碟容量大小不一樣
	for i := range r.disks {

		r.disks[i] = NewDisk(len)

	}

	for i, v := range r.disks {
		fmt.Printf("read 1 第 %v 硬碟 長度 %v \n", i+1, v.Len())
	}

	return r
}

// 寫入
func (r *Raid1) Write(position int, str string) {
	bIdx := 0

	distLen := r.disks[0].Len()

	for i := position; i < position+len(str); i++ {
		idx := i % distLen
		for _, v := range r.disks {
			v.Write(idx, []byte{str[bIdx]})
		}
		bIdx++
	}

}

// 讀取
func (r *Raid1) Read(position int, count int) string {
	var result strings.Builder

	distLen := r.disks[0].Len()

	for i := position; i < position+count; i++ {
		if i >= distLen {
			break
		}
		idx := i % distLen

		for _, v := range r.disks {
			b := v.Read(idx)

			if len(b) != 0 {
				for _, v := range b {
					result.WriteByte(v)
				}
				break
			}

		}

	}

	return result.String()
}

// 清除
func (r *Raid1) Clear(idx int) {
	if idx < 0 || idx >= len(r.disks) {
		fmt.Println("Clear err idx nil")
		return
	}

	r.disks[idx] = NewDisk(int64(r.disks[idx].Len()))
}
