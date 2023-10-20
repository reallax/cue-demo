package main

import (
	"os"
	"fmt"
	"time"
	"strings"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/protobuf/proto"

	"reallx/cue-demo/cmd/kafka/kk"
)

func main() {
	// prod:
	brokerServer := "10.27.9.202:9092"
	topic := "dmg_metric"

	// dev:
	// brokerServer := "10.16.244.71:8092"
	// topic := "dmc_alarm"


	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    brokerServer,
		"group.id":             "tmp-hpa-group",
		"auto.offset.reset":    "smallest"})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%% Error1: %v\n", err)
	}
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%% Error2: %v\n", err)
	}

	loop:
	for {
		ev := consumer.Poll(1000)
		switch e := ev.(type) {
		case *kafka.Message:
			// application-specific processing
			key, inst, val, flg := analysisSinkMessageMQ(e.Key, e.Value)
			if key == "dmall-open/dmall-open-trace-logger" {
				fmt.Printf("%% Message on %s:\n%s\n%d\n%v\n",
				key, inst, val, flg)
			}
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error3: %v\n", e)
			break loop
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}
	
	consumer.Close()
}


func analysisSinkMessageMQ(k, v []byte) (key string, instance string, value int, valid bool) {
	point := &kk.Point{}
	err := proto.Unmarshal(v, point)
	if err != nil {
		fmt.Printf("Unmarshal Sink.Message failed: %s", err)
		return "", "", 0, false
	}

	projectCode := point.Labels["projectCode"]
	appCode := point.Labels["appCode"]
	topic := point.Labels["topic"]
	broker := point.Labels["broker"]
	partition := point.Labels["partition"]
	group := point.Labels["group"]

	if topic == "kb_open_tracelogger_req" {
		fmt.Printf("get message content: %+v. focus.\n",
			point,
		)
	}

	if point.Table != "mq_consumer_group_metric" {
		// fmt.Printf("Table should be %s instead of %s, content: %+v. Skip it.",
		// 	"mq_consumer_group_metric",
		// 	point.Table,
		// 	point,
		// )
		return "", "", 0, false
	}

	pointExpiredTime := time.Minute * 5
	expired := isExpired(point.Timestamp, pointExpiredTime)
	if expired {
		fmt.Printf("MQ point data's time %s is expired over %s, content: %+v. Skip it.",
			point.Timestamp,
			pointExpiredTime,
			point,
		)
		return "", "", 0, false
	}

	metricValue := int(point.Values["lag"])
	if metricValue < 0 {
		return "", "", 0, false
	}

	modified := int64(point.Values["modified"] / 1000)
	diffHours := int((time.Now().Unix() - modified) / 3600)
	if diffHours >= 24 {
		fmt.Printf("MQ point data's modified time %dh is expired over %dh, content: %+v. Skip it.",
			diffHours,
			24,
			point,
		)
		return "", "", 0, false
	}

	key, err = generateKey(projectCode, appCode)
	if err != nil {
		fmt.Printf("generateKey err: %+v, projectCode: %s, appCode: %s", err, projectCode, appCode)
		return "", "", 0, false
	}

	instance, err = generateKey(topic, broker, partition, group)
	if err != nil {
		fmt.Printf("generateInstance err: %+v, topic: %s, broker: %s, partition: %s, group: %s", err, topic, broker, partition, group)
		return "", "", 0, false
	}

	return key, instance, metricValue, true
}

func isExpired(timestamp string, d time.Duration) bool {
	i, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		fmt.Printf("isExpired parse timestamp err: %v", err)
		return false
	}
	tm := time.Unix(i, 0)

	return time.Now().Add(-d).After(tm)
}

func generateKey(keys ...string) (string, error) {
	for i, key := range keys {
		// java null string
		if key == "" || key == "null" {
			return "", fmt.Errorf("generateKey args %d key is empty", i)
		}
	}

	return strings.Join(keys, "/"), nil
}