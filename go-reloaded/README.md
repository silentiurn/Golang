# go-reloaded


## Task Description

This utility takes a text file as input, applies a set of transformations to its contents, and writes the result to an output file.

The program supports:

* number conversions from **hex** and **bin** to decimal
* word case modifiers **low**, **up**, **cap**
* punctuation normalization
* proper quote formatting
* automatic article correction (**a → an**)

Only **standard Go packages** are allowed.

---

## Project Structure

```
go-reloaded/
├── articles.go       # Article correction logic (a / an)
├── conversions.go    # Number conversions and word case modifiers
├── punctuation.go    # Punctuation normalization
├── quotes.go         # Single quote handling
├── processor.go     # Coordinates all transformations in order
├── main.go           # Program entry point
├── sample.txt         # Sample input file
├── result.txt        # Sample output file
├── go.mod            # Go module definition
└── README.md         # Documentation
```

---

## Usage

The program takes **two arguments**:

1. input file name
2. output file name

### Run

```bash
go run . <input_file> <output_file>
```

### Example

```bash
go run . sample.txt result.txt
```

---

## Supported Transformations

### 1. Number Conversions

* `(hex)` — replaces the previous word (a hexadecimal number) with its decimal value
* `(bin)` — replaces the previous word (a binary number) with its decimal value

**Example:**

```
1E (hex) files were added
→
30 files were added
```

---

### 2. Case Modifiers

* `(up)` — converts the previous word to **UPPERCASE**
* `(low)` — converts the previous word to **lowercase**
* `(cap)` — converts the previous word to **Capitalized** form

Modifiers also support a number of words:

* `(up, N)`
* `(low, N)`
* `(cap, N)`

**Example:**

```
This is so exciting (up, 2)
→
This is SO EXCITING
```

---

### 3. Punctuation

Supported punctuation marks:

```
. , ! ? : ;
```

Rules:

* punctuation marks must be placed **right after the previous word**
* exactly **one space** must follow the punctuation

**Example:**

```
Hello ,world !
→
Hello, world!
```

Punctuation groups such as `...` or `!?` are preserved:

```
I was thinking ... You were right
→
I was thinking... You were right
```

---

### 4. Quotes `'`

* single quotes always appear in **pairs**
* extra spaces inside quotes are removed

**Examples:**

```
' awesome '
→
'awesome'
```

```
' I am the most well-known homosexual in the world '
→
'I am the most well-known homosexual in the world'
```

---

### 5. Articles (a / an)

The article `a` is automatically replaced with `an` if the next word starts with:

```
a, e, i, o, u, h
```

**Example:**

```
A amazing rock
→
An amazing rock
```

---

## Examples

### Example 1 — Case & Punctuation

**Input:**

```
it (cap) was the best of times, it was the worst of times (up) ,
```

**Output:**

```
It was the best of times, it was the worst of TIMES,
```

---

### Example 2 — Numbers

**Input:**

```
Simply add 42 (hex) and 10 (bin)
```

**Output:**

```
Simply add 66 and 2
```

---

### Example 3 — Articles

**Input:**

```
bearing a untold story
```

**Output:**

```
bearing an untold story
```

---

### Example 4 — Punctuation

**Input:**

```
Punctuation tests are ... kinda boring ,what do you think ?
```

**Output:**

```
Punctuation tests are... kinda boring, what do you think?
```

---

## Dependencies

This project **does not use any external libraries**.

Only standard Go packages are used, including.

---

## Notes

This project was implemented as part of a learning assignment and follows strictly defined text processing rules.
