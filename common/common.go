/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package common

import (
	"crypto/md5"
	"crypto/rand"
	sha1_ "crypto/sha1"
	sha256_ "crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	mathrand "math/rand"
	"net/url"
	"os"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/bwmarrin/snowflake"

	"github.com/ca17/teamsacs/common/log"
)

var (
	EmptyList = []interface{}{}
)

const (
	NA       = "N/A"
	ENABLED  = "enabled"
	DISABLED = "disabled"
)


func init() {
}

// print usage
func Usage(str string) {
	fmt.Fprintf(os.Stderr, str)
	flag.PrintDefaults()
}

// 创建目录
func MakeDir(path string) {
	f, err := os.Stat(path)
	if err != nil || f.IsDir() == false {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			log.Error("create dir fail！", err)
			return
		}
	}
}


func FileExists(file string) bool {
	info, err := os.Stat(file)
	return err == nil && !info.IsDir()
}

func DirExists(file string) bool {
	info, err := os.Stat(file)
	return err == nil && info.IsDir()
}

// panic error
func Must(err error) {
	if err != nil {
		panic(errors.WithStack(err))
	}
}

// panic error
func MustDebug(err error, debug bool) {
	if err != nil {
		if debug {
			panic(errors.WithStack(err))
		}else{
			panic(err)
		}
	}
}

func MustCallBefore(err error, callbefore func()) {
	if err != nil {
		callbefore()
		panic(errors.WithStack(err))
	}
}


func Must2(v interface{}, err error) interface{} {
	Must(err)
	return v
}

func IgnoreError(v interface{}, err error) interface{} {
	return v
}


func UUID() string {
	unix32bits := uint32(time.Now().UTC().Unix())
	buff := make([]byte, 12)
	numRead, err := rand.Read(buff)
	if numRead != len(buff) || err != nil {
		Must(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x-%x", unix32bits, buff[0:2], buff[2:4], buff[4:6], buff[6:8], buff[8:])
}

var snowflakeNode, _ = snowflake.NewNode(int64(mathrand.Intn(1000)))

// Generate int64
func UUIDint64() int64 {
	return snowflakeNode.Generate().Int64()
}

func UUIDBase32() (string, error) {
	id := snowflakeNode.Generate()
	// Print out the ID in a few different ways.
	return id.Base32(), nil
}


// Convert to Big Hump format
func ToCamelCase(str string) string {
	temp := strings.Split(str, "_")
	for i, r := range temp {
		temp[i] = strings.Title(r)
	}
	return strings.Join(temp, "")
}

var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// Convert to underlined format
func ToSnakeCase(str string) string {
	snake := matchAllCap.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(snake)
}


func Sha1Hash(src string) string {
	h := sha1_.New()
	h.Write([]byte(src))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func Sha256Hash(src string) string {
	h := sha256_.New()
	h.Write([]byte(src))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}


func Sha256HashWithSalt(src string, salt string) string {
	h := sha256_.New()
	h.Write([]byte(src))
	h.Write([]byte(salt))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}

// Determine if the string is in the list.
func InSlice(v string, sl []string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func IfEmpty(src interface{}, defval interface{}) interface{} {
	if IsEmpty(src) {
		return defval
	}
	return src
}

func IfNA(src string, defval string) string {
	if src == "N/A" || src == "" {
		return defval
	}
	return src
}

func EmptyToNA(src string) string {
	if strings.TrimSpace(src) == "" {
		return "N/A"
	}
	return src
}

func IfEmptyStr(src string, defval string) string {
	if src == "" {
		return defval
	}
	return src
}

// IsEmpty checks if a value is empty or not.
// A value is considered empty if
// - integer, float: zero
// - bool: false
// - string, array: len() == 0
// - slice, map: nil or len() == 0
// - interface, pointer: nil or the referenced value is empty
func IsEmpty(value interface{}) bool {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String, reflect.Array, reflect.Map, reflect.Slice:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Invalid:
		return true
	case reflect.Interface, reflect.Ptr:
		if v.IsNil() {
			return true
		}
		return IsEmpty(v.Elem().Interface())
	case reflect.Struct:
		v, ok := value.(time.Time)
		if ok && v.IsZero() {
			return true
		}
	}

	return false
}

func IsNotEmpty(value interface{}) bool {
	return !IsEmpty(value)
}

func split(s string, size int) []string {
	ss := make([]string, 0, len(s)/size+1)
	for len(s) > 0 {
		if len(s) < size {
			size = len(s)
		}
		ss, s = append(ss, s[:size]), s[size:]

	}
	return ss
}

func File2Base64(file string) string {
	data := Must2(ioutil.ReadFile(file))
	return base64.StdEncoding.EncodeToString(data.([]byte))
}

func Base642file(b64str string, file string) error {
	data, err := base64.StdEncoding.DecodeString(b64str)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, data, 777)
}

func GetCaptchaFont(dir string) string {
	var fpath = path.Join(dir, "comic.ttf")
	if !FileExists(fpath) {
		MakeDir(dir)
		_ = Base642file(Fontb64, fpath)
	}
	return fpath
}

func parseWithLocation(name string, timeStr string) (time.Time, error) {
	locationName := name
	if l, err := time.LoadLocation(locationName); err != nil {
		println(err.Error())
		return time.Time{}, err
	} else {
		lt, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, l)
		fmt.Println(locationName, lt)
		return lt, nil
	}
}


var mobileRe, _ = regexp.Compile("(?i:Mobile|iPod|iPhone|Android|Opera Mini|BlackBerry|webOS|UCWEB|Blazer|PSP)")

func MobileAgent(userAgent string) string {
	return mobileRe.FindString(userAgent)
}

// 校验码生成
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	mathrand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[mathrand.Intn(r)])
	}
	return sb.String()
}

