package vendorcoder

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

/*

# Juniper KBs
- http://kb.juniper.net/InfoCenter/index?page=content&id=KB23834
- http://kb.juniper.net/InfoCenter/index?page=content&id=KB27200

# IETF RFCs
- https://tools.ietf.org/html/rfc3925
- https://tools.ietf.org/html/rfc3396
- https://tools.ietf.org/html/rfc2939

Junos Configs
option space ztp-ops;
option ztp-ops.image-file-name code 0 = text;
option ztp-ops.config-file-name code 1 = text;
option ztp-ops.image-file-type code 2 = text;
option ztp-ops.transfer-mode code 3 = text;
option ztp-ops-encap code 43 = encapsulate ztp-ops;
*/

//VendorSubOption a data structure that defines a DHCP type 43 vendor sub option
type VendorSubOption struct {
	Code uint8
	Len  uint8
	Data []byte
}

//SetOption sets the option, data, and length of the VendorSubOption
func (vo *VendorSubOption) SetOption(code uint8, data []byte) {
	vo.Code = code
	vo.Data = data
	vo.Len = uint8(len(vo.Data))
}

//Bytes returns the structure as []byte
func (vo *VendorSubOption) Bytes() []byte {
	var finalBytes []byte
	code := byte(vo.Code)
	len := byte(vo.Len)
	finalBytes = append(finalBytes, code)
	finalBytes = append(finalBytes, len)
	finalBytes = bytes.Join([][]byte{finalBytes, vo.Data}, nil)
	return finalBytes
}

//Hex returns a hex-encoded string
//hex bool returns hex encoded with 0x
//delim the deliminator between hex bytes
func (vo *VendorSubOption) Hex(hex bool, delim string) string {
	var buffer bytes.Buffer
	var finalBytes []byte
	code := byte(vo.Code)
	len := byte(vo.Len)
	finalBytes = append(finalBytes, code)
	finalBytes = append(finalBytes, len)
	finalBytes = bytes.Join([][]byte{finalBytes, vo.Data}, nil)
	for item := range finalBytes {
		var finalString string
		if hex {
			finalString = fmt.Sprintf("%#x", finalBytes[item])
		} else {
			finalString = fmt.Sprintf("%x", finalBytes[item])
		}
		if utf8.RuneCountInString(finalString) == 1 {
			finalString = fmt.Sprintf("0%s", finalString)
		}
		buffer.WriteString(fmt.Sprintf("%s%s", finalString, delim))
	}
	return buffer.String()
	//return fmt.Sprintf("% x%s", finalBytes, delim)
}

//ByteString returns a string representation of base 10 bytes
//delim the deliminator between bytes
func (vo *VendorSubOption) ByteString(delim string) string {
	var buffer bytes.Buffer
	var finalBytes []byte
	code := byte(vo.Code)
	len := byte(vo.Len)
	finalBytes = append(finalBytes, code)
	finalBytes = append(finalBytes, len)
	finalBytes = bytes.Join([][]byte{finalBytes, vo.Data}, nil)
	for item := range finalBytes {
		finalString := fmt.Sprintf("%d", finalBytes[item])
		if utf8.RuneCountInString(finalString) == 1 {
			finalString = fmt.Sprintf("0%s", finalString)
		}
		buffer.WriteString(fmt.Sprintf("%s%s", finalString, delim))
	}
	return buffer.String()
}

//String returns the string representation of the VendorSubOption
func (vo *VendorSubOption) String() string {
	return fmt.Sprintf("Code: %d Length: %d Data: %s", vo.Code, vo.Len, vo.Data)
}
