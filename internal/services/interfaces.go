package services

import "log"

type Services struct {
	AssignmentService *assignmentService
}

func NewServices(logger *log.Logger) Services {
	as := newAssignmentService(logger)

	return Services{
		AssignmentService: as,
	}
}
