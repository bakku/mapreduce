package mapreduce

import (
	"fmt"
	"log"
)

type Worker struct {
	id string
}

type JobScheduler struct {
	Workers []Worker
}

func (js JobScheduler) RegisterWorker(args struct{}, reply *string) error {
	id, err := GenerateRandomString(20)
	if err != nil {
		return fmt.Errorf("could not register worker: %w", err)
	}

	js.Workers = append(js.Workers, Worker{id})

	*reply = id

	return nil
}

type GetNextJobArgs struct {
	WorkerId string
}

type MapperArgs struct {
	Key   string
	Value string
}

type ReducerArgs struct {
	Key    string
	Values []string
}

type ScheduledJob struct {
	JobId        string
	MapperPhase  bool
	MapperArgs   MapperArgs
	ReducerPhase bool
	ReducerArgs  ReducerArgs
}

func (js JobScheduler) GetNextJob(args GetNextJobArgs, reply *ScheduledJob) error {
	log.Printf("worker %s is asking for next job\n", args.WorkerId)

	*reply = ScheduledJob{
		JobId:        "wordcount",
		MapperPhase:  true,
		MapperArgs:   MapperArgs{"file1", "Hello world!"},
		ReducerPhase: false,
		ReducerArgs:  ReducerArgs{},
	}

	log.Printf("scheduled job %s for worker %s\n", reply.JobId, args.WorkerId)

	return nil
}
