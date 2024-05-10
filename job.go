package mapreduce

type Job struct {
	Mapper  func(k, v string) map[string][]string
	Reducer func(k string, vs []string) string
}

var Jobs map[string]Job

func RegisterJob(jobId string, job Job) {
	if Jobs == nil {
		Jobs = make(map[string]Job)
	}

	Jobs[jobId] = job
}
