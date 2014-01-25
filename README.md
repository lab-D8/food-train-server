Food Train
=====

Food Train REST API and client.

Just so we are all on par, this is written in GO. As such, we are going to follow the GO guidelines when it comes to where to put GO-code.

When you install GO, you have to give a GO_PATH, such as `~/GO`. Once you have this directory, in the GO_PATH, then you can run the following:

```
go get github.com/lab-D8/food-train-server
```

If you get the error ```bzr not installed``` or something of that sort, install bzr (brew install for macs, and your distro flavor for linux). We do not use bzr, but one of the packages we rely on does. 

Afterwards, your directory structure should resemble something like this...
```
---/GO
------/bin
---------/stuff
------/src
---------/github.com/lab-D8
------------/food-train-server
---------------/router.go
---------------/router_test.go
---------------/user
-------------------/user.go
-------------------/user_test.go
...
```


Build & Run Instructions
=========
Added to the wiki [here](https://github.com/iph/catan/wiki/Build-&-Run-Instructions).

Design Doc
=========

By working on the food train API, you agree to develop the API with the following mindset:

Law of Security
-----------------
If at any point, there is a risk that personal user data can be compromised, we will not go that route. For any action that changes board state, friend state or user state for a particular user, we will make sure that the person is Authenticated and Authorized to do that specific action.

Law of Simplicity (Keep It Simple Stupid)
-------------
If at any point, there is an easier solution to use, go with that. No need to make things convoluted. This rule is only superceded by security. Security is not simple.

Law of Speed
------
If two solutions are easy to make, choose the one that is faster. This is lower than KISS because easy solutions are generally easier to undestand, easier to debug and just as fast as a "faster" yet convoluted decision. This philosophy is why we chose Go as our language of choice: Simplicity first (over C), and Speed second.
