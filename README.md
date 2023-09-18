这个仓库里面是一些[walk](https://github.com/lxn/walk)框架的用例，因为[walk](https://github.com/lxn/walk)原项目已经很久没有维护了，里面一些用例有些过时，故此创建该仓库。
## 什么是[walk](https://github.com/lxn/walk)
[walk](https://github.com/lxn/walk)是一款基于windows gui封装的golang组件库，用于开发windows桌面应用。
[walk官方文档](https://pkg.go.dev/github.com/lxn/walk)
## 准备工作
在使用walk框架之前，您需要了解windows gui的基本工作原理，microsoft官网上有详细的讲解:[https://learn.microsoft.com/zh-cn/windows/win32/](https://learn.microsoft.com/zh-cn/windows/win32/) \
网上walk相关的资料较少，您可能需要研究项目源代码进行开发。笔者并不建议完全没有图形界面开发经验的人使用walk框架，您应该至少熟悉另外一种图形界面框架，不论是html/css亦或者qt。
## 为什么选择walk框架
尽管walk框架相关资料非常少，使用起来非常困难，但它是基于原生的windows api进行封装的纯go语言库，可以支持一些底层操作。因为不需要大量依赖第三方组件，编译出的二进制文件体积较小，优化更好。 \
~~以目前的go语言生态来看本来也没有什么成熟的ui框架，直接用原生框架可能反而更加友好一些~~
## TODO

- [X] action
- [X] clipboard
- [X] databinding
- [X] drawing
- [ ] draw svg
- [X] dropfiles
- [X] externalwidgets
- [ ] filebrowser
- [X] gradientcomposite
- [X] imageicon
- [X] imageview
- [X] imageviewer
- [X] img
- [X] linklabel
- [X] listbox
- [X] listbox_ownerdrawing
- [X] logview
- [X] multiplepages
- [X] notifyicon
- [ ] progressindicator [ui2walk工具不再维护](https://github.com/lxn/walk/issues/469)，用例无法运行
- [X] radiobutton
- [X] settings
- [X] slider
- [X] statubar
- [ ] tableview
- [X] webview 内核较老
- [X] webview_events
- [ ] 响应式布局
- [ ] 设计模式应用
- [ ] 更多用例