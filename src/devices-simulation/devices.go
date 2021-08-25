package main

import "fmt"

func TopMAC(num int) ([]string, error) {

	if num < 0 || num > 50 {
		return nil, fmt.Errorf("num of desired MAC Adresses out of range [1 - 50]: passed %d", num)
	}

	mac := []string{
		"7c:7f:01:2c:10:cd",
		"68:e0:32:c9:e2:47",
		"83:50:4a:5b:b7:1d",
		"30:c2:65:0a:5d:7a",
		"10:3c:4a:e3:36:eb",
		"e8:ad:b9:1d:9d:9e",
		"57:13:5a:c3:47:3d",
		"0b:97:75:77:33:6c",
		"ea:b7:ba:76:b1:2c",
		"f0:06:8f:5b:54:72",
		"f3:64:42:7c:e9:8b",
		"6f:d8:1a:e4:0f:e7",
		"da:cd:3f:11:fb:4b",
		"ca:c1:95:7b:03:d5",
		"94:de:57:51:02:f4",
		"fc:a9:2e:b3:7c:38",
		"39:63:df:31:7b:d2",
		"1e:7f:c9:03:74:c6",
		"15:21:b8:82:bf:2f",
		"96:e2:c1:11:3a:9e",
		"25:0a:1a:b4:e9:c0",
		"72:ad:f8:9e:a5:b0",
		"81:1a:f0:0e:1b:63",
		"67:18:ef:08:7c:e5",
		"bf:bd:18:70:67:e6",
		"d3:32:8c:e0:7e:5c",
		"58:cc:43:ed:36:ec",
		"54:15:48:e8:b0:35",
		"35:4d:7b:39:b5:cb",
		"fa:cd:9e:eb:91:c6",
		"a6:11:4a:1e:79:6e",
		"38:13:3a:52:74:12",
		"3e:35:19:ad:3e:ce",
		"77:9e:71:ca:56:fe",
		"1b:6c:5e:2e:c3:ee",
		"7a:5e:4c:8f:7d:70",
		"6c:5b:22:0f:92:e5",
		"09:7f:e8:52:65:1b",
		"ce:b1:44:de:e3:2b",
		"6a:a4:10:c4:8b:42",
		"56:2c:5f:23:c8:21",
		"a1:64:70:a8:f6:ab",
		"f3:29:5b:32:65:55",
		"e3:91:5b:b9:fa:9c",
		"49:69:83:38:ae:8d",
		"a3:79:bf:e4:0f:c4",
		"af:7b:c5:32:38:6e",
		"ce:59:19:52:37:35",
		"53:51:e2:c7:30:cc",
		"da:6a:0b:8f:81:59",
	}

	return mac[:num], nil
}
