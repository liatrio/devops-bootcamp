# Code styling / formatting
“Formatting issues are the most contentious but the least consequential. People can adapt to different formatting styles but it's better if they don't have to, and less time is devoted to the topic if everyone adheres to the same style.” https://golang.org/doc/effective_go.html#formatting

When writing or dealing with a lot of code, it can become quite cumbersome when reading and writing code of different styles. Some people might use snake_case, some people might use camelCase. Some people might capitalize their class names, and some people might not. To solve this, a particular style is adopted for the code base. Some languages such as Golang are built around a specific code style provided by the gofmt tool. Other languages such as C++ have a myriad of different styles, all with the same general goals https://google.github.io/styleguide/cppguide.html#Goals

Oftentimes, the style guides are long and in-depth. To solve this, specific tools known as formatters and linters have been developed to help developers adhere to specific style guides. 

To learn more about the difference between formatters and linters, read https://prettier.io/docs/en/comparison.html

## Golang gofmt
	Gofmt is a standard code formatting tool for the Go language. Using tools like these can create a consistency and predictability for your source code. This means a codebase that has gone through the gofmt will be easier to read and easier to maintain. Running gofmt -d . will print the differences between the code in the current directory and what the gofmt standard would look like. Running gofmt -w . will write the changes proposed by gofmt into your code.
