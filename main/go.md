# go

## basic
- array
- slice 
- map
- struct
- chan
- ptr

---

## 接口

--- 

## 并发 
### go并发与并行
- 并发 concurrency: 关注任务切分
- 并行 parallelism: 关注同时执行
### 设计
#### 线程模型
- 实现: 混合型线程模型
- 区别: 线程与内核调度实体KSE(Kernel Scheduling Entity)之间的对应关系上
- 分类
    - 用户级线程模型
        - 用户线程与KSE是1对1关系(1:1)。大部分编程语言的线程库(如linux的pthread，Java的java.lang.Thread，C++11的std::thread等等)都是对操作系统的线程（内核级线程）的一层封装，创建出来的每个线程与一个不同的KSE静态关联，因此其调度完全由OS调度器来做。这种方式实现简单，直接借助OS提供的线程能力，并且不同用户线程之间一般也不会相互影响。但其创建，销毁以及多个线程之间的上下文切换等操作都是直接由OS层面亲自来做，在需要使用大量线程的场景下对OS的性能影响会很大。
    - 内核级线程模型
        - 用户线程与KSE是多对1关系(M:1)，这种线程的创建，销毁以及多个线程之间的协调等操作都是由用户自己实现的线程库来负责，对OS内核透明，一个进程中所有创建的线程都与同一个KSE在运行时动态关联。现在有许多语言实现的 协程 基本上都属于这种方式。这种实现方式相比内核级线程可以做的很轻量级，对系统资源的消耗会小很多，因此可以创建的数量与上下文切换所花费的代价也会小得多。但该模型有个致命的缺点，如果我们在某个用户线程上调用阻塞式系统调用(如用阻塞方式read网络IO)，那么一旦KSE因阻塞被内核调度出CPU的话，剩下的所有对应的用户线程全都会变为阻塞状态（整个进程挂起）。
        - 所以这些语言的协程库会把自己一些阻塞的操作重新封装为完全的非阻塞形式，然后在以前要阻塞的点上，主动让出自己，并通过某种方式通知或唤醒其他待执行的用户线程在该KSE上运行，从而避免了内核调度器由于KSE阻塞而做上下文切换，这样整个进程也不会被阻塞了。
    - 混合型线程模型 
        - 用户线程与KSE是多对多关系(M:N), 这种实现综合了前两种模型的优点，为一个进程中创建多个KSE，并且线程可以与不同的KSE在运行时进行动态关联，当某个KSE由于其上工作的线程的阻塞操作被内核调度出CPU时，当前与其关联的其余用户线程可以重新与其他KSE建立关联关系。当然这种动态关联机制的实现很复杂，也需要用户自己去实现，这算是它的一个缺点吧。Go语言中的并发就是使用的这种实现方式，Go为了实现该模型自己实现了一个运行时调度器来负责Go中的"线程"与KSE的动态关联。此模型有时也被称为 两级线程模型，即用户调度器实现用户线程到KSE的“调度”，内核调度器实现KSE到CPU上的调度。
#### 设计思想
- 不要以共享内存的方式来通信，相反，要通过通信来共享内存。(Do not communicate by sharing memory; instead, share memory by communicating.)
#### 设计模型
- CSP: CSP并发模型  communicating sequential processes
#### 设计实现
- G：Goroutine的简称，上面用go关键字加函数调用的代码就是创建了一个G对象，是对一个要并发执行的任务的封装，也可以称作用户态线程。属于用户级资源，对OS透明，具备轻量级，可以大量创建，上下文切换成本低等特点。
- M：Machine的简称，在linux平台上是用clone系统调用创建的，其与用linux pthread库创建出来的线程本质上是一样的，都是利用系统调用创建出来的OS线程实体。M的作用就是执行G中包装的并发任务。Go运行时系统中的调度器的主要职责就是将G公平合理的安排到多个M上去执行。其属于OS资源，可创建的数量上也受限了OS，通常情况下G的数量都多于活跃的M的。
- P：Processor的简称，逻辑处理器，主要作用是管理G对象（每个P都有一个G队列），并为G在M上的运行提供本地化资源。
#### 调度过程
- 单M
    - 单P绑定单M, 顺序执行G
- 多M
    - 多P绑定多M, 并行执行G
