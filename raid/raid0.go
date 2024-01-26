package raid

import (
	"fmt"
	"graidtechHW/util"
	"strings"
)

type Raid0 struct {
	disks []*Disk

	totalDiskPos []DiskPos // 所有儲存位置
}

func NewRaid0() IRaid {

	r := new(Raid0)

	// 產生硬碟
	r.disks = make([]*Disk, 2)
	var totalLen int64
	var maxLen int64
	// 硬碟容量大小不一樣
	for i := range r.disks {
		// len := util.RandRange(util.DISK_LEN_MIN, util.DISK_LEN_MAX)
		len := int64(util.DISK_LEN_MIN)

		r.disks[i] = NewDisk(len)
		totalLen += len

		if len > maxLen {
			maxLen = len
		}
	}

	r.totalDiskPos = make([]DiskPos, totalLen)

	idx := 0
	for i := 0; i < int(maxLen); i++ {
		for j, v := range r.disks {
			if i < v.Len() {
				r.totalDiskPos[idx] = DiskPos{
					X: j,
					Y: i,
				}
				idx++
			}
		}
	}

	for i, v := range r.disks {
		fmt.Printf("read 0 第 %v 硬碟 長度 %v \n", i+1, v.Len())
	}

	fmt.Printf("read 0 總長度 %v \n", len(r.totalDiskPos))

	return r
}

// 寫入
func (r *Raid0) Write(position int, str string) {
	bIdx := 0

	for i := position; i < position+len(str); i++ {
		idx := i % len(r.totalDiskPos)
		pos := r.totalDiskPos[idx]
		r.disks[pos.X].Write(pos.Y, []byte{str[bIdx]})
		bIdx++
	}
}

// 讀取
func (r *Raid0) Read(position int, count int) string {
	var result strings.Builder

	for i := position; i < position+count; i++ {
		idx := i % len(r.totalDiskPos)
		pos := r.totalDiskPos[idx]
		b := r.disks[pos.X].Read(pos.Y)

		// raid 0 假如有資料不見 就無法找到
		if len(b) == 0 {
			return ""
		}

		for _, v := range b {
			result.WriteByte(v)
		}

	}

	return result.String()
}

// 清除
func (r *Raid0) Clear(idx int) {
	if idx < 0 || idx >= len(r.disks) {
		fmt.Println("Clear err idx nil")
		return
	}

	r.disks[idx] = NewDisk(int64(r.disks[idx].Len()))
}
