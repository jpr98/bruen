// Code generated from Bruen.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Bruen

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 55, 606,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9,
	60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65, 9, 65,
	4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9, 70, 4,
	71, 9, 71, 4, 72, 9, 72, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 149, 10, 2, 12,
	2, 14, 2, 152, 11, 2, 3, 2, 3, 2, 3, 3, 7, 3, 157, 10, 3, 12, 3, 14, 3,
	160, 11, 3, 3, 3, 7, 3, 163, 10, 3, 12, 3, 14, 3, 166, 11, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7,
	5, 181, 10, 5, 12, 5, 14, 5, 184, 11, 5, 3, 5, 3, 5, 3, 6, 3, 6, 7, 6,
	190, 10, 6, 12, 6, 14, 6, 193, 11, 6, 3, 7, 3, 7, 5, 7, 197, 10, 7, 5,
	7, 199, 10, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 208, 10,
	8, 12, 8, 14, 8, 211, 11, 8, 3, 8, 7, 8, 214, 10, 8, 12, 8, 14, 8, 217,
	11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 5, 9, 224, 10, 9, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 13, 3, 13, 5, 13, 254, 10, 13, 5, 13, 256, 10, 13,
	3, 13, 3, 13, 3, 13, 5, 13, 261, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 15, 3, 15, 5, 15, 271, 10, 15, 3, 15, 5, 15, 274, 10, 15,
	3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 5, 17, 282, 10, 17, 3, 18, 3,
	18, 3, 18, 5, 18, 287, 10, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 19, 5, 19, 296, 10, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 302, 10,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 308, 10, 19, 3, 20, 3, 20, 3, 20,
	3, 20, 5, 20, 314, 10, 20, 3, 20, 3, 20, 3, 20, 5, 20, 319, 10, 20, 3,
	20, 3, 20, 3, 21, 3, 21, 3, 21, 7, 21, 326, 10, 21, 12, 21, 14, 21, 329,
	11, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 7, 23, 337, 10, 23, 12,
	23, 14, 23, 340, 11, 23, 3, 23, 7, 23, 343, 10, 23, 12, 23, 14, 23, 346,
	11, 23, 3, 23, 3, 23, 3, 24, 3, 24, 5, 24, 352, 10, 24, 3, 24, 3, 24, 3,
	25, 3, 25, 7, 25, 358, 10, 25, 12, 25, 14, 25, 361, 11, 25, 3, 25, 3, 25,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 372, 10, 26, 3,
	27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 382, 10, 27,
	3, 28, 3, 28, 3, 28, 5, 28, 387, 10, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3,
	29, 3, 29, 3, 29, 5, 29, 396, 10, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30,
	3, 31, 3, 31, 3, 31, 5, 31, 406, 10, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 5, 32, 413, 10, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 34, 3, 34, 3, 34, 5, 34, 425, 10, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 36, 3, 36, 3, 36, 5, 36, 436, 10, 36, 3, 36, 3, 36, 3, 37,
	3, 37, 3, 37, 5, 37, 443, 10, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3,
	38, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 5, 40, 457, 10, 40, 3, 41,
	3, 41, 3, 41, 3, 41, 5, 41, 463, 10, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3,
	43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3,
	49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 7, 51, 499,
	10, 51, 12, 51, 14, 51, 502, 11, 51, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53,
	7, 53, 509, 10, 53, 12, 53, 14, 53, 512, 11, 53, 3, 54, 3, 54, 3, 54, 3,
	55, 3, 55, 5, 55, 519, 10, 55, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 7, 57,
	526, 10, 57, 12, 57, 14, 57, 529, 11, 57, 3, 58, 3, 58, 3, 58, 3, 59, 3,
	59, 7, 59, 536, 10, 59, 12, 59, 14, 59, 539, 11, 59, 3, 60, 3, 60, 3, 60,
	3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 5, 61, 549, 10, 61, 3, 61, 5, 61, 552,
	10, 61, 3, 62, 3, 62, 3, 62, 3, 62, 3, 63, 3, 63, 3, 63, 3, 63, 3, 63,
	5, 63, 563, 10, 63, 3, 64, 3, 64, 3, 65, 3, 65, 3, 66, 3, 66, 3, 67, 3,
	67, 3, 68, 3, 68, 7, 68, 575, 10, 68, 12, 68, 14, 68, 578, 11, 68, 3, 68,
	3, 68, 3, 69, 3, 69, 3, 69, 3, 69, 3, 69, 3, 70, 3, 70, 7, 70, 589, 10,
	70, 12, 70, 14, 70, 592, 11, 70, 3, 70, 7, 70, 595, 10, 70, 12, 70, 14,
	70, 598, 11, 70, 3, 70, 3, 70, 3, 71, 3, 71, 3, 72, 3, 72, 3, 72, 3, 576,
	2, 73, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
	72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104,
	106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134,
	136, 138, 140, 142, 2, 6, 3, 2, 26, 27, 3, 2, 24, 25, 3, 2, 49, 52, 3,
	2, 20, 23, 2, 605, 2, 144, 3, 2, 2, 2, 4, 158, 3, 2, 2, 2, 6, 170, 3, 2,
	2, 2, 8, 175, 3, 2, 2, 2, 10, 187, 3, 2, 2, 2, 12, 198, 3, 2, 2, 2, 14,
	202, 3, 2, 2, 2, 16, 223, 3, 2, 2, 2, 18, 225, 3, 2, 2, 2, 20, 231, 3,
	2, 2, 2, 22, 240, 3, 2, 2, 2, 24, 255, 3, 2, 2, 2, 26, 262, 3, 2, 2, 2,
	28, 270, 3, 2, 2, 2, 30, 275, 3, 2, 2, 2, 32, 278, 3, 2, 2, 2, 34, 283,
	3, 2, 2, 2, 36, 292, 3, 2, 2, 2, 38, 309, 3, 2, 2, 2, 40, 322, 3, 2, 2,
	2, 42, 330, 3, 2, 2, 2, 44, 334, 3, 2, 2, 2, 46, 349, 3, 2, 2, 2, 48, 355,
	3, 2, 2, 2, 50, 371, 3, 2, 2, 2, 52, 381, 3, 2, 2, 2, 54, 386, 3, 2, 2,
	2, 56, 392, 3, 2, 2, 2, 58, 399, 3, 2, 2, 2, 60, 405, 3, 2, 2, 2, 62, 407,
	3, 2, 2, 2, 64, 416, 3, 2, 2, 2, 66, 424, 3, 2, 2, 2, 68, 426, 3, 2, 2,
	2, 70, 435, 3, 2, 2, 2, 72, 442, 3, 2, 2, 2, 74, 444, 3, 2, 2, 2, 76, 450,
	3, 2, 2, 2, 78, 456, 3, 2, 2, 2, 80, 458, 3, 2, 2, 2, 82, 464, 3, 2, 2,
	2, 84, 468, 3, 2, 2, 2, 86, 470, 3, 2, 2, 2, 88, 473, 3, 2, 2, 2, 90, 479,
	3, 2, 2, 2, 92, 483, 3, 2, 2, 2, 94, 485, 3, 2, 2, 2, 96, 489, 3, 2, 2,
	2, 98, 493, 3, 2, 2, 2, 100, 496, 3, 2, 2, 2, 102, 503, 3, 2, 2, 2, 104,
	506, 3, 2, 2, 2, 106, 513, 3, 2, 2, 2, 108, 516, 3, 2, 2, 2, 110, 520,
	3, 2, 2, 2, 112, 523, 3, 2, 2, 2, 114, 530, 3, 2, 2, 2, 116, 533, 3, 2,
	2, 2, 118, 540, 3, 2, 2, 2, 120, 551, 3, 2, 2, 2, 122, 553, 3, 2, 2, 2,
	124, 562, 3, 2, 2, 2, 126, 564, 3, 2, 2, 2, 128, 566, 3, 2, 2, 2, 130,
	568, 3, 2, 2, 2, 132, 570, 3, 2, 2, 2, 134, 572, 3, 2, 2, 2, 136, 581,
	3, 2, 2, 2, 138, 586, 3, 2, 2, 2, 140, 601, 3, 2, 2, 2, 142, 603, 3, 2,
	2, 2, 144, 145, 7, 54, 2, 2, 145, 146, 7, 55, 2, 2, 146, 150, 7, 28, 2,
	2, 147, 149, 5, 16, 9, 2, 148, 147, 3, 2, 2, 2, 149, 152, 3, 2, 2, 2, 150,
	148, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 153, 3, 2, 2, 2, 152, 150,
	3, 2, 2, 2, 153, 154, 5, 4, 3, 2, 154, 3, 3, 2, 2, 2, 155, 157, 5, 6, 4,
	2, 156, 155, 3, 2, 2, 2, 157, 160, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 158,
	159, 3, 2, 2, 2, 159, 164, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 161, 163,
	5, 38, 20, 2, 162, 161, 3, 2, 2, 2, 163, 166, 3, 2, 2, 2, 164, 162, 3,
	2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 167, 3, 2, 2, 2, 166, 164, 3, 2, 2,
	2, 167, 168, 5, 136, 69, 2, 168, 169, 7, 2, 2, 3, 169, 5, 3, 2, 2, 2, 170,
	171, 7, 37, 2, 2, 171, 172, 7, 55, 2, 2, 172, 173, 5, 8, 5, 2, 173, 174,
	7, 28, 2, 2, 174, 7, 3, 2, 2, 2, 175, 176, 7, 3, 2, 2, 176, 177, 5, 10,
	6, 2, 177, 178, 5, 14, 8, 2, 178, 182, 7, 39, 2, 2, 179, 181, 5, 12, 7,
	2, 180, 179, 3, 2, 2, 2, 181, 184, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 182,
	183, 3, 2, 2, 2, 183, 185, 3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 185, 186,
	7, 4, 2, 2, 186, 9, 3, 2, 2, 2, 187, 191, 7, 38, 2, 2, 188, 190, 5, 24,
	13, 2, 189, 188, 3, 2, 2, 2, 190, 193, 3, 2, 2, 2, 191, 189, 3, 2, 2, 2,
	191, 192, 3, 2, 2, 2, 192, 11, 3, 2, 2, 2, 193, 191, 3, 2, 2, 2, 194, 199,
	7, 41, 2, 2, 195, 197, 7, 42, 2, 2, 196, 195, 3, 2, 2, 2, 196, 197, 3,
	2, 2, 2, 197, 199, 3, 2, 2, 2, 198, 194, 3, 2, 2, 2, 198, 196, 3, 2, 2,
	2, 199, 200, 3, 2, 2, 2, 200, 201, 5, 38, 20, 2, 201, 13, 3, 2, 2, 2, 202,
	203, 7, 40, 2, 2, 203, 204, 7, 30, 2, 2, 204, 205, 7, 31, 2, 2, 205, 209,
	7, 3, 2, 2, 206, 208, 5, 16, 9, 2, 207, 206, 3, 2, 2, 2, 208, 211, 3, 2,
	2, 2, 209, 207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 215, 3, 2, 2, 2,
	211, 209, 3, 2, 2, 2, 212, 214, 5, 50, 26, 2, 213, 212, 3, 2, 2, 2, 214,
	217, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 218,
	3, 2, 2, 2, 217, 215, 3, 2, 2, 2, 218, 219, 7, 4, 2, 2, 219, 15, 3, 2,
	2, 2, 220, 224, 5, 18, 10, 2, 221, 224, 5, 20, 11, 2, 222, 224, 5, 22,
	12, 2, 223, 220, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 223, 222, 3, 2, 2, 2,
	224, 17, 3, 2, 2, 2, 225, 226, 7, 48, 2, 2, 226, 227, 7, 55, 2, 2, 227,
	228, 7, 5, 2, 2, 228, 229, 5, 28, 15, 2, 229, 230, 7, 28, 2, 2, 230, 19,
	3, 2, 2, 2, 231, 232, 7, 48, 2, 2, 232, 233, 7, 55, 2, 2, 233, 234, 7,
	5, 2, 2, 234, 235, 7, 6, 2, 2, 235, 236, 7, 13, 2, 2, 236, 237, 7, 7, 2,
	2, 237, 238, 5, 140, 71, 2, 238, 239, 7, 28, 2, 2, 239, 21, 3, 2, 2, 2,
	240, 241, 7, 48, 2, 2, 241, 242, 7, 55, 2, 2, 242, 243, 7, 5, 2, 2, 243,
	244, 7, 6, 2, 2, 244, 245, 7, 13, 2, 2, 245, 246, 7, 8, 2, 2, 246, 247,
	7, 13, 2, 2, 247, 248, 7, 7, 2, 2, 248, 249, 5, 140, 71, 2, 249, 250, 7,
	28, 2, 2, 250, 23, 3, 2, 2, 2, 251, 256, 7, 41, 2, 2, 252, 254, 7, 42,
	2, 2, 253, 252, 3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254, 256, 3, 2, 2, 2,
	255, 251, 3, 2, 2, 2, 255, 253, 3, 2, 2, 2, 256, 260, 3, 2, 2, 2, 257,
	261, 5, 26, 14, 2, 258, 261, 5, 20, 11, 2, 259, 261, 5, 22, 12, 2, 260,
	257, 3, 2, 2, 2, 260, 258, 3, 2, 2, 2, 260, 259, 3, 2, 2, 2, 261, 25, 3,
	2, 2, 2, 262, 263, 7, 48, 2, 2, 263, 264, 7, 55, 2, 2, 264, 265, 7, 5,
	2, 2, 265, 266, 5, 140, 71, 2, 266, 267, 7, 28, 2, 2, 267, 27, 3, 2, 2,
	2, 268, 271, 5, 140, 71, 2, 269, 271, 7, 55, 2, 2, 270, 268, 3, 2, 2, 2,
	270, 269, 3, 2, 2, 2, 271, 273, 3, 2, 2, 2, 272, 274, 5, 30, 16, 2, 273,
	272, 3, 2, 2, 2, 273, 274, 3, 2, 2, 2, 274, 29, 3, 2, 2, 2, 275, 276, 7,
	29, 2, 2, 276, 277, 5, 100, 51, 2, 277, 31, 3, 2, 2, 2, 278, 281, 7, 55,
	2, 2, 279, 280, 7, 9, 2, 2, 280, 282, 7, 55, 2, 2, 281, 279, 3, 2, 2, 2,
	281, 282, 3, 2, 2, 2, 282, 33, 3, 2, 2, 2, 283, 286, 7, 55, 2, 2, 284,
	285, 7, 9, 2, 2, 285, 287, 7, 55, 2, 2, 286, 284, 3, 2, 2, 2, 286, 287,
	3, 2, 2, 2, 287, 288, 3, 2, 2, 2, 288, 289, 7, 6, 2, 2, 289, 290, 5, 100,
	51, 2, 290, 291, 7, 7, 2, 2, 291, 35, 3, 2, 2, 2, 292, 295, 7, 55, 2, 2,
	293, 294, 7, 9, 2, 2, 294, 296, 7, 55, 2, 2, 295, 293, 3, 2, 2, 2, 295,
	296, 3, 2, 2, 2, 296, 301, 3, 2, 2, 2, 297, 298, 7, 6, 2, 2, 298, 299,
	5, 100, 51, 2, 299, 300, 7, 7, 2, 2, 300, 302, 3, 2, 2, 2, 301, 297, 3,
	2, 2, 2, 301, 302, 3, 2, 2, 2, 302, 307, 3, 2, 2, 2, 303, 304, 7, 6, 2,
	2, 304, 305, 5, 100, 51, 2, 305, 306, 7, 7, 2, 2, 306, 308, 3, 2, 2, 2,
	307, 303, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 37, 3, 2, 2, 2, 309, 310,
	7, 45, 2, 2, 310, 311, 7, 55, 2, 2, 311, 313, 7, 30, 2, 2, 312, 314, 5,
	40, 21, 2, 313, 312, 3, 2, 2, 2, 313, 314, 3, 2, 2, 2, 314, 315, 3, 2,
	2, 2, 315, 318, 7, 31, 2, 2, 316, 319, 5, 140, 71, 2, 317, 319, 7, 53,
	2, 2, 318, 316, 3, 2, 2, 2, 318, 317, 3, 2, 2, 2, 319, 320, 3, 2, 2, 2,
	320, 321, 5, 44, 23, 2, 321, 39, 3, 2, 2, 2, 322, 327, 5, 42, 22, 2, 323,
	324, 7, 10, 2, 2, 324, 326, 5, 42, 22, 2, 325, 323, 3, 2, 2, 2, 326, 329,
	3, 2, 2, 2, 327, 325, 3, 2, 2, 2, 327, 328, 3, 2, 2, 2, 328, 41, 3, 2,
	2, 2, 329, 327, 3, 2, 2, 2, 330, 331, 7, 55, 2, 2, 331, 332, 7, 5, 2, 2,
	332, 333, 5, 140, 71, 2, 333, 43, 3, 2, 2, 2, 334, 338, 7, 3, 2, 2, 335,
	337, 5, 16, 9, 2, 336, 335, 3, 2, 2, 2, 337, 340, 3, 2, 2, 2, 338, 336,
	3, 2, 2, 2, 338, 339, 3, 2, 2, 2, 339, 344, 3, 2, 2, 2, 340, 338, 3, 2,
	2, 2, 341, 343, 5, 52, 27, 2, 342, 341, 3, 2, 2, 2, 343, 346, 3, 2, 2,
	2, 344, 342, 3, 2, 2, 2, 344, 345, 3, 2, 2, 2, 345, 347, 3, 2, 2, 2, 346,
	344, 3, 2, 2, 2, 347, 348, 7, 4, 2, 2, 348, 45, 3, 2, 2, 2, 349, 351, 7,
	46, 2, 2, 350, 352, 5, 100, 51, 2, 351, 350, 3, 2, 2, 2, 351, 352, 3, 2,
	2, 2, 352, 353, 3, 2, 2, 2, 353, 354, 7, 28, 2, 2, 354, 47, 3, 2, 2, 2,
	355, 359, 7, 3, 2, 2, 356, 358, 5, 52, 27, 2, 357, 356, 3, 2, 2, 2, 358,
	361, 3, 2, 2, 2, 359, 357, 3, 2, 2, 2, 359, 360, 3, 2, 2, 2, 360, 362,
	3, 2, 2, 2, 361, 359, 3, 2, 2, 2, 362, 363, 7, 4, 2, 2, 363, 49, 3, 2,
	2, 2, 364, 372, 5, 54, 28, 2, 365, 372, 5, 68, 35, 2, 366, 372, 5, 74,
	38, 2, 367, 372, 5, 80, 41, 2, 368, 372, 5, 88, 45, 2, 369, 372, 5, 94,
	48, 2, 370, 372, 5, 98, 50, 2, 371, 364, 3, 2, 2, 2, 371, 365, 3, 2, 2,
	2, 371, 366, 3, 2, 2, 2, 371, 367, 3, 2, 2, 2, 371, 368, 3, 2, 2, 2, 371,
	369, 3, 2, 2, 2, 371, 370, 3, 2, 2, 2, 372, 51, 3, 2, 2, 2, 373, 382, 5,
	54, 28, 2, 374, 382, 5, 68, 35, 2, 375, 382, 5, 74, 38, 2, 376, 382, 5,
	80, 41, 2, 377, 382, 5, 88, 45, 2, 378, 382, 5, 94, 48, 2, 379, 382, 5,
	98, 50, 2, 380, 382, 5, 46, 24, 2, 381, 373, 3, 2, 2, 2, 381, 374, 3, 2,
	2, 2, 381, 375, 3, 2, 2, 2, 381, 376, 3, 2, 2, 2, 381, 377, 3, 2, 2, 2,
	381, 378, 3, 2, 2, 2, 381, 379, 3, 2, 2, 2, 381, 380, 3, 2, 2, 2, 382,
	53, 3, 2, 2, 2, 383, 387, 5, 32, 17, 2, 384, 387, 5, 34, 18, 2, 385, 387,
	5, 36, 19, 2, 386, 383, 3, 2, 2, 2, 386, 384, 3, 2, 2, 2, 386, 385, 3,
	2, 2, 2, 387, 388, 3, 2, 2, 2, 388, 389, 7, 29, 2, 2, 389, 390, 5, 100,
	51, 2, 390, 391, 7, 28, 2, 2, 391, 55, 3, 2, 2, 2, 392, 393, 7, 55, 2,
	2, 393, 395, 7, 30, 2, 2, 394, 396, 5, 58, 30, 2, 395, 394, 3, 2, 2, 2,
	395, 396, 3, 2, 2, 2, 396, 397, 3, 2, 2, 2, 397, 398, 7, 31, 2, 2, 398,
	57, 3, 2, 2, 2, 399, 400, 5, 100, 51, 2, 400, 401, 5, 60, 31, 2, 401, 59,
	3, 2, 2, 2, 402, 403, 7, 10, 2, 2, 403, 406, 5, 58, 30, 2, 404, 406, 3,
	2, 2, 2, 405, 402, 3, 2, 2, 2, 405, 404, 3, 2, 2, 2, 406, 61, 3, 2, 2,
	2, 407, 408, 7, 55, 2, 2, 408, 409, 7, 9, 2, 2, 409, 410, 7, 55, 2, 2,
	410, 412, 7, 30, 2, 2, 411, 413, 5, 58, 30, 2, 412, 411, 3, 2, 2, 2, 412,
	413, 3, 2, 2, 2, 413, 414, 3, 2, 2, 2, 414, 415, 7, 31, 2, 2, 415, 63,
	3, 2, 2, 2, 416, 417, 7, 11, 2, 2, 417, 418, 7, 55, 2, 2, 418, 419, 7,
	30, 2, 2, 419, 420, 7, 31, 2, 2, 420, 65, 3, 2, 2, 2, 421, 425, 5, 56,
	29, 2, 422, 425, 5, 62, 32, 2, 423, 425, 5, 64, 33, 2, 424, 421, 3, 2,
	2, 2, 424, 422, 3, 2, 2, 2, 424, 423, 3, 2, 2, 2, 425, 67, 3, 2, 2, 2,
	426, 427, 7, 44, 2, 2, 427, 428, 7, 30, 2, 2, 428, 429, 5, 70, 36, 2, 429,
	430, 7, 31, 2, 2, 430, 431, 7, 28, 2, 2, 431, 69, 3, 2, 2, 2, 432, 436,
	5, 32, 17, 2, 433, 436, 5, 34, 18, 2, 434, 436, 5, 36, 19, 2, 435, 432,
	3, 2, 2, 2, 435, 433, 3, 2, 2, 2, 435, 434, 3, 2, 2, 2, 436, 437, 3, 2,
	2, 2, 437, 438, 5, 72, 37, 2, 438, 71, 3, 2, 2, 2, 439, 440, 7, 10, 2,
	2, 440, 443, 5, 70, 36, 2, 441, 443, 3, 2, 2, 2, 442, 439, 3, 2, 2, 2,
	442, 441, 3, 2, 2, 2, 443, 73, 3, 2, 2, 2, 444, 445, 7, 43, 2, 2, 445,
	446, 7, 30, 2, 2, 446, 447, 5, 76, 39, 2, 447, 448, 7, 31, 2, 2, 448, 449,
	7, 28, 2, 2, 449, 75, 3, 2, 2, 2, 450, 451, 5, 100, 51, 2, 451, 452, 5,
	78, 40, 2, 452, 77, 3, 2, 2, 2, 453, 454, 7, 10, 2, 2, 454, 457, 5, 76,
	39, 2, 455, 457, 3, 2, 2, 2, 456, 453, 3, 2, 2, 2, 456, 455, 3, 2, 2, 2,
	457, 79, 3, 2, 2, 2, 458, 459, 7, 32, 2, 2, 459, 460, 5, 82, 42, 2, 460,
	462, 5, 84, 43, 2, 461, 463, 5, 86, 44, 2, 462, 461, 3, 2, 2, 2, 462, 463,
	3, 2, 2, 2, 463, 81, 3, 2, 2, 2, 464, 465, 7, 30, 2, 2, 465, 466, 5, 100,
	51, 2, 466, 467, 7, 31, 2, 2, 467, 83, 3, 2, 2, 2, 468, 469, 5, 48, 25,
	2, 469, 85, 3, 2, 2, 2, 470, 471, 7, 33, 2, 2, 471, 472, 5, 48, 25, 2,
	472, 87, 3, 2, 2, 2, 473, 474, 7, 35, 2, 2, 474, 475, 5, 90, 46, 2, 475,
	476, 7, 36, 2, 2, 476, 477, 5, 92, 47, 2, 477, 478, 5, 48, 25, 2, 478,
	89, 3, 2, 2, 2, 479, 480, 7, 55, 2, 2, 480, 481, 7, 29, 2, 2, 481, 482,
	5, 100, 51, 2, 482, 91, 3, 2, 2, 2, 483, 484, 5, 100, 51, 2, 484, 93, 3,
	2, 2, 2, 485, 486, 7, 34, 2, 2, 486, 487, 5, 96, 49, 2, 487, 488, 5, 48,
	25, 2, 488, 95, 3, 2, 2, 2, 489, 490, 7, 30, 2, 2, 490, 491, 5, 100, 51,
	2, 491, 492, 7, 31, 2, 2, 492, 97, 3, 2, 2, 2, 493, 494, 5, 100, 51, 2,
	494, 495, 7, 28, 2, 2, 495, 99, 3, 2, 2, 2, 496, 500, 5, 104, 53, 2, 497,
	499, 5, 102, 52, 2, 498, 497, 3, 2, 2, 2, 499, 502, 3, 2, 2, 2, 500, 498,
	3, 2, 2, 2, 500, 501, 3, 2, 2, 2, 501, 101, 3, 2, 2, 2, 502, 500, 3, 2,
	2, 2, 503, 504, 7, 18, 2, 2, 504, 505, 5, 104, 53, 2, 505, 103, 3, 2, 2,
	2, 506, 510, 5, 108, 55, 2, 507, 509, 5, 106, 54, 2, 508, 507, 3, 2, 2,
	2, 509, 512, 3, 2, 2, 2, 510, 508, 3, 2, 2, 2, 510, 511, 3, 2, 2, 2, 511,
	105, 3, 2, 2, 2, 512, 510, 3, 2, 2, 2, 513, 514, 7, 19, 2, 2, 514, 515,
	5, 108, 55, 2, 515, 107, 3, 2, 2, 2, 516, 518, 5, 112, 57, 2, 517, 519,
	5, 110, 56, 2, 518, 517, 3, 2, 2, 2, 518, 519, 3, 2, 2, 2, 519, 109, 3,
	2, 2, 2, 520, 521, 5, 142, 72, 2, 521, 522, 5, 112, 57, 2, 522, 111, 3,
	2, 2, 2, 523, 527, 5, 116, 59, 2, 524, 526, 5, 114, 58, 2, 525, 524, 3,
	2, 2, 2, 526, 529, 3, 2, 2, 2, 527, 525, 3, 2, 2, 2, 527, 528, 3, 2, 2,
	2, 528, 113, 3, 2, 2, 2, 529, 527, 3, 2, 2, 2, 530, 531, 9, 2, 2, 2, 531,
	532, 5, 116, 59, 2, 532, 115, 3, 2, 2, 2, 533, 537, 5, 120, 61, 2, 534,
	536, 5, 118, 60, 2, 535, 534, 3, 2, 2, 2, 536, 539, 3, 2, 2, 2, 537, 535,
	3, 2, 2, 2, 537, 538, 3, 2, 2, 2, 538, 117, 3, 2, 2, 2, 539, 537, 3, 2,
	2, 2, 540, 541, 9, 3, 2, 2, 541, 542, 5, 120, 61, 2, 542, 119, 3, 2, 2,
	2, 543, 552, 5, 122, 62, 2, 544, 552, 5, 124, 63, 2, 545, 549, 5, 32, 17,
	2, 546, 549, 5, 34, 18, 2, 547, 549, 5, 36, 19, 2, 548, 545, 3, 2, 2, 2,
	548, 546, 3, 2, 2, 2, 548, 547, 3, 2, 2, 2, 549, 552, 3, 2, 2, 2, 550,
	552, 5, 66, 34, 2, 551, 543, 3, 2, 2, 2, 551, 544, 3, 2, 2, 2, 551, 548,
	3, 2, 2, 2, 551, 550, 3, 2, 2, 2, 552, 121, 3, 2, 2, 2, 553, 554, 7, 30,
	2, 2, 554, 555, 5, 100, 51, 2, 555, 556, 7, 31, 2, 2, 556, 123, 3, 2, 2,
	2, 557, 563, 5, 126, 64, 2, 558, 563, 5, 128, 65, 2, 559, 563, 5, 130,
	66, 2, 560, 563, 5, 134, 68, 2, 561, 563, 5, 132, 67, 2, 562, 557, 3, 2,
	2, 2, 562, 558, 3, 2, 2, 2, 562, 559, 3, 2, 2, 2, 562, 560, 3, 2, 2, 2,
	562, 561, 3, 2, 2, 2, 563, 125, 3, 2, 2, 2, 564, 565, 7, 13, 2, 2, 565,
	127, 3, 2, 2, 2, 566, 567, 7, 14, 2, 2, 567, 129, 3, 2, 2, 2, 568, 569,
	7, 15, 2, 2, 569, 131, 3, 2, 2, 2, 570, 571, 7, 16, 2, 2, 571, 133, 3,
	2, 2, 2, 572, 576, 7, 12, 2, 2, 573, 575, 11, 2, 2, 2, 574, 573, 3, 2,
	2, 2, 575, 578, 3, 2, 2, 2, 576, 577, 3, 2, 2, 2, 576, 574, 3, 2, 2, 2,
	577, 579, 3, 2, 2, 2, 578, 576, 3, 2, 2, 2, 579, 580, 7, 12, 2, 2, 580,
	135, 3, 2, 2, 2, 581, 582, 7, 47, 2, 2, 582, 583, 7, 30, 2, 2, 583, 584,
	7, 31, 2, 2, 584, 585, 5, 138, 70, 2, 585, 137, 3, 2, 2, 2, 586, 590, 7,
	3, 2, 2, 587, 589, 5, 16, 9, 2, 588, 587, 3, 2, 2, 2, 589, 592, 3, 2, 2,
	2, 590, 588, 3, 2, 2, 2, 590, 591, 3, 2, 2, 2, 591, 596, 3, 2, 2, 2, 592,
	590, 3, 2, 2, 2, 593, 595, 5, 52, 27, 2, 594, 593, 3, 2, 2, 2, 595, 598,
	3, 2, 2, 2, 596, 594, 3, 2, 2, 2, 596, 597, 3, 2, 2, 2, 597, 599, 3, 2,
	2, 2, 598, 596, 3, 2, 2, 2, 599, 600, 7, 4, 2, 2, 600, 139, 3, 2, 2, 2,
	601, 602, 9, 4, 2, 2, 602, 141, 3, 2, 2, 2, 603, 604, 9, 5, 2, 2, 604,
	143, 3, 2, 2, 2, 51, 150, 158, 164, 182, 191, 196, 198, 209, 215, 223,
	253, 255, 260, 270, 273, 281, 286, 295, 301, 307, 313, 318, 327, 338, 344,
	351, 359, 371, 381, 386, 395, 405, 412, 424, 435, 442, 456, 462, 500, 510,
	518, 527, 537, 548, 551, 562, 576, 590, 596,
}
var literalNames = []string{
	"", "'{'", "'}'", "':'", "'['", "']'", "']['", "'.'", "','", "'new'", "'\"'",
	"", "", "", "", "", "'||'", "'&&'", "'>'", "'<'", "'=='", "'!='", "'*'",
	"'/'", "'+'", "'-'", "';'", "'='", "'('", "')'", "'if'", "'else'", "'while'",
	"'for'", "'in'", "'class'", "'attributes'", "'methods'", "'init'", "'private'",
	"'public'", "'write'", "'read'", "'function'", "'return'", "'main'", "'var'",
	"'int'", "'float'", "'char'", "'bool'", "'void'", "'program'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "INT", "FLOAT", "CHAR", "BOOL",
	"WS", "OR", "AND", "GT", "LT", "EQ", "NEQ", "MUL", "DIV", "ADD", "SUB",
	"SEMI", "ASSIGN", "LPAREN", "RPAREN", "IF", "ELSE", "WHILE", "FOR", "IN",
	"CLASS", "ATTRIBUTES", "METHODS", "INIT", "PRIVATE", "PUBLIC", "WRITE",
	"READ", "FUNCTION", "RETURN", "MAIN", "VAR", "INT_TYPE", "FLOAT_TYPE",
	"CHAR_TYPE", "BOOL_TYPE", "VOID", "PROGRAM", "ID",
}

