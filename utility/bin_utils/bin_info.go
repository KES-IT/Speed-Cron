package binInfo

import (
	"fmt"
	"runtime"
	"strings"
)

// 初始化为 unknown，如果编译时没有传入这些值，则为 unknown
var (
	GitTag         = "unknown"
	GitCommitLog   = "unknown"
	GitStatus      = "cleanly"
	BuildTime      = "unknown"
	BuildGoVersion = "unknown"
)

var (
	VersionString = "GitTag:" + GitTag + "\n" +
		"GitCommitLog:" + GitCommitLog + "\n" +
		"GitStatus:" + GitStatus + "\n" +
		"BuildTime:" + BuildTime + "\n" +
		"BuildGoVersion:" + BuildGoVersion + "\n"
)

/*// StringifySingleLine 返回单行格式
func StringifySingleLine() string {
	return fmt.Sprintf("GitTag=%s. GitCommitLog=%s. GitStatus=%s. BuildTime=%s. GoVersion=%s. runtime=%s/%s.",
		GitTag, GitCommitLog, GitStatus, BuildTime, BuildGoVersion, runtime.GOOS, runtime.GOARCH)
}*/

// StringifyMultiLine 返回多行格式
func StringifyMultiLine() string {
	return fmt.Sprintf("GitTag=%s\nGitCommitLog=%s\nGitStatus=%s\nBuildTime=%s\nGoVersion=%s\nruntime=%s/%s\n",
		GitTag, GitCommitLog, GitStatus, BuildTime, BuildGoVersion, runtime.GOOS, runtime.GOARCH)
}

func init() {
	beauty()
}

// 对一些值做美化处理
func beauty() {
	if GitStatus == "" {
		// GitStatus 为空时，说明本地源码与最近的 commit 记录一致，无修改
		// 为它赋一个特殊值
		GitStatus = "cleanly"
	} else {
		// 将多行结果合并为一行
		GitStatus = strings.Replace(strings.Replace(GitStatus, "\r\n", " |", -1), "\n", " |", -1)
	}
}
