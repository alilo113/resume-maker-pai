package model

type Resume struct {
    ID          string       `json:"id"` 
    FirstName   string       `json:"first_name"`
    LastName    string       `json:"last_name"`
    Email       string       `json:"email"`
    Phone       string       `json:"phone"`
    Address     string       `json:"address"`
    Summary     string       `json:"summary"`
    Education   []Education  `json:"education"`
    Experience  []Experience `json:"experience"`
    Skills      []string     `json:"skills"`
    Projects    []Project    `json:"projects"`
    Certifications []string  `json:"certifications,omitempty"`
    Links       []string     `json:"links,omitempty"`
}

type Education struct {
    School      string `json:"school"`
    Degree      string `json:"degree"`
    Field       string `json:"field"`
    StartYear   int    `json:"start_year"`
    EndYear     int    `json:"end_year"`
}

type Experience struct { 
    Company     string `json:"company"`
    Title       string `json:"title"`
    StartYear   int    `json:"start_year"`
    EndYear     int    `json:"end_year"`
    Description string `json:"description"`
}

type Project struct {
    Name        string `json:"name"`
    Description string `json:"description"`
    Link        string `json:"link,omitempty"`
}