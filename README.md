# MKLister

A command line tool for listing the contents of a directory.

## Installation

This tool is written in Go. If you don't have Go installed on your machine yet, follow these instructions to get up and running: [Go Installation Guide](https://golang.org/doc/install)

## Usage

MKLister supports the following CLI flags:

- --help (prints help info)
- --path=, -p= (the path to the directory, this is _required_)
- --recursive, -r (when enabled, print contents recursively. defaults to false)
- --output=, -o= (currently supports text, json, and yaml. defaults to text)

###Example Usage

```shell
$ mklister --path=/foo -r --output=json
```

###Example Response

```json
[
 {
  "Name": "bar",
  "Size": 136,
  "ModifiedTime": "2016-07-19T19:14:38-04:00",
  "IsLink": false,
  "IsDir": true,
  "LinksTo": "",
  "Children": [
   {
    "Name": "afolder",
    "Size": 136,
    "ModifiedTime": "2016-07-19T19:14:38-04:00",
    "IsLink": false,
    "IsDir": true,
    "LinksTo": "",
    "Children": [
     {
      "Name": "file1",
      "Size": 0,
      "ModifiedTime": "2016-07-18T18:30:08-04:00",
      "IsLink": false,
      "IsDir": false,
      "LinksTo": "",
      "Children": null
     },
     {
      "Name": "file2",
      "Size": 0,
      "ModifiedTime": "2016-07-18T18:30:10-04:00",
      "IsLink": false,
      "IsDir": false,
      "LinksTo": "",
      "Children": null
     }
    ]
   },
   {
    "Name": "todo",
    "Size": 19,
    "ModifiedTime": "2016-07-18T18:31:57-04:00",
    "IsLink": true,
    "IsDir": false,
    "LinksTo": "/Users/MaxKohl/todo",
    "Children": null
   }
  ]
 },
 {
  "Name": "hello",
  "Size": 0,
  "ModifiedTime": "2016-07-18T18:29:57-04:00",
  "IsLink": false,
  "IsDir": false,
  "LinksTo": "",
  "Children": null
 }
]
```