var ruleNames = []string{
	"program", "program2", "classDef", "classBlock", "classAttributes", "classMethod",
	"classInit", "variableDeclaration", "varsDec", "varsDecArray", "varsDecMat",
	"attributesDeclaration", "attributesDec", "varsTypeInit", "varsTypeInit2",
	"vars", "varArray", "varMat", "functions", "parameters", "parameter", "functionBlock",
	"returnRule", "block", "statutesNoReturn", "statutes", "assignation", "functionCall",
	"arguments", "arguments2", "methodCall", "initCall", "call", "read", "read2",
	"read3", "write", "w_arguments", "w_arguments2", "conditional", "conditional2",
	"conditional3", "conditional4", "forLoop", "forLoop2", "forLoop3", "whileLoop",
	"whileLoop2", "expression", "exp", "exp2", "t_exp", "t_exp2", "g_exp",
	"g_exp2", "m_exp", "m_exp2", "term", "term2", "factor", "factor2", "varCte",
	"cte_i", "cte_f", "cte_c", "cte_b", "cte_s", "main", "mainBlock", "typeRule",
	"relop",
}

type BruenParser struct {
	*antlr.BaseParser
}

// NewBruenParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *BruenParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewBruenParser(input antlr.TokenStream) *BruenParser {
	this := new(BruenParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Bruen.g4"

	return this
}

