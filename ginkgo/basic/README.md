# Motivation

this test case is used to make it clear that how to use ginkgo, and the relationship among all containers block, `Describe`, `Context`

# how to use ginkgo

http://onsi.github.io/ginkgo/

## start a test suite

test suite is the top level, which includes many test spec

```
$ cd path/to/books
$ ginkgo bootstrap

```

## create test spec

create a collection of ginkgo test spec

each It container correspondes to an individual spec.

```
$ ginkgo generate <package_name>
```






