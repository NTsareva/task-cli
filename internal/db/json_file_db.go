package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"syscall"
	"task-cli/internal/models"
)

const FILE_PATH = "./tasks-cli.json"
const BACKUP_FILE_PATH = "./backup-tasks-cli.json"

type JsonDB struct {
	TasksInFile []models.Task
	TasksMap    map[int]models.Task
	MaxTaskID   int
}

func readTaskFile() ([]models.Task, error) {
	var data []models.Task

	file, err := os.Open(FILE_PATH) // TODO: create file copy in any case

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []models.Task{}, nil
		}

		return nil, fmt.Errorf("open file %s: %w", FILE_PATH, err)
	}

	err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX) //in case of opening another instance of app
	if err != nil {
		panic("Could not open file: " + FILE_PATH)
	}

	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("read file: %s", err)
	}

	return data, nil
}

func createOriginalFileBackup() error {
	body, err := os.ReadFile(FILE_PATH)
	if err != nil {
		return fmt.Errorf("read from original file: %s\n", err)
	}

	err = os.WriteFile(BACKUP_FILE_PATH, body, 0644)
	if err != nil {
		return fmt.Errorf("create backup file: %s\n", err)
	}

	return nil
}

func writeChangesToFile(jsb JsonDB) error {
	file, err := os.OpenFile(FILE_PATH, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("open file file %s: %w", FILE_PATH, err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("	", " ")

	if err := encoder.Encode(jsb.TasksInFile); err != nil {
		return fmt.Errorf("write changes to file: %s", FILE_PATH)
	}

	return nil
}

func NewJsonDB() (*JsonDB, error) {
	var tasksMap = make(map[int]models.Task)
	fileJsonBody, err := readTaskFile()
	if err != nil {
		return nil, err
	}

	maxID := 0

	for i, task := range fileJsonBody {
		tasksMap[task.ID] = fileJsonBody[i]

		if task.ID > maxID {
			maxID = task.ID
		}
	}

	return &JsonDB{
		TasksInFile: fileJsonBody,
		TasksMap:    tasksMap,
		MaxTaskID:   maxID,
	}, nil
}

func (jdb *JsonDB) CreateTask(task models.Task) (*models.Task, error) {
	newTask := models.Task{
		ID:     jdb.MaxTaskID + 1,
		Name:   task.Name,
		Status: models.TODO,
	}

	jdb.TasksInFile = append(jdb.TasksInFile, newTask)

	err := createOriginalFileBackup()
	if err != nil {
		return nil, err
	}

	err = writeChangesToFile(*jdb)
	if err != nil {
		return nil, err
	}

	return &newTask, nil
}

//file, err := os.OpenFile(FILE_PATH, os.o_|os.O_CREATE, 0600)
