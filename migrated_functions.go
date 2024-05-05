package sprout

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/hmac"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"hash/adler32"
	"io"
	"math"
	"math/big"
	mathrand "math/rand"
	"net"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"dario.cat/mergo"
	sv2 "github.com/Masterminds/semver/v3"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
	bcrypt_lib "golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

// ! FUTURE: Rename this function to be more explicit
func (fh *FunctionHandler) SplitList(sep string, str string) []string {
	return strings.Split(str, sep)
}

func (fh *FunctionHandler) Strval(v any) string {
	switch v := v.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case error:
		return v.Error()
	case fmt.Stringer:
		return v.String()
	default:
		// Handles any other types by leveraging fmt.Sprintf for a string representation.
		return fmt.Sprintf("%v", v)
	}
}

func (fh *FunctionHandler) Strslice(v any) []string {
	if v == nil {
		return []string{}
	}

	// Handle []string type efficiently without reflection.
	if strs, ok := v.([]string); ok {
		return strs
	}

	// For slices of any, convert each element to a string, skipping nil values.
	if interfaces, ok := v.([]any); ok {
		var result []string
		for _, s := range interfaces {
			if s != nil {
				result = append(result, fh.Strval(s))
			}
		}
		return result
	}

	// Use reflection for other slice types to convert them to []string.
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		var result []string
		for i := 0; i < val.Len(); i++ {
			value := val.Index(i).Interface()
			if value != nil {
				result = append(result, fh.Strval(value))
			}
		}
		return result
	}

	// If it's not a slice, array, or nil, return a slice with the string representation of v.
	return []string{fh.Strval(v)}
}

func (fh *FunctionHandler) FillMapWithParts(parts []string) map[string]string {
	res := make(map[string]string, len(parts))
	for i, v := range parts {
		res[fmt.Sprintf("_%d", i)] = v
	}
	return res
}

func (fh *FunctionHandler) Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func (fh *FunctionHandler) Base64Decode(s string) string {
	bytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

func (fh *FunctionHandler) Base32Encode(s string) string {
	return base32.StdEncoding.EncodeToString([]byte(s))
}

func (fh *FunctionHandler) Base32Decode(s string) string {
	bytes, err := base32.StdEncoding.DecodeString(s)
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

func (fh *FunctionHandler) DictGetOrEmpty(dict map[string]any, key string) string {
	value, ok := dict[key]
	if !ok {
		return ""
	}
	tp := reflect.TypeOf(value).Kind()
	if tp != reflect.String {
		panic(fmt.Sprintf("unable to parse %s key, must be of type string, but %s found", key, tp.String()))
	}
	return reflect.ValueOf(value).String()
}

func (fh *FunctionHandler) UrlParse(v string) map[string]any {
	dict := map[string]any{}
	parsedURL, err := url.Parse(v)
	if err != nil {
		panic(fmt.Sprintf("unable to parse url: %s", err))
	}
	dict["scheme"] = parsedURL.Scheme
	dict["host"] = parsedURL.Host
	dict["hostname"] = parsedURL.Hostname()
	dict["path"] = parsedURL.Path
	dict["query"] = parsedURL.RawQuery
	dict["opaque"] = parsedURL.Opaque
	dict["fragment"] = parsedURL.Fragment
	if parsedURL.User != nil {
		dict["userinfo"] = parsedURL.User.String()
	} else {
		dict["userinfo"] = ""
	}

	return dict
}

func (fh *FunctionHandler) UrlJoin(d map[string]any) string {
	resURL := url.URL{
		Scheme:   fh.DictGetOrEmpty(d, "scheme"),
		Host:     fh.DictGetOrEmpty(d, "host"),
		Path:     fh.DictGetOrEmpty(d, "path"),
		RawQuery: fh.DictGetOrEmpty(d, "query"),
		Opaque:   fh.DictGetOrEmpty(d, "opaque"),
		Fragment: fh.DictGetOrEmpty(d, "fragment"),
	}
	userinfo := fh.DictGetOrEmpty(d, "userinfo")
	var user *url.Userinfo
	if userinfo != "" {
		tempURL, err := url.Parse(fmt.Sprintf("proto://%s@host", userinfo))
		if err != nil {
			panic(fmt.Sprintf("unable to parse userinfo in dict: %s", err))
		}
		user = tempURL.User
	}

	resURL.User = user
	return resURL.String()
}

func (fh *FunctionHandler) RegexMatch(regex string, s string) bool {
	match, _ := regexp.MatchString(regex, s)
	return match
}

func (fh *FunctionHandler) MustRegexMatch(regex string, s string) (bool, error) {
	return regexp.MatchString(regex, s)
}

func (fh *FunctionHandler) RegexFindAll(regex string, s string, n int) []string {
	r := regexp.MustCompile(regex)
	return r.FindAllString(s, n)
}

func (fh *FunctionHandler) MustRegexFindAll(regex string, s string, n int) ([]string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return []string{}, err
	}
	return r.FindAllString(s, n), nil
}

func (fh *FunctionHandler) RegexFind(regex string, s string) string {
	r := regexp.MustCompile(regex)
	return r.FindString(s)
}

func (fh *FunctionHandler) MustRegexFind(regex string, s string) (string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	return r.FindString(s), nil
}

func (fh *FunctionHandler) RegexReplaceAll(regex string, s string, repl string) string {
	r := regexp.MustCompile(regex)
	return r.ReplaceAllString(s, repl)
}

func (fh *FunctionHandler) MustRegexReplaceAll(regex string, s string, repl string) (string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	return r.ReplaceAllString(s, repl), nil
}

func (fh *FunctionHandler) RegexReplaceAllLiteral(regex string, s string, repl string) string {
	r := regexp.MustCompile(regex)
	return r.ReplaceAllLiteralString(s, repl)
}

func (fh *FunctionHandler) MustRegexReplaceAllLiteral(regex string, s string, repl string) (string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	return r.ReplaceAllLiteralString(s, repl), nil
}

func (fh *FunctionHandler) RegexSplit(regex string, s string, n int) []string {
	r := regexp.MustCompile(regex)
	return r.Split(s, n)
}

func (fh *FunctionHandler) MustRegexSplit(regex string, s string, n int) ([]string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return []string{}, err
	}
	return r.Split(s, n), nil
}

