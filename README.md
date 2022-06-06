# colors

Easily add ANSI colors to your text and symbols in the terminal.

## Install

```sh
$ go get -u github.com/davipatricio/colors
```

## Example

```go
package main

import (
    "fmt"

    "github.com/davipatricio/colors/backgrounds"
    "github.com/davipatricio/colors/bright"
    "github.com/davipatricio/colors/colors"
    "github.com/davipatricio/colors/styles"
)

func main() {
    fmt.Println(colors.Red("Red"))
    fmt.Println(colors.Green("Green"))

    fmt.Println(styles.Bold("Bold"))
    fmt.Println(styles.Underline("Underline"))

    fmt.Println(bright.Yellow("Bright Yellow"))
    fmt.Println(bright.Magenta("Bright Magenta"))

    fmt.Println(backgrounds.Cyan("Cyan Background"))
    fmt.Println(backgrounds.White("White Background"))

    // Mixed
    fmt.Println(styles.Bold(colors.Red("Bold Red")))
    fmt.Println(backgrounds.White(colors.Black("White background, black text")))
}

```

## Available styles

**Note** that bright and bright-background colors are not always supported.


| Colors | Background Colors | Bright Colors |
|-------|-------------------|--------------|
| Black | Black | Black |
| Red | Red | Red |
| Green | Green | Green |
| Yellow | Yellow | Yellow |
| Blue | Blue | Blue |
| Magenta | Magenta | Magenta |
| Cyan | Cyan | Cyan |
| White | White | White |
| Gray | |
| Grey | |

(`gray` is the U.S. spelling, `grey` is more commonly used in the Canada and U.K.)

## Style modifiers

Those methods are available in the `github.com/davipatricio/colors/styles` package.

- Dim
- **Bold**
- Hidden
- _Italic_
- __Underline__
- ~~Strikethrough~~
- Reset

## License

MIT
