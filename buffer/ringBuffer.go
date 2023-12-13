package buffer

import "sync"

// Структура кольцеого буфера
type RingIntBuffer struct {
	data  []int //данные
	sendx int   //индекс ячейки куда писать
	resvx int   //индекс ячейки откуда начинать читать
	size  int   //размер буфера
	count int   //количество данных в буфере
	m     sync.Mutex
}

func (r *RingIntBuffer) Push(item int) {
	r.m.Lock()
	defer r.m.Unlock()

	if r.count == r.size {
		r.resvx++
		r.count--
	}

	r.data[r.sendx] = item
	r.sendx++
	r.count++

	if r.sendx == r.size {
		r.sendx = 0
	}
	if r.resvx == r.size {
		r.resvx = 0
	}
}

func (r *RingIntBuffer) Get() []int {
	if r.count == 0 {
		return nil
	}

	r.m.Lock()
	defer r.m.Unlock()

	retData := make([]int, r.count)
	for i := range retData {
		retData[i] = r.data[r.resvx]
		r.resvx++
		if r.resvx == r.size {
			r.resvx = 0
		}
	}
	r.count = 0
	return retData
}

func NewRingIntBuffer(size int) *RingIntBuffer {
	return &RingIntBuffer{
		data:  make([]int, size),
		sendx: 0,
		resvx: 0,
		size:  size,
		count: 0,
		m:     sync.Mutex{},
	}
}
