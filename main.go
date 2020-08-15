package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	primes "github.com/Azorej/go-prime-generator"
)

// Response object format
type Response struct {
	PrimeNumber int `json:"prime_number"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &Response{PrimeNumber: ProducePrime()})
	})
	e.Logger.Fatal(e.Start(":3000"))
}

// ProducePrime returns a single prime number from a sieve
// Note: The seive could have been precomputed, but the aim of this
// API is to stress test CPU resources
func ProducePrime() int {
	count := 10000000
	primes := generate(count, time.Second*2)
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r1.Intn(len(primes))
	return primes[index]
}

func generate(maxN int, timeout time.Duration) []int {
	ch := make(chan int, 100)
	enough := primes.NewChanSignal()

	go primes.Generate(maxN, ch, enough)

	timer := time.NewTimer(timeout)
	res := make([]int, 0)

	for {
		select {
		case v, more := <-ch:
			if !more {
				return res
			}
			res = append(res, v)

		case <-timer.C:
			enough <- primes.NewSignal()
			println("Timeout :(")
			return res
		}
	}
}
