package main

import "errors"

type VersionMetaData struct {
	Version      int
	ECL          ErrorCorrectionLevel
	NumericMode  int
	AlphaMode    int
	ByteMode     int
	KanjiMode    int
	MaxCodewords int
	ECPerBlock   int
	NumBlocks    int
	NumCW        int
	NumBlocksG2  int
	NumCWG2      int
}

var charCapTable []VersionMetaData = []VersionMetaData{
	{1, L, 41, 25, 17, 10, 19, 7, 1, 19, 0, 0},
	{1, M, 34, 20, 14, 8, 16, 10, 1, 16, 0, 0},
	{1, Q, 27, 16, 11, 7, 13, 13, 1, 13, 0, 0},
	{1, H, 17, 10, 7, 4, 9, 17, 1, 9, 0, 0},
	{2, L, 77, 47, 32, 20, 34, 10, 1, 34, 0, 0},
	{2, M, 63, 38, 26, 16, 28, 16, 1, 28, 0, 0},
	{2, Q, 48, 29, 20, 12, 22, 22, 1, 22, 0, 0},
	{2, H, 34, 20, 14, 8, 16, 28, 1, 16, 0, 0},
	{3, L, 127, 77, 53, 32, 55, 15, 1, 55, 0, 0},
	{3, M, 101, 61, 42, 26, 44, 26, 1, 44, 0, 0},
	{3, Q, 77, 47, 32, 20, 34, 18, 2, 17, 0, 0},
	{3, H, 58, 35, 24, 15, 26, 22, 2, 13, 0, 0},
	{4, L, 187, 114, 78, 48, 80, 20, 1, 80, 0, 0},
	{4, M, 149, 90, 62, 38, 64, 18, 2, 32, 0, 0},
	{4, Q, 111, 67, 46, 28, 48, 26, 2, 24, 0, 0},
	{4, H, 82, 50, 34, 21, 36, 16, 4, 9, 0, 0},
	{5, L, 255, 154, 106, 65, 108, 26, 1, 108, 0, 0},
	{5, M, 202, 122, 84, 52, 86, 24, 2, 43, 0, 0},
	{5, Q, 144, 87, 60, 37, 62, 18, 2, 15, 2, 16},
	{5, H, 106, 64, 44, 27, 46, 22, 2, 11, 2, 12},
	{6, L, 322, 195, 134, 82, 136, 18, 2, 68, 0, 0},
	{6, M, 255, 154, 106, 65, 108, 16, 4, 27, 0, 0},
	{6, Q, 178, 108, 74, 45, 76, 24, 4, 19, 0, 0},
	{6, H, 139, 84, 58, 36, 60, 28, 4, 15, 0, 0},
	{7, L, 370, 224, 154, 95, 156, 20, 2, 78, 0, 0},
	{7, M, 293, 178, 122, 75, 124, 18, 4, 31, 0, 0},
	{7, Q, 207, 125, 86, 53, 88, 18, 2, 14, 4, 15},
	{7, H, 154, 93, 64, 39, 66, 26, 4, 13, 1, 14},
	{8, L, 461, 279, 192, 118, 194, 24, 2, 97, 0, 0},
	{8, M, 365, 221, 152, 93, 154, 22, 2, 38, 2, 39},
	{8, Q, 259, 157, 108, 66, 110, 22, 4, 18, 2, 19},
	{8, H, 202, 122, 84, 52, 86, 26, 4, 14, 2, 15},
	{9, L, 552, 335, 230, 141, 232, 30, 2, 116, 0, 0},
	{9, M, 432, 262, 180, 111, 182, 22, 3, 36, 2, 37},
	{9, Q, 312, 189, 130, 80, 132, 20, 4, 16, 4, 17},
	{9, H, 235, 143, 98, 60, 100, 24, 4, 12, 4, 13},
	{10, L, 652, 395, 271, 167, 274, 18, 2, 68, 2, 69},
	{10, M, 513, 311, 213, 131, 216, 26, 4, 43, 1, 44},
	{10, Q, 364, 221, 151, 93, 154, 24, 6, 19, 2, 20},
	{10, H, 288, 174, 119, 74, 122, 28, 6, 15, 2, 16},
	{11, L, 772, 468, 321, 198, 324, 20, 4, 81, 0, 0},
	{11, M, 604, 366, 251, 155, 254, 30, 1, 50, 4, 51},
	{11, Q, 427, 259, 177, 109, 180, 28, 4, 22, 4, 23},
	{11, H, 331, 200, 137, 85, 140, 24, 3, 12, 8, 13},
	{12, L, 883, 535, 367, 226, 370, 24, 2, 92, 2, 93},
	{12, M, 691, 419, 287, 177, 290, 22, 6, 36, 2, 37},
	{12, Q, 489, 296, 203, 125, 206, 26, 4, 20, 6, 21},
	{12, H, 374, 227, 155, 96, 158, 28, 7, 14, 4, 15},
	{13, L, 1022, 619, 425, 262, 428, 26, 4, 107, 0, 0},
	{13, M, 796, 483, 331, 204, 334, 22, 8, 37, 1, 38},
	{13, Q, 580, 352, 241, 149, 244, 24, 8, 20, 4, 21},
	{13, H, 427, 259, 177, 109, 180, 22, 12, 11, 4, 12},
	{14, L, 1101, 667, 458, 282, 461, 30, 3, 115, 1, 116},
	{14, M, 871, 528, 362, 223, 365, 24, 4, 40, 5, 41},
	{14, Q, 621, 376, 258, 159, 261, 20, 11, 16, 5, 17},
	{14, H, 468, 283, 194, 120, 197, 24, 11, 12, 5, 13},
	{15, L, 1250, 758, 520, 320, 523, 22, 5, 87, 1, 88},
	{15, M, 991, 600, 412, 254, 415, 24, 5, 41, 5, 42},
	{15, Q, 703, 426, 292, 180, 295, 30, 5, 24, 7, 25},
	{15, H, 530, 321, 220, 136, 223, 24, 11, 12, 7, 13},
	{16, L, 1408, 854, 586, 361, 589, 24, 5, 98, 1, 99},
	{16, M, 1082, 656, 450, 277, 453, 28, 7, 45, 3, 46},
	{16, Q, 775, 470, 322, 198, 325, 24, 15, 19, 2, 20},
	{16, H, 602, 365, 250, 154, 253, 30, 3, 15, 13, 16},
	{17, L, 1548, 938, 644, 397, 647, 28, 1, 107, 5, 108},
	{17, M, 1212, 734, 504, 310, 507, 28, 10, 46, 1, 47},
	{17, Q, 876, 531, 364, 224, 367, 28, 1, 22, 15, 23},
	{17, H, 674, 408, 280, 173, 283, 28, 2, 14, 17, 15},
	{18, L, 1725, 1046, 718, 442, 721, 30, 5, 120, 1, 121},
	{18, M, 1346, 816, 560, 345, 563, 26, 9, 43, 4, 44},
	{18, Q, 948, 574, 394, 243, 397, 28, 17, 22, 1, 23},
	{18, H, 746, 452, 310, 191, 313, 28, 2, 14, 19, 15},
	{19, L, 1903, 1153, 792, 488, 795, 28, 3, 113, 4, 114},
	{19, M, 1500, 909, 624, 384, 627, 26, 3, 44, 11, 45},
	{19, Q, 1063, 644, 442, 272, 445, 26, 17, 21, 4, 22},
	{19, H, 813, 493, 338, 208, 341, 26, 9, 13, 16, 14},
	{20, L, 2061, 1249, 858, 528, 861, 28, 3, 107, 5, 108},
	{20, M, 1600, 970, 666, 410, 669, 26, 3, 41, 13, 42},
	{20, Q, 1159, 702, 482, 297, 485, 30, 15, 24, 5, 25},
	{20, H, 919, 557, 382, 235, 385, 28, 15, 15, 10, 16},
	{21, L, 2232, 1352, 929, 572, 932, 28, 4, 116, 4, 117},
	{21, M, 1708, 1035, 711, 438, 714, 26, 17, 42, 0, 0},
	{21, Q, 1224, 742, 509, 314, 512, 28, 17, 22, 6, 23},
	{21, H, 969, 587, 403, 248, 406, 30, 19, 16, 6, 17},
	{22, L, 2409, 1460, 1003, 618, 1006, 28, 2, 111, 7, 112},
	{22, M, 1872, 1134, 779, 480, 782, 28, 17, 46, 0, 0},
	{22, Q, 1358, 823, 565, 348, 568, 30, 7, 24, 16, 25},
	{22, H, 1056, 640, 439, 270, 442, 24, 34, 13, 0, 0},
	{23, L, 2620, 1588, 1091, 672, 1094, 30, 4, 121, 5, 122},
	{23, M, 2059, 1248, 857, 528, 860, 28, 4, 47, 14, 48},
	{23, Q, 1468, 890, 611, 376, 614, 30, 11, 24, 14, 25},
	{23, H, 1108, 672, 461, 284, 464, 30, 16, 15, 14, 16},
	{24, L, 2812, 1704, 1171, 721, 1174, 30, 6, 117, 4, 118},
	{24, M, 2188, 1326, 911, 561, 914, 28, 6, 45, 14, 46},
	{24, Q, 1588, 963, 661, 407, 664, 30, 11, 24, 16, 25},
	{24, H, 1228, 744, 511, 315, 514, 30, 30, 16, 2, 17},
	{25, L, 3057, 1853, 1273, 784, 1276, 26, 8, 106, 4, 107},
	{25, M, 2395, 1451, 997, 614, 1000, 28, 8, 47, 13, 48},
	{25, Q, 1718, 1041, 715, 440, 718, 30, 7, 24, 22, 25},
	{25, H, 1286, 779, 535, 330, 538, 30, 22, 15, 13, 16},
	{26, L, 3283, 1990, 1367, 842, 1370, 28, 10, 114, 2, 115},
	{26, M, 2544, 1542, 1059, 652, 1062, 28, 19, 46, 4, 47},
	{26, Q, 1804, 1094, 751, 462, 754, 28, 28, 22, 6, 23},
	{26, H, 1425, 864, 593, 365, 596, 30, 33, 16, 4, 17},
	{27, L, 3517, 2132, 1465, 902, 1468, 30, 8, 122, 4, 123},
	{27, M, 2701, 1637, 1125, 692, 1128, 28, 22, 45, 3, 46},
	{27, Q, 1933, 1172, 805, 496, 808, 30, 8, 23, 26, 24},
	{27, H, 1501, 910, 625, 385, 628, 30, 12, 15, 28, 16},
	{28, L, 3669, 2223, 1528, 940, 1531, 30, 3, 117, 10, 118},
	{28, M, 2857, 1732, 1190, 732, 1193, 28, 3, 45, 23, 46},
	{28, Q, 2085, 1263, 868, 534, 871, 30, 4, 24, 31, 25},
	{28, H, 1581, 958, 658, 405, 661, 30, 11, 15, 31, 16},
	{29, L, 3909, 2369, 1628, 1002, 1631, 30, 7, 116, 7, 117},
	{29, M, 3035, 1839, 1264, 778, 1267, 28, 21, 45, 7, 46},
	{29, Q, 2181, 1322, 908, 559, 911, 30, 1, 23, 37, 24},
	{29, H, 1677, 1016, 698, 430, 701, 30, 19, 15, 26, 16},
	{30, L, 4158, 2520, 1732, 1066, 1735, 30, 5, 115, 10, 116},
	{30, M, 3289, 1994, 1370, 843, 1373, 28, 19, 47, 10, 48},
	{30, Q, 2358, 1429, 982, 604, 985, 30, 15, 24, 25, 25},
	{30, H, 1782, 1080, 742, 457, 745, 30, 23, 15, 25, 16},
	{31, L, 4417, 2677, 1840, 1132, 1843, 30, 13, 115, 3, 116},
	{31, M, 3486, 2113, 1452, 894, 1455, 28, 2, 46, 29, 47},
	{31, Q, 2473, 1499, 1030, 634, 1033, 30, 42, 24, 1, 25},
	{31, H, 1897, 1150, 790, 486, 793, 30, 23, 15, 28, 16},
	{32, L, 4686, 2840, 1952, 1201, 1955, 30, 17, 115, 0, 0},
	{32, M, 3693, 2238, 1538, 947, 1541, 28, 10, 46, 23, 47},
	{32, Q, 2670, 1618, 1112, 684, 1115, 30, 10, 24, 35, 25},
	{32, H, 2022, 1226, 842, 518, 845, 30, 19, 15, 35, 16},
	{33, L, 4965, 3009, 2068, 1273, 2071, 30, 17, 115, 1, 116},
	{33, M, 3909, 2369, 1628, 1002, 1631, 28, 14, 46, 21, 47},
	{33, Q, 2805, 1700, 1168, 719, 1171, 30, 29, 24, 19, 25},
	{33, H, 2157, 1307, 898, 553, 901, 30, 11, 15, 46, 16},
	{34, L, 5253, 3183, 2188, 1347, 2191, 30, 13, 115, 6, 116},
	{34, M, 4134, 2506, 1722, 1060, 1725, 28, 14, 46, 23, 47},
	{34, Q, 2949, 1787, 1228, 756, 1231, 30, 44, 24, 7, 25},
	{34, H, 2301, 1394, 958, 590, 961, 30, 59, 16, 1, 17},
	{35, L, 5529, 3351, 2303, 1417, 2306, 30, 12, 121, 7, 122},
	{35, M, 4343, 2632, 1809, 1113, 1812, 28, 12, 47, 26, 48},
	{35, Q, 3081, 1867, 1283, 790, 1286, 30, 39, 24, 14, 25},
	{35, H, 2361, 1431, 983, 605, 986, 30, 22, 15, 41, 16},
	{36, L, 5836, 3537, 2431, 1496, 2434, 30, 6, 121, 14, 122},
	{36, M, 4588, 2780, 1911, 1176, 1914, 28, 6, 47, 34, 48},
	{36, Q, 3244, 1966, 1351, 832, 1354, 30, 46, 24, 10, 25},
	{36, H, 2524, 1530, 1051, 647, 1054, 30, 2, 15, 64, 16},
	{37, L, 6153, 3729, 2563, 1577, 2566, 30, 17, 122, 4, 123},
	{37, M, 4775, 2894, 1989, 1224, 1992, 28, 29, 46, 14, 47},
	{37, Q, 3417, 2071, 1423, 876, 1426, 30, 49, 24, 10, 25},
	{37, H, 2625, 1591, 1093, 673, 1096, 30, 24, 15, 46, 16},
	{38, L, 6479, 3927, 2699, 1661, 2702, 30, 4, 122, 18, 123},
	{38, M, 5039, 3054, 2099, 1292, 2102, 28, 13, 46, 32, 47},
	{38, Q, 3599, 2181, 1499, 923, 1502, 30, 48, 24, 14, 25},
	{38, H, 2735, 1658, 1139, 701, 1142, 30, 42, 15, 32, 16},
	{39, L, 6743, 4087, 2809, 1729, 2812, 30, 20, 117, 4, 118},
	{39, M, 5313, 3220, 2213, 1362, 2216, 28, 40, 47, 7, 48},
	{39, Q, 3791, 2298, 1579, 972, 1582, 30, 43, 24, 22, 25},
	{39, H, 2927, 1774, 1219, 750, 1222, 30, 10, 15, 67, 16},
	{40, L, 7089, 4296, 2953, 1817, 2956, 30, 19, 118, 6, 119},
	{40, M, 5596, 3391, 2331, 1435, 2334, 28, 18, 47, 31, 48},
	{40, Q, 3993, 2420, 1663, 1024, 1666, 30, 34, 24, 34, 25},
	{40, H, 3057, 1852, 1273, 784, 1276, 30, 20, 15, 61, 16},
}

func SmallestQRVersion(byteLength int, el ErrorCorrectionLevel) (*VersionMetaData, error) {
	capsWithErrorLevel := make([]VersionMetaData, 0)

	for _, v := range charCapTable {
		if v.ECL == el {
			capsWithErrorLevel = append(capsWithErrorLevel, v)
		}
	}

	for _, v := range capsWithErrorLevel {
		if v.ByteMode >= byteLength {
			return &v, nil
		}
	}
	return nil, errors.New("Cannot support the amount of data")
}