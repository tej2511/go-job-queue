package job

import "context"

type Job struct {
	ID 			string
	Execute 	func(ctx context.Context) error
	Attempts 	int
	MaxRetries 	int
}