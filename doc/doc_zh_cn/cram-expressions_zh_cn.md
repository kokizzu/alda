# Cram表达式 (Cram expressions)

*本页面翻译自[Cram Expressions](../cram-expressions.md)*

**Cram表达式**是Alda中表示节奏的另一种方式 它对n连音和复节奏很有用

这个想法是 您可以将一堆[音符](notes_zh_cn.md)"塞进"单个音符的时值 例如 您可能想要把5个音符"塞进"二分音符的持续时间:

```alda
{c d e f g}2
```

与一般音符一样 如果您没有在cram表达式的末尾指定其时值 它会使用最后指定过时值的音符的时值 例如 在下面的例子中 大括号之间的音符会被塞进全音符的持续时间 因为这是最后指定过的时值

```alda
c1 e {g a b} > c
```

您还可以在cram表达式内部添加上音符长度 这样可以使较长的音符相对于整个cram所分配的时间更长 整个cram的持续时间不会改变

```alda
{c d e}2 {c2 d4 e} {c1 d4 e}
```

> 默认情况下 每个cram表达式的第一个音符都是四分音符

Cram表达式可以嵌套 每个子cram将在它的父cram表达式中占用适当的空间

```alda
{c e {g a b}}1 c
```

