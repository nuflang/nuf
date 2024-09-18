# Nuf language

## Commands

### Generate

Converts `nuf` file into `html`

`generate <input-file-pathname> <output-file-pathname>` where:
- `<input-file-pathname>` - `.nuf` file
- `<output-file-pathname>` - `.html` file

```sh
go run main.go generate ./test_data/text.nuf ./test_data/text.html
```

## Syntax

### Paragraph

Input:

```nuf
"Paragraph 1";
"Paragraph 2";
```

Output:

```html
<p>Paragraph 1</p>
<p>Paragraph 2</p>
```

### Heading

Input:

```nuf
section_title("Heading 1");
section_title("Heading 2");
```

Output:

```html
<h1>Heading 1</h1>
<h1>Heading 2</h1>
```

---

## Attribution

Builds on top of
- [Introduction to Writing Modern Parsers (YouTube playlist)](https://www.youtube.com/watch?v=V77J9l8N-P8&list=PL_2VhOvlMk4XDeq2eOOSDQMrbZj9zIU_b) by [tylerlaceby (YouTube channel)](https://www.youtube.com/@tylerlaceby)
- [Writing An Interpreter In Go book](https://interpreterbook.com/) by [Thorsten Ball](https://x.com/thorstenball)
