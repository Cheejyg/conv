package conv

import (
	"fmt"
	"strconv"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import conv "github.com/Cheejyg/conv/Go"
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// fmt
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ToString(interface_ interface{}) (string_ string) {
	return fmt.Sprint(interface_)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// strconv
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Atoi is equivalent to strconv.ParseInt(s, 10, 0), converted to type int.
func Atoi(string_ string) (int_ int, err error) {
	return strconv.Atoi(string_)
}

// Itoa is equivalent to strconv.FormatInt(int64(int_), 10).
func Itoa(int_ int) (string_ string) {
	return strconv.Itoa(int_)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Int64 conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Int64ToString parses int64 int64_ to string.
func Int64ToString(int64_ int64) (string_ string) {
	return fmt.Sprint(int64_)
}

// Int64ToPtr is equivalent to Int64ToPtrInt64(int64_).
func Int64ToPtr(int64_ int64) (ptr_int64_ *int64) {
	return Int64ToPtrInt64(int64_)
}

// Int64ToPtrString parses int64 int64_ to *string.
func Int64ToPtrString(int64_ int64) (ptr_string_ *string) {
	ptr_string_ = new(string)

	*ptr_string_ = fmt.Sprint(int64_)

	return
}

// Int64ToPtrUint64 parses int64 int64_ to *uint64.
func Int64ToPtrUint64(int64_ int64) (ptr_uint64_ *uint64) {
	ptr_uint64_ = new(uint64)

	*ptr_uint64_ = uint64(int64_)

	return
}

// Int64ToPtrUint32 parses int64 int64_ to *uint32.
func Int64ToPtrUint32(int64_ int64) (ptr_uint32_ *uint32) {
	ptr_uint32_ = new(uint32)

	*ptr_uint32_ = uint32(int64_)

	return
}

// Int64ToPtrUint16 parses int64 int64_ to *uint16.
func Int64ToPtrUint16(int64_ int64) (ptr_uint16_ *uint16) {
	ptr_uint16_ = new(uint16)

	*ptr_uint16_ = uint16(int64_)

	return
}

// Int64ToPtrUint8 parses int64 int64_ to *uint8.
func Int64ToPtrUint8(int64_ int64) (ptr_uint8_ *uint8) {
	ptr_uint8_ = new(uint8)

	*ptr_uint8_ = uint8(int64_)

	return
}

// Int64ToPtrUint parses int64 int64_ to *uint.
func Int64ToPtrUint(int64_ int64) (ptr_uint_ *uint) {
	ptr_uint_ = new(uint)

	*ptr_uint_ = uint(int64_)

	return
}

// Int64ToPtrInt64 parses int64 int64_ to *int64.
func Int64ToPtrInt64(int64_ int64) (ptr_int64_ *int64) {
	return &int64_
}

// Int64ToPtrInt32 parses int64 int64_ to *int32.
func Int64ToPtrInt32(int64_ int64) (ptr_int32_ *int32) {
	ptr_int32_ = new(int32)

	*ptr_int32_ = int32(int64_)

	return
}

// Int64ToPtrInt16 parses int64 int64_ to *int16.
func Int64ToPtrInt16(int64_ int64) (ptr_int16_ *int16) {
	ptr_int16_ = new(int16)

	*ptr_int16_ = int16(int64_)

	return
}

// Int64ToPtrInt8 parses int64 int64_ to *int8.
func Int64ToPtrInt8(int64_ int64) (ptr_int8_ *int8) {
	ptr_int8_ = new(int8)

	*ptr_int8_ = int8(int64_)

	return
}

// Int64ToPtrInt parses int64 int64_ to *int.
func Int64ToPtrInt(int64_ int64) (ptr_int_ *int) {
	ptr_int_ = new(int)

	*ptr_int_ = int(int64_)

	return
}

// Int64ToPtrBool parses int64 int64_ to *bool.
func Int64ToPtrBool(int64_ int64) (ptr_bool_ *bool) {
	ptr_bool_ = new(bool)

	if int64_ == 0 { // || int_ < 0 {
		*ptr_bool_ = false
	} else if int64_ == 1 { // || int_ > 0 {
		*ptr_bool_ = true
	} else {
		ptr_bool_ = nil
	}

	return
}

// Int64ToUint64 parses int64 int64_ to uint64.
func Int64ToUint64(int64_ int64) (uint64_ uint64) {
	uint64_ = uint64(int64_)

	return
}

// Int64ToUint32 parses int64 int64_ to uint32.
func Int64ToUint32(int64_ int64) (uint32_ uint32) {
	uint32_ = uint32(int64_)

	return
}

// Int64ToUint16 parses int64 int64_ to uint16.
func Int64ToUint16(int64_ int64) (uint16_ uint16) {
	uint16_ = uint16(int64_)

	return
}

// Int64ToUint8 parses int64 int64_ to uint8.
func Int64ToUint8(int64_ int64) (uint8_ uint8) {
	uint8_ = uint8(int64_)

	return
}

// Int64ToUint parses int64 int64_ to uint.
func Int64ToUint(int64_ int64) (uint_ uint) {
	uint_ = uint(int64_)

	return
}

// Int64ToInt64 returns int64 int64_.
func Int64ToInt64(int64_ int64) int64 {
	return int64_
}

// Int64ToInt32 parses int64 int64_ to int32.
func Int64ToInt32(int64_ int64) (int32_ int32) {
	int32_ = int32(int64_)

	return
}

// Int64ToInt16 parses int64 int64_ to int16.
func Int64ToInt16(int64_ int64) (int16_ int16) {
	int16_ = int16(int64_)

	return
}

// Int64ToInt8 parses int64 int64_ to int8.
func Int64ToInt8(int64_ int64) (int8_ int8) {
	int8_ = int8(int64_)

	return
}

// Int64ToInt parses int64 int64_ to int.
func Int64ToInt(int64_ int64) (int_ int) {
	int_ = int(int64_)

	return
}

// Int64ToBool parses int64 int64_ to bool.
func Int64ToBool(int64_ int64) (bool_ bool) {
	if int64_ == 0 { // || int_ < 0 {
		bool_ = false
	} else if int64_ == 1 { // || int_ > 0 {
		bool_ = true
	} else {
	}

	return
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Int32 conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Int32ToString parses int32 int32_ to string.
func Int32ToString(int32_ int32) (string_ string) {
	return fmt.Sprint(int32_)
}

// Int32ToPtr is equivalent to Int32ToPtrInt32(int32_).
func Int32ToPtr(int32_ int32) (ptr_int32_ *int32) {
	return Int32ToPtrInt32(int32_)
}

// Int32ToPtrString parses int32 int32_ to *string.
func Int32ToPtrString(int32_ int32) (ptr_string_ *string) {
	ptr_string_ = new(string)

	*ptr_string_ = fmt.Sprint(int32_)

	return
}

// Int32ToPtrUint64 parses int32 int32_ to *uint64.
func Int32ToPtrUint64(int32_ int32) (ptr_uint64_ *uint64) {
	ptr_uint64_ = new(uint64)

	*ptr_uint64_ = uint64(int32_)

	return
}

// Int32ToPtrUint32 parses int32 int32_ to *uint32.
func Int32ToPtrUint32(int32_ int32) (ptr_uint32_ *uint32) {
	ptr_uint32_ = new(uint32)

	*ptr_uint32_ = uint32(int32_)

	return
}

// Int32ToPtrUint16 parses int32 int32_ to *uint16.
func Int32ToPtrUint16(int32_ int32) (ptr_uint16_ *uint16) {
	ptr_uint16_ = new(uint16)

	*ptr_uint16_ = uint16(int32_)

	return
}

// Int32ToPtrUint8 parses int32 int32_ to *uint8.
func Int32ToPtrUint8(int32_ int32) (ptr_uint8_ *uint8) {
	ptr_uint8_ = new(uint8)

	*ptr_uint8_ = uint8(int32_)

	return
}

// Int32ToPtrUint parses int32 int32_ to *uint.
func Int32ToPtrUint(int32_ int32) (ptr_uint_ *uint) {
	ptr_uint_ = new(uint)

	*ptr_uint_ = uint(int32_)

	return
}

// Int32ToPtrInt64 parses int32 int32_ to *int64.
func Int32ToPtrInt64(int32_ int32) (ptr_int64_ *int64) {
	ptr_int64_ = new(int64)

	*ptr_int64_ = int64(int32_)

	return
}

// Int32ToPtrInt32 parses int32 int32_ to *int32.
func Int32ToPtrInt32(int32_ int32) (ptr_int32_ *int32) {
	return &int32_
}

// Int32ToPtrInt16 parses int32 int32_ to *int16.
func Int32ToPtrInt16(int32_ int32) (ptr_int16_ *int16) {
	ptr_int16_ = new(int16)

	*ptr_int16_ = int16(int32_)

	return
}

// Int32ToPtrInt8 parses int32 int32_ to *int8.
func Int32ToPtrInt8(int32_ int32) (ptr_int8_ *int8) {
	ptr_int8_ = new(int8)

	*ptr_int8_ = int8(int32_)

	return
}

// Int32ToPtrInt parses int32 int32_ to *int.
func Int32ToPtrInt(int32_ int32) (ptr_int_ *int) {
	ptr_int_ = new(int)

	*ptr_int_ = int(int32_)

	return
}

// Int32ToPtrBool parses int32 int32_ to *bool.
func Int32ToPtrBool(int32_ int32) (ptr_bool_ *bool) {
	ptr_bool_ = new(bool)

	if int32_ == 0 { // || int_ < 0 {
		*ptr_bool_ = false
	} else if int32_ == 1 { // || int_ > 0 {
		*ptr_bool_ = true
	} else {
		ptr_bool_ = nil
	}

	return
}

// Int32ToUint64 parses int32 int32_ to uint64.
func Int32ToUint64(int32_ int32) (uint64_ uint64) {
	uint64_ = uint64(int32_)

	return
}

// Int32ToUint32 parses int32 int32_ to uint32.
func Int32ToUint32(int32_ int32) (uint32_ uint32) {
	uint32_ = uint32(int32_)

	return
}

// Int32ToUint16 parses int32 int32_ to uint16.
func Int32ToUint16(int32_ int32) (uint16_ uint16) {
	uint16_ = uint16(int32_)

	return
}

// Int32ToUint8 parses int32 int32_ to uint8.
func Int32ToUint8(int32_ int32) (uint8_ uint8) {
	uint8_ = uint8(int32_)

	return
}

// Int32ToUint parses int32 int32_ to uint.
func Int32ToUint(int32_ int32) (uint_ uint) {
	uint_ = uint(int32_)

	return
}

// Int32ToInt64 parses int32 int32_ to int64.
func Int32ToInt64(int32_ int32) (int64_ int64) {
	int64_ = int64(int32_)

	return
}

// Int32ToInt32 returns int32 int32_.
func Int32ToInt32(int32_ int32) int32 {
	return int32_
}

// Int32ToInt16 parses int32 int32_ to int16.
func Int32ToInt16(int32_ int32) (int16_ int16) {
	int16_ = int16(int32_)

	return
}

// Int32ToInt8 parses int32 int32_ to int8.
func Int32ToInt8(int32_ int32) (int8_ int8) {
	int8_ = int8(int32_)

	return
}

// Int32ToInt parses int32 int32_ to int.
func Int32ToInt(int32_ int32) (int_ int) {
	int_ = int(int32_)

	return
}

// Int32ToBool parses int32 int32_ to bool.
func Int32ToBool(int32_ int32) (bool_ bool) {
	if int32_ == 0 { // || int_ < 0 {
		bool_ = false
	} else if int32_ == 1 { // || int_ > 0 {
		bool_ = true
	} else {
	}

	return
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Int16 conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Int16ToString parses int16 int16_ to string.
func Int16ToString(int16_ int16) (string_ string) {
	return fmt.Sprint(int16_)
}

// Int16ToPtr is equivalent to Int16ToPtrInt16(int16_).
func Int16ToPtr(int16_ int16) (ptr_int16_ *int16) {
	return Int16ToPtrInt16(int16_)
}

// Int16ToPtrString parses int16 int16_ to *string.
func Int16ToPtrString(int16_ int16) (ptr_string_ *string) {
	ptr_string_ = new(string)

	*ptr_string_ = fmt.Sprint(int16_)

	return
}

// Int16ToPtrUint64 parses int16 int16_ to *uint64.
func Int16ToPtrUint64(int16_ int16) (ptr_uint64_ *uint64) {
	ptr_uint64_ = new(uint64)

	*ptr_uint64_ = uint64(int16_)

	return
}

// Int16ToPtrUint32 parses int16 int16_ to *uint32.
func Int16ToPtrUint32(int16_ int16) (ptr_uint32_ *uint32) {
	ptr_uint32_ = new(uint32)

	*ptr_uint32_ = uint32(int16_)

	return
}

// Int16ToPtrUint16 parses int16 int16_ to *uint16.
func Int16ToPtrUint16(int16_ int16) (ptr_uint16_ *uint16) {
	ptr_uint16_ = new(uint16)

	*ptr_uint16_ = uint16(int16_)

	return
}

// Int16ToPtrUint8 parses int16 int16_ to *uint8.
func Int16ToPtrUint8(int16_ int16) (ptr_uint8_ *uint8) {
	ptr_uint8_ = new(uint8)

	*ptr_uint8_ = uint8(int16_)

	return
}

// Int16ToPtrUint parses int16 int16_ to *uint.
func Int16ToPtrUint(int16_ int16) (ptr_uint_ *uint) {
	ptr_uint_ = new(uint)

	*ptr_uint_ = uint(int16_)

	return
}

// Int16ToPtrInt64 parses int16 int16_ to *int64.
func Int16ToPtrInt64(int16_ int16) (ptr_int64_ *int64) {
	ptr_int64_ = new(int64)

	*ptr_int64_ = int64(int16_)

	return
}

// Int16ToPtrInt32 parses int16 int16_ to *int32.
func Int16ToPtrInt32(int16_ int16) (ptr_int32_ *int32) {
	ptr_int32_ = new(int32)

	*ptr_int32_ = int32(int16_)

	return
}

// Int16ToPtrInt16 parses int16 int16_ to *int16.
func Int16ToPtrInt16(int16_ int16) (ptr_int16_ *int16) {
	return &int16_
}

// Int16ToPtrInt8 parses int16 int16_ to *int8.
func Int16ToPtrInt8(int16_ int16) (ptr_int8_ *int8) {
	ptr_int8_ = new(int8)

	*ptr_int8_ = int8(int16_)

	return
}

// Int16ToPtrInt parses int16 int16_ to *int.
func Int16ToPtrInt(int16_ int16) (ptr_int_ *int) {
	ptr_int_ = new(int)

	*ptr_int_ = int(int16_)

	return
}

// Int16ToPtrBool parses int16 int16_ to *bool.
func Int16ToPtrBool(int16_ int16) (ptr_bool_ *bool) {
	ptr_bool_ = new(bool)

	if int16_ == 0 { // || int_ < 0 {
		*ptr_bool_ = false
	} else if int16_ == 1 { // || int_ > 0 {
		*ptr_bool_ = true
	} else {
		ptr_bool_ = nil
	}

	return
}

// Int16ToUint64 parses int16 int16_ to uint64.
func Int16ToUint64(int16_ int16) (uint64_ uint64) {
	uint64_ = uint64(int16_)

	return
}

// Int16ToUint32 parses int16 int16_ to uint32.
func Int16ToUint32(int16_ int16) (uint32_ uint32) {
	uint32_ = uint32(int16_)

	return
}

// Int16ToUint16 parses int16 int16_ to uint16.
func Int16ToUint16(int16_ int16) (uint16_ uint16) {
	uint16_ = uint16(int16_)

	return
}

// Int16ToUint8 parses int16 int16_ to uint8.
func Int16ToUint8(int16_ int16) (uint8_ uint8) {
	uint8_ = uint8(int16_)

	return
}

// Int16ToUint parses int16 int16_ to uint.
func Int16ToUint(int16_ int16) (uint_ uint) {
	uint_ = uint(int16_)

	return
}

// Int16ToInt64 parses int16 int16_ to int64.
func Int16ToInt64(int16_ int16) (int64_ int64) {
	int64_ = int64(int16_)

	return
}

// Int16ToInt32 parses int16 int16_ to int32.
func Int16ToInt32(int16_ int16) (int32_ int32) {
	int32_ = int32(int16_)

	return
}

// Int16ToInt16 returns int16 int16_.
func Int16ToInt16(int16_ int16) int16 {
	return int16_
}

// Int16ToInt8 parses int16 int16_ to int8.
func Int16ToInt8(int16_ int16) (int8_ int8) {
	int8_ = int8(int16_)

	return
}

// Int16ToInt parses int16 int16_ to int.
func Int16ToInt(int16_ int16) (int_ int) {
	int_ = int(int16_)

	return
}

// Int16ToBool parses int16 int16_ to bool.
func Int16ToBool(int16_ int16) (bool_ bool) {
	if int16_ == 0 { // || int_ < 0 {
		bool_ = false
	} else if int16_ == 1 { // || int_ > 0 {
		bool_ = true
	} else {
	}

	return
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Int8 conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Int8ToString parses int8 int8_ to string.
func Int8ToString(int8_ int8) (string_ string) {
	return fmt.Sprint(int8_)
}

// Int8ToPtr is equivalent to Int8ToPtrInt8(int8_).
func Int8ToPtr(int8_ int8) (ptr_int_ *int8) {
	return Int8ToPtrInt8(int8_)
}

// Int8ToPtrString parses int8 int8_ to *string.
func Int8ToPtrString(int8_ int8) (ptr_string_ *string) {
	ptr_string_ = new(string)

	*ptr_string_ = fmt.Sprint(int8_)

	return
}

// Int8ToPtrUint64 parses int8 int8_ to *uint64
func Int8ToPtrUint64(int8_ int8) (ptr_uint64_ *uint64) {
	ptr_uint64_ = new(uint64)

	*ptr_uint64_ = uint64(int8_)

	return
}

// Int8ToPtrUint32 parses int8 int8_ to *uint32
func Int8ToPtrUint32(int8_ int8) (ptr_uint32_ *uint32) {
	ptr_uint32_ = new(uint32)

	*ptr_uint32_ = uint32(int8_)

	return
}

// Int8ToPtrUint16 parses int8 int8_ to *uint16
func Int8ToPtrUint16(int8_ int8) (ptr_uint16_ *uint16) {
	ptr_uint16_ = new(uint16)

	*ptr_uint16_ = uint16(int8_)

	return
}

// Int8ToPtrUint8 parses int8 int8_ to *uint8
func Int8ToPtrUint8(int8_ int8) (ptr_uint8_ *uint8) {
	ptr_uint8_ = new(uint8)

	*ptr_uint8_ = uint8(int8_)

	return
}

// Int8ToPtrUint parses int8 int8_ to *uint
func Int8ToPtrUint(int8_ int8) (ptr_uint_ *uint) {
	ptr_uint_ = new(uint)

	*ptr_uint_ = uint(int8_)

	return
}

// Int8ToPtrInt64 parses int8 int8_ to *int64
func Int8ToPtrInt64(int8_ int8) (ptr_int64_ *int64) {
	ptr_int64_ = new(int64)

	*ptr_int64_ = int64(int8_)

	return
}

// Int8ToPtrInt32 parses int8 int8_ to *int32
func Int8ToPtrInt32(int8_ int8) (ptr_int32_ *int32) {
	ptr_int32_ = new(int32)

	*ptr_int32_ = int32(int8_)

	return
}

// Int8ToPtrInt16 parses int8 int8_ to *int16
func Int8ToPtrInt16(int8_ int8) (ptr_int16_ *int16) {
	ptr_int16_ = new(int16)

	*ptr_int16_ = int16(int8_)

	return
}

// Int8ToPtrInt8 parses int8 int8_ to *int8
func Int8ToPtrInt8(int8_ int8) (ptr_int8_ *int8) {
	return &int8_
}

// Int8ToPtrInt parses int8 int8_ to *int
func Int8ToPtrInt(int8_ int8) (ptr_int_ *int) {
	ptr_int_ = new(int)

	*ptr_int_ = int(int8_)

	return
}

// Int8ToPtrBool parses int8 int8_ to *bool
func Int8ToPtrBool(int8_ int8) (ptr_bool_ *bool) {
	ptr_bool_ = new(bool)

	if int8_ == 0 { // || int_ < 0 {
		*ptr_bool_ = false
	} else if int8_ == 1 { // || int_ > 0 {
		*ptr_bool_ = true
	} else {
		ptr_bool_ = nil
	}

	return
}

// Int8ToUint64 parses int8 int8_ to uint64.
func Int8ToUint64(int8_ int8) (uint64_ uint64) {
	uint64_ = uint64(int8_)

	return
}

// Int8ToUint32 parses int8 int8_ to uint32.
func Int8ToUint32(int8_ int8) (uint32_ uint32) {
	uint32_ = uint32(int8_)

	return
}

// Int8ToUint16 parses int8 int8_ to uint16.
func Int8ToUint16(int8_ int8) (uint16_ uint16) {
	uint16_ = uint16(int8_)

	return
}

// Int8ToUint8 parses int8 int8_ to uint8.
func Int8ToUint8(int8_ int8) (uint8_ uint8) {
	uint8_ = uint8(int8_)

	return
}

// Int8ToUint parses int8 int8_ to uint.
func Int8ToUint(int8_ int8) (uint_ uint) {
	uint_ = uint(int8_)

	return
}

// Int8ToInt64 parses int8 int8_ to int64.
func Int8ToInt64(int8_ int8) (int64_ int64) {
	int64_ = int64(int8_)

	return
}

// Int8ToInt32 parses int8 int8_ to int32.
func Int8ToInt32(int8_ int8) (int32_ int32) {
	int32_ = int32(int8_)

	return
}

// Int8ToInt16 parses int8 int8_ to int16.
func Int8ToInt16(int8_ int8) (int16_ int16) {
	int16_ = int16(int8_)

	return
}

// Int8ToInt8 returns int8 int8_.
func Int8ToInt8(int8_ int8) int8 {
	return int8_
}

// Int8ToInt parses int8 int8_ to int.
func Int8ToInt(int8_ int8) (int_ int) {
	int_ = int(int8_)

	return
}

// Int8ToBool parses int8 int8_ to bool.
func Int8ToBool(int8_ int8) (bool_ bool) {
	if int8_ == 0 { // || int_ < 0 {
		bool_ = false
	} else if int8_ == 1 { // || int_ > 0 {
		bool_ = true
	} else {
	}

	return
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Int conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// IntToString parses int int_ to string.
func IntToString(int_ int) (string_ string) {
	return fmt.Sprint(int_)
}

// IntToPtr is equivalent to IntToPtrInt(int_).
func IntToPtr(int_ int) (ptr_int_ *int) {
	return IntToPtrInt(int_)
}

// IntToPtrString parses int int_ to *string.
func IntToPtrString(int_ int) (ptr_string_ *string) {
	ptr_string_ = new(string)

	*ptr_string_ = fmt.Sprint(int_)

	return
}

// IntToPtrUint64 parses int int_ to *uint64.
func IntToPtrUint64(int_ int) (ptr_uint64_ *uint64) {
	ptr_uint64_ = new(uint64)

	*ptr_uint64_ = uint64(int_)

	return
}

// IntToPtrUint32 parses int int_ to *uint32.
func IntToPtrUint32(int_ int) (ptr_uint32_ *uint32) {
	ptr_uint32_ = new(uint32)

	*ptr_uint32_ = uint32(int_)

	return
}

// IntToPtrUint16 parses int int_ to *uint16.
func IntToPtrUint16(int_ int) (ptr_uint16_ *uint16) {
	ptr_uint16_ = new(uint16)

	*ptr_uint16_ = uint16(int_)

	return
}

// IntToPtrUint8 parses int int_ to *uint8.
func IntToPtrUint8(int_ int) (ptr_uint8_ *uint8) {
	ptr_uint8_ = new(uint8)

	*ptr_uint8_ = uint8(int_)

	return
}

// IntToPtrUint parses int int_ to *uint.
func IntToPtrUint(int_ int) (ptr_uint_ *uint) {
	ptr_uint_ = new(uint)

	*ptr_uint_ = uint(int_)

	return
}

// IntToPtrInt64 parses int int_ to *int64.
func IntToPtrInt64(int_ int) (ptr_int64_ *int64) {
	ptr_int64_ = new(int64)

	*ptr_int64_ = int64(int_)

	return
}

// IntToPtrInt32 parses int int_ to *int32.
func IntToPtrInt32(int_ int) (ptr_int32_ *int32) {
	ptr_int32_ = new(int32)

	*ptr_int32_ = int32(int_)

	return
}

// IntToPtrInt16 parses int int_ to *int16.
func IntToPtrInt16(int_ int) (ptr_int16_ *int16) {
	ptr_int16_ = new(int16)

	*ptr_int16_ = int16(int_)

	return
}

// IntToPtrInt8 parses int int_ to *int8.
func IntToPtrInt8(int_ int) (ptr_int8_ *int8) {
	ptr_int8_ = new(int8)

	*ptr_int8_ = int8(int_)

	return
}

// IntToPtrInt parses int int_ to *int.
func IntToPtrInt(int_ int) (ptr_int_ *int) {
	return &int_
}

// IntToPtrBool parses int int_ to *bool.
func IntToPtrBool(int_ int) (ptr_bool_ *bool) {
	ptr_bool_ = new(bool)

	if int_ == 0 { // || int_ < 0 {
		*ptr_bool_ = false
	} else if int_ == 1 { // || int_ > 0 {
		*ptr_bool_ = true
	} else {
		ptr_bool_ = nil
	}

	return
}

// IntToUint64 parses int int_ to uint64.
func IntToUint64(int_ int) (uint64_ uint64) {
	return uint64(int_)
}

// IntToUint32 parses int int_ to uint32.
func IntToUint32(int_ int) (uint32_ uint32) {
	return uint32(int_)
}

// IntToUint16 parses int int_ to uint16.
func IntToUint16(int_ int) (uint16_ uint16) {
	return uint16(int_)
}

// IntToUint8 parses int int_ to uint8.
func IntToUint8(int_ int) (uint8_ uint8) {
	return uint8(int_)
}

// IntToUint parses int int_ to uint.
func IntToUint(int_ int) (uint_ uint) {
	return uint(int_)
}

// IntToInt64 parses int int_ to int64.
func IntToInt64(int_ int) (int64_ int64) {
	return int64(int_)
}

// IntToInt32 parses int int_ to int32.
func IntToInt32(int_ int) (int32_ int32) {
	return int32(int_)
}

// IntToInt16 parses int int_ to int16.
func IntToInt16(int_ int) (int16_ int16) {
	return int16(int_)
}

// IntToInt8 parses int int_ to int8.
func IntToInt8(int_ int) (int8_ int8) {
	return int8(int_)
}

// IntToInt returns int int_.
func IntToInt(int_ int) int {
	return int_
}

// IntToBool parses int int_ to bool.
func IntToBool(int_ int) (bool_ bool) {
	if int_ == 0 { // || int_ < 0 {
		bool_ = false
	} else if int_ == 1 { // || int_ > 0 {
		bool_ = true
	} else {
	}

	return
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Bool conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// BoolToString parses bool bool_ to string.
func BoolToString(bool_ bool) (string_ string) {
	return strconv.FormatBool(bool_)
}

// BoolToPtr is equivalent to BoolToPtrBool(bool_).
func BoolToPtr(bool_ bool) (ptr_bool_ *bool) {
	return BoolToPtrBool(bool_)
}

// BoolToPtrString parses bool bool_ to *string.
func BoolToPtrString(bool_ bool) (ptr_string_ *string) {
	ptr_string_ = new(string)

	*ptr_string_ = strconv.FormatBool(bool_)

	return
}

// BoolToPtrUint parses bool bool_ to *uint.
func BoolToPtrUint(bool_ bool) (ptr_uint_ *uint) {
	ptr_uint_ = new(uint)

	switch bool_ {
	case false:
		*ptr_uint_ = 0
	case true:
		*ptr_uint_ = 1
	default:
		ptr_uint_ = nil
	}

	return
}

// BoolToPtrInt parses bool bool_ to *int.
func BoolToPtrInt(bool_ bool) (ptr_int_ *int) {
	ptr_int_ = new(int)

	switch bool_ {
	case false:
		*ptr_int_ = 0
	case true:
		*ptr_int_ = 1
	default:
		ptr_int_ = nil
	}

	return
}

// BoolToPtrBool parses bool bool_ to *bool.
func BoolToPtrBool(bool_ bool) (ptr_bool_ *bool) {
	return &bool_
}

// BoolToUint parses bool bool_ to uint.
func BoolToUint(bool_ bool) (uint_ uint) {
	switch bool_ {
	case false:
		uint_ = 0
	case true:
		uint_ = 1
	default:
	}

	return
}

// BoolToInt parses bool bool_ to int.
func BoolToInt(bool_ bool) (int_ int) {
	switch bool_ {
	case false:
		int_ = 0
	case true:
		int_ = 1
	default:
	}

	return
}

// BoolToBool returns bool bool_.
func BoolToBool(bool_ bool) bool {
	return bool_
}
