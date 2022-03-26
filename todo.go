package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"todo/driver"
)

// Todo App
type Todo struct {
	defaultDriver driver.RepositoryDriver
	config        *Config
}

// TodoListItem Json
type TodoListItem struct {
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
	Status  int    `json:"status"`
}

// NewTodo 创建Todo
func NewTodo(config *Config) *Todo {
	return &Todo{
		config: config,
	}
}

// CreateTodo 创建
func (t *Todo) CreateTodo(str string) error {
	d, err := t.GetDriver()
	if err != nil {
		return err
	}
	item := &TodoListItem{
		Content: str,
		UserId:  0,
		Status:  InCompleteStatus,
	}
	return d.Insert(item)
}

func (t *Todo) ChangeTodoStatus(id int, status int) error {
	d, err := t.GetDriver()
	if err != nil {
		return err
	}
	itemJson, err := d.Find(id)
	if err != nil {
		return errors.New("找不到目标数据")
	}
	item := &TodoListItem{}
	err = json.Unmarshal([]byte(itemJson), item)
	if err != nil {
		return errors.New("数据存在异常")
	}
	item.Status = status
	return d.Update(id, item)
}

func (t *Todo) DeleteTodo(id int) error {
	d, err := t.GetDriver()
	if err != nil {
		return err
	}
	return d.Delete(id)
}

func (t *Todo) SelectByStatus(status ...int) ([]string, error) {
	d, err := t.GetDriver()
	if err != nil {
		return nil, err
	}
	targetStatus := -1
	for i := range status {
		targetStatus = status[i]
		break
	}

	list, _ := d.Select()
	res := make([]string, 0, 20)
	for i := range list {
		temp := &TodoListItem{}
		err := json.Unmarshal([]byte(list[i]), temp)
		if err != nil {
			continue
		}
		if targetStatus == -1 || temp.Status == targetStatus {
			statusLabel := "未完成"
			if temp.Status == CompleteStatus {
				statusLabel = "已完成"
			}
			res = append(res, fmt.Sprintf("\t\t%d\t|\t%s\t|\t%s\t", i, temp.Content, statusLabel))
		}
	}
	return res, nil
}

// GetDriver 获取驱动
func (t *Todo) GetDriver(driver ...string) (driver.RepositoryDriver, error) {
	if len(driver) == 0 {
		var err error
		if t.defaultDriver == nil {
			t.defaultDriver, err = t.NewDriver(t.config.Default)
		}
		return t.defaultDriver, err
	}
	for i := range driver {
		return t.NewDriver(driver[i])
	}
	return nil, nil
}

// NewDriver 创建驱动
func (t *Todo) NewDriver(driverName string) (driver.RepositoryDriver, error) {
	if driverName == "" {
		driverName = "file"
	}
	switch driverName {
	case "file":
		return driver.NewFileDriver()
	case "database":
		return driver.NewDatabaseDriver()
	default:
		return nil, errors.New(fmt.Sprintf("不存在的driver [%s]", driverName))
	}
}

func (t *Todo) Done() error {
	d, err := t.GetDriver()
	if err != nil {
		log.Fatalf("获取驱动失败:%s", err)
		return err
	}
	d.Free()
	return nil
}