// BruenParser tokens.
const (
	BruenParserEOF        = antlr.TokenEOF
	BruenParserT__0       = 1
	BruenParserT__1       = 2
	BruenParserT__2       = 3
	BruenParserT__3       = 4
	BruenParserT__4       = 5
	BruenParserT__5       = 6
	BruenParserT__6       = 7
	BruenParserT__7       = 8
	BruenParserT__8       = 9
	BruenParserT__9       = 10
	BruenParserINT        = 11
	BruenParserFLOAT      = 12
	BruenParserCHAR       = 13
	BruenParserBOOL       = 14
	BruenParserWS         = 15
	BruenParserOR         = 16
	BruenParserAND        = 17
	BruenParserGT         = 18
	BruenParserLT         = 19
	BruenParserEQ         = 20
	BruenParserNEQ        = 21
	BruenParserMUL        = 22
	BruenParserDIV        = 23
	BruenParserADD        = 24
	BruenParserSUB        = 25
	BruenParserSEMI       = 26
	BruenParserASSIGN     = 27
	BruenParserLPAREN     = 28
	BruenParserRPAREN     = 29
	BruenParserIF         = 30
	BruenParserELSE       = 31
	BruenParserWHILE      = 32
	BruenParserFOR        = 33
	BruenParserIN         = 34
	BruenParserCLASS      = 35
	BruenParserATTRIBUTES = 36
	BruenParserMETHODS    = 37
	BruenParserINIT       = 38
	BruenParserPRIVATE    = 39
	BruenParserPUBLIC     = 40
	BruenParserWRITE      = 41
	BruenParserREAD       = 42
	BruenParserFUNCTION   = 43
	BruenParserRETURN     = 44
	BruenParserMAIN       = 45
	BruenParserVAR        = 46
	BruenParserINT_TYPE   = 47
	BruenParserFLOAT_TYPE = 48
	BruenParserCHAR_TYPE  = 49
	BruenParserBOOL_TYPE  = 50
	BruenParserVOID       = 51
	BruenParserPROGRAM    = 52
	BruenParserID         = 53
)

// BruenParser rules.
const (
	BruenParserRULE_program               = 0
	BruenParserRULE_program2              = 1
	BruenParserRULE_classDef              = 2
	BruenParserRULE_classBlock            = 3
	BruenParserRULE_classAttributes       = 4
	BruenParserRULE_classMethod           = 5
	BruenParserRULE_classInit             = 6
	BruenParserRULE_variableDeclaration   = 7
	BruenParserRULE_varsDec               = 8
	BruenParserRULE_varsDecArray          = 9
	BruenParserRULE_varsDecMat            = 10
	BruenParserRULE_attributesDeclaration = 11
	BruenParserRULE_attributesDec         = 12
	BruenParserRULE_varsTypeInit          = 13
	BruenParserRULE_varsTypeInit2         = 14
	BruenParserRULE_vars                  = 15
	BruenParserRULE_varArray              = 16
	BruenParserRULE_varMat                = 17
	BruenParserRULE_functions             = 18
	BruenParserRULE_parameters            = 19
	BruenParserRULE_parameter             = 20
	BruenParserRULE_functionBlock         = 21
	BruenParserRULE_returnRule            = 22
	BruenParserRULE_block                 = 23
	BruenParserRULE_statutesNoReturn      = 24
	BruenParserRULE_statutes              = 25
	BruenParserRULE_assignation           = 26
	BruenParserRULE_functionCall          = 27
	BruenParserRULE_arguments             = 28
	BruenParserRULE_arguments2            = 29
	BruenParserRULE_methodCall            = 30
	BruenParserRULE_initCall              = 31
	BruenParserRULE_call                  = 32
	BruenParserRULE_read                  = 33
	BruenParserRULE_read2                 = 34
	BruenParserRULE_read3                 = 35
	BruenParserRULE_write                 = 36
	BruenParserRULE_w_arguments           = 37
	BruenParserRULE_w_arguments2          = 38
	BruenParserRULE_conditional           = 39
	BruenParserRULE_conditional2          = 40
	BruenParserRULE_conditional3          = 41
	BruenParserRULE_conditional4          = 42
	BruenParserRULE_forLoop               = 43
	BruenParserRULE_forLoop2              = 44
	BruenParserRULE_forLoop3              = 45
	BruenParserRULE_whileLoop             = 46
	BruenParserRULE_whileLoop2            = 47
	BruenParserRULE_expression            = 48
	BruenParserRULE_exp                   = 49
	BruenParserRULE_exp2                  = 50
	BruenParserRULE_t_exp                 = 51
	BruenParserRULE_t_exp2                = 52
	BruenParserRULE_g_exp                 = 53
	BruenParserRULE_g_exp2                = 54
	BruenParserRULE_m_exp                 = 55
	BruenParserRULE_m_exp2                = 56
	BruenParserRULE_term                  = 57
	BruenParserRULE_term2                 = 58
	BruenParserRULE_factor                = 59
	BruenParserRULE_factor2               = 60
	BruenParserRULE_varCte                = 61
	BruenParserRULE_cte_i                 = 62
	BruenParserRULE_cte_f                 = 63
	BruenParserRULE_cte_c                 = 64
	BruenParserRULE_cte_b                 = 65
	BruenParserRULE_cte_s                 = 66
	BruenParserRULE_main                  = 67
	BruenParserRULE_mainBlock             = 68
	BruenParserRULE_typeRule              = 69
	BruenParserRULE_relop                 = 70
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) PROGRAM() antlr.TerminalNode {
	return s.GetToken(BruenParserPROGRAM, 0)
}

func (s *ProgramContext) ID() antlr.TerminalNode {
	return s.GetToken(BruenParserID, 0)
}

func (s *ProgramContext) SEMI() antlr.TerminalNode {
	return s.GetToken(BruenParserSEMI, 0)
}

func (s *ProgramContext) Program2() IProgram2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProgram2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProgram2Context)
}

func (s *ProgramContext) AllVariableDeclaration() []IVariableDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem())
	var tst = make([]IVariableDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDeclarationContext)
		}
	}

	return tst
}

func (s *ProgramContext) VariableDeclaration(i int) IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *BruenParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BruenParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(BruenParserPROGRAM)
	}
	{
		p.SetState(143)
		p.Match(BruenParserID)
	}
	{
		p.SetState(144)
		p.Match(BruenParserSEMI)
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BruenParserVAR {
		{
			p.SetState(145)
			p.VariableDeclaration()
		}

		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(151)
		p.Program2()
	}

	return localctx
}

// IProgram2Context is an interface to support dynamic dispatch.
type IProgram2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgram2Context differentiates from other interfaces.
	IsProgram2Context()
}

type Program2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgram2Context() *Program2Context {
	var p = new(Program2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_program2
	return p
}

func (*Program2Context) IsProgram2Context() {}

func NewProgram2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Program2Context {
	var p = new(Program2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_program2

	return p
}

func (s *Program2Context) GetParser() antlr.Parser { return s.parser }

func (s *Program2Context) Main() IMainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMainContext)
}

func (s *Program2Context) EOF() antlr.TerminalNode {
	return s.GetToken(BruenParserEOF, 0)
}

func (s *Program2Context) AllClassDef() []IClassDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassDefContext)(nil)).Elem())
	var tst = make([]IClassDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassDefContext)
		}
	}

	return tst
}

func (s *Program2Context) ClassDef(i int) IClassDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassDefContext)
}

func (s *Program2Context) AllFunctions() []IFunctionsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionsContext)(nil)).Elem())
	var tst = make([]IFunctionsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionsContext)
		}
	}

	return tst
}

func (s *Program2Context) Functions(i int) IFunctionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionsContext)
}

func (s *Program2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Program2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Program2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterProgram2(s)
	}
}

func (s *Program2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitProgram2(s)
	}
}

func (p *BruenParser) Program2() (localctx IProgram2Context) {
	localctx = NewProgram2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BruenParserRULE_program2)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BruenParserCLASS {
		{
			p.SetState(153)
			p.ClassDef()
		}

		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BruenParserFUNCTION {
		{
			p.SetState(159)
			p.Functions()
		}

		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(165)
		p.Main()
	}
	{
		p.SetState(166)
		p.Match(BruenParserEOF)
	}

	return localctx
}

// IClassDefContext is an interface to support dynamic dispatch.
type IClassDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassDefContext differentiates from other interfaces.
	IsClassDefContext()
}

type ClassDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassDefContext() *ClassDefContext {
	var p = new(ClassDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_classDef
	return p
}

func (*ClassDefContext) IsClassDefContext() {}

func NewClassDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDefContext {
	var p = new(ClassDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_classDef

	return p
}

func (s *ClassDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDefContext) CLASS() antlr.TerminalNode {
	return s.GetToken(BruenParserCLASS, 0)
}

func (s *ClassDefContext) ID() antlr.TerminalNode {
	return s.GetToken(BruenParserID, 0)
}

func (s *ClassDefContext) ClassBlock() IClassBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassBlockContext)
}

func (s *ClassDefContext) SEMI() antlr.TerminalNode {
	return s.GetToken(BruenParserSEMI, 0)
}

func (s *ClassDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterClassDef(s)
	}
}

func (s *ClassDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitClassDef(s)
	}
}

func (p *BruenParser) ClassDef() (localctx IClassDefContext) {
	localctx = NewClassDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BruenParserRULE_classDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(BruenParserCLASS)
	}
	{
		p.SetState(169)
		p.Match(BruenParserID)
	}
	{
		p.SetState(170)
		p.ClassBlock()
	}
	{
		p.SetState(171)
		p.Match(BruenParserSEMI)
	}

	return localctx
}

// IClassBlockContext is an interface to support dynamic dispatch.
type IClassBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassBlockContext differentiates from other interfaces.
	IsClassBlockContext()
}

type ClassBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassBlockContext() *ClassBlockContext {
	var p = new(ClassBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_classBlock
	return p
}

func (*ClassBlockContext) IsClassBlockContext() {}

func NewClassBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassBlockContext {
	var p = new(ClassBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_classBlock

	return p
}

func (s *ClassBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassBlockContext) ClassAttributes() IClassAttributesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassAttributesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassAttributesContext)
}

func (s *ClassBlockContext) ClassInit() IClassInitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassInitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassInitContext)
}

func (s *ClassBlockContext) METHODS() antlr.TerminalNode {
	return s.GetToken(BruenParserMETHODS, 0)
}

func (s *ClassBlockContext) AllClassMethod() []IClassMethodContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassMethodContext)(nil)).Elem())
	var tst = make([]IClassMethodContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassMethodContext)
		}
	}

	return tst
}

func (s *ClassBlockContext) ClassMethod(i int) IClassMethodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassMethodContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassMethodContext)
}

func (s *ClassBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterClassBlock(s)
	}
}

func (s *ClassBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitClassBlock(s)
	}
}

func (p *BruenParser) ClassBlock() (localctx IClassBlockContext) {
	localctx = NewClassBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BruenParserRULE_classBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Match(BruenParserT__0)
	}
	{
		p.SetState(174)
		p.ClassAttributes()
	}
	{
		p.SetState(175)
		p.ClassInit()
	}
	{
		p.SetState(176)
		p.Match(BruenParserMETHODS)
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(BruenParserPRIVATE-39))|(1<<(BruenParserPUBLIC-39))|(1<<(BruenParserFUNCTION-39)))) != 0 {
		{
			p.SetState(177)
			p.ClassMethod()
		}

		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(183)
		p.Match(BruenParserT__1)
	}

	return localctx
}

// IClassAttributesContext is an interface to support dynamic dispatch.
type IClassAttributesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassAttributesContext differentiates from other interfaces.
	IsClassAttributesContext()
}

type ClassAttributesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassAttributesContext() *ClassAttributesContext {
	var p = new(ClassAttributesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_classAttributes
	return p
}

func (*ClassAttributesContext) IsClassAttributesContext() {}

func NewClassAttributesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassAttributesContext {
	var p = new(ClassAttributesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_classAttributes

	return p
}

func (s *ClassAttributesContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassAttributesContext) ATTRIBUTES() antlr.TerminalNode {
	return s.GetToken(BruenParserATTRIBUTES, 0)
}

func (s *ClassAttributesContext) AllAttributesDeclaration() []IAttributesDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttributesDeclarationContext)(nil)).Elem())
	var tst = make([]IAttributesDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttributesDeclarationContext)
		}
	}

	return tst
}

func (s *ClassAttributesContext) AttributesDeclaration(i int) IAttributesDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributesDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttributesDeclarationContext)
}

func (s *ClassAttributesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassAttributesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassAttributesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterClassAttributes(s)
	}
}

func (s *ClassAttributesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitClassAttributes(s)
	}
}

func (p *BruenParser) ClassAttributes() (localctx IClassAttributesContext) {
	localctx = NewClassAttributesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BruenParserRULE_classAttributes)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Match(BruenParserATTRIBUTES)
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(BruenParserPRIVATE-39))|(1<<(BruenParserPUBLIC-39))|(1<<(BruenParserVAR-39)))) != 0 {
		{
			p.SetState(186)
			p.AttributesDeclaration()
		}

		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IClassMethodContext is an interface to support dynamic dispatch.
type IClassMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassMethodContext differentiates from other interfaces.
	IsClassMethodContext()
}

type ClassMethodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassMethodContext() *ClassMethodContext {
	var p = new(ClassMethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_classMethod
	return p
}

func (*ClassMethodContext) IsClassMethodContext() {}

func NewClassMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassMethodContext {
	var p = new(ClassMethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_classMethod

	return p
}

func (s *ClassMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassMethodContext) Functions() IFunctionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionsContext)
}

func (s *ClassMethodContext) PRIVATE() antlr.TerminalNode {
	return s.GetToken(BruenParserPRIVATE, 0)
}

func (s *ClassMethodContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(BruenParserPUBLIC, 0)
}

func (s *ClassMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterClassMethod(s)
	}
}

func (s *ClassMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitClassMethod(s)
	}
}

func (p *BruenParser) ClassMethod() (localctx IClassMethodContext) {
	localctx = NewClassMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BruenParserRULE_classMethod)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(196)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BruenParserPRIVATE:
		{
			p.SetState(192)
			p.Match(BruenParserPRIVATE)
		}

	case BruenParserPUBLIC, BruenParserFUNCTION:
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BruenParserPUBLIC {
			{
				p.SetState(193)
				p.Match(BruenParserPUBLIC)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(198)
		p.Functions()
	}

	return localctx
}

// IClassInitContext is an interface to support dynamic dispatch.
type IClassInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassInitContext differentiates from other interfaces.
	IsClassInitContext()
}

type ClassInitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassInitContext() *ClassInitContext {
	var p = new(ClassInitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_classInit
	return p
}

func (*ClassInitContext) IsClassInitContext() {}

func NewClassInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassInitContext {
	var p = new(ClassInitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_classInit

	return p
}

func (s *ClassInitContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassInitContext) INIT() antlr.TerminalNode {
	return s.GetToken(BruenParserINIT, 0)
}

func (s *ClassInitContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserLPAREN, 0)
}

func (s *ClassInitContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserRPAREN, 0)
}

func (s *ClassInitContext) AllVariableDeclaration() []IVariableDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem())
	var tst = make([]IVariableDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDeclarationContext)
		}
	}

	return tst
}

func (s *ClassInitContext) VariableDeclaration(i int) IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *ClassInitContext) AllStatutesNoReturn() []IStatutesNoReturnContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatutesNoReturnContext)(nil)).Elem())
	var tst = make([]IStatutesNoReturnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatutesNoReturnContext)
		}
	}

	return tst
}

func (s *ClassInitContext) StatutesNoReturn(i int) IStatutesNoReturnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatutesNoReturnContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatutesNoReturnContext)
}

func (s *ClassInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterClassInit(s)
	}
}

func (s *ClassInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitClassInit(s)
	}
}

func (p *BruenParser) ClassInit() (localctx IClassInitContext) {
	localctx = NewClassInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BruenParserRULE_classInit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(BruenParserINIT)
	}
	{
		p.SetState(201)
		p.Match(BruenParserLPAREN)
	}
	{
		p.SetState(202)
		p.Match(BruenParserRPAREN)
	}
	{
		p.SetState(203)
		p.Match(BruenParserT__0)
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BruenParserVAR {
		{
			p.SetState(204)
			p.VariableDeclaration()
		}

		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BruenParserT__8)|(1<<BruenParserT__9)|(1<<BruenParserINT)|(1<<BruenParserFLOAT)|(1<<BruenParserCHAR)|(1<<BruenParserBOOL)|(1<<BruenParserLPAREN)|(1<<BruenParserIF))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BruenParserWHILE-32))|(1<<(BruenParserFOR-32))|(1<<(BruenParserWRITE-32))|(1<<(BruenParserREAD-32))|(1<<(BruenParserID-32)))) != 0) {
		{
			p.SetState(210)
			p.StatutesNoReturn()
		}

		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(216)
		p.Match(BruenParserT__1)
	}

	return localctx
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) VarsDec() IVarsDecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarsDecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarsDecContext)
}

func (s *VariableDeclarationContext) VarsDecArray() IVarsDecArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarsDecArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarsDecArrayContext)
}

func (s *VariableDeclarationContext) VarsDecMat() IVarsDecMatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarsDecMatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarsDecMatContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (p *BruenParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BruenParserRULE_variableDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(218)
			p.VarsDec()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(219)
			p.VarsDecArray()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(220)
			p.VarsDecMat()
		}

	}

	return localctx
}

// IVarsDecContext is an interface to support dynamic dispatch.
type IVarsDecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarsDecContext differentiates from other interfaces.
	IsVarsDecContext()
}

type VarsDecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarsDecContext() *VarsDecContext {
	var p = new(VarsDecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_varsDec
	return p
}

func (*VarsDecContext) IsVarsDecContext() {}

func NewVarsDecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsDecContext {
	var p = new(VarsDecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_varsDec

	return p
}

func (s *VarsDecContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsDecContext) VAR() antlr.TerminalNode {
	return s.GetToken(BruenParserVAR, 0)
}

func (s *VarsDecContext) ID() antlr.TerminalNode {
	return s.GetToken(BruenParserID, 0)
}

func (s *VarsDecContext) VarsTypeInit() IVarsTypeInitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarsTypeInitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarsTypeInitContext)
}

func (s *VarsDecContext) SEMI() antlr.TerminalNode {
	return s.GetToken(BruenParserSEMI, 0)
}

func (s *VarsDecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsDecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsDecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterVarsDec(s)
	}
}

func (s *VarsDecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitVarsDec(s)
	}
}

func (p *BruenParser) VarsDec() (localctx IVarsDecContext) {
	localctx = NewVarsDecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BruenParserRULE_varsDec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(BruenParserVAR)
	}
	{
		p.SetState(224)
		p.Match(BruenParserID)
	}
	{
		p.SetState(225)
		p.Match(BruenParserT__2)
	}
	{
		p.SetState(226)
		p.VarsTypeInit()
	}
	{
		p.SetState(227)
		p.Match(BruenParserSEMI)
	}

	return localctx
}

// IVarsDecArrayContext is an interface to support dynamic dispatch.
type IVarsDecArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarsDecArrayContext differentiates from other interfaces.
	IsVarsDecArrayContext()
}

type VarsDecArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarsDecArrayContext() *VarsDecArrayContext {
	var p = new(VarsDecArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_varsDecArray
	return p
}

func (*VarsDecArrayContext) IsVarsDecArrayContext() {}

func NewVarsDecArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsDecArrayContext {
	var p = new(VarsDecArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_varsDecArray

	return p
}

func (s *VarsDecArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsDecArrayContext) VAR() antlr.TerminalNode {
	return s.GetToken(BruenParserVAR, 0)
}

func (s *VarsDecArrayContext) ID() antlr.TerminalNode {
	return s.GetToken(BruenParserID, 0)
}

func (s *VarsDecArrayContext) INT() antlr.TerminalNode {
	return s.GetToken(BruenParserINT, 0)
}

func (s *VarsDecArrayContext) TypeRule() ITypeRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRuleContext)
}

func (s *VarsDecArrayContext) SEMI() antlr.TerminalNode {
	return s.GetToken(BruenParserSEMI, 0)
}

func (s *VarsDecArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsDecArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsDecArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterVarsDecArray(s)
	}
}

func (s *VarsDecArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitVarsDecArray(s)
	}
}

func (p *BruenParser) VarsDecArray() (localctx IVarsDecArrayContext) {
	localctx = NewVarsDecArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BruenParserRULE_varsDecArray)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Match(BruenParserVAR)
	}
	{
		p.SetState(230)
		p.Match(BruenParserID)
	}
	{
		p.SetState(231)
		p.Match(BruenParserT__2)
	}
	{
		p.SetState(232)
		p.Match(BruenParserT__3)
	}
	{
		p.SetState(233)
		p.Match(BruenParserINT)
	}
	{
		p.SetState(234)
		p.Match(BruenParserT__4)
	}
	{
		p.SetState(235)
		p.TypeRule()
	}
	{
		p.SetState(236)
		p.Match(BruenParserSEMI)
	}

	return localctx
}

// IVarsDecMatContext is an interface to support dynamic dispatch.
type IVarsDecMatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarsDecMatContext differentiates from other interfaces.
	IsVarsDecMatContext()
}

type VarsDecMatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarsDecMatContext() *VarsDecMatContext {
	var p = new(VarsDecMatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_varsDecMat
	return p
}

func (*VarsDecMatContext) IsVarsDecMatContext() {}

func NewVarsDecMatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsDecMatContext {
	var p = new(VarsDecMatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_varsDecMat

	return p
}

func (s *VarsDecMatContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsDecMatContext) VAR() antlr.TerminalNode {
	return s.GetToken(BruenParserVAR, 0)
}

func (s *VarsDecMatContext) ID() antlr.TerminalNode {
	return s.GetToken(BruenParserID, 0)
}

func (s *VarsDecMatContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(BruenParserINT)
}

func (s *VarsDecMatContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(BruenParserINT, i)
}

func (s *VarsDecMatContext) TypeRule() ITypeRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRuleContext)
}

func (s *VarsDecMatContext) SEMI() antlr.TerminalNode {
	return s.GetToken(BruenParserSEMI, 0)
}

func (s *VarsDecMatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsDecMatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsDecMatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterVarsDecMat(s)
	}
}

func (s *VarsDecMatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitVarsDecMat(s)
	}
}

func (p *BruenParser) VarsDecMat() (localctx IVarsDecMatContext) {
	localctx = NewVarsDecMatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BruenParserRULE_varsDecMat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Match(BruenParserVAR)
	}
	{
		p.SetState(239)
		p.Match(BruenParserID)
	}
	{
		p.SetState(240)
		p.Match(BruenParserT__2)
	}
	{
		p.SetState(241)
		p.Match(BruenParserT__3)
	}
	{
		p.SetState(242)
		p.Match(BruenParserINT)
	}
	{
		p.SetState(243)
		p.Match(BruenParserT__5)
	}
	{
		p.SetState(244)
		p.Match(BruenParserINT)
	}
	{
		p.SetState(245)
		p.Match(BruenParserT__4)
	}
	{
		p.SetState(246)
		p.TypeRule()
	}
	{
		p.SetState(247)
		p.Match(BruenParserSEMI)
	}

	return localctx
}

// IAttributesDeclarationContext is an interface to support dynamic dispatch.
type IAttributesDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributesDeclarationContext differentiates from other interfaces.
	IsAttributesDeclarationContext()
}

type AttributesDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributesDeclarationContext() *AttributesDeclarationContext {
	var p = new(AttributesDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_attributesDeclaration
	return p
}

func (*AttributesDeclarationContext) IsAttributesDeclarationContext() {}

func NewAttributesDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributesDeclarationContext {
	var p = new(AttributesDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_attributesDeclaration

	return p
}

func (s *AttributesDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributesDeclarationContext) PRIVATE() antlr.TerminalNode {
	return s.GetToken(BruenParserPRIVATE, 0)
}

func (s *AttributesDeclarationContext) AttributesDec() IAttributesDecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributesDecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributesDecContext)
}

func (s *AttributesDeclarationContext) VarsDecArray() IVarsDecArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarsDecArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarsDecArrayContext)
}

func (s *AttributesDeclarationContext) VarsDecMat() IVarsDecMatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarsDecMatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarsDecMatContext)
}

func (s *AttributesDeclarationContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(BruenParserPUBLIC, 0)
}

func (s *AttributesDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributesDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributesDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterAttributesDeclaration(s)
	}
}

func (s *AttributesDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitAttributesDeclaration(s)
	}
}

func (p *BruenParser) AttributesDeclaration() (localctx IAttributesDeclarationContext) {
	localctx = NewAttributesDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BruenParserRULE_attributesDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(253)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BruenParserPRIVATE:
		{
			p.SetState(249)
			p.Match(BruenParserPRIVATE)
		}

	case BruenParserPUBLIC, BruenParserVAR:
		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == BruenParserPUBLIC {
			{
				p.SetState(250)
				p.Match(BruenParserPUBLIC)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(255)
			p.AttributesDec()
		}

	case 2:
		{
			p.SetState(256)
			p.VarsDecArray()
		}

	case 3:
		{
			p.SetState(257)
			p.VarsDecMat()
		}

	}

	return localctx
}

// IAttributesDecContext is an interface to support dynamic dispatch.
type IAttributesDecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributesDecContext differentiates from other interfaces.
	IsAttributesDecContext()
}

type AttributesDecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributesDecContext() *AttributesDecContext {
	var p = new(AttributesDecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_attributesDec
	return p
}

func (*AttributesDecContext) IsAttributesDecContext() {}

func NewAttributesDecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributesDecContext {
	var p = new(AttributesDecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_attributesDec

	return p
}

func (s *AttributesDecContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributesDecContext) VAR() antlr.TerminalNode {
	return s.GetToken(BruenParserVAR, 0)
}

func (s *AttributesDecContext) ID() antlr.TerminalNode {
	return s.GetToken(BruenParserID, 0)
}

func (s *AttributesDecContext) TypeRule() ITypeRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRuleContext)
}

func (s *AttributesDecContext) SEMI() antlr.TerminalNode {
	return s.GetToken(BruenParserSEMI, 0)
}

func (s *AttributesDecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributesDecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributesDecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterAttributesDec(s)
	}
}

func (s *AttributesDecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitAttributesDec(s)
	}
}

func (p *BruenParser) AttributesDec() (localctx IAttributesDecContext) {
	localctx = NewAttributesDecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BruenParserRULE_attributesDec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		p.Match(BruenParserVAR)
	}
	{
		p.SetState(261)
		p.Match(BruenParserID)
	}
	{
		p.SetState(262)
		p.Match(BruenParserT__2)
	}
	{
		p.SetState(263)
		p.TypeRule()
	}
	{
		p.SetState(264)
		p.Match(BruenParserSEMI)
	}

	return localctx
}

// IVarsTypeInitContext is an interface to support dynamic dispatch.
type IVarsTypeInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarsTypeInitContext differentiates from other interfaces.
	IsVarsTypeInitContext()
}

type VarsTypeInitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarsTypeInitContext() *VarsTypeInitContext {
	var p = new(VarsTypeInitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_varsTypeInit
	return p
}

func (*VarsTypeInitContext) IsVarsTypeInitContext() {}

func NewVarsTypeInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsTypeInitContext {
	var p = new(VarsTypeInitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_varsTypeInit

	return p
}

func (s *VarsTypeInitContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsTypeInitContext) TypeRule() ITypeRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRuleContext)
}

func (s *VarsTypeInitContext) ID() antlr.TerminalNode {
	return s.GetToken(BruenParserID, 0)
}

func (s *VarsTypeInitContext) VarsTypeInit2() IVarsTypeInit2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarsTypeInit2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarsTypeInit2Context)
}

func (s *VarsTypeInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsTypeInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsTypeInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterVarsTypeInit(s)
	}
}

func (s *VarsTypeInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitVarsTypeInit(s)
	}
}

func (p *BruenParser) VarsTypeInit() (localctx IVarsTypeInitContext) {
	localctx = NewVarsTypeInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BruenParserRULE_varsTypeInit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(268)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BruenParserINT_TYPE, BruenParserFLOAT_TYPE, BruenParserCHAR_TYPE, BruenParserBOOL_TYPE:
		{
			p.SetState(266)
			p.TypeRule()
		}

	case BruenParserID:
		{
			p.SetState(267)
			p.Match(BruenParserID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BruenParserASSIGN {
		{
			p.SetState(270)
			p.VarsTypeInit2()
		}

	}

	return localctx
}

// IVarsTypeInit2Context is an interface to support dynamic dispatch.
type IVarsTypeInit2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarsTypeInit2Context differentiates from other interfaces.
	IsVarsTypeInit2Context()
}

type VarsTypeInit2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarsTypeInit2Context() *VarsTypeInit2Context {
	var p = new(VarsTypeInit2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_varsTypeInit2
	return p
}

func (*VarsTypeInit2Context) IsVarsTypeInit2Context() {}

func NewVarsTypeInit2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsTypeInit2Context {
	var p = new(VarsTypeInit2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_varsTypeInit2

	return p
}

func (s *VarsTypeInit2Context) GetParser() antlr.Parser { return s.parser }

func (s *VarsTypeInit2Context) ASSIGN() antlr.TerminalNode {
	return s.GetToken(BruenParserASSIGN, 0)
}

func (s *VarsTypeInit2Context) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *VarsTypeInit2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsTypeInit2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsTypeInit2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterVarsTypeInit2(s)
	}
}

func (s *VarsTypeInit2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitVarsTypeInit2(s)
	}
}

func (p *BruenParser) VarsTypeInit2() (localctx IVarsTypeInit2Context) {
	localctx = NewVarsTypeInit2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BruenParserRULE_varsTypeInit2)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		p.Match(BruenParserASSIGN)
	}
	{
		p.SetState(274)
		p.Exp()
	}

	return localctx
}

// IVarsContext is an interface to support dynamic dispatch.
type IVarsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarsContext differentiates from other interfaces.
	IsVarsContext()
}

type VarsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarsContext() *VarsContext {
	var p = new(VarsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_vars
	return p
}

func (*VarsContext) IsVarsContext() {}

func NewVarsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsContext {
	var p = new(VarsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_vars

	return p
}

func (s *VarsContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(BruenParserID)
}

func (s *VarsContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(BruenParserID, i)
}

func (s *VarsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterVars(s)
	}
}

func (s *VarsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitVars(s)
	}
}

func (p *BruenParser) Vars() (localctx IVarsContext) {
	localctx = NewVarsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BruenParserRULE_vars)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Match(BruenParserID)
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BruenParserT__6 {
		{
			p.SetState(277)
			p.Match(BruenParserT__6)
		}
		{
			p.SetState(278)
			p.Match(BruenParserID)
		}

	}

	return localctx
}

