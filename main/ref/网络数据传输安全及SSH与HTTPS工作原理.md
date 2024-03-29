# 网络数据传输安全及SSH与HTTPS工作原理

## 一、网络数据传输安全概述

我们说的数据加密与解密通常是为了保证数据在网络传输过程中的安全性.在网络发展初期, 网络的数据安全性是没有被足够的重视的.事实上, 当时为了实现数据可以通过网络进行传输已经耗费了科学家大部分男细胞, 因此在TCP/IP协议设计的初期, 他们也实在没有太多精力去过多考虑数据在网络传输过程中可能存在的安全性问题.随着TCP/IP协议及相关技术的日渐成熟, 网络数据传输技术越来越稳定, 人们才慢慢开始重视这个问题, 美国国家标准与技术研究院(National Institue of Standard and Technology, 简称NIST)也开始制定相关的安全标准.

网络安全涉及到很多个方面, 我们这里仅仅讨论下网络数据传输过程中可能受到的威胁, 其中常见的有: 

* 数据窃听
* 数据篡改
* 身份伪装

针对以上威胁, 我们介绍下网络数据传输的安全性涉及的几个方面: 

### 1\. 机密性

机密性是指对要传输的数据进行加密和解密, 防止第三方看到通信数据的明文内容.其对应的通信过程如下: 

数据发送方: 

    plaintext(明文) ==> 转换算法 ==> ciphertext(密文)

数据接收方: 

    ciphertext(密文) ==> 转换算法 ==> plaintext(明文)

### 2\. 完整性

数据完整性是指不允许数据在传输过程中被修改(第三方恶意篡改或电平信号造成的部分数据丢失), 但是它不要求数据的机密性, 也就是说允许其他人看到明文数据.我们通常通过以不可逆的算法对数据提取特征码(也叫数据指纹), 通过验证特征码的一致性来判断数据是否被修改过, 通信过程如下: 

数据发送发: 

    plaintext(明文) ==> 转换算法 ==> plaintext(明文) + footprint(数据指纹A)

数据接收方: 

    plaintext(明文) + footprint(数据指纹A) ==> 转换算法 ==> footprint(数据指纹B) ==> 对比数据指纹A与B是否一致

### 3\. 身份验证

身份验证通常是指数据接收方需要确认发送数据给自己的数据是自己想要通信的那一方, 防止他人冒充通信对方的身份进行通信.身份验证的大体原理是: 数据发送方与数据接收方约定一种特殊的数据加解密方式, 数据发送方将一个通过约定的加密方式进行加密后的数据发送给数据接收方, 数据接收方如能按照约定的加密方式正确解密该数据就表示对数据发送方的身份验证成功.其对应的通信过程如下: 

数据发送方: 

    plaintext(明文) ==> 转换算法 ==> ciphertext(密文)

数据接收方: 

    ciphertext(密文) ==> 转换算法 ==> plaintext(明文)

## 二、数据加密算法分类

* * *

上面提到的网络数据传输所涉及到的几个方面都需要特定的转换算法来实现, 常用的转换算法(数据加密/解密算法)大体上可以分为以下几类: 

### 1\. 对称加密

对称加密是指数据加密与解密使用相同的密钥.

#### 主要功能

通常用于保证数据的机密性.

#### 常用的算法实现

* **_DES: _** Data Encryption Standard, 秘钥长度为56位, 2003年左右被破解--秘钥可以暴力破解.
* **_3DES: _** DES的改进版本.
* **_AES: _** Advanced Encryption Standard, 支持的秘钥长度包括 128bits, 192bits, 258bits, 384bits, 512bits.

> 需要说明的是, 秘钥长度越长, 数据加密与解密的时间就越久.

#### 特点

* 加密与解密使用的密钥相同.
* 在一定程度上实现了数据的机密性, 且简单、快速.
* 但是由于算法一般都是公开的, 因此机密性几乎完全依赖于密钥.
* 同一发送方与不同接收方进行通信时应使用不同的密钥, 防止数据被窃听或拦截后被解密.

