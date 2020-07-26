[![Run on Repl.it](https://repl.it/badge/github/matt4biz/go-class-counts)](https://repl.it/github/matt4biz/go-class-counts)

# Go class: line counts example
The first example counts all words in a file (equivalent to `wc -w`).

```shell
$ go run ./counts rich2.txt
2558 total words

$ wc -w rich2.txt
    2558 rich2.txt
```

The second counts only _unique_ words.

```shell
$ go run ./uniq rich2.txt
1197 unique words
```

This can be checked by using AWK to find the words and then running them through a pipeline of other Unix commands:

```shell
$ awk '{for (i=1; i<=NF; i++) {print $i}}' rich2.txt|sort|uniq|wc -l
    1197
```

The file `rich2.txt` is taken from Shakespeare, Richard II act 2 scene 1.