func (fh *FunctionHandler) RegexQuoteMeta(s string) string {
	return regexp.QuoteMeta(s)
}

func (fh *FunctionHandler) ToFloat64(v any) float64 {
	return cast.ToFloat64(v)
}

func (fh *FunctionHandler) ToInt(v any) int {
	return cast.ToInt(v)
}

func (fh *FunctionHandler) ToInt64(v any) int64 {
	return cast.ToInt64(v)
}

func (fh *FunctionHandler) ToDecimal(v any) int64 {
	result, err := strconv.ParseInt(fmt.Sprint(v), 8, 64)
	if err != nil {
		return 0
	}
	return result
}

func (fh *FunctionHandler) IntArrayToString(slice []int, delimeter string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), delimeter), "[]")
}

func (fh *FunctionHandler) ExecDecimalOp(a any, b []any, f func(d1, d2 decimal.Decimal) decimal.Decimal) float64 {
	prt := decimal.NewFromFloat(fh.ToFloat64(a))
	for _, x := range b {
		dx := decimal.NewFromFloat(fh.ToFloat64(x))
		prt = f(prt, dx)
	}
	rslt, _ := prt.Float64()
	return rslt
}

func (fh *FunctionHandler) GetHostByName(name string) string {
	addrs, _ := net.LookupHost(name)
	//TODO: add error handing when release v3 comes out
	return addrs[mathrand.Intn(len(addrs))]
}

func (fh *FunctionHandler) List(v ...any) []any {
	return v
}

func (fh *FunctionHandler) Append(list any, v any) []any {
	l, err := fh.MustAppend(list, v)
	if err != nil {
		panic(err)
	}

	return l
}

func (fh *FunctionHandler) MustAppend(list any, v any) ([]any, error) {
	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(list)

		l := l2.Len()
		nl := make([]any, l)
		for i := 0; i < l; i++ {
			nl[i] = l2.Index(i).Interface()
		}

		return append(nl, v), nil

	default:
		return nil, fmt.Errorf("Cannot push on type %s", tp)
	}
}

func (fh *FunctionHandler) Prepend(list any, v any) []any {
	l, err := fh.MustPrepend(list, v)
	if err != nil {
		panic(err)
	}

	return l
}

func (fh *FunctionHandler) MustPrepend(list any, v any) ([]any, error) {
	//return append([]any{v}, list...)

	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(list)

		l := l2.Len()
		nl := make([]any, l)
		for i := 0; i < l; i++ {
			nl[i] = l2.Index(i).Interface()
		}

		return append([]any{v}, nl...), nil

	default:
		return nil, fmt.Errorf("Cannot prepend on type %s", tp)
	}
}

func (fh *FunctionHandler) Chunk(size int, list any) [][]any {
	l, err := fh.MustChunk(size, list)
	if err != nil {
		panic(err)
	}

	return l
}

func (fh *FunctionHandler) MustChunk(size int, list any) ([][]any, error) {
	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(list)

		l := l2.Len()

		cs := int(math.Floor(float64(l-1)/float64(size)) + 1)
		nl := make([][]any, cs)

		for i := 0; i < cs; i++ {
			clen := size
			if i == cs-1 {
				clen = int(math.Floor(math.Mod(float64(l), float64(size))))
				if clen == 0 {
					clen = size
				}
			}

			nl[i] = make([]any, clen)

			for j := 0; j < clen; j++ {
				ix := i*size + j
				nl[i][j] = l2.Index(ix).Interface()
			}
		}

		return nl, nil

	default:
		return nil, fmt.Errorf("Cannot chunk type %s", tp)
	}
}

func (fh *FunctionHandler) Last(list any) any {
	l, err := fh.MustLast(list)
	if err != nil {
		panic(err)
	}

	return l
}

func (fh *FunctionHandler) MustLast(list any) (any, error) {
	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(list)

		l := l2.Len()
		if l == 0 {
			return nil, nil
		}

		return l2.Index(l - 1).Interface(), nil
	default:
		return nil, fmt.Errorf("Cannot find last on type %s", tp)
	}
}

func (fh *FunctionHandler) First(list any) any {
	l, err := fh.MustFirst(list)
	if err != nil {
		panic(err)
	}

	return l
}

func (fh *FunctionHandler) MustFirst(list any) (any, error) {
	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(list)

		l := l2.Len()
		if l == 0 {
			return nil, nil
		}

		return l2.Index(0).Interface(), nil
	default:
		return nil, fmt.Errorf("Cannot find first on type %s", tp)
	}
}

func (fh *FunctionHandler) Rest(list any) []any {
	l, err := fh.MustRest(list)
	if err != nil {
		panic(err)
	}

	return l
}

func (fh *FunctionHandler) MustRest(list any) ([]any, error) {
	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(list)

		l := l2.Len()
		if l == 0 {
			return nil, nil
		}

		nl := make([]any, l-1)
		for i := 1; i < l; i++ {
			nl[i-1] = l2.Index(i).Interface()
		}

		return nl, nil
	default:
		return nil, fmt.Errorf("Cannot find rest on type %s", tp)
	}
}

func (fh *FunctionHandler) Initial(list any) []any {
	l, err := fh.MustInitial(list)
	if err != nil {
		panic(err)
	}

	return l
}

