package service

import (
	"errors"
	"myweb/datasource"
	"myweb/model"
)

type TestcaseService struct{}

// 获取所有测试用例
func (ts *TestcaseService) GetAllTestcase() ([]model.TesetCase, error) {
	testcases := make([]model.TesetCase, 0)
	engine, err := datasource.GetEngine()
	if err != nil {
		return nil, err
	}
	err = engine.Find(&testcases)
	return testcases, err
}

//通过ID查询用例

func (ts *TestcaseService) GetTestCaseByID(id int64) (*model.TesetCase, error) {
	testcases := &model.TesetCase{}
	engine, err := datasource.GetEngine()
	if err != nil {
		return nil, err
	}
	has, err := engine.Id(id).Get(testcases)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未查询到该项目")
	}
	return testcases, nil
}

// 创建测试用例
func (ts *TestcaseService) CreateTestCase(testcase *model.TesetCase) error {

	engine, err := datasource.GetEngine()
	if err != nil {
		return err
	}
	_, err = engine.Insert(testcase)
	return err
}

// 通过ID删除测试用例
func (ts *TestcaseService) DeleteTestCaseByID(id int64) error {
	engine, err := datasource.GetEngine()
	if err != nil {
		return err
	}
	testcases := &model.TesetCase{ID: id}
	_, err = engine.Id(id).Delete(testcases)
	return err
}
