package mapreduce

// This file basically just contains job definitions which will be immediately
// defined through the init() call.

import (
	"fmt"
	"strings"
)

func init() {
	RegisterJob("wordcount", Job{
		Mapper: func(k, v string) map[string][]string {
			result := make(map[string][]string)
			words := strings.Split(v, " ")

			for _, word := range words {
				slice, ok := result[word]
				if !ok {
					result[word] = []string{"1"}
				}

				slice = append(slice, "1")
			}

			return result
		},
		Reducer: func(k string, vs []string) string {
			return fmt.Sprintf("%d", len(vs))
		},
	})
}
