# Universal Note Converter

unc is a command line tool to convert between various note formats. It currently supports the following, fairly arbitrary list of formats:

* [Evernote](https://evernote.com) (Export only)
* Markdown (Export only)
* [Paper](https://paper.dropbox.com) (Import only)
* Vesper (Import only)

[![Go](https://github.com/cbguder/unc/actions/workflows/go.yaml/badge.svg)](https://github.com/cbguder/unc/actions/workflows/go.yaml)
[![Maintainability](https://api.codeclimate.com/v1/badges/179a50ac83c139b246e5/maintainability)](https://codeclimate.com/github/cbguder/unc/maintainability)

## Installation

```
go get github.com/cbguder/unc
```

## Usage

```
unc -f <source format> -t <destination format> -i <input path> -o <output path>
```
