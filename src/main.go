/* Go concurrency test
*
* author: ashraf minhaj
* mail  : ashraf_minhaj@yahoo.com
 */

package main

import (
	"log"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	start := time.Now()
	log.Println("hey man")

	// a for loop
	salarySheet := [3]int{400, 500, 600}

	for indx, salary := range salarySheet {
		log.Println(indx, salary)

		wg.Add(1)
		go sendSalary(salary)
	}
	wg.Wait()
	defer log.Println("Execution Time:", time.Since(start))
}

func sendSalary(salary int) {
	time.Sleep(10 * time.Second)
	log.Println("Sending salary, amount:", salary)
	wg.Done()
}
