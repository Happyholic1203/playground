# Playground
A playground for practicing SQL, SQLAlchemy and stuff.

# Usage
`./go <mysql|sa|sqlalchemy|tmux|bash>`

* Modify `scripts/mysql-default-data` to insert your default data
* `mysql`: enters mysql shell using the custom db `mydb`
* `sa|sqlalchemy`: enters ipython with custom tables and instances initialized
* `tmux`: start with tmux
* `bash`: start with bash

# Git Playground

## Useful Commands
* `git merge-base`
* `git describe`

```
              + ---------------------+
              |                      |
              |      + --- m2_1 --- m2_2 <<< v2.0
              |      |
  fix-M3 >>> fM3     |        c3 --- c4 <<< feature-c
              |      |       /
M1 --- M2 --- M3 --- M4 --- M5 <<< master
       | \
       |  a1 --- a2 --- a3 <<< feature-a
       |    \
       |     b1 --- b2 --- b3 <<< feature-b
       |
       + --- m1_1 --- m1_2 --- m1_3 <<< v1.0
                       |
                      fM1 <<< fix-M1


              + ---------------------+
              |                      |
              w                      v
              |      + --> m2_1 --> m2_2 <<< v2.0
              |      |
 fix-M3 >>>  fM3     |      c3 --> c4 <<< feature-c
              ^      |       ^
              |      |       |
M1 --> M2 --> M3 --> M4 --> M5 --> M6 <<< master
       |\                          ^
       | \                         |
       |  \                        +-------------------------+
       |   \                                                 |
       |    a1 --- a2 --- a3 <<< feature-a                   |
       |      \                                              |
       |       b1 --- b2 --- b3 <<< feature-b                |
       |                                                     |
       + --- m1_1 --- m1_2 --- m1_3 --- m1_4 <<< v1.0 |
                       |                  ^                  |
                       v                  |                  |
          fix-M1 >>>  fM1 ----------------+------------------+
```

## Rebase
Rebase is *recreating* your work on a certain branch onto another.
`rebase <dst> <src>`
* Idea: save the diff on `src`, and generate new commits on top of `dst` with
    those changes applied.
    * Ex. `rebase master feature-a`
        * Generates new commits: `a1'`, `a2'`, `a3'`
        * Those commits are *rebased* on `master`, yielding:
            * `...M4 --- M5 --- a1' --- a2' --- a3'`
            * `master` still at `M5`
            * `feature-a` is now at `a3'`
* Rebase v.s. merged
    * Rebase local changes to cleanup your dev. *story*,
      but do NOT rebase the changes you've made public

## Practice
* Apply `fix-M3` to all descendent branches
    * Already merged into `m2_2` because it's urgent.
    * But we also want it on: `master`, `feature-c`
        * Future branches should include this fix
* Apply `fix-M1` to `v1.0` and `v2.0`
    * `fix-M1` is a bug since `M1`, but you fixed it on `m1_2`
    * Apply the patch on `m1_2` and rebase
* We want `feature-b` in the next 3.0 release
    * But it's not stable and untested after `a2`
    * `b1` depends on the stable `a1`
    * Need to apply the changes `a1...b3`

## Possible Solutions
* Apply `fix-M3` to all descendent branches (`master`, `feature-c`)
    * 
* Apply `fix-M1` to `v1.0` and `v2.0`
    * `git co -b fix-M1-1.0 fix-M1`
    * `git co -b fix-M1-2.0 fix-M1`
    * `git rebase v1.0 fix-M1-1.0`
    * `git co v1.0 && git merge fix-M1-1.0 # fast forward`
    * `git rebase v2.0 fix-M1-2.0`
    * `git co v2.0 && git merge fix-M1-2.0 # fast forward`
* We want `feature-b` in the next 3.0 release
    * `git rebase --onto master feature-a feature-b`
    * `git checkout master && git merge feature-b`