### 存在的问题

* 当通信对象很多时会面临众多秘钥的有效管理问题.
* 对于一个新的数据通信对象, 密钥怎样进行传输的问题.

### 2\. 单向加密

单向加密是指只能对明文数据进行加密, 而不能解密数据.

#### 单向加密主要功能

通常用于保证数据的完整性.

#### 单向加密常用的算法实现

* **_MD5: _** 128bits
* **_SHA: _** SHA1(160bits), SHA224, SHA256, SHA384

#### 单向加密特点

* 不可逆: 无法根据数据指纹/特征码还原原来的数据.
* 输入相同, 输出必然相同.
* 雪崩效应: 输入的微小改变, 将会引起结果的巨大改变.
* 定长输出: 无论原始数据有多长, 结果的长度是相同的.

#### 单向加密存在的问题

可能出现中间人攻击, 中间人可以对原始内容进行修改之后重新生成数据指纹, 数据接收方验证数据指纹时会发现数据是正常的.此时, 数据发送方只能把生成的数据指纹进行加密后再发送给数据接收方, **_那么问题就又回到了加密密钥的传输和管理上._**

### 3\. 公钥加密(也叫非对称加密)

公钥加密, 也被称作非对称加密, 也就是说加密和解密所使用的密钥是不同的.

#### 主要作用

通常用于保证身份验证.

#### 常用的公钥加密算法有

* **_RSA: _** 可以实现数字签名 和 数据加密
* **_DSA: _** 只能实现数字签名, 不能实现数据加密

#### 公钥加密特点

* 加密与解密使用的不同的密钥.
* 实际上它所使用的密钥是一对儿, 一个交公钥, 一个叫私钥.这对密钥不是独立的, 公钥是从私钥中提炼出来, 因此私钥是很长的, 968位、1024位、2048位、4096位的都有.
* 通常公钥是公开的, 所有人都可以得到; 私钥是不能公开的, 只有自己才有.
* 用公钥机密的内容只能用与之对应的私钥才能解密, 反之亦然, 这个特点尤为重要.

我们发现公钥加密“貌似”已经解决了密钥管理的问题--所有人只需要知道自己的那一对儿密钥即可, 需要跟谁通信就去获取对方的公钥, 然后通过这个公钥对数据进行加密和机密就可以了.我们可以用它来完成以下两件事情: 

* **_用自己的私钥加密, _** 可以保证身份验证, 因为用你的私钥加密的数据只能用你的公开的公钥才能解密数据; 但是不能保证数据的机密性, 因为所有人都知道你的公钥.浏览器检查CA证书合法性时, 验证CA机构的数字签名时就是通过这种方式进行的.
* **_用对方的公钥加密, _** 可以保证数据的机密性, 因为只有用对方的私钥才能解密, 而对方的私钥只有他一个人有.HTTPS通信时, 通过密钥协商技术得到的密钥进行传输时就是通过这种方式来保证机密性的.其实用对方公钥加密也可以用于用于身份验证, 验证过程是: A用B的公钥加密数据后将密文传输给B, B用自己的私钥进行解密并将明文发送回给A, A对比B返回的明文和自己加密前的明文一致则表示对B完成了身份验证, 通过SSH进行免密钥登录时就是通过这种方式来完成用户身份验证的.

> 事实上, **_公钥加密算法很少用于数据加密, 它通常只是用来做身份认证_**, 因为它的密钥太长, 加密速度太慢--公钥加密算法的速度甚至比对称加密算法的速度慢上3个数量级(1000倍).

#### 公钥加密存在的问题

* 既然公钥加密通常只用于身份验证, 而不是用于保证数据的机密性, 也就意味着这个密钥对儿并不能完全作为加密和解密数据的秘钥来用.那么, 秘钥的管理和传输问题依然存在着, 这个问题到底怎样来解决呢？
* 另外还有个问题就是, 如果有人伪造了一对儿密钥, 把其中的公钥发送给别人怎么办？怎样验证以获取公钥的合法性呢？

