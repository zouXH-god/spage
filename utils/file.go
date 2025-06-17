package utils

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
)

// IsValidZipFile 检查 multipart.FileHeader 是否为合法的 ZIP 文件
func IsValidZipFile(fileHeader *multipart.FileHeader) (bool, error) {
	// 打开上传的文件
	file, err := fileHeader.Open()
	if err != nil {
		return false, err
	}
	defer file.Close()

	// 读取文件的前 512 字节用于检测（ZIP 文件头通常在文件开头）
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		return false, err
	}

	// 检查文件签名是否是 ZIP 文件
	// ZIP 文件签名通常是 "PK\x03\x04" 或其他变体
	if !isZipFileSignature(buffer) {
		return false, nil
	}

	// 重置文件读取位置，因为上面已经读取了一部分
	if _, err = file.Seek(0, io.SeekStart); err != nil {
		return false, err
	}

	// 使用 archive/zip 包尝试读取 ZIP 文件
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return false, err
	}

	reader := bytes.NewReader(fileBytes)
	zipReader, err := zip.NewReader(reader, fileHeader.Size)
	if err != nil {
		return false, nil // 不是有效的 ZIP 文件
	}

	// 检查 ZIP 文件中是否至少有一个文件
	if len(zipReader.File) == 0 {
		return false, nil
	}

	return true, nil
}

// isZipFileSignature 检查字节切片是否包含 ZIP 文件签名
func isZipFileSignature(data []byte) bool {
	// ZIP 文件签名可以是以下几种:
	// - PK\x03\x04 (最常见的 ZIP 文件)
	// - PK\x05\x06 (空归档)
	// - PK\x07\x08 (分卷归档)
	if len(data) < 4 {
		return false
	}

	return data[0] == 'P' && data[1] == 'K' &&
		(data[2] == 0x03 && data[3] == 0x04 || // 常规文件
			data[2] == 0x05 && data[3] == 0x06 || // 空归档
			data[2] == 0x07 && data[3] == 0x08) // 分卷归档
}

// FileHash 计算文件的哈希值并返回十六进制字符串
func FileHash(filePath string) (string, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 创建哈希计算器
	hash := sha256.New()

	// 将文件内容拷贝到哈希计算器
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	// 计算哈希值并转换为十六进制字符串
	hashInBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)

	return hashString, nil
}

func FileHashFromStream(file multipart.File) (string, error) {
	// 创建哈希计算器
	hash := sha256.New()

	// 将文件流内容拷贝到哈希计算器
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	// 计算哈希值并转换为十六进制字符串
	hashInBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)

	return hashString, nil
}

// FilePath 根据哈希值生成文件路径，前4位为目录位hash[0:4]/hash
func FilePath(hash string) (dir, file string) {
	dir = hash[0:4]
	file = hash
	return
}
