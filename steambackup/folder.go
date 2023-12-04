package steambackup

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyFolder(source, destination string) error {
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(destination, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		sourceFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer sourceFile.Close()

		destFile, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, sourceFile)
		return err
	})
}

func CopyFile(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("error opeoning source file: %v", err)
	}
	defer sourceFile.Close()
	
	destinationFile, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("error opeoning destination file: %v", err)
	}
	defer destinationFile.Close()
	
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return fmt.Errorf("error copy file contents: %v", err)
	}


	return nil

}

func DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("error deleting file : %v", err)
	}
	return nil
}

func ZipFolder(zipFileName, sourceFolder string) error {
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	err = filepath.Walk(sourceFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(sourceFolder, path)
		if err != nil {
			return err
		}

		zipPath := filepath.ToSlash(relPath)

		if info.IsDir() {
			zipPath += "/"
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = zipPath

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
		}

		return err
	})

	return err
}

func UnzipFolder(zipFileName, extractFolder string) error {
	zipReader, err := zip.OpenReader(zipFileName)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	err = os.MkdirAll(extractFolder, 0755)
	if err != nil {
		return err
	}

	for _, file := range zipReader.File {
		err := extractFile(file, extractFolder)
		if err != nil {
			return err
		}
	}

	return nil
}

func extractFile(file *zip.File, extractFolder string) error {
	zippedFile, err := file.Open()
	if err != nil {
		return err
	}
	defer zippedFile.Close()

	extractPath := filepath.Join(extractFolder, file.Name)
	if file.FileInfo().IsDir() {
		os.MkdirAll(extractPath, file.Mode())
	} else {
		os.MkdirAll(filepath.Dir(extractPath), file.Mode())

		extractFile, err := os.Create(extractPath)
		if err != nil {
			return err
		}
		defer extractFile.Close()

		_, err = io.Copy(extractFile, zippedFile)
		if err != nil {
			return err
		}
	}

	return nil
}
