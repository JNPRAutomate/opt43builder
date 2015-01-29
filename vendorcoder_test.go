package vendorcoder

import "testing"

func TestBasicVendorSuboption(t *testing.T) {
	newOption0 := &VendorSubOption{}
	newOption0.SetOption(00, []byte("junos-image.tgz"))
	t.Log("Option 0")
	t.Log(newOption0.Bytes())
	t.Log(newOption0.Hex(false, " "))
	t.Log(newOption0.Hex(false, ""))
	t.Log(newOption0.Hex(true, " "))
	t.Log(newOption0.Hex(true, ""))
	t.Log(newOption0.ByteString(" "))
	t.Log(newOption0.ByteString(""))
	t.Log(newOption0.String())

	newOption1 := &VendorSubOption{}
	newOption1.SetOption(01, []byte("MS-UC-Client"))
	t.Log("Option 1")
	t.Log(newOption1.Bytes())
	t.Log(newOption1.Hex(false, " "))
	t.Log(newOption1.Hex(false, ""))
	t.Log(newOption1.Hex(true, " "))
	t.Log(newOption1.Hex(true, ""))
	t.Log(newOption1.ByteString(" "))
	t.Log(newOption1.ByteString(""))
	t.Log(newOption1.String())

	newOption2 := &VendorSubOption{}
	newOption2.SetOption(02, []byte("https"))
	t.Log("Option 2")
	t.Log(newOption2.Bytes())
	t.Log(newOption2.Hex(false, " "))
	t.Log(newOption2.Hex(false, ""))
	t.Log(newOption2.Hex(true, " "))
	t.Log(newOption2.Hex(true, ""))
	t.Log(newOption2.ByteString(" "))
	t.Log(newOption2.ByteString(""))
	t.Log(newOption2.String())

	newOption3 := &VendorSubOption{}
	newOption3.SetOption(03, []byte("Ebdc1ucpool1.dimensions-uk.org"))
	t.Log("Option 3")
	t.Log(newOption3.Bytes())
	t.Log(newOption3.Hex(false, " "))
	t.Log(newOption3.Hex(false, ""))
	t.Log(newOption3.Hex(true, " "))
	t.Log(newOption3.Hex(true, ""))
	t.Log(newOption3.ByteString(" "))
	t.Log(newOption3.ByteString(""))
	t.Log(newOption3.String())

	newOption4 := &VendorSubOption{}
	newOption4.SetOption(04, []byte("443"))
	t.Log("Option 4")
	t.Log(newOption4.Bytes())
	t.Log(newOption4.Hex(false, " "))
	t.Log(newOption4.Hex(false, ""))
	t.Log(newOption4.Hex(true, " "))
	t.Log(newOption4.Hex(true, ""))
	t.Log(newOption4.ByteString(" "))
	t.Log(newOption4.ByteString(""))
	t.Log(newOption4.String())

	newOption5 := &VendorSubOption{}
	newOption5.SetOption(05, []byte("/CertProv/CertProvisioningService.svc"))
	t.Log("Option 5")
	t.Log(newOption5.Bytes())
	t.Log(newOption5.Hex(false, " "))
	t.Log(newOption5.Hex(false, ""))
	t.Log(newOption5.Hex(true, " "))
	t.Log(newOption5.Hex(true, ""))
	t.Log(newOption5.ByteString(" "))
	t.Log(newOption5.ByteString(""))
	t.Log(newOption5.String())

}
