package utils

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	EXT_HTML = ".html"
	EXT_XML  = ".xml"
	EXT_JSON = ".json"
	EXT_ZIP  = ".zip"
	EXT_IDOC = ".idoc"
	EXT_ZED  = ".zed"
	EXT_PDF  = ".pdf"
)

/*GetFilelist
golang 文件处理公共函数库
https://www.cnblogs.com/chengbiwei/p/9672896.html
*/
func GetFilelist(path string, pattern string) []string {
	fileList := make([]string, 0, 1000)
	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if ok, _ := filepath.Match(pattern, f.Name()); ok {
			fileList = append(fileList, path)
		}
		return nil
	})
	return fileList
}

func ForeachFile(path string, pattern string, onFindFile func(fileName string) error) error {
	return filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if ok, _ := filepath.Match(pattern, f.Name()); ok {
			err := onFindFile(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func ReadFileContext(path string) string {
	fi, err := os.Open(path)
	NoneError(err)
	defer fi.Close()

	fd, err := ioutil.ReadAll(fi)
	NoneError(err)
	return string(fd)
}

func SaveFile(file string, content string) {
	f := ForceCreateFile(file)
	defer f.Close()

	_, err := f.WriteString(content)
	NoneError(err)
}

func SaveUtf8BomFile(file string, content string) {
	fs := ForceCreateFile(file)
	defer fs.Close()

	fs.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	fs.WriteString(content)
}

func ForceCreateFile(filename string) *os.File {
	dir := filepath.Dir(filename)
	dir = UNCPath(dir)
	err := os.MkdirAll(dir, 0755)
	ShowError(err)

	filename = UNCPath(filename)
	fs, err := os.Create(filename)
	ShowError(err)
	return fs
}

func SetLastModified(file string, time time.Time) {
	os.Chtimes(file, time, time)
}

func GetlastModified(file string) time.Time {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return time.Time{}
	}
	return fileInfo.ModTime()
}

func CopyFile(dst, src string) error {
	if !FileExists(src) {
		return errors.New("file not exist `" + src + "`")
	}
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out := ForceCreateFile(dst)
	if err != nil {
		fmt.Println("copy ", src)
		fmt.Println("to   ", dst)
		fmt.Println("Failed")
		return nil
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	cerr := out.Close()
	if err != nil {
		return err
	}
	return cerr
}

func CopyDir(dst, src string) error {
	dst = UNCPath(dst)
	src = UNCPath(src)
	return filepath.Walk(src,
		func(path string, f os.FileInfo, err error) error {
			if err != nil {
				NoneError(err)
			}
			if !f.IsDir() {
				rel, err := filepath.Rel(src, path)
				NoneError(err)
				return CopyFile(filepath.Join(dst, rel), path)
			}
			return nil
		})
}

func PathExists(path string) bool {
	finfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return finfo.IsDir()
}

func FileExists(filename string) bool {
	finfo, err := os.Stat(filename)
	if err == nil && !finfo.IsDir() {
		return true
	}
	return false
}

func WriteFile(file string, data []byte) {
	fs := ForceCreateFile(file)
	defer fs.Close()
	fs.Write(data)
}

func ReadJson2Slice(filename string) (records []map[string]interface{}, err error) {
	// Read json buffer from jsonFile
	byteValue, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	// We have known the outer json object is a map, so we define  result as map.
	// otherwise, result could be defined as slice if outer is an array
	var obj map[string]interface{}
	err = json.Unmarshal(byteValue, &obj)
	if err != nil {
		return
	}
	for _, v := range obj {
		records = append(records, v.(map[string]interface{}))
	}
	return
}

func ReadExcelsheet2Slice(sheet *xlsx.Sheet, dataStartRow int) (records []map[string]interface{}, err error) {
	// We have known the outer json object is a map, so we define  result as map.
	// otherwise, result could be defined as slice if outer is an array
	var obj map[string]interface{}
	rowsNum := len(sheet.Rows) //- dataStartRow + 2
	for rowIndex := dataStartRow - 1; rowIndex < rowsNum; rowIndex++ {

	}

	for _, v := range obj {
		records = append(records, v.(map[string]interface{}))
	}
	return
}

func SaveJson(filename string, obj interface{}) error {
	fs := ForceCreateFile(filename)
	defer fs.Close()

	data, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}

	_, err = fs.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func ReadXml(obj interface{}, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(data, obj)
	if err != nil {
		return err
	}
	return nil
}
func SaveXml(filename string, obj interface{}) error {
	fs := ForceCreateFile(filename)
	defer fs.Close()

	data, err := xml.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}

	_, err = fs.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func GetFileInfo(filename string) (name, ext string, en bool) {
	file := filepath.Base(filename)
	ext = filepath.Ext(file)

	name = strings.TrimSuffix(file, ext)

	if strings.ToLower(filepath.Ext(name)) == ".en" {
		en = true
		ext = filepath.Ext(name) + ext
		name = strings.TrimSuffix(file, ext)
	}
	return name, ext, en
}

func FileToUrlPath(file string) string {
	xUrl, err := url.Parse(file)
	NoneError(err)
	return xUrl.EscapedPath()
}

func UrlToLocalFile(Url string) string {
	xUrl, err := url.Parse(Url)
	NoneError(err)
	return SHA1EncodeFileName(xUrl.Path)
}

func SHA1EncodeFileName(filename string) string {
	//for en picture: xxx.en.png

	dir := filepath.Dir(filename)

	filenameOnly, ext, _ := GetFileInfo(filename)

	t := sha1.New()
	_, err := io.WriteString(t, filenameOnly)
	NoneError(err)

	return dir + "/" + fmt.Sprintf("%x", t.Sum(nil)) + ext
}

func RemoveFile(file string) {
	if !FileExists(file) {
		return
	}
	err := os.Remove(file)
	if err != nil {
		log.Println(file, " file remove Error!\n")
		log.Printf("%s", err)
	} else {
		log.Print(file, " file remove OK!\n")
	}
}

func FileOpen(file_name string) ([]byte, error) {
	var (
		open               *os.File
		file_data          []byte
		open_err, read_err error
	)
	open, open_err = os.Open(file_name)
	if open_err != nil {
		return nil, open_err
	}
	defer open.Close()
	file_data, read_err = ioutil.ReadAll(open)
	if read_err != nil {
		return nil, read_err
	}
	return file_data, nil
}

func CurDir() string {
	_, file, _, _ := runtime.Caller(1)
	return filepath.Dir(file)
}

func IsFileSame(left, right string) bool {
	sFile, err := os.Open(left)
	if err != nil {
		return false
	}
	defer sFile.Close()

	dFile, err := os.Open(right)
	if err != nil {
		return false
	}
	defer dFile.Close()

	sBuffer := make([]byte, 512)
	dBuffer := make([]byte, 512)
	for {
		_, sErr := sFile.Read(sBuffer)
		_, dErr := dFile.Read(dBuffer)
		if sErr != nil || dErr != nil {
			if sErr != dErr {
				return false
			}
			if sErr == io.EOF {
				break
			}
		}
		if bytes.Equal(sBuffer, dBuffer) {
			continue
		}
		return false
	}
	return true
}

func DiffFile(left, right string) (diff string, same bool, err error) {
	_, err = exec.LookPath("diff")
	if err != nil {
		return "", false, err
	}

	cmdName, _ := exec.LookPath("diff")
	cmd := exec.Command(cmdName, left, right)
	output, errOfCmd := cmd.CombinedOutput()
	if errOfCmd != nil {
		return string(output), false, nil
	}
	return "", true, nil
}

// Pattern to match a windows absolute path: "c:\" and similar
var isAbsWinDrive = regexp.MustCompile(`^[a-zA-Z]\:\\`)

// UNCPath converts an absolute Windows path
// to a UNC long path.
//https://blog.klauspost.com/long-windows-paths-unc-paths-in-go/
//\\?\UNC\server\share\dir\file.txt
func UNCPath(s string) string {
	// UNC can NOT use "/", so convert all to "\"
	s = strings.Replace(s, `/`, `\`, -1)

	// If prefix is "\\", we already have a UNC path or server.
	if strings.HasPrefix(s, `\\`) {

		// If already long path, just keep it
		if strings.HasPrefix(s, `\\?\`) {
			return s
		}

		// Trim "\\" from path and add UNC prefix.
		return "\\\\?\\UNC\\" + strings.TrimPrefix(s, `\\`)
	}

	if isAbsWinDrive.MatchString(s) {
		return `\\?\` + s
	}

	return s
}

// 使用Value.Kind()类型判断，并获取真实值进行打印
func GetStringValue(i interface{}) (result string) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Int:
		result = strconv.Itoa(int(v.Int()))
	case reflect.Float64:
		result = strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.String:
		result = v.String()
	case reflect.Bool:
		result = strconv.FormatBool(v.Bool())
	default:
		result = v.String()
	}
	return
}
