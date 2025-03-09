package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Response struct {
	res int
	err error
}

func main() {
	start := time.Now()

	//ctx := context.Background()
	ctx := context.WithValue(context.Background(), "foo", Response{res: 111, err: fmt.Errorf("something bad happened")})
	userId := 10

	res, err := fetchUserData(ctx, userId)

	if err != nil {
		log.Fatal(err)
	}

	println("data:", res)
	println("request took:", time.Since(start))

}

func fetchUserData(ctx context.Context, userId int) (int, error) {
	val := ctx.Value("foo")

	fmt.Printf("value from context: %s\n", val)
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	respCh := make(chan Response)

	go func() {
		val, err := fetchThirdPartyStuffWhichCanBeSlow()

		respCh <- Response{
			res: val,
			err: err,
		}

	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching data from third party took too long")

		case resp := <-respCh:
			return resp.res, resp.err

		}
	}

}

func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(150 * time.Millisecond)
	return 200, nil
}