// IVarArrayContext is an interface to support dynamic dispatch.
type IVarArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarArrayContext differentiates from other interfaces.
	IsVarArrayContext()
}

type VarArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarArrayContext() *VarArrayContext {
	var p = new(VarArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_varArray
	return p
}

func (*VarArrayContext) IsVarArrayContext() {}

func NewVarArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarArrayContext {
	var p = new(VarArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_varArray

	return p
}

func (s *VarArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *VarArrayContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(BruenParserID)
}

func (s *VarArrayContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(BruenParserID, i)
}

func (s *VarArrayContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *VarArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterVarArray(s)
	}
}

func (s *VarArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitVarArray(s)
	}
}

func (p *BruenParser) VarArray() (localctx IVarArrayContext) {
	localctx = NewVarArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BruenParserRULE_varArray)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.Match(BruenParserID)
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BruenParserT__6 {
		{
			p.SetState(282)
			p.Match(BruenParserT__6)
		}
		{
			p.SetState(283)
			p.Match(BruenParserID)
		}

	}
	{
		p.SetState(286)
		p.Match(BruenParserT__3)
	}
	{
		p.SetState(287)
		p.Exp()
	}
	{
		p.SetState(288)
		p.Match(BruenParserT__4)
	}

	return localctx
}

// IVarMatContext is an interface to support dynamic dispatch.
type IVarMatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarMatContext differentiates from other interfaces.
	IsVarMatContext()
}

type VarMatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarMatContext() *VarMatContext {
	var p = new(VarMatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_varMat
	return p
}

func (*VarMatContext) IsVarMatContext() {}

func NewVarMatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarMatContext {
	var p = new(VarMatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_varMat

	return p
}

func (s *VarMatContext) GetParser() antlr.Parser { return s.parser }

func (s *VarMatContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(BruenParserID)
}

func (s *VarMatContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(BruenParserID, i)
}

func (s *VarMatContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *VarMatContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *VarMatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarMatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarMatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterVarMat(s)
	}
}

func (s *VarMatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitVarMat(s)
	}
}

func (p *BruenParser) VarMat() (localctx IVarMatContext) {
	localctx = NewVarMatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BruenParserRULE_varMat)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Match(BruenParserID)
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BruenParserT__6 {
		{
			p.SetState(291)
			p.Match(BruenParserT__6)
		}
		{
			p.SetState(292)
			p.Match(BruenParserID)
		}

	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(295)
			p.Match(BruenParserT__3)
		}
		{
			p.SetState(296)
			p.Exp()
		}
		{
			p.SetState(297)
			p.Match(BruenParserT__4)
		}

	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BruenParserT__3 {
		{
			p.SetState(301)
			p.Match(BruenParserT__3)
		}
		{
			p.SetState(302)
			p.Exp()
		}
		{
			p.SetState(303)
			p.Match(BruenParserT__4)
		}

	}

	return localctx
}

// IFunctionsContext is an interface to support dynamic dispatch.
type IFunctionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionsContext differentiates from other interfaces.
	IsFunctionsContext()
}

type FunctionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionsContext() *FunctionsContext {
	var p = new(FunctionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_functions
	return p
}

func (*FunctionsContext) IsFunctionsContext() {}

func NewFunctionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionsContext {
	var p = new(FunctionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_functions

	return p
}

func (s *FunctionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionsContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(BruenParserFUNCTION, 0)
}

func (s *FunctionsContext) ID() antlr.TerminalNode {
	return s.GetToken(BruenParserID, 0)
}

func (s *FunctionsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserLPAREN, 0)
}

func (s *FunctionsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserRPAREN, 0)
}

func (s *FunctionsContext) FunctionBlock() IFunctionBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBlockContext)
}

func (s *FunctionsContext) TypeRule() ITypeRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRuleContext)
}

func (s *FunctionsContext) VOID() antlr.TerminalNode {
	return s.GetToken(BruenParserVOID, 0)
}

func (s *FunctionsContext) Parameters() IParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *FunctionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterFunctions(s)
	}
}

func (s *FunctionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitFunctions(s)
	}
}

func (p *BruenParser) Functions() (localctx IFunctionsContext) {
	localctx = NewFunctionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BruenParserRULE_functions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.Match(BruenParserFUNCTION)
	}
	{
		p.SetState(308)
		p.Match(BruenParserID)
	}
	{
		p.SetState(309)
		p.Match(BruenParserLPAREN)
	}
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BruenParserID {
		{
			p.SetState(310)
			p.Parameters()
		}

	}
	{
		p.SetState(313)
		p.Match(BruenParserRPAREN)
	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BruenParserINT_TYPE, BruenParserFLOAT_TYPE, BruenParserCHAR_TYPE, BruenParserBOOL_TYPE:
		{
			p.SetState(314)
			p.TypeRule()
		}

	case BruenParserVOID:
		{
			p.SetState(315)
			p.Match(BruenParserVOID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(318)
		p.FunctionBlock()
	}

	return localctx
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) AllParameter() []IParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterContext)(nil)).Elem())
	var tst = make([]IParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterContext)
		}
	}

	return tst
}

func (s *ParametersContext) Parameter(i int) IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (p *BruenParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BruenParserRULE_parameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.Parameter()
	}
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BruenParserT__7 {
		{
			p.SetState(321)
			p.Match(BruenParserT__7)
		}
		{
			p.SetState(322)
			p.Parameter()
		}

		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) ID() antlr.TerminalNode {
	return s.GetToken(BruenParserID, 0)
}

func (s *ParameterContext) TypeRule() ITypeRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRuleContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (p *BruenParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, BruenParserRULE_parameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		p.Match(BruenParserID)
	}
	{
		p.SetState(329)
		p.Match(BruenParserT__2)
	}
	{
		p.SetState(330)
		p.TypeRule()
	}

	return localctx
}

// IFunctionBlockContext is an interface to support dynamic dispatch.
type IFunctionBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionBlockContext differentiates from other interfaces.
	IsFunctionBlockContext()
}

type FunctionBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionBlockContext() *FunctionBlockContext {
	var p = new(FunctionBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_functionBlock
	return p
}

func (*FunctionBlockContext) IsFunctionBlockContext() {}

func NewFunctionBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionBlockContext {
	var p = new(FunctionBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_functionBlock

	return p
}

func (s *FunctionBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionBlockContext) AllVariableDeclaration() []IVariableDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem())
	var tst = make([]IVariableDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDeclarationContext)
		}
	}

	return tst
}

func (s *FunctionBlockContext) VariableDeclaration(i int) IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *FunctionBlockContext) AllStatutes() []IStatutesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatutesContext)(nil)).Elem())
	var tst = make([]IStatutesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatutesContext)
		}
	}

	return tst
}

func (s *FunctionBlockContext) Statutes(i int) IStatutesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatutesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatutesContext)
}

func (s *FunctionBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterFunctionBlock(s)
	}
}

func (s *FunctionBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitFunctionBlock(s)
	}
}

func (p *BruenParser) FunctionBlock() (localctx IFunctionBlockContext) {
	localctx = NewFunctionBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BruenParserRULE_functionBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
		p.Match(BruenParserT__0)
	}
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BruenParserVAR {
		{
			p.SetState(333)
			p.VariableDeclaration()
		}

		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BruenParserT__8)|(1<<BruenParserT__9)|(1<<BruenParserINT)|(1<<BruenParserFLOAT)|(1<<BruenParserCHAR)|(1<<BruenParserBOOL)|(1<<BruenParserLPAREN)|(1<<BruenParserIF))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BruenParserWHILE-32))|(1<<(BruenParserFOR-32))|(1<<(BruenParserWRITE-32))|(1<<(BruenParserREAD-32))|(1<<(BruenParserRETURN-32))|(1<<(BruenParserID-32)))) != 0) {
		{
			p.SetState(339)
			p.Statutes()
		}

		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(345)
		p.Match(BruenParserT__1)
	}

	return localctx
}

// IReturnRuleContext is an interface to support dynamic dispatch.
type IReturnRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnRuleContext differentiates from other interfaces.
	IsReturnRuleContext()
}

type ReturnRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnRuleContext() *ReturnRuleContext {
	var p = new(ReturnRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_returnRule
	return p
}

func (*ReturnRuleContext) IsReturnRuleContext() {}

func NewReturnRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnRuleContext {
	var p = new(ReturnRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_returnRule

	return p
}

func (s *ReturnRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnRuleContext) RETURN() antlr.TerminalNode {
	return s.GetToken(BruenParserRETURN, 0)
}

func (s *ReturnRuleContext) SEMI() antlr.TerminalNode {
	return s.GetToken(BruenParserSEMI, 0)
}

func (s *ReturnRuleContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ReturnRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterReturnRule(s)
	}
}

func (s *ReturnRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitReturnRule(s)
	}
}

func (p *BruenParser) ReturnRule() (localctx IReturnRuleContext) {
	localctx = NewReturnRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, BruenParserRULE_returnRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.Match(BruenParserRETURN)
	}
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BruenParserT__8)|(1<<BruenParserT__9)|(1<<BruenParserINT)|(1<<BruenParserFLOAT)|(1<<BruenParserCHAR)|(1<<BruenParserBOOL)|(1<<BruenParserLPAREN))) != 0) || _la == BruenParserID {
		{
			p.SetState(348)
			p.Exp()
		}

	}
	{
		p.SetState(351)
		p.Match(BruenParserSEMI)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStatutes() []IStatutesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatutesContext)(nil)).Elem())
	var tst = make([]IStatutesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatutesContext)
		}
	}

	return tst
}

func (s *BlockContext) Statutes(i int) IStatutesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatutesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatutesContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *BruenParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, BruenParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(353)
		p.Match(BruenParserT__0)
	}
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BruenParserT__8)|(1<<BruenParserT__9)|(1<<BruenParserINT)|(1<<BruenParserFLOAT)|(1<<BruenParserCHAR)|(1<<BruenParserBOOL)|(1<<BruenParserLPAREN)|(1<<BruenParserIF))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BruenParserWHILE-32))|(1<<(BruenParserFOR-32))|(1<<(BruenParserWRITE-32))|(1<<(BruenParserREAD-32))|(1<<(BruenParserRETURN-32))|(1<<(BruenParserID-32)))) != 0) {
		{
			p.SetState(354)
			p.Statutes()
		}

		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(360)
		p.Match(BruenParserT__1)
	}

	return localctx
}

// IStatutesNoReturnContext is an interface to support dynamic dispatch.
type IStatutesNoReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatutesNoReturnContext differentiates from other interfaces.
	IsStatutesNoReturnContext()
}

type StatutesNoReturnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatutesNoReturnContext() *StatutesNoReturnContext {
	var p = new(StatutesNoReturnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_statutesNoReturn
	return p
}

func (*StatutesNoReturnContext) IsStatutesNoReturnContext() {}

func NewStatutesNoReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatutesNoReturnContext {
	var p = new(StatutesNoReturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_statutesNoReturn

	return p
}

func (s *StatutesNoReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *StatutesNoReturnContext) Assignation() IAssignationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignationContext)
}

func (s *StatutesNoReturnContext) Read() IReadContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReadContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReadContext)
}

func (s *StatutesNoReturnContext) Write() IWriteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWriteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWriteContext)
}

func (s *StatutesNoReturnContext) Conditional() IConditionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionalContext)
}

func (s *StatutesNoReturnContext) ForLoop() IForLoopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForLoopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForLoopContext)
}

func (s *StatutesNoReturnContext) WhileLoop() IWhileLoopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileLoopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhileLoopContext)
}

func (s *StatutesNoReturnContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatutesNoReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatutesNoReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatutesNoReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterStatutesNoReturn(s)
	}
}

func (s *StatutesNoReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitStatutesNoReturn(s)
	}
}

func (p *BruenParser) StatutesNoReturn() (localctx IStatutesNoReturnContext) {
	localctx = NewStatutesNoReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, BruenParserRULE_statutesNoReturn)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(362)
			p.Assignation()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(363)
			p.Read()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(364)
			p.Write()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(365)
			p.Conditional()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(366)
			p.ForLoop()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(367)
			p.WhileLoop()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(368)
			p.Expression()
		}

	}

	return localctx
}

// IStatutesContext is an interface to support dynamic dispatch.
type IStatutesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatutesContext differentiates from other interfaces.
	IsStatutesContext()
}

type StatutesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatutesContext() *StatutesContext {
	var p = new(StatutesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_statutes
	return p
}

func (*StatutesContext) IsStatutesContext() {}

func NewStatutesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatutesContext {
	var p = new(StatutesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_statutes

	return p
}

func (s *StatutesContext) GetParser() antlr.Parser { return s.parser }

func (s *StatutesContext) Assignation() IAssignationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignationContext)
}

func (s *StatutesContext) Read() IReadContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReadContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReadContext)
}

func (s *StatutesContext) Write() IWriteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWriteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWriteContext)
}

func (s *StatutesContext) Conditional() IConditionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionalContext)
}

func (s *StatutesContext) ForLoop() IForLoopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForLoopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForLoopContext)
}

func (s *StatutesContext) WhileLoop() IWhileLoopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileLoopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhileLoopContext)
}

func (s *StatutesContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatutesContext) ReturnRule() IReturnRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnRuleContext)
}

func (s *StatutesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatutesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatutesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterStatutes(s)
	}
}

func (s *StatutesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitStatutes(s)
	}
}

func (p *BruenParser) Statutes() (localctx IStatutesContext) {
	localctx = NewStatutesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, BruenParserRULE_statutes)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(371)
			p.Assignation()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(372)
			p.Read()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(373)
			p.Write()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(374)
			p.Conditional()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(375)
			p.ForLoop()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(376)
			p.WhileLoop()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(377)
			p.Expression()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(378)
			p.ReturnRule()
		}

	}

	return localctx
}

// IAssignationContext is an interface to support dynamic dispatch.
type IAssignationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignationContext differentiates from other interfaces.
	IsAssignationContext()
}

type AssignationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignationContext() *AssignationContext {
	var p = new(AssignationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_assignation
	return p
}

func (*AssignationContext) IsAssignationContext() {}

func NewAssignationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignationContext {
	var p = new(AssignationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_assignation

	return p
}

func (s *AssignationContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(BruenParserASSIGN, 0)
}

func (s *AssignationContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AssignationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(BruenParserSEMI, 0)
}

func (s *AssignationContext) Vars() IVarsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarsContext)
}

func (s *AssignationContext) VarArray() IVarArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarArrayContext)
}

func (s *AssignationContext) VarMat() IVarMatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarMatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarMatContext)
}

func (s *AssignationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterAssignation(s)
	}
}

func (s *AssignationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitAssignation(s)
	}
}

func (p *BruenParser) Assignation() (localctx IAssignationContext) {
	localctx = NewAssignationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, BruenParserRULE_assignation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(381)
			p.Vars()
		}

	case 2:
		{
			p.SetState(382)
			p.VarArray()
		}

	case 3:
		{
			p.SetState(383)
			p.VarMat()
		}

	}
	{
		p.SetState(386)
		p.Match(BruenParserASSIGN)
	}
	{
		p.SetState(387)
		p.Exp()
	}
	{
		p.SetState(388)
		p.Match(BruenParserSEMI)
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) ID() antlr.TerminalNode {
	return s.GetToken(BruenParserID, 0)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserRPAREN, 0)
}

func (s *FunctionCallContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *BruenParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, BruenParserRULE_functionCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.Match(BruenParserID)
	}
	{
		p.SetState(391)
		p.Match(BruenParserLPAREN)
	}
	p.SetState(393)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BruenParserT__8)|(1<<BruenParserT__9)|(1<<BruenParserINT)|(1<<BruenParserFLOAT)|(1<<BruenParserCHAR)|(1<<BruenParserBOOL)|(1<<BruenParserLPAREN))) != 0) || _la == BruenParserID {
		{
			p.SetState(392)
			p.Arguments()
		}

	}
	{
		p.SetState(395)
		p.Match(BruenParserRPAREN)
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ArgumentsContext) Arguments2() IArguments2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArguments2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArguments2Context)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *BruenParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, BruenParserRULE_arguments)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(397)
		p.Exp()
	}
	{
		p.SetState(398)
		p.Arguments2()
	}

	return localctx
}

