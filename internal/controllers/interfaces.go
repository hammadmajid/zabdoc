package controllers

import "log"

type Services struct {
	AssignmentCtrl *assignmentCtrl
	LabTaskCtrl    *labTaskCtrl
}

func NewServices(logger *log.Logger) Services {
	ac := newAssignmentCtrl(logger)
	ltc := newLabTaskCtrl(logger)

	return Services{
		AssignmentCtrl: ac,
		LabTaskCtrl:    ltc,
	}
}
