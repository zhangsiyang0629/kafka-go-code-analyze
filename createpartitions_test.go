package kafka

import (
	"context"
	"testing"

	ktesting "github.com/segmentio/kafka-go/testing"
)

func TestClientCreatePartitions(t *testing.T) {
	if !ktesting.KafkaIsAtLeast("1.0.1") {
		return
	}

	client, shutdown := newLocalClient()
	defer shutdown()

	topic := makeTopic()
	createTopic(t, topic, 1)
	defer deleteTopic(t, topic)

	res, err := client.CreatePartitions(context.Background(), &CreatePartitionsRequest{
		Topics: []TopicPartitionsConfig{
			{
				Name:  topic,
				Count: 2,
				TopicPartitionAssignments: []TopicPartitionAssignment{
					{
						BrokerIDs: []int32{1},
					},
				},
			},
		},
		ValidateOnly: false,
	})
	if err != nil {
		t.Fatal(err)
	}

	if err := res.Errors[topic]; err != nil {
		t.Error(err)
	}
}

func TestClientCreatePartitionsNoAssignments(t *testing.T) {
	if !ktesting.KafkaIsAtLeast("1.0.1") {
		return
	}

	client, shutdown := newLocalClient()
	defer shutdown()

	topic := makeTopic()
	createTopic(t, topic, 1)
	defer deleteTopic(t, topic)

	res, err := client.CreatePartitions(context.Background(), &CreatePartitionsRequest{
		Topics: []TopicPartitionsConfig{
			{
				Name:  topic,
				Count: 2,
			},
		},
		ValidateOnly: false,
	})
	if err != nil {
		t.Fatal(err)
	}

	if err := res.Errors[topic]; err != nil {
		t.Error(err)
	}
}
