package driver

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

type FileDriver struct {
	dataFile     *os.File
	indexFile    *os.File
	currentIndex int
	currentUser  string
}

func (f *FileDriver) Logout() error {
	//TODO implement me
	// 删除记录当前用户的文件
	panic("implement me")
}

func (f *FileDriver) Login(username, password string) error {
	//TODO implement me
	// 创建当前登陆用户的文件
	panic("implement me")
}

func (f *FileDriver) Find(id int) (string, error) {
	if id < 0 {
		return "", nil
	}
	_, _ = f.dataFile.Seek(0, 0)
	reader := bufio.NewReader(f.dataFile)
	for true {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		if len(line) == 0 {
			continue
		}
		temp := &Data{}
		err = json.Unmarshal(line, temp)
		if err != nil {
			continue
		}
		if temp.Id == id {
			return temp.Data, nil
		}
	}
	return "", nil
}

func (f *FileDriver) Select() (map[int]string, error) {
	data := map[int]string{}
	_, _ = f.dataFile.Seek(0, 0)
	reader := bufio.NewReader(f.dataFile)
	for true {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		if len(line) == 0 {
			continue
		}
		temp := &Data{}
		err = json.Unmarshal(line, temp)
		if err != nil {
			continue
		}
		data[temp.Id] = temp.Data
	}
	return data, nil
}

func (f *FileDriver) Insert(data interface{}) error {
	f.currentIndex++
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	log.Println(string(b))
	f.writeData(f.currentIndex, string(b))
	return nil
}

func (f *FileDriver) Update(id int, data interface{}) error {
	if id < 0 {
		return nil
	}
	_, _ = f.dataFile.Seek(0, 0)
	reader := bufio.NewReader(f.dataFile)
	seekLine := 0
	for true {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		if len(line) == 0 {
			continue
		}
		temp := &Data{}
		err = json.Unmarshal(line, temp)
		if err == nil {
			if temp.Id == id {
				//此处更新数据
				data, _ := json.Marshal(data)
				temp.Data = string(data)
				tempJson, _ := json.Marshal(temp)
				log.Printf("修改后数据为:%s", string(tempJson))
				log.Printf("大小为:%d", seekLine)
				_, _ = f.dataFile.Seek(int64(seekLine), 0)
				_, err := f.dataFile.WriteString(fmt.Sprintf("%s\n", string(tempJson)))
				if err != nil {
					panic(err)
				}
				return nil
			}
		}
		seekLine += len(line) + 1
	}
	return errors.New("找不到数据")
}

func (f *FileDriver) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func (f *FileDriver) Free() error {
	defer func(dataFile *os.File) {
		_ = dataFile.Close()
	}(f.dataFile)
	defer func(indexFile *os.File) {
		_ = indexFile.Close()
	}(f.indexFile)
	_, _ = f.indexFile.Seek(0, 0)
	_, _ = f.indexFile.WriteString(fmt.Sprintf("%d", f.currentIndex))
	return nil
}

func (f *FileDriver) Init() error {
	//TODO 初始化登陆用户对内容
	panic("implement me")
}

func getFile(file string) (*os.File, error) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		// 创建文件
		_, err := os.Create(file)
		if err != nil {
			return nil, err
		}
	}
	f, err := os.OpenFile(file, os.O_RDWR, 0777)
	return f, err
}

func (f *FileDriver) writeData(id int, str string) {
	data, _ := json.Marshal(&Data{
		Id:   id,
		Data: str,
	})
	_, _ = f.dataFile.Seek(0, 2)
	_, err := f.dataFile.WriteString(fmt.Sprintf("%s\n", string(data)))
	if err != nil {
		log.Fatalf("写入数据失败 %s", err)
	}
}

func (f *FileDriver) BootStrap() {
	// 获取数据文件
	err := f.getUser()
	if err != nil {
		panic(err)
	}
	err = f.getData()
	if err != nil {
		panic(err)
	}
	err = f.getIndex()
	if err != nil {
		panic(err)
	}

}

func (f *FileDriver) getUser() error {
	// .todo-users
	f.currentUser = "default"
	// .todo-current-user
	return nil
}

func (f *FileDriver) getIndex() error {
	file, err := getFile(fmt.Sprintf(".%s-todo-index", f.currentUser))
	if err != nil {
		log.Fatalf("打开文件失败: %s", err)
		return err
	}
	f.indexFile = file
	_, _ = f.indexFile.Seek(0, 0)
	reader := bufio.NewReader(f.indexFile)
	readBuf, _ := reader.ReadBytes(10)
	if len(readBuf) == 0 {
		f.currentIndex = 0
	} else {
		index, err := strconv.Atoi(string(readBuf))
		log.Printf("读取到当前index %d", index)
		if err != nil {
			log.Fatalf("index转换失败: %s", err)
			return err
		}
		f.currentIndex = index
	}
	return nil
}

func (f *FileDriver) getData() error {
	file, err := getFile(fmt.Sprintf(".%s-todo-list", f.currentUser))
	if err != nil {
		return err
	}
	f.dataFile = file
	return nil
}

var fileDriver RepositoryDriver = &FileDriver{}

func NewFileDriver() (RepositoryDriver, error) {
	res := &FileDriver{}
	res.BootStrap()
	return res, nil
}

// uuid