// IArguments2Context is an interface to support dynamic dispatch.
type IArguments2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArguments2Context differentiates from other interfaces.
	IsArguments2Context()
}

type Arguments2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArguments2Context() *Arguments2Context {
	var p = new(Arguments2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_arguments2
	return p
}

func (*Arguments2Context) IsArguments2Context() {}

func NewArguments2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arguments2Context {
	var p = new(Arguments2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_arguments2

	return p
}

func (s *Arguments2Context) GetParser() antlr.Parser { return s.parser }

func (s *Arguments2Context) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *Arguments2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arguments2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arguments2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterArguments2(s)
	}
}

func (s *Arguments2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitArguments2(s)
	}
}

func (p *BruenParser) Arguments2() (localctx IArguments2Context) {
	localctx = NewArguments2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, BruenParserRULE_arguments2)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(403)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BruenParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(400)
			p.Match(BruenParserT__7)
		}
		{
			p.SetState(401)
			p.Arguments()
		}

	case BruenParserRPAREN:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMethodCallContext is an interface to support dynamic dispatch.
type IMethodCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodCallContext differentiates from other interfaces.
	IsMethodCallContext()
}

type MethodCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodCallContext() *MethodCallContext {
	var p = new(MethodCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_methodCall
	return p
}

func (*MethodCallContext) IsMethodCallContext() {}

func NewMethodCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallContext {
	var p = new(MethodCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_methodCall

	return p
}

func (s *MethodCallContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(BruenParserID)
}

func (s *MethodCallContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(BruenParserID, i)
}

func (s *MethodCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserLPAREN, 0)
}

func (s *MethodCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserRPAREN, 0)
}

func (s *MethodCallContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *MethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterMethodCall(s)
	}
}

func (s *MethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitMethodCall(s)
	}
}

func (p *BruenParser) MethodCall() (localctx IMethodCallContext) {
	localctx = NewMethodCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, BruenParserRULE_methodCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(405)
		p.Match(BruenParserID)
	}
	{
		p.SetState(406)
		p.Match(BruenParserT__6)
	}
	{
		p.SetState(407)
		p.Match(BruenParserID)
	}
	{
		p.SetState(408)
		p.Match(BruenParserLPAREN)
	}
	p.SetState(410)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BruenParserT__8)|(1<<BruenParserT__9)|(1<<BruenParserINT)|(1<<BruenParserFLOAT)|(1<<BruenParserCHAR)|(1<<BruenParserBOOL)|(1<<BruenParserLPAREN))) != 0) || _la == BruenParserID {
		{
			p.SetState(409)
			p.Arguments()
		}

	}
	{
		p.SetState(412)
		p.Match(BruenParserRPAREN)
	}

	return localctx
}

// IInitCallContext is an interface to support dynamic dispatch.
type IInitCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitCallContext differentiates from other interfaces.
	IsInitCallContext()
}

type InitCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitCallContext() *InitCallContext {
	var p = new(InitCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_initCall
	return p
}

func (*InitCallContext) IsInitCallContext() {}

func NewInitCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitCallContext {
	var p = new(InitCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_initCall

	return p
}

func (s *InitCallContext) GetParser() antlr.Parser { return s.parser }

func (s *InitCallContext) ID() antlr.TerminalNode {
	return s.GetToken(BruenParserID, 0)
}

func (s *InitCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserLPAREN, 0)
}

func (s *InitCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserRPAREN, 0)
}

func (s *InitCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterInitCall(s)
	}
}

func (s *InitCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitInitCall(s)
	}
}

func (p *BruenParser) InitCall() (localctx IInitCallContext) {
	localctx = NewInitCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, BruenParserRULE_initCall)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(414)
		p.Match(BruenParserT__8)
	}
	{
		p.SetState(415)
		p.Match(BruenParserID)
	}
	{
		p.SetState(416)
		p.Match(BruenParserLPAREN)
	}
	{
		p.SetState(417)
		p.Match(BruenParserRPAREN)
	}

	return localctx
}

// ICallContext is an interface to support dynamic dispatch.
type ICallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCallContext differentiates from other interfaces.
	IsCallContext()
}

type CallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallContext() *CallContext {
	var p = new(CallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_call
	return p
}

func (*CallContext) IsCallContext() {}

func NewCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallContext {
	var p = new(CallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_call

	return p
}

func (s *CallContext) GetParser() antlr.Parser { return s.parser }

func (s *CallContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *CallContext) MethodCall() IMethodCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *CallContext) InitCall() IInitCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitCallContext)
}

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterCall(s)
	}
}

func (s *CallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitCall(s)
	}
}

func (p *BruenParser) Call() (localctx ICallContext) {
	localctx = NewCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, BruenParserRULE_call)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(419)
			p.FunctionCall()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(420)
			p.MethodCall()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(421)
			p.InitCall()
		}

	}

	return localctx
}

// IReadContext is an interface to support dynamic dispatch.
type IReadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReadContext differentiates from other interfaces.
	IsReadContext()
}

type ReadContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReadContext() *ReadContext {
	var p = new(ReadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_read
	return p
}

func (*ReadContext) IsReadContext() {}

func NewReadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReadContext {
	var p = new(ReadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_read

	return p
}

func (s *ReadContext) GetParser() antlr.Parser { return s.parser }

func (s *ReadContext) READ() antlr.TerminalNode {
	return s.GetToken(BruenParserREAD, 0)
}

func (s *ReadContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserLPAREN, 0)
}

func (s *ReadContext) Read2() IRead2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRead2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRead2Context)
}

func (s *ReadContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserRPAREN, 0)
}

func (s *ReadContext) SEMI() antlr.TerminalNode {
	return s.GetToken(BruenParserSEMI, 0)
}

func (s *ReadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterRead(s)
	}
}

func (s *ReadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitRead(s)
	}
}

func (p *BruenParser) Read() (localctx IReadContext) {
	localctx = NewReadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, BruenParserRULE_read)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(424)
		p.Match(BruenParserREAD)
	}
	{
		p.SetState(425)
		p.Match(BruenParserLPAREN)
	}
	{
		p.SetState(426)
		p.Read2()
	}
	{
		p.SetState(427)
		p.Match(BruenParserRPAREN)
	}
	{
		p.SetState(428)
		p.Match(BruenParserSEMI)
	}

	return localctx
}

// IRead2Context is an interface to support dynamic dispatch.
type IRead2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRead2Context differentiates from other interfaces.
	IsRead2Context()
}

type Read2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRead2Context() *Read2Context {
	var p = new(Read2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_read2
	return p
}

func (*Read2Context) IsRead2Context() {}

func NewRead2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Read2Context {
	var p = new(Read2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_read2

	return p
}

func (s *Read2Context) GetParser() antlr.Parser { return s.parser }

func (s *Read2Context) Read3() IRead3Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRead3Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRead3Context)
}

func (s *Read2Context) Vars() IVarsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarsContext)
}

func (s *Read2Context) VarArray() IVarArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarArrayContext)
}

func (s *Read2Context) VarMat() IVarMatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarMatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarMatContext)
}

func (s *Read2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Read2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Read2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterRead2(s)
	}
}

func (s *Read2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitRead2(s)
	}
}

func (p *BruenParser) Read2() (localctx IRead2Context) {
	localctx = NewRead2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, BruenParserRULE_read2)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(430)
			p.Vars()
		}

	case 2:
		{
			p.SetState(431)
			p.VarArray()
		}

	case 3:
		{
			p.SetState(432)
			p.VarMat()
		}

	}
	{
		p.SetState(435)
		p.Read3()
	}

	return localctx
}

// IRead3Context is an interface to support dynamic dispatch.
type IRead3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRead3Context differentiates from other interfaces.
	IsRead3Context()
}

type Read3Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRead3Context() *Read3Context {
	var p = new(Read3Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_read3
	return p
}

func (*Read3Context) IsRead3Context() {}

func NewRead3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Read3Context {
	var p = new(Read3Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_read3

	return p
}

func (s *Read3Context) GetParser() antlr.Parser { return s.parser }

func (s *Read3Context) Read2() IRead2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRead2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRead2Context)
}

func (s *Read3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Read3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Read3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterRead3(s)
	}
}

func (s *Read3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitRead3(s)
	}
}

func (p *BruenParser) Read3() (localctx IRead3Context) {
	localctx = NewRead3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, BruenParserRULE_read3)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(440)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BruenParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(437)
			p.Match(BruenParserT__7)
		}
		{
			p.SetState(438)
			p.Read2()
		}

	case BruenParserRPAREN:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IWriteContext is an interface to support dynamic dispatch.
type IWriteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWriteContext differentiates from other interfaces.
	IsWriteContext()
}

type WriteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWriteContext() *WriteContext {
	var p = new(WriteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_write
	return p
}

func (*WriteContext) IsWriteContext() {}

func NewWriteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WriteContext {
	var p = new(WriteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_write

	return p
}

func (s *WriteContext) GetParser() antlr.Parser { return s.parser }

func (s *WriteContext) WRITE() antlr.TerminalNode {
	return s.GetToken(BruenParserWRITE, 0)
}

func (s *WriteContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserLPAREN, 0)
}

func (s *WriteContext) W_arguments() IW_argumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IW_argumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IW_argumentsContext)
}

func (s *WriteContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserRPAREN, 0)
}

func (s *WriteContext) SEMI() antlr.TerminalNode {
	return s.GetToken(BruenParserSEMI, 0)
}

func (s *WriteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WriteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WriteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterWrite(s)
	}
}

func (s *WriteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitWrite(s)
	}
}

func (p *BruenParser) Write() (localctx IWriteContext) {
	localctx = NewWriteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, BruenParserRULE_write)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(442)
		p.Match(BruenParserWRITE)
	}
	{
		p.SetState(443)
		p.Match(BruenParserLPAREN)
	}
	{
		p.SetState(444)
		p.W_arguments()
	}
	{
		p.SetState(445)
		p.Match(BruenParserRPAREN)
	}
	{
		p.SetState(446)
		p.Match(BruenParserSEMI)
	}

	return localctx
}

// IW_argumentsContext is an interface to support dynamic dispatch.
type IW_argumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsW_argumentsContext differentiates from other interfaces.
	IsW_argumentsContext()
}

type W_argumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyW_argumentsContext() *W_argumentsContext {
	var p = new(W_argumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_w_arguments
	return p
}

func (*W_argumentsContext) IsW_argumentsContext() {}

func NewW_argumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *W_argumentsContext {
	var p = new(W_argumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_w_arguments

	return p
}

func (s *W_argumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *W_argumentsContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *W_argumentsContext) W_arguments2() IW_arguments2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IW_arguments2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IW_arguments2Context)
}

func (s *W_argumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *W_argumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *W_argumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterW_arguments(s)
	}
}

func (s *W_argumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitW_arguments(s)
	}
}

func (p *BruenParser) W_arguments() (localctx IW_argumentsContext) {
	localctx = NewW_argumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, BruenParserRULE_w_arguments)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(448)
		p.Exp()
	}
	{
		p.SetState(449)
		p.W_arguments2()
	}

	return localctx
}

// IW_arguments2Context is an interface to support dynamic dispatch.
type IW_arguments2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsW_arguments2Context differentiates from other interfaces.
	IsW_arguments2Context()
}

type W_arguments2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyW_arguments2Context() *W_arguments2Context {
	var p = new(W_arguments2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_w_arguments2
	return p
}

func (*W_arguments2Context) IsW_arguments2Context() {}

func NewW_arguments2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *W_arguments2Context {
	var p = new(W_arguments2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_w_arguments2

	return p
}

func (s *W_arguments2Context) GetParser() antlr.Parser { return s.parser }

func (s *W_arguments2Context) W_arguments() IW_argumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IW_argumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IW_argumentsContext)
}

func (s *W_arguments2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *W_arguments2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *W_arguments2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterW_arguments2(s)
	}
}

func (s *W_arguments2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitW_arguments2(s)
	}
}

func (p *BruenParser) W_arguments2() (localctx IW_arguments2Context) {
	localctx = NewW_arguments2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, BruenParserRULE_w_arguments2)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(454)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BruenParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(451)
			p.Match(BruenParserT__7)
		}
		{
			p.SetState(452)
			p.W_arguments()
		}

	case BruenParserRPAREN:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConditionalContext is an interface to support dynamic dispatch.
type IConditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionalContext differentiates from other interfaces.
	IsConditionalContext()
}

type ConditionalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalContext() *ConditionalContext {
	var p = new(ConditionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_conditional
	return p
}

func (*ConditionalContext) IsConditionalContext() {}

func NewConditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalContext {
	var p = new(ConditionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_conditional

	return p
}

func (s *ConditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalContext) IF() antlr.TerminalNode {
	return s.GetToken(BruenParserIF, 0)
}

func (s *ConditionalContext) Conditional2() IConditional2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditional2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditional2Context)
}

func (s *ConditionalContext) Conditional3() IConditional3Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditional3Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditional3Context)
}

func (s *ConditionalContext) Conditional4() IConditional4Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditional4Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditional4Context)
}

func (s *ConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterConditional(s)
	}
}

func (s *ConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitConditional(s)
	}
}

func (p *BruenParser) Conditional() (localctx IConditionalContext) {
	localctx = NewConditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, BruenParserRULE_conditional)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(456)
		p.Match(BruenParserIF)
	}
	{
		p.SetState(457)
		p.Conditional2()
	}
	{
		p.SetState(458)
		p.Conditional3()
	}
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == BruenParserELSE {
		{
			p.SetState(459)
			p.Conditional4()
		}

	}

	return localctx
}

// IConditional2Context is an interface to support dynamic dispatch.
type IConditional2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditional2Context differentiates from other interfaces.
	IsConditional2Context()
}

type Conditional2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditional2Context() *Conditional2Context {
	var p = new(Conditional2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_conditional2
	return p
}

func (*Conditional2Context) IsConditional2Context() {}

func NewConditional2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Conditional2Context {
	var p = new(Conditional2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_conditional2

	return p
}

func (s *Conditional2Context) GetParser() antlr.Parser { return s.parser }

func (s *Conditional2Context) LPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserLPAREN, 0)
}

func (s *Conditional2Context) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *Conditional2Context) RPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserRPAREN, 0)
}

func (s *Conditional2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Conditional2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Conditional2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterConditional2(s)
	}
}

func (s *Conditional2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitConditional2(s)
	}
}

func (p *BruenParser) Conditional2() (localctx IConditional2Context) {
	localctx = NewConditional2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, BruenParserRULE_conditional2)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(462)
		p.Match(BruenParserLPAREN)
	}
	{
		p.SetState(463)
		p.Exp()
	}
	{
		p.SetState(464)
		p.Match(BruenParserRPAREN)
	}

	return localctx
}

// IConditional3Context is an interface to support dynamic dispatch.
type IConditional3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditional3Context differentiates from other interfaces.
	IsConditional3Context()
}

type Conditional3Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditional3Context() *Conditional3Context {
	var p = new(Conditional3Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_conditional3
	return p
}

func (*Conditional3Context) IsConditional3Context() {}

func NewConditional3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Conditional3Context {
	var p = new(Conditional3Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_conditional3

	return p
}

func (s *Conditional3Context) GetParser() antlr.Parser { return s.parser }

func (s *Conditional3Context) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Conditional3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Conditional3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Conditional3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterConditional3(s)
	}
}

func (s *Conditional3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitConditional3(s)
	}
}

func (p *BruenParser) Conditional3() (localctx IConditional3Context) {
	localctx = NewConditional3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, BruenParserRULE_conditional3)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(466)
		p.Block()
	}

	return localctx
}

// IConditional4Context is an interface to support dynamic dispatch.
type IConditional4Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditional4Context differentiates from other interfaces.
	IsConditional4Context()
}

type Conditional4Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditional4Context() *Conditional4Context {
	var p = new(Conditional4Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_conditional4
	return p
}

func (*Conditional4Context) IsConditional4Context() {}

func NewConditional4Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Conditional4Context {
	var p = new(Conditional4Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_conditional4

	return p
}

func (s *Conditional4Context) GetParser() antlr.Parser { return s.parser }

func (s *Conditional4Context) ELSE() antlr.TerminalNode {
	return s.GetToken(BruenParserELSE, 0)
}

func (s *Conditional4Context) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Conditional4Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Conditional4Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Conditional4Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterConditional4(s)
	}
}

func (s *Conditional4Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitConditional4(s)
	}
}

