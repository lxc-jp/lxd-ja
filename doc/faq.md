# FAQ
<!-- # Frequently Asked Questions -->

## コンテナー起動時の問題 <!-- Container Startup Issues -->

もしコンテナーが起動しない場合や、期待通りの動きをしない場合に最初にすべきことは、コンテナーが生成したコンソールログを見ることです。
これには `lxc console --show-log CONTAINERNAME` コマンドを使います。
<!--
If your container is not starting, or not behaving as you would expect,
the first thing to do is to look at the console logs generated by the
container, using the `lxc console --show-log CONTAINERNAME` command.
-->

次の例では、`systemd` が起動しない RHEL 7 システムを調べています。
<!--
In this example, we will investigate a RHEL 7 system in which `systemd`
can not start.
-->

    # lxc console --show-log systemd
    Console log:
    
    Failed to insert module 'autofs4'
    Failed to insert module 'unix'
    Failed to mount sysfs at /sys: Operation not permitted
    Failed to mount proc at /proc: Operation not permitted
    [!!!!!!] Failed to mount API filesystems, freezing.

ここでのエラーは、/sys と /proc がマウントできないというエラーです。これは非特権コンテナーでは正しい動きです。
しかし、LXD は _可能であれば_ 自動的にこれらのファイルシステムをマウントします。
<!--
The errors here say that /sys and /proc can not be mounted - which is
correct in an unprivileged container.  However, LXD does mount these
filesystems automatically _if it can_. 
-->

[コンテナーの要件](container-environment.md) では、コンテナーには `/sbin/init` が存在するだけでなく、空の `/dev`、`/proc`、`/sys` フォルダーが存在していなければならないと定められています。
もしこれらのフォルダーが存在しなければ、LXD はこれらをマウントできません。そして、systemd がこれらをマウントしようとします。
非特権コンテナーでは、systemd はこれを行う権限はなく、フリーズしてしまいます。
<!--
The [container requirements](container-environment.md) specify that
every container must come with an empty `/dev`, `/proc`, and `/sys`
folder, as well as `/sbin/init` existing.  If those folders don't
exist, LXD will be unable to mount to them, and systemd will then
try to. As this is an unprivileged container, systemd does not have
the ability to do this, and it then freezes.
-->

何かが変更される前に環境を見ることはできます。`raw.lxc` 設定パラメーターを使って、明示的にコンテナー内の init を変更できます。
これは Linux カーネルコマンドラインに `init=/bin/bash` を設定するのと同じです。
<!--
So you can see the environment before anything is changed, you can
explicitly change the init in a container using the `raw.lxc` config
param.  This is equivalent to setting `init=/bin/bash` on the linux
kernel commandline.
-->

    lxc config set systemd raw.lxc 'lxc.init.cmd = /bin/bash'

次のようになります:
<!--
Here is what it looks like:
-->

    root@lxc-01:~# lxc config set systemd raw.lxc 'lxc.init.cmd = /bin/bash'
    root@lxc-01:~# lxc start systemd
    root@lxc-01:~# lxc console --show-log systemd
    
    Console log:

    [root@systemd /]#
    root@lxc-01:~#

コンテナーが起動しましたので、コンテナー内で期待通りに動いていないことを確認できます。
<!--
Now that the container has started, you can look in it and see that things are
not running as well as expected.
-->

    root@lxc-01:~# lxc exec systemd bash
    [root@systemd ~]# ls
    [root@systemd ~]# mount
    mount: failed to read mtab: No such file or directory
    [root@systemd ~]# cd /
    [root@systemd /]# ls /proc/
    sys
    [root@systemd /]# exit

LXD は自動修復を試みますので、起動時に作成されたフォルダもあります。コンテナーをシャットダウンして再起動すると問題は解決されます。
しかし問題の根源は依然として存在しています。**テンプレートに必要なファイルが含まれていないという問題です**。
<!--
Because LXD tries to auto-heal, it *did* create some of the folders when it was
starting up. Shutting down and restarting the container will fix the problem, but
the original cause is still there - the **template does not contain the required
files**.
-->

## ネットワークの問題 <!-- Networking Issues -->

大規模な[プロダクション環境](production-setup.md)では、複数の VLAN を持ち、LXD クライアントを直接それらの VLAN に接続するのが一般的です。
netplan と systemd-networkd を使っている場合、いくつかの最悪の問題を引き起こす可能性があるバグに遭遇するでしょう。
<!--
In a larger [Production Environment](production-setup.md), it is common to have
multiple vlans and have LXD clients attached directly to those vlans. Be aware that
if you are using netplan and system-networkd, you will encounter some bugs that
could cause catastrophic issues
-->

### VLAN ベースのブリッジでは netplan で systemd-networkd が使えない <!-- Do not use system-networkd with netplan and bridges based on vlans -->

