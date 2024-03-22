package git

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/shipengqi/golib/cliutil"
	"github.com/shipengqi/golib/fsutil"
)

const (
	SubCmdPrefix = "git-"
)

func InitSubCmd(src, subname string) (string, error) {
	dir, err := ExecPath()
	if err != nil {
		return "", err
	}

	dst := filepath.Join(dir, SubCmdPrefix+subname)
	if err = fsutil.CopyFile(src, dst); err != nil {
		return "", err
	}
	return dst, nil
}

func IsGitRepo() (bool, error) {
	_, err := cliutil.ExecContext(context.TODO(), "git", "remote")
	if err != nil {
		return false, err
	}
	return true, nil
}

// WorkingTreeRoot return path of the top-level directory of the working tree
func WorkingTreeRoot() (path string, err error) {
	output, err := cliutil.ExecContext(context.TODO(), "git", "rev-parse", "--show-toplevel")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(output), nil
}

func ExecPath() (string, error) {
	stdout, err := cliutil.ExecContext(context.TODO(), "git", "--exec-path")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(stdout), nil
}

func Commit(msg []byte, opts *Options) (string, error) {
	temp, err := os.CreateTemp("", "COMMIT_MESSAGE_")
	if err != nil {
		return "", err
	}
	defer func() { _ = os.Remove(temp.Name()) }()
	if _, err = temp.Write(msg); err != nil {
		return "", err
	}

	args := opts.Combination(temp.Name())
	stdout, err := cliutil.ExecContext(context.TODO(), "git", args...)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(stdout), nil
}
