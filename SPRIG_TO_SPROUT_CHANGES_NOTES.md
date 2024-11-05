# Migration Notes for Sprout Library
This document outlines the key differences and migration changes between the
Sprig and Sprout libraries. The changes are designed to enhance stability and 
usability in the Sprout library.

This document will help contributors and maintainers understand the changes made
between the fork date and version 1.0.0 of the Sprout library. 

It will be updated to reflect changes in future versions of the library. 

This document will also assist in creating the migration guide when version 1.0 is ready.

Any changes between the fork date and version 1.0.0 of the Sprout library not 
migrated yet to the [Migration from Sprig](https://docs.atom.codes/sprout/migration-from-sprig) will be documented here.


## Templating Differences
### Bugs fixed
In Sprig, some functions have bugs that have been fixed in Sprout:

1. `{{ "foooboooooo" | abbrevboth 4 9 }}`
  - Sprig: `fooobo...`
  - Sprout: `...boo...`
2. `{{ "FoO  bar" | camelcase }}` | `{{ "foo  bar" | camelcase }}`
  - Sprig: `FoO Bar` | `Foo Bar`
  - Sprout: `FoOBar` | `FooBar`
3. `{{ camelcase "___complex__case_" }}` | `{{ camelcase "_camel_case" }}`
  - Sprig: `___Complex_Case_` | `_CamelCase`
  - Sprout: `ComplexCase` | `CamelCase`
4. `{{ "foo  bar" | kebabcase }}`
  - Sprig: `foo--bar`
  - Sprout: `foo-bar`
5. `{{ "foo  bar" | snakecase }}`
  - Sprig: `foo__bar`
  - Sprout: `foo_bar`
6. `{{ snakecase "Duration2m3s" }}` | `{{ snakecase "Bld4Floor3rd" }}`
  - Sprig: `duration_2m3s` | `bld4_floor_3rd`
  - Sprout: `duration_2m_3s` | `bld_4floor_3rd`
7. `{{ "foobar" | substr 0 -3 }}` | `{{ "foobar" | substr -3 6 }}`
  - Sprig: `foobar` | `foobar`
  - Sprout: `foo` | `bar`
8. `{{ .duration | durationRound }}`
  - Sprig: `0s` (only when not cast to `time.Duration` before, `{{ .duration | duration | durationRound }}` works as expected with `5s`)
  - Sprout: `5s`
