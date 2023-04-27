package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Person struct {
	sal, level int
	rwm        sync.RWMutex
}

func (p *Person) promote() {
	p.rwm.Lock()

	log.Printf("**orginal** ")
	p.sal += 1000
	p.level++
	log.Printf("==changed== ")

	p.rwm.Unlock()
}

func (p *Person) String() string {
	p.rwm.RLock()
	personStr := fmt.Sprintf("Person{sal = %d, level = %d}", p.sal, p.level)
	p.rwm.RUnlock()
	return personStr
}

func main() {
	p := &Person{sal: 1000, level: 1}

	for i := 0; i < 3; i++ {
		go log.Println(p)
	}

	go p.promote()

	time.Sleep(time.Second)
}
