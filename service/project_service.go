package service

import (
	"errors"
	"myweb/datasource"
	"myweb/model"
)

/*
项目管理service层
*/

type ProjectService struct{}

// CreateProject 创建项目
func (ps *ProjectService) CreateProject(project *model.Project) error {
	engine, err := datasource.GetEngine()
	if err != nil {
		return err
	}
	_, err = engine.Insert(project)
	return err
}

// GetById 根据ID获取项目
func (ps *ProjectService) GetById(id int64) (*model.Project, error) {
	project := &model.Project{}
	engine, err := datasource.GetEngine()
	if err != nil {
		return nil, err
	}
	has, err := engine.Id(id).Get(project)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未查询到该项目")
	}
	return project, nil
}

// DeleteProject 通过ID删除项目
func (ps *ProjectService) DeleteProject(id int64) error {
	engine, err := datasource.GetEngine()
	if err != nil {
		return err
	}
	project := &model.Project{ID: id}
	_, err = engine.Delete(project)
	return err
}

// GetAllProject 获取所有项目
func (ps *ProjectService) GetAllProject() ([]model.Project, error) {
	projects := make([]model.Project, 0)
	engine, err := datasource.GetEngine()
	if err != nil {
		return nil, err
	}
	err = engine.Find(&projects)
	return projects, err
}
