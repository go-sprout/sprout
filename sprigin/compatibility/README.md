## compatibility

This directory contains compatibility tests with Sprig in a separate Go module
to avoid adding Sprig as a dependency to Sprout.

To run these tests, change to this directory and run:

    go test

Tests that are known to fail are skipped.

**Total Test Cases: 2053+** (CODE_000001 through CODE_002053)

### Functions Covered by Category

| Category | Functions | Test Codes |
  |----------|-----------|------------|
| **String - Trim** | trim, trimAll, trimall, trimPrefix, trimSuffix | 000001-000043 |
| **String - Case** | upper, lower, title, untitle | 000044-000083 |
| **String - Manipulation** | repeat, substr, nospace, trunc | 000084-000123 |
| **String - Abbrev/Wrap** | abbrev, abbrevboth, initials, wrap, wrapWith | 000124-000173 |
| **String - Case Conv** | snakecase, camelcase, kebabcase, swapcase | 000174-000213 |
| **String - Search** | contains, hasPrefix, hasSuffix | 000214-000243 |
| **String - Format** | quote, squote, cat, replace, indent, nindent, plural, shuffle | 000244-000323 |
| **Random** | randAlphaNum, randAlpha, randAscii, randNumeric | 000324-000363 |
| **Hash** | sha1sum, sha256sum, sha512sum, adler32sum | 000364-000403 |
| **Math Int** | add, add1, sub, div, mod, mul, max, min, biggest, randInt | 000404-000503 |
| **Math Seq** | until, untilStep, seq | 000504-000533 |
| **Math Float** | addf, add1f, subf, divf, mulf, maxf, minf, floor, ceil, round | 000534-000633 |
| **Type Conv** | toString, atoi, int, int64, float64, toDecimal, toStrings | 000634-000703 |
| **Encoding** | b64enc, b64dec, b32enc, b32dec | 000704-000743 |
| **List - Basic** | list, tuple, first, mustFirst, rest, mustRest | 000744-000803 |
| **List - Access** | last, mustLast, initial, mustInitial | 000804-000843 |
| **List - Add** | append, mustAppend, push, mustPush, prepend, mustPrepend | 000844-000903 |
| **List - Transform** | concat, reverse, mustReverse, uniq, mustUniq | 000904-000953 |
| **List - Filter** | without, mustWithout, has, mustHas | 000954-000993 |
| **List - Slice** | compact, mustCompact, slice, mustSlice, chunk, mustChunk, sortAlpha | 000994-001063 |
| **String List** | split, splitList, splitn, join | 001064-001103 |
| **Dict - Basic** | dict, get, set, unset, hasKey | 001104-001153 |
| **Dict - Access** | pluck, keys, pick, omit, values | 001154-001203 |
| **Dict - Merge** | merge, mustMerge, mergeOverwrite, mustMergeOverwrite | 001204-001243 |
| **Dict - Utility** | deepCopy, mustDeepCopy, dig | 001244-001273 |
| **Default/Logic** | default, empty, coalesce, all, any, ternary | 001274-001333 |
| **JSON - Basic** | fromJson, mustFromJson, toJson, mustToJson | 001334-001373 |
| **JSON - Format** | toPrettyJson, mustToPrettyJson, toRawJson, mustToRawJson | 001374-001413 |
| **Reflection** | typeOf, typeIs, typeIsLike, kindOf, kindIs, deepEqual | 001414-001473 |
| **Path** | base, dir, clean, ext, isAbs | 001474-001523 |
| **Path OS** | osBase, osDir, osClean, osExt, osIsAbs | 001524-001573 |
| **OS** | env, expandenv | 001574-001593 |
| **UUID** | uuidv4 | 001594-001603 |
| **Semver** | semver, semverCompare | 001604-001623 |
| **Regex - Match** | regexMatch, mustRegexMatch, regexFind, mustRegexFind | 001624-001663 |
| **Regex - Find/Replace** | regexFindAll, mustRegexFindAll, regexReplaceAll, mustRegexReplaceAll | 001664-001703 |
| **Regex - Split/Quote** | regexReplaceAllLiteral, mustRegexReplaceAllLiteral, regexSplit, mustRegexSplit, regexQuoteMeta | 001704-001753 |
| **URL** | urlParse, urlJoin | 001754-001773 |
| **Crypto - Basic** | bcrypt, htpasswd, randBytes | 001774-001803 |
| **Crypto - Keys** | genPrivateKey, derivePassword | 001804-001823 |
| **Crypto - AES** | encryptAES, decryptAES | 001824-001843 |
| **Crypto - CA** | genCA, genCAWithKey | 001844-001853 |
| **Crypto - Self-Signed** | genSelfSignedCert, genSelfSignedCertWithKey | 001854-001863 |
| **Crypto - Signed** | genSignedCert, genSignedCertWithKey | 001864-001873 |
| **Date - Basic** | now, ago, date | 001874-001903 |
| **Date - Zone** | dateInZone, date_in_zone | 001904-001923 |
| **Date - Modify** | dateModify, date_modify, mustDateModify, must_date_modify | 001924-001963 |
| **Date - HTML** | htmlDate, htmlDateInZone | 001964-001983 |
| **Date - Parse** | toDate, mustToDate | 001984-002003 |
| **Date - Epoch/Duration** | unixEpoch, duration, durationRound | 002004-002033 |
| **Utility** | hello | 002034-002043 |
| **Network** | getHostByName | 002044-002046 |
| **Crypto - Custom** | buildCustomCert | 002050-002053 |
| **Flow Control** | fail | (commented out - causes errors) |

