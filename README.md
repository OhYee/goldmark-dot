# goldmark-dot

[![Sync to Gitee](https://github.com/OhYee/goldmark-dot/workflows/Sync%20to%20Gitee/badge.svg)](https://gitee.com/OhYee/goldmark-dot) [![w
orkflow state](https://github.com/OhYee/goldmark-dot/workflows/test/badge.svg)](https://github.com/OhYee/goldmark-dot/actions) [![codecov](https://codecov.io/gh/OhYee/goldmark-dot/branch/master/graph/badge.svg)](https://codecov.io/gh/OhYee/goldmark-dot) [![version](https://img.shields.io/github/v/tag/OhYee/goldmark-dot)](https://github.com/OhYee/goldmark-dot/tags)

goldmark-dot is an extension for [goldmark](https://github.com/yuin/goldmark).  

You can dot language to build svg image in your markdown like [mume](https://github.com/shd101wyy/mume)

## screenshot

There are two demo(using `'` instead of &#8242; in the code block)

1. default config

[Demo1](demo/demo1/main.go)
[Output1](demo/demo1/output.html)

```markdown
'''go
package main

import ()

func main(){}
'''

'''dot
digraph{a->b}
'''
```

![](img/default.png)

2. using `dot-svg` and [goldmark-highlighting extension](https://github.com/yuin/goldmark-highlighting)

[Demo2](demo/demo1/main.go)
[Output2](demo/demo1/output.html)

```markdown
'''go
package main

import ()

func main(){}
'''

'''dot-svg
digraph{a->b}
'''
```

![](img/highlighting.png)

## Installation

```bash
go get -u github.com/OhYee/goldmark-dot
```

## License

[MIT](LICENSE)
