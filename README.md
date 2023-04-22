# go-training


## installing present tool
    go install golang.org/x/tools/cmd/present

## before using the present tool
- make sure go mod is initialized on the parent folder
- and run the below
    go get golang.org/x/tools/cmd/present

## running presentation slides

### local
% present
2022/11/15 08:20:33 Open your web browser and visit http://127.0.0.1:3999

### go talks
https://go-talks.appspot.com/github.com/muly/go-training/slice.slide [server error]

working example: https://go-talks.appspot.com/github.com/davecheney/presentations/reproducible-builds.slide#1

not working example: 403 error from github api
https://go-talks.appspot.com/github.com/muly/go-training/slice.slide#1
    Error accessing api.github.com.
        403: (https://api.github.com/repos/muly/go-training)
    Error accessing api.github.com.
        get https://api.github.com/repos/muly/go-training/contents/_concepts/03_slice/07forLoop/6range_withValueIgnoredInSimpleForm.go -> 403


sometimes not working:
https://go-talks.appspot.com/github.com/davecheney/presentations/5nines.slide#26
    similar 403 error

open issue related to this:
    https://github.com/golang/go/issues/58906