### Complete Function List (170+ functions)

**String Functions (41):**
`trim`, `trimAll`, `trimall`, `trimPrefix`, `trimSuffix`, `upper`, `lower`, `title`, `untitle`, `repeat`, `substr`, `nospace`, `trunc`, `abbrev`, `abbrevboth`, `initials`, `wrap`, `wrapWith`, `snakecase`, `camelcase`, `kebabcase`, `swapcase`,
`contains`, `hasPrefix`, `hasSuffix`, `quote`, `squote`, `cat`, `replace`, `indent`, `nindent`, `plural`, `shuffle`, `randAlphaNum`, `randAlpha`, `randAscii`, `randNumeric`, `split`, `splitList`, `splitn`, `join`

**Math Functions (23):**
`add`, `add1`, `sub`, `div`, `mod`, `mul`, `max`, `min`, `biggest`, `randInt`, `until`, `untilStep`, `seq`, `addf`, `add1f`, `subf`, `divf`, `mulf`, `maxf`, `minf`, `floor`, `ceil`, `round`

**Type/Encoding Functions (11):**
`toString`, `atoi`, `int`, `int64`, `float64`, `toDecimal`, `toStrings`, `b64enc`, `b64dec`, `b32enc`, `b32dec`

**Hash Functions (4):**
`sha1sum`, `sha256sum`, `sha512sum`, `adler32sum`

**List Functions (32):**
`list`, `tuple`, `first`, `mustFirst`, `rest`, `mustRest`, `last`, `mustLast`, `initial`, `mustInitial`, `append`, `mustAppend`, `push`, `mustPush`, `prepend`, `mustPrepend`, `concat`, `reverse`, `mustReverse`, `uniq`, `mustUniq`, `without`,
`mustWithout`, `has`, `mustHas`, `compact`, `mustCompact`, `slice`, `mustSlice`, `chunk`, `mustChunk`, `sortAlpha`

**Dict Functions (17):**
`dict`, `get`, `set`, `unset`, `hasKey`, `pluck`, `keys`, `pick`, `omit`, `values`, `merge`, `mustMerge`, `mergeOverwrite`, `mustMergeOverwrite`, `deepCopy`, `mustDeepCopy`, `dig`

**Default/Logic Functions (6):**
`default`, `empty`, `coalesce`, `all`, `any`, `ternary`

**JSON Functions (8):**
`fromJson`, `mustFromJson`, `toJson`, `mustToJson`, `toPrettyJson`, `mustToPrettyJson`, `toRawJson`, `mustToRawJson`

**Reflection Functions (6):**
`typeOf`, `typeIs`, `typeIsLike`, `kindOf`, `kindIs`, `deepEqual`

**Path Functions (10):**
`base`, `dir`, `clean`, `ext`, `isAbs`, `osBase`, `osDir`, `osClean`, `osExt`, `osIsAbs`

**OS Functions (2):**
`env`, `expandenv`

**Crypto Functions (14):**
`bcrypt`, `htpasswd`, `randBytes`, `genPrivateKey`, `derivePassword`, `encryptAES`, `decryptAES`, `buildCustomCert`, `genCA`, `genCAWithKey`, `genSelfSignedCert`, `genSelfSignedCertWithKey`, `genSignedCert`, `genSignedCertWithKey`

**Date Functions (16):**
`now`, `ago`, `date`, `dateInZone`, `date_in_zone`, `dateModify`, `date_modify`, `mustDateModify`, `must_date_modify`, `htmlDate`, `htmlDateInZone`, `toDate`, `mustToDate`, `unixEpoch`, `duration`, `durationRound`

**Regex Functions (13):**
`regexMatch`, `mustRegexMatch`, `regexFind`, `mustRegexFind`, `regexFindAll`, `mustRegexFindAll`, `regexReplaceAll`, `mustRegexReplaceAll`, `regexReplaceAllLiteral`, `mustRegexReplaceAllLiteral`, `regexSplit`, `mustRegexSplit`, `regexQuoteMeta`

**Semver Functions (2):**
`semver`, `semverCompare`

**URL Functions (2):**
`urlParse`, `urlJoin`

**UUID Functions (1):**
`uuidv4`

**Network Functions (1):**
`getHostByName`

**Flow Control Functions (1):**
`fail`

**Utility Functions (1):**
`hello`






Deterministic vs Non-Deterministic Functions

DETERMINISTIC Functions (produce same output every run):
┌─────────────────────┬───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│      Category       │                                                               Functions                                                               │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ String Trim         │ trim, trimAll, trimall, trimPrefix, trimSuffix                                                                                        │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ String Case         │ upper, lower, title, untitle, snakecase, camelcase, kebabcase, swapcase                                                               │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ String Manipulation │ repeat, substr, nospace, trunc, abbrev, abbrevboth, initials, wrap, wrapWith                                                          │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ String Search       │ contains, hasPrefix, hasSuffix                                                                                                        │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ String Format       │ quote, squote, cat, replace, indent, nindent, plural                                                                                  │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Hash                │ sha1sum, sha256sum, sha512sum, adler32sum                                                                                             │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Math (Int)          │ add, add1, sub, div, mod, mul, max, min, biggest                                                                                      │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Math (Float)        │ addf, add1f, subf, divf, mulf, maxf, minf, floor, ceil, round                                                                         │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Sequences           │ until, untilStep, seq                                                                                                                 │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Type Conversion     │ toString, atoi, int, int64, float64, toDecimal, toStrings                                                                             │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Encoding            │ b64enc, b64dec, b32enc, b32dec                                                                                                        │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ List                │ list, tuple, first, rest, last, initial, append, prepend, push, concat, reverse, uniq, without, has, compact, slice, chunk, sortAlpha │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ String List         │ split, splitList, splitn, join                                                                                                        │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Dict                │ dict, get, set, unset, hasKey, pluck, keys, pick, omit, values, merge, mergeOverwrite, deepCopy, dig                                  │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Default/Logic       │ default, empty, coalesce, all, any, ternary                                                                                           │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ JSON                │ fromJson, toJson, toPrettyJson, toRawJson (and must* variants)                                                                        │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Reflection          │ typeOf, typeIs, typeIsLike, kindOf, kindIs, deepEqual                                                                                 │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Path                │ base, dir, clean, ext, isAbs, osBase, osDir, osClean, osExt, osIsAbs                                                                  │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Semver              │ semver, semverCompare                                                                                                                 │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Regex               │ regexMatch, regexFind, regexFindAll, regexReplaceAll, regexReplaceAllLiteral, regexSplit, regexQuoteMeta (and must* variants)         │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ URL                 │ urlParse, urlJoin                                                                                                                     │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Duration            │ duration, durationRound                                                                                                               │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Utility             │ hello                                                                                                                                 │
├─────────────────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ Fixed Date Parse    │ toDate, mustToDate                                                                                                                    │
└─────────────────────┴───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
NON-DETERMINISTIC Functions (different output each run):
┌───────────────┬──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┬───────────────────────────────────────────────────────────────┐
│   Category    │                                                                Functions                                                                 │                            Reason                             │
├───────────────┼──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼───────────────────────────────────────────────────────────────┤
│ Random String │ randAlphaNum, randAlpha, randAscii, randNumeric                                                                                          │ Random generation                                             │
├───────────────┼──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼───────────────────────────────────────────────────────────────┤
│ Random Number │ randInt                                                                                                                                  │ Random generation                                             │
├───────────────┼──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼───────────────────────────────────────────────────────────────┤
│ Random Bytes  │ randBytes                                                                                                                                │ Random generation                                             │
├───────────────┼──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼───────────────────────────────────────────────────────────────┤
│ Shuffle       │ shuffle                                                                                                                                  │ Random reordering                                             │
├───────────────┼──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼───────────────────────────────────────────────────────────────┤
│ UUID          │ uuidv4                                                                                                                                   │ Random UUID generation                                        │
├───────────────┼──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼───────────────────────────────────────────────────────────────┤
│ Crypto Hash   │ bcrypt, htpasswd                                                                                                                         │ Random salt                                                   │
├───────────────┼──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼───────────────────────────────────────────────────────────────┤
│ Crypto Keys   │ genPrivateKey                                                                                                                            │ Random key generation                                         │
├───────────────┼──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼───────────────────────────────────────────────────────────────┤
│ Crypto Certs  │ genCA, genCAWithKey, genSelfSignedCert, genSelfSignedCertWithKey, genSignedCert, genSignedCertWithKey, buildCustomCert                   │ Random key material                                           │
├───────────────┼──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼───────────────────────────────────────────────────────────────┤
│ AES           │ encryptAES                                                                                                                               │ Random IV                                                     │
├───────────────┼──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼───────────────────────────────────────────────────────────────┤
│ Password      │ derivePassword                                                                                                                           │ Algorithm produces deterministic output but depends on inputs │
├───────────────┼──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┼───────────────────────────────────────────────────────────────┤
│ Time          │ now, ago, date, dateInZone, date_in_zone, dateModify, date_modify, mustDateModify, must_date_modify, htmlDate, htmlDateInZone, unixEpoch │ Current time dependent                                        │
└───────────────┴──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┴───────────────────────────────────────────────────────────────┘
ENVIRONMENT-DEPENDENT Functions:
┌────────────────┬─────────────────────────────┐
│   Functions    │           Reason            │
├────────────────┼─────────────────────────────┤
│ env, expandenv │ Reads environment variables │
├────────────────┼─────────────────────────────┤
│ getHostByName  │ DNS resolution              │
└────────────────┴─────────────────────────────┘


