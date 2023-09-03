module poehelper

go 1.20

require (
	github.com/AllenDang/cimgui-go v0.0.0-20230825100656-022a31c244c1
	github.com/hpcloud/tail v1.0.0
	github.com/sirupsen/logrus v1.9.3
	golang.org/x/exp v0.0.0-20230817173708-d852ddb80c63
)

require (
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
)

replace github.com/AllenDang/cimgui-go => ..\cimgui-go
