package raid

type DiskPos struct {
	X int
	Y int
}
type Disk struct {
	// Datas []byte
	Datas [][]byte
}

func NewDisk(len int64) *Disk {
	return &Disk{
		Datas: make([][]byte, len),
	}
}

func (d *Disk) Len() int {
	return len(d.Datas)
}

func (d *Disk) Read(idx int) []byte {
	if idx >= len(d.Datas) {
		return nil
	}

	return d.Datas[idx]
}
func (d *Disk) Write(pos int, b []byte) {
	d.Datas[pos] = make([]byte, len(b))

	for i, v := range b {
		d.Datas[pos][i] = v
	}
}
