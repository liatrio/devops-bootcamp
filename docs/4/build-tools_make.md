# Make

`make` is a lightweight command line utility which allows users to run compiling,
testing, or installing steps in logical orders. Make  rules such as 'build'
or 'install' are defined by the user in a `Makefile`

## Makefile Structure

Let's walk through an example of a simple Makefile to explore how they look.

```
# The build target executable
TARGET = my-go-project

all: build hello

build:
	go build -o ${TARGET}

test: build
	go test -v

hello:
	echo Hello World

clean:
	rm ${TARGET}
```

Makefiles have a fairly basic structure. The file often begins with macros, which
define variables or commands that can be used later in the file. Then the Makefile
begins a list of rules, which look like the following:
```
target: dependencies
	command
```

When Make is run, the user either gives the target to run, such as `make test`
or simply runs `make` in which case the first rule will be run.

Whenever a rule is run any listed dependencies will also be run beforehand. This
allows Make to build even complex chains of dependencies with fairly simple
Makefiles. Files can also be used as dependencies, which can sometimes allow
make to skip steps if dependent files haven't changed.

It should also be noted there while there are common targets in many Makefiles
such as 'build', 'test', and 'clean', there's nothing special about the names
of these targets. Make is simply running the dependent steps and the specified
command for each of the rules given.

## Practice

Take a bit to look through a more complex Makefile, such as the Makefile for
[Rode](https://github.com/liatrio/rode/blob/master/Makefile). See how much you
can understand.

If you want, you can read more information about Make to help figure parts out:
- https://en.wikipedia.org/wiki/Make_(software)
- https://www.gnu.org/software/make/manual/html_node/Introduction.html
- https://www.tutorialspoint.com/makefile/index.htm

It's okay if you can't understand every part of the Makefile. This section's focus
is on getting you familiar with the larger concepts of Make and understanding
how the tool can be used.

# Deliverable

Discuss the purpose of Make and Makefiles.

Talk about what you could figure out from the more complex Makefile.

Discuss when you might use Make, and when you might avoid using it.
