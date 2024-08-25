CLI tool to recreate and entire Git repository from a deleted commit hash, branch or tag.

To use the `git-restore` binary, do:

```
$ go install github.com/Patel-Raj/git-restore@latest
$ git-restore [/path/to/source/repo] [/path/to/destination/directory] [commit hash/branch/tag]

// Eaxmple
$ git-restore . ./copy 0c7dd4e4b626a23632e7b4a54ccfcb91e5c9960f
```

To build this binary from source locally, do:

```
$ git clone git@github.com:Patel-Raj/git-restore.git
$ cd git-restore
$ go build
$ ./git-restore . ./copy 0c7dd4e4b626a23632e7b4a54ccfcb91e5c9960f
```

How is this tool useful?

Consider the following scenario in your development workflow:

1. You are working on the `feature` branch.
2. You make a couple of changes and commit them to the `feature` branch with the commit message `commit-message-1`.
3. You continue your development and create another commit on the `feature` branch with the commit message `commit-message-2`.
4. However, it turns out that you don't need the changes from `commit-message-2` and want to remove the entire commit, moving the `HEAD` back to `commit-message-1`.
5. You mistakenly execute `git reset --hard HEAD~2` instead of `git reset --hard HEAD~1`. As a result, `HEAD` moves back by 2 commits, and the changes from `commit-message-1` are deleted as well.
6. This is where the `git-restore` CLI tool can be handy. You can fetch the commit hash of `commit-message-1` using `git reflog` by identifying the commit message, and then use the tool to restore the entire state of the repository to that commit hash.

```
$ git reflog
9d82c22 (HEAD -> main) HEAD@{0}: reset: moving to HEAD~2
3ca40e5 HEAD@{1}: commit: commit-message-2
7cb88eb HEAD@{2}: commit: commit-message-1

$ git-restore . . 7cb88eb
```

If you find bugs, improvements or have questions about applyinig this approach, feel free to open a [GitHub issue](https://github.com/Patel-Raj/git-restore/issues).

