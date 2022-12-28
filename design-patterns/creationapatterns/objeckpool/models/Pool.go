package models

import (
	"creationapatterns/objeckpool/interfaces"
	"fmt"
	"sync"
)

type Pool struct {
	Idle     []interfaces.IPoolObject
	Active   []interfaces.IPoolObject
	Capacity int
	Mulock   *sync.Mutex
}

func InitPool(poolObject []interfaces.IPoolObject) (*Pool, error) {
	if len(poolObject) == 0 {
		return nil, fmt.Errorf("cannot create a pool of 0 length")
	}

	active := make([]interfaces.IPoolObject, 0)
	pool := &Pool{
		Idle:     poolObject,
		Active:   active,
		Capacity: len(poolObject),
		Mulock:   new(sync.Mutex),
	}

	return pool, nil
}

func (p *Pool) Loan() (interfaces.IPoolObject, error) {
	p.Mulock.Lock()
	defer p.Mulock.Unlock()
	if len(p.Idle) == 0 {
		return nil, fmt.Errorf("no pool object free. Please request after sometime")
	}
	obj := p.Idle[0]
	p.Idle = p.Idle[1:]
	p.Active = append(p.Active, obj)
	fmt.Printf("Loan Pool Object with ID: %s\n", obj.GetID())
	return obj, nil
}

func (p *Pool) Receive(target interfaces.IPoolObject) error {
	p.Mulock.Lock()
	defer p.Mulock.Unlock()
	err := p.Remove(target)
	if err != nil {
		return err
	}
	p.Idle = append(p.Idle, target)
	fmt.Printf("Return Pool Object with ID: %s\n", target.GetID())
	return nil
}

func (p *Pool) Remove(target interfaces.IPoolObject) error {
	currentActiveLength := len(p.Active)
	for i, obj := range p.Active {
		if obj.GetID() == target.GetID() {
			p.Active[currentActiveLength-1], p.Active[i] = p.Active[i], p.Active[currentActiveLength-1]
			p.Active = p.Active[:currentActiveLength-1]
			return nil
		}
	}
	return fmt.Errorf("target pool object doesn't belong to the pool")
}
