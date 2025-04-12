package upload

import (
	"hestia/internal/logger"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// ╔═══════════════════════════════════════════════════════════╗
// ║                       FONCTIONS PRIVATE                   ║
// ╚═══════════════════════════════════════════════════════════╝

// ╔═══════════════════════════════════════════════════════════╗
// ║                       FONCTIONS PUBLIC                    ║
// ╚═══════════════════════════════════════════════════════════╝

/*
Vérifie si le fichier existe à l'emplacement donné.

Params :
	- path : string not null
Return :
	- bool
	- error
*/
func Exist(path string) (bool, error) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		logger.Warnf("[upload:Exist] Fichier %v n'existe pas", path)
		return false, nil
	} 
	
	if err != nil {
		logger.Errorf("[upload:Exist] Erreur existance du fichier %v", path)
		return false, err
	}

	return true, nil
}

/*
Prépare le chemin du fichier en créant les répertoires nécessaires.

Params :
	- basePath : string not null
	- customPath : string not null
	- fileName : string not null
return :
	- string
	- erreur
	- create folder /tmp/upload/2023/01/01 if not exist
	- return ex: /tmp/upload/2023/01/01/test.txt
*/
func PrepareFilePath(basePath string, customPath string, fileName string) (string, error) {
	pathFinal := filepath.Join(basePath, customPath)
	if customPath == "" {
		pathFinal = basePath
	}

	if err := os.MkdirAll(pathFinal, os.ModePerm); err != nil {
		logger.Warnf("[upload:PrepareFilePath] Erreur lors de la création du dossier %v", pathFinal)
		return "", err
	}

	return filepath.Join(pathFinal, fileName), nil
}

/*
Sauvegarde ou écrase un fichier à l'emplacement donné.

Params :
	- file : pointeur file not null
	- pathFinal : string not null
Return :
	- error
*/
func SaveOrOverwrite(file *multipart.FileHeader, pathFinal string) error {
	src, err := file.Open()
	if err != nil {
		logger.Warnf("[upload:SaveOrOverwrite] Erreur lors de l'ouverture du fichier %v : %v", pathFinal, err)
		return err
	}
	defer src.Close()

	newFile, err := os.Create(pathFinal)
	if err != nil {
		logger.Warnf("[upload:SaveOrOverwrite] Erreur lors de la création du fichier %v : %v", pathFinal, err)
		return err
	}
	defer newFile.Close()

	if _, err := io.Copy(newFile, src); err != nil {
		logger.Warnf("[upload:SaveOrOverwrite] Erreur à l'écriture du fichier dans %v : %v", pathFinal, err)
		return err
	}

	return nil
}

/*
Supprime in préfix du chemin donné.

Params :
	- path : string not null
	- prefix : string not null
Return :
	- string
	- return ex: /tmp/upload/2023/01/01/test.txt => 2023/01/01/test.txt
*/
func RemovePrefixPath(path string, prefix string) string {
	if len(path) >= len(prefix) && (path)[:len(prefix)] == prefix {
		return path[len(prefix):]
	}
	return path
}

/*
Supprime le fichier situé au chemin donné.  
  
Params :
	- path : string not null
Return :  
	- error
*/
func DeleteFile(path string) error {
	return os.Remove(path)
}