# Open Shell

A cross platform shell program with an open plugin system.

# Road Map

## Main

- [ ] server-client architecture, behaves like tmux by default
- [ ] socket based cross-language plugin system, osh uses [grpc](https://grpc.io/) to talk with plugins
- [ ] plugin downloader and manager

## Plugins

- [ ] vscode like goto anywhere for history and custom commands
- [ ] vscode like autocomplete dropdown while typing each char
- [ ] syntax highlight while typing


## Osh Script

Strong typed shell script. The goal is be easy and safe to use,
performance is not the main target.

It has two layers, the core layer is lisp, the sugar layer is for
`bash` like syntax which will convert to core layer before the execution.

### Types

```
function
number
string
array
map
boolean
```

No null type, nothing can be null


### Function call

just like bash, use `#` to comment.

```
foo 1 2

# foo(1 2)
```

```
1 + 2

# vm.add(1 2)
```

```
1 * (2 + 3)

# vm.multiply(1 vm.add(2 3))
```

```
if true {
    echo 'yes'
} else {
    echo 'no'
    echo 'no'
}

# vm.if(true, {
#     echo('yes')
# }, {
#     echo('no')
#     echo('no')
# })
```

String use symmetric single quote to quote block

```

a = 'test'

# vm.def(a 'test')

a = '
    test
'

a = ''
    'tes'
''

a = '''''
    '''tes'''
'''''

```


```
foo = (a b) -> {
    ret a + b
}

# vm.def_fun(foo [a b] {
#     vm.ret(vm.add(a b))
# })

foo foo(1 2) foo(1 2)
```