Real Compatibility Differences Found

The test failures now reveal actual differences between sprig and sprigin:
┌────────────────────────┬───────────────────┬─────────────────────┐
│         Issue          │       Sprig       │       Sprigin       │
├────────────────────────┼───────────────────┼─────────────────────┤
│ title "HELLO WORLD"    │ HELLO WORLD       │ Hello World         │
├────────────────────────┼───────────────────┼─────────────────────┤
│ camelcase "UPPER_CASE" │ UpperCase         │ UPPERCASE           │
├────────────────────────┼───────────────────┼─────────────────────┤
│ keys/values order      │ Non-deterministic │ Non-deterministic   │
├────────────────────────┼───────────────────┼─────────────────────┤
│ b64dec invalid input   │ Error message     │ Empty string        │
├────────────────────────┼───────────────────┼─────────────────────┤
│ kindOf nil             │ invalid           │ Empty               │
├────────────────────────┼───────────────────┼─────────────────────┤
│ Float precision        │ 0.3               │ 0.30000000000000004 │
├────────────────────────┼───────────────────┼─────────────────────┤
│ Date timezone          │ Local (CEST)      │ UTC                 │
└────────────────────────┴───────────────────┴─────────────────────┘


[CODE_001451]    {{ kindOf nil }} // IGNORE: invalid are a bug of sprig
[CODE_000067]    {{ title "HELLO WORLD" }} // IGNORE: title are camelcase
[CODE_000073]    {{ title "it's a test" }} // Bug resolve
[CODE_000112]    {{ nospace "α β γ" }} // fix bug with special chars
[CODE_000131]    {{ abbrev 6 "αβγδεζ" }} // fix bug with special chars
[CODE_000143]    {{ abbrevboth 3 12 "αβγδεζηθικλμ" }} // fx bug with special chars
[CODE_000141]    {{ abbrevboth 1 10 "testing string here" }} // fix bug on abbrvboth not working at left