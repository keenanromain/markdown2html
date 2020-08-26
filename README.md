# markdown2html


## Summary
This is a small Go program that takes in markdown file as an argument and produces a matching HTML file. 
The corresponding HTML file is written into an `output` directory inside of this directory once the program completes.
This project was written under a time contraint so the features included are limited in scope. It takes the following:

```
# Header one

Hello there

How are you?
What's going on?

## Another Header

This is a paragraph [with an inline link](http://google.com). Neat, eh?

## This is a header [with a link](http://yahoo.com)
```
and produce HTML that looks like:

```
<h1>Header one</h1>

<p>Hello there</p>

<p>How are you?
What's going on?</p>

<h2>Another Header</h2>

<p>This is a paragraph <a href="http://google.com">with an inline link</a>. Neat, eh?</p>

<h2>This is a header <a href="http://yahoo.com">with a link</a></h2>
```

## Features
The full list of supported features is below:

| Markdown                               | HTML                                              |
| -------------------------------------- | ------------------------------------------------- |
| `# Heading 1`                          | `<h1>Heading 1</h1>`                              | 
| `## Heading 2`                         | `<h2>Heading 2</h2>`                              | 
| `...`                                  | `...`                                             | 
| `###### Heading 6`                     | `<h6>Heading 6</h6>`                              | 
| `Unformatted text`                     | `<p>Unformatted text</p>`                         | 
| `[Link text](https://www.example.com)` | `<a href="https://www.example.com">Link text</a>` | 
| `Blank line`                           | `Ignored`                                         | 
| `** Bold **`                           | `<strong>Bold</strong>`                           |
| `* Italitcs *`                         | `<em>Italics</em>`                                |

## Improvements
There are always points of improvements with any project. Given enough time, I'd would have liked to have done the following:

- Include CSS styling on the rendered HTML 
- Include support for pulling a markdown file off the internet
- Include multiple markdown files as input
- Divide the Go file across multiple files
- Incorporate more in-depth testing
- Include support for ordered and unordered HTML lists
