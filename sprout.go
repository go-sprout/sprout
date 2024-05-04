package sprout

import (
	"log/slog"
	"text/template"
)

// ErrHandling defines the strategy for handling errors within FunctionHandler.
// It supports returning default values, panicking, or sending errors to a
// specified channel.
type ErrHandling int

const (
	// ErrHandlingReturnDefaultValue indicates that a default value should be
	// returned on error (default).
	ErrHandlingReturnDefaultValue ErrHandling = iota + 1
	// ErrHandlingPanic indicates that a panic should be raised on error.
	ErrHandlingPanic
	// ErrHandlingErrorChannel indicates that errors should be sent to an error
	// channel.
	ErrHandlingErrorChannel
)

// FunctionHandler manages function execution with configurable error handling
// and logging.
type FunctionHandler struct {
	ErrHandling ErrHandling
	errChan     chan error
	Logger      *slog.Logger
	funcMap     template.FuncMap
	funcsAlias  FunctionAliasMap
}

// FunctionHandlerOption defines a type for functional options that configure
// FunctionHandler.
type FunctionHandlerOption func(*FunctionHandler)

// NewFunctionHandler creates a new FunctionHandler with the provided options.
func NewFunctionHandler(opts ...FunctionHandlerOption) *FunctionHandler {
	fnHandler := &FunctionHandler{
		ErrHandling: ErrHandlingReturnDefaultValue,
		errChan:     make(chan error),
		Logger:      slog.New(&slog.TextHandler{}),
		funcMap:     make(template.FuncMap),
		funcsAlias:  make(FunctionAliasMap),
	}

	for _, opt := range opts {
		opt(fnHandler)
	}

	return fnHandler
}

// WithErrHandling sets the error handling strategy for a FunctionHandler.
func WithErrHandling(eh ErrHandling) FunctionHandlerOption {
	return func(p *FunctionHandler) {
		p.ErrHandling = eh
	}
}

// WithLogger sets the logger used by a FunctionHandler.
func WithLogger(l *slog.Logger) FunctionHandlerOption {
	return func(p *FunctionHandler) {
		p.Logger = l
	}
}

// WithErrorChannel sets the error channel for a FunctionHandler.
func WithErrorChannel(ec chan error) FunctionHandlerOption {
	return func(p *FunctionHandler) {
		p.errChan = ec
	}
}

// WithFunctionHandler updates a FunctionHandler with settings from another FunctionHandler.
// This is useful for copying configurations between handlers.
func WithFunctionHandler(new *FunctionHandler) FunctionHandlerOption {
	return func(fnh *FunctionHandler) {
		*fnh = *new
	}
}