#### 密钥管理的解决方案

实际上, 已经存在一种专门用于秘钥交换的算法--Diffie-Hellman加密算法.该加密算法本身仅限于秘钥的交换用途, 被许多商用产品用作秘钥交换技术.这种秘钥交换技术的目的在于使得两个用户安全的交换一个密钥, 以便用于之后的数据对称加密.也就是说, 通信双方可以通过这个技术, 动态的协商生成一个用于对称加密的密钥, 而不用管理很多静态的密钥, 这样就解决了密钥的管理问题.

> 需要说明的是, 在通过秘钥交互技术动态协商生成密钥之前, 通常需要先通过公钥加密算法对对方的身份进行验证.实际上, https就是这样工作的.

#### 防止公钥被伪造的解决方案

公钥实际上也是一段文本, 验证公钥的合法性涉及到两个方面: 

* 1)该公钥的发布者身份是否合法
* 2)该公钥的内容是否被篡改过

其实, 这个已经不是靠纯技术能解决的问题了, 这需要借助一些机构和人为约定来解决.常见的解决方案有两种: 

* 1)**_公钥的合法拥有者, 通过官方渠道声明其密钥的数据指纹: _** 既然时官方发布的信息, 那么身份的合法性是有保证的; 用户在获取公钥后也生成一个数据指纹, 通过对比这两个数据指纹就知道公钥内容是否被修改过; SSH的身份验证实际上就是这个原理.
* 2)**_通过一些权威的机构来完成这些验证: _** 比如https使用的证书就是由CA机构签发的, 这个在后面讲https原理时再做具体介绍.

我们常见的对于上面这些加密算法的经典应用就是ssh和https了, 它们都是使用这些加密算法实现的网络协议.下面我们对ssh和https的工作原理进行下介绍, 一方面当做上面这些加密算法的实例讲解, 帮助大家了解这些算法的经典应用; 另一方面, 也帮助大家更深入的理解ssh和https是什么, 以及它们是怎样工作的.

## 三、SSH工作原理

### 1\. SSH是什么

简单来说, **_SSH就是一种网络协议, 主要用于计算机之间的加密登录与数据传输_**, 使用方式如下: 

    # ssh user@host

表示要以user这个用户的身份登录host这台网络机器.也可以省略前面的user, 这样来用`ssh host`, 表示以当前本地登录的用户名登录host这台网络机器.

早期, 人们主要是通过telnet协议进行计算机之间的登录操作, 但是它有一个很严重的安全隐患就是“数据是明文传输的”, 登录时传输的包括用户名和密码在内的所有信息都有可能会被恶意拦截而暴露.而SSH则是将登录信息全部加密后进行传输的, 因此使用SSH进行登录时安全的, 即使数据在传输过程中被截获, 里面的密码已经被加密而不会泄露.

> 现在SSH作为互联网安全的一个基本解决方案, 已经在全世界获得推广, 且目前已经成为Linux系统的标准配置.需要说明的是, SSH只是一种协议, 它有多种软件实现, 既有商业的, 也有开源的.OpenSSH是当前使用最为广泛的一个SSH协议的开源实现.

### 2\. SSH工作原理

其实SSH是充分利用了公钥加密/非对称机密 、对称加密 和 单向加密 来实现数据安全登录的.在使用SSH进行通信时, 通信过程分为以下几个步骤: 

* 1)**_生成会话密钥: _** 这个会话密钥, 不是密钥对儿中公钥或私钥, 而是通过密钥协商技术生成密钥.这个密钥会被通过被登录机器的密钥对进行加密后传输, 用于后续所有(通过对称加密方式进行的)加密通信.
* 2)**_用户身份认证(登录): _**\* 这个对登录者进行身份验证的过程是通过登录者的密钥对儿对数据进行加解密验证实现的, 这个过程中传输的所有数据都是通过上一步生成的密钥加密过的.
* 3)**_数据加密通信: _** 后面就行基于第1步生成的密钥进行数据加密传输的通信过程了.

