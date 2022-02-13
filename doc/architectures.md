# アーキテクチャ
LXD はちょうど LXC と同じように Linux カーネルと Go でサポートされる
あらゆるアーキテクチャで稼働することができます。

コンテナー、コンテナーのスナップショットやイメージのように LXD のいくつか
のオブジェクトはアーキテクチャに依存しています。

このドキュメントではサポートされているアーキテクチャ、それらの
（データベースで使われる）ユニークな識別子、それらがどのように名前付け
されるべきかといくつかの注釈をリストアップします。


LXD が問題とするのはカーネルアーキテクチャであり、ツールチェインで
決定される特定のユーザースペースのフレーバーではないことに注意してください。

これは LXD は armv7 hard-float を armv7 soft-float と同じとして扱い、
両方を "armv7" として参照することを意味します。もしユーザーにとって有用で
あれば正確なユーザースペースの ABI がイメージとコンテナープロパティとして
設定可能となり、簡単に問い合わせすることを許可します。

## アーキテクチャ

ID    | Name          | Notes                           | Personalities
:---  | :---          | :----                           | :------------
1     | i686          | 32bit Intel x86                 |
2     | x86\_64       | 64bit Intel x86                 | x86
3     | armv7l        | 32bit ARMv7 little-endian       |
4     | aarch64       | 64bit ARMv8 little-endian       | armv7 (optional)
5     | ppc           | 32bit PowerPC big-endian        |
6     | ppc64         | 64bit PowerPC big-endian        | powerpc
7     | ppc64le       | 64bit PowerPC little-endian     |
8     | s390x         | 64bit ESA/390 big-endian        |
9     | mips          | 32bit MIPS                      |
10    | mips64        | 64bit MIPS                      | mips
11    | riscv32       | 32bit RISC-V little-endian      |
12    | riscv64       | 64bit RISC-V little-endian      |

上記のアーキテクチャ名は通常は Linux のカーネルアーキテクチャ名と
揃えてあります。