func (fh *FunctionHandler) MustInitial(list any) ([]any, error) {
	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(list)

		l := l2.Len()
		if l == 0 {
			return nil, nil
		}

		nl := make([]any, l-1)
		for i := 0; i < l-1; i++ {
			nl[i] = l2.Index(i).Interface()
		}

		return nl, nil
	default:
		return nil, fmt.Errorf("Cannot find initial on type %s", tp)
	}
}

func (fh *FunctionHandler) SortAlpha(list any) []string {
	k := reflect.Indirect(reflect.ValueOf(list)).Kind()
	switch k {
	case reflect.Slice, reflect.Array:
		a := fh.Strslice(list)
		s := sort.StringSlice(a)
		s.Sort()
		return s
	}
	return []string{fh.Strval(list)}
}

func (fh *FunctionHandler) Reverse(v any) []any {
	l, err := fh.MustReverse(v)
	if err != nil {
		panic(err)
	}

	return l
}

func (fh *FunctionHandler) MustReverse(v any) ([]any, error) {
	tp := reflect.TypeOf(v).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(v)

		l := l2.Len()
		// We do not sort in place because the incoming array should not be altered.
		nl := make([]any, l)
		for i := 0; i < l; i++ {
			nl[l-i-1] = l2.Index(i).Interface()
		}

		return nl, nil
	default:
		return nil, fmt.Errorf("Cannot find reverse on type %s", tp)
	}
}

func (fh *FunctionHandler) Compact(list any) []any {
	l, err := fh.MustCompact(list)
	if err != nil {
		panic(err)
	}

	return l
}

func (fh *FunctionHandler) MustCompact(list any) ([]any, error) {
	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(list)

		l := l2.Len()
		nl := []any{}
		var item any
		for i := 0; i < l; i++ {
			item = l2.Index(i).Interface()
			if !fh.Empty(item) {
				nl = append(nl, item)
			}
		}

		return nl, nil
	default:
		return nil, fmt.Errorf("Cannot compact on type %s", tp)
	}
}

func (fh *FunctionHandler) Uniq(list any) []any {
	l, err := fh.MustUniq(list)
	if err != nil {
		panic(err)
	}

	return l
}

func (fh *FunctionHandler) MustUniq(list any) ([]any, error) {
	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(list)

		l := l2.Len()
		dest := []any{}
		var item any
		for i := 0; i < l; i++ {
			item = l2.Index(i).Interface()
			if !fh.InList(dest, item) {
				dest = append(dest, item)
			}
		}

		return dest, nil
	default:
		return nil, fmt.Errorf("Cannot find uniq on type %s", tp)
	}
}

func (fh *FunctionHandler) InList(haystack []any, needle any) bool {
	for _, h := range haystack {
		if reflect.DeepEqual(needle, h) {
			return true
		}
	}
	return false
}

func (fh *FunctionHandler) Without(list any, omit ...any) []any {
	l, err := fh.MustWithout(list, omit...)
	if err != nil {
		panic(err)
	}

	return l
}

func (fh *FunctionHandler) MustWithout(list any, omit ...any) ([]any, error) {
	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(list)

		l := l2.Len()
		res := []any{}
		var item any
		for i := 0; i < l; i++ {
			item = l2.Index(i).Interface()
			if !fh.InList(omit, item) {
				res = append(res, item)
			}
		}

		return res, nil
	default:
		return nil, fmt.Errorf("Cannot find without on type %s", tp)
	}
}

func (fh *FunctionHandler) Has(needle any, haystack any) bool {
	l, err := fh.MustHas(needle, haystack)
	if err != nil {
		panic(err)
	}

	return l
}

func (fh *FunctionHandler) MustHas(needle any, haystack any) (bool, error) {
	if haystack == nil {
		return false, nil
	}
	tp := reflect.TypeOf(haystack).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(haystack)
		var item any
		l := l2.Len()
		for i := 0; i < l; i++ {
			item = l2.Index(i).Interface()
			if reflect.DeepEqual(needle, item) {
				return true, nil
			}
		}

		return false, nil
	default:
		return false, fmt.Errorf("Cannot find has on type %s", tp)
	}
}

func (fh *FunctionHandler) Slice(list any, indices ...any) any {
	l, err := fh.MustSlice(list, indices...)
	if err != nil {
		panic(err)
	}

	return l
}

func (fh *FunctionHandler) MustSlice(list any, indices ...any) (any, error) {
	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		l2 := reflect.ValueOf(list)

		l := l2.Len()
		if l == 0 {
			return nil, nil
		}

		var start, end int
		if len(indices) > 0 {
			start = fh.ToInt(indices[0])
		}
		if len(indices) < 2 {
			end = l
		} else {
			end = fh.ToInt(indices[1])
		}

		return l2.Slice(start, end).Interface(), nil
	default:
		return nil, fmt.Errorf("list should be type of slice or array but %s", tp)
	}
}

func (fh *FunctionHandler) Concat(lists ...any) any {
	var res []any
	for _, list := range lists {
		tp := reflect.TypeOf(list).Kind()
		switch tp {
		case reflect.Slice, reflect.Array:
			l2 := reflect.ValueOf(list)
			for i := 0; i < l2.Len(); i++ {
				res = append(res, l2.Index(i).Interface())
			}
		default:
			panic(fmt.Sprintf("Cannot concat type %s as list", tp))
		}
	}
	return res
}

func (fh *FunctionHandler) Get(d map[string]any, key string) any {
	if val, ok := d[key]; ok {
		return val
	}
	return ""
}

func (fh *FunctionHandler) Set(d map[string]any, key string, value any) map[string]any {
	d[key] = value
	return d
}

func (fh *FunctionHandler) Unset(d map[string]any, key string) map[string]any {
	delete(d, key)
	return d
}

func (fh *FunctionHandler) HasKey(d map[string]any, key string) bool {
	_, ok := d[key]
	return ok
}

