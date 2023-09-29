const rawQuizdown = `
# This is the quiz for Chapter 1.1?

- [x] True
- [ ] False

# Who is the person in the picture?

![](https://upload.wikimedia.org/wikipedia/commons/thumb/9/9d/Sir_Tim_Berners-Lee.jpg/330px-Sir_Tim_Berners-Lee.jpg)

> In 1990, he published the [worlds first website](http://info.cern.ch/hypertext/WWW/TheProject.html).

1. [x] Tim Berners-Lee
    > This is the correct answer!
1. [ ] Alan Turing
1. [ ] Barbara Liskov
1. [ ] Larry Page
    > This is not Larry Page ;)

# Select your superpower!

There exist many superpowers in the world but one of them is better than everything else. Do you find it?

1. [ ] Enhanced Strength
1. [ ] Levitation
1. [x] Shapeshifting
    > Correct. This the best superpower!


# Please bring the following into sequence!

That's **super easy**!

> Three is larger than one...

1. One
2. Two
3. Three
4. Four
5. Five

# Python Lists

Testing syntax highlighting

> Python lists are mutable!

\`\`\`python
# Python
x = [2, 3, 4]
y = "string"
x[2] = 4
print(x[2])
\`\`\`

\`\`\`bash
# Bash
y = "string"
echo $y
rm -rf /
aws-vault exec nick-is-awesome -- aws s3 ls
\`\`\`

\`\`\`js
// Javascript
const x = [2, 3, 4]
x[2] = 4
let y = 'string'
console.log(x[2])
\`\`\`

\`\`\`json
{
  "name":"John",
  "age":30,
  "car":null
}
\`\`\`

\`\`\`go
// GO
package main
import "fmt"
func main() {
    fmt.Println("hello world")
}
\`\`\`

\`\`\`yaml
on:
  # Runs on pushes targeting the default branch
  push:
    branches: ["master"]

permissions:
  contents: read
  pages: write
  id-token: write
\`\`\`

- [ ] I'm done playing
`;

export { rawQuizdown }