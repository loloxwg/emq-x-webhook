package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"math/rand"
	"time"
)

func main() {
	//// You can generate a Token from the "Tokens Tab" in the UI
	//const token = "F-QFQpmCL9UkR3qyoXnLkzWj03s6m4eCvYgDl1ePfHBf9ph7yxaSgQ6WN0i9giNgRTfONwVMK1f977r_g71oNQ=="
	//const bucket = "users_business_events"
	//const org = "iot"
	//u1 := uuid.Must(uuid.NewV4(), nil)
	//fmt.Printf("UUIDv4: %s\n", u1)
	//
	//client := influxdb2.NewClient("http://124.222.47.219:8086", token)
	//// always close client at the end
	//defer client.Close()
	//
	//// get non-blocking write client
	//writeAPI := client.WriteAPI(org, bucket)
	//
	//// create point using fluent style
	//i := 0
	////start := time.Now()
	//for i < 10000 {
	//	i++
	//	fmt.Println(i)
	//
	//	//rand.Seed(time.Now().Unix())
	//	//p := influxdb2.NewPointWithMeasurement("devices").
	//	//	AddTag("client_id", u1.String()).
	//	//	AddField("temperature", rand.Float64()*50).
	//	//	AddField("humidity", rand.Intn(50)).
	//	//	SetTime(time.Now())
	//	//// write point asynchronously
	//	//writeAPI.WritePoint(p)
	//	//// Flush writes
	//	//writeAPI.Flush()
	//
	//}
	//
	////end := time.Now()
	////fmt.Printf("\npoints read %s\n", end.Sub(start))
	//fmt.Printf("\npoints read %s\n", "points read 3m21.77373ms")
	for {
		Writesingle()
	}

}

func Writesingle() {
	rand.Seed(time.Now().Unix())
	// You can generate a Token from the "Tokens Tab" in the UI
	const token = "F-QFQpmCL9UkR3qyoXnLkzWj03s6m4eCvYgDl1ePfHBf9ph7yxaSgQ6WN0i9giNgRTfONwVMK1f977r_g71oNQ=="
	const bucket = "users_business_events"
	const org = "iot"

	client := influxdb2.NewClient("http://124.222.47.219:8086", token)
	// always close client at the end
	defer client.Close()
	writeAPI := client.WriteAPI(org, bucket)

	//p := influxdb2.NewPoint("",
	//	map[string]string{"unit": "temperature"},
	//	map[string]interface{}{"avg": 24.5, "max": 45},
	//	time.Now())
	//// write point asynchronously
	//writeAPI.WritePoint(p)
	// create point using fluent style
	p := influxdb2.NewPointWithMeasurement("sensor").
		AddTag("unit", "th").
		//AddField("temperature", 22.6).
		//AddField("humidity", 43).
		AddField("temperature", rand.Float64()*50).
		AddField("humidity", rand.Intn(50)).
		SetTime(time.Now())
	// write point asynchronously
	writeAPI.WritePoint(p)
	// Flush writes
	writeAPI.Flush()
	fmt.Println(rand.Float64()*50, rand.Intn(50))
	fmt.Println("finished writing point")

	////////////////////////////////
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "124.222.47.219:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	pipe := rdb.TxPipeline()
	incr := pipe.Incr(ctx, "counter")
	_, exer := pipe.Exec(ctx)
	if exer != nil {
		fmt.Println(exer)
	}
	fmt.Println(incr.Val())
	pipe.Expire(ctx, "key", time.Hour)

	time.Sleep(10 * time.Second)
}
