[使用Gradle构建Go语言项目](http://www.jianshu.com/p/fd77b461b87c?from=timeline&isappinstalled=0)

1. gradle [gradle 原生指令]
2. gradle gogland [gogland 相关指令]

我们可以在build.gradle中声明所需的依赖，Gogradle会自动检索、下载所有的依赖包以及传递性依赖。下列代码给出了一些声明依赖的方式
```
dependencies {
    golang {
        build 'github.com/user/project'  // 不指定版本，默认使用最新版本
        build name:'github.com/user/project' // 和上一行等价

        build 'github.com/user/project@1.0.0-RELEASE' // 指定tag
        build name:'github.com/user/project', tag:'1.0.0-RELEASE' // 和上一行等价

        build name: 'github.com/user/project', url:'https://github.com/user/project.git', tag:'v1.0.0' // 指定url，例如镜像仓库

        test 'github.com/user/project#d3fbe10ecf7294331763e5c219bb5aa3a6a86e80' // 指定commit
        test name:'github.com/user/project', commit:'d3fbe10ecf7294331763e5c219bb5aa3a6a86e80' // 和上一行等价

        // 语义化版本：
        build 'github.com/user/project@1.*'  // Equivalent to >=1.0.0 & <2.0.0
        build 'github.com/user/project@1.x'  // Equivalent to last line
        build 'github.com/user/project@1.X'  // Equivalent to last line
        build 'github.com/user/project@~1.5' // Equivalent to >=1.5.0 & <1.6.0
        build 'github.com/user/project@1.0-2.0' // Equivalent to >=1.0.0 & <=2.0.0
        build 'github.com/user/project@^0.2.3' // Equivalent to >=0.2.3 & <0.3.0
        build 'github.com/user/project@1' // Equivalent to 1.X or >=1.0.0 & <2.0.0
        build 'github.com/user/project@!(1.x)' // Equivalent to <1.0.0 & >=2.0.0
        build 'github.com/user/project@ ~1.3 | (1.4.* & !=1.4.5) | ~2' // Very complicated expression

        build 'github.com/a/b@1.0.0', 'github.com/c/d@2.0.0', 'github.com/e/f#commitId' // 同时声明多个依赖

        // 声明一个依赖，禁止其所有传递性依赖
        build('github.com/user/project') {
            transitive = false
        }

        // 声明一个依赖，排除部分传递性依赖
        build('github.com/a/b') {
            exclude name:'github.com/c/d'
            exclude name:'github.com/c/d', tag: 'v1.0.0'
        }

        build name: 'github.com/big/package', subpackages: ['.', 'sub1', 'sub2/subsub'] // 只依赖这个包的部分子包
    }
}
```

单元测试 以 _test结尾，方法需要以Test开头 参数(t *testing.T)