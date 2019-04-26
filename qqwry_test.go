package qqwry

import (
	"log"
	"testing"
)

func TestQQwry(t *testing.T) {
	q := NewQQwry("qqwry.dat")

	for i := 0; i < 5; i++ {
		ip := q.Find("183.224.52.133")
		log.Printf("ip:%v, country:%v, city:%v\n", ip.Ip, ip.Country, ip.City)
	}
	ip := q.Find("8.8.8.8")
	log.Printf("ip:%v, country:%v, city:%v\n", ip.Ip, ip.Country, ip.City)
	ip = q.Find("114.114.114.114")
	log.Printf("ip:%v, country:%v, city:%v\n", ip.Ip, ip.Country, ip.City)
}