func (fh *FunctionHandler) Pluck(key string, d ...map[string]any) []any {
	res := []any{}
	for _, dict := range d {
		if val, ok := dict[key]; ok {
			res = append(res, val)
		}
	}
	return res
}

func (fh *FunctionHandler) Keys(dicts ...map[string]any) []string {
	k := []string{}
	for _, dict := range dicts {
		for key := range dict {
			k = append(k, key)
		}
	}
	return k
}

func (fh *FunctionHandler) Pick(dict map[string]any, keys ...string) map[string]any {
	res := map[string]any{}
	for _, k := range keys {
		if v, ok := dict[k]; ok {
			res[k] = v
		}
	}
	return res
}

func (fh *FunctionHandler) Omit(dict map[string]any, keys ...string) map[string]any {
	res := map[string]any{}

	omit := make(map[string]bool, len(keys))
	for _, k := range keys {
		omit[k] = true
	}

	for k, v := range dict {
		if _, ok := omit[k]; !ok {
			res[k] = v
		}
	}
	return res
}

func (fh *FunctionHandler) Dict(v ...any) map[string]any {
	dict := map[string]any{}
	lenv := len(v)
	for i := 0; i < lenv; i += 2 {
		key := fh.Strval(v[i])
		if i+1 >= lenv {
			dict[key] = ""
			continue
		}
		dict[key] = v[i+1]
	}
	return dict
}

func (fh *FunctionHandler) Merge(dst map[string]any, srcs ...map[string]any) any {
	for _, src := range srcs {
		if err := mergo.Merge(&dst, src); err != nil {
			// Swallow errors inside of a template.
			return ""
		}
	}
	return dst
}

func (fh *FunctionHandler) MustMerge(dst map[string]any, srcs ...map[string]any) (any, error) {
	for _, src := range srcs {
		if err := mergo.Merge(&dst, src); err != nil {
			return nil, err
		}
	}
	return dst, nil
}

func (fh *FunctionHandler) MergeOverwrite(dst map[string]any, srcs ...map[string]any) any {
	for _, src := range srcs {
		if err := mergo.MergeWithOverwrite(&dst, src); err != nil {
			// Swallow errors inside of a template.
			return ""
		}
	}
	return dst
}

func (fh *FunctionHandler) MustMergeOverwrite(dst map[string]any, srcs ...map[string]any) (any, error) {
	for _, src := range srcs {
		if err := mergo.MergeWithOverwrite(&dst, src); err != nil {
			return nil, err
		}
	}
	return dst, nil
}

func (fh *FunctionHandler) Values(dict map[string]any) []any {
	values := []any{}
	for _, value := range dict {
		values = append(values, value)
	}

	return values
}

func (fh *FunctionHandler) Dig(ps ...any) (any, error) {
	if len(ps) < 3 {
		panic("dig needs at least three arguments")
	}
	dict := ps[len(ps)-1].(map[string]any)
	def := ps[len(ps)-2]
	ks := make([]string, len(ps)-2)
	for i := 0; i < len(ks); i++ {
		ks[i] = ps[i].(string)
	}

	return fh.DigFromDict(dict, def, ks)
}

func (fh *FunctionHandler) DigFromDict(dict map[string]any, d any, ks []string) (any, error) {
	k, ns := ks[0], ks[1:]
	step, has := dict[k]
	if !has {
		return d, nil
	}
	if len(ns) == 0 {
		return step, nil
	}
	return fh.DigFromDict(step.(map[string]any), d, ns)
}

func (fh *FunctionHandler) FromJson(v string) any {
	output, _ := fh.MustFromJson(v)
	return output
}

func (fh *FunctionHandler) MustFromJson(v string) (any, error) {
	var output any
	err := json.Unmarshal([]byte(v), &output)
	return output, err
}

func (fh *FunctionHandler) ToJson(v any) string {
	output, _ := json.Marshal(v)
	return string(output)
}

