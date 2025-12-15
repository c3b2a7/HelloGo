package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// FindMissingFiles 找出dir1中不存在但dir2中存在的文件
//
// 参数:
//
//	dir1: 基准目录
//	dir2: 比较目录
//
// 返回:
//
//	[]string: dir1中不存在但dir2中存在的文件列表（相对于dir2的路径）
//	error: 错误信息
func FindMissingFiles(dir1, dir2 string) ([]string, error) {
	// 检查目录是否存在
	if !dirExists(dir1) {
		return nil, fmt.Errorf("目录 %s 不存在", dir1)
	}
	if !dirExists(dir2) {
		return nil, fmt.Errorf("目录 %s 不存在", dir2)
	}

	// 获取两个目录中的所有文件
	files1, err := getAllFiles(dir1)
	if err != nil {
		return nil, fmt.Errorf("无法读取目录 %s: %v", dir1, err)
	}

	files2, err := getAllFiles(dir2)
	if err != nil {
		return nil, fmt.Errorf("无法读取目录 %s: %v", dir2, err)
	}

	// 找出dir2中存在但dir1中不存在的文件
	var missing []string
	for file := range files2 {
		if !files1[file] {
			missing = append(missing, file)
		}
	}

	return missing, nil
}

// FindMissingFilesWithFullPath 找出dir1中不存在但dir2中存在的文件（返回完整路径）
//
// 参数:
//
//	dir1: 基准目录
//	dir2: 比较目录
//
// 返回:
//
//	[]string: dir1中不存在但dir2中存在的文件列表（完整路径）
//	error: 错误信息
func FindMissingFilesWithFullPath(dir1, dir2 string) ([]string, error) {
	missingFiles, err := FindMissingFiles(dir1, dir2)
	if err != nil {
		return nil, err
	}

	// 转换为完整路径
	var fullPaths []string
	for _, file := range missingFiles {
		fullPath := filepath.Join(dir2, file)
		fullPaths = append(fullPaths, filepath.ToSlash(fullPath))
	}

	return fullPaths, nil
}

// dirExists 检查目录是否存在
func dirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// getAllFiles 递归获取目录中的所有文件（相对路径）
func getAllFiles(dir string) (map[string]bool, error) {
	files := make(map[string]bool)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 只处理文件，忽略目录
		if !info.IsDir() {
			// 获取相对于根目录的路径
			relPath, err := filepath.Rel(dir, path)
			if err != nil {
				return err
			}
			// 统一使用正斜杠作为路径分隔符
			relPath = filepath.ToSlash(relPath)
			files[relPath] = true
		}

		return nil
	})

	return files, err
}

// 示例用法
func main() {
	errors.Join()
	if len(os.Args) != 3 {
		fmt.Println("用法: go run filecmp.go <dir1> <dir2>")
		fmt.Println("功能: 找出dir1中不存在但dir2中存在的文件")
		return
	}

	dir1 := os.Args[1]
	dir2 := os.Args[2]

	// 使用相对路径版本
	missingFiles, err := FindMissingFiles(dir1, dir2)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}

	sort.Strings(missingFiles)

	if len(missingFiles) == 0 {
		fmt.Println("没有找到dir1中不存在但dir2中存在的文件")
	} else {
		fmt.Printf("dir1(%s)中不存在但dir2(%s)中存在的文件:\n", dir1, dir2)
		for _, file := range missingFiles {
			fmt.Printf("  %s\n", file)
		}
		fmt.Printf("\n总计: %d 个文件\n", len(missingFiles))
	}
}