- 多M多P
    - 如果我们在一个Goroutine中通过go关键字创建了大量G，这些G虽然暂时会被放在同一个队列, 但如果这时还有空闲P（系统内P的数量默认等于系统cpu核心数），Go运行时系统始终能保证至少有一个（通常也只有一个）活跃的M与空闲P绑定去各种G队列去寻找可运行的G任务，该种M称为自旋的M。一般寻找顺序为：自己绑定的P的队列，全局队列，然后其他P队列。如果自己P队列找到就拿出来开始运行，否则去全局队列看看，由于全局队列需要锁保护，如果里面有很多任务，会转移一批到本地P队列中，避免每次都去竞争锁。如果全局队列还是没有，就要开始玩狠的了，直接从其他P队列偷任务了（偷一半任务回来）。这样就保证了在还有可运行的G任务的情况下，总有与CPU核心数相等的M+P组合 在执行G任务或在执行G的路上(寻找G任务)。
- 某个M在执行G的过程中被G中的系统调用阻塞
    - 在这种情况下，这个M将会被内核调度器调度出CPU并处于阻塞状态，与该M关联的其他G就没有办法继续执行了，但Go运行时系统的一个监控线程(sysmon线程)能探测到这样的M，并把与该M绑定的P剥离，寻找其他空闲或新建M接管该P，然后继续运行其中的G，大致过程如下图所示。然后等到该M从阻塞状态恢复，需要重新找一个空闲P来继续执行原来的G，如果这时系统正好没有空闲的P，就把原来的G放到全局队列当中，等待其他M+P组合发掘并执行。
- 如果某一个G在M运行时间过长，有没有办法做抢占式调度，让该M上的其他G获得一定的运行时间，以保证调度系统的公平性
    - 我们知道linux的内核调度器主要是基于时间片和优先级做调度的。对于相同优先级的线程，内核调度器会尽量保证每个线程都能获得一定的执行时间。为了防止有些线程"饿死"的情况，内核调度器会发起抢占式调度将长期运行的线程中断并让出CPU资源，让其他线程获得执行机会。当然在Go的运行时调度器中也有类似的抢占机制，但并不能保证抢占能成功，因为Go运行时系统并没有内核调度器的中断能力，它只能通过向运行时间过长的G中设置抢占flag的方法温柔的让运行的G自己主动让出M的执行权。 
    - 说到这里就不得不提一下Goroutine在运行过程中可以动态扩展自己线程栈的能力，可以从初始的2KB大小扩展到最大1G（64bit系统上），因此在每次调用函数之前需要先计算该函数调用需要的栈空间大小，然后按需扩展（超过最大值将导致运行时异常）。Go抢占式调度的机制就是利用在判断要不要扩栈的时候顺便查看以下自己的抢占flag，决定是否继续执行，还是让出自己。
    - 运行时系统的监控线程会计时并设置抢占flag到运行时间过长的G，然后G在有函数调用的时候会检查该抢占flag，如果已设置就将自己放入全局队列，这样该M上关联的其他G就有机会执行了。但如果正在执行的G是个很耗时的操作且没有任何函数调用(如只是for循环中的计算操作)，即使抢占flag已经被设置，该G还是将一直霸占着当前M直到执行完自己的任务。
- 对网络IO的优化
    - 将标准库中的网络库全部封装为非阻塞形式，防止其阻塞底层的M并导致内核调度器切换上下文带来的系统开销。
    - 运行时系统加入epoll机制(针对Linux系统)，当某一个Goroutine在进行网络IO操作时，如果网络IO未就绪，就将其该Goroutine封装一下，放入epoll的等待队列中，当前G挂起，与其关联的M可以继续运行其他G。当相应的网络IO就绪后，Go运行时系统会将等待网络IO就绪的G从epoll就绪队列中取出（主要在两个地方从epoll中获取已网络IO就绪的G列表，一是sysmon监控线程中，二是自旋的M中），再由调度器将它们像普通的G一样分配给各个M去执行。


![G-P-M](res/gpm)

---

## 单元测试

---

## 标准库
- archive	    	
    - tar	    	tar包实现了tar格式压缩文件的存取.
    - zip	    	zip包提供了zip档案文件的读写服务.
