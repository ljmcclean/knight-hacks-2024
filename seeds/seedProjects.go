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
			IsRemote:    true,
			Location:    "Orlando, Florida",
			Skills:      []string{"programming", "art"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 2",
			Description: "Some description for Project 2",
			IsRemote:    false,
			Location:    "New York, New York",
			Skills:      []string{"design", "writing"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 3",
			Description: "Some description for Project 3",
			IsRemote:    true,
			Location:    "Los Angeles, California",
			Skills:      []string{"programming", "music"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 4",
			Description: "Some description for Project 4",
			IsRemote:    false,
			Location:    "Austin, Texas",
			Skills:      []string{"photography", "video"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 5",
			Description: "Some description for Project 5",
			IsRemote:    true,
			Location:    "Seattle, Washington",
			Skills:      []string{"programming", "data science"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 6",
			Description: "Some description for Project 6",
			IsRemote:    true,
			Location:    "Miami, Florida",
			Skills:      []string{"marketing", "sales"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 7",
			Description: "Some description for Project 7",
			IsRemote:    false,
			Location:    "Chicago, Illinois",
			Skills:      []string{"design", "photography"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 8",
			Description: "Some description for Project 8",
			IsRemote:    true,
			Location:    "Boston, Massachusetts",
			Skills:      []string{"writing", "programming"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 9",
			Description: "Some description for Project 9",
			IsRemote:    false,
			Location:    "Denver, Colorado",
			Skills:      []string{"art", "music"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 10",
			Description: "Some description for Project 10",
			IsRemote:    true,
			Location:    "Phoenix, Arizona",
			Skills:      []string{"data analysis", "business"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 11",
			Description: "Some description for Project 11",
			IsRemote:    false,
			Location:    "San Francisco, California",
			Skills:      []string{"programming", "design"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 12",
			Description: "Some description for Project 12",
			IsRemote:    true,
			Location:    "Atlanta, Georgia",
			Skills:      []string{"photography", "art"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 13",
			Description: "Some description for Project 13",
			IsRemote:    false,
			Location:    "Dallas, Texas",
			Skills:      []string{"programming", "music"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 14",
			Description: "Some description for Project 14",
			IsRemote:    true,
			Location:    "Philadelphia, Pennsylvania",
			Skills:      []string{"business", "marketing"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 15",
			Description: "Some description for Project 15",
			IsRemote:    false,
			Location:    "Portland, Oregon",
			Skills:      []string{"video", "art"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 16",
			Description: "Some description for Project 16",
			IsRemote:    true,
			Location:    "San Diego, California",
			Skills:      []string{"programming", "data analysis"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 17",
			Description: "Some description for Project 17",
			IsRemote:    false,
			Location:    "Las Vegas, Nevada",
			Skills:      []string{"marketing", "sales"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 18",
			Description: "Some description for Project 18",
			IsRemote:    true,
			Location:    "New Orleans, Louisiana",
			Skills:      []string{"design", "business"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 19",
			Description: "Some description for Project 19",
			IsRemote:    false,
			Location:    "Nashville, Tennessee",
			Skills:      []string{"music", "art"},
			UserID:      uuid.New(),
		},
		{
			Name:        "Project 20",
			Description: "Some description for Project 20",
			IsRemote:    true,
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
