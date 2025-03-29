package upload

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func FileInUploadFolderWithCustomPath(file *multipart.FileHeader, customPath *string, fileName *string, oldFileName *string) (string, error) {
	path, err := definePath("app/views/assets/upload", customPath, fileName)
	if err != nil {
		return "" , err
	}

	if oldFileName != nil {
		if err := replaceIfFileExist(&path); err != nil {
			return "", err
		}
	}

	if err := saveUploadedFile(file, path); err != nil {
		return "", err
	}

	return path, nil
}

func definePath(basePath string, customPath *string, fileName *string) (string, error) {
	dynamicPath := filepath.Join(basePath, *customPath)

	if err := os.MkdirAll(dynamicPath, os.ModePerm); err != nil {
		return "", err
	}

	return filepath.Join(dynamicPath, *fileName), nil
}

func replaceIfFileExist(path *string) error {
	if _, err := os.Stat(*path); err == nil {
		if err := os.Remove(*path); err != nil {
			return err
		}
	}

	return nil
}

func saveUploadedFile(file *multipart.FileHeader, fullPath string) error {
	// ouvre le fichier qu'on reçois de la request
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// creer le fichier à l'endroit qu'on veut
	newFile, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	// ecris le contenu du fichier qu'on reçois dans la request dans le nouveau fichier
	_, err = io.Copy(newFile, src)
	return err
}