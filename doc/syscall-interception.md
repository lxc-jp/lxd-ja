# システムコールのインターセプション
<!-- System call interception -->
<!--
LXD supports intercepting some specific system calls from unprivileged
containers and if they're considered to be safe, will executed with
elevated privileges on the host.
-->
LXD では非特権コンテナで、いくつか特定のシステムコールをインターセプトできます。もし、それが安全であると見なせるのであれば、ホスト上で特権を昇格させて実行します。

<!--
Doing so comes with a performance impact for the syscall in question and
will cause some work for LXD to evaluate the request and if allowed,
process it with elevated privileges.
-->
これを行うことで、対象のシステムコールではパフォーマンスに影響があり、LXD ではリクエストを評価するための作業が必要となり、もし許可されれば昇格した特権で実行されます。

# 利用できるシステムコール
<!-- Available system calls -->
## mknod / mknodat
<!--
The `mknod` and `mknodat` system calls can be used to create a variety of special files.
-->
`mknod` と `mknodat` システムコールを使用して、色々なスペシャルファイルを作成できます。

<!--
Most commonly inside containers, they may be called to create block or character devices.
Creating such devices isn't allowed in unprivileged containers as this
is a very easy way to escalate privileges by allowing direct write
access to resources like disks or memory.
-->
もっとも一般的にはコンテナ内部で、ブロックデバイスやキャラクターデバイスを作成するために呼び出されます。このようなデバイスを作成することは、非特権コンテナ内では許可されません。これは、ディスクやメモリのようなリソースに直接書き込みのアクセスを許可することになり、特権を昇格するのに非常に簡単な方法であるためです。　

<!--
But there are files which are safe to create. For those, intercepting
this syscall may unblock some specific workloads and allow them to run
inside an unprivileged containers.
-->
しかし、作成しても安全であるファイルもあります。このような場合に、システムコールをインターセプトすることで、特定の処理のブロックが解除され、非特権コンテナ内部で実行できるようになります。

<!--
The devices which are currently allowed are:
-->
現時点で許可されているデバイスは次のものです:

 - overlayfs whiteout (char 0:0)
 - /dev/console (char 5:1)
 - /dev/full (char 1:7)
 - /dev/null (char 1:3)
 - /dev/random (char 1:8)
 - /dev/tty (char 5:0)
 - /dev/urandom (char 1:9)
 - /dev/zero (char 1:5)

<!--
All file types other than character devices are currently sent to the
kernel as usual, so enabling this feature doesn't change their behavior
at all.
-->
キャラクターデバイス以外のすべてのファイルタイプは、現時点では通常通りカーネルに送られるので、この機能を有効にしても動作は全く変わりません。

<!--
This can be enabled by setting `security.syscalls.intercept.mknod` to `true`.
-->
この機能は `security.syscalls.intercept.mknod` を `true` に設定することで有効に出来ます。

## setxattr
<!--
The `setxattr` system call is used to set extended attributes on files.
-->
`setxattr` システムコールは、拡張ファイル属性を設定するのに使われます。

<!--
The attributes which are handled by this currently are:
-->
現時点で、これにより処理される属性は次のものです:

 - trusted.overlay.opaque (overlayfs directory whiteout)

<!--
Note that because the mediation must happen on a number of character
strings, there is no easy way at present to only intercept the few
attributes we care about. As we only allow the attributes above, this
may result in breakage for other attributes that would have been
previously allowed by the kernel.
-->
この介入は多数の文字列で行う必要があるため、現在のところ、対象の少数の属性のみインターセプトする簡単な方法がありません。上記の属性のみを許可しているため、カーネルが以前に許可していた他の属性を破損する可能性があります。

<!--
This can be enabled by setting `security.syscalls.intercept.setxattr` to `true`.
-->
この機能は `security.syscalls.intercept.setxattr` を `true` に設定することで有効にできます。