執筆時点（2019-03-05）では、netplan は VLAN にアタッチされたブリッジにランダムな MAC アドレスを割り当てられません。
常に同じ MAC アドレスを選択するため、同じネットワークセグメントに複数のマシンが存在する場合、レイヤー 2 の問題が発生します。
複数のブリッジを作成することも困難です。代わりに `network-manager` を使ってください。
設定例は次のようになります。管理アドレスが 10.61.0.25 で、VLAN102 をクライアントのトラフィックに使います。
<!--
At time of writing (2019-03-05), netplan can not assign a random MAC address to
a bridge attached to a vlan. It always picks the same MAC address, which causes
layer2 issues when you have more than one machine on the same network segment.
It also has difficultly creating multiple bridges.  Make sure you use
`network-manager` instead. An example config is below, with a management
address of 10.61.0.25, and VLAN102 being used for client traffic.
-->

    network:
      version: 2
      renderer: NetworkManager
      ethernets:
        eth0:
          dhcp4: no
          accept-ra: no
          # This is the 'Management Address'
          addresses: [ 10.61.0.25/24 ]
          gateway4: 10.61.0.1
          nameservers:
            addresses: [ 1.1.1.1, 8.8.8.8 ]
        eth1:
          dhcp4: no
          accept-ra: no
          # A bogus IP address is required to ensure the link state is up
          addresses: [ 10.254.254.25/32 ]
    
      vlans:
        vlan102:
          accept-ra: no
          dhcp4: no
          id: 102
          link: eth1

      bridges:
        br102:
          accept-ra: no
          dhcp4: no
          interfaces: [ "vlan102" ]
          # A bogus IP address is required to ensure the link state is up
          addresses: [ 10.254.102.25/32 ]
          parameters:
            stp: false

#### 注意事項 <!-- Things to note -->

* eth0 はデフォルトゲートウェイの指定がある管理インターフェースです <!-- eth0 is the Management interface, with the default gateway. -->
* vlan102 は eth1 を使います <!-- vlan102 uses eth1. -->
* br102 は vlan102 を使います。そして __bogus な /32 の IP アドレスが割り当てられています__ <!-- br102 uses vlan102, and _has a bogus /32 IP address assigned to it_ -->

他に重要なこととして、`stp: false` を設定することがあります。そうしなければ、ブリッジは最大で 10 秒間 `learning` 状態となります。これはほとんどの DHCP リクエストが投げられる期間よりも長いです。
クロスコネクトされてループを引き起こす可能性はありませんので、このように設定しても安全です。
<!--
The other important thing is to set `stp: false`, otherwise the bridge will sit
in `learning` state for up to 10 seconds, which is longer than most DHCP requests
last. As there is no possibility of cross-connecting and causing loops, this is
safe to do.
-->

### 'port security' に気をつける <!-- Beware of 'port security' -->

スイッチは MAC アドレスの変更を許さず、不正な MAC アドレスのトラフィックをドロップするか、ポートを完全に無効にするものが多いです。
ホストから LXD インスタンスに ping できたとしても、_異なった_ ホストから ping できない場合は、これが原因の可能性があります。
この原因を突き止める方法は、アップリンク（この場合は eth1）で tcpdump を実行することです。
すると、応答は送るが ACK を取得できない 'ARP Who has xx.xx.xx.xx tell yy.yy.yy.yy'、もしくは ICMP パケットが行き来しているものの、決して他のホストで受け取られないのが見えるでしょう。
<!--
Many switches do *not* allow MAC address changes, and will either drop traffic
with an incorrect MAC, or, disable the port totally. If you can ping a LXD instance
from the host, but are not able to ping it from a _different_ host, this could be
the cause.  The way to diagnose this is to run a tcpdump on the uplink (in this case,
eth1), and you will see either 'ARP Who has xx.xx.xx.xx tell yy.yy.yy.yy', with you
sending responses but them not getting acknowledged, or, ICMP packets going in and
out successfully, but never being received by the other host.
-->

### 不必要に特権コンテナーを実行しない <!-- Do not run privileged containers unless necessary -->

特権コンテナーはホスト全体に影響する処理を行うことができます。例えば、ネットワークカードをリセットするために、/sys 内のものを使えます。
これは **ホスト全体** に対してリセットを行い、ネットワークの切断を引き起こします。
ほぼすべてのことが非特権コンテナーで実行できます。コンテナー内から NFS マウントしたいというような、通常とは異なる特権が必要なケースでは、バインドマウントを使う必要があるかもしれません。
<!--
A privileged container can do things that effect the entire host - for example, it
can use things in /sys to reset the network card, which will reset it for **the entire
host**, causing network blips. Almost everything can be run in an unprivileged container,
or - in cases of things that require unusual privileges, like wanting to mount NFS
filesystems inside the container, you may need to use bind mounts.
-->

