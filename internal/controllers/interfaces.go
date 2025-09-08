package controllers

import "log"

type Services struct {
	AssignmentCtrl *assignmentCtrl
}

func NewServices(logger *log.Logger) Services {
	ac := newAssignmentCtrl(logger)

	return Services{
		AssignmentCtrl: ac,
	}
}
