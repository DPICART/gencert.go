# gencert.go
## Crashcourse - Creating a 'certificate' generator in go
Small project based on the online course "Le language Go | Formation complète"

https://www.udemy.com/course/le-langage-go-formation-complete

## How to use ?
You'll need go installed on your machine.

To generate multiple certificates from a csv file: 

PDF format: ```go run main.go -file students.csv -type pdf```

HTML format: ```go run main.go -file students.csv -type html```