package store

import (
	"context"

	"github.com/amirhnajafiz/personal-website/back-end/internal/model"
)

const (
	projectsCollection = "projects"
)

type ProjectsCollection struct{}

func (p *ProjectsCollection) Upsert(c context.Context, project *model.Project) error {
	return nil
}

func (p *ProjectsCollection) GetAll(c context.Context) []*model.Project {
	return nil
}

func (p *ProjectsCollection) GetAllAvailable(c context.Context) []*model.Project {
	return nil
}

func (p *ProjectsCollection) GetSingle(c context.Context) *model.Project {
	return nil
}
