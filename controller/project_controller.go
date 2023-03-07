package controller

import (
	"github.com/kataras/iris/v12"
	"myweb/model"
	"myweb/service"
	"strconv"
)

type ProjectController struct{}

// GetProjectByid 根据ID查询
func (pc *ProjectController) GetProjectByid(ctx iris.Context) {
	projectID, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64) //将字符串解析成整数，10进制 int64
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"status":  "0",
			"success": "查询失败",
			"message": "请求参数错误",
		})
		return
	}
	projectService := &service.ProjectService{}
	project, err := projectService.GetById(projectID)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"status":  "0",
			"success": "查询失败",
			"message": err.Error(),
		})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{
		"status":  "1",
		"success": "查询成功",
		"message": "获取测试项目成功",
		"project": project,
	})
}

// CreateProject 创建项目
func (pc *ProjectController) CreateProject(ctx iris.Context) {
	var project model.Project
	if err := ctx.ReadJSON(&project); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"status":  "0",
			"success": "创建项目失败",
			"message": "请求格式错误",
		})
		return
	}
	projectService := &service.ProjectService{}
	if err := projectService.CreateProject(&project); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"status":  "0",
			"success": "创建项目失败",
			"message": err.Error(),
		})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{
		"status":  "1",
		"success": "创建项目成功",
		"message": "项目创建成功",
	})
}

// 删除项目
func (pc *ProjectController) DeleteProjectByid(ctx iris.Context) {
	projectID, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64) //将字符串解析成整数，10进制 int64
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"status":  "0",
			"success": "删除失败",
			"message": "请求参数错误",
		})
		return
	}
	projectService := &service.ProjectService{}
	if err := projectService.DeleteProject(projectID); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"status":  "0",
			"success": "删除失败",
			"message": err.Error(),
		})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{
		"status":  "1",
		"success": "删除成功",
		"message": "删除成功",
	})
}

// 获取所有项目
func (pc *ProjectController) GetAllProject(ctx iris.Context) {
	projectService := &service.ProjectService{}
	projects, err := projectService.GetAllProject()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"status":  "0",
			"success": "列表查询失败",
			"message": err.Error(),
		})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{
		"status":   "1",
		"success":  "查询成功",
		"message":  "查询成功",
		"projects": projects,
	})
}
