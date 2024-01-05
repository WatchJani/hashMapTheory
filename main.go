package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

// brzo jeste ali ne i dovoljno
// teorija pada u vodu, nije sortirano
func main() {
	memTable := New(24)

	memTable.Append(41, rand.Intn(10000))
	memTable.Append(4, rand.Intn(10000))
	memTable.Append(1, rand.Intn(10000))

	value, err := memTable.GetValue(41)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(memTable.GetTable())

	fmt.Println(value)

	memTable.Read()
}

type MemTable struct {
	table    map[int]int
	counter  int
	capacity int
}

func New(capacity int) *MemTable {
	return &MemTable{
		table:    make(map[int]int, capacity),
		capacity: capacity,
	}
}

func (m *MemTable) Append(key, value int) {
	m.counterIncrease()

	if m.counter > m.capacity {
		m.reset()
		// m.resetTable()
		m.table = make(map[int]int, m.capacity)
	}

	m.insert(key, value)
}

func (m *MemTable) resetTable() {
	m.table = nil
}

func (m *MemTable) reset() {
	m.counter = 0
}

func (m *MemTable) insert(key, value int) {
	m.table[key] = value
}

func (m *MemTable) counterIncrease() {
	m.counter++
}

func (m MemTable) GetTable() map[int]int {
	return m.table
}

func (m MemTable) GetValue(key int) (int, error) {
	if value, ok := m.table[key]; ok {
		return value, nil
	}

	return -1, errors.New("Cant find this key in memTable")
}

func (m MemTable) Read() {
	for key, value := range m.table {
		fmt.Println("key:", key, "value:", value)
	}
}
