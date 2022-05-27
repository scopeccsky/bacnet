# Bacnet

一个基于go的bacnet/ip库，无其他依赖
在https://github.com/REQUEA/bacnet 库基础上修改而来， 可以参照原库的main里面的demo

# 实现功能
- Who Is
- Read Property
- Write Property（目前仅支持real- float32数据写入）


# License
This library is heavily based on the gobacnet library from @REQUEA based on @alextran which is itself based on the BACnet-Stack library originally written by Steve Karg and therefore is released under the same license as his
project.  This includes the exception which allows for this library to
be linked by proprietary code without that code becoming GPL. This
exception was taken from the original BACnet stack. 

The exception is as follows:
```
    "As a special exception, if other files instantiate
     templates or use macros or inline functions from
     this file, or you compile this file and link it
     with other works to produce a work based on this file,
     this file does not by itself cause the resulting work
     to be covered by the GNU General Public License.
     However the source code for this file must still be
     made available in accordance with section (3) of the
     GNU General Public License."
```