// FuncMap returns a template.FuncMap for use with text/template or html/template.
// It provides backward compatibility with sprig.FuncMap and integrates
// additional configured functions.
// FOR BACKWARD COMPATIBILITY ONLY
func FuncMap(opts ...FunctionHandlerOption) template.FuncMap {
	fnHandler := NewFunctionHandler(opts...)

	// BACKWARD COMPATIBILITY
	// Fallback to FuncMap() to get the unmigrated functions
	for k, v := range TxtFuncMap() {
		fnHandler.funcMap[k] = v
	}

	// Added migrated functions
	fnHandler.funcMap["hello"] = fnHandler.Hello

	// Added functions not migrated totally yet
	fnHandler.funcMap["ago"] = fnHandler.DateAgo
	fnHandler.funcMap["date"] = fnHandler.Date
	fnHandler.funcMap["dateModify"] = fnHandler.DateModify
	fnHandler.funcMap["dateInZone"] = fnHandler.DateInZone
	fnHandler.funcMap["duration"] = fnHandler.Duration
	fnHandler.funcMap["durationRound"] = fnHandler.DurationRound
	fnHandler.funcMap["htmlDate"] = fnHandler.HtmlDate
	fnHandler.funcMap["htmlDateInZone"] = fnHandler.HtmlDateInZone
	fnHandler.funcMap["mustDateModify"] = fnHandler.MustDateModify
	fnHandler.funcMap["mustToDate"] = fnHandler.MustToDate
	fnHandler.funcMap["now"] = fnHandler.Now
	fnHandler.funcMap["toDate"] = fnHandler.ToDate
	fnHandler.funcMap["unixEpoch"] = fnHandler.UnixEpoch
	fnHandler.funcMap["ellipsis"] = fnHandler.Ellipsis
	fnHandler.funcMap["ellipsisBoth"] = fnHandler.EllipsisBoth
	fnHandler.funcMap["toUpper"] = fnHandler.ToUpper
	fnHandler.funcMap["toLower"] = fnHandler.ToLower
	fnHandler.funcMap["title"] = fnHandler.ToTitleCase
	fnHandler.funcMap["untitle"] = fnHandler.Untitle
	fnHandler.funcMap["substr"] = fnHandler.Substring
	fnHandler.funcMap["repeat"] = fnHandler.Repeat
	fnHandler.funcMap["trunc"] = fnHandler.Trunc
	fnHandler.funcMap["trim"] = fnHandler.Trim
	fnHandler.funcMap["trimAll"] = fnHandler.TrimAll
	fnHandler.funcMap["trimPrefix"] = fnHandler.TrimPrefix
	fnHandler.funcMap["trimSuffix"] = fnHandler.TrimSuffix
	fnHandler.funcMap["nospace"] = fnHandler.Nospace
	fnHandler.funcMap["initials"] = fnHandler.Initials
	fnHandler.funcMap["randAlphaNum"] = fnHandler.RandAlphaNumeric
	fnHandler.funcMap["randAlpha"] = fnHandler.RandAlpha
	fnHandler.funcMap["randAscii"] = fnHandler.RandAscii
	fnHandler.funcMap["randNumeric"] = fnHandler.RandNumeric
	fnHandler.funcMap["swapcase"] = fnHandler.SwapCase
	fnHandler.funcMap["shuffle"] = fnHandler.Shuffle
	fnHandler.funcMap["snakecase"] = fnHandler.ToSnakeCase
	fnHandler.funcMap["camelcase"] = fnHandler.ToCamelCase
	fnHandler.funcMap["kebabcase"] = fnHandler.ToKebabCase
	fnHandler.funcMap["pascalcase"] = fnHandler.ToPascalCase
	fnHandler.funcMap["titlecase"] = fnHandler.ToTitleCase
	fnHandler.funcMap["sentencecase"] = fnHandler.ToSentenceCase
	fnHandler.funcMap["dotcase"] = fnHandler.ToDotCase
	fnHandler.funcMap["pathcase"] = fnHandler.ToPathCase
	fnHandler.funcMap["constantcase"] = fnHandler.ToConstantCase
	fnHandler.funcMap["wrap"] = fnHandler.Wrap
	fnHandler.funcMap["wrapWith"] = fnHandler.WrapWith
	fnHandler.funcMap["contains"] = fnHandler.Contains
	fnHandler.funcMap["hasPrefix"] = fnHandler.HasPrefix
	fnHandler.funcMap["hasSuffix"] = fnHandler.HasSuffix
	fnHandler.funcMap["quote"] = fnHandler.Quote
	fnHandler.funcMap["squote"] = fnHandler.Squote
	fnHandler.funcMap["cat"] = fnHandler.Cat
	fnHandler.funcMap["indent"] = fnHandler.Indent
	fnHandler.funcMap["nindent"] = fnHandler.Nindent
	fnHandler.funcMap["replace"] = fnHandler.Replace
	fnHandler.funcMap["plural"] = fnHandler.Plural
	fnHandler.funcMap["sha1sum"] = fnHandler.Sha1sum
	fnHandler.funcMap["sha256sum"] = fnHandler.Sha256sum
	fnHandler.funcMap["adler32sum"] = fnHandler.Adler32sum
	fnHandler.funcMap["toString"] = fnHandler.Strval
	fnHandler.funcMap["int64"] = fnHandler.ToInt64
	fnHandler.funcMap["int"] = fnHandler.ToInt
	fnHandler.funcMap["float64"] = fnHandler.ToFloat64
	fnHandler.funcMap["seq"] = fnHandler.Seq
	fnHandler.funcMap["toDecimal"] = fnHandler.ToDecimal
	fnHandler.funcMap["split"] = fnHandler.Split
	fnHandler.funcMap["splitList"] = fnHandler.SplitList
	fnHandler.funcMap["splitn"] = fnHandler.Splitn
	fnHandler.funcMap["toStrings"] = fnHandler.Strslice // fnHandler.ToStrings
	fnHandler.funcMap["until"] = fnHandler.Until
	fnHandler.funcMap["untilStep"] = fnHandler.UntilStep
	fnHandler.funcMap["add1"] = fnHandler.Add1
	fnHandler.funcMap["add"] = fnHandler.Add
	fnHandler.funcMap["sub"] = fnHandler.Sub
	fnHandler.funcMap["div"] = fnHandler.DivInt
	fnHandler.funcMap["divf"] = fnHandler.Divf
	fnHandler.funcMap["mod"] = fnHandler.Mod
	fnHandler.funcMap["mul"] = fnHandler.MulInt
	fnHandler.funcMap["mulf"] = fnHandler.Mulf
	fnHandler.funcMap["randInt"] = fnHandler.RandInt
	fnHandler.funcMap["max"] = fnHandler.Max
	fnHandler.funcMap["min"] = fnHandler.Min
	fnHandler.funcMap["maxf"] = fnHandler.Maxf
	fnHandler.funcMap["minf"] = fnHandler.Minf
	fnHandler.funcMap["ceil"] = fnHandler.Ceil
	fnHandler.funcMap["floor"] = fnHandler.Floor
	fnHandler.funcMap["round"] = fnHandler.Round
	fnHandler.funcMap["join"] = fnHandler.Join
	fnHandler.funcMap["sortAlpha"] = fnHandler.SortAlpha
	fnHandler.funcMap["default"] = fnHandler.Default
	fnHandler.funcMap["empty"] = fnHandler.Empty
	fnHandler.funcMap["coalesce"] = fnHandler.Coalesce
	fnHandler.funcMap["all"] = fnHandler.All
	fnHandler.funcMap["any"] = fnHandler.Any
	fnHandler.funcMap["compact"] = fnHandler.Compact
	fnHandler.funcMap["mustCompact"] = fnHandler.MustCompact
	fnHandler.funcMap["fromJson"] = fnHandler.FromJson
	fnHandler.funcMap["toJson"] = fnHandler.ToJson
	fnHandler.funcMap["toPrettyJson"] = fnHandler.ToPrettyJson
	fnHandler.funcMap["toRawJson"] = fnHandler.ToRawJson
	fnHandler.funcMap["mustFromJson"] = fnHandler.MustFromJson
	fnHandler.funcMap["mustToJson"] = fnHandler.MustToJson
	fnHandler.funcMap["mustToPrettyJson"] = fnHandler.MustToPrettyJson
	fnHandler.funcMap["mustToRawJson"] = fnHandler.MustToRawJson
	fnHandler.funcMap["ternary"] = fnHandler.Ternary
	fnHandler.funcMap["deepCopy"] = fnHandler.DeepCopy
	fnHandler.funcMap["mustDeepCopy"] = fnHandler.MustDeepCopy
	fnHandler.funcMap["typeOf"] = fnHandler.TypeOf
	fnHandler.funcMap["typeIs"] = fnHandler.TypeIs
	fnHandler.funcMap["typeIsLike"] = fnHandler.TypeIsLike
	fnHandler.funcMap["kindOf"] = fnHandler.KindOf
	fnHandler.funcMap["kindIs"] = fnHandler.KindIs
	fnHandler.funcMap["deepEqual"] = fnHandler.DeepEqual
	fnHandler.funcMap["regexMatch"] = fnHandler.RegexMatch
	fnHandler.funcMap["mustRegexMatch"] = fnHandler.MustRegexMatch
	fnHandler.funcMap["regexFindAll"] = fnHandler.RegexFindAll
	fnHandler.funcMap["mustRegexFindAll"] = fnHandler.MustRegexFindAll
	fnHandler.funcMap["regexFind"] = fnHandler.RegexFind
	fnHandler.funcMap["mustRegexFind"] = fnHandler.MustRegexFind
	fnHandler.funcMap["regexReplaceAll"] = fnHandler.RegexReplaceAll
	fnHandler.funcMap["mustRegexReplaceAll"] = fnHandler.MustRegexReplaceAll
	fnHandler.funcMap["regexReplaceAllLiteral"] = fnHandler.RegexReplaceAllLiteral
	fnHandler.funcMap["mustRegexReplaceAllLiteral"] = fnHandler.MustRegexReplaceAllLiteral
	fnHandler.funcMap["regexSplit"] = fnHandler.RegexSplit
	fnHandler.funcMap["mustRegexSplit"] = fnHandler.MustRegexSplit
	fnHandler.funcMap["regexQuoteMeta"] = fnHandler.RegexQuoteMeta
	fnHandler.funcMap["append"] = fnHandler.Push
	fnHandler.funcMap["mustAppend"] = fnHandler.MustPush
	fnHandler.funcMap["prepend"] = fnHandler.Prepend
	fnHandler.funcMap["mustPrepend"] = fnHandler.MustPrepend
	fnHandler.funcMap["first"] = fnHandler.First
	fnHandler.funcMap["mustFirst"] = fnHandler.MustFirst
	fnHandler.funcMap["rest"] = fnHandler.Rest
	fnHandler.funcMap["mustRest"] = fnHandler.MustRest
	fnHandler.funcMap["last"] = fnHandler.Last
	fnHandler.funcMap["mustLast"] = fnHandler.MustLast
	fnHandler.funcMap["initial"] = fnHandler.Initial
	fnHandler.funcMap["mustInitial"] = fnHandler.MustInitial
	fnHandler.funcMap["reverse"] = fnHandler.Reverse
	fnHandler.funcMap["mustReverse"] = fnHandler.MustReverse
	fnHandler.funcMap["uniq"] = fnHandler.Uniq
	fnHandler.funcMap["mustUniq"] = fnHandler.MustUniq
	fnHandler.funcMap["without"] = fnHandler.Without
	fnHandler.funcMap["mustWithout"] = fnHandler.MustWithout
	fnHandler.funcMap["has"] = fnHandler.Has
	fnHandler.funcMap["mustHas"] = fnHandler.MustHas
	fnHandler.funcMap["slice"] = fnHandler.Slice
	fnHandler.funcMap["mustSlice"] = fnHandler.MustSlice
	fnHandler.funcMap["concat"] = fnHandler.Concat
	fnHandler.funcMap["dig"] = fnHandler.Dig
	fnHandler.funcMap["chunk"] = fnHandler.Chunk
	fnHandler.funcMap["mustChunk"] = fnHandler.MustChunk
	fnHandler.funcMap["list"] = fnHandler.List
	fnHandler.funcMap["dict"] = fnHandler.Dict
	fnHandler.funcMap["get"] = fnHandler.Get
	fnHandler.funcMap["set"] = fnHandler.Set
	fnHandler.funcMap["unset"] = fnHandler.Unset
	fnHandler.funcMap["hasKey"] = fnHandler.HasKey
	fnHandler.funcMap["pluck"] = fnHandler.Pluck
	fnHandler.funcMap["keys"] = fnHandler.Keys
	fnHandler.funcMap["pick"] = fnHandler.Pick
	fnHandler.funcMap["omit"] = fnHandler.Omit
	fnHandler.funcMap["merge"] = fnHandler.Merge
	fnHandler.funcMap["mergeOverwrite"] = fnHandler.MergeOverwrite
	fnHandler.funcMap["mustMerge"] = fnHandler.MustMerge
	fnHandler.funcMap["mustMergeOverwrite"] = fnHandler.MustMergeOverwrite
	fnHandler.funcMap["values"] = fnHandler.Values
	fnHandler.funcMap["bcrypt"] = fnHandler.Bcrypt
	fnHandler.funcMap["htpasswd"] = fnHandler.Htpasswd
	fnHandler.funcMap["genPrivateKey"] = fnHandler.GeneratePrivateKey
	fnHandler.funcMap["derivePassword"] = fnHandler.DerivePassword
	fnHandler.funcMap["buildCustomCert"] = fnHandler.BuildCustomCertificate
	fnHandler.funcMap["genCA"] = fnHandler.GenerateCertificateAuthority
	fnHandler.funcMap["genCAWithKey"] = fnHandler.GenerateCertificateAuthorityWithPEMKey
	fnHandler.funcMap["genSelfSignedCert"] = fnHandler.GenerateSelfSignedCertificate
	fnHandler.funcMap["genSelfSignedCertWithKey"] = fnHandler.GenerateSelfSignedCertificateWithPEMKey
	fnHandler.funcMap["genSignedCert"] = fnHandler.GenerateSignedCertificate
	fnHandler.funcMap["genSignedCertWithKey"] = fnHandler.GenerateSignedCertificateWithPEMKey
	fnHandler.funcMap["encryptAES"] = fnHandler.EncryptAES
	fnHandler.funcMap["decryptAES"] = fnHandler.DecryptAES
	fnHandler.funcMap["randBytes"] = fnHandler.RandBytes
	fnHandler.funcMap["base"] = fnHandler.PathBase
	fnHandler.funcMap["dir"] = fnHandler.PathDir
	fnHandler.funcMap["clean"] = fnHandler.PathClean
	fnHandler.funcMap["ext"] = fnHandler.PathExt
	fnHandler.funcMap["isAbs"] = fnHandler.PathIsAbs
	fnHandler.funcMap["osBase"] = fnHandler.OsBase
	fnHandler.funcMap["osClean"] = fnHandler.OsClean
	fnHandler.funcMap["osDir"] = fnHandler.OsDir
	fnHandler.funcMap["osExt"] = fnHandler.OsExt
	fnHandler.funcMap["osIsAbs"] = fnHandler.OsIsAbs
	fnHandler.funcMap["env"] = fnHandler.Env
	fnHandler.funcMap["expandenv"] = fnHandler.ExpandEnv
	fnHandler.funcMap["getHostByName"] = fnHandler.GetHostByName
	fnHandler.funcMap["uuidv4"] = fnHandler.Uuidv4
	fnHandler.funcMap["semver"] = fnHandler.Semver
	fnHandler.funcMap["semverCompare"] = fnHandler.SemverCompare
	fnHandler.funcMap["fail"] = fnHandler.Fail
	fnHandler.funcMap["urlParse"] = fnHandler.UrlParse
	fnHandler.funcMap["urlJoin"] = fnHandler.UrlJoin
	fnHandler.funcMap["b64enc"] = fnHandler.Base64Encode
	fnHandler.funcMap["b64dec"] = fnHandler.Base64Decode
	fnHandler.funcMap["b32enc"] = fnHandler.Base32Encode
	fnHandler.funcMap["b32dec"] = fnHandler.Base32Decode

	// Register aliases for functions
	fnHandler.registerAliases()
	return fnHandler.funcMap
}
