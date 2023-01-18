# go-Golang-go
 using go for what it's claimed good for - concurrency.


![img](https://www.freecodecamp.org/news/content/images/2022/03/dancing-gopher.gif)


# Situation:
 You have to send salary to 4 men. You can send the salaries one by one using a for loop. But you want to make a faster parallel operation so that you can go home early.

## Simple For loop solution
 ```
 /* Go concurrency test
 *
 * author: ashraf minhaj
 * mail  : ashraf_minhaj@yahoo.com
  */

 package main
 
 import (
 	"log"
 	"time"
 )

 func main() {
 	start := time.Now()
 	log.Println("hey man")
 
 	// a for loop
 	salarySheet := [3]int{400, 500, 600}
 
 	for indx, salary := range salarySheet {
 		log.Println(indx, salary)
 		sendSalary(salary)
 	}
 	log.Println("Execution Time:", time.Since(start))
 }

 func sendSalary(salary int) {
 	log.Println("Sending salary, amount:", salary)
 }
```

 output:
 ```
 2023/01/19 00:24:43 hey man
 2023/01/19 00:24:43 0 400
 2023/01/19 00:24:43 Sending salary, amount: 400
 2023/01/19 00:24:43 1 500
 2023/01/19 00:24:43 Sending salary, amount: 500
 2023/01/19 00:24:43 2 600
 2023/01/19 00:24:43 Sending salary, amount: 600
 2023/01/19 00:24:43 Execution Time: 328.584µs
 ```
## with concurrency
 Strange!
 ```
 2023/01/19 00:36:38 hey man
 2023/01/19 00:36:38 0 400
 2023/01/19 00:36:38 1 500
 2023/01/19 00:36:38 2 600
 2023/01/19 00:36:38 Sending salary, amount: 600
 2023/01/19 00:36:38 Sending salary, amount: 400
 2023/01/19 00:36:38 Sending salary, amount: 500
 2023/01/19 00:36:38 Execution Time: 430.5µs
 ```

## Let's give it some load
 Now we add a little `delay` of 10 seconds using the `time.sleep()` function.

 Now it get's a little heavy and outputs - 
 ```
 2023/01/19 00:40:47 hey man
 2023/01/19 00:40:47 0 400
 2023/01/19 00:40:57 Sending salary, amount: 400
 2023/01/19 00:40:57 1 500
 2023/01/19 00:41:07 Sending salary, amount: 500
 2023/01/19 00:41:07 2 600
 2023/01/19 00:41:17 Sending salary, amount: 600
 2023/01/19 00:41:17 Execution Time: 30.003529167s
 ```
 with added concurrency with go...
 ```
 2023/01/19 00:44:02 hey man
 2023/01/19 00:44:02 0 400
 2023/01/19 00:44:02 1 500
 2023/01/19 00:44:02 2 600
 2023/01/19 00:44:12 Sending salary, amount: 400
 2023/01/19 00:44:12 Sending salary, amount: 600
 2023/01/19 00:44:12 Sending salary, amount: 500
 2023/01/19 00:44:12 Execution Time: 10.001873666s
 ```

 Interesting, isn't it? What took 30 seconds to process is now done in only 10!
 

(c) Ashraf Minhaj