func (fh *FunctionHandler) MustToJson(v any) (string, error) {
	output, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func (fh *FunctionHandler) ToPrettyJson(v any) string {
	output, _ := json.MarshalIndent(v, "", "  ")
	return string(output)
}

func (fh *FunctionHandler) MustToPrettyJson(v any) (string, error) {
	output, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func (fh *FunctionHandler) ToRawJson(v any) string {
	output, err := fh.MustToRawJson(v)
	if err != nil {
		panic(err)
	}
	return string(output)
}

func (fh *FunctionHandler) MustToRawJson(v any) (string, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(&v)
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(buf.String(), "\n"), nil
}

func (fh *FunctionHandler) Date(fmt string, date any) string {
	return fh.DateInZone(fmt, date, "Local")
}

func (fh *FunctionHandler) HtmlDate(date any) string {
	return fh.DateInZone("2006-01-02", date, "Local")
}

func (fh *FunctionHandler) HtmlDateInZone(date any, zone string) string {
	return fh.DateInZone("2006-01-02", date, zone)
}

func (fh *FunctionHandler) DateInZone(fmt string, date any, zone string) string {
	var t time.Time
	switch date := date.(type) {
	default:
		t = time.Now()
	case time.Time:
		t = date
	case *time.Time:
		t = *date
	case int64:
		t = time.Unix(date, 0)
	case int:
		t = time.Unix(int64(date), 0)
	case int32:
		t = time.Unix(int64(date), 0)
	}

	loc, err := time.LoadLocation(zone)
	if err != nil {
		loc, _ = time.LoadLocation("UTC")
	}

	return t.In(loc).Format(fmt)
}

func (fh *FunctionHandler) DateModify(fmt string, date time.Time) time.Time {
	d, err := time.ParseDuration(fmt)
	if err != nil {
		return date
	}
	return date.Add(d)
}

func (fh *FunctionHandler) MustDateModify(fmt string, date time.Time) (time.Time, error) {
	d, err := time.ParseDuration(fmt)
	if err != nil {
		return time.Time{}, err
	}
	return date.Add(d), nil
}

func (fh *FunctionHandler) DateAgo(date any) string {
	var t time.Time

	switch date := date.(type) {
	default:
		t = time.Now()
	case time.Time:
		t = date
	case int64:
		t = time.Unix(date, 0)
	case int:
		t = time.Unix(int64(date), 0)
	}
	// Drop resolution to seconds
	duration := time.Since(t).Round(time.Second)
	return duration.String()
}

func (fh *FunctionHandler) Duration(sec any) string {
	var n int64
	switch value := sec.(type) {
	default:
		n = 0
	case string:
		n, _ = strconv.ParseInt(value, 10, 64)
	case int64:
		n = value
	}
	return (time.Duration(n) * time.Second).String()
}

func (fh *FunctionHandler) DurationRound(duration any) string {
	var d time.Duration
	switch duration := duration.(type) {
	default:
		d = 0
	case string:
		d, _ = time.ParseDuration(duration)
	case int64:
		d = time.Duration(duration)
	case time.Time:
		d = time.Since(duration)
	}

	u := uint64(d)
	neg := d < 0
	if neg {
		u = -u
	}

	var (
		year   = uint64(time.Hour) * 24 * 365
		month  = uint64(time.Hour) * 24 * 30
		day    = uint64(time.Hour) * 24
		hour   = uint64(time.Hour)
		minute = uint64(time.Minute)
		second = uint64(time.Second)
	)
	switch {
	case u > year:
		return strconv.FormatUint(u/year, 10) + "y"
	case u > month:
		return strconv.FormatUint(u/month, 10) + "mo"
	case u > day:
		return strconv.FormatUint(u/day, 10) + "d"
	case u > hour:
		return strconv.FormatUint(u/hour, 10) + "h"
	case u > minute:
		return strconv.FormatUint(u/minute, 10) + "m"
	case u > second:
		return strconv.FormatUint(u/second, 10) + "s"
	}
	return "0s"
}

func (fh *FunctionHandler) ToDate(fmt, str string) time.Time {
	t, _ := time.ParseInLocation(fmt, str, time.Local)
	return t
}

func (fh *FunctionHandler) MustToDate(fmt, str string) (time.Time, error) {
	return time.ParseInLocation(fmt, str, time.Local)
}

func (fh *FunctionHandler) UnixEpoch(date time.Time) string {
	return strconv.FormatInt(date.Unix(), 10)
}

func (fh *FunctionHandler) SemverCompare(constraint, version string) (bool, error) {
	c, err := sv2.NewConstraint(constraint)
	if err != nil {
		return false, err
	}

	v, err := sv2.NewVersion(version)
	if err != nil {
		return false, err
	}

	return c.Check(v), nil
}

func (fh *FunctionHandler) Semver(version string) (*sv2.Version, error) {
	return sv2.NewVersion(version)
}

func (fh *FunctionHandler) Now() time.Time {
	return time.Now()
}

func (fh *FunctionHandler) PathBase(str string) string {
	return path.Base(str)
}

func (fh *FunctionHandler) PathDir(str string) string {
	return path.Dir(str)
}

func (fh *FunctionHandler) PathExt(str string) string {
	return path.Ext(str)
}

func (fh *FunctionHandler) PathClean(str string) string {
	return path.Clean(str)
}

func (fh *FunctionHandler) PathIsAbs(str string) bool {
	return path.IsAbs(str)
}

func (fh *FunctionHandler) OsBase(str string) string {
	return filepath.Base(str)
}

func (fh *FunctionHandler) OsDir(str string) string {
	return filepath.Dir(str)
}

func (fh *FunctionHandler) OsExt(str string) string {
	return filepath.Ext(str)
}

func (fh *FunctionHandler) OsClean(str string) string {
	return filepath.Clean(str)
}

func (fh *FunctionHandler) OsIsAbs(str string) bool {
	return filepath.IsAbs(str)
}

func (fh *FunctionHandler) Env(key string) string {
	return os.Getenv(key)
}

func (fh *FunctionHandler) ExpandEnv(str string) string {
	return os.ExpandEnv(str)
}

// //////////
// CRYPTO //
// //////////
func (fh *FunctionHandler) Sha256sum(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

func (fh *FunctionHandler) Sha1sum(input string) string {
	hash := sha1.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func (fh *FunctionHandler) Adler32sum(input string) string {
	hash := adler32.Checksum([]byte(input))
	return fmt.Sprintf("%d", hash)
}

func (fh *FunctionHandler) Bcrypt(input string) string {
	hash, err := bcrypt_lib.GenerateFromPassword([]byte(input), bcrypt_lib.DefaultCost)
	if err != nil {
		return fmt.Sprintf("failed to encrypt string with bcrypt: %s", err)
	}

	return string(hash)
}

func (fh *FunctionHandler) Htpasswd(username string, password string) string {
	if strings.Contains(username, ":") {
		return fmt.Sprintf("invalid username: %s", username)
	}
	return fmt.Sprintf("%s:%s", username, fh.Bcrypt(password))
}

var masterPasswordSeed = "com.lyndir.masterpassword"

var passwordTypeTemplates = map[string][][]byte{
	"maximum": {[]byte("anoxxxxxxxxxxxxxxxxx"), []byte("axxxxxxxxxxxxxxxxxno")},
	"long": {[]byte("CvcvnoCvcvCvcv"), []byte("CvcvCvcvnoCvcv"), []byte("CvcvCvcvCvcvno"), []byte("CvccnoCvcvCvcv"), []byte("CvccCvcvnoCvcv"),
		[]byte("CvccCvcvCvcvno"), []byte("CvcvnoCvccCvcv"), []byte("CvcvCvccnoCvcv"), []byte("CvcvCvccCvcvno"), []byte("CvcvnoCvcvCvcc"),
		[]byte("CvcvCvcvnoCvcc"), []byte("CvcvCvcvCvccno"), []byte("CvccnoCvccCvcv"), []byte("CvccCvccnoCvcv"), []byte("CvccCvccCvcvno"),
		[]byte("CvcvnoCvccCvcc"), []byte("CvcvCvccnoCvcc"), []byte("CvcvCvccCvccno"), []byte("CvccnoCvcvCvcc"), []byte("CvccCvcvnoCvcc"),
		[]byte("CvccCvcvCvccno")},
	"medium": {[]byte("CvcnoCvc"), []byte("CvcCvcno")},
	"short":  {[]byte("Cvcn")},
	"basic":  {[]byte("aaanaaan"), []byte("aannaaan"), []byte("aaannaaa")},
	"pin":    {[]byte("nnnn")},
}

var templateCharacters = map[byte]string{
	'V': "AEIOU",
	'C': "BCDFGHJKLMNPQRSTVWXYZ",
	'v': "aeiou",
	'c': "bcdfghjklmnpqrstvwxyz",
	'A': "AEIOUBCDFGHJKLMNPQRSTVWXYZ",
	'a': "AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz",
	'n': "0123456789",
	'o': "@&%?,=[]_:-+*$#!'^~;()/.",
	'x': "AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz0123456789!@#$%^&*()",
}

func (fh *FunctionHandler) DerivePassword(counter uint32, passwordType, password, user, site string) string {
	var templates = passwordTypeTemplates[passwordType]
	if templates == nil {
		return fmt.Sprintf("cannot find password template %s", passwordType)
	}

	var buffer bytes.Buffer
	buffer.WriteString(masterPasswordSeed)
	binary.Write(&buffer, binary.BigEndian, uint32(len(user)))
	buffer.WriteString(user)

	salt := buffer.Bytes()
	key, err := scrypt.Key([]byte(password), salt, 32768, 8, 2, 64)
	if err != nil {
		return fmt.Sprintf("failed to derive password: %s", err)
	}

	buffer.Truncate(len(masterPasswordSeed))
	binary.Write(&buffer, binary.BigEndian, uint32(len(site)))
	buffer.WriteString(site)
	binary.Write(&buffer, binary.BigEndian, counter)

	var hmacv = hmac.New(sha256.New, key)
	hmacv.Write(buffer.Bytes())
	var seed = hmacv.Sum(nil)
	var temp = templates[int(seed[0])%len(templates)]

	buffer.Truncate(0)
	for i, element := range temp {
		passChars := templateCharacters[element]
		passChar := passChars[int(seed[i+1])%len(passChars)]
		buffer.WriteByte(passChar)
	}

	return buffer.String()
}

func (fh *FunctionHandler) GeneratePrivateKey(typ string) string {
	var priv interface{}
	var err error
	switch typ {
	case "", "rsa":
		// good enough for government work
		priv, err = rsa.GenerateKey(cryptorand.Reader, 4096)
	case "dsa":
		key := new(dsa.PrivateKey)
		// again, good enough for government work
		if err = dsa.GenerateParameters(&key.Parameters, cryptorand.Reader, dsa.L2048N256); err != nil {
			return fmt.Sprintf("failed to generate dsa params: %s", err)
		}
		err = dsa.GenerateKey(key, cryptorand.Reader)
		priv = key
	case "ecdsa":
		// again, good enough for government work
		priv, err = ecdsa.GenerateKey(elliptic.P256(), cryptorand.Reader)
	case "ed25519":
		_, priv, err = ed25519.GenerateKey(cryptorand.Reader)
	default:
		return "Unknown type " + typ
	}
	if err != nil {
		return fmt.Sprintf("failed to generate private key: %s", err)
	}

	return string(pem.EncodeToMemory(fh.PemBlockForKey(priv)))
}

// DSAKeyFormat stores the format for DSA keys.
// Used by pemBlockForKey
type DSAKeyFormat struct {
	Version       int
	P, Q, G, Y, X *big.Int
}

func (fh *FunctionHandler) PemBlockForKey(priv interface{}) *pem.Block {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}
	case *dsa.PrivateKey:
		val := DSAKeyFormat{
			P: k.P, Q: k.Q, G: k.G,
			Y: k.Y, X: k.X,
		}
		bytes, _ := asn1.Marshal(val)
		return &pem.Block{Type: "DSA PRIVATE KEY", Bytes: bytes}
	case *ecdsa.PrivateKey:
		b, _ := x509.MarshalECPrivateKey(k)
		return &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}
	default:
		// attempt PKCS#8 format for all other keys
		b, err := x509.MarshalPKCS8PrivateKey(k)
		if err != nil {
			return nil
		}
		return &pem.Block{Type: "PRIVATE KEY", Bytes: b}
	}
}

func (fh *FunctionHandler) ParsePrivateKeyPEM(pemBlock string) (crypto.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemBlock))
	if block == nil {
		return nil, errors.New("no PEM data in input")
	}

	if block.Type == "PRIVATE KEY" {
		priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("decoding PEM as PKCS#8: %s", err)
		}
		return priv, nil
	} else if !strings.HasSuffix(block.Type, " PRIVATE KEY") {
		return nil, fmt.Errorf("no private key data in PEM block of type %s", block.Type)
	}

	switch block.Type[:len(block.Type)-12] { // strip " PRIVATE KEY"
	case "RSA":
		priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("parsing RSA private key from PEM: %s", err)
		}
		return priv, nil
	case "EC":
		priv, err := x509.ParseECPrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("parsing EC private key from PEM: %s", err)
		}
		return priv, nil
	case "DSA":
		var k DSAKeyFormat
		_, err := asn1.Unmarshal(block.Bytes, &k)
		if err != nil {
			return nil, fmt.Errorf("parsing DSA private key from PEM: %s", err)
		}
		priv := &dsa.PrivateKey{
			PublicKey: dsa.PublicKey{
				Parameters: dsa.Parameters{
					P: k.P, Q: k.Q, G: k.G,
				},
				Y: k.Y,
			},
			X: k.X,
		}
		return priv, nil
	default:
		return nil, fmt.Errorf("invalid private key type %s", block.Type)
	}
}