下面来看具体解析.

#### 账号密码安全登录的实现

上面提到, SSH是通过对数据进行加密后进行传输来保证数据安全的.但是, SSH的数据加密采用的是对称加密算法, 只是对称加密所使用的密钥是通过公钥加密/非对称加密实现加密后的安全传输的.另外, 每台Linux机器都有自己的密钥对儿(通常放在/etc/ssh目录下), 这个密钥对儿跟具体的用户无关.其工作流程是: 

* 1)在主机A上向主机B发送连接请求; 
* 2)主机B在与用户建立连接后, 把自己的公钥发送给主机A; 
* 3)主机A通过密钥协商技术产生一个随机密钥, 然后使用主机B的公钥对这个随机密钥进行加密后发送给主机B;
* 4) 主机B接收到主机A发送过来的密文形式的密钥后, 通过自己的私钥进行解密, 得到对称加密使用的密钥明文; 至此, 会话密钥已经生成完毕了; 
* 5)主机A通过生成的会话密钥对账号和密码等信息进行加密然后发送给主机B; 
* 6)主机B接收到加密信息后, 使用会话密钥进行解密, 从而得到明文的账号和密码进行账号验证; 
* 7)主机B在验证账号和密码后通知主机A是否登录成功; 

这样即便有人结果了账号密码信息, 也是密文信息, 并不能知道里面是什么内容.貌似已经OK了, 但是, 主机A怎么验证主机B的身份呢？如果有主机C冒充主机B截获了登录请求, 将自己伪造的公钥发送给主机A, 怎么办？尽管信息是加密过的, 通信过程也是合法的, 但是通信信息都被主机C截获了, 其实这就是所谓的“中间人攻击”(Man-in-the-middle attack).其实, **_对主机B进行验证就是对主机B发送过来的公钥的合法性进行验证的过程._**

#### 公钥合法性验证的实现

上面我们提到过, 验证公钥的合法性有两种方式: 

* 1)验证公钥的官方发布的公钥数据指纹
* 2)通过权威的结构进行验证

SSH主要用于机器之间的安全登录, 因此通常不会通过权威的机构去签发证书, 它主要是通过验证数据指纹的方式来验证公钥的合法性的.公钥的合法性验证是发生在主机A接收到主机B发送的公钥之后, 主机A向主机B协商产生会话密钥之前, 也就是上个部分所列举的数据机密时间的第2个步骤和第3个步骤之间.具体的工作流程如下: 

