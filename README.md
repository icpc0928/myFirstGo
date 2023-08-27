

* 格式化輸出
  * %s 字串
  * %d 十進制整數
  * %f 浮點數  如 %.3f 表示保留小數點後3位 (四捨五入)
  * %b 2進制數
  * %o 8進制數
  * %x 16進制數
  * %c 顯示ASCII碼
  * %% 輸出百分號(%) 浮點類型帶入的話就會自動轉成百分比
  * %T 數據類型
  * %v 採用默認格式表示
  
# 系統相關整數類型

| 型別      | 占用空間     |
|---------|----------|
| int     | 與平台相關    |
| int8    | 8        |
| int16   | 16       |
| int 32  | 32       |
| int 64  | 644      |
| uint    | 與平台相關    |
| uint8   | 8        |
| uint16  | 16       |
| uint32  | 32       |
| uint64  | 64       |
| byte    | 等價uint8  |
| rune    | 等價int32  |
| uintptr | 無符號的指針   |

# 浮點類型
* float32 | float64  不指定的話默認float64
# 複數類型
* complex128(64位實數和虛數)
* complex64 (32位實數和虛數)
* 聲明變量ex:
  * var complex1 complex128 = complex(2, -3)

# 切片 slice
  * 當切片的長度大於切片的容量時, 切片的容量會自動擴容
    * 切片使用append時會檢查其容量是否夠用, 如果不夠用則會調用 runtime\slice.go的 growslice方法