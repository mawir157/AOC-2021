package main

import AH "./adventhelper"

import (
	"encoding/hex"
)

// -- (version, type, value, subpackets)
type Packet struct {
	version    int
	typeID     int
	value      int
	subPackets []Packet
}

func htob(s string) (bits []bool) {
	hexData, _ := hex.DecodeString(s)
	for _, byte := range hexData {
		for i := 0; i < 8; i++ {
			bit := byte & 0x80 > 0
			bits = append(bits, bit)
			byte = byte << 1
		}
	}
	return
}

func btoi(bs []bool) (i int) {
	i = 0
	for j := 0; j < len(bs); j++ {
		i *= 2

		if bs[j] {
			i += 1
		}
	}

	return
}

func parseInput(bs []bool, p uint) ([]Packet, uint) {
	pckts := []Packet{}

	version := btoi( bs[p:(p + 3)] )
	p += 3
	typeID := btoi( bs[p:(p + 3)] )
	p += 3

	if typeID == 4 { // literal packet
		value := 0
		for ; true; {
			last := !(bs[p])
			p += 1
			value *= 16
			value += btoi( bs[p:(p+4)] )

			p += 4
			if last {
				break
			}
		}
		pckt := Packet{version:version, typeID:typeID, value:value}
		pckts = append(pckts, pckt)
	} else { // an operator packet
		var subPackets []Packet
		var temp []Packet
		lengthId := bs[p]
		p += 1
		if ( !lengthId ) { // bit count given
			next := uint(btoi(bs[p:(p+15)])) // how many bits for the subpacket
			p += 15
			end := p + next
			
			for ; p < end; {
				temp, p = parseInput(bs, p)
				subPackets = append(subPackets, temp...)
			}
			pckt := Packet{version:version, typeID:typeID, subPackets:subPackets}
			pckts = append(pckts, pckt)
		} else { // number of packets given
			length := uint(btoi(bs[p:(p+11)])) // how many sub packets
			p += 11
			for i := uint(0); i < length; i++ {
				temp, p = parseInput(bs, p)
				subPackets = append(subPackets, temp...)
			}
			pckt := Packet{version:version, typeID:typeID, subPackets:subPackets}
			pckts = append(pckts, pckt)
		}
	}

	
	return pckts, p
}

func score1(pckts []Packet) (sum int) {
	sum = 0
	for _, p := range pckts {
		sum += p.version
		sum += score1(p.subPackets)
	}

	return
}

func score2(p Packet) (score int) {
	score = 0
	temp := 0
	if len(p.subPackets) == 0 {
		return p.value
	} else {
		switch p.typeID {
			case 0: // sum packet
				for _, sp := range p.subPackets {
					temp += score2(sp)
				}
			case 1: // product packet
				temp = 1;
				for _, sp := range p.subPackets {
					temp *= score2(sp)
				}
			case 2: // minimum packet
				temp = 1 << 30;
				for _, sp := range p.subPackets {
					sc := score2(sp)
					if sc < temp {
						temp = sc
					}
				}
			case 3: // maximum packet
				for _, sp := range p.subPackets {
					sc := score2(sp)
					if sc > temp {
						temp = sc
					}
				}
			// case 4:
			case 5: // greater than packet
				p0 := score2(p.subPackets[0])
				p1 := score2(p.subPackets[1])
				if p0 > p1 {
					temp = 1
				}
			case 6: // less than packet
				p0 := score2(p.subPackets[0])
				p1 := score2(p.subPackets[1])
				if p0 < p1 {
					temp = 1
				}
			case 7: // equal packet
				p0 := score2(p.subPackets[0])
				p1 := score2(p.subPackets[1])
				if p0 == p1 {
					temp = 1
				}
			}
			score += temp
	}

	return
}

func part1(s string) int {
	pckt, _ := parseInput(htob(s), 0)
	return score1(pckt)
}

func part2(s string) int {
	pckt, _ := parseInput(htob(s), 0)
	return score2(pckt[0])
}

func main() {
	ss, _ := AH.ReadStrFile("../input/input16.txt")
	AH.PrintSoln(16, part1(ss[0]), part2(ss[0]))

	return
}
