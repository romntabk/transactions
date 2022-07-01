package worker


import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"database/sql"
	"fmt"
)

func validate(command string, name string, amount int) bool {
	if amount <= 0 {
		return false
	}
	if command != "and" && command != "withdraw"{
		return false
	}
	if nameExists(name) {
		return true
	}
	return false
}

func main() {
	db = GetDb()
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
	     "bootstrap.servers":    "kafka_server:9092",
	     "group.id":             "group_id"
	 })
	
	for run == true {
		// get from kafka first unreaded message from each topic/user with name command amount 
	    ev := consumer.Poll(0)
	    switch e := ev.(type) {
	    case kafka.AssignedPartitions:
	        fmt.Fprintf(os.Stderr, "%% %v\n", e)
	        c.Assign(e.Partitions)
	    case kafka.RevokedPartitions:
	        fmt.Fprintf(os.Stderr, "%% %v\n", e)
	        c.Unassign()
	    case *kafka.Message:

	    	// e.Value
	    	// e.TopicPartition
	    	// get name
	    	// get command 
	    	// get amount
	    	if validate(command, name, amount) != true {
	    		continue
	    	}
	    	if command == "add" {
	    		db.Exec(fmt.Sprintf("update account set balance = balance + %s where name = %s", amount, name))# add to db module
	    		consumer.Commit()
	    		continue
	    	}
	    	actual_amount := GetAmount(db, name)
	    	if actual_amount < amount{
	    		continue
	    	}
	    	db.Exec(fmt.Sprintf("update account set balance = balance - %s where name = %s", amount, name)) # add to db module
	        consumer.Commit()
	    case kafka.PartitionEOF:
	        fmt.Printf("%% Reached %v\n", e)
	    case kafka.Error:
	        fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
	        run = false 
	    default:
	        fmt.Printf("Ignored %v\n", e)
	    }
	}
}

