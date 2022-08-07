package main

import "fmt"

type Kisi struct {
	id     int
	adi    string
	soyadi string
}

// Mysql
type Mysql struct {
	kisi Kisi
}

func (m Mysql) Kaydet() {
	fmt.Println(m.kisi.adi, " ya ait veriler mysql e kaydedildi.")
}

// Oracle
type Oracle struct {
	kisi Kisi
}

func (o Oracle) Kaydet() {
	fmt.Println(o.kisi.adi, " ya ait veriler Oracle e kaydedildi.")
}

type Kaydedici interface {
	Kaydet()
}

func vtKaydet(k Kaydedici) {
	k.Kaydet()
}

func main() {
	k := Kisi{
		id:     1,
		adi:    "Mustafa Çağlar",
		soyadi: "KARA",
	}
	m := Oracle{k}
	m2 := Mysql{k}

	vtKaydet(m)
	vtKaydet(m2)
	//var m Kaydedici = Mysql{k}
	//var l Kaydedici = Oracle{k}
	//
	//m.Kaydet()
	//l.Kaydet()

}
