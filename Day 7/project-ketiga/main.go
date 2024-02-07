package main

import (
	"fmt"
)

func main() {
	var u1 kc = unit{nama: "Unit Benda Raya", alamat: "Kemang"}
	var k1 kc = kcp{nama: "KCP Patrajasa", alamat: "Gatot Subroto"}

	fmt.Println(u1.kur())
	fmt.Println(u1.tabungan())
	fmt.Println(u1.(unit).simpedes())
	fmt.Println(u1.(unit).kupedes())
	fmt.Println(k1.kur())
	fmt.Println(k1.tabungan())
}

// Day 7 Exercises
type kc interface {
	kur() string
	tabungan() string
}

type unit struct {
	nama, alamat string
}

type kcp struct {
	nama, alamat string
}

func (u unit) kur() string {
	return "Kredit Usaha Rakyat oleh " + u.nama
}

func (u unit) tabungan() string {
	return "Simpanan Umum oleh " + u.nama
}

func (u unit) kupedes() string {
	return "Kredit Usaha Pedesaan oleh " + u.nama
}

func (u unit) simpedes() string {
	return "Simpanan Masyarakat Pedesaan oleh " + u.nama
}

func (k kcp) kur() string {
	return "Kredit Usaha Rakyat oleh " + k.nama
}

func (k kcp) tabungan() string {
	return "Simpanan Umum oleh " + k.nama
}
