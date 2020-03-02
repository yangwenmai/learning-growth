### Go image history

### image 包的描述：

image实现了基本的2D图片库。

基本接口叫作Image。

图片的色彩定义在image/color包。

Image接口可以通过调用如NewRGBA和NewPaletted函数等获得；也可以通过调用Decode函数解码包含GIF、JPEG或PNG格式图像数据的输入流获得。解码任何具体图像类型之前都必须注册对应类型的解码函数。注册过程一般是作为包初始化的副作用，放在包的init函数里。因此，要解码PNG图像，只需在程序的main包里嵌入如下代码：

```
import _ "image/png"
```

_表示导入包但不使用包中的变量/函数/类型，只是为了包初始化函数的副作用。

官方博客文档：https://blog.golang.org/go-image-package 这一篇文章讲的非常好。

>Nigel Tao
>21 September 2011
>
>A common mistake is assuming that an `Image`'s bounds start at (0, 0)



### image 迁移之前的 commit

*[https://github.com/golang/go/commits/a2c30fe6481fa5b18331bb54a660928f6816cc44?after=a2c30fe6481fa5b18331bb54a660928f6816cc44+174&path%5B%5D=src&path%5B%5D=pkg&path%5B%5D=image](https://github.com/golang/go/commits/a2c30fe6481fa5b18331bb54a660928f6816cc44?after=a2c30fe6481fa5b18331bb54a660928f6816cc44+174&path[]=src&path[]=pkg&path[]=image)*



## 开发 image 相关的时间线

### Introduce the image package: https://github.com/golang/go/commit/5cbc96d958e7026b6d2c82c947e24e6159e57564

> commit date: 2009.08.27 12:51:00

image 的第一个版本： `src/pkg/image/color.go` 和 `src/pkg/image/image.go` 两个文件。

image 的第二个版本：png decoder for go；

image 刚开始实现的时候并没有 gofmt？那gofmt 是什么时候才有的呢？

> [apply gofmt to go, gob, hash, http, image, io, json, log](https://github.com/golang/go/commit/c2ec9583a0abe0a74d5cb8dd01631ca0911aa3aa#diff-800b7d7968ca70cafdfa0a33819de7da) 
>
> 是在这个coomit 吗？
>
> 2009.10.07 10:41:00

rsc 一直都在处理一下全局项目的事情，也会涉及到 image 库。

[griesemer](https://github.com/griesemer) 在 2010-05-16 修复了 image 库的拼写错误。

[mpl](https://github.com/mpl) 在2010-08-10 给 image 库增加了 grayscale 支持。

[rogpeppe](https://github.com/rogpeppe) 在 2010-09-23 给 image 库增加了 mul, div, eq size 函数。

[adg](https://github.com/adg) 在 2010-10-01 给 image 库修改了合适的 fmt.Errorf 处理。

2010-10-28 才由 rsc 提交 commit 支持使用 append，也就是在这个时候， append 才开发完成？

2011-02-22 由 Mikael Tillenius 支持了更多格式。（带 grayscale+alpha读取 image，深度为 1，2，4的图片）

2011-04-30 由 bradfitz 添加 png 和 jpeg 的 benchmark 测试用例。

2011-05-02 由 bradfitz 又进行了优化，RGBA 编码效果提升 ~50%，（https://github.com/golang/go/commit/807ce2719c243f6d8640de774599e3335883eacb#diff-800b7d7968ca70cafdfa0a33819de7da）

> 避免 image.At(), color.RGBA(), 8位移位和min函数的循环调用。

2011-05-03 由 robpikie 优化，image: 添加类型特定的设置方法，并在解码PNG时使用它们，可以使得 png 解码速度提高约 20%。（https://github.com/golang/go/commit/2398a74bd80ea945f687b3750fa3e18c258945eb#diff-800b7d7968ca70cafdfa0a33819de7da）

> CL 中，没有任何 commit, LGTM 即可。

2011-05-05 [dchest](https://github.com/dchest) 支持使用Alpha通道对调色板图像进行编码。

2011-05-05 [bsiegert](https://github.com/bsiegert) 支持 image/tiff 。

2011-05-08 [rob pike](https://github.com/robpike) 支持 gif decoder。

rob pike 也有逻辑 bug （https://github.com/golang/go/commit/45ea58746bf7010d7d993e942d136394b780325c#diff-800b7d7968ca70cafdfa0a33819de7da）

2011-05-12 nigeltao 实现 bmp decoder。

2011-05-19 nigeltao 对 image / jpeg：用于编码的小内存布局优化。

> 将 uBits uint8 改为 uint32

2011-06-02 nigeltao 支持 SubImage 函数。

2011-07-12 nigeltao 添加的 image/draw benchmark 跟最终的模式比较像了。

2011-08-26 bsiegert 对 image/tiff：decoder 优化（RGB 图像提升 3倍，NRGBA 图像提升 6倍。）

2011-10-04 nigeltao 拆分出 color 相关到 color package 中。

2012-01-31 nigeltao 把 image/bmp 和 image/tiff 从标准库中移除。（https://github.com/golang/go/commit/10498f4d335f6bf0089791b263e795233ff79ec5#diff-800b7d7968ca70cafdfa0a33819de7da）

2012-02-03 nigeltao 给标准库添加结构体字段标签未加标签的文本（包括 image package）。(std: add struct field tags to untagged literals.)

2012-02-06 robpike 避免 bytes.NewBuffer(nil) ，改为 var buf bytes.Buffer 和 new(bytes.Buffer) 让人们相信这是最佳实践。

2012-04-27 nigeltao 为通用的颜色模型：Gray, NRGBA, 调色板，RGBA 提升 png 编码性能。

2012-05-25, 2012-05-30 nigeltao 分别优化 paeth 相关，提升相关性能。

2012-06-28 mpl 支持读取 4:4:0 二次抽样（subsampling）。（来自于一次 issues: https://github.com/golang/go/issues/2362）

2012-09-13 nigeltao 优化编码 image.Gray 和 image.NRGBA 图片。

2012-10-07 nigeltao 做了3次优化，提升了不同的逻辑性能，很小的优化点。（变量预计算，变量在 struct 的位置等）

2012-10-19 remyoudompheng 做了 TestDCT 性能提升；

2012-10-30 nigeltao 改变 block [64]int 为 [64]int32，性能提升；

2012-10-31 nigeltao 减少 ensureNBits 的调用，提升性能；



### 优化一：image/png: speed up paletted encoding ~25%

> https://github.com/golang/go/commit/ad7d24ac4b66f9ee83e1cb5f81c3ae4cb3eb5610#diff-800b7d7968ca70cafdfa0a33819de7da

怎么优化的呢？就是将 for 循环改为一次性 copy 处理，减少许多不必要的 bounds checks.

这个优化是由 bradfitz 在 2010.10.27 提交的。

> review 的过程非常快，第一个commit是 2010-10-27 06:40:13 UTC, 06:43:31 UTC 分别得到2个 LGTM ，11:48:36 UTC 就合并了。

### 优化二：image/png: speed up PNG decoding for common color models: Gray, NRGBA, Paletted, RGBA

https://github.com/golang/go/commit/dd294fbd5a6eea9574df8c3f842342a8cd10f2c6#diff-800b7d7968ca70cafdfa0a33819de7da

>2012-04-27
>
>提升1.4倍以上。

```
benchmark                       old ns/op    new ns/op    delta
BenchmarkDecodeGray               3681144      2536049  -31.11%
BenchmarkDecodeNRGBAGradient     12108660     10020650  -17.24%
BenchmarkDecodeNRGBAOpaque       10699230      8677165  -18.90%
BenchmarkDecodePaletted           2562806      1458798  -43.08%
BenchmarkDecodeRGB                8468175      7180730  -15.20%

benchmark                        old MB/s     new MB/s  speedup
BenchmarkDecodeGray                 17.80        25.84    1.45x
BenchmarkDecodeNRGBAGradient        21.65        26.16    1.21x
BenchmarkDecodeNRGBAOpaque          24.50        30.21    1.23x
BenchmarkDecodePaletted             25.57        44.92    1.76x
BenchmarkDecodeRGB                  30.96        36.51    1.18x

$ file $GOROOT/src/pkg/image/png/testdata/bench*
benchGray.png:           PNG image, 256 x 256, 8-bit grayscale, non-interlaced
benchNRGBA-gradient.png: PNG image, 256 x 256, 8-bit/color RGBA, non-interlaced
benchNRGBA-opaque.png:   PNG image, 256 x 256, 8-bit/color RGBA, non-interlaced
benchPaletted.png:       PNG image, 256 x 256, 8-bit colormap, non-interlaced
benchRGB.png:            PNG image, 256 x 256, 8-bit/color RGB, non-interlaced
```

for 循环改为 copy 预设值。

### 优化三：image/png: optimize the paeth filter implementation.

https://github.com/golang/go/commit/1423ecb1266c9af288caa2723988a326adf7118e#diff-800b7d7968ca70cafdfa0a33819de7da

> 2012-05-25

```
image/png benchmarks:
benchmark                       old ns/op    new ns/op    delta
BenchmarkPaeth                         10            7  -29.21%
BenchmarkDecodeGray               2381745      2241620   -5.88%
BenchmarkDecodeNRGBAGradient      9535555      8835100   -7.35%
BenchmarkDecodeNRGBAOpaque        8189590      7611865   -7.05%
BenchmarkDecodePaletted           1300688      1301940   +0.10%
BenchmarkDecodeRGB                6760146      6317082   -6.55%
BenchmarkEncodePaletted           6048596      6122666   +1.22%
BenchmarkEncodeRGBOpaque         18891140     19474230   +3.09%
BenchmarkEncodeRGBA              78945350     78552600   -0.50%
```

有一些案例会因此而增长了时间，但是也确实有案例因此而减少了时间。

### 优化四：image/png: optimize paeth some more.

https://github.com/golang/go/commit/dbcdce5866502aadc9b3e70e06c92d2afb22e1e1#diff-800b7d7968ca70cafdfa0a33819de7da

> 2012-05-30

```
benchmark                       old ns/op    new ns/op    delta
BenchmarkDecodeGray               3139636      2812531  -10.42%
BenchmarkDecodeNRGBAGradient     12341520     10971680  -11.10%
BenchmarkDecodeNRGBAOpaque       10740780      9612455  -10.51%
BenchmarkDecodePaletted           1819535      1818913   -0.03%
BenchmarkDecodeRGB                8974695      8178070   -8.88%
```

filterPaeth 使用 []byte 参数来替代 byte 参数，避免一些先前像素在循环体内的冗余计算。

### 优化五：image/png: optimize encoding image.Gray and image.NRGBA images.

https://github.com/golang/go/commit/237ee3926906ad08a048e764920e036ecdb08b11#diff-800b7d7968ca70cafdfa0a33819de7da

> 2012-09-13
>
> Delta 76%

```
benchmark                    old ns/op    new ns/op    delta
BenchmarkEncodeGray           23616080      5624558  -76.18%
BenchmarkEncodeNRGBOpaque     34181260     17144380  -49.84%
BenchmarkEncodeNRGBA          41235820     20345990  -50.66%
BenchmarkEncodePaletted        5594652      5620362   +0.46%
BenchmarkEncodeRGBOpaque      17242210     17168820   -0.43%
BenchmarkEncodeRGBA           66515720     67243560   +1.09%
```

### 优化六：image/jpeg: add DCT tests, do a small optimization (common sub-expression elimination) in idct.go

> 2012-10-07

```
benchmark                   old ns/op    new ns/op    delta
BenchmarkIDCT                    5649         5610   -0.69%
BenchmarkDecodeRGBOpaque      2948607      2941051   -0.26%
```

y8 := y * 8 替换掉每次单独计算。

### 优化七：image/jpeg: move the level-shift and clip out of the idct function

> 2012-10-07

```
benchmark                   old ns/op    new ns/op    delta
BenchmarkFDCT                    3842         3837   -0.13%
BenchmarkIDCT                    5611         3478  -38.01%
BenchmarkDecodeRGBOpaque      2932785      2929751   -0.10%
```

### 优化八：image/jpeg: move the huffman bit decoder state higher up in the decoder struct.

> 2012-10-07

```
benchmark          old ns/op    new ns/op    delta
BenchmarkDecode      2943204      2746360   -6.69%
```

move the huffman bit decoder state higher up in the decoder struct, inside the unmappedzero limit, to eliminate some TESTB instructions in the inner decoding loop.

将 b bits 移到 decoder struct 的最前面，减少 TESTB 指令在内部循环的开销。

### 优化九：image/jpeg: make TestDCT faster.

> 2012-10-19

math.cos 提前计算，可以使用全局缓存数组来替代每一次的重新计算。

### 优化十：image/jpeg: change block from [64]int to [64]int32.

https://github.com/golang/go/commit/daf43ba476ac29c9c15b59169a9458900efa0e1d#diff-800b7d7968ca70cafdfa0a33819de7da

> 2012-10-30

```
On 6g/linux:
benchmark                     old ns/op    new ns/op    delta
BenchmarkFDCT                      4606         4241   -7.92%
BenchmarkIDCT                      4187         3923   -6.31%
BenchmarkDecodeBaseline         3154864      3170224   +0.49%
BenchmarkDecodeProgressive      4072812      4017132   -1.37%
BenchmarkEncode                39406920     34596760  -12.21%
```

### 优化十一：image/jpeg: don't call ensureNBits unless we have to.

https://github.com/golang/go/commit/ad487dad75faca0c5cd6a152d9f04d9ff93aaff5#diff-800b7d7968ca70cafdfa0a33819de7da

> 2012-10-31

```
benchmark                     old ns/op    new ns/op    delta
BenchmarkDecodeBaseline         3155638      2783998  -11.78%
BenchmarkDecodeProgressive      4008088      3660310   -8.68%
```

除非必要，否则不必调用 ensureNBits。



----

> 2014-09-08开始 Go 从 src/pkg/image 迁移到 src/image/ 目录，去掉了 pkg 目录。

### image/color: optimize YCbCrToRGB

https://github.com/golang/go/commit/f0c5b8b9c9c7900033ddb11b584da6198d599454#diff-d90c5b811de11d0d53e5d77f2162f6eb

> 2016.4.13

```
name     old time/op  new time/op  delta
YCbCr-2  1.12ms ± 0%  0.64ms ± 0%  -43.01%  (p=0.000 n=48+47)

name              old time/op  new time/op  delta
YCbCrToRGB/0-2    5.52ns ± 0%  5.77ns ± 0%  +4.48%  (p=0.000 n=50+49)
YCbCrToRGB/128-2  6.05ns ± 0%  5.52ns ± 0%  -8.69%  (p=0.000 n=39+50)
YCbCrToRGB/255-2  5.80ns ± 0%  5.77ns ± 0%  -0.58%  (p=0.000 n=50+49)
```

使用一个比较来同时检测下溢和上溢。

使用移位，按位补码和uint8类型转换来处理夹紧到上下限而无其他分支。

### image/color: improve speed of RGBA methods

https://github.com/golang/go/commit/2113c9ad0d82b3d1a734c0b5fc0efc9c44a920d5#diff-d90c5b811de11d0d53e5d77f2162f6eb

> 2016.10.18

```
YCbCrToRGBA/0-4      6.32ns ± 3%  6.58ns ± 2%   +4.15%  (p=0.000 n=20+19)
YCbCrToRGBA/128-4    8.02ns ± 2%  5.89ns ± 2%  -26.57%  (p=0.000 n=20+19)
YCbCrToRGBA/255-4    8.06ns ± 2%  6.59ns ± 3%  -18.18%  (p=0.000 n=20+20)
NYCbCrAToRGBA/0-4    8.71ns ± 2%  8.78ns ± 2%   +0.86%  (p=0.036 n=19+20)
NYCbCrAToRGBA/128-4  10.3ns ± 4%   7.9ns ± 2%  -23.44%  (p=0.000 n=20+20)
NYCbCrAToRGBA/255-4  9.64ns ± 2%  8.79ns ± 3%   -8.80%  (p=0.000 n=20+20)
```

这个优化是应用：image/color: optimize YCbCrToRGB 的优化实现。



### image/jpeg: improve performance when encoding *image.YCbCr

https://github.com/golang/go/commit/435450bf3c6efcc65111e96a42fc1c8acd3081e3#diff-d90c5b811de11d0d53e5d77f2162f6eb

> 2017.1.2

```
benchmark                  old ns/op     new ns/op     delta
BenchmarkEncodeYCbCr-4     43990846      24201148      -44.99%

benchmark                  old MB/s     new MB/s     speedup
BenchmarkEncodeYCbCr-4     20.95        38.08        1.82x
```

现有的实现方式会回落到使用 image.At() 对于编码 * image.YCbCr 时的每个像素效率低下并导致许多内存分配。

增加 `yCbCrToYCbCr()` 从 image.YCbCr 转换为 toYCbCr 的特殊处理。

### image/gif: write fewer, bigger blocks

https://github.com/golang/go/commit/8b220d8ef1ad8fdedd2728fe047ec7c2f55e8aa6#diff-d90c5b811de11d0d53e5d77f2162f6eb

> 2017.10.5

```
BenchmarkEncode indicates slight improvements:

name      old time/op    new time/op    delta
Encode-8    7.71ms ± 0%    7.38ms ± 0%   -4.27%  (p=0.008 n=5+5)

name      old speed      new speed      delta
Encode-8   159MB/s ± 0%   167MB/s ± 0%   +4.46%  (p=0.008 n=5+5)

name      old alloc/op   new alloc/op   delta
Encode-8    84.1kB ± 0%    80.0kB ± 0%   -4.94%  (p=0.008 n=5+5)

name      old allocs/op  new allocs/op  delta
Encode-8      9.00 ± 0%      7.00 ± 0%  -22.22%  (p=0.008 n=5+5)
```



### image: optimize bounds checking for At and Set methods

https://github.com/golang/go/commit/b57ccdf992b46b15c33cf4672de4a7911d667617#diff-d90c5b811de11d0d53e5d77f2162f6eb

> 2018.9.23

```
name                   old time/op    new time/op    delta
At/rgba-8              18.8ns ± 2%    18.5ns ± 1%   -1.49%  (p=0.026 n=10+10)
At/rgba64-8            22.2ns ± 2%    21.1ns ± 3%   -4.51%  (p=0.000 n=10+10)
At/nrgba-8             18.8ns ± 2%    18.7ns ± 2%     ~     (p=0.467 n=10+10)
At/nrgba64-8           21.9ns ± 2%    21.0ns ± 2%   -4.15%  (p=0.000 n=10+9)
At/alpha-8             14.3ns ± 1%    14.3ns ± 2%     ~     (p=0.543 n=10+10)
At/alpha16-8           6.43ns ± 1%    6.47ns ± 1%     ~     (p=0.053 n=10+10)
At/gray-8              14.4ns ± 2%    14.6ns ± 5%     ~     (p=0.194 n=10+10)
At/gray16-8            6.52ns ± 1%    6.55ns ± 2%     ~     (p=0.610 n=10+10)
At/paletted-8          4.17ns ± 1%    4.21ns ± 2%     ~     (p=0.095 n=9+10)
Set/rgba-8             39.2ns ± 2%    40.1ns ± 4%   +2.45%  (p=0.007 n=10+10)
Set/rgba64-8           46.2ns ± 3%    43.3ns ± 3%   -6.11%  (p=0.000 n=10+10)
Set/nrgba-8            39.2ns ± 1%    39.7ns ± 5%     ~     (p=0.407 n=10+10)
Set/nrgba64-8          45.6ns ± 3%    42.9ns ± 3%   -5.83%  (p=0.000 n=9+10)
Set/alpha-8            35.0ns ± 3%    34.1ns ± 2%   -2.43%  (p=0.017 n=10+10)
Set/alpha16-8          36.3ns ± 4%    35.8ns ± 5%     ~     (p=0.254 n=10+10)
Set/gray-8             19.8ns ± 1%    19.7ns ± 0%   -0.69%  (p=0.002 n=8+6)
Set/gray16-8           36.0ns ± 1%    36.4ns ± 2%   +1.08%  (p=0.037 n=10+10)
Set/paletted-8         39.1ns ± 0%    39.6ns ± 1%   +1.30%  (p=0.000 n=10+10)
RGBAAt-8               3.72ns ± 1%    3.58ns ± 1%   -3.76%  (p=0.000 n=9+10)
RGBASetRGBA-8          4.35ns ± 1%    3.70ns ± 1%  -14.92%  (p=0.000 n=10+10)
RGBA64At-8             5.08ns ± 1%    3.69ns ± 1%  -27.40%  (p=0.000 n=9+9)
RGBA64SetRGBA64-8      6.65ns ± 2%    3.63ns ± 0%  -45.35%  (p=0.000 n=10+9)
NRGBAAt-8              3.72ns ± 1%    3.59ns ± 1%   -3.55%  (p=0.000 n=10+10)
NRGBASetNRGBA-8        4.05ns ± 0%    3.71ns ± 1%   -8.57%  (p=0.000 n=9+10)
NRGBA64At-8            4.99ns ± 1%    3.69ns ± 0%  -26.07%  (p=0.000 n=10+9)
NRGBA64SetNRGBA64-8    6.66ns ± 1%    3.64ns ± 1%  -45.40%  (p=0.000 n=10+10)
AlphaAt-8              1.44ns ± 1%    1.44ns ± 0%     ~     (p=0.176 n=10+7)
AlphaSetAlpha-8        1.60ns ± 2%    1.56ns ± 0%   -2.62%  (p=0.000 n=10+6)
Alpha16At-8            2.87ns ± 1%    2.92ns ± 2%   +1.67%  (p=0.001 n=10+10)
AlphaSetAlpha16-8      3.26ns ± 1%    3.35ns ± 1%   +2.68%  (p=0.012 n=8+3)
```

针对此优化有一个讨论：https://github.com/golang/go/issues/27857

### image/jpeg: reduce bound checks from idct and fdct

https://github.com/golang/go/commit/64f22e4bd6a1c7fe8a2dcf52cc8ac4c39d5abbb4#diff-d90c5b811de11d0d53e5d77f2162f6eb

> 2019.1.8

```
name                 old time/op    new time/op    delta
FDCT-4                 1.85µs ± 2%    1.74µs ± 1%  -5.95%  (p=0.000 n=10+10)
IDCT-4                 1.94µs ± 2%    1.89µs ± 1%  -2.67%  (p=0.000 n=10+9)
DecodeBaseline-4       1.45ms ± 2%    1.46ms ± 1%    ~     (p=0.156 n=9+10)
DecodeProgressive-4    2.21ms ± 1%    2.21ms ± 1%    ~     (p=0.796 n=10+10)
EncodeRGBA-4           24.9ms ± 1%    25.0ms ± 1%    ~     (p=0.075 n=10+10)
EncodeYCbCr-4          26.1ms ± 1%    26.2ms ± 1%    ~     (p=0.573 n=8+10)

name                 old speed      new speed      delta
DecodeBaseline-4     42.5MB/s ± 2%  42.4MB/s ± 1%    ~     (p=0.162 n=9+10)
DecodeProgressive-4  27.9MB/s ± 1%  27.9MB/s ± 1%    ~     (p=0.796 n=10+10)
EncodeRGBA-4         49.4MB/s ± 1%  49.1MB/s ± 1%    ~     (p=0.066 n=10+10)
EncodeYCbCr-4        35.3MB/s ± 1%  35.2MB/s ± 1%    ~     (p=0.586 n=8+10)

name                 old alloc/op   new alloc/op   delta
DecodeBaseline-4       63.0kB ± 0%    63.0kB ± 0%    ~     (all equal)
DecodeProgressive-4     260kB ± 0%     260kB ± 0%    ~     (all equal)
EncodeRGBA-4           4.40kB ± 0%    4.40kB ± 0%    ~     (all equal)
EncodeYCbCr-4          4.40kB ± 0%    4.40kB ± 0%    ~     (all equal)

name                 old allocs/op  new allocs/op  delta
DecodeBaseline-4         5.00 ± 0%      5.00 ± 0%    ~     (all equal)
DecodeProgressive-4      13.0 ± 0%      13.0 ± 0%    ~     (all equal)
EncodeRGBA-4             4.00 ± 0%      4.00 ± 0%    ~     (all equal)
EncodeYCbCr-4            4.00 ± 0%      4.00 ± 0%    ~     (all equal)
```

通过 `go build -gcflags="-d=ssa/check_bce/debug=1" fdct.go idct.go` 来寻找需要进行边界检查消除的代码。

引申出来一个有关性能问题的权衡思考：

> go/issues/24499: image/jpeg: Decode is slow
>
> https://github.com/golang/go/issues/24499



### image/png: hoist repetitive pixels per byte out of loop in Encode

https://github.com/golang/go/commit/f04f594a1fd1e8cf1bd01f9dc62599d5bd6e1d92#diff-d90c5b811de11d0d53e5d77f2162f6eb

> 2019.7.24

```
name                        old time/op    new time/op     delta
EncodeGray-4                  2.16ms ± 1%     2.16ms ± 1%    -0.28%  (p=0.000 n=86+84)
EncodeGrayWithBufferPool-4    1.99ms ± 0%     1.97ms ± 0%    -0.72%  (p=0.000 n=97+92)
EncodeNRGBOpaque-4            6.51ms ± 1%     6.48ms ± 1%    -0.45%  (p=0.000 n=90+85)
EncodeNRGBA-4                 7.33ms ± 1%     7.28ms ± 0%    -0.69%  (p=0.000 n=89+87)
EncodePaletted-4              5.10ms ± 1%     2.29ms ± 0%   -55.11%  (p=0.000 n=90+85)
EncodeRGBOpaque-4             6.51ms ± 1%     6.51ms ± 0%      ~     (p=0.311 n=94+88)
EncodeRGBA-4                  24.3ms ± 2%     24.1ms ± 1%    -0.87%  (p=0.000 n=91+91)

name                        old speed      new speed       delta
EncodeGray-4                 142MB/s ± 1%    143MB/s ± 1%    +0.26%  (p=0.000 n=86+85)
EncodeGrayWithBufferPool-4   154MB/s ± 0%    156MB/s ± 0%    +0.73%  (p=0.000 n=97+92)
EncodeNRGBOpaque-4           189MB/s ± 1%    190MB/s ± 1%    +0.44%  (p=0.000 n=90+86)
EncodeNRGBA-4                168MB/s ± 1%    169MB/s ± 0%    +0.69%  (p=0.000 n=89+87)
EncodePaletted-4            60.3MB/s ± 1%  134.2MB/s ± 0%  +122.74%  (p=0.000 n=90+85)
EncodeRGBOpaque-4            189MB/s ± 1%    189MB/s ± 0%      ~     (p=0.326 n=94+88)
EncodeRGBA-4                50.6MB/s ± 2%   51.1MB/s ± 1%    +0.87%  (p=0.000 n=91+91)

name                        old alloc/op   new alloc/op    delta
EncodeGray-4                   852kB ± 0%      852kB ± 0%    +0.00%  (p=0.000 n=100+100)
EncodeGrayWithBufferPool-4    1.49kB ± 2%     1.47kB ± 1%    -0.88%  (p=0.000 n=95+90)
EncodeNRGBOpaque-4             860kB ± 0%      860kB ± 0%    +0.00%  (p=0.003 n=98+58)
EncodeNRGBA-4                  864kB ± 0%      864kB ± 0%    +0.00%  (p=0.021 n=100+99)
EncodePaletted-4               849kB ± 0%      849kB ± 0%    +0.00%  (p=0.040 n=100+100)
EncodeRGBOpaque-4              860kB ± 0%      860kB ± 0%      ~     (p=0.062 n=66+98)
EncodeRGBA-4                  3.32MB ± 0%     3.32MB ± 0%    -0.00%  (p=0.044 n=99+99)

name                        old allocs/op  new allocs/op   delta
EncodeGray-4                    32.0 ± 0%       32.0 ± 0%      ~     (all equal)
EncodeGrayWithBufferPool-4      3.00 ± 0%       3.00 ± 0%      ~     (all equal)
EncodeNRGBOpaque-4              32.0 ± 0%       32.0 ± 0%      ~     (all equal)
EncodeNRGBA-4                   32.0 ± 0%       32.0 ± 0%      ~     (all equal)
EncodePaletted-4                36.0 ± 0%       36.0 ± 0%      ~     (all equal)
EncodeRGBOpaque-4               32.0 ± 0%       32.0 ± 0%      ~     (all equal)
EncodeRGBA-4                    614k ± 0%       614k ± 0%      ~     (all equal)
```

现有的实现已为每个像素计算了每个字节的像素。

> 减少每字节像素的计算。



----

以下主要是 image/draw 的部分：

### 从 Go 1.4 src/image/draw 说起

image/draw 的第一个 commit 是由 Russ Cox 在 2014年9月8日 由 `src/pkg/image/draw` 迁移到 `src/image/draw`。

为什么要迁移呢？Russ Cox June 2014 专门写了一个文档来阐述《Go 1.4 src/pkg -> src》：https://docs.google.com/document/d/12w7ksAXmo3sF1EWlbAp4adrm1kJoe_8WIFlaZshIHWM/edit

> 文档的来源：build: move package sources from src/pkg to src
>
> ```
> Preparation was in CL 134570043.
> This CL contains only the effect of 'hg mv src/pkg/* src'.
> For more about the move, see golang.org/s/go14nopkg.
> ```

从这个 review commit 可以看出来，Go 开始是使用 hg 的，他使用的命令是：`hg mv src/pkg/* src`

文档的主要内容：

> 在2008年7月，我们开始将更大的包(如syscall)放在它们自己的目录中。小的包保存在单个文件中。
>
> 在2009年5月，我们决定将所有包(无论大小)放在它们自己的目录中。
>
> 2009年6月，我们将src/lib重命名为src/pkg，与编译包所在的目录名匹配。在其他src目录中仍然有相当数量的非go。
>
> 2011年4月，经过几个月的讨论，我们引入了$GOPATH，允许开发人员处理标准树之外的代码。$GOPATH目录的形式模仿了主要的 repo，但是因为它是所有 Go，所以我们将 src/pkg 缩短为src。额外的“pkg”是go命令中的一个特例，它使 $GOROOT 工作区与 $GOPATH 工作区不同。



了解了其中的一些花絮，接下来我们就来了解一下 `image/draw` 的历史吧。



### 一个 nil 空指针的问题，引发的相关改进



image/draw 的 nil 空指针问题 （https://github.com/golang/go/commit/64e6fe2d29c3f8ae1e3e38a09bfb92b0452ad51b#diff-1cbc7137c670aec930ba891437b4b520）

> 处理办法，当然就是判断后再使用咯。

但是对于这个使用，我们发现一个很有意思的优化解释：

```
(*x).y is equivalent to x.y for a pointer-typed x, but the latter is cleaner.
from: https://go-review.googlesource.com/c/go/+/2304/
```



### 优化一：image/draw: optimize drawFillSrc.



```
benchmark                    old ns/op     new ns/op     delta
BenchmarkFillSrc             46781         46000         -1.67%
```

https://github.com/golang/go/commit/66c4031ee97289f26922126111e0510398fb43fd#diff-1cbc7137c670aec930ba891437b4b520



for 循环内的重复计算改为for 循环外一次计算，重复使用。

> 这是一个非常简单的优化，但是带来的提升还是可以的。



### 发现一



在 image/draw 中的 benchmark 是单独的一个 `bench_test.go` 文件 。

### fast path 优化 提升 97.68% ？

```
Grayscale PNG and JPEG images are not uncommon. We should have a fast path.

Also add a benchmark for the recently added CMYK fast path.

benchmark                    old ns/op     new ns/op     delta
BenchmarkGray                13960348      324152        -97.68%

from: https://github.com/golang/go/commit/c20323d2bb44146c49f71f583a1f4a089360d062#diff-1cbc7137c670aec930ba891437b4b520
```

但是我没有看懂这个 97.68% 是怎么对比测试的呢？

`drawGray`  是新加的处理逻辑，怎么得到性能对比的呢？是跟谁对比的呢？

### 优化三：image/draw: fix double-draw when the dst is paletted.

```
The second (fallback) draw is a no-op, but it's a non-trivial amount of work.

Fixes #11550.

benchmark               old ns/op     new ns/op     delta
BenchmarkPaletted-4     16301219      7309568       -55.16%
```

一个 return 就提升 55% 的性能

### 优化四：image/draw: optimize out some bounds checks.

```
We could undoubtedly squeeze even more out of these loops, and in the
long term, a better compiler would be smarter with bounds checks, but in
the short term, this small change is an easy win.

benchmark                      old ns/op     new ns/op     delta
BenchmarkFillOver-8            1619470       1323192       -18.29%
BenchmarkCopyOver-8            1129369       1062787       -5.90%
BenchmarkGlyphOver-8           420070        378608        -9.87%

On github.com/golang/freetype/truetype's BenchmarkDrawString:
benchmark                 old ns/op     new ns/op     delta
BenchmarkDrawString-8     9561435       8807019       -7.89%
```

https://github.com/golang/go/commit/e424d5968006167d5c99b3b9959e61950aa50cf8#diff-1cbc7137c670aec930ba891437b4b520



### 优化五：image/draw: optimize drawFillOver as drawFillSrc for opaque fills.



```
Benchmarks are much better for opaque fills and slightly worse on non
opaque fills. I think that on balance, this is still a win.

When the source is uniform(color.RGBA{0x11, 0x22, 0x33, 0xff}):
name        old time/op  new time/op  delta
FillOver-8   966µs ± 1%    32µs ± 1%  -96.67%  (p=0.000 n=10+10)
FillSrc-8   32.4µs ± 1%  32.2µs ± 1%     ~      (p=0.053 n=9+10)

When the source is uniform(color.RGBA{0x11, 0x22, 0x33, 0x44}):
name        old time/op  new time/op  delta
FillOver-8   962µs ± 0%  1018µs ± 0%  +5.85%   (p=0.000 n=9+10)
FillSrc-8   32.2µs ± 1%  32.1µs ± 0%    ~     (p=0.148 n=10+10)
```



取巧的一次优化，带来了一个96%的提升，另外一个5.8%的下降，但是整体来说是更优的。



### 优化五：image/draw, image/color: optimize hot path sqDiff function

https://github.com/golang/go/commit/7a8e8b2f19c423b07a86adcd41b91575b7ecd875#diff-d90c5b811de11d0d53e5d77f2162f6eb

```
Function sqDiff is called multiple times in the hot path (x, y loops) of
drawPaletted from the image/draw package; number of sqDiff calls is
between 4×width×height and 4×width×height×len(palette) for each
drawPaletted call.

Simplify this function by removing arguments comparison and relying
instead on signed to unsigned integer conversion rules and properties of
unsigned integer values operations guaranteed by the spec:

> For unsigned integer values, the operations +, -, *, and << are
> computed modulo 2n, where n is the bit width of the unsigned integer's
> type. Loosely speaking, these unsigned integer operations discard high
> bits upon overflow, and programs may rely on ``wrap around''.

image/gif package benchmark that depends on the code updated shows
throughput improvements:

name               old time/op    new time/op    delta
QuantizedEncode-4     788ms ± 2%     468ms ± 9%  -40.58%  (p=0.000 n=9+10)

name               old speed      new speed      delta
QuantizedEncode-4  1.56MB/s ± 2%  2.63MB/s ± 8%  +68.47%  (p=0.000 n=9+10)

Closes #22375.
```



没太看懂？



```
var d uint32

if x > y {
	d = x - y
} else {
	d = y - x
}
```

简化为：

```
func sqDiff(x, y int32) uint32 {
        d := uint32(x - y)
        return (d * d) >> 2
}
```

需要测试一下。。。



### 优化六：image/draw: optimize bounds checks in loops

https://github.com/golang/go/commit/5bba5053675f102a3a81242e8f7551791ae5a56e#diff-d90c5b811de11d0d53e5d77f2162f6eb

https://go-review.googlesource.com/c/go/+/136935/2

> 2018.9.24

```
name               old time/op  new time/op  delta
FillOver-8          607µs ± 1%   609µs ± 1%     ~     (p=0.447 n=9+10)
FillSrc-8          23.0µs ± 1%  22.9µs ± 2%     ~     (p=0.412 n=9+10)
CopyOver-8          647µs ± 0%   560µs ± 0%  -13.43%  (p=0.000 n=9+10)
CopySrc-8          19.3µs ± 1%  19.1µs ± 2%   -0.66%  (p=0.029 n=10+10)
NRGBAOver-8         697µs ± 1%   651µs ± 1%   -6.64%  (p=0.000 n=10+10)
NRGBASrc-8          405µs ± 1%   347µs ± 0%  -14.23%  (p=0.000 n=10+10)
YCbCr-8             432µs ± 2%   431µs ± 1%     ~     (p=0.764 n=10+9)
Gray-8              164µs ± 1%   139µs ± 1%  -15.44%  (p=0.000 n=10+10)
CMYK-8              498µs ± 0%   461µs ± 0%   -7.49%  (p=0.000 n=10+9)
GlyphOver-8         220µs ± 0%   199µs ± 0%   -9.52%  (p=0.000 n=9+10)
RGBA-8             3.81ms ± 5%  3.79ms ± 5%     ~     (p=0.549 n=9+10)
Paletted-8         1.73ms ± 0%  1.73ms ± 1%     ~     (p=0.278 n=10+9)
GenericOver-8      11.0ms ± 2%  11.0ms ± 1%     ~     (p=0.842 n=9+10)
GenericMaskOver-8  5.29ms ± 1%  5.30ms ± 0%     ~     (p=0.182 n=9+10)
GenericSrc-8       4.24ms ± 1%  4.24ms ± 0%     ~     (p=0.436 n=9+9)
GenericMaskSrc-8   7.89ms ± 1%  7.90ms ± 2%     ~     (p=0.631 n=10+10)
```

引申出来一篇文档：《边界值检查消除 Bounds Checking Elimination（BCE）- https://docs.google.com/document/d/1vdAEAjYdzjnPA9WDOQ1e4e05cYVMpqSxJYZT33Cqw2g/edit# 》

> BCE is important because: it speeds up the code, makes the binary smaller. 
>
> To see where the bound checks are introduced use -gcflags="-d=ssa/check_bce/debug=1" 

```go
s := spix[i : i+4 : i+4]
s[0]
s[1]
s[2]
s[3]
```

自测发现并没有提升，待确定，是否是 go 版本升级后带来的效果？

边界检查消除：https://gfw.go101.org/article/bounds-check-elimination.html

go 中的边界检查消除：https://www.ardanlabs.com/blog/2018/04/bounds-check-elimination-in-go.html

其他一些边界检查消除带来的性能提升：

go 标准包中一些经过 bound checks 后的 CL：
https://go-review.googlesource.com/c/go/+/136935
https://go-review.googlesource.com/c/go/+/164966

针对边界检查消除，还有一个讨论：https://golang.org/issue/27857

### 优化七：image/draw: reduce drawPaletted allocations for special source cases



```
drawPaletted has to discover R,G,B,A color values of each source image
pixel in a given rectangle. Doing that by calling image.Image.At()
method returning color.Color interface is quite taxing allocation-wise
since interface values go through heap. Introduce special cases for some
concrete source types by fetching color values using type-specific
methods.

name        old time/op    new time/op    delta
Paletted-4    7.62ms ± 4%    3.72ms ± 3%   -51.20%  (p=0.008 n=5+5)

name        old alloc/op   new alloc/op   delta
Paletted-4     480kB ± 0%       0kB ± 0%   -99.99%  (p=0.000 n=4+5)

name        old allocs/op  new allocs/op  delta
Paletted-4      120k ± 0%        0k ± 0%  -100.00%  (p=0.008 n=5+5)

Updates #15759.
```



特殊情况下的快速路径，可避免过度使用颜色。

以上基于：https://github.com/golang/go/commits/master/src/image/draw/draw.go history。

----



### Reference

1. https://blog.golang.org/go-image-package
2. https://blog.golang.org/go-imagedraw-package
   1. 中译本：https://www.cnblogs.com/ghj1976/p/3443638.html