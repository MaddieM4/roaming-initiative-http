roaming-initiative-http
=======================

The front-facing site interface for roaming-initiative.com, not including the blog.

While this is intended to be useful as a "head start" for making similar sites,
it must primarily service its actual use on the author's Linode-hosted site at
http://roaming-initiative.com/, so while I'll happily merge in changes that
make this run on a wider range of system configurations, I'll only be developing
and testing with Debian Stable, and writing for my own needs.

This used to be a Python project. I decided I hated uwsgi, so I ported
it to Golang. If you use this, be mindful that if your server platform
is a different than that of your dev env, you will need to cross
compile.

DEPENDENCIES
============

 * nginx >= 1.2.1 (via squeeze-backports)
 * Golang 1.1+ (I use 1.2)
 * http://github.com/codegangsta/martini
 * http://github.com/codegangsta/martini-contrib/render
