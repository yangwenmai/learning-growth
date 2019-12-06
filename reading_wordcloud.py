#coding=utf-8
import jieba
import wordcloud
# 构建并配置词云对象w
w = wordcloud.WordCloud(width=1000,
                        height=700,
                        background_color='white',
                        font_path='STHeiti Light.ttc',
                        scale=3,
                        stopwords={'我们','一个','可以', '这个','那个','自己','就是','问题','如果','用户','可能','些','这些','那些','需要','不同', '现在','通过', '不是', '还是', '所以', '能','因为','使用','他们','大家','应该','或者','时候','然后','一样','成为','其实','开始'})

# stopwords 貌似不能生效。

# 调用jieba的lcut()方法对原始文本进行中文分词，得到string
# 读取 reading 文章的所有文本内容
f = open('data.txt', 'r')
txt = f.read()
txtlist = jieba.lcut(txt)
string = " ".join(txtlist)

# 将string变量传入w的generate()方法，给词云输入文字
w.generate(string)

# 将词云图片导出到当前文件夹
w.to_file('reading_wordcloud_output.png')