func (fh *FunctionHandler) GetPublicKey(priv crypto.PrivateKey) (crypto.PublicKey, error) {
	switch k := priv.(type) {
	case interface{ Public() crypto.PublicKey }:
		return k.Public(), nil
	case *dsa.PrivateKey:
		return &k.PublicKey, nil
	default:
		return nil, fmt.Errorf("unable to get public key for type %T", priv)
	}
}

type certificate struct {
	Cert string
	Key  string
}

func (fh *FunctionHandler) BuildCustomCertificate(b64cert string, b64key string) (certificate, error) {
	crt := certificate{}

	cert, err := base64.StdEncoding.DecodeString(b64cert)
	if err != nil {
		return crt, errors.New("unable to decode base64 certificate")
	}

	key, err := base64.StdEncoding.DecodeString(b64key)
	if err != nil {
		return crt, errors.New("unable to decode base64 private key")
	}

	decodedCert, _ := pem.Decode(cert)
	if decodedCert == nil {
		return crt, errors.New("unable to decode certificate")
	}
	_, err = x509.ParseCertificate(decodedCert.Bytes)
	if err != nil {
		return crt, fmt.Errorf(
			"error parsing certificate: decodedCert.Bytes: %s",
			err,
		)
	}

	_, err = fh.ParsePrivateKeyPEM(string(key))
	if err != nil {
		return crt, fmt.Errorf(
			"error parsing private key: %s",
			err,
		)
	}

	crt.Cert = string(cert)
	crt.Key = string(key)

	return crt, nil
}