* 1)上个部分所列举的数据加密实现的第1-2步; 
* 2)主机A会去当前用户家目录下的.ssh/known_hosts文件中查找是否存在该机器的公钥, 如果不存在, 表示主机A是第一次与该主机进行通信, 那么主机A会计算出该公钥的数据指纹并要求用户对该指纹进行合法性确认.就是我们经常看到的的这样子:   
![img](https://images2015.cnblogs.com/blog/1063221/201706/1063221-20170610110552778-2137595593.png)
* 3)用户需要把目标主机管理员公布的公钥的数据指纹与主机A计算得到的数据指纹进行比对, 如果一致, 则说明该公钥是合法的; 如果不一致则说明不合法; 
* 4)用户如果确认该公钥是合法的, 则输入yes表示继续后面的连接, 主机A则会把这个公钥的内容保存到当前用户家目录下的.ssh/known_hosts文件中, 然后提示用户输入密码, 如下图所示:   
![img](https://images2015.cnblogs.com/blog/1063221/201706/1063221-20170610110717965-1712933521.png)下次再登录, 执行到步骤2时, 主机A发现该公钥已经在.ssh/known_hosts文件中存在了, 就不用要求再次确认了, 而是会直接提示输出密码:   
![img](https://images2015.cnblogs.com/blog/1063221/201706/1063221-20170610110310809-1975093095.png)
* 5)至此, 主机B的身份合法性验证就结束了.

> 每个用户都有自己的kown\_hosts文件, 它们是相互独立的.我们也可以为所有用户保存一份公共的可信赖的远程主机的公钥, 这个文件通常是/etc/ssh/ssh\_known_hosts.

#### 问题: 假如之前我们通过ssh登录过的一台机器的IP被绑定到其他机器上了会出现什么情况

当机器A接收到机器B的公钥指纹时, 发现knowns_hosts文件中虽然有机器B的公钥, 但是计算得出的公钥指纹与机器B发送过来的公钥指纹不一致.这肯定是不一致的, 因为每台机器的密钥对都是随机生成的, 几乎不可能出现重复.因此, 我们会看到如下提示信息: 

![img](https://images2015.cnblogs.com/blog/1063221/201706/1063221-20170615172106587-1711264707.png)

上面的大概意思是, 主机A发现主机B的公钥指纹对不上了, 怀疑我们正在遭受中间人攻击(即有人在冒充主机B), 并且密码验证方式和键盘交互验证方式都被禁止使用了.其实, 我们自己知道是因为IP被绑定到其他机器上引起的这个问题, 所以我们如果想继续登录新的主机B, **_只需要在.ssh/known_hosts文件中把原来保存的主机B的公钥删掉就可以了_**.

### 3\. SSH免密钥登录的实现

#### 使用SSH免密钥登录的优点

大家都知道, SSH免密钥登录是通过公钥认证的, 用户登录时只需要提供用户名, 而不需要输入密码.其实其优点不止这一个, 我们来总结下: 

* 1)使用账号和密码进行登录时, 由于用户无法设置空密码, 因此每次登录都要输入密码.而且即使系统允许给用户设置空密码, 也是十分危险的行为.而公钥认证允许用户给私钥设置空密码, 同时还能保证安全性.
* 2)使用账号和密码进行登录时密码容易被人看到, 且密码也容易被猜到; 而公钥认证所使用的密钥不用手动输入, 而且内容很长, 因此安全性比较高.
* 3)使用账号和密码进行登录时, 服务器上的一个账号如果想给多个人同时使用, 机器密码维护工作会变得很繁琐, 因为他们所有人都需要知道密码是什么, 当修改密码也要通知他们每个人.而使用公钥认证只需要把它们的公钥保存在服务器上, 如果要取消某个人的操作权限, 只需要把这个人的公钥删掉, 而不需要修改服务器密码.

#### SSH免密钥登录过程

其实**_登录的过程就是被登录端对登录用户进行“身份验证”的过程_**, 前面是通过账号和密码来验证用户身份, 因为密码应该只有该账号的拥有者才知道.而我们知道公钥加密算法中, 用公钥加密的数据只能由与其配对的私钥才能解密, 而私钥只有用户自己才有.那么, 我们是否可以通过这种方式来验证用户身份呢？实际上SSH免密钥登录就是这样的原理.比如, 我们想在主机A上以root用户以SSH免密钥的方式登录主机B, 登录验证过程是这样的: 

* 1)主机A与主机B协商产生会话密钥; 
* 2)主机A会向主机B发送一个登录请求(如: `root@192.168.1.2`), 发送的信息包括用户名root和root的公钥指纹, 且所有信息都是通过会话密钥加密过的.
* 3)主机B通过会话密钥解密主机A发送的数据得到请求登录的用户名root和root的公钥指纹, 然后读取root用户家目录下的所有公钥数据(/root/.ssh/autorized_keys文件中), 并分别通过单向加密算法获取各公钥的数据指纹与主机A发送过来的数据指纹做对比, 从而找到主机A上的root用户的公钥; 
* 4)主机B使用找到的root用户的公钥对一个随机数进行加密发送发送给主机A; 
* 5)主机A使用root用户的私钥对主机B发送的随机数密文进行解密, 然后把解密结果发送给主机B;
* 6)主机B验证主机A解密后的数据与自己发送的数据一致, 则对root用户的身份验证成功; 

**_那么主机A是怎样获取root用户的私钥的呢？主机B又是怎样获取root用户的公钥的呢？_** 这个就是实现SSH免密钥登录所要配置的内容: 

