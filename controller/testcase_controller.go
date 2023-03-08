package controller

import (
	"github.com/kataras/iris/v12"
	"myweb/model"
	"myweb/service"
	"strconv"
)

type TestcaseController struct{}

//通过ID查询

func (tc *TestcaseController) GetCasesById(ctx iris.Context) {
	CaseID, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"status":  "0",
			"success": "查询失败",
			"message": "请求参数错误",
		})
		return
	}
	testcaseService := &service.TestcaseService{}
	testcase, err := testcaseService.GetTestCaseByID(CaseID)
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
		"status":   "1",
		"success":  "查询成功",
		"testcase": testcase,
	})

}

// 创建测试用例
func (tc *TestcaseController) CreateCase(ctx iris.Context) {
	var testcase model.TesetCase
	if err := ctx.ReadJSON(&testcase); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"status":  "0",
			"success": "创建项目失败",
			"message": "请求格式错误",
		})
		return
	}
	testcaseService := &service.TestcaseService{}
	if err := testcaseService.CreateTestCase(&testcase); err != nil {
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
		"success": "测试成功",
		"message": "项目创建成功",
	})
}

// 删除测试用例
func (tc *TestcaseController) DeleteCase(ctx iris.Context) {
	testcaseID, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"status":  "0",
			"success": "删除项目失败",
			"message": "请求格式错误",
		})
		return
	}
	testcaseService := &service.TestcaseService{}
	if err := testcaseService.DeleteTestCaseByID(testcaseID); err != nil {
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
		"success": "测试成功",
		"message": "测试成功",
	})
}

// 获取所有测试用例
func (tc *TestcaseController) GetAllCase(ctx iris.Context) {
	testcaseService := &service.TestcaseService{}
	testcase, err := testcaseService.GetAllTestcase()
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
		"status":   "1",
		"success":  "查询成功",
		"message":  "查询成功",
		"testcase": testcase,
	})
}
