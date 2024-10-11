package files

import "os"

// func FileExists(filePath string) bool {
// 	_, err := os.Stat(filePath)
// 	return !os.IsNotExist(err)
// }

func FileExists(filePath string) bool {
    info, err := os.Stat(filePath)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}