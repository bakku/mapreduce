package main

import (
	"flag"
	"log"
	"mapreduce"
	"net/rpc"
)

var masterHost = flag.String("master", "127.0.0.1:1234", "the url to the master RPC handler")

func main() {
	client, err := rpc.DialHTTP("tcp", *masterHost)
	if err != nil {
		log.Fatalln("error connecting to master host:", err)
	}

	var workerId *string
	err = client.Call("JobScheduler.RegisterWorker", struct{}{}, &workerId)
	if err != nil {
		log.Fatalln("error registering via RPC:", err)
	}

	log.Println("got worker id from master:", *workerId)

	var result *mapreduce.ScheduledJob

	err = client.Call("JobScheduler.GetNextJob", mapreduce.GetNextJobArgs{WorkerId: *workerId}, &result)
	if err != nil {
		log.Fatalln("error running RPC Add:", err)
	}

	HandleJob(*result)
}

func HandleJob(scheduledJob mapreduce.ScheduledJob) {
	jobDefinition, ok := mapreduce.Jobs[scheduledJob.JobId]
	if !ok {
		log.Printf("scheduledJob %s is not registered\n", scheduledJob.JobId)
		return
	}

	if scheduledJob.MapperPhase {
		result := jobDefinition.Mapper(scheduledJob.MapperArgs.Key, scheduledJob.MapperArgs.Value)
		log.Printf("result of mapping: %v\n", result)
	} else if scheduledJob.ReducerPhase {
		result := jobDefinition.Reducer(scheduledJob.ReducerArgs.Key, scheduledJob.ReducerArgs.Values)
		log.Printf("result of reducer: %v\n", result)
	} else {
		log.Printf("no active phase set in scheduled job\n")
	}
}
