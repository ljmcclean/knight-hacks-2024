package seeds

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/ljmcclean/knight-hacks-2024/services"
)

func SeedProjects(ctx context.Context, ps services.ProjectService) {
	projects := []*services.Project{
		{
			Name:        "Project 1",
			Description: "Some description for Project 1",
			IsRemote:    1,
			Location:    "Orlando, Florida",
			Skills:      []string{"programming", "art"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 2",
			Description: "Some description for Project 2",
			IsRemote:    0,
			Location:    "New York, New York",
			Skills:      []string{"design", "writing"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 3",
			Description: "Some description for Project 3",
			IsRemote:    1,
			Location:    "Los Angeles, California",
			Skills:      []string{"programming", "music"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 4",
			Description: "Some description for Project 4",
			IsRemote:    0,
			Location:    "Austin, Texas",
			Skills:      []string{"photography", "video"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 5",
			Description: "Some description for Project 5",
			IsRemote:    1,
			Location:    "Seattle, Washington",
			Skills:      []string{"programming", "data science"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 6",
			Description: "Some description for Project 6",
			IsRemote:    1,
			Location:    "Miami, Florida",
			Skills:      []string{"marketing", "sales"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 7",
			Description: "Some description for Project 7",
			IsRemote:    0,
			Location:    "Chicago, Illinois",
			Skills:      []string{"design", "photography"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 8",
			Description: "Some description for Project 8",
			IsRemote:    1,
			Location:    "Boston, Massachusetts",
			Skills:      []string{"writing", "programming"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 9",
			Description: "Some description for Project 9",
			IsRemote:    0,
			Location:    "Denver, Colorado",
			Skills:      []string{"art", "music"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 10",
			Description: "Some description for Project 10",
			IsRemote:    1,
			Location:    "Phoenix, Arizona",
			Skills:      []string{"data analysis", "business"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 11",
			Description: "Some description for Project 11",
			IsRemote:    0,
			Location:    "San Francisco, California",
			Skills:      []string{"programming", "design"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 12",
			Description: "Some description for Project 12",
			IsRemote:    1,
			Location:    "Atlanta, Georgia",
			Skills:      []string{"photography", "art"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 13",
			Description: "Some description for Project 13",
			IsRemote:    0,
			Location:    "Dallas, Texas",
			Skills:      []string{"programming", "music"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 14",
			Description: "Some description for Project 14",
			IsRemote:    1,
			Location:    "Philadelphia, Pennsylvania",
			Skills:      []string{"business", "marketing"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 15",
			Description: "Some description for Project 15",
			IsRemote:    0,
			Location:    "Portland, Oregon",
			Skills:      []string{"video", "art"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 16",
			Description: "Some description for Project 16",
			IsRemote:    1,
			Location:    "San Diego, California",
			Skills:      []string{"programming", "data analysis"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 17",
			Description: "Some description for Project 17",
			IsRemote:    0,
			Location:    "Las Vegas, Nevada",
			Skills:      []string{"marketing", "sales"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 18",
			Description: "Some description for Project 18",
			IsRemote:    1,
			Location:    "New Orleans, Louisiana",
			Skills:      []string{"design", "business"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 19",
			Description: "Some description for Project 19",
			IsRemote:    0,
			Location:    "Nashville, Tennessee",
			Skills:      []string{"music", "art"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 20",
			Description: "Some description for Project 20",
			IsRemote:    1,
			Location:    "Charlotte, North Carolina",
			Skills:      []string{"programming", "photography"},
			UserID:      uuid.New(),
		},
	}

	for _, project := range projects {
		if err := ps.PostProject(ctx, project); err != nil {
			log.Printf("error seeding project %s: %s", project.Name, err)
		}
	}
}
