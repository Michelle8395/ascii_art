## 1. Initial State: Hardcoded Banner and CLI-Only Program

i started with a **command-line ASCII art program** written in Go.

Key characteristics of the initial version:

* The program accepted **text via `os.Args`**
* The banner file was **hardcoded**:

  ```go
  os.Open("shadow.txt")
  ```
* No flexibility. No banner choice. No web interface.
* Logic and rendering worked, but the design was rigid.

### Problems at this stage

* You could not switch between `standard.txt`, `shadow.txt`, or `thinkertoy.txt`
* Any new banner required editing source code
* The program logic was tightly coupled to the banner file

This was fine for passing early tasks, but terrible for extensibility.

---

## 2. Removing the Default Banner Dependency

The first real fix was **eliminating the hardcoded banner**.

### What changed conceptually

Instead of:

* “The program decides which banner to use”

We moved to:

* “The user decides which banner to use”

This required:

* Accepting the banner filename as an input
* Passing that filename into the rendering logic

---

## 3. Extracting Rendering Logic into `RenderAscii`

This was the most important architectural change.

### Before

* `main.go` did everything:

  * Read input
  * Open banner file
  * Build the banner map
  * Render ASCII
  * Print output



### After

I created a reusable function in `utils.go`:

```go
func RenderAscii(input string, bannerFile string) (string, error)
```

### What `RenderAscii` does

1. Opens **any banner file** passed to it
2. Builds the rune → ASCII mapping
3. Splits input by newlines
4. Renders ASCII output
5. Returns the result as a string

### Why this mattered

* Rendering logic became **reusable**
* CLI and Web could now share the same core logic
* Banner choice became dynamic

---

## 4. Separating CLI and Web Responsibilities

 I then ran into this error:

```
main redeclared in this block
```

### The fix

You separated concerns properly:

* `main.go`

  * CLI version
  * Uses `RenderAscii`
* `server.go`

  * Web server
  * Uses `RenderAscii`
* `utils.go`

  * Shared logic
  * No `main()` function

### Result

* No duplicate `main`
* Clean compilation
* Ability to run:

  * CLI: `go run main.go utils.go`
  * Web: `go run server.go utils.go`



## 5. First Web Version (Text Input + Banner Choice)

I then moved to the browser.

### What the web server initially did

* Served a form at `localhost:8080`
* Allowed:

  * Text input
  * Banner selection
* Rendered ASCII output using `RenderAscii`

### Initial issues

* Text disappeared when switching banners
* Banner selection did not persist
* Page state reset on every request



## 6. Fixing the “Text Gets Lost” Problem

It was a **state management bug**.

### Root cause

* The server was not sending previous form values back to the template

### Fix

I introduced a proper data structure:

```go
type PageData struct {
	Input   string
	Banner  string
	Result  string
	Error   string
	Banners []string
}
```

And ensured:

* The textarea re-renders with `{{.Input}}`
* The selected banner persists using `selected`
* Output only changes when the user submits

### Result

* Users can compare:

  * Standard
  * Shadow
  * Thinkertoy
* Without retyping text every time



## 7. Converting Banner Selection into a Dropdown

Instead of free-text banner names, I moved to **controlled input**.

### What changed

* Banner choices are predefined:

  ```go
  []string{"standard.txt", "shadow.txt", "thinkertoy.txt"}
  ```
* Rendered as a `<select>` dropdown
* Prevents invalid filenames
* Eliminates user error

## 8. Adding CSS Styling

Up to this point, the app worked but looked like a crime scene.

### What the CSS added

* Dark theme
* Monospace font for ASCII output
* Proper spacing and layout
* Styled textarea, dropdown, and button
* Scrollable `<pre>` block for large ASCII art

### Why inline CSS was acceptable here

* Small project
* No static file server yet
* Focus was on functionality first


## 9. Final Architecture Overview

At the end, my project looks like this:

### Files

* `utils.go`

  * Banner parsing
  * ASCII rendering
  * `RenderAscii`
* `main.go`

  * CLI entry point
* `server.go`

  * Web server
  * HTML template
  * CSS
* `standard.txt`
* `shadow.txt`
* `thinkertoy.txt`


       