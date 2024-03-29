<div align="center">
<pre>
          __                           ____ 
   _____ / /__ ____   ____  _      __ / __ \
  / ___// //_// __ \ / __ \| | /| / // /_/ /
 (__  )/ ,<  / / / // /_/ /| |/ |/ // _, _/ 
/____//_/|_|/_/ /_/ \____/ |__/|__//_/ |_|  
--------------------------------------------
A CLI tool for making HTTP requests written in Go
</pre>

[![Go Reference](https://pkg.go.dev/badge/github.com/sknoww/sknowR@latest.svg)](https://pkg.go.dev/github.com/sknoww/sknowR) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

</div>

## Installation

```sh
go install github.com/sknoww/sknowR@latest
```

## Usage
### Help and options info
```sh
sknowR -h
sknowR --help
```

### Command line usage
```sh
sknowR -f example/folder/request.json
```

### Command line usage with outfile
```sh
sknowR -f example/folder/request.json -o example/folder/response.json
```

## IO Redirection
Output can be redirected to a file. The response body is written to stdout and status/headers are written to stderr. 
```sh
sknowR -f example/folder/request.json > response.json 2> response_headers.json
```
This can also be done to the same file.

#### Note: There are issues with this in powershell but it works in cmd/command prompt. Also, this must be a .txt file if writing to the same file.
```sh
sknowR -f example/folder/request.json > response.txt 2>&1
```

## File Downloads
You can also use sknowR to download files such as PDFs. However, you must specify an output file with the proper file extension.
```sh
sknowR -f example/folder/request.json -o response.pdf
```

## Sample requests (`request.json / request.yaml`)

### GET
```js
{
    "url": "https://jsonplaceholder.typicode.com/posts",
    "method": "GET"
}
 ```
```yml
url: https://jsonplaceholder.typicode.com/posts
method: GET
```

### GET (File Download)
#### Note: Must specify an output file
```js
{
    "url": "https://www.golang-book.com/public/pdf/gobook.pdf",
    "method": "GET"
}
```
```yml
url: https://www.golang-book.com/public/pdf/gobook.pdf
method: GET
```

### POST
```js
{
    "url": "https://jsonplaceholder.typicode.com/posts",
    "method": "POST",
    "headers": {
        "Content-Type": "application/json; charset=UTF-8"
    },
    "body": {
        "title": "foo",
        "body": "bar",
        "userId": 1
  }
}
 ```
```yml
url: https://jsonplaceholder.typicode.com/posts
method: POST
headers:
  Content-Type: application/json; charset=UTF-8
body:
  title: foo
  body: bar
  userId: 1
```

### PUT
```js
{
    "url": "https://jsonplaceholder.typicode.com/posts/1",
    "method": "PUT",
    "headers": {
        "Content-Type": "application/json; charset=UTF-8"
    },
    "body": {
        "id": 1,
        "title": "foo",
        "body": "bar",
        "userId": 1
    }
}
 ```
```yml
url: https://jsonplaceholder.typicode.com/posts/1
method: PUT
headers:
  Content-Type: application/json; charset=UTF-8
body:
  id: 1
  title: foo
  body: bar
  userId: 1
```

### DELETE
```js
{
    "url": "https://jsonplaceholder.typicode.com/posts/1",
    "method": "DELETE"
}
 ```
```yml
url: https://jsonplaceholder.typicode.com/posts/1
method: DELETE
```

### Request with all compatible fields
```js
{
    "url": "XXX", // (REQUIRED)
    "method": "XXX", // (REQUIRED)
    "headers": {
        "Content-Type": "application/json",
        "Authorization": "XXX"
    },
    "params": {
        "XXX": "XXX"
    },
    "body": {
        "XXX": "XXX"
    },
    "cookies": {
        "XXX": "XXX"
    }
}
```
```yml
url: XXX # (REQUIRED)
method: XXX # (REQUIRED)
headers:
    Content-Type: application/json
    Authorization: XXX
params:
    XXX: XXX
body:
    XXX: XXX
cookies:
    XXX: XXX
```

### Metadata
H. Sullivan - sknow.codes@gmail.com

### License
MIT License, reference `LICENSE` for details.

### Roadmap
- Html(.html) response support
- Xml(.xml) response support
