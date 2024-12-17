package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

var client *resty.Client

type Course struct {
	ID          string `json:"kcrwdm"`
	Name        string `json:"kcmc"`
	Description string `json:"xmmc"`
	Teacher     string `json:"teaxm"`
}

type CourseResponse struct {
	Total int      `json:"total"`
	Data  []Course `json:"rows"`
}

func GetCourses() ([]Course, error) {
	courseResponse := CourseResponse{}

	res, err := client.R().
		SetBody("page=1&rows=60&sort=kcrwdm&order=asc").
		Post("https://jxfw.gdut.edu.cn/xsxklist!getDataList.action")

	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	err = json.Unmarshal(res.Body(), &courseResponse)

	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	fmt.Println(courseResponse)

	return courseResponse.Data, nil
}

func main() {
	client = resty.New()

	var cookie string

	//输入cookie
	fmt.Println("请输入cookie(形如:JSESSIONID=xxxxx):")
	_, err := fmt.Scanln(&cookie)
	if err != nil {
		return
	}

	headers := map[string]string{
		"Host":             "jxfw.gdut.edu.cn",
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0",
		"Accept":           "application/json, text/javascript, */*; q=0.01",
		"Accept-Language":  "en-US,en;q=0.5",
		"Accept-Encoding":  "gzip, deflate, br",
		"Content-Type":     "application/x-www-form-urlencoded; charset=UTF-8",
		"X-Requested-With": "XMLHttpRequest",
		"Content-Length":   "36",
		"Origin":           "https://jxfw.gdut.edu.cn",
		"DNT":              "1",
		"Connection":       "keep-alive",
		"Referer":          "https://jxfw.gdut.edu.cn/xskjcjxx!kjcjList.action",
		"Cookie":           cookie,
		"Sec-Fetch-Dest":   "empty",
		"Sec-Fetch-Mode":   "cors",
		"Sec-Fetch-Site":   "same-origin",
	}

	//设置headers
	client.SetHeaders(headers)

	//获取课程信息
	fmt.Println("正在获取课程信息...")

	courses, err := GetCourses()
	if err != nil {
		fmt.Println("获取课程信息失败:", err)
		return
	}
	fmt.Println("获取课程信息成功")
	fmt.Println()

	for i, course := range courses {
		fmt.Println("编号:", i)
		fmt.Println("课程代码:", course.ID)
		fmt.Println("课程名称:", course.Name)
		fmt.Println("项目名称:", course.Description)
		fmt.Println("教师姓名:", course.Teacher)
		fmt.Println()
	}
	fmt.Println("请输入编号开始抢课:")

	//输入编号
	var index int
	_, err = fmt.Scanln(&index)

	if err != nil {
		fmt.Println("输入错误")
		return
	}

	//选课
	for {
		res, err := client.R().
			SetBody("kcrwdm=" + courses[index].ID + "&kcmc=" + courses[index].Name).
			Post("https://jxfw.gdut.edu.cn/xsxklist!getAdd.action")
		if err != nil {
			fmt.Println("选课失败:", err)
			return
		}
		fmt.Println(res.String())

		time.Sleep(time.Millisecond * 500)
	}
}
