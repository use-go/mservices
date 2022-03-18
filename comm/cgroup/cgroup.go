package cgroup

import "os"

// Share defined cgroup isol
func Isol(pid int, cpu uint64, mem int64) error {
	return nil
}

// Share defined cgroup share
func Share(pid []int, cpu uint64, mem int64) error {
	return nil
}

func IsolWithDefault() error {
	return Isol(os.Getpid(), 2000, 1024*500)
}

func ShareWithDefault() error {
	return Share([]int{os.Getpid()}, 2000, 1024*500)
}
