因为go里面没有继承这一说法, 而又要使用gin内部的
绑定数据方法却又不想使用自带的v10校验的功能, 此外:
mapUri,mapForm等函数是私有的函数, 外部无法调用. 所以只能把源代码
copy出来自己做做修改, 方法比较笨...
没找到比较好的方法. 