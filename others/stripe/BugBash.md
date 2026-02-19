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

  1. Issue 1 确认存在
     位置：mako/lookup.py:257
     问题：用 os.path.exists(srcfile) 判断后直接 _load，目录也会被当成模板文件加载。
     复现结果：lookup.get_template('subdir') 抛 PermissionError: [Errno 13] Permission denied（目录被当文件读）。
     影响：应当是 lookup miss/继续下个目录，而不是直接因为目录报错；和 test/test_lookup.py:38、test/test_lookup.py:61 的
     期望一致。
  2. Issue 2 确认存在
     位置：mako/_ast_util.py:392、mako/_ast_util.py:404、mako/_ast_util.py:282
     问题：SourceGenerator.signature() 对参数节点调用 self.visit(arg)，但缺少 visit_arg，回退到 generic_visit 不会写出参
     数名。
     复现结果：ArgumentList('lambda x, y=1: x+y').args 得到 ['lambda , =1: (x + y)']。
     影响：参数文本序列化丢失，args 写入不完整（可见于 mako/pyparser.py:186 的写入路径）。
  3. Issue 3 确认存在
     位置：mako/codegen.py:799、mako/codegen.py:818、mako/parsetree.py:31、mako/parsetree.py:165
     问题：visitControlLine 自动补 pass 的条件没覆盖“子节点全是 comment”的情况。Comment 节点生成阶段不产出语句（且无
     visitComment）。
     复现结果：模板
     % if True: / ## only comment / % endif
     渲染时报 IndentationError: expected an indented block...。
     影响：控制块仅注释时会生成非法 Python 代码。

## Java moshi

可以参考一下 branch moshi_10, 问题出在结束 object 或 array 的时候 indices[] 要加一 
Bug squash， Java的题是Moshi，找bug很简单，同一个bug在两个function里出现了，修起来也特别简单，一共写了两短行代码。这轮时间还是花了很多在setup上，最后还是Command Line跑的test，IDE里查函数改文件。找bug修bug一共10分钟，还是我看的很细的情况