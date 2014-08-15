ef
====

Make `cd` a little more <strong>ef</strong>-ficient.

Works just like `cd`, but you can specify a directory name or path (not necessarily relative to your working directory),
and `ef` will search your home directory tree for directories that match. `ef` intelligently sorts matches and `cd`'s you
into the best match.

Usage
----
```bash
~$ ls
bin src pkg tmp
~$ ef src                # works just like cd!
~/src$ ef
~$ ef flask              # jump to directory by name
~/src/github.com/mitsuhiko/flask$ 
~/src/github.com/mitsuhiko/flask$ ef mux
~/src/github.com/gorilla/mux$ 
~/src/github.com/gorilla/mux$ ef beyang/ef  # type not just a name, but a partial path
~/src/github.com/beyang/ef$ ef
~$ 
```

Install
-----

1. `go get github.com/beyang/ef/...`
1. Add this to your `.bashrc`:
```
source path/to/this/repository/ef.sh
```
