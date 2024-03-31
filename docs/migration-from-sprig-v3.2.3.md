---
description: A compatibility table between sprig v3.2.3 and sprout v1.0
---

# 👷 Migration from sprig v3.2.3

{% hint style="info" %}
:white\_check\_mark: Implemented - The implementation will be available through v1

:hourglass:To rework - This function are present but aren't reworked with Sprout vision

:warning: Deprecated - This function are deprecated and not be migrated to sprout

:heavy\_multiplication\_x: Will not be implemented - This function will not be migrated to sprout v1

:x: Dont exist - This function don't exist on this version

:man\_shrugging: TBD : To Be Determinate for v1

***

This page will be updated each time function are re-implemented correctly in Sprout v1
{% endhint %}



| Function                   | Sprig v3.2.3               | Spout v1.0 (to define)                    | Notes                  |
| -------------------------- | -------------------------- | ----------------------------------------- | ---------------------- |
| `hello`                    | :white\_check\_mark:       | :white\_check\_mark:                      |                        |
| `ago`                      | :white\_check\_mark:       | :hourglass:                               |                        |
| `date`                     | :white\_check\_mark:       | :hourglass:                               |                        |
| `date_modify`              | :warning: DEPRECATED       | :heavy\_multiplication\_x:                |                        |
| `dateModify`               | :white\_check\_mark:       | :hourglass:                               |                        |
| `date_in_zone`             | :warning: DEPRECATED       | :heavy\_multiplication\_x:                |                        |
| `dateInZone`               | :white\_check\_mark:       | :hourglass:                               |                        |
| duration                   | :white\_check\_mark:       | :hourglass:                               |                        |
| durationRound              | :white\_check\_mark:       | :hourglass:                               |                        |
| htmlDate                   | :white\_check\_mark:       | :hourglass:                               |                        |
| htmlDateInZone             | :white\_check\_mark:       | :hourglass:                               |                        |
| must\_date\_modify         | :warning: DEPRECATED       | :heavy\_multiplication\_x:                |                        |
| mustDateModify             | :white\_check\_mark:       | :hourglass:                               |                        |
| mustToDate                 | :white\_check\_mark:       | :hourglass:                               |                        |
| now                        | :white\_check\_mark:       | :hourglass:                               |                        |
| toDate                     | :white\_check\_mark:       | :hourglass:                               |                        |
| unixEpoch                  | :white\_check\_mark:       | :hourglass:                               |                        |
| abbrev                     | :white\_check\_mark:       | :heavy\_multiplication\_x: (ellipsis)     |                        |
| ellipsis                   | :heavy\_multiplication\_x: | :white\_check\_mark:                      |                        |
| abbrevboth                 | :white\_check\_mark:       | :heavy\_multiplication\_x: (ellipsisBoth) |                        |
| ellipsisBoth               | :heavy\_multiplication\_x: | :white\_check\_mark:                      |                        |
| trunc                      | :white\_check\_mark:       | :hourglass:                               |                        |
| trim                       | :white\_check\_mark:       | :hourglass:                               |                        |
| upper                      | :white\_check\_mark:       | :hourglass:                               |                        |
| lower                      | :white\_check\_mark:       | :hourglass:                               |                        |
| title                      | :white\_check\_mark:       | :hourglass:                               |                        |
| untitle                    | :white\_check\_mark:       | :hourglass:                               |                        |
| substr                     | :white\_check\_mark:       | :hourglass:                               |                        |
| repeat                     | :white\_check\_mark:       | :hourglass:                               |                        |
| trimall                    | :warning: DEPRECATED       | :heavy\_multiplication\_x: (trimAll)      |                        |
| trimAll                    | :white\_check\_mark:       | :hourglass:                               |                        |
| trimSuffix                 | :white\_check\_mark:       | :hourglass:                               |                        |
| trimPrefix                 | :white\_check\_mark:       | :hourglass:                               |                        |
| nospace                    | :white\_check\_mark:       | :white\_check\_mark:                      |                        |
| initials                   | :white\_check\_mark:       | :white\_check\_mark:                      |                        |
| randAlphaNum               | :white\_check\_mark:       | :white\_check\_mark:                      |                        |
| randAlpha                  | :white\_check\_mark:       | :white\_check\_mark:                      |                        |
| randAscii                  | :white\_check\_mark:       | :white\_check\_mark:                      |                        |
| randNumeric                | :white\_check\_mark:       | :white\_check\_mark:                      |                        |
| swapcase                   | :white\_check\_mark:       | :white\_check\_mark:                      |                        |
| shuffle                    | :white\_check\_mark:       | :hourglass:                               |                        |
| snakecase                  | :white\_check\_mark:       | :hourglass:                               |                        |
| camelcase                  | :white\_check\_mark:       | :hourglass:                               |                        |
| kebabcase                  | :white\_check\_mark:       | :hourglass:                               |                        |
| wrap                       | :white\_check\_mark:       | :hourglass:                               |                        |
| wrapWith                   | :white\_check\_mark:       | :hourglass:                               |                        |
| contains                   | :white\_check\_mark:       | :hourglass:                               |                        |
| hasPrefix                  | :white\_check\_mark:       | :hourglass:                               |                        |
| hasSuffix                  | :white\_check\_mark:       | :hourglass:                               |                        |
| quote                      | :white\_check\_mark:       | :hourglass:                               |                        |
| squote                     | :white\_check\_mark:       | :hourglass:                               |                        |
| cat                        | :white\_check\_mark:       | :man\_shrugging: TBD                      |                        |
| indent                     | :white\_check\_mark:       | :hourglass:                               |                        |
| nindent                    | :white\_check\_mark:       | :hourglass:                               |                        |
| replace                    | :white\_check\_mark:       | :hourglass:                               |                        |
| plural                     | :white\_check\_mark:       | :hourglass:                               |                        |
| sha1sum                    | :white\_check\_mark:       | :hourglass:                               |                        |
| sha256sum                  | :white\_check\_mark:       | :hourglass:                               |                        |
| adler32sum                 | :white\_check\_mark:       | :hourglass:                               |                        |
| toString                   | :white\_check\_mark:       | :hourglass:                               |                        |
| atoi                       | :white\_check\_mark:       | :heavy\_multiplication\_x: (toInt)        |                        |
| int64                      | :white\_check\_mark:       | :heavy\_multiplication\_x: (toInt64)      |                        |
| int                        | :white\_check\_mark:       | :heavy\_multiplication\_x: (toInt)        |                        |
| float64                    | :white\_check\_mark:       | :heavy\_multiplication\_x: (toFloat64)    |                        |
| seq                        | :white\_check\_mark:       | :hourglass:                               |                        |
| toDecimal                  | :white\_check\_mark:       | :man\_shrugging: TBD                      | From octale to decimal |
| _gt_                       | :x:                        | :man\_shrugging: TBD                      |                        |
| gte                        | :x:                        | :man\_shrugging: TBD                      |                        |
| lt                         | :x:                        | :man\_shrugging: TBD                      |                        |
| lte                        | :x:                        | :man\_shrugging: TBD                      |                        |
| split                      | :white\_check\_mark:       | :hourglass:                               |                        |
| splitList                  | :white\_check\_mark:       | :hourglass:                               |                        |
| splitn                     | :white\_check\_mark:       | :hourglass:                               |                        |
| toStrings                  | :white\_check\_mark:       | :man\_shrugging: TBD                      |                        |
| until                      | :white\_check\_mark:       | :hourglass:                               |                        |
| untilStep                  | :white\_check\_mark:       | :hourglass:                               |                        |
| add1                       | :white\_check\_mark:       | :hourglass:                               |                        |
| add                        | :white\_check\_mark:       | :hourglass:                               |                        |
| sub                        | :white\_check\_mark:       | :hourglass:                               |                        |
| div                        | :white\_check\_mark:       | :hourglass:                               |                        |
| mod                        | :white\_check\_mark:       | :hourglass:                               |                        |
| mul                        | :white\_check\_mark:       | :hourglass:                               |                        |
| randInt                    | :white\_check\_mark:       | :hourglass:                               |                        |
| add1f                      | :white\_check\_mark:       | :heavy\_multiplication\_x: (add1)         |                        |
| addf                       | :white\_check\_mark:       | :heavy\_multiplication\_x: (add)          |                        |
| subf                       | :white\_check\_mark:       | :heavy\_multiplication\_x: (sub)          |                        |
| divf                       | :white\_check\_mark:       | :heavy\_multiplication\_x: (div)          |                        |
| mulf                       | :white\_check\_mark:       | :heavy\_multiplication\_x: (mul)          |                        |
| biggest                    | :white\_check\_mark:       | :heavy\_multiplication\_x: (max)          |                        |
| max                        | :white\_check\_mark:       | :hourglass:                               |                        |
| min                        | :white\_check\_mark:       | :hourglass:                               |                        |
| maxf                       | :white\_check\_mark:       | :heavy\_multiplication\_x: (max)          |                        |
| minf                       | :white\_check\_mark:       | :heavy\_multiplication\_x: (min)          |                        |
| ceil                       | :white\_check\_mark:       | :hourglass:                               |                        |
| floor                      | :white\_check\_mark:       | :hourglass:                               |                        |
| round                      | :white\_check\_mark:       | :hourglass:                               |                        |
| join                       | :white\_check\_mark:       | :hourglass:                               |                        |
| sortAlpha                  | :white\_check\_mark:       | :hourglass:                               |                        |
| default                    | :white\_check\_mark:       | :hourglass:                               |                        |
| empty                      | :white\_check\_mark:       | :hourglass:                               |                        |
| coalesce                   | :white\_check\_mark:       | :hourglass:                               |                        |
| all                        | :white\_check\_mark:       | :hourglass:                               |                        |
| any                        | :white\_check\_mark:       | :hourglass:                               |                        |
| compact                    | :white\_check\_mark:       | :hourglass:                               |                        |
| mustCompact                | :white\_check\_mark:       | :hourglass:                               |                        |
| fromJson                   | :white\_check\_mark:       | :hourglass:                               |                        |
| toJson                     | :white\_check\_mark:       | :hourglass:                               |                        |
| toPrettyJson               | :white\_check\_mark:       | :hourglass:                               |                        |
| toRawJson                  | :white\_check\_mark:       | :hourglass:                               |                        |
| mustFromJson               | :white\_check\_mark:       | :hourglass:                               |                        |
| mustToJson                 | :white\_check\_mark:       | :hourglass:                               |                        |
| mustToPrettyJson           | :white\_check\_mark:       | :hourglass:                               |                        |
| mustToRawJson              | :white\_check\_mark:       | :hourglass:                               |                        |
| ternary                    | :white\_check\_mark:       | :hourglass:                               |                        |
| deepCopy                   | :white\_check\_mark:       | :hourglass:                               |                        |
| mustDeepCopy               | :white\_check\_mark:       | :hourglass:                               |                        |
| typeOf                     | :white\_check\_mark:       | :hourglass:                               |                        |
| typeIs                     | :white\_check\_mark:       | :hourglass:                               |                        |
| typeIsLike                 | :white\_check\_mark:       | :hourglass:                               |                        |
| kindOf                     | :white\_check\_mark:       | :hourglass:                               |                        |
| kindIs                     | :white\_check\_mark:       | :hourglass:                               |                        |
| deepEqual                  | :white\_check\_mark:       | :hourglass:                               |                        |
| env                        | :white\_check\_mark:       | :hourglass:                               |                        |
| expandenv                  | :white\_check\_mark:       | :hourglass:                               |                        |
| getHostByName              | :white\_check\_mark:       | :man\_shrugging: TBD                      |                        |
| base                       | :white\_check\_mark:       | :hourglass:                               |                        |
| dir                        | :white\_check\_mark:       | :hourglass:                               |                        |
| clean                      | :white\_check\_mark:       | :hourglass:                               |                        |
| ext                        | :white\_check\_mark:       | :hourglass:                               |                        |
| isAbs                      | :white\_check\_mark:       | :hourglass:                               |                        |
| osBase                     | :white\_check\_mark:       | :hourglass:                               |                        |
| osClean                    | :white\_check\_mark:       | :hourglass:                               |                        |
| osDir                      | :white\_check\_mark:       | :hourglass:                               |                        |
| osExt                      | :white\_check\_mark:       | :hourglass:                               |                        |
| osIsAbs                    | :white\_check\_mark:       | :hourglass:                               |                        |
| b64enc                     | :white\_check\_mark:       | :hourglass:                               |                        |
| b64dec                     | :white\_check\_mark:       | :hourglass:                               |                        |
| b63enc                     | :white\_check\_mark:       | :hourglass:                               |                        |
| b63dec                     | :white\_check\_mark:       | :hourglass:                               |                        |
| tuple                      | :white\_check\_mark:       | :hourglass:                               |                        |
| list                       | :white\_check\_mark:       | :hourglass:                               |                        |
| dict                       | :white\_check\_mark:       | :hourglass:                               |                        |
| get                        | :white\_check\_mark:       | :hourglass:                               |                        |
| set                        | :white\_check\_mark:       | :hourglass:                               |                        |
| unset                      | :white\_check\_mark:       | :hourglass:                               |                        |
| hasKey                     | :white\_check\_mark:       | :hourglass:                               |                        |
| pluck                      | :white\_check\_mark:       | :hourglass:                               |                        |
| keys                       | :white\_check\_mark:       | :hourglass:                               |                        |
| pick                       | :white\_check\_mark:       | :hourglass:                               |                        |
| omit                       | :white\_check\_mark:       | :hourglass:                               |                        |
| merge                      | :white\_check\_mark:       | :hourglass:                               |                        |
| mergeOverwrite             | :white\_check\_mark:       | :hourglass:                               |                        |
| mustMerge                  | :white\_check\_mark:       | :hourglass:                               |                        |
| mustMergeOverwrite         | :white\_check\_mark:       | :hourglass:                               |                        |
| values                     | :white\_check\_mark:       | :hourglass:                               |                        |
| append                     | :white\_check\_mark:       | :hourglass:                               |                        |
| push                       | :white\_check\_mark:       | :hourglass:                               |                        |
| mustAppend                 | :white\_check\_mark:       | :hourglass:                               |                        |
| mustPush                   | :white\_check\_mark:       | :hourglass:                               |                        |
| prepend                    | :white\_check\_mark:       | :hourglass:                               |                        |
| mustPrepend                | :white\_check\_mark:       | :hourglass:                               |                        |
| first                      | :white\_check\_mark:       | :hourglass:                               |                        |
| mustFirst                  | :white\_check\_mark:       | :hourglass:                               |                        |
| rest                       | :white\_check\_mark:       | :hourglass:                               |                        |
| mustRest                   | :white\_check\_mark:       | :hourglass:                               |                        |
| last                       | :white\_check\_mark:       | :hourglass:                               |                        |
| mustLast                   | :white\_check\_mark:       | :hourglass:                               |                        |
| initial                    | :white\_check\_mark:       | :hourglass:                               |                        |
| mustInitial                | :white\_check\_mark:       | :hourglass:                               |                        |
| reverse                    | :white\_check\_mark:       | :hourglass:                               |                        |
| mustReverse                | :white\_check\_mark:       | :hourglass:                               |                        |
| uniq                       | :white\_check\_mark:       | :hourglass:                               |                        |
| mustUniq                   | :white\_check\_mark:       | :hourglass:                               |                        |
| without                    | :white\_check\_mark:       | :hourglass:                               |                        |
| mustWithout                | :white\_check\_mark:       | :hourglass:                               |                        |
| has                        | :white\_check\_mark:       | :hourglass:                               |                        |
| mustHas                    | :white\_check\_mark:       | :hourglass:                               |                        |
| slice                      | :white\_check\_mark:       | :hourglass:                               |                        |
| mustSlice                  | :white\_check\_mark:       | :hourglass:                               |                        |
| concat                     | :white\_check\_mark:       | :hourglass:                               |                        |
| dig                        | :white\_check\_mark:       | :hourglass:                               |                        |
| chunk                      | :white\_check\_mark:       | :hourglass:                               |                        |
| mustChunk                  | :white\_check\_mark:       | :hourglass:                               |                        |
| bcrypt                     | :white\_check\_mark:       | :hourglass:                               |                        |
| htpasswd                   | :white\_check\_mark:       | :hourglass:                               |                        |
| genPrivateKey              | :white\_check\_mark:       | :man\_shrugging: TBD                      |                        |
| derivePassword             | :white\_check\_mark:       | :man\_shrugging: TBD                      |                        |
| buildCustomCert            | :white\_check\_mark:       | :man\_shrugging: TBD                      |                        |
| genCA                      | :white\_check\_mark:       | :man\_shrugging: TBD                      |                        |
| genCAWithKey               | :white\_check\_mark:       | :man\_shrugging: TBD                      |                        |
| genSelfSignedCert          | :white\_check\_mark:       | :man\_shrugging: TBD                      |                        |
| genSelfSignedCertWithKey   | :white\_check\_mark:       | :man\_shrugging: TBD                      |                        |
| genSignedCert              | :white\_check\_mark:       | :man\_shrugging: TBD                      |                        |
| genSignedCertWithKey       | :white\_check\_mark:       | :man\_shrugging: TBD                      |                        |
| encryptAES                 | :white\_check\_mark:       | :man\_shrugging: TBD                      |                        |
| decryptAES                 | :white\_check\_mark:       | :man\_shrugging: TBD                      |                        |
| randBytes                  | :white\_check\_mark:       | :hourglass:                               |                        |
| uuidv4                     | :white\_check\_mark:       | :hourglass:                               |                        |
| semver                     | :white\_check\_mark:       | :man\_shrugging: TBD                      |                        |
| semverCompare              | :white\_check\_mark:       | :man\_shrugging: TBD                      |                        |
| fail                       | :white\_check\_mark:       | :man\_shrugging: TBD                      |                        |
| regexMatch                 | :white\_check\_mark:       | :hourglass:                               |                        |
| mustRegexMatch             | :white\_check\_mark:       | :hourglass:                               |                        |
| regexFindAll               | :white\_check\_mark:       | :hourglass:                               |                        |
| mustRegexFindAll           | :white\_check\_mark:       | :hourglass:                               |                        |
| regexFind                  | :white\_check\_mark:       | :hourglass:                               |                        |
| mustRegexFind              | :white\_check\_mark:       | :hourglass:                               |                        |
| regexReplaceAll            | :white\_check\_mark:       | :hourglass:                               |                        |
| mustRegexReplaceAll        | :white\_check\_mark:       | :hourglass:                               |                        |
| regexReplaceAllLiteral     | :white\_check\_mark:       | :hourglass:                               |                        |
| mustRegexReplaceAllLiteral | :white\_check\_mark:       | :hourglass:                               |                        |
| regexSplit                 | :white\_check\_mark:       | :hourglass:                               |                        |
| mustRegexSplit             | :white\_check\_mark:       | :hourglass:                               |                        |
| regexQuoteMeta             | :white\_check\_mark:       | :hourglass:                               |                        |
| urlParse                   | :white\_check\_mark:       | :hourglass:                               |                        |
| urlJoin                    | :white\_check\_mark:       | :hourglass:                               |                        |

