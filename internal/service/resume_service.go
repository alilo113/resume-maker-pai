package service

import "resume-maker-pai/internal/model" // Adjust path as needed

// Define the Response struct if you haven't already
type Response struct {
    Resume     model.Resume     `json:"resume"`
    Education  model.Education  `json:"education"`
    Experience model.Experience `json:"experience"`
    Project    model.Project    `json:"project"`
}

func CreateResume() Response {
    // 1. Create the detailed Resume
    res := model.Resume{
        ID:        "1",
        FirstName: "John",
        LastName:  "Doe",
        Email:     "john@example.com",
        Skills:    []string{"Go", "Python"},
    }

    // 2. Create the standalone samples
    edu := model.Education{
        School: "Tech University",
        Degree: "B.S. CS",
    }

    exp := model.Experience{
        Company: "Google",
        Title:   "Software Engineer",
    }

    proj := model.Project{
        Name:        "Resume Builder",
        Description: "A Go-based API",
    }

    // 3. Wrap them in the Response struct
    return Response{
        Resume:     res,
        Education:  edu,
        Experience: exp,
        Project:    proj,
    }
}