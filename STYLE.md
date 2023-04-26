# Liatrio's DevOps Bootcamp

## Style Guide

### HTML

Docsify and some Markdown viewers supports mixing HTML and Markdown elements in the same document. There must be a blank line between the two formats.

```html
<center>

  ![](../../img/goals.png)

</center>
```

Note the blank line between the `<center>` HTML tag and the `![](../img/goals.png)` Markdown tag.

### Multi Column Layout

Multi column layouts can be created using HTML div tags with following CSS classes:

- `grid2`, `grid3`, `grid4`: Define a section with 2, 3 or 4 columns
- `col`: Define a column inside of a grid section

Here is an example of a 3 column layout

```html
<div class="grid3"><div class="col">
<center>

**Column 1**

</center>

1. one
2. two
3. three

</div><div class="col">
<center>

**Column 2**

</center>

* one
* two
* three

</div><div class="col">
<center>

**Column 3**

</center>

>blaa blaa blaa

</div></div>
```