- bufio	    	bufio 包实现了带缓存的I/O操作.
- builtin	    	builtin 包为Go的预声明标识符提供了文档.
- bytes	    	bytes包实现了操作[]byte的常用函数.
- compress	    	
    - bzip2	    	bzip2包实现bzip2的解压缩.
    - flate	    	flate包实现了deflate压缩数据格式，参见RFC 1951.
    - gzip	    	gzip包实现了gzip格式压缩文件的读写，参见RFC 1952.
    - lzw	    	lzw包实现了Lempel-Ziv-Welch数据压缩格式，这是一种T. A. Welch在“A Technique for High-Performance Data Compression”一文（Computer, 17(6) (June 1984), pp 8-19）提出的一种压缩格式.
    - zlib	    	zlib包实现了对zlib格式压缩数据的读写，参见RFC 1950.
- container	    	
    - heap	    	heap包提供了对任意类型（实现了heap.Interface接口）的堆操作.
    - list	    	list包实现了双向链表.
    - ring	    	ring实现了环形链表的操作.
- context	    	Package context defines the Context type, which carries deadlines, cancelation signals, and other request-scoped values across API boundaries and between processes.
- crypto	    	crypto包搜集了常用的密码（算法）常量.
    - aes	    	aes包实现了AES加密算法，参见U.S. Federal Information Processing Standards Publication 197.
    - cipher	    	cipher包实现了多个标准的用于包装底层块加密算法的加密算法实现.
    - des	    	des包实现了DES标准和TDEA算法，参见U.S. Federal Information Processing Standards Publication 46-3.
    - dsa	    	dsa包实现FIPS 186-3定义的数字签名算法（Digital Signature Algorithm），即DSA算法.
    - ecdsa	    	ecdsa包实现了椭圆曲线数字签名算法，参见FIPS 186-3.
    - elliptic	    	elliptic包实现了几条覆盖素数有限域的标准椭圆曲线.
    - hmac	    	hmac包实现了U.S. Federal Information Processing Standards Publication 198规定的HMAC（加密哈希信息认证码）.
    - md5	    	md5包实现了MD5哈希算法，参见RFC 1321.
    - rand	    	rand包实现了用于加解密的更安全的随机数生成器.
    - rc4	    	rc4包实现了RC4加密算法，参见Bruce Schneier's Applied Cryptography.
    - rsa	    	rsa包实现了PKCS#1规定的RSA加密算法.
    - sha1	    	sha1包实现了SHA1哈希算法，参见RFC 3174.
    - sha256	    	sha256包实现了SHA224和SHA256哈希算法，参见FIPS 180-4.
    - sha512	    	sha512包实现了SHA384和SHA512哈希算法，参见FIPS 180-2.
    - subtle	    	Package subtle implements functions that are often useful in cryptographic code but require careful thought to use correctly.
    - tls	    	tls包实现了TLS 1.2，细节参见RFC 5246.
    - x509	    	x509包解析X.509编码的证书和密钥.
        - pkix	    	pkix包提供了共享的、低层次的结构体，用于ASN.1解析和X.509证书、CRL、OCSP的序列化.
- database	    	
    - sql	    	sql 包提供了通用的SQL（或类SQL）数据库接口.
        - driver	    	driver包定义了应被数据库驱动实现的接口，这些接口会被sql包使用.
- debug	    	
    - dwarf	    	Package dwarf provides access to DWARF debugging information loaded from executable files, as defined in the DWARF 2.0 Standard at http://dwarfstd.org/doc/dwarf-2.0.0.pdf
    - elf	    	Package elf implements access to ELF object files.
    - gosym	    	Package gosym implements access to the Go symbol and line number tables embedded in Go binaries generated by the gc compilers.
    - macho	    	Package macho implements access to Mach-O object files.
    - pe	    	Package pe implements access to PE (Microsoft Windows Portable Executable) files.
    - plan9obj	    	Package plan9obj implements access to Plan 9 a.out object files.
