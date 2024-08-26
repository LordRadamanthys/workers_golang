package main

import (
	"fmt"
	"sync"
	"time"
)

type Cars struct {
	Name       string
	Engine     float64
	Type       string
	Aspiration string
	Status     string
}

func (c *Cars) update(wg *sync.WaitGroup) {
	c.Status = "Updated"
	wg.Done()
}

func worker(jobs <-chan *Cars, results chan<- Cars) {
	for j := range jobs {
		j.Status = "updated"
		results <- *j
	}

}

func main() {
	carsList := []*Cars{
		{
			Name:       "Supra",
			Engine:     3.6,
			Type:       "hatch",
			Aspiration: "turbo",
		},
		{
			Name:       "Evo 7",
			Engine:     2.5,
			Type:       "sedan",
			Aspiration: "turbo",
		},
		{
			Name:       "GTR 34",
			Engine:     3.5,
			Type:       "sedan",
			Aspiration: "biturbo",
		},
	}

	//carWorkers(carsList)
	carUpdated(carsList)
}

func carWorkers(carsList []*Cars) {
	var carsListUpdated []Cars
	ch := make(chan *Cars)
	results := make(chan Cars, len(carsList))
	timeInit := time.Now()

	for i := 0; i < 1; i++ {
		go worker(ch, results)
	}

	for _, car := range carsList {
		ch <- car
	}

	for i := 0; i < len(carsList); i++ {
		carsListUpdated = append(carsListUpdated, <-results)
	}

	close(ch)
	close(results)
	fmt.Println(carsListUpdated)
	timeEnd := time.Now()
	fmt.Println("Time elapsed carWorkers: ", timeEnd.Sub(timeInit))
}

func carUpdated(carList []*Cars) {
	wg := new(sync.WaitGroup)
	wg.Add(len(carList))
	timeInit := time.Now()
	for _, car := range carList {
		go car.update(wg)
	}
	wg.Wait()

	fmt.Println(carList[0])
	timeEnd := time.Now()
	fmt.Println("Time elapsed carUpdated: ", timeEnd.Sub(timeInit))
}