* 1)生成密钥对儿: 在当前机器A上, 可以通过`ssh-keygen`命令生成一个ssh密钥对儿, 一路回车就可以; 生成的密钥对儿默认保存在当前登录用户家目录下的.ssh目录, 也可以指定保存目录.我们当前是以root用户登录, 因此是保存在/root/.ssh目录:   
![img](https://images2015.cnblogs.com/blog/1063221/201706/1063221-20170610122423965-60669097.png)
* 2)我们可以把这个密钥对儿中的两个文件复制到其他用户家目录的.ssh目录下(如/home/wader/.ssh/目录), 也可以复制到其他任意目录.需要说明的是一定要注意目录和文件的权限: .ssh 目录的权限必须是0700, authorized_keys 文件权限必须是0600.
* 3)当在主机A上通过 `ssh root@hostB`进行登录时, 主机A会尝试读取登录用户的家目录下的私钥文件(这里是以root用户登录主机B, 因此主机A会读取/root/.ssh/id_rsa文件作为私钥), 也可以通过-i选项指定要使用的私钥文件; 
* 4)我们需要手动把公钥的内容复制到要登录机器B的相应用户(如root)家目录下的指定文件中: /home/root/.ssh/autorized_keys; 可以使用`ssh-copy-id root@hostB`命令直接完成这个操作, 也可以通过复制粘贴的方式来完成; 
* 5)在当前机器上就可以通过ssh私钥使用root用户登录机器B了.

### 4\. ssh免密钥登录在git中的使用

我们在管理git仓库中的项目时, 可以使用http/https协议, 也可以使用ssh协议来管理我们的项目代码: 

http/https协议: 

    http://192.168.1.1/GROUP_OR_USER/PROJECT_NAME.git
    https://192.168.1.1/GROUP_OR_USER/PROJECT_NAME.git

ssh协议:

    ssh://git@192.168.1.1/GROUP_OR_USER/PROJECT_NAME.git

无论使用http/https协议还是ssh协议来管理项目仓库, 对于非公开的仓库都是需要进行登录(即账户身份验证)的.如果我们使用http/https协议的话, 就需要提供用户名和密码进行验证; 如果我们使用ssh协议的话, 就可以把我们公钥保存到项目仓库机器的指定位置, 来通过非对称加密的方式进行身份验证, 验证的原理上面已经详细说明过了.

## 四、HTTPS工作原理

HTTPS实际上就是HTTP协议和SSL/TSL协议的组合, 可以把HTTPS大致理解为“HTTP over SSL”或“HTTP over TSL”.关于它们的相关介绍, 可以参考[这篇文章](http://www.techug.com/post/https-ssl-tls.html).对于HTTPS我们应该有以下几个认知: 

* 1)使用HTTPS传输数据是安全的, 因为数据都是被加密传输的; 
* 2)使用HTTPS需要在服务器端配置密钥对; 
* 3)使用HTTPS需要花钱找专业的权威机构进行CA证书的签发.

那么使用HTTPS与网站服务器进行交互的流程和原理到底是怎样的呢？让我们先以逆向思考的方式来进行说明: 

* 我们说过, 公钥加密/非对称加密方式虽然安全, 但是由于密钥过长, 加密和解密速度都远远低于对称加密.因此, 出于对性能方面的考虑, HTTPS并不是把所有传输的数据都使用公钥加密的方式进行机密性的保护, 而是继续使用对称加密的方式来加密数据.还有一个原因就是, 使用公钥机密算法来保证数据机密性的话, 需要通信双方都要有密钥对儿, 否则总有一方发出的数据是能被对方公布的公钥解密的.

* 既然时使用对称加密的方式加密数据, 就需要有一个通信双方都知道的加解密所使用的密钥.HTTPS是通过上面提到的密钥交换技术来动态协商这个密钥的, 实际上就是由客户端生成一个随机密钥, 然后发送给服务器端, 这样就解决了密钥的管理问题.
  