func (fh *FunctionHandler) GenerateCertificateAuthority(
	cn string,
	daysValid int,
) (certificate, error) {
	priv, err := rsa.GenerateKey(cryptorand.Reader, 2048)
	if err != nil {
		return certificate{}, fmt.Errorf("error generating rsa key: %s", err)
	}

	return fh.GenerateCertificateAuthorityWithKeyInternal(cn, daysValid, priv)
}

func (fh *FunctionHandler) GenerateCertificateAuthorityWithPEMKey(
	cn string,
	daysValid int,
	privPEM string,
) (certificate, error) {
	priv, err := fh.ParsePrivateKeyPEM(privPEM)
	if err != nil {
		return certificate{}, fmt.Errorf("parsing private key: %s", err)
	}
	return fh.GenerateCertificateAuthorityWithKeyInternal(cn, daysValid, priv)
}

func (fh *FunctionHandler) GenerateCertificateAuthorityWithKeyInternal(
	cn string,
	daysValid int,
	priv crypto.PrivateKey,
) (certificate, error) {
	ca := certificate{}

	template, err := fh.GetBaseCertTemplate(cn, nil, nil, daysValid)
	if err != nil {
		return ca, err
	}
	// Override KeyUsage and IsCA
	template.KeyUsage = x509.KeyUsageKeyEncipherment |
		x509.KeyUsageDigitalSignature |
		x509.KeyUsageCertSign
	template.IsCA = true

	ca.Cert, ca.Key, err = fh.GetCertAndKey(template, priv, template, priv)

	return ca, err
}

func (fh *FunctionHandler) GenerateSelfSignedCertificate(
	cn string,
	ips []interface{},
	alternateDNS []interface{},
	daysValid int,
) (certificate, error) {
	priv, err := rsa.GenerateKey(cryptorand.Reader, 2048)
	if err != nil {
		return certificate{}, fmt.Errorf("error generating rsa key: %s", err)
	}
	return fh.GenerateSelfSignedCertificateWithKeyInternal(cn, ips, alternateDNS, daysValid, priv)
}

func (fh *FunctionHandler) GenerateSelfSignedCertificateWithPEMKey(
	cn string,
	ips []interface{},
	alternateDNS []interface{},
	daysValid int,
	privPEM string,
) (certificate, error) {
	priv, err := fh.ParsePrivateKeyPEM(privPEM)
	if err != nil {
		return certificate{}, fmt.Errorf("parsing private key: %s", err)
	}
	return fh.GenerateSelfSignedCertificateWithKeyInternal(cn, ips, alternateDNS, daysValid, priv)
}

func (fh *FunctionHandler) GenerateSelfSignedCertificateWithKeyInternal(
	cn string,
	ips []interface{},
	alternateDNS []interface{},
	daysValid int,
	priv crypto.PrivateKey,
) (certificate, error) {
	cert := certificate{}

	template, err := fh.GetBaseCertTemplate(cn, ips, alternateDNS, daysValid)
	if err != nil {
		return cert, err
	}

	cert.Cert, cert.Key, err = fh.GetCertAndKey(template, priv, template, priv)

	return cert, err
}

func (fh *FunctionHandler) GenerateSignedCertificate(
	cn string,
	ips []interface{},
	alternateDNS []interface{},
	daysValid int,
	ca certificate,
) (certificate, error) {
	priv, err := rsa.GenerateKey(cryptorand.Reader, 2048)
	if err != nil {
		return certificate{}, fmt.Errorf("error generating rsa key: %s", err)
	}
	return fh.GenerateSignedCertificateWithKeyInternal(cn, ips, alternateDNS, daysValid, ca, priv)
}

func (fh *FunctionHandler) GenerateSignedCertificateWithPEMKey(
	cn string,
	ips []interface{},
	alternateDNS []interface{},
	daysValid int,
	ca certificate,
	privPEM string,
) (certificate, error) {
	priv, err := fh.ParsePrivateKeyPEM(privPEM)
	if err != nil {
		return certificate{}, fmt.Errorf("parsing private key: %s", err)
	}
	return fh.GenerateSignedCertificateWithKeyInternal(cn, ips, alternateDNS, daysValid, ca, priv)
}