- encoding	    	encoding包定义了供其它包使用的可以将数据在字节水平和文本表示之间转换的接口.
    - ascii85	    	ascii85 包是对 ascii85 的数据编码的实现.
    - asn1	    	asn1包实现了DER编码的ASN.1数据结构的解析，参见ITU-T Rec X.690.
    - base32	    	base32包实现了RFC 4648规定的base32编码.
    - base64	    	base64实现了RFC 4648规定的base64编码.
    - binary	    	binary包实现了简单的数字与字节序列的转换以及变长值的编解码.
    - csv	    	csv读写逗号分隔值（csv）的文件.
    - gob	    	gob包管理gob流——在编码器（发送器）和解码器（接受器）之间交换的binary值.
    - hex	    	hex包实现了16进制字符表示的编解码.
    - json	    	json包实现了json对象的编解码，参见RFC 4627.
    - pem	    	pem包实现了PEM数据编码（源自保密增强邮件协议）.
    - xml	    	Package xml implements a simple XML 1.0 parser that understands XML name spaces.
- errors	    	error 包实现了用于错误处理的函数.
- expvar	    	expvar包提供了公共变量的标准接口，如服务的操作计数器.
- flag	    	flag 包实现命令行标签解析.
- fmt	    	fmt 包实现了格式化I/O函数，类似于C的 printf 和 scanf.
- go	    	
    - ast	    	Package ast declares the types used to represent syntax trees for Go packages.
    - build	    	Package build gathers information about Go packages.
    - constant	    	Package constant implements Values representing untyped Go constants and their corresponding operations.
    - doc	    	Package doc extracts source code documentation from a Go AST.
    - format	    	Package format implements standard formatting of Go source.
    - importer	    	Package importer provides access to export data importers.
    - parser	    	Package parser implements a parser for Go source files.
    - printer	    	Package printer implements printing of AST nodes.
    - scanner	    	Package scanner implements a scanner for Go source text.
    - token	    	Package token defines constants representing the lexical tokens of the Go programming language and basic operations on tokens (printing, predicates).
    - types	    	Package types declares the data types and implements the algorithms for type-checking of Go packages.
- hash	    	hash包提供hash函数的接口.
    - adler32	    	adler32包实现了Adler-32校验和算法，参见RFC 1950.
    - crc32	    	crc32包实现了32位循环冗余校验（CRC-32）的校验和算法.
    - crc64	    	Package crc64 implements the 64-bit cyclic redundancy check, or CRC-64, checksum.
    - fnv	    	fnv包实现了FNV-1和FNV-1a（非加密hash函数）.
- html	    	html包提供了用于转义和解转义HTML文本的函数.
    - template	    	template包（html/template）实现了数据驱动的模板，用于生成可对抗代码注入的安全HTML输出.
- image	    	image实现了基本的2D图片库.
    - color	    	color 包实现了基本的颜色库。
        - palette	    	palette包提供了标准的调色板.
    - draw	    	draw 包提供组装图片的方法.
    - gif	    	gif 包实现了GIF图片的解码.
    - jpeg	    	jpeg包实现了jpeg格式图像的编解码.
    - png	    	png 包实现了PNG图像的编码和解码.
- index	    	
    - suffixarray	    	suffixarrayb包通过使用内存中的后缀树实现了对数级时间消耗的子字符串搜索.
- io	    	io 包为I/O原语提供了基础的接口.
    - ioutil	    	ioutil 实现了一些I/O的工具函数。
- log	    	log包实现了简单的日志服务.
    - syslog	    	syslog包提供一个简单的系统日志服务的接口.
- math	    	math 包提供了基本常数和数学函数。
    - big	    	big 包实现了（大数的）高精度运算.
    - cmplx	    	cmplx 包为复数提供了基本的常量和数学函数.
    - rand	    	rand 包实现了伪随机数生成器.
- mime	    	mime实现了MIME的部分规定.
    - multipart	    	multipart实现了MIME的multipart解析，参见RFC 2046.
    - quotedprintable	    	Package quotedprintable implements quoted-printable encoding as specified by RFC 2045.