func (p *BruenParser) Conditional4() (localctx IConditional4Context) {
	localctx = NewConditional4Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, BruenParserRULE_conditional4)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(468)
		p.Match(BruenParserELSE)
	}
	{
		p.SetState(469)
		p.Block()
	}

	return localctx
}

// IForLoopContext is an interface to support dynamic dispatch.
type IForLoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForLoopContext differentiates from other interfaces.
	IsForLoopContext()
}

type ForLoopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForLoopContext() *ForLoopContext {
	var p = new(ForLoopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_forLoop
	return p
}

func (*ForLoopContext) IsForLoopContext() {}

func NewForLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForLoopContext {
	var p = new(ForLoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_forLoop

	return p
}

func (s *ForLoopContext) GetParser() antlr.Parser { return s.parser }

func (s *ForLoopContext) FOR() antlr.TerminalNode {
	return s.GetToken(BruenParserFOR, 0)
}

func (s *ForLoopContext) ForLoop2() IForLoop2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForLoop2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForLoop2Context)
}

func (s *ForLoopContext) IN() antlr.TerminalNode {
	return s.GetToken(BruenParserIN, 0)
}

func (s *ForLoopContext) ForLoop3() IForLoop3Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForLoop3Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForLoop3Context)
}

func (s *ForLoopContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForLoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForLoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForLoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterForLoop(s)
	}
}

func (s *ForLoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitForLoop(s)
	}
}

func (p *BruenParser) ForLoop() (localctx IForLoopContext) {
	localctx = NewForLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, BruenParserRULE_forLoop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(471)
		p.Match(BruenParserFOR)
	}
	{
		p.SetState(472)
		p.ForLoop2()
	}
	{
		p.SetState(473)
		p.Match(BruenParserIN)
	}
	{
		p.SetState(474)
		p.ForLoop3()
	}
	{
		p.SetState(475)
		p.Block()
	}

	return localctx
}

// IForLoop2Context is an interface to support dynamic dispatch.
type IForLoop2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForLoop2Context differentiates from other interfaces.
	IsForLoop2Context()
}

type ForLoop2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForLoop2Context() *ForLoop2Context {
	var p = new(ForLoop2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_forLoop2
	return p
}

func (*ForLoop2Context) IsForLoop2Context() {}

func NewForLoop2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForLoop2Context {
	var p = new(ForLoop2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_forLoop2

	return p
}

func (s *ForLoop2Context) GetParser() antlr.Parser { return s.parser }

func (s *ForLoop2Context) ID() antlr.TerminalNode {
	return s.GetToken(BruenParserID, 0)
}

func (s *ForLoop2Context) ASSIGN() antlr.TerminalNode {
	return s.GetToken(BruenParserASSIGN, 0)
}

func (s *ForLoop2Context) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ForLoop2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForLoop2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForLoop2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterForLoop2(s)
	}
}

func (s *ForLoop2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitForLoop2(s)
	}
}

func (p *BruenParser) ForLoop2() (localctx IForLoop2Context) {
	localctx = NewForLoop2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, BruenParserRULE_forLoop2)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(477)
		p.Match(BruenParserID)
	}
	{
		p.SetState(478)
		p.Match(BruenParserASSIGN)
	}
	{
		p.SetState(479)
		p.Exp()
	}

	return localctx
}

// IForLoop3Context is an interface to support dynamic dispatch.
type IForLoop3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForLoop3Context differentiates from other interfaces.
	IsForLoop3Context()
}

type ForLoop3Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForLoop3Context() *ForLoop3Context {
	var p = new(ForLoop3Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_forLoop3
	return p
}

func (*ForLoop3Context) IsForLoop3Context() {}

func NewForLoop3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForLoop3Context {
	var p = new(ForLoop3Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_forLoop3

	return p
}

func (s *ForLoop3Context) GetParser() antlr.Parser { return s.parser }

func (s *ForLoop3Context) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ForLoop3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForLoop3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForLoop3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterForLoop3(s)
	}
}

func (s *ForLoop3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitForLoop3(s)
	}
}

func (p *BruenParser) ForLoop3() (localctx IForLoop3Context) {
	localctx = NewForLoop3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, BruenParserRULE_forLoop3)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(481)
		p.Exp()
	}

	return localctx
}

// IWhileLoopContext is an interface to support dynamic dispatch.
type IWhileLoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileLoopContext differentiates from other interfaces.
	IsWhileLoopContext()
}

type WhileLoopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileLoopContext() *WhileLoopContext {
	var p = new(WhileLoopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_whileLoop
	return p
}

func (*WhileLoopContext) IsWhileLoopContext() {}

func NewWhileLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileLoopContext {
	var p = new(WhileLoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_whileLoop

	return p
}

func (s *WhileLoopContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileLoopContext) WHILE() antlr.TerminalNode {
	return s.GetToken(BruenParserWHILE, 0)
}

func (s *WhileLoopContext) WhileLoop2() IWhileLoop2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileLoop2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhileLoop2Context)
}

func (s *WhileLoopContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileLoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileLoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileLoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterWhileLoop(s)
	}
}

func (s *WhileLoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitWhileLoop(s)
	}
}

func (p *BruenParser) WhileLoop() (localctx IWhileLoopContext) {
	localctx = NewWhileLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, BruenParserRULE_whileLoop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(483)
		p.Match(BruenParserWHILE)
	}
	{
		p.SetState(484)
		p.WhileLoop2()
	}
	{
		p.SetState(485)
		p.Block()
	}

	return localctx
}

// IWhileLoop2Context is an interface to support dynamic dispatch.
type IWhileLoop2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileLoop2Context differentiates from other interfaces.
	IsWhileLoop2Context()
}

type WhileLoop2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileLoop2Context() *WhileLoop2Context {
	var p = new(WhileLoop2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_whileLoop2
	return p
}

func (*WhileLoop2Context) IsWhileLoop2Context() {}

func NewWhileLoop2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileLoop2Context {
	var p = new(WhileLoop2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_whileLoop2

	return p
}

func (s *WhileLoop2Context) GetParser() antlr.Parser { return s.parser }

func (s *WhileLoop2Context) LPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserLPAREN, 0)
}

func (s *WhileLoop2Context) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *WhileLoop2Context) RPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserRPAREN, 0)
}

func (s *WhileLoop2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileLoop2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileLoop2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterWhileLoop2(s)
	}
}

func (s *WhileLoop2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitWhileLoop2(s)
	}
}

func (p *BruenParser) WhileLoop2() (localctx IWhileLoop2Context) {
	localctx = NewWhileLoop2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, BruenParserRULE_whileLoop2)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(487)
		p.Match(BruenParserLPAREN)
	}
	{
		p.SetState(488)
		p.Exp()
	}
	{
		p.SetState(489)
		p.Match(BruenParserRPAREN)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpressionContext) SEMI() antlr.TerminalNode {
	return s.GetToken(BruenParserSEMI, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BruenParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, BruenParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(491)
		p.Exp()
	}
	{
		p.SetState(492)
		p.Match(BruenParserSEMI)
	}

	return localctx
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) T_exp() IT_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IT_expContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IT_expContext)
}

func (s *ExpContext) AllExp2() []IExp2Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExp2Context)(nil)).Elem())
	var tst = make([]IExp2Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExp2Context)
		}
	}

	return tst
}

func (s *ExpContext) Exp2(i int) IExp2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp2Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExp2Context)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitExp(s)
	}
}

func (p *BruenParser) Exp() (localctx IExpContext) {
	localctx = NewExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, BruenParserRULE_exp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(494)
		p.T_exp()
	}
	p.SetState(498)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BruenParserOR {
		{
			p.SetState(495)
			p.Exp2()
		}

		p.SetState(500)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExp2Context is an interface to support dynamic dispatch.
type IExp2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExp2Context differentiates from other interfaces.
	IsExp2Context()
}

type Exp2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExp2Context() *Exp2Context {
	var p = new(Exp2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_exp2
	return p
}

func (*Exp2Context) IsExp2Context() {}

func NewExp2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp2Context {
	var p = new(Exp2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_exp2

	return p
}

func (s *Exp2Context) GetParser() antlr.Parser { return s.parser }

func (s *Exp2Context) OR() antlr.TerminalNode {
	return s.GetToken(BruenParserOR, 0)
}

func (s *Exp2Context) T_exp() IT_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IT_expContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IT_expContext)
}

func (s *Exp2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterExp2(s)
	}
}

func (s *Exp2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitExp2(s)
	}
}

func (p *BruenParser) Exp2() (localctx IExp2Context) {
	localctx = NewExp2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, BruenParserRULE_exp2)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(501)
		p.Match(BruenParserOR)
	}
	{
		p.SetState(502)
		p.T_exp()
	}

	return localctx
}

// IT_expContext is an interface to support dynamic dispatch.
type IT_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsT_expContext differentiates from other interfaces.
	IsT_expContext()
}

type T_expContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyT_expContext() *T_expContext {
	var p = new(T_expContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_t_exp
	return p
}

func (*T_expContext) IsT_expContext() {}

func NewT_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *T_expContext {
	var p = new(T_expContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_t_exp

	return p
}

func (s *T_expContext) GetParser() antlr.Parser { return s.parser }

func (s *T_expContext) G_exp() IG_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IG_expContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IG_expContext)
}

func (s *T_expContext) AllT_exp2() []IT_exp2Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IT_exp2Context)(nil)).Elem())
	var tst = make([]IT_exp2Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IT_exp2Context)
		}
	}

	return tst
}

func (s *T_expContext) T_exp2(i int) IT_exp2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IT_exp2Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IT_exp2Context)
}

func (s *T_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *T_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *T_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterT_exp(s)
	}
}

func (s *T_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitT_exp(s)
	}
}

func (p *BruenParser) T_exp() (localctx IT_expContext) {
	localctx = NewT_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, BruenParserRULE_t_exp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(504)
		p.G_exp()
	}
	p.SetState(508)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BruenParserAND {
		{
			p.SetState(505)
			p.T_exp2()
		}

		p.SetState(510)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IT_exp2Context is an interface to support dynamic dispatch.
type IT_exp2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsT_exp2Context differentiates from other interfaces.
	IsT_exp2Context()
}

type T_exp2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyT_exp2Context() *T_exp2Context {
	var p = new(T_exp2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_t_exp2
	return p
}

func (*T_exp2Context) IsT_exp2Context() {}

func NewT_exp2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *T_exp2Context {
	var p = new(T_exp2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_t_exp2

	return p
}

func (s *T_exp2Context) GetParser() antlr.Parser { return s.parser }

func (s *T_exp2Context) AND() antlr.TerminalNode {
	return s.GetToken(BruenParserAND, 0)
}

func (s *T_exp2Context) G_exp() IG_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IG_expContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IG_expContext)
}

func (s *T_exp2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *T_exp2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *T_exp2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterT_exp2(s)
	}
}

func (s *T_exp2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitT_exp2(s)
	}
}

func (p *BruenParser) T_exp2() (localctx IT_exp2Context) {
	localctx = NewT_exp2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, BruenParserRULE_t_exp2)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(511)
		p.Match(BruenParserAND)
	}
	{
		p.SetState(512)
		p.G_exp()
	}

	return localctx
}

// IG_expContext is an interface to support dynamic dispatch.
type IG_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsG_expContext differentiates from other interfaces.
	IsG_expContext()
}

type G_expContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyG_expContext() *G_expContext {
	var p = new(G_expContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_g_exp
	return p
}

func (*G_expContext) IsG_expContext() {}

func NewG_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *G_expContext {
	var p = new(G_expContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_g_exp

	return p
}

func (s *G_expContext) GetParser() antlr.Parser { return s.parser }

func (s *G_expContext) M_exp() IM_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IM_expContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IM_expContext)
}

func (s *G_expContext) G_exp2() IG_exp2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IG_exp2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IG_exp2Context)
}

func (s *G_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *G_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *G_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterG_exp(s)
	}
}

func (s *G_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitG_exp(s)
	}
}

func (p *BruenParser) G_exp() (localctx IG_expContext) {
	localctx = NewG_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, BruenParserRULE_g_exp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(514)
		p.M_exp()
	}
	p.SetState(516)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BruenParserGT)|(1<<BruenParserLT)|(1<<BruenParserEQ)|(1<<BruenParserNEQ))) != 0 {
		{
			p.SetState(515)
			p.G_exp2()
		}

	}

	return localctx
}

// IG_exp2Context is an interface to support dynamic dispatch.
type IG_exp2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsG_exp2Context differentiates from other interfaces.
	IsG_exp2Context()
}

type G_exp2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyG_exp2Context() *G_exp2Context {
	var p = new(G_exp2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_g_exp2
	return p
}

func (*G_exp2Context) IsG_exp2Context() {}

func NewG_exp2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *G_exp2Context {
	var p = new(G_exp2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_g_exp2

	return p
}

func (s *G_exp2Context) GetParser() antlr.Parser { return s.parser }

func (s *G_exp2Context) Relop() IRelopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelopContext)
}

func (s *G_exp2Context) M_exp() IM_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IM_expContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IM_expContext)
}

func (s *G_exp2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *G_exp2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *G_exp2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterG_exp2(s)
	}
}

func (s *G_exp2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitG_exp2(s)
	}
}

func (p *BruenParser) G_exp2() (localctx IG_exp2Context) {
	localctx = NewG_exp2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, BruenParserRULE_g_exp2)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(518)
		p.Relop()
	}
	{
		p.SetState(519)
		p.M_exp()
	}

	return localctx
}

// IM_expContext is an interface to support dynamic dispatch.
type IM_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsM_expContext differentiates from other interfaces.
	IsM_expContext()
}

type M_expContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyM_expContext() *M_expContext {
	var p = new(M_expContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_m_exp
	return p
}

func (*M_expContext) IsM_expContext() {}

func NewM_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *M_expContext {
	var p = new(M_expContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_m_exp

	return p
}

func (s *M_expContext) GetParser() antlr.Parser { return s.parser }

func (s *M_expContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *M_expContext) AllM_exp2() []IM_exp2Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IM_exp2Context)(nil)).Elem())
	var tst = make([]IM_exp2Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IM_exp2Context)
		}
	}

	return tst
}

func (s *M_expContext) M_exp2(i int) IM_exp2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IM_exp2Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IM_exp2Context)
}

func (s *M_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *M_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *M_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterM_exp(s)
	}
}

func (s *M_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitM_exp(s)
	}
}

func (p *BruenParser) M_exp() (localctx IM_expContext) {
	localctx = NewM_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, BruenParserRULE_m_exp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(521)
		p.Term()
	}
	p.SetState(525)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BruenParserADD || _la == BruenParserSUB {
		{
			p.SetState(522)
			p.M_exp2()
		}

		p.SetState(527)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IM_exp2Context is an interface to support dynamic dispatch.
type IM_exp2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsM_exp2Context differentiates from other interfaces.
	IsM_exp2Context()
}

type M_exp2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyM_exp2Context() *M_exp2Context {
	var p = new(M_exp2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_m_exp2
	return p
}

func (*M_exp2Context) IsM_exp2Context() {}

func NewM_exp2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *M_exp2Context {
	var p = new(M_exp2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_m_exp2

	return p
}

func (s *M_exp2Context) GetParser() antlr.Parser { return s.parser }

func (s *M_exp2Context) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *M_exp2Context) ADD() antlr.TerminalNode {
	return s.GetToken(BruenParserADD, 0)
}

func (s *M_exp2Context) SUB() antlr.TerminalNode {
	return s.GetToken(BruenParserSUB, 0)
}

func (s *M_exp2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *M_exp2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *M_exp2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterM_exp2(s)
	}
}

func (s *M_exp2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitM_exp2(s)
	}
}

func (p *BruenParser) M_exp2() (localctx IM_exp2Context) {
	localctx = NewM_exp2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, BruenParserRULE_m_exp2)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(528)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BruenParserADD || _la == BruenParserSUB) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(529)
		p.Term()
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Factor() IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContext) AllTerm2() []ITerm2Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITerm2Context)(nil)).Elem())
	var tst = make([]ITerm2Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITerm2Context)
		}
	}

	return tst
}

