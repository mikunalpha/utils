# utils

### Get new ptr of common type
```Go
func NewPtrString(s string) *string {}
func NewPtrBool(s bool) *bool {}
func NewPtrInt(i int) *int {}
func NewPtrInt8(i int8) *int8 {}
func NewPtrInt16(i int16) *int16 {}
func NewPtrInt32(i int32) *int32 {}
func NewPtrInt64(i int32) *int64 {}
func NewPtrUint(i uint) *uint {}
func NewPtrUint8(i uint8) *uint8 {}
func NewPtrUint16(i uint16) *uint16 {}
func NewPtrUint32(i uint32) *uint32 {}
func NewPtrUint64(i uint32) *uint64 {}
func NewPtrFloat32(f float32) *float32 {}
func NewPtrFloat64(f float64) *float64 {}
```

### Utils of slice of common type
```Go
type SliceOfString []string
func (sos SliceOfString) Unique() []string {}
func (sos SliceOfString) InSlice(t string) bool {}

type SliceOfPtrString []*string
func (sos SliceOfPtrString) Unique() []*string {}
func (sos SliceOfPtrString) InSlice(t string) bool {}

type SliceOfBool []bool
func (sos SliceOfBool) Unique() []bool {}
func (sos SliceOfBool) InSlice(t bool) bool {}

type SliceOfPtrBool []*bool
func (sos SliceOfPtrBool) Unique() []*bool {}
func (sos SliceOfPtrBool) InSlice(t bool) bool {}

type SliceOfInt []int
func (soi SliceOfInt) Unique() []int {}
func (soi SliceOfInt) InSlice(t int) bool {}

type SliceOfPtrInt []*int
func (sopi SliceOfPtrInt) Unique() []*int {}
func (sopi SliceOfPtrInt) InSlice(t int) bool {}

// type SliceOfInt8 ...
// type SliceOfInt16 ...
// type SliceOfInt32 ...
// type SliceOfInt64 ...
// type SliceOfPtrInt8 ...
// type SliceOfPtrInt16 ...
// type SliceOfPtrInt32 ...
// type SliceOfPtrInt64 ...

// type SliceOfUint ...
// type SliceOfUint8 ...
// type SliceOfUint16 ...
// type SliceOfUint32 ...
// type SliceOfUint64 ...
// type SliceOfPtrUint ...
// type SliceOfPtrUint8 ...
// type SliceOfPtrUint16 ...
// type SliceOfPtrUint32 ...
// type SliceOfPtrUint64 ...

// type SliceOfFloat32 ...
// type SliceOfFloat64 ...
// type SliceOfPtrFloat32 ...
// type SliceOfPtrFloat64 ...
```

### Get copy of http Request body
```Go
func GetRequestBody(r *http.Request) []byte {}
```

### Get variable from ENV or from file
```Go
func LoadEnvs(path string) {}
func EnvString(key string, spare string) string {}
func EnvBool(key string, spare bool) bool {}
func EnvInt(key string, spare int64) int64 {}
func EnvFloat(key string, spare float64) float64 {}
```
The file format as below:
```
VAR1=abc
VAR2=true
VAR3=10
VAR4=1.5
```