- net	    	net包提供了可移植的网络I/O接口，包括TCP/IP、UDP、域名解析和Unix域socket.
    - http	    	http包提供了HTTP客户端和服务端的实现.
        - cgi	    	cgi 包实现了RFC3875协议描述的CGI（公共网关接口）.
        - cookiejar	    	cookiejar包实现了保管在内存中的符合RFC 6265标准的http.CookieJar接口.
        - fcgi	    	fcgi 包实现了FastCGI协议.
        - httptest	    	httptest 包提供HTTP测试的单元工具.
        - httptrace	    	Package httptrace provides mechanisms to trace the events within HTTP client requests.
        - httputil	    	httputil包提供了HTTP公用函数，是对net/http包的更常见函数的补充.
        - pprof	    	pprof 包通过提供HTTP服务返回runtime的统计数据，这个数据是以pprof可视化工具规定的返回格式返回的.
    - mail	    	mail 包实现了解析邮件消息的功能.
    - rpc	    	rpc 包提供了一个方法来通过网络或者其他的I/O连接进入对象的外部方法.
        - jsonrpc	    	jsonrpc 包使用了rpc的包实现了一个JSON-RPC的客户端解码器和服务端的解码器.
    - smtp	    	smtp包实现了简单邮件传输协议（SMTP），参见RFC 5321.
    - textproto	    	textproto实现了对基于文本的请求/回复协议的一般性支持，包括HTTP、NNTP和SMTP.
    - url	    	url包解析URL并实现了查询的逸码，参见RFC 3986.
- os	    	os包提供了操作系统函数的不依赖平台的接口.
    - exec	    	exec包执行外部命令.
    - signal	    	signal包实现了对输入信号的访问.
    - user	    	user包允许通过名称或ID查询用户帐户.
- path	    	path实现了对斜杠分隔的路径的实用操作函数.
    - filepath	    	filepath包实现了兼容各操作系统的文件路径的实用操作函数.
- plugin	    	Package plugin implements loading and symbol resolution of Go plugins.
- reflect	    	reflect包实现了运行时反射，允许程序操作任意类型的对象.
- regexp	    	regexp包实现了正则表达式搜索.
    - syntax	    	Package syntax parses regular expressions into parse trees and compiles parse trees into programs.
- runtime	    	TODO(osc): 需更新 runtime 包含与Go的运行时系统进行交互的操作，例如用于控制Go程的函数.
    - cgo	    	cgo 包含有 cgo 工具生成的代码的运行时支持.
    - debug	    	debug 包含有程序在运行时调试其自身的功能.
    - pprof	    	pprof 包按照可视化工具 pprof 所要求的格式写出运行时分析数据.
    - race	    	race 包实现了数据竞争检测逻辑.
    - trace	    	Go execution tracer.
- sort	    	sort 包为切片及用户定义的集合的排序操作提供了原语.
- strconv	    	strconv包实现了基本数据类型和其字符串表示的相互转换.
- strings	    	strings包实现了用于操作字符的简单函数.
- sync	    	sync 包提供了互斥锁这类的基本的同步原语.
    - atomic	    	atomic 包提供了底层的原子性内存原语，这对于同步算法的实现很有用.
- syscall	    	Package syscall contains an interface to the low-level operating system primitives.
- testing	    	Package testing provides support for automated testing of Go packages.
    - iotest	    	Package iotest implements Readers and Writers useful mainly for testing.
    - quick	    	Package quick implements utility functions to help with black box testing.
- text	    	
    - scanner	    	scanner包提供对utf-8文本的token扫描服务.
    - tabwriter	    	tabwriter包实现了写入过滤器（tabwriter.Writer），可以将输入的缩进修正为正确的对齐文本.
    - template	    	template包实现了数据驱动的用于生成文本输出的模板.
        - parse	    	Package parse builds parse trees for templates as defined by text/template and html/template.
- time	    	time包提供了时间的显示和测量用的函数.
- unicode	    	unicode 包提供了一些测试Unicode码点属性的数据和函数.
    - utf16	    	utf16 包实现了对UTF-16序列的编码和解码。
    - utf8	    	utf8 包实现了支持UTF-8文本编码的函数和常量.
- unsafe	    	unsafe 包含有关于Go程序类型安全的所有操作.

---

## go翻墙
### Require
- shadowsocks: 翻墙  
- polipo: socks5协议转http/https代理  

### configuration
查看shadowsocks的http代理监听端口(1087), 并设置环境变量  
```
export http_proxy="http://127.0.0.1:1087"
export https_proxy=$http_proxy
```
用完不需要的时候记得注释掉...

### test
```
curl -sL -k -vv www.google.com
```
--- 

## related
[slice底层实现](https://blog.csdn.net/lengyuezuixue/article/details/81197691)