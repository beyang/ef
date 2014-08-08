ef
====

Make `cd` a little more *ef*ficient.

Works just like `cd`, but you can specify a partial path (not necessarily relative to your working directory), and `ef`
will search your home directory tree for directories matching that path. It intelligently sorts matches and `cd`s you
into the best match.

Usage
----
```bash
~$ ls
bin src pkg tmp
~$ ef src              # works just like cd!
~/src$ ef
~$ ef flask            # jump to directory by name
~/src/github.com/mitsuhiko/flask$ ef sourcegraph
~/src/sourcegraph.com/sourcegraph$ ef beyang/ef       # specify not just a name, but a partial path
~/src/github.com/beyang/ef$ ef
~$ 
```

Install
-----

1. `go get github.com/beyang/ef/...`
1. Add this to your `.bashrc`:
```
source path/to/cd.sh
```
1. If you really like it, you can also add to `.bashrc`:
```
alias cd=ef
```
