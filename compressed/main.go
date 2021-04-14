package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"unsafe"
)

// func main() {
// 	var uinthoge uint64 = 31312
// 	buf := make([]byte, binary.MaxVarintLen64)
// 	n := binary.PutUvarint(buf, uinthoge)
// 	fmt.Println(buf[:n])

// 	bytes := make([]byte, 41241)
// 	binary.BigEndian.PutUint64(bytes, uinthoge)
// 	fmt.Println(bytes)
// }

// func main() {
// 	h1 := NewHoge([]uint64{1, 2, 3, 4}, []uint64{3, 4, 5, 6})
// 	pp.Println(h1)

// 	buf := bytes.NewBuffer(nil)
// 	if err := gob.NewEncoder(buf).Encode(&h1); err != nil {
// 		log.Fatalln(err)
// 	}

// 	var h2 hoge
// 	ret := bytes.NewBuffer(buf.Bytes())
// 	if err := gob.NewDecoder(ret).Decode(&h2); err != nil {
// 		log.Fatalln(err)
// 	}
// 	pp.Println(h2)
// }

type hoge struct {
	F1 []uint64
	F2 []uint64
}

func NewHoge(f1, f2 []uint64) hoge {
	return hoge{
		F1: f1,
		F2: f2,
	}
}

func main() {
	// h1 := NewFuga([]uint64{4512351251235132523, 412421, 4512351251235132523, 452345235},
	// 	NewFuga([]uint64{4512351251235132523, 4512351251235132523, 4512351251235132523, 24135125, 4512351235132523, 4512351251235132523, 4512351251235132523, 4512351251235132523, 24135125, 4512351251235132523, 4512351251235132523, 4512351251235132523, 4512351251235132523, 24135125, 4512351251235132523, 4512351251235132523},
	// 		NewFuga([]uint64{6, 2, 7, 4512351251235132523, 4125125, 5215231521, 62435476645, 85685468797549, 84568465, 9879979, 5234},
	// 			NewFuga([]uint64{1512353125123512, 2, 32545423151521512, 4512351251235132523, 4125125, 5215231521, 62435476645, 85685468797549, 84568465, 9879979, 5234},
	// 				NewFuga([]uint64{7, 2, 32545423151521512, 4512351251235132523, 4125125, 5215231521, 62435476645, 85685468797549, 84568465, 9879979, 5234},
	// 					NewFuga([]uint64{1, 2, 51352, 24135125, 2513, 4}, nil))))))
	// 422

	h1 := NewFuga([]uint64{1, 23, 3, 4},
		NewFuga([]uint64{4, 5, 5, 24135125, 5, 8, 432847384728, 432847384728, 243284738472835125, 432847384728, 432847384728, 432847384728, 432847384728, 243284738472835125, 432847384728, 432847384728},
			NewFuga([]uint64{6, 2, 7, 432847384728, 21, 2133, 31, 3231, 84568465, 9879979, 5234},
				NewFuga([]uint64{21, 2, 32, 432847384728, 42, 32, 32, 42, 42, 42, 5234},
					NewFuga([]uint64{7, 2, 42, 432847384728, 43284738472825125, 42, 42, 42, 84568465, 9879979, 5234},
						NewFuga([]uint64{1, 2, 51352, 243284738472835125, 2513, 4}, nil))))))
	// pp.Println(h1)

	buf := bytes.NewBuffer(nil)
	if err := gob.NewEncoder(buf).Encode(&h1); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(len(buf.Bytes()))
	fmt.Println(unsafe.Sizeof(buf.Bytes()))
	fmt.Println(unsafe.Sizeof(buf.Cap()))

	var h2 fuga
	ret := bytes.NewBuffer(buf.Bytes())
	if err := gob.NewDecoder(ret).Decode(&h2); err != nil {
		log.Fatalln(err)
	}
	// pp.Println(h2)
}

type fuga struct {
	F1   []uint64
	Next *fuga
}

func NewFuga(f1 []uint64, n *fuga) *fuga {
	return &fuga{
		F1:   f1,
		Next: n,
	}
}