func (fh *FunctionHandler) GenerateSignedCertificateWithKeyInternal(
	cn string,
	ips []interface{},
	alternateDNS []interface{},
	daysValid int,
	ca certificate,
	priv crypto.PrivateKey,
) (certificate, error) {
	cert := certificate{}

	decodedSignerCert, _ := pem.Decode([]byte(ca.Cert))
	if decodedSignerCert == nil {
		return cert, errors.New("unable to decode certificate")
	}
	signerCert, err := x509.ParseCertificate(decodedSignerCert.Bytes)
	if err != nil {
		return cert, fmt.Errorf(
			"error parsing certificate: decodedSignerCert.Bytes: %s",
			err,
		)
	}
	signerKey, err := fh.ParsePrivateKeyPEM(ca.Key)
	if err != nil {
		return cert, fmt.Errorf(
			"error parsing private key: %s",
			err,
		)
	}

	template, err := fh.GetBaseCertTemplate(cn, ips, alternateDNS, daysValid)
	if err != nil {
		return cert, err
	}

	cert.Cert, cert.Key, err = fh.GetCertAndKey(
		template,
		priv,
		signerCert,
		signerKey,
	)

	return cert, err
}

func (fh *FunctionHandler) GetCertAndKey(
	template *x509.Certificate,
	signeeKey crypto.PrivateKey,
	parent *x509.Certificate,
	signingKey crypto.PrivateKey,
) (string, string, error) {
	signeePubKey, err := fh.GetPublicKey(signeeKey)
	if err != nil {
		return "", "", fmt.Errorf("error retrieving public key from signee key: %s", err)
	}
	derBytes, err := x509.CreateCertificate(
		cryptorand.Reader,
		template,
		parent,
		signeePubKey,
		signingKey,
	)
	if err != nil {
		return "", "", fmt.Errorf("error creating certificate: %s", err)
	}

	certBuffer := bytes.Buffer{}
	if err := pem.Encode(
		&certBuffer,
		&pem.Block{Type: "CERTIFICATE", Bytes: derBytes},
	); err != nil {
		return "", "", fmt.Errorf("error pem-encoding certificate: %s", err)
	}

	keyBuffer := bytes.Buffer{}
	if err := pem.Encode(
		&keyBuffer,
		fh.PemBlockForKey(signeeKey),
	); err != nil {
		return "", "", fmt.Errorf("error pem-encoding key: %s", err)
	}

	return certBuffer.String(), keyBuffer.String(), nil
}

func (fh *FunctionHandler) GetBaseCertTemplate(
	cn string,
	ips []interface{},
	alternateDNS []interface{},
	daysValid int,
) (*x509.Certificate, error) {
	ipAddresses, err := fh.GetNetIPs(ips)
	if err != nil {
		return nil, err
	}
	dnsNames, err := fh.GetAlternateDNSStrs(alternateDNS)
	if err != nil {
		return nil, err
	}
	serialNumberUpperBound := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := cryptorand.Int(cryptorand.Reader, serialNumberUpperBound)
	if err != nil {
		return nil, err
	}
	return &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName: cn,
		},
		IPAddresses: ipAddresses,
		DNSNames:    dnsNames,
		NotBefore:   time.Now(),
		NotAfter:    time.Now().Add(time.Hour * 24 * time.Duration(daysValid)),
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
			x509.ExtKeyUsageClientAuth,
		},
		BasicConstraintsValid: true,
	}, nil
}

func (fh *FunctionHandler) GetNetIPs(ips []interface{}) ([]net.IP, error) {
	if ips == nil {
		return []net.IP{}, nil
	}
	var ipStr string
	var ok bool
	var netIP net.IP
	netIPs := make([]net.IP, len(ips))
	for i, ip := range ips {
		ipStr, ok = ip.(string)
		if !ok {
			return nil, fmt.Errorf("error parsing ip: %v is not a string", ip)
		}
		netIP = net.ParseIP(ipStr)
		if netIP == nil {
			return nil, fmt.Errorf("error parsing ip: %s", ipStr)
		}
		netIPs[i] = netIP
	}
	return netIPs, nil
}

func (fh *FunctionHandler) GetAlternateDNSStrs(alternateDNS []interface{}) ([]string, error) {
	if alternateDNS == nil {
		return []string{}, nil
	}
	var dnsStr string
	var ok bool
	alternateDNSStrs := make([]string, len(alternateDNS))
	for i, dns := range alternateDNS {
		dnsStr, ok = dns.(string)
		if !ok {
			return nil, fmt.Errorf(
				"error processing alternate dns name: %v is not a string",
				dns,
			)
		}
		alternateDNSStrs[i] = dnsStr
	}
	return alternateDNSStrs, nil
}

func (fh *FunctionHandler) EncryptAES(password string, plaintext string) (string, error) {
	if plaintext == "" {
		return "", nil
	}

	key := make([]byte, 32)
	copy(key, []byte(password))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	content := []byte(plaintext)
	blockSize := block.BlockSize()
	padding := blockSize - len(content)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	content = append(content, padtext...)

	ciphertext := make([]byte, aes.BlockSize+len(content))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(cryptorand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], content)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (fh *FunctionHandler) DecryptAES(password string, crypt64 string) (string, error) {
	if crypt64 == "" {
		return "", nil
	}

	key := make([]byte, 32)
	copy(key, []byte(password))

	crypt, err := base64.StdEncoding.DecodeString(crypt64)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := crypt[:aes.BlockSize]
	crypt = crypt[aes.BlockSize:]
	decrypted := make([]byte, len(crypt))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, crypt)

	return string(decrypted[:len(decrypted)-int(decrypted[len(decrypted)-1])]), nil
}
