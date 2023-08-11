module poehelper

go 1.20

require (
	github.com/AllenDang/cimgui-go v0.0.0-20230809030658-606eb8908b6c
	github.com/hpcloud/tail v1.0.0
	github.com/sirupsen/logrus v1.9.3
	golang.org/x/exp v0.0.0-20230809150735-7b3493d9a819
)

require (
	github.com/BurntSushi/toml v1.3.2
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
)

replace github.com/AllenDang/cimgui-go => ..\cimgui-go