func (s *TermContext) Term2(i int) ITerm2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerm2Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITerm2Context)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *BruenParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, BruenParserRULE_term)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(531)
		p.Factor()
	}
	p.SetState(535)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BruenParserMUL || _la == BruenParserDIV {
		{
			p.SetState(532)
			p.Term2()
		}

		p.SetState(537)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITerm2Context is an interface to support dynamic dispatch.
type ITerm2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTerm2Context differentiates from other interfaces.
	IsTerm2Context()
}

type Term2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerm2Context() *Term2Context {
	var p = new(Term2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_term2
	return p
}

func (*Term2Context) IsTerm2Context() {}

func NewTerm2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Term2Context {
	var p = new(Term2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_term2

	return p
}

func (s *Term2Context) GetParser() antlr.Parser { return s.parser }

func (s *Term2Context) Factor() IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *Term2Context) MUL() antlr.TerminalNode {
	return s.GetToken(BruenParserMUL, 0)
}

func (s *Term2Context) DIV() antlr.TerminalNode {
	return s.GetToken(BruenParserDIV, 0)
}

func (s *Term2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Term2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Term2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterTerm2(s)
	}
}

func (s *Term2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitTerm2(s)
	}
}

func (p *BruenParser) Term2() (localctx ITerm2Context) {
	localctx = NewTerm2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, BruenParserRULE_term2)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(538)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BruenParserMUL || _la == BruenParserDIV) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(539)
		p.Factor()
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) Factor2() IFactor2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactor2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFactor2Context)
}

func (s *FactorContext) VarCte() IVarCteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarCteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarCteContext)
}

func (s *FactorContext) Vars() IVarsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarsContext)
}

func (s *FactorContext) VarArray() IVarArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarArrayContext)
}

func (s *FactorContext) VarMat() IVarMatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarMatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarMatContext)
}

func (s *FactorContext) Call() ICallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *BruenParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, BruenParserRULE_factor)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(549)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(541)
			p.Factor2()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(542)
			p.VarCte()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(546)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(543)
				p.Vars()
			}

		case 2:
			{
				p.SetState(544)
				p.VarArray()
			}

		case 3:
			{
				p.SetState(545)
				p.VarMat()
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(548)
			p.Call()
		}

	}

	return localctx
}

// IFactor2Context is an interface to support dynamic dispatch.
type IFactor2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactor2Context differentiates from other interfaces.
	IsFactor2Context()
}

type Factor2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactor2Context() *Factor2Context {
	var p = new(Factor2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_factor2
	return p
}

func (*Factor2Context) IsFactor2Context() {}

func NewFactor2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Factor2Context {
	var p = new(Factor2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_factor2

	return p
}

func (s *Factor2Context) GetParser() antlr.Parser { return s.parser }

func (s *Factor2Context) LPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserLPAREN, 0)
}

func (s *Factor2Context) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *Factor2Context) RPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserRPAREN, 0)
}

func (s *Factor2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Factor2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Factor2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterFactor2(s)
	}
}

func (s *Factor2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitFactor2(s)
	}
}

func (p *BruenParser) Factor2() (localctx IFactor2Context) {
	localctx = NewFactor2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, BruenParserRULE_factor2)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(551)
		p.Match(BruenParserLPAREN)
	}
	{
		p.SetState(552)
		p.Exp()
	}
	{
		p.SetState(553)
		p.Match(BruenParserRPAREN)
	}

	return localctx
}

// IVarCteContext is an interface to support dynamic dispatch.
type IVarCteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarCteContext differentiates from other interfaces.
	IsVarCteContext()
}

type VarCteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarCteContext() *VarCteContext {
	var p = new(VarCteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_varCte
	return p
}

func (*VarCteContext) IsVarCteContext() {}

func NewVarCteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarCteContext {
	var p = new(VarCteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_varCte

	return p
}

func (s *VarCteContext) GetParser() antlr.Parser { return s.parser }

func (s *VarCteContext) Cte_i() ICte_iContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICte_iContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICte_iContext)
}

func (s *VarCteContext) Cte_f() ICte_fContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICte_fContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICte_fContext)
}

func (s *VarCteContext) Cte_c() ICte_cContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICte_cContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICte_cContext)
}

func (s *VarCteContext) Cte_s() ICte_sContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICte_sContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICte_sContext)
}

func (s *VarCteContext) Cte_b() ICte_bContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICte_bContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICte_bContext)
}

func (s *VarCteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarCteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarCteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterVarCte(s)
	}
}

func (s *VarCteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitVarCte(s)
	}
}

func (p *BruenParser) VarCte() (localctx IVarCteContext) {
	localctx = NewVarCteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, BruenParserRULE_varCte)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(560)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BruenParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(555)
			p.Cte_i()
		}

	case BruenParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(556)
			p.Cte_f()
		}

	case BruenParserCHAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(557)
			p.Cte_c()
		}

	case BruenParserT__9:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(558)
			p.Cte_s()
		}

	case BruenParserBOOL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(559)
			p.Cte_b()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICte_iContext is an interface to support dynamic dispatch.
type ICte_iContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCte_iContext differentiates from other interfaces.
	IsCte_iContext()
}

type Cte_iContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCte_iContext() *Cte_iContext {
	var p = new(Cte_iContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_cte_i
	return p
}

func (*Cte_iContext) IsCte_iContext() {}

func NewCte_iContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cte_iContext {
	var p = new(Cte_iContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_cte_i

	return p
}

func (s *Cte_iContext) GetParser() antlr.Parser { return s.parser }

func (s *Cte_iContext) INT() antlr.TerminalNode {
	return s.GetToken(BruenParserINT, 0)
}

func (s *Cte_iContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cte_iContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cte_iContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterCte_i(s)
	}
}

func (s *Cte_iContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitCte_i(s)
	}
}

func (p *BruenParser) Cte_i() (localctx ICte_iContext) {
	localctx = NewCte_iContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, BruenParserRULE_cte_i)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(562)
		p.Match(BruenParserINT)
	}

	return localctx
}

// ICte_fContext is an interface to support dynamic dispatch.
type ICte_fContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCte_fContext differentiates from other interfaces.
	IsCte_fContext()
}

type Cte_fContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCte_fContext() *Cte_fContext {
	var p = new(Cte_fContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_cte_f
	return p
}

func (*Cte_fContext) IsCte_fContext() {}

func NewCte_fContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cte_fContext {
	var p = new(Cte_fContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_cte_f

	return p
}

func (s *Cte_fContext) GetParser() antlr.Parser { return s.parser }

func (s *Cte_fContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(BruenParserFLOAT, 0)
}

func (s *Cte_fContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cte_fContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cte_fContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterCte_f(s)
	}
}

func (s *Cte_fContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitCte_f(s)
	}
}

func (p *BruenParser) Cte_f() (localctx ICte_fContext) {
	localctx = NewCte_fContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, BruenParserRULE_cte_f)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(564)
		p.Match(BruenParserFLOAT)
	}

	return localctx
}

// ICte_cContext is an interface to support dynamic dispatch.
type ICte_cContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCte_cContext differentiates from other interfaces.
	IsCte_cContext()
}

type Cte_cContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCte_cContext() *Cte_cContext {
	var p = new(Cte_cContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_cte_c
	return p
}

func (*Cte_cContext) IsCte_cContext() {}

func NewCte_cContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cte_cContext {
	var p = new(Cte_cContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_cte_c

	return p
}

func (s *Cte_cContext) GetParser() antlr.Parser { return s.parser }

func (s *Cte_cContext) CHAR() antlr.TerminalNode {
	return s.GetToken(BruenParserCHAR, 0)
}

func (s *Cte_cContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cte_cContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cte_cContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterCte_c(s)
	}
}

func (s *Cte_cContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitCte_c(s)
	}
}

func (p *BruenParser) Cte_c() (localctx ICte_cContext) {
	localctx = NewCte_cContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, BruenParserRULE_cte_c)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(566)
		p.Match(BruenParserCHAR)
	}

	return localctx
}

// ICte_bContext is an interface to support dynamic dispatch.
type ICte_bContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCte_bContext differentiates from other interfaces.
	IsCte_bContext()
}

type Cte_bContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCte_bContext() *Cte_bContext {
	var p = new(Cte_bContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_cte_b
	return p
}

func (*Cte_bContext) IsCte_bContext() {}

func NewCte_bContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cte_bContext {
	var p = new(Cte_bContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_cte_b

	return p
}

func (s *Cte_bContext) GetParser() antlr.Parser { return s.parser }

func (s *Cte_bContext) BOOL() antlr.TerminalNode {
	return s.GetToken(BruenParserBOOL, 0)
}

func (s *Cte_bContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cte_bContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cte_bContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterCte_b(s)
	}
}

func (s *Cte_bContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitCte_b(s)
	}
}

func (p *BruenParser) Cte_b() (localctx ICte_bContext) {
	localctx = NewCte_bContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, BruenParserRULE_cte_b)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(568)
		p.Match(BruenParserBOOL)
	}

	return localctx
}

// ICte_sContext is an interface to support dynamic dispatch.
type ICte_sContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCte_sContext differentiates from other interfaces.
	IsCte_sContext()
}

type Cte_sContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCte_sContext() *Cte_sContext {
	var p = new(Cte_sContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_cte_s
	return p
}

func (*Cte_sContext) IsCte_sContext() {}

func NewCte_sContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cte_sContext {
	var p = new(Cte_sContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_cte_s

	return p
}

func (s *Cte_sContext) GetParser() antlr.Parser { return s.parser }
func (s *Cte_sContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cte_sContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cte_sContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterCte_s(s)
	}
}

func (s *Cte_sContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitCte_s(s)
	}
}

func (p *BruenParser) Cte_s() (localctx ICte_sContext) {
	localctx = NewCte_sContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, BruenParserRULE_cte_s)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(570)
		p.Match(BruenParserT__9)
	}
	p.SetState(574)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(571)
			p.MatchWildcard()

		}
		p.SetState(576)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())
	}
	{
		p.SetState(577)
		p.Match(BruenParserT__9)
	}

	return localctx
}

// IMainContext is an interface to support dynamic dispatch.
type IMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMainContext differentiates from other interfaces.
	IsMainContext()
}

type MainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainContext() *MainContext {
	var p = new(MainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_main
	return p
}

func (*MainContext) IsMainContext() {}

func NewMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainContext {
	var p = new(MainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_main

	return p
}

func (s *MainContext) GetParser() antlr.Parser { return s.parser }

func (s *MainContext) MAIN() antlr.TerminalNode {
	return s.GetToken(BruenParserMAIN, 0)
}

func (s *MainContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserLPAREN, 0)
}

func (s *MainContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BruenParserRPAREN, 0)
}

func (s *MainContext) MainBlock() IMainBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMainBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMainBlockContext)
}

func (s *MainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterMain(s)
	}
}

func (s *MainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitMain(s)
	}
}

func (p *BruenParser) Main() (localctx IMainContext) {
	localctx = NewMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, BruenParserRULE_main)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(579)
		p.Match(BruenParserMAIN)
	}
	{
		p.SetState(580)
		p.Match(BruenParserLPAREN)
	}
	{
		p.SetState(581)
		p.Match(BruenParserRPAREN)
	}
	{
		p.SetState(582)
		p.MainBlock()
	}

	return localctx
}

// IMainBlockContext is an interface to support dynamic dispatch.
type IMainBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMainBlockContext differentiates from other interfaces.
	IsMainBlockContext()
}

type MainBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainBlockContext() *MainBlockContext {
	var p = new(MainBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_mainBlock
	return p
}

func (*MainBlockContext) IsMainBlockContext() {}

func NewMainBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainBlockContext {
	var p = new(MainBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_mainBlock

	return p
}

func (s *MainBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *MainBlockContext) AllVariableDeclaration() []IVariableDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem())
	var tst = make([]IVariableDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDeclarationContext)
		}
	}

	return tst
}

func (s *MainBlockContext) VariableDeclaration(i int) IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *MainBlockContext) AllStatutes() []IStatutesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatutesContext)(nil)).Elem())
	var tst = make([]IStatutesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatutesContext)
		}
	}

	return tst
}

func (s *MainBlockContext) Statutes(i int) IStatutesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatutesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatutesContext)
}

func (s *MainBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterMainBlock(s)
	}
}

func (s *MainBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitMainBlock(s)
	}
}

func (p *BruenParser) MainBlock() (localctx IMainBlockContext) {
	localctx = NewMainBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, BruenParserRULE_mainBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(584)
		p.Match(BruenParserT__0)
	}
	p.SetState(588)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BruenParserVAR {
		{
			p.SetState(585)
			p.VariableDeclaration()
		}

		p.SetState(590)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(594)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BruenParserT__8)|(1<<BruenParserT__9)|(1<<BruenParserINT)|(1<<BruenParserFLOAT)|(1<<BruenParserCHAR)|(1<<BruenParserBOOL)|(1<<BruenParserLPAREN)|(1<<BruenParserIF))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BruenParserWHILE-32))|(1<<(BruenParserFOR-32))|(1<<(BruenParserWRITE-32))|(1<<(BruenParserREAD-32))|(1<<(BruenParserRETURN-32))|(1<<(BruenParserID-32)))) != 0) {
		{
			p.SetState(591)
			p.Statutes()
		}

		p.SetState(596)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(597)
		p.Match(BruenParserT__1)
	}

	return localctx
}

// ITypeRuleContext is an interface to support dynamic dispatch.
type ITypeRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeRuleContext differentiates from other interfaces.
	IsTypeRuleContext()
}

type TypeRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeRuleContext() *TypeRuleContext {
	var p = new(TypeRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_typeRule
	return p
}

func (*TypeRuleContext) IsTypeRuleContext() {}

func NewTypeRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeRuleContext {
	var p = new(TypeRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_typeRule

	return p
}

func (s *TypeRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeRuleContext) INT_TYPE() antlr.TerminalNode {
	return s.GetToken(BruenParserINT_TYPE, 0)
}

func (s *TypeRuleContext) FLOAT_TYPE() antlr.TerminalNode {
	return s.GetToken(BruenParserFLOAT_TYPE, 0)
}

func (s *TypeRuleContext) CHAR_TYPE() antlr.TerminalNode {
	return s.GetToken(BruenParserCHAR_TYPE, 0)
}

func (s *TypeRuleContext) BOOL_TYPE() antlr.TerminalNode {
	return s.GetToken(BruenParserBOOL_TYPE, 0)
}

func (s *TypeRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterTypeRule(s)
	}
}

func (s *TypeRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitTypeRule(s)
	}
}

func (p *BruenParser) TypeRule() (localctx ITypeRuleContext) {
	localctx = NewTypeRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, BruenParserRULE_typeRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(599)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(BruenParserINT_TYPE-47))|(1<<(BruenParserFLOAT_TYPE-47))|(1<<(BruenParserCHAR_TYPE-47))|(1<<(BruenParserBOOL_TYPE-47)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRelopContext is an interface to support dynamic dispatch.
type IRelopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelopContext differentiates from other interfaces.
	IsRelopContext()
}

type RelopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelopContext() *RelopContext {
	var p = new(RelopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BruenParserRULE_relop
	return p
}

func (*RelopContext) IsRelopContext() {}

func NewRelopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelopContext {
	var p = new(RelopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BruenParserRULE_relop

	return p
}

func (s *RelopContext) GetParser() antlr.Parser { return s.parser }

func (s *RelopContext) GT() antlr.TerminalNode {
	return s.GetToken(BruenParserGT, 0)
}

func (s *RelopContext) LT() antlr.TerminalNode {
	return s.GetToken(BruenParserLT, 0)
}

func (s *RelopContext) EQ() antlr.TerminalNode {
	return s.GetToken(BruenParserEQ, 0)
}

func (s *RelopContext) NEQ() antlr.TerminalNode {
	return s.GetToken(BruenParserNEQ, 0)
}

func (s *RelopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.EnterRelop(s)
	}
}

func (s *RelopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BruenListener); ok {
		listenerT.ExitRelop(s)
	}
}

func (p *BruenParser) Relop() (localctx IRelopContext) {
	localctx = NewRelopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, BruenParserRULE_relop)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(601)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BruenParserGT)|(1<<BruenParserLT)|(1<<BruenParserEQ)|(1<<BruenParserNEQ))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