func SetEmptyStrToNA(t interface{}) {
	d := reflect.TypeOf(t).Elem()
	for j := 0; j < d.NumField(); j++ {
		ctype := d.Field(j).Type.String()
		if ctype == "string" {
			val := reflect.ValueOf(t).Elem().Field(j)
			if val.String() == "" {
				val.SetString(NA)
			}
		}
	}
}

func IsEmptyOrNA(val string) bool {
	return val == "" || val == NA
}

func IsNotEmptyAndNA(val string) bool {
	return val != "" && val != NA
}

func Md5HashFile(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}

func UrlJoin(hurl string, elm ...string) string {
	u, err := url.Parse(hurl)
	Must(err)
	u.Path = path.Join(u.Path, path.Join(elm...))
	return u.String()
}



var notfloat = errors.New("not float value")

func ParseFloat64(v interface{}) (float64,error) {
	switch v.(type) {
	case float64:
		return v.(float64),nil
	case int64:
		return float64(v.(int64)),nil
	case int:
		return float64(v.(int)),nil
	case string:
		fv, err := strconv.ParseFloat(v.(string), 64)
		if err != nil {
			return 0, err
		}
		return fv, nil
	}
	return 0, notfloat
}

var notint = errors.New("not int value")

func ParseInt64(v interface{}) (int64,error) {
	switch v.(type) {
	case float64:
		return int64(v.(float64)),nil
	case int64:
		return v.(int64), nil
	case int:
		return int64(v.(int)),nil
	case string:
		ival, err := strconv.ParseInt(v.(string), 10,64)
		if err != nil {
			return 0, err
		}
		return ival, nil
	}
	return 0, notint
}

func ParseString(v interface{}) (string,error) {
	switch v.(type) {
	case float64:
		return strconv.FormatFloat(v.(float64), 'f', 2, 64),nil
	case int64:
		return strconv.FormatInt(v.(int64), 10), nil
	case int:
		return strconv.Itoa(v.(int)),nil
	case string:
		return v.(string), nil
	case nil:
		return "", nil
	case time.Time:
		return v.(time.Time).Format("2006-01-02 15:04:05"), nil
	default:
		return fmt.Sprintf("%v", v), nil
	}
}
