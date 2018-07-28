# What is bk & jm

## bk
bk bookmarks your current directory.

```bash
% pwd
/tmp
% bk save
% bk show
/tmp
% cd /var
% bk save
% bk show
/tmp
/var
```

```bash
% bk del
QUERY>                                                                                                                                                                                  IgnoreCase [2 (1/1)]
/tmp     /* choose /tmp */
/var
% bk show
/var
```

## jm
jump from current directory to a directory which is one of saved by bk command. 

```bash
% . jm   /* MUST NEED dot command OR make alias */
QUERY>                                                                                                                                                                                  IgnoreCase [2 (1/1)]
/var     /* choose /var */
/tmp
% pwd
/var     /* change current directory */
```

alias example in ~/.bash_profile
```bash
alias jm='. jm'
```

# Install

### required

- golang

- peco.

```
brew isntall peco
```

build go src and make symbolic link.

```
make build
make install
```

Completed!! :tada:
