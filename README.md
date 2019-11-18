# table
```
# install
go install github.com/scott-x/table
```
### usage
```
table
table -c
```
If you use `table`, it will create a html named `tb.html` in your current folder.

If you use `table -c`, -c here means "customed", first it will check whether `a.txt` exists or not? If not, it will throw tips as following:

```
You can define a.txt with syntax of Markdown, the cell content is separated by "|":
Suplier|Name|Email|Phone
QTOP|Lillian/Sue|xxx@qtop.com.cn / yyy@qtop.com.cn|Tel:(0)574 8921 xxxx
Danson|Lynn|zzz@dansondecor.com|Mobile : (86)15019468xxx Tel : (86) 0755-8272 xxxx
Merry Art|Tina|xxx@merryart.cn|TEL：+86 574 8790xxxx EXT：xxx
```
copy above txt into `a.txt`, then you will get a `table.html` in your current folder. Opening with chrome, see attached:

![](../imgs/1.png)

If `a.txt` exists, the code will parse it. Note that it must written with Markdown.