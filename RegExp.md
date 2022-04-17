# Regular Expression Tutorial

Reference: [Wiki](https://en.wikipedia.org/wiki/Regular_expression), [Zhihu](https://zhuanlan.zhihu.com/p/308626540)

## Syntax

### Basic Syntax

- Boolean “or”: `grey|gray`

- Grouping: `gr(a|e)y` 

- Quantifier: after an element (such as a token, character, or group)

  | Metacharacter |       Description       |
  | :-----------: | :---------------------: |
  |      `?`      |      once or never      |
  |      `*`      | some times (0 included) |
  |      `+`      |      once or more       |
  |     `{n}`     |        `n` times        |
  |   `{min,}`    |    least `min` times    |
  |   `{,max}`    |    most `max ` times    |
  |  `{min,max}`  |   `min` - `max` times   |

  

- Wildcard: `.` matches ***any*** character.

### Delimiters

> When entering a regex in a programming language, they may be represented as a usual string literal, hence usually quoted; this is common in C, Java, and Python for instance, where the regex `re` is entered as `"re"`. However, they are often written with slashes as [delimiters](https://en.wikipedia.org/wiki/Delimiter), as in `/re/` for the regex `re`. This originates in [ed](https://en.wikipedia.org/wiki/Ed_(text_editor)), where `/` is the editor command for searching, and an expression `/re/` can be used to specify a range of lines (matching the pattern), which can be combined with other commands on either side, most famously `g/re/p` as in [grep](https://en.wikipedia.org/wiki/Grep) ("global regex print"), which is included in most [Unix](https://en.wikipedia.org/wiki/Unix)-based operating systems, such as [Linux](https://en.wikipedia.org/wiki/Linux) distributions. A similar convention is used in [sed](https://en.wikipedia.org/wiki/Sed), where search and replace is given by `s/re/replacement/` and patterns can be joined with a comma to specify a range of lines as in `/re1/,/re2/`. This notation is particularly well known due to its use in [Perl](https://en.wikipedia.org/wiki/Perl), where it forms part of the syntax distinct from normal string literals. In some cases, such as sed and Perl, alternative delimiters can be used to avoid collision with contents, and to avoid having to escape occurrences of the delimiter character in the contents. For example, in sed the command `s,/,X,` will replace a `/` with an `X`, using commas as delimiters.

### Metacharacter

| Metacharacter |                         Description                          |
| :-----------: | :----------------------------------------------------------: |
|      `^`      |    Matches the starting ***position*** within the string.    |
|      `.`      | Matches a character, but in POSIX bracket, it matches “.” itself. |
|`[]`|A bracket expression. Matches a single character that is contained within the brackets.|
|     `[^]`     | Matches a single character that is not contained within the brackets. |
|      `$`      | Matches the ending ***position*** of the string or the ***position*** just before a string-ending newline. |
|     `()`      |               Defines a marked subexpression.                |
|     `\n`      | Matches what the *n*th marked subexpression matched, where `n` is a digit from 1 to 9. |
|      `*`      |      Matches the preceding element zero or more times.       |
|    `{m,n}`    | Matches the preceding element at least `m` and not more than `n` times |

| Metacharacter |       Description       |
| :-----------: | :---------------------: |
|      `?`      |      once or never      |
|      `*`      | some times (0 included) |
|      `+`      |      once or more       |
|     `{n}`     |        `n` times        |
|   `{min,}`    |    least `min` times    |
|   `{,max}`    |    most `max ` times    |
|  `{min,max}`  |   `min` - `max` times   |

| Metacharacter |                         Description                          |
| :-----------: | :----------------------------------------------------------: |
|      `?`      | Modifies the `*`, `+`, `?` or `{M,N}`'d regex that comes before to match as few times as possible. |
|     `\b`      | Matches a zero-width boundary between a word-class character (see next) and either a non-word class character or an edge. |
|     `\w`      |      Matches an alphanumeric character, including "_";       |
|     `\W`      |    Matches a *non*-alphanumeric character, excluding "_";    |
|     `\s`      | Matches a whitespace character, which in ASCII are tab, line feed, form feed, carriage return, and space; |
|     `\S`      |             Matches anything *but* a whitespace.             |
|     `\d`      |                       Matches a digit.                       |
|     `\D`      |                          non-digit                           |
|     `\A`      | Matches the beginning of a string (but not an internal line). |
|     `\z`      |   Matches the end of a string (but not an internal line).    |

| Metacharacter |               Description               |
| :-----------: | :-------------------------------------: |
|     `\u`      |                 Unicode                 |
|     `\x`      |           Hexadecimal digits            |
|     `\p`      | Visible characters and space characters |

### Flags

| Flags |                         Description                          |
| :---: | :----------------------------------------------------------: |
|  `g`  |                            global                            |
|  `i`  |                            ignore                            |
|  `m`  | multiline, let `^ $` matches the beginning and end of each ***line***. |
|  `s`  |               singleline, let `.` include `\n`               |

## Further Syntax

### In brackets

#### Backreference (custom name of a group)

`(?<Word>\w+)` or `(?'word'\w+)`, to use backreference, use `\k<Word>`.

We also can make a group not to be tagged: `(?:exp)`.

#### Assertions

|    Code    |            Description             |
| :--------: | :--------------------------------: |
| `(?=exp)`  |     the position before `exp`      |
| `(?<=exp)` |      the position after `exp`      |
| `(?!exp)`  | the position followed by non-`exp` |
| `(?<!exp)` |    the position after non-`exp`    |

#### Comments

Use `(?#comment)` to add a comment in a regexp.

#### Stack

In fact, when we use `(?'word'\w+)`, we make what `\w+` matches into a stack called “word”, there are more syntax for this feature:

|        Code        |                         Description                          |
| :----------------: | :----------------------------------------------------------: |
|   `(?'-group')`    |                             pop                              |
| `(?(group)yes|no)` | if there are some items in “group” stack, then match `yes`, otherwise `no` |



