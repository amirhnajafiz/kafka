package store

import (
	"context"

	"github.com/amirhnajafiz/personal-website/back-end/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	projectsCollection = "projects"
)

type ProjectsCollection struct {
	DB *mongo.Database
}

func (p *ProjectsCollection) Upsert(c context.Context, project *model.Project) error {
	opts := options.Update().SetUpsert(true)

	if _, err := p.DB.Collection(projectsCollection).UpdateOne(c, nil, bson.M{"$set": project}, opts); err != nil {
		return err
	}

	return nil
}

func (p *ProjectsCollection) GetAll(c context.Context, filter interface{}) ([]*model.Project, error) {
	var projects []*model.Project

	cur, err := p.DB.Collection(projectsCollection).Find(c, filter)
	if err != nil {
		return nil, err
	}

	var project *model.Project

	for cur.Next(c) {
		if err := cur.Decode(&project); err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	return projects, nil
}

func (p *ProjectsCollection) GetSingle(c context.Context, title string) (*model.Project, error) {
	filter := bson.M{
		"title": title,
	}

	result := p.DB.Collection(projectsCollection).FindOne(c, filter)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var project *model.Project

	if err := result.Decode(&project); err != nil {
		return nil, err
	}

	return project, nil
}

func (p *ProjectsCollection) Delete(c context.Context, title string) error {
	filter := bson.M{
		"title": title,
	}

	_, err := p.DB.Collection(projectsCollection).DeleteOne(c, filter)

	return err
}
