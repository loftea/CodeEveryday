package leetcode

type UndergroundSystem struct {
	checkedTimes map[string]map[string]int
	total        map[string]map[string]int
	persons      map[int]*Person
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{checkedTimes: map[string]map[string]int{}, total: map[string]map[string]int{}, persons: map[int]*Person{}}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.persons[id] = &Person{checkIn: stationName, inTime: t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	checkIn := this.persons[id].checkIn
	delta := t - this.persons[id].inTime
	delete(this.persons, id)
	if _, ok := this.checkedTimes[checkIn]; !ok {
		this.checkedTimes[checkIn] = map[string]int{}
	}
	this.checkedTimes[checkIn][stationName]++
	if _, ok := this.total[checkIn]; !ok {
		this.total[checkIn] = map[string]int{}
	}
	this.total[checkIn][stationName] += delta
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	return float64(this.total[startStation][endStation]) / float64(this.checkedTimes[startStation][endStation])
}

type Info struct {
	id   int
	time int
}

type Person struct {
	checkIn string
	inTime  int
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
