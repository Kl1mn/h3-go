package h3

type baseCellOrient struct {
	baseCell, ccwRot60 int
}

type baseCellData struct {
	homeFijk     faceIJK
	isPentagon   bool
	cwOffsetPent [2]int
}

var (
	faceIJKBaseCells = [NUM_ICOSA_FACES][3][3][3]baseCellOrient{
		{ // face 0
			{
				// i 0
				{{16, 0}, {18, 0}, {24, 0}}, // j 0
				{{33, 0}, {30, 0}, {32, 3}}, // j 1
				{{49, 1}, {48, 3}, {50, 3}}, // j 2
			},
			{
				// i 1
				{{8, 0}, {5, 5}, {10, 5}},   // j 0
				{{22, 0}, {16, 0}, {18, 0}}, // j 1
				{{41, 1}, {33, 0}, {30, 0}}, // j 2
			},
			{
				// i 2
				{{4, 0}, {0, 5}, {2, 5}},    // j 0
				{{15, 1}, {8, 0}, {5, 5}},   // j 1
				{{31, 1}, {22, 0}, {16, 0}}, // j 2
			},
		},
		{ // face 1
			{
				// i 0
				{{2, 0}, {6, 0}, {14, 0}},   // j 0
				{{10, 0}, {11, 0}, {17, 3}}, // j 1
				{{24, 1}, {23, 3}, {25, 3}}, // j 2
			},
			{
				// i 1
				{{0, 0}, {1, 5}, {9, 5}},    // j 0
				{{5, 0}, {2, 0}, {6, 0}},    // j 1
				{{18, 1}, {10, 0}, {11, 0}}, // j 2
			},
			{
				// i 2
				{{4, 1}, {3, 5}, {7, 5}},  // j 0
				{{8, 1}, {0, 0}, {1, 5}},  // j 1
				{{16, 1}, {5, 0}, {2, 0}}, // j 2
			},
		},
		{ // face 2
			{
				// i 0
				{{7, 0}, {21, 0}, {38, 0}},  // j 0
				{{9, 0}, {19, 0}, {34, 3}},  // j 1
				{{14, 1}, {20, 3}, {36, 3}}, // j 2
			},
			{
				// i 1
				{{3, 0}, {13, 5}, {29, 5}}, // j 0
				{{1, 0}, {7, 0}, {21, 0}},  // j 1
				{{6, 1}, {9, 0}, {19, 0}},  // j 2
			},
			{
				// i 2
				{{4, 2}, {12, 5}, {26, 5}}, // j 0
				{{0, 1}, {3, 0}, {13, 5}},  // j 1
				{{2, 1}, {1, 0}, {7, 0}},   // j 2
			},
		},
		{ // face 3
			{
				// i 0
				{{26, 0}, {42, 0}, {58, 0}}, // j 0
				{{29, 0}, {43, 0}, {62, 3}}, // j 1
				{{38, 1}, {47, 3}, {64, 3}}, // j 2
			},
			{
				// i 1
				{{12, 0}, {28, 5}, {44, 5}}, // j 0
				{{13, 0}, {26, 0}, {42, 0}}, // j 1
				{{21, 1}, {29, 0}, {43, 0}}, // j 2
			},
			{
				// i 2
				{{4, 3}, {15, 5}, {31, 5}}, // j 0
				{{3, 1}, {12, 0}, {28, 5}}, // j 1
				{{7, 1}, {13, 0}, {26, 0}}, // j 2
			},
		},
		{ // face 4
			{
				// i 0
				{{31, 0}, {41, 0}, {49, 0}}, // j 0
				{{44, 0}, {53, 0}, {61, 3}}, // j 1
				{{58, 1}, {65, 3}, {75, 3}}, // j 2
			},
			{
				// i 1
				{{15, 0}, {22, 5}, {33, 5}}, // j 0
				{{28, 0}, {31, 0}, {41, 0}}, // j 1
				{{42, 1}, {44, 0}, {53, 0}}, // j 2
			},
			{
				// i 2
				{{4, 4}, {8, 5}, {16, 5}},   // j 0
				{{12, 1}, {15, 0}, {22, 5}}, // j 1
				{{26, 1}, {28, 0}, {31, 0}}, // j 2
			},
		},
		{ // face 5
			{
				// i 0
				{{50, 0}, {48, 0}, {49, 3}}, // j 0
				{{32, 0}, {30, 3}, {33, 3}}, // j 1
				{{24, 3}, {18, 3}, {16, 3}}, // j 2
			},
			{
				// i 1
				{{70, 0}, {67, 0}, {66, 3}}, // j 0
				{{52, 3}, {50, 0}, {48, 0}}, // j 1
				{{37, 3}, {32, 0}, {30, 3}}, // j 2
			},
			{
				// i 2
				{{83, 0}, {87, 3}, {85, 3}}, // j 0
				{{74, 3}, {70, 0}, {67, 0}}, // j 1
				{{57, 1}, {52, 3}, {50, 0}}, // j 2
			},
		},
		{ // face 6
			{
				// i 0
				{{25, 0}, {23, 0}, {24, 3}}, // j 0
				{{17, 0}, {11, 3}, {10, 3}}, // j 1
				{{14, 3}, {6, 3}, {2, 3}},   // j 2
			},
			{
				// i 1
				{{45, 0}, {39, 0}, {37, 3}}, // j 0
				{{35, 3}, {25, 0}, {23, 0}}, // j 1
				{{27, 3}, {17, 0}, {11, 3}}, // j 2
			},
			{
				// i 2
				{{63, 0}, {59, 3}, {57, 3}}, // j 0
				{{56, 3}, {45, 0}, {39, 0}}, // j 1
				{{46, 3}, {35, 3}, {25, 0}}, // j 2
			},
		},
		{ // face 7
			{
				// i 0
				{{36, 0}, {20, 0}, {14, 3}}, // j 0
				{{34, 0}, {19, 3}, {9, 3}},  // j 1
				{{38, 3}, {21, 3}, {7, 3}},  // j 2
			},
			{
				// i 1
				{{55, 0}, {40, 0}, {27, 3}}, // j 0
				{{54, 3}, {36, 0}, {20, 0}}, // j 1
				{{51, 3}, {34, 0}, {19, 3}}, // j 2
			},
			{
				// i 2
				{{72, 0}, {60, 3}, {46, 3}}, // j 0
				{{73, 3}, {55, 0}, {40, 0}}, // j 1
				{{71, 3}, {54, 3}, {36, 0}}, // j 2
			},
		},
		{ // face 8
			{
				// i 0
				{{64, 0}, {47, 0}, {38, 3}}, // j 0
				{{62, 0}, {43, 3}, {29, 3}}, // j 1
				{{58, 3}, {42, 3}, {26, 3}}, // j 2
			},
			{
				// i 1
				{{84, 0}, {69, 0}, {51, 3}}, // j 0
				{{82, 3}, {64, 0}, {47, 0}}, // j 1
				{{76, 3}, {62, 0}, {43, 3}}, // j 2
			},
			{
				// i 2
				{{97, 0}, {89, 3}, {71, 3}}, // j 0
				{{98, 3}, {84, 0}, {69, 0}}, // j 1
				{{96, 3}, {82, 3}, {64, 0}}, // j 2
			},
		},
		{ // face 9
			{
				// i 0
				{{75, 0}, {65, 0}, {58, 3}}, // j 0
				{{61, 0}, {53, 3}, {44, 3}}, // j 1
				{{49, 3}, {41, 3}, {31, 3}}, // j 2
			},
			{
				// i 1
				{{94, 0}, {86, 0}, {76, 3}}, // j 0
				{{81, 3}, {75, 0}, {65, 0}}, // j 1
				{{66, 3}, {61, 0}, {53, 3}}, // j 2
			},
			{
				// i 2
				{{107, 0}, {104, 3}, {96, 3}}, // j 0
				{{101, 3}, {94, 0}, {86, 0}},  // j 1
				{{85, 3}, {81, 3}, {75, 0}},   // j 2
			},
		},
		{ // face 10
			{
				// i 0
				{{57, 0}, {59, 0}, {63, 3}}, // j 0
				{{74, 0}, {78, 3}, {79, 3}}, // j 1
				{{83, 3}, {92, 3}, {95, 3}}, // j 2
			},
			{
				// i 1
				{{37, 0}, {39, 3}, {45, 3}}, // j 0
				{{52, 0}, {57, 0}, {59, 0}}, // j 1
				{{70, 3}, {74, 0}, {78, 3}}, // j 2
			},
			{
				// i 2
				{{24, 0}, {23, 3}, {25, 3}}, // j 0
				{{32, 3}, {37, 0}, {39, 3}}, // j 1
				{{50, 3}, {52, 0}, {57, 0}}, // j 2
			},
		},
		{ // face 11
			{
				// i 0
				{{46, 0}, {60, 0}, {72, 3}}, // j 0
				{{56, 0}, {68, 3}, {80, 3}}, // j 1
				{{63, 3}, {77, 3}, {90, 3}}, // j 2
			},
			{
				// i 1
				{{27, 0}, {40, 3}, {55, 3}}, // j 0
				{{35, 0}, {46, 0}, {60, 0}}, // j 1
				{{45, 3}, {56, 0}, {68, 3}}, // j 2
			},
			{
				// i 2
				{{14, 0}, {20, 3}, {36, 3}}, // j 0
				{{17, 3}, {27, 0}, {40, 3}}, // j 1
				{{25, 3}, {35, 0}, {46, 0}}, // j 2
			},
		},
		{ // face 12
			{
				// i 0
				{{71, 0}, {89, 0}, {97, 3}},  // j 0
				{{73, 0}, {91, 3}, {103, 3}}, // j 1
				{{72, 3}, {88, 3}, {105, 3}}, // j 2
			},
			{
				// i 1
				{{51, 0}, {69, 3}, {84, 3}}, // j 0
				{{54, 0}, {71, 0}, {89, 0}}, // j 1
				{{55, 3}, {73, 0}, {91, 3}}, // j 2
			},
			{
				// i 2
				{{38, 0}, {47, 3}, {64, 3}}, // j 0
				{{34, 3}, {51, 0}, {69, 3}}, // j 1
				{{36, 3}, {54, 0}, {71, 0}}, // j 2
			},
		},
		{ // face 13
			{
				// i 0
				{{96, 0}, {104, 0}, {107, 3}}, // j 0
				{{98, 0}, {110, 3}, {115, 3}}, // j 1
				{{97, 3}, {111, 3}, {119, 3}}, // j 2
			},
			{
				// i 1
				{{76, 0}, {86, 3}, {94, 3}},  // j 0
				{{82, 0}, {96, 0}, {104, 0}}, // j 1
				{{84, 3}, {98, 0}, {110, 3}}, // j 2
			},
			{
				// i 2
				{{58, 0}, {65, 3}, {75, 3}}, // j 0
				{{62, 3}, {76, 0}, {86, 3}}, // j 1
				{{64, 3}, {82, 0}, {96, 0}}, // j 2
			},
		},
		{ // face 14
			{
				// i 0
				{{85, 0}, {87, 0}, {83, 3}},    // j 0
				{{101, 0}, {102, 3}, {100, 3}}, // j 1
				{{107, 3}, {112, 3}, {114, 3}}, // j 2
			},
			{
				// i 1
				{{66, 0}, {67, 3}, {70, 3}},   // j 0
				{{81, 0}, {85, 0}, {87, 0}},   // j 1
				{{94, 3}, {101, 0}, {102, 3}}, // j 2
			},
			{
				// i 2
				{{49, 0}, {48, 3}, {50, 3}}, // j 0
				{{61, 3}, {66, 0}, {67, 3}}, // j 1
				{{75, 3}, {81, 0}, {85, 0}}, // j 2
			},
		},
		{ // face 15
			{
				// i 0
				{{95, 0}, {92, 0}, {83, 0}}, // j 0
				{{79, 0}, {78, 0}, {74, 3}}, // j 1
				{{63, 1}, {59, 3}, {57, 3}}, // j 2
			},
			{
				// i 1
				{{109, 0}, {108, 0}, {100, 5}}, // j 0
				{{93, 1}, {95, 0}, {92, 0}},    // j 1
				{{77, 1}, {79, 0}, {78, 0}},    // j 2
			},
			{
				// i 2
				{{117, 4}, {118, 5}, {114, 5}}, // j 0
				{{106, 1}, {109, 0}, {108, 0}}, // j 1
				{{90, 1}, {93, 1}, {95, 0}},    // j 2
			},
		},
		{ // face 16
			{
				// i 0
				{{90, 0}, {77, 0}, {63, 0}}, // j 0
				{{80, 0}, {68, 0}, {56, 3}}, // j 1
				{{72, 1}, {60, 3}, {46, 3}}, // j 2
			},
			{
				// i 1
				{{106, 0}, {93, 0}, {79, 5}}, // j 0
				{{99, 1}, {90, 0}, {77, 0}},  // j 1
				{{88, 1}, {80, 0}, {68, 0}},  // j 2
			},
			{
				// i 2
				{{117, 3}, {109, 5}, {95, 5}}, // j 0
				{{113, 1}, {106, 0}, {93, 0}}, // j 1
				{{105, 1}, {99, 1}, {90, 0}},  // j 2
			},
		},
		{ // face 17
			{
				// i 0
				{{105, 0}, {88, 0}, {72, 0}}, // j 0
				{{103, 0}, {91, 0}, {73, 3}}, // j 1
				{{97, 1}, {89, 3}, {71, 3}},  // j 2
			},
			{
				// i 1
				{{113, 0}, {99, 0}, {80, 5}},  // j 0
				{{116, 1}, {105, 0}, {88, 0}}, // j 1
				{{111, 1}, {103, 0}, {91, 0}}, // j 2
			},
			{
				// i 2
				{{117, 2}, {106, 5}, {90, 5}},  // j 0
				{{121, 1}, {113, 0}, {99, 0}},  // j 1
				{{119, 1}, {116, 1}, {105, 0}}, // j 2
			},
		},
		{ // face 18
			{
				// i 0
				{{119, 0}, {111, 0}, {97, 0}}, // j 0
				{{115, 0}, {110, 0}, {98, 3}}, // j 1
				{{107, 1}, {104, 3}, {96, 3}}, // j 2
			},
			{
				// i 1
				{{121, 0}, {116, 0}, {103, 5}}, // j 0
				{{120, 1}, {119, 0}, {111, 0}}, // j 1
				{{112, 1}, {115, 0}, {110, 0}}, // j 2
			},
			{
				// i 2
				{{117, 1}, {113, 5}, {105, 5}}, // j 0
				{{118, 1}, {121, 0}, {116, 0}}, // j 1
				{{114, 1}, {120, 1}, {119, 0}}, // j 2
			},
		},
		{ // face 19
			{
				// i 0
				{{114, 0}, {112, 0}, {107, 0}}, // j 0
				{{100, 0}, {102, 0}, {101, 3}}, // j 1
				{{83, 1}, {87, 3}, {85, 3}},    // j 2
			},
			{
				// i 1
				{{118, 0}, {120, 0}, {115, 5}}, // j 0
				{{108, 1}, {114, 0}, {112, 0}}, // j 1
				{{92, 1}, {100, 0}, {102, 0}},  // j 2
			},
			{
				// i 2
				{{117, 0}, {121, 5}, {119, 5}}, // j 0
				{{109, 1}, {118, 0}, {120, 0}}, // j 1
				{{95, 1}, {108, 1}, {114, 0}},  // j 2
			},
		},
	}

	baseCellsData = [NUM_BASE_CELLS]baseCellData{

		{
			homeFijk:     faceIJK{face: 1, coord: coordIJK{1, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 0

		{
			homeFijk:     faceIJK{face: 2, coord: coordIJK{1, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 1
		{
			homeFijk:     faceIJK{face: 1, coord: coordIJK{0, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 2
		{
			homeFijk:     faceIJK{face: 2, coord: coordIJK{1, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 3
		{
			homeFijk:     faceIJK{face: 0, coord: coordIJK{2, 0, 0}},
			isPentagon:   true,
			cwOffsetPent: [2]int{-1, -1},
		}, // base cell 4
		{
			homeFijk:     faceIJK{face: 1, coord: coordIJK{1, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 5
		{
			homeFijk:     faceIJK{face: 1, coord: coordIJK{0, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 6
		{
			homeFijk:     faceIJK{face: 2, coord: coordIJK{0, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 7
		{
			homeFijk:     faceIJK{face: 0, coord: coordIJK{1, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 8
		{
			homeFijk:     faceIJK{face: 2, coord: coordIJK{0, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 9
		{
			homeFijk:     faceIJK{face: 1, coord: coordIJK{0, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 10
		{
			homeFijk:     faceIJK{face: 1, coord: coordIJK{0, 1, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 11
		{
			homeFijk:     faceIJK{face: 3, coord: coordIJK{1, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 12
		{
			homeFijk:     faceIJK{face: 3, coord: coordIJK{1, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 13
		{
			homeFijk:     faceIJK{face: 11, coord: coordIJK{2, 0, 0}},
			isPentagon:   true,
			cwOffsetPent: [2]int{2, 6},
		}, // base cell 14
		{
			homeFijk:     faceIJK{face: 4, coord: coordIJK{1, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 15
		{
			homeFijk:     faceIJK{face: 0, coord: coordIJK{0, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 16
		{
			homeFijk:     faceIJK{face: 6, coord: coordIJK{0, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 17
		{
			homeFijk:     faceIJK{face: 0, coord: coordIJK{0, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 18
		{
			homeFijk:     faceIJK{face: 2, coord: coordIJK{0, 1, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 19
		{
			homeFijk:     faceIJK{face: 7, coord: coordIJK{0, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 20
		{
			homeFijk:     faceIJK{face: 2, coord: coordIJK{0, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 21
		{
			homeFijk:     faceIJK{face: 0, coord: coordIJK{1, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 22
		{
			homeFijk:     faceIJK{face: 6, coord: coordIJK{0, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 23
		{
			homeFijk:     faceIJK{face: 10, coord: coordIJK{2, 0, 0}},
			isPentagon:   true,
			cwOffsetPent: [2]int{1, 5},
		}, // base cell 24
		{
			homeFijk:     faceIJK{face: 6, coord: coordIJK{0, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 25
		{
			homeFijk:     faceIJK{face: 3, coord: coordIJK{0, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 26
		{
			homeFijk:     faceIJK{face: 11, coord: coordIJK{1, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 27
		{
			homeFijk:     faceIJK{face: 4, coord: coordIJK{1, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 28
		{
			homeFijk:     faceIJK{face: 3, coord: coordIJK{0, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 29
		{
			homeFijk:     faceIJK{face: 0, coord: coordIJK{0, 1, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 30
		{
			homeFijk:     faceIJK{face: 4, coord: coordIJK{0, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 31
		{
			homeFijk:     faceIJK{face: 5, coord: coordIJK{0, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 32
		{
			homeFijk:     faceIJK{face: 0, coord: coordIJK{0, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 33
		{
			homeFijk:     faceIJK{face: 7, coord: coordIJK{0, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 34
		{
			homeFijk:     faceIJK{face: 11, coord: coordIJK{1, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 35
		{
			homeFijk:     faceIJK{face: 7, coord: coordIJK{0, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 36
		{
			homeFijk:     faceIJK{face: 10, coord: coordIJK{1, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 37
		{
			homeFijk:     faceIJK{face: 12, coord: coordIJK{2, 0, 0}},
			isPentagon:   true,
			cwOffsetPent: [2]int{3, 7},
		}, // base cell 38
		{
			homeFijk:     faceIJK{face: 6, coord: coordIJK{1, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 39
		{
			homeFijk:     faceIJK{face: 7, coord: coordIJK{1, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 40
		{
			homeFijk:     faceIJK{face: 4, coord: coordIJK{0, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 41
		{
			homeFijk:     faceIJK{face: 3, coord: coordIJK{0, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 42
		{
			homeFijk:     faceIJK{face: 3, coord: coordIJK{0, 1, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 43
		{
			homeFijk:     faceIJK{face: 4, coord: coordIJK{0, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 44
		{
			homeFijk:     faceIJK{face: 6, coord: coordIJK{1, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 45
		{
			homeFijk:     faceIJK{face: 11, coord: coordIJK{0, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 46
		{
			homeFijk:     faceIJK{face: 8, coord: coordIJK{0, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 47
		{
			homeFijk:     faceIJK{face: 5, coord: coordIJK{0, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 48
		{
			homeFijk:     faceIJK{face: 14, coord: coordIJK{2, 0, 0}},
			isPentagon:   true,
			cwOffsetPent: [2]int{0, 9},
		}, // base cell 49
		{
			homeFijk:     faceIJK{face: 5, coord: coordIJK{0, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 50
		{
			homeFijk:     faceIJK{face: 12, coord: coordIJK{1, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 51
		{
			homeFijk:     faceIJK{face: 10, coord: coordIJK{1, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 52
		{
			homeFijk:     faceIJK{face: 4, coord: coordIJK{0, 1, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 53
		{
			homeFijk:     faceIJK{face: 12, coord: coordIJK{1, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 54
		{
			homeFijk:     faceIJK{face: 7, coord: coordIJK{1, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 55
		{
			homeFijk:     faceIJK{face: 11, coord: coordIJK{0, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 56
		{
			homeFijk:     faceIJK{face: 10, coord: coordIJK{0, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 57
		{
			homeFijk:     faceIJK{face: 13, coord: coordIJK{2, 0, 0}},
			isPentagon:   true,
			cwOffsetPent: [2]int{4, 8},
		}, // base cell 58
		{
			homeFijk:     faceIJK{face: 10, coord: coordIJK{0, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 59
		{
			homeFijk:     faceIJK{face: 11, coord: coordIJK{0, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 60
		{
			homeFijk:     faceIJK{face: 9, coord: coordIJK{0, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 61
		{
			homeFijk:     faceIJK{face: 8, coord: coordIJK{0, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 62
		{
			homeFijk:     faceIJK{face: 6, coord: coordIJK{2, 0, 0}},
			isPentagon:   true,
			cwOffsetPent: [2]int{11, 15},
		}, // base cell 63
		{
			homeFijk:     faceIJK{face: 8, coord: coordIJK{0, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 64
		{
			homeFijk:     faceIJK{face: 9, coord: coordIJK{0, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 65
		{
			homeFijk:     faceIJK{face: 14, coord: coordIJK{1, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 66
		{
			homeFijk:     faceIJK{face: 5, coord: coordIJK{1, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 67
		{
			homeFijk:     faceIJK{face: 16, coord: coordIJK{0, 1, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 68
		{
			homeFijk:     faceIJK{face: 8, coord: coordIJK{1, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 69
		{
			homeFijk:     faceIJK{face: 5, coord: coordIJK{1, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 70
		{
			homeFijk:     faceIJK{face: 12, coord: coordIJK{0, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 71
		{
			homeFijk:     faceIJK{face: 7, coord: coordIJK{2, 0, 0}},
			isPentagon:   true,
			cwOffsetPent: [2]int{12, 16},
		}, // base cell 72
		{
			homeFijk:     faceIJK{face: 12, coord: coordIJK{0, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 73
		{
			homeFijk:     faceIJK{face: 10, coord: coordIJK{0, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 74
		{
			homeFijk:     faceIJK{face: 9, coord: coordIJK{0, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 75
		{
			homeFijk:     faceIJK{face: 13, coord: coordIJK{1, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 76
		{
			homeFijk:     faceIJK{face: 16, coord: coordIJK{0, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 77
		{
			homeFijk:     faceIJK{face: 15, coord: coordIJK{0, 1, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 78
		{
			homeFijk:     faceIJK{face: 15, coord: coordIJK{0, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 79
		{
			homeFijk:     faceIJK{face: 16, coord: coordIJK{0, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 80
		{
			homeFijk:     faceIJK{face: 14, coord: coordIJK{1, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 81
		{
			homeFijk:     faceIJK{face: 13, coord: coordIJK{1, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 82
		{
			homeFijk:     faceIJK{face: 5, coord: coordIJK{2, 0, 0}},
			isPentagon:   true,
			cwOffsetPent: [2]int{10, 19},
		}, // base cell 83
		{
			homeFijk:     faceIJK{face: 8, coord: coordIJK{1, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 84
		{
			homeFijk:     faceIJK{face: 14, coord: coordIJK{0, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 85
		{
			homeFijk:     faceIJK{face: 9, coord: coordIJK{1, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 86
		{
			homeFijk:     faceIJK{face: 14, coord: coordIJK{0, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 87
		{
			homeFijk:     faceIJK{face: 17, coord: coordIJK{0, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 88
		{
			homeFijk:     faceIJK{face: 12, coord: coordIJK{0, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 89
		{
			homeFijk:     faceIJK{face: 16, coord: coordIJK{0, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 90
		{
			homeFijk:     faceIJK{face: 17, coord: coordIJK{0, 1, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 91
		{
			homeFijk:     faceIJK{face: 15, coord: coordIJK{0, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 92
		{
			homeFijk:     faceIJK{face: 16, coord: coordIJK{1, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 93
		{
			homeFijk:     faceIJK{face: 9, coord: coordIJK{1, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 94
		{
			homeFijk:     faceIJK{face: 15, coord: coordIJK{0, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 95
		{
			homeFijk:     faceIJK{face: 13, coord: coordIJK{0, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 96
		{
			homeFijk:     faceIJK{face: 8, coord: coordIJK{2, 0, 0}},
			isPentagon:   true,
			cwOffsetPent: [2]int{13, 17},
		}, // base cell 97
		{
			homeFijk:     faceIJK{face: 13, coord: coordIJK{0, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 98
		{
			homeFijk:     faceIJK{face: 17, coord: coordIJK{1, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 99
		{
			homeFijk:     faceIJK{face: 19, coord: coordIJK{0, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 100
		{
			homeFijk:     faceIJK{face: 14, coord: coordIJK{0, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 101
		{
			homeFijk:     faceIJK{face: 19, coord: coordIJK{0, 1, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 102
		{
			homeFijk:     faceIJK{face: 17, coord: coordIJK{0, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 103
		{
			homeFijk:   faceIJK{face: 13, coord: coordIJK{0, 0, 1}},
			isPentagon: false, cwOffsetPent: [2]int{0, 0},
		}, // base cell 104
		{
			homeFijk:     faceIJK{face: 17, coord: coordIJK{0, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 105
		{
			homeFijk:     faceIJK{face: 16, coord: coordIJK{1, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 106
		{
			homeFijk:     faceIJK{face: 9, coord: coordIJK{2, 0, 0}},
			isPentagon:   true,
			cwOffsetPent: [2]int{14, 18},
		}, // base cell 107
		{
			homeFijk:     faceIJK{face: 15, coord: coordIJK{1, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 108
		{
			homeFijk:     faceIJK{face: 15, coord: coordIJK{1, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 109
		{
			homeFijk:     faceIJK{face: 18, coord: coordIJK{0, 1, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 110
		{
			homeFijk:     faceIJK{face: 18, coord: coordIJK{0, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 111
		{
			homeFijk:     faceIJK{face: 19, coord: coordIJK{0, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 112
		{
			homeFijk:     faceIJK{face: 17, coord: coordIJK{1, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 113
		{
			homeFijk:     faceIJK{face: 19, coord: coordIJK{0, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 114
		{
			homeFijk:     faceIJK{face: 18, coord: coordIJK{0, 1, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 115
		{
			homeFijk:     faceIJK{face: 18, coord: coordIJK{1, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 116
		{
			homeFijk:     faceIJK{face: 19, coord: coordIJK{2, 0, 0}},
			isPentagon:   true,
			cwOffsetPent: [2]int{-1, -1},
		}, // base cell 117
		{
			homeFijk:     faceIJK{face: 19, coord: coordIJK{1, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 118
		{
			homeFijk:     faceIJK{face: 18, coord: coordIJK{0, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 119
		{
			homeFijk:     faceIJK{face: 19, coord: coordIJK{1, 0, 1}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 120
		{
			homeFijk:     faceIJK{face: 18, coord: coordIJK{1, 0, 0}},
			isPentagon:   false,
			cwOffsetPent: [2]int{0, 0},
		}, // base cell 121
	}
)

func isBaseCellPentagon(baseCell int) bool {
	return baseCellsData[baseCell].isPentagon
}
