package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

var (
	daemon = flag.Bool("d", false, "run as a daemon")
)

func testMinimal() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	uid, err := strconv.ParseUint(u.Uid, 10, 64)
	if err != nil {
		panic(err)
	}
	gid, err := strconv.ParseUint(u.Gid, 10, 64)
	if err != nil {
		panic(err)
	}
	process, err := os.StartProcess("/bin/busybox", []string{"sh", "init"}, &os.ProcAttr{
		Dir:   "/",
		Env:   []string{"PATH=/bin"},
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Sys: &syscall.SysProcAttr{
			Chroot:     wd + "/var/linux-minimal",
			Noctty:     false,
			Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWNS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNET | syscall.CLONE_NEWUSER,
			UidMappings: []syscall.SysProcIDMap{
				{0, int(uid), 1},
			},
			GidMappings: []syscall.SysProcIDMap{
				{0, int(gid), 1},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	if _, err := process.Wait(); err != nil {
		panic(err)
	}
}

func runContainer(exe string, daemon bool) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	uid, err := strconv.ParseUint(u.Uid, 10, 64)
	if err != nil {
		panic(err)
	}
	gid, err := strconv.ParseUint(u.Gid, 10, 64)
	if err != nil {
		panic(err)
	}
	var files []*os.File = nil
	if !daemon {
		files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
	}
	process, err := os.StartProcess("/bin/busybox", []string{"sh", "init", "-c", "/exe/" + exe}, &os.ProcAttr{
		Dir:   "/",
		Env:   []string{"PATH=/bin"},
		Files: files,
		Sys: &syscall.SysProcAttr{
			Chroot:     wd + "/var/linux-minimal",
			Cloneflags: syscall.CLONE_NEWNET | syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWNS | syscall.CLONE_NEWPID | syscall.CLONE_NEWUSER,
			UidMappings: []syscall.SysProcIDMap{
				{0, int(uid), 1},
			},
			GidMappings: []syscall.SysProcIDMap{
				{0, int(gid), 1},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	if !daemon {
		if _, err := process.Wait(); err != nil {
			panic(err)
		}
	} else {
		fmt.Println(process.Pid)
	}
}

func main() {
	// flag.Parse()
	// executable := flag.Arg(0)
	// filename := filepath.Base(executable)
	// wd, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }
	// src, err := os.Open(executable)
	// if err != nil {
	// 	panic(err)
	// }
	// exe := fmt.Sprintf("%s/var/linux-minimal/exe/%s", wd, filename)
	// dst, err := os.OpenFile(exe, os.O_WRONLY|os.O_CREATE, 0777)
	// if err != nil {
	// 	panic(err)
	// }
	// _, err = io.Copy(dst, src)
	// if err != nil {
	// 	panic(err)
	// }
	// src.Close()
	// dst.Close()

	testMinimal()
}
