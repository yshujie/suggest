# suggest


## 简介
suggest 是一个搜索框联想系统，基于 TernaryTree 和 HashTable 实现，支持中文、拼音、拼音缩写、英文、手机号码等多种类型的联想。

## 模型设计
suggest 系统的设计中，最基本的模型元素为：WordBank（词库）、Dictionary（词典）、Entry（词条）   
而 suggest 系统又拆分成两个模块：suggest-core、suggest-search
 
### 模型元素
#### WordBank
- 定义：
  - 词库是一类数据的全集
- 举例：
  - 我们对外提供 医院词库、医生词库 等

#### Dictionary
- 定义：
  - 词典是一类数据的不同实现
- 举例：
  - 针对医院词库，我们会提供 HashDictionary、TernaryTreeDictionary 等不同的实现

#### Entry
- 定义：
  - 词条是词库中的最小元素
- 举例：
  - 医院词库中的词条为一家具体的医院，医院信息包含：医院名称、医院别名、医院地址、医院电话等


### suggest-core
- 定义：
  - suggest-core 是 suggest 系统的核心模块，提供了词库、词典、词条的基本实现
  - suggest-core 提供词库、词典、词条数据的维护、查询、删除等基本操作

### suggest-search
- 定义：
  - suggest-search 是 suggest 系统的搜索模块，提供了对 suggest-core 的封装，以及对 suggest-core 的扩展
  - suggest-search 提供了对 suggest-core 的封装，以便于用户快速使用 suggest 系统
