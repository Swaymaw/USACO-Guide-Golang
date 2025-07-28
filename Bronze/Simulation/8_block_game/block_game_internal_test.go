package main

import "testing"

func compareWordCounts(counts1, counts2 map[rune]int) bool {
	for _, val := range "abcdefghijklmnopqrstuvwxyz" {
		if counts1[val] != counts2[val] {
			return false
		}
	}
	return true
}

func TestGetNumWords(t *testing.T) {
	type testCase struct {
		blocks []Block
		want   map[rune]int
	}

	tt := map[string]testCase{
		"sample input": {
			blocks: []Block{{front: "fox", back: "box"}, {front: "dog", back: "cat"}, {front: "car", back: "bus"}},
			want:   map[rune]int{'a': 2, 'b': 2, 'c': 2, 'd': 1, 'f': 1, 'g': 1, 'o': 2, 'r': 1, 's': 1, 't': 1, 'u': 1, 'x': 1},
		},
		"example": {
			blocks: []Block{{front: "wlrb", back: "mqbhcd"}, {front: "r", back: "owkk"}, {front: "hiddqsc", back: "xrjm"}, {front: "wfr", back: "sjyb"}},
			want:   map[rune]int{'b': 2, 'c': 2, 'd': 3, 'f': 1, 'h': 2, 'i': 1, 'j': 2, 'k': 2, 'l': 1, 'm': 2, 'o': 1, 'q': 2, 'r': 4, 's': 2, 'w': 3, 'x': 1, 'y': 1},
		},
		"example2": {
			blocks: []Block{{front: "js", back: "koel"}, {front: "ynzeav", back: "aceazui"}, {front: "ydypvmgyx", back: "lhppuunk"}, {front: "tkqtmv", back: "nuuvjva"}, {front: "mvvuvsvh", back: "kywh"}, {front: "gchqvdcqd", back: "mz"}, {front: "xwn", back: "ikzfgta"}, {front: "tnlpwzva", back: "nvkplpfaot"}, {front: "ngexrfcz", back: "dmus"}, {front: "lo", back: "iokvkwkx"}, {front: "rxbl", back: "ot"}, {front: "omeqlftx", back: "lzkb"}, {front: "sqmncia", back: "vrzyeupyvd"}, {front: "btwhp", back: "gcwteylw"}, {front: "byubxru", back: "szshxpm"}, {front: "rhfawdibzb", back: "ypdksb"}, {front: "ta", back: "pzs"}, {front: "r", back: "nj"}, {front: "zcxecvjm", back: "jqdjhgvzj"}, {front: "ukf", back: "jzac"}, {front: "pnsoppvtle", back: "jpy"}, {front: "yfvuefyy", back: "dj"}, {front: "d", back: "egbsyp"}, {front: "omfbrn", back: "ilzhsvb"}, {front: "czwtdfvir", back: "osvmjfcy"}, {front: "d", back: "zq"}, {front: "kpge", back: "jaojyjofe"}, {front: "wim", back: "dac"}, {front: "dawitxysj", back: "zncipttnc"}, {front: "tjhrtvkwwo", back: "bqhjjfkb"}, {front: "acc", back: "nrxihcsan"}, {front: "tgxdcttn", back: "j"}, {front: "fs", back: "rwqtyuy"}, {front: "mxwvbqxorq", back: "owzhpmdzj"}, {front: "rlcncyoywb", back: "vzhxpen"}, {front: "vivthf", back: "vkhfx"}, {front: "qaquyetw", back: "fthns"}, {front: "rggoqbhxiq", back: "v"}, {front: "zscqutms", back: "fqjnmtae"}, {front: "tmykcbr", back: "kj"}, {front: "hltzn", back: "ui"}, {front: "okf", back: "vsto"}, {front: "z", back: "qmeao"}, {front: "r", back: "sdmzohydt"}, {front: "otjyy", back: "ttlknmqdy"}, {front: "dpbxftatuq", back: "stvphoa"}, {front: "psdifnxr", back: "bqaghd"}, {front: "oyhhsxhntd", back: "ctc"}, {front: "iupqzeqsj", back: "kmzr"}, {front: "hfoooioy", back: "jdxnwbzhvq"}, {front: "wuprgibsit", back: "pazf"}, {front: "itpf", back: "sfsqgrx"}, {front: "k", back: "cgmtmvvg"}, {front: "qq", back: "spdjxza"}, {front: "dukv", back: "qpkuzjzhws"}, {front: "tpcafsyai", back: "jham"}, {front: "egwbtpq", back: "lrnkbld"}, {front: "guzgcsecci", back: "li"}, {front: "yogw", back: "z"}, {front: "ifxc", back: "hdgmanjztl"}, {front: "anviaj", back: "chhkdxlrf"}, {front: "ohmx", back: "smmhv"}, {front: "dihdvfnxo", back: "bxmlclxvro"}, {front: "acjp", back: "xbe"}, {front: "rhcbm", back: "gzwwgyvtlz"}, {front: "ivxyg", back: "appzo"}, {front: "dik", back: "mlwltxi"}, {front: "dioytnfeie", back: "pehgvgv"}, {front: "jfavsnt", back: "iqnbvxpu"}, {front: "cz", back: "nfdc"}, {front: "k", back: "h"}, {front: "hxdnz", back: "hormwjcxf"}, {front: "ob", back: "rcveh"}, {front: "itnsdgacjp", back: "i"}, {front: "sbmrnoq", back: "sfvoyxkeg"}, {front: "ma", back: "g"}, {front: "gihq", back: "nazgdmzq"}, {front: "piuud", back: "ucvyjimliv"}, {front: "pdzhf", back: "hevksudvjl"}, {front: "grcavx", back: "ehlrqgjh"}, {front: "jqt", back: "zztjsnopy"}, {front: "get", back: "fq"}, {front: "exq", back: "oi"}, {front: "yro", back: "hjhgia"}, {front: "cpgrni", back: "s"}, {front: "hztp", back: "qhwh"}, {front: "yfio", back: "p"}, {front: "xvgxniovab", back: "atgj"}, {front: "zazsi", back: "zzwei"}, {front: "ux", back: "r"}, {front: "qqkzef", back: "hi"}, {front: "dqmx", back: "mx"}, {front: "wt", back: "w"}, {front: "ogckdl", back: "adtkczpf"}, {front: "zikp", back: "j"}, {front: "jg", back: "fbb"}, {front: "trhvcnif", back: "mbaqapyj"}, {front: "gvgdfp", back: "lir"}, {front: "jvgaed", back: "t"}, {front: "kjclj", back: "naqzr"}},
			want:   map[rune]int{'a': 40, 'b': 30, 'c': 42, 'd': 41, 'e': 30, 'f': 39, 'g': 40, 'h': 45, 'i': 43, 'j': 44, 'k': 31, 'l': 28, 'm': 35, 'n': 39, 'o': 38, 'p': 42, 'q': 41, 'r': 34, 's': 39, 't': 49, 'u': 26, 'v': 50, 'w': 28, 'x': 39, 'y': 40, 'z': 49},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := getNumWords(tc.blocks)

			if !compareWordCounts(got, tc.want) {
				t.Fatalf("failed for test %s with got: %v and\nexpected:%v\n", name, got, tc.want)
			}
		})
	}
}
