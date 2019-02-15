# Universal Note Converter

unc is a command line tool to convert between various note formats. It currently supports the following, fairly arbitrary list of formats:

* [Evernote](https://evernote.com) (Export only)
* Markdown (Export only)
* [Paper](https://paper.dropbox.com) (Import only)
* Vesper (Import only)

[![Build Status](https://travis-ci.org/cbguder/unc.svg?branch=master)](https://travis-ci.org/cbguder/unc)

## Installation

```
go get github.com/cbguder/unc
```

## Usage

```
unc -f <source format> -t <destination format> -i <input path> -o <output path>
```
