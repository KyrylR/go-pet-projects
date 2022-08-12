# My grep clone

A custom `grep' that is capable of searching for entries in files.

## Description

A grep clone that can do a simple substring search
within files. It can also automatically search subdirectories.

## Purpose

To consolidate my knowledge in working with Go Modules & Packages,
Goroutines, Channels, Wait Groups and bufio package

## Product functions (brief description)

* Used goroutines to search for substrings in files.
* Displaying matches on the terminal as they are found
* Displayed string number, file path and entire string containing the match.
* Recursive search for matches in any subdirectories.
* A synchronisation method was used to ensure that all files
  were found and all results were displayed before the program
  is terminated.


### Executing program

The program is invoked according to the scheme:
```
go run ./mgrep [--searchterm SEARCHTERM] [--searchdir SEARCHDIR]
```

[//]: # (## Content of the system &#40;system boundaries&#41;)

[//]: # ()
[//]: # (## Interaction &#40;potential&#41; of the product &#40;with other products and components&#41;)


[//]: # ()
[//]: # (## Security requirements)

[//]: # ()
[//]: # (## User characteristics &#40;who is the end user of the system&#41;)

[//]: # ()
[//]: # (## Limitations)


