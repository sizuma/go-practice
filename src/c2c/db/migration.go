package db

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

type byFileName struct {
	files []os.FileInfo
	asc   bool
}

func (st *byFileName) Len() int {
	return len(st.files)
}

func (st *byFileName) Less(i, j int) bool {
	iName := st.files[i].Name()
	jName := st.files[j].Name()
	iStringNumber := strings.Split(iName, "_")[0]
	jStringNumber := strings.Split(jName, "_")[0]
	iNumber, _ := strconv.Atoi(iStringNumber)
	jNumber, _ := strconv.Atoi(jStringNumber)
	result := iNumber > jNumber
	if st.asc {
		result = !result
	}
	return result
}

func (st *byFileName) Swap(i, j int) {
	temp := st.files[i]
	st.files[i] = st.files[j]
	st.files[j] = temp
}

// Migration migrate db
func Migration(ups, drops bool) error {
	db, connectError := Connect()
	if connectError != nil {
		return connectError
	}
	defer db.Close()
	_, filename, _, _ := runtime.Caller(0)
	pwd := path.Join(path.Dir(filename))
	files, readDirError := ioutil.ReadDir(filepath.Join(pwd, "..", "resource", "migration"))
	if readDirError != nil {
		return readDirError
	}
	upRegex := regexp.MustCompile(`.*\.up\.sql`)
	downRegex := regexp.MustCompile(`.*\.down\.sql`)
	sort.Sort(&byFileName{
		files: files[:],
		asc:   true,
	})

	if drops {
		for _, file := range files {
			name := file.Name()
			if downRegex.MatchString(name) {
				path := filepath.Join(pwd, "..", "resource", "migration/", name)
				file, fileReadError := ioutil.ReadFile(path)
				if fileReadError != nil {
					return fileReadError
				}
				content := string(file[:])
				_, migrationError := db.Exec(content)
				if migrationError != nil {
					if strings.Index(migrationError.Error(), "Unknown table") != -1 {
						log.Println("undo unknown table(ignored)")
					} else {
						return migrationError
					}
				} else {
					log.Printf("exec down migration : %s", content)
				}
			}
		}
	}

	if ups {
		for index := len(files) - 1; index >= 0; index-- {
			file := files[index]
			name := file.Name()
			if upRegex.MatchString(name) {
				path := filepath.Join(pwd, "..", "resource", "migration/", name)
				file, fileReadError := ioutil.ReadFile(path)
				if fileReadError != nil {
					return fileReadError
				}
				content := string(file[:])
				_, migrationError := db.Exec(content)
				if migrationError != nil {
					if strings.Index(migrationError.Error(), "Table") != -1 &&
						strings.Index(migrationError.Error(), "already exists") != -1 {
						log.Println("migrate already exists table(ignored)")
					} else {
						return migrationError
					}
				} else {
					log.Printf("exec up migration : %s", content)
				}
			}
		}
	}
	return nil
}
