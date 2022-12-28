package objeckpool

import (
	"creationapatterns/objeckpool/interfaces"
	"creationapatterns/objeckpool/models"
	"log"
	"strconv"
)

func ObjeckPool() {
	connections := make([]interfaces.IPoolObject, 0)
	for i := 0; i < 3; i++ {
		c := &models.Connection{Id: strconv.Itoa(i)}
		connections = append(connections, c)
	}
	pool, err := models.InitPool(connections)
	if err != nil {
		log.Fatalf("Init Pool Error: %s", err)
	}
	conn1, err := pool.Loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	conn2, err := pool.Loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	pool.Receive(conn1)
	pool.Receive(conn2)
}
