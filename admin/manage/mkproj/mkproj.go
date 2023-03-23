package mkproj

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type t_manager struct {
	dir string
}

func mkProjDirErr(title, backpath string, err error) {
	if err1 := os.Chdir(backpath); err1 != nil {
		log.Fatalln(err, err1)
	}
	if err1 := os.RemoveAll(title); err1 != nil {
		log.Fatalln(err, err1)
	}
}

func mkProjSetting(title string, perm fs.FileMode) {
	if err := os.Mkdir("settings", perm); err != nil {
		mkProjDirErr(title, "..", err)
	}
}

func mkProjNote(title string, perm fs.FileMode) {
	if err := os.Mkdir("notes", perm); err != nil {
		mkProjDirErr(title, "..", err)
	}
}

func mkProjMain(title string, perm fs.FileMode) {
	if err := os.Mkdir("main", perm); err != nil {
		mkProjDirErr(title, "..", err)
	}
	if err := os.Chdir("main"); err != nil {
		mkProjDirErr(title, "..", err)
	}
	if err := ioutil.WriteFile("README.md", []byte("# "+title+"\n"), perm); err != nil {
		mkProjDirErr(title, filepath.Join("..", ".."), err)
	}
	if err := os.Chdir(".."); err != nil {
		mkProjDirErr(title, filepath.Join("..", ".."), err)
	}
}

func MkProjDir(title string, perm fs.FileMode) {
	if err := os.Mkdir(title, perm); err != nil {
		log.Fatalln(err)
	}
	if err := os.Chdir(title); err != nil {
		if err1 := os.RemoveAll(title); err1 != nil {
			log.Fatalln(err, err1)
		}
	}
	mkProjMain(title, perm)
	mkProjNote(title, perm)
	mkProjSetting(title, perm)
}

func IsProjDir(path string) bool {
	if _, err := os.Stat(filepath.Join(path, "settings")); err != nil {
		return false
	}
	if _, err := os.Stat(filepath.Join(path, "notes")); err != nil {
		return false
	}
	if _, err := os.Stat(filepath.Join(path, "main")); err != nil {
		return false
	}
	if _, err := os.Stat(filepath.Join(path, "main", "README.md")); err != nil {
		return false
	}
	return true
}

func isDir(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	return file.IsDir
}

func copyDirAll()

func CopyDirAll(frompath string, topath string) error {
	if isDir(frompath) && isDir(topath) {
		return errors.New("Either frompath or topath are not directory path.")
	}
	title := filepath.Abs(frompath)
	
	files, err := ioutil.ReadDir()
}

func RecognizeProj(path, title string, manage t_manager) error {
	if IsProjDir(path) {
		return errors.New("the directory cannot recognize as the project")
	}
	return os.(path, manage.dir)
}
