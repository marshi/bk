# What is bk & jm

## bk
bk bookmarks your current directory.

```
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

```
% bk del
QUERY>                                                                                                                                                                                  IgnoreCase [2 (1/1)]
/tmp     /* choose /tmp */
/var
% bk show
/var
```

## jm
jm changes current directory to a directory which is one of saved by bk command. 

```
% . jm   /* MUST NEED dot command OR make alias */
QUERY>                                                                                                                                                                                  IgnoreCase [2 (1/1)]
/var     /* choose /var */
/tmp
% pwd
/var     /* change current directory */
```

alias example in ~/.bash_profile
```
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
make 
```

Completed!! :tada:
