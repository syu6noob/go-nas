package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

type ApiType struct {
	Info     string `json:"info,omitempty"`
	Open     string `json:"open,omitempty"`
	Download string `json:"download,omitempty"`
}

type StatType struct {
	Name  string  `json:"name"`
	Size  int64   `json:"size"`
	Path  string  `json:"path"`
	Mime  string  `json:"mime"`
	IsDir bool    `json:"isDir"`
	Api   ApiType `json:"api"`
}

type InfoType struct {
	IsRoot     bool        `json:"isRoot"`
	Name       string      `json:"name"`
	Size       int64       `json:"size"`
	Path       string      `json:"path"`
	Mime       string      `json:"mime"`
	IsDir      bool        `json:"isDir"`
	Api        ApiType     `json:"api"`
	Link       string      `json:"link"`
	ParentLink string      `json:"parent_link"`
	Parent     string      `json:"parent"`
	Children   *[]StatType `json:"children,omitempty"`
}

func GetContentDir() (string, error) {
	contentsFolder := os.Getenv("CONTENTS_FOLDER")
	current, err := os.Getwd()
	if err != nil {
		return "", err
	} else {
		return path.Join(current, contentsFolder), nil
	}
}

func getApiLink(t string, arg string) string {
	host := os.Getenv("API_HOST")
	hostUrl, _ := url.Parse(host)

	linkUrl := *hostUrl
	if t != "" {
		linkUrl.Path = path.Join(hostUrl.Path, arg, t)
	}
	linkUrl.Path = path.Clean(linkUrl.Path)

	return linkUrl.String()
}

func getMime(file *os.File) string {
	buffer := make([]byte, 512)
	_, _ = file.Read(buffer)
	file.Seek(0, io.SeekStart)

	mime := http.DetectContentType(buffer)
	return mime
}

func GetRawTarget(t string) (string, error) {
	contentDir, contentDirErr := GetContentDir()
	if contentDirErr != nil {
		return "", contentDirErr
	}
	target := filepath.Join(contentDir, t)

	absTarget, err := filepath.Abs(target)
	if err != nil {
		return "", err
	}
	absContentDir, err := filepath.Abs(contentDir)
	if err != nil {
		return "", err
	}
	if !filepath.HasPrefix(absTarget, absContentDir) {
		return "", fmt.Errorf("invalid path access")
	}
	return absTarget, nil
}

func GetStat(t string) (*StatType, error) {
	rt, rtErr := GetRawTarget(t)
	if rtErr != nil {
		return nil, rtErr
	}

	info, err := os.Lstat(rt)
	if err != nil {
		return nil, err
	}

	var mime string
	if !info.IsDir() {
		file, err := os.Open(rt)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		mime = getMime(file)
	}

	var stat *StatType
	if info.IsDir() {
		stat = &StatType{
			Name:  info.Name(),
			Size:  info.Size(),
			IsDir: true,
			Path:  t,
			Mime:  "",
			Api: ApiType{
				Info: getApiLink(t, "info"),
			},
		}
	} else {
		stat = &StatType{
			Name:  info.Name(),
			Size:  info.Size(),
			IsDir: false,
			Path:  t,
			Mime:  mime,
			Api: ApiType{
				Info:     getApiLink(t, "info"),
				Open:     getApiLink(t, "open"),
				Download: getApiLink(t, "raw"),
			},
		}
	}
	return stat, nil
}

func GetDirList(t string) (*[]StatType, error) {
	rt, rtErr := GetRawTarget(t)
	if rtErr != nil {
		return nil, rtErr
	}

	files, filesErr := os.ReadDir(rt)
	if filesErr != nil {
		return nil, filesErr
	}

	var dirList []StatType

	for _, file := range files {
		target := path.Join(t, file.Name())

		// fmt.Printf("%s\n", target)

		stat, statErr := GetStat(target)
		if statErr != nil {
			return nil, statErr
		}
		dirList = append(dirList, *stat)
	}

	return &dirList, nil
}

func GetInfo(t string) (*InfoType, error) {
	t = path.Clean(t)

	parentDir := path.Clean(path.Dir(t))

	// link := getAppLink(t)
	// parentLink := getAppLink(parentDir)

	stat, statErr := GetStat(t)
	if statErr != nil {
		return nil, statErr
	}

	var children *[]StatType
	var childrenErr error

	var info InfoType

	if stat.IsDir {
		children, childrenErr = GetDirList(t)
		if childrenErr != nil {
			return nil, childrenErr
		}

		if t == "/" {
			info = InfoType{
				IsRoot: true,
				Name:   "",
				Size:   stat.Size,
				Path:   t,
				IsDir:  stat.IsDir,
				Parent: parentDir,
				Api: ApiType{
					Info: getApiLink(t, "info"),
				},
				Children: children,
			}
		} else {
			info = InfoType{
				IsRoot: false,
				Name:   stat.Name,
				Size:   stat.Size,
				Path:   t,
				IsDir:  stat.IsDir,
				Parent: parentDir,
				Api: ApiType{
					Info: getApiLink(t, "info"),
				},
				Children: children,
			}
		}
	} else {
		info = InfoType{
			IsRoot: false,
			Name:   stat.Name,
			Size:   stat.Size,
			Mime:   stat.Mime,
			IsDir:  stat.IsDir,
			Parent: path.Dir(t),
			Api: ApiType{
				Info:     getApiLink(t, "info"),
				Open:     getApiLink(t, "open"),
				Download: getApiLink(t, "raw"),
			},
			Children: children,
		}
	}

	return &info, nil
}
