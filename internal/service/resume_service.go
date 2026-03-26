package sercvice 

import (
	"resume-maker-pai/internal/model"
)

type Response struct {
	Resume model.Resume
	Education model.Education
	Experience model.Experience
	Project model.Project
}

func CreateResume() Response{
	return Response
}