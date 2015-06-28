package encoding

import (
	"testing"
)

type compressedValues struct {
	EncodedUri string
	Base64     string
}

var testingValues map[string]compressedValues = map[string]compressedValues{
	"a":      {"IZA", "IZA="},
	"aa":     {"IbI", "IbI="},
	"ààà":    {"Aco", "Aco="},
	"àààààà": {"Ac1A", "Ac1A"},
	"toto":   {"C4e1Q", "C4e1Q==="},
	"Å∑⁄¢∆Ê‚√∫ﬂı":                                                                                {"KORCJCIEBFhgRAU4WASCwiQ1ESCBvwjIBA", "KORCJCIEBFhgRAU4WASCwiQ1ESCBvwjIBA=="},
	"ce train est à destination":                                                                 {"MYUwBALgTghglgOzCAzhMADsATVFEz4D2CQA", "MYUwBALgTghglgOzCAzhMADsATVFEz4D2CQA"},
	"Ce train est à destination de PARIS SAINT LAZARE: Vous en êtes à 45%; Oui c'est très long.": {"MIUwBALgTghglgOzCAzhMADsATVFEz4D2SuYACgIIBKAkgMpj2W0ByAKmADKUBaNAUQBcYAGpEArimRIAVxFSYwAFgCsAUgDcYAPIS4YAMYByPJCgAL6QBsSAcwB0QA", "MIUwBALgTghglgOzCAzhMADsATVFEz4D2SuYACgIIBKAkgMpj2W0ByAKmADKUBaNAUQBcYAGpEArimRIAVxFSYwAFgCsAUgDcYAPIS4YAMYByPJCgAL6QBsSAcwB0QA="},
	"{\"id\":\"1234534625254\",\"at\":2,\"tmax\":120,\"imp\":[{\"id\":\"1\",\"banner\":{\"w\":1583,\"h\":1095,\"pos\":7,\"battr\":[13]}}, {\"id\":\"2\",\"banner\":{\"w\":784,\"h\":100,\"pos\":4,\"battr\":[13]}}],\"badv\":[\"company1.com\",\"company2.com\"],\"site\":{\"id\":\"234563\",\"name\":\"Site ABCD\",\"domain\":\"www.auchandrive.fr\",\"pagecat\":[\"3700610\"],\"sectioncat\":[\"3700624\"],\"privacypolicy\":1,\"page\":\"http://siteabcd.com/page.htm\",\"ref\":\"http://referringsite.com/referringpage.htm\",\"publisher\":{\"id\":\"pub12345\",\"name\":\"Publisher A\"},\"content\":{\"keywords\":[\"keyword a\",\"keyword b\",\"keyword c\"]}},\"device\":{\"ip\":\"64.124.253.1\",\"ua\":\"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.2.16) Gecko/20110319 Firefox/3.6.16\",\"os\":\"OS X\",\"flashver\":\"10.1\",\"js\":1},\"user\":{\"id\":\"45asdf987656789adfad4678rew656789\",\"buyeruid\":\"5df678asd8987656asdf78987654\",\"data\":[{\"name\":\"AuchanDrive\", \"segment\":[{\"name\":\"basket\", \"value\":\"[]\"}]}]}}": {"N4IglgJiBcIIwCYDMAWArKgbAtOUgBoQBDAFxgSNIFtiAPGRABiLGoAcYBtUSGeQiABGxAHaiApgCcYoAO6M0ADiREAFoyYBONEXYB7AM4wA7ERGlSM6FzhIAugF9HBAAS8osBIJHjpskAVoEyUUdU0WEANjaDDhMituOydHe3NiCAA3bhAAY30OMQBPOAA6fOpBCvZihHKCkDSQQzBSCQC+L1Q0TCRBUWJqdtgAZVaJVwBBACEAYQARQQgC4jBRfjlN0uIAV1y1MQgpMEyJUoAzGT1iAHMJXLIcpBMmJkw4JkaiQ3vSMH1RA9yDYQM9Xth8E12MdMsRckUDAAbMDwxjXO78NSWdjQAD0uJabWIQlyEHq1FxNTupSxlSIUgk50x2LxuIZ52kx1EN0JZwqbMZnLWNypZ1pgnYOyEyMMan80A8-ElQkQ3X6g2GIAAClKZXKpFMQC48gC2qJgaAANYSIpyfRSCAxLgga22+0QVzEQSuu0O1xCb0230e3KNZxECASTIo4a8TiwTAoUqIJM4JDJwQ7L2wACy+gAXmBEYjiLi0KUmK4ABQ5uFrUhGNQAblcAFUWwBJc0SRGuWu5VwAeRGrgAGq4PqVMC2JKIALStkYtqSZaBlLSlOpwTAASlcAHF7pb9LiEEw4B8kHAtK4AGJgdn6Oi49OYZOYQRGfjDseCc4l2VTmseAmAzIgACsYjgY0dh+axFVgdBiEMCBzi0JQTEwHoQi0DJzgyFBMBCBk5CwoilC0HwdiKaQdk6EA0FQ8jkIgCiMLIljzhw9i0HwCMyGzHgQAGIZ+EmPYDlEeYYXaNxmgkG4hnNbhQBEzUREMa1yDk2FER2TUuHsI0nBSIA", "N4IglgJiBcIIwCYDMAWArKgbAtOUgBoQBDAFxgSNIFtiAPGRABiLGoAcYBtUSGeQiABGxAHaiApgCcYoAO6M0ADiREAFoyYBONEXYB7AM4wA7ERGlSM6FzhIAugF9HBAAS8osBIJHjpskAVoEyUUdU0WEANjaDDhMituOydHe3NiCAA3bhAAY30OMQBPOAA6fOpBCvZihHKCkDSQQzBSCQC+L1Q0TCRBUWJqdtgAZVaJVwBBACEAYQARQQgC4jBRfjlN0uIAV1y1MQgpMEyJUoAzGT1iAHMJXLIcpBMmJkw4JkaiQ3vSMH1RA9yDYQM9Xth8E12MdMsRckUDAAbMDwxjXO78NSWdjQAD0uJabWIQlyEHq1FxNTupSxlSIUgk50x2LxuIZ52kx1EN0JZwqbMZnLWNypZ1pgnYOyEyMMan80A8/ElQkQ3X6g2GIAAClKZXKpFMQC48gC2qJgaAANYSIpyfRSCAxLgga22+0QVzEQSuu0O1xCb0230e3KNZxECASTIo4a8TiwTAoUqIJM4JDJwQ7L2wACy+gAXmBEYjiLi0KUmK4ABQ5uFrUhGNQAblcAFUWwBJc0SRGuWu5VwAeRGrgAGq4PqVMC2JKIALStkYtqSZaBlLSlOpwTAASlcAHF7pb9LiEEw4B8kHAtK4AGJgdn6Oi49OYZOYQRGfjDseCc4l2VTmseAmAzIgACsYjgY0dh+axFVgdBiEMCBzi0JQTEwHoQi0DJzgyFBMBCBk5CwoilC0HwdiKaQdk6EA0FQ8jkIgCiMLIljzhw9i0HwCMyGzHgQAGIZ+EmPYDlEeYYXaNxmgkG4hnNbhQBEzUREMa1yDk2FER2TUuHsI0nBSIA=="},
}

func TestDecompress(t *testing.T) {
	for expected, compressed := range testingValues {
		result, err := DecompressFromEncodedUriComponent(compressed.EncodedUri)
		if err != nil {
			t.Errorf("Unexpected error", err)
		}
		if result != expected {
			t.Errorf("Encoded URI result should be %s instead of %s\n", expected, result)
		}
		result, err = DecompressFromBase64(compressed.Base64)
		if err != nil {
			t.Errorf("Unexpected error", err)
		}
		if result != expected {
			t.Errorf("Encoded URI result should be %s instead of %s\n", expected, result)
		}
	}
}

func BenchmarkDecompress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, k := range testingValues {
			DecompressFromEncodedUriComponent(k.EncodedUri)
		}
	}
}