* 既然说HTTPS是安全的, 那么客户端生成的这个随机密钥肯定不能以明文的方式发送给服务器端啊.是的, 当客户端以https的方式访问一个站点时, 该站点会自动下发其公钥信息.客户端会使用这个公钥对产生的随机密钥进行加密, 然后传送给服务器端.服务器端以自己的私钥对这个密文进行解密, 然后得到这个密钥的明文内容.至此, 客户端与服务端用于对称加密和解密的密钥协商与传输工作已经安全的完成了.
  
* 那么要通过网络获取服务器端的公钥信息, 那么怎么验证该公钥信息的合法性呢？我们上面说过, 不是所有问题都能依赖技术来解决的.这里要验证公钥信息的合法性就要依靠CA证书签发机构了, 网站服务的提供者必须找一个大家都信任的机构来对他提供的公钥进行签名, 用户得到一个网站下发的公钥后看到有这个机构的签名就认为这个公钥是合法的, 是可信赖的.
  
* 那么CA机构的签名要以什么样的形式来提供呢？实际上网站服务器下发给客户端(通常是浏览器)的公钥已经不仅仅是密钥对儿中公钥的内容了, 而是包含了证书签发机构写入的其他信息的CA证书.这个CA证书中包括证书签发机构的标识和公钥的数据指纹, 当然还有包含网站服务提供者的公钥信息以及证书到期时间等等.但是, 我们前面提到过, 单向加密只能保证数据的完整性, 不能保证数据机密性.CA证书的伪造者完全可以伪造公钥信息并生成相应的数据指纹, 然后发送给用户.那么现在的问题就变成了要验证CA证书中公钥的合法性以及CA证书提供者的身份了.貌似问题只是转移了, 而没有被解决.
  
* 其实每个CA证书的签发机构也都有自己的密钥对儿, 他们放在CA证书中的公钥的数据指纹时通过自己的私钥加密过的, 而这些CA证书签发机构的公钥是被各浏览器厂商内置在浏览器内部的.当浏览器接收到某网站服务器下发的CA证书后会根据CA证书中签发机构的标识来读取浏览器内置的相应CA签发机构的公钥信息, 通过这个公钥信息对公钥数据指纹的密文进行解密就可以得到CA证书中包含的公钥信息的真实数据指纹.浏览器再通过单向加密的方式自己计算一次CA证书中包含的公钥信息的数据指纹, 两个数据指纹一致则说明这个CA证书确实是该CA机构签发的, 同时也证明了CA证书中的公钥信息没有被篡改过.至此, 所有的问题就都解决了.
  
现在我们再来以正常的顺序描述一下使用HTTPS与网站服务器进行交互的过程: 

* 1)浏览器A与网站服务器B通过三次握手后建立网络连接.
* 2)浏览器A告诉网站服务器B: 我想跟你通过HTTPS协议进行秘密交流.
* 3)网站服务器B把包含自己公钥信息的CA证书下发给浏览器A, 并告诉浏览器A这个CA证书里有我的公钥信息, 你决定一个对称加密使用的秘钥串, 然后通过这个公钥加密后发送给我.
* 4) 浏览器A接收到网站服务器B下发的CA证书后, 对这个CA证书的及其包含的公钥信息的合法性表示怀疑.于是根据CA证书中包含的证书签发机构的标识找到自身内置的该签发机构的公钥对CA证书中公钥的数据指纹进行解密, 然后再自己计算一下CA证书中公钥的数据指纹, 对了一下这两个数据指纹是一致的.浏览器A放心了, 知道这个CA证书是合法的, CA证书中的公钥也没有被篡改过.
* 5)然后浏览器A通过通过密钥协商技术产生了一个随机的字符串作为与网站服务器B进行秘密通信的密钥, 并把这个密钥通过CA证书中包含的公钥进行加密后发送给网站服务器B.
* 6)网站服务器B接收到密文格式的密钥后, 通过自己的私钥进行解密得到密钥的明文内容.
* 7)浏览器A和网站服务器B开始了秘密交流.
