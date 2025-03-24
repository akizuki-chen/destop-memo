package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) ReadData() string {
	filePath, err := getPath()
	if err != nil {
		return ""
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return ""
	}
	return string(data)
}

func (a *App) SaveData(str string) string {
	filePath, err := getPath()
	if err != nil {
		return err.Error()
	}
	err = os.WriteFile(filePath, []byte(str), 0644)
	if err != nil {
		return err.Error()
	}
	return ""
}

func getPath() (string, error) {
	var dataPath string
	currentOs := runtime.GOOS
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("无法获取用户主目录: %v", err)
	}
	switch currentOs {
	case "windows":
		dataPath = filepath.Join(userHomeDir, "memo.json")
	case "darwin", "linux":
		dataPath = filepath.Join(userHomeDir, "memo.json")
	default:
		return "", fmt.Errorf("不支持的系统")
	}
	return dataPath, nil
}
