package registry

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"provsim-prototype/core"

	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

//LocalFileSystemStorageProvider ...
type LocalFileSystemStorageProvider struct {
	FilePath string
}

//Fetch Configfiles from the filesystem and load the configs
//TODO: is this an optimized way to send the map as return value.
func (S LocalFileSystemStorageProvider) Fetch() (map[uuid.UUID]core.Provider, error) {
	b := time.Now()

	path, err := os.Getwd()
	if err != nil {
		log.Printf("Cant find the working directory, failed with following error %v", err)
		return nil, err
	}
	log.Printf("The working directory is %s", path)

	files, err := readDir(path + S.FilePath)
	if nil != err {
		log.Fatalf("while reading the data folder: %v", err)
		return nil, err
	}
	log.Printf("Number of files in the list %d", len(files))

	configs := make(map[uuid.UUID]core.Provider)
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".yaml") {
			continue
		}

		var pr core.Provider
		d, err := ioutil.ReadFile(path + S.FilePath + file.Name())
		if nil != err {
			log.Fatalf("while reading the data folder: %v", err)
			return nil, err
		}

		yaml.Unmarshal([]byte(d), &pr)
		configs[pr.ID] = pr
	}

	e := time.Now()
	elapsed := e.Sub(b)

	fmt.Printf("started at %s ended at %s, so the elapsed is %d to load number of %d configs\n",
		b, e, elapsed/1000000, len(configs))

	return configs, nil
}

// ReadDir reads the directory named by dirname and returns
// a list of directory entries sorted by filename.
func readDir(dirname string) ([]os.FileInfo, error) {
	log.Printf("Trying to read files from the director %s", dirname)

	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}

	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}

	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, nil
}
