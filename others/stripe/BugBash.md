## Bug Squash：新题 snakeyaml

2 bugs:
1. 失败测试主要是没办法把“flag: On" parse成正确的Boolean.True value 
2. parse csv报错

面试官其实人很好，有很努力地在引导我，但楼主实在是太菜没有精确找到出错的那一部分，一个都没解决。这一轮可以多主动问面试官的，毕竟短时间了解一个repo实话说挺难的。这道java的debug感觉和地里的两道老题难度要高，jackson core和moshi的报错比较好定位到

## Mako

3 bugs:
Issue 1：有一处地方漏了检查读的是folder还是file
Issue 2：漏了一个visit function导致args没被写入
Issue 3：这个面试官明确表明不用做 但是答案是有一处地方漏了检测ControlLine是不是只包含comments

## Java moshi

可以参考一下 branch moshi_10, 问题出在结束 object 或 array 的时候 indices[] 要加一 