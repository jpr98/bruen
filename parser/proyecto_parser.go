// Code generated from Proyecto.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Proyecto

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 53, 586,
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
	4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 3, 2, 3, 2, 3,
	2, 3, 2, 7, 2, 143, 10, 2, 12, 2, 14, 2, 146, 11, 2, 3, 2, 3, 2, 3, 3,
	7, 3, 151, 10, 3, 12, 3, 14, 3, 154, 11, 3, 3, 3, 7, 3, 157, 10, 3, 12,
	3, 14, 3, 160, 11, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5,
	4, 170, 10, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 180,
	10, 5, 12, 5, 14, 5, 183, 11, 5, 3, 5, 3, 5, 3, 6, 3, 6, 7, 6, 189, 10,
	6, 12, 6, 14, 6, 192, 11, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 199, 10,
	7, 12, 7, 14, 7, 202, 11, 7, 3, 7, 7, 7, 205, 10, 7, 12, 7, 14, 7, 208,
	11, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 5, 8, 215, 10, 8, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 5, 12, 246, 10, 12, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 5, 13, 253, 10, 13, 3, 13, 3, 13, 3, 14, 3, 14, 5, 14, 259,
	10, 14, 3, 14, 5, 14, 262, 10, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	16, 5, 16, 270, 10, 16, 3, 17, 3, 17, 3, 17, 5, 17, 275, 10, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 5, 18, 284, 10, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 5, 18, 290, 10, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18,
	296, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 302, 10, 19, 3, 19, 3,
	19, 3, 19, 5, 19, 307, 10, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 7, 20,
	314, 10, 20, 12, 20, 14, 20, 317, 11, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	22, 3, 22, 7, 22, 325, 10, 22, 12, 22, 14, 22, 328, 11, 22, 3, 22, 7, 22,
	331, 10, 22, 12, 22, 14, 22, 334, 11, 22, 3, 22, 3, 22, 3, 23, 3, 23, 5,
	23, 340, 10, 23, 3, 23, 3, 23, 3, 24, 3, 24, 7, 24, 346, 10, 24, 12, 24,
	14, 24, 349, 11, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	25, 3, 25, 5, 25, 360, 10, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 5, 26, 370, 10, 26, 3, 27, 3, 27, 3, 27, 5, 27, 375, 10,
	27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 5, 28, 384, 10, 28,
	3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 5, 30, 394, 10,
	30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 401, 10, 31, 3, 31, 3, 31,
	3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 5, 33, 413, 10,
	33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 5, 35,
	424, 10, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 5, 36, 431, 10, 36, 3,
	37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 5, 38,
	443, 10, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3,
	41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43,
	3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3,
	47, 3, 47, 3, 47, 3, 48, 3, 48, 7, 48, 479, 10, 48, 12, 48, 14, 48, 482,
	11, 48, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 7, 50, 489, 10, 50, 12, 50,
	14, 50, 492, 11, 50, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 5, 52, 499, 10,
	52, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 7, 54, 506, 10, 54, 12, 54, 14,
	54, 509, 11, 54, 3, 55, 3, 55, 3, 55, 3, 56, 3, 56, 7, 56, 516, 10, 56,
	12, 56, 14, 56, 519, 11, 56, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58, 3, 58,
	3, 58, 3, 58, 5, 58, 529, 10, 58, 3, 58, 5, 58, 532, 10, 58, 3, 59, 3,
	59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 5, 60, 543, 10, 60,
	3, 61, 3, 61, 3, 62, 3, 62, 3, 63, 3, 63, 3, 64, 3, 64, 3, 65, 3, 65, 7,
	65, 555, 10, 65, 12, 65, 14, 65, 558, 11, 65, 3, 65, 3, 65, 3, 66, 3, 66,
	3, 66, 3, 66, 3, 66, 3, 67, 3, 67, 7, 67, 569, 10, 67, 12, 67, 14, 67,
	572, 11, 67, 3, 67, 7, 67, 575, 10, 67, 12, 67, 14, 67, 578, 11, 67, 3,
	67, 3, 67, 3, 68, 3, 68, 3, 69, 3, 69, 3, 69, 3, 556, 2, 70, 2, 4, 6, 8,
	10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
	46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80,
	82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112,
	114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 136, 2, 6, 3, 2,
	26, 27, 3, 2, 24, 25, 3, 2, 47, 50, 3, 2, 20, 23, 2, 585, 2, 138, 3, 2,
	2, 2, 4, 152, 3, 2, 2, 2, 6, 164, 3, 2, 2, 2, 8, 174, 3, 2, 2, 2, 10, 186,
	3, 2, 2, 2, 12, 193, 3, 2, 2, 2, 14, 214, 3, 2, 2, 2, 16, 216, 3, 2, 2,
	2, 18, 222, 3, 2, 2, 2, 20, 231, 3, 2, 2, 2, 22, 245, 3, 2, 2, 2, 24, 247,
	3, 2, 2, 2, 26, 258, 3, 2, 2, 2, 28, 263, 3, 2, 2, 2, 30, 266, 3, 2, 2,
	2, 32, 271, 3, 2, 2, 2, 34, 280, 3, 2, 2, 2, 36, 297, 3, 2, 2, 2, 38, 310,
	3, 2, 2, 2, 40, 318, 3, 2, 2, 2, 42, 322, 3, 2, 2, 2, 44, 337, 3, 2, 2,
	2, 46, 343, 3, 2, 2, 2, 48, 359, 3, 2, 2, 2, 50, 369, 3, 2, 2, 2, 52, 374,
	3, 2, 2, 2, 54, 380, 3, 2, 2, 2, 56, 387, 3, 2, 2, 2, 58, 393, 3, 2, 2,
	2, 60, 395, 3, 2, 2, 2, 62, 404, 3, 2, 2, 2, 64, 412, 3, 2, 2, 2, 66, 414,
	3, 2, 2, 2, 68, 423, 3, 2, 2, 2, 70, 430, 3, 2, 2, 2, 72, 432, 3, 2, 2,
	2, 74, 438, 3, 2, 2, 2, 76, 444, 3, 2, 2, 2, 78, 448, 3, 2, 2, 2, 80, 450,
	3, 2, 2, 2, 82, 453, 3, 2, 2, 2, 84, 459, 3, 2, 2, 2, 86, 463, 3, 2, 2,
	2, 88, 465, 3, 2, 2, 2, 90, 469, 3, 2, 2, 2, 92, 473, 3, 2, 2, 2, 94, 476,
	3, 2, 2, 2, 96, 483, 3, 2, 2, 2, 98, 486, 3, 2, 2, 2, 100, 493, 3, 2, 2,
	2, 102, 496, 3, 2, 2, 2, 104, 500, 3, 2, 2, 2, 106, 503, 3, 2, 2, 2, 108,
	510, 3, 2, 2, 2, 110, 513, 3, 2, 2, 2, 112, 520, 3, 2, 2, 2, 114, 531,
	3, 2, 2, 2, 116, 533, 3, 2, 2, 2, 118, 542, 3, 2, 2, 2, 120, 544, 3, 2,
	2, 2, 122, 546, 3, 2, 2, 2, 124, 548, 3, 2, 2, 2, 126, 550, 3, 2, 2, 2,
	128, 552, 3, 2, 2, 2, 130, 561, 3, 2, 2, 2, 132, 566, 3, 2, 2, 2, 134,
	581, 3, 2, 2, 2, 136, 583, 3, 2, 2, 2, 138, 139, 7, 52, 2, 2, 139, 140,
	7, 53, 2, 2, 140, 144, 7, 28, 2, 2, 141, 143, 5, 14, 8, 2, 142, 141, 3,
	2, 2, 2, 143, 146, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 144, 145, 3, 2, 2,
	2, 145, 147, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 147, 148, 5, 4, 3, 2, 148,
	3, 3, 2, 2, 2, 149, 151, 5, 6, 4, 2, 150, 149, 3, 2, 2, 2, 151, 154, 3,
	2, 2, 2, 152, 150, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 158, 3, 2, 2,
	2, 154, 152, 3, 2, 2, 2, 155, 157, 5, 36, 19, 2, 156, 155, 3, 2, 2, 2,
	157, 160, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159,
	161, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 161, 162, 5, 130, 66, 2, 162, 163,
	7, 2, 2, 3, 163, 5, 3, 2, 2, 2, 164, 165, 7, 37, 2, 2, 165, 169, 7, 53,
	2, 2, 166, 167, 7, 21, 2, 2, 167, 168, 7, 53, 2, 2, 168, 170, 7, 20, 2,
	2, 169, 166, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171,
	172, 5, 8, 5, 2, 172, 173, 7, 28, 2, 2, 173, 7, 3, 2, 2, 2, 174, 175, 7,
	3, 2, 2, 175, 176, 5, 10, 6, 2, 176, 177, 5, 12, 7, 2, 177, 181, 7, 39,
	2, 2, 178, 180, 5, 36, 19, 2, 179, 178, 3, 2, 2, 2, 180, 183, 3, 2, 2,
	2, 181, 179, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 184, 3, 2, 2, 2, 183,
	181, 3, 2, 2, 2, 184, 185, 7, 4, 2, 2, 185, 9, 3, 2, 2, 2, 186, 190, 7,
	38, 2, 2, 187, 189, 5, 22, 12, 2, 188, 187, 3, 2, 2, 2, 189, 192, 3, 2,
	2, 2, 190, 188, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 11, 3, 2, 2, 2,
	192, 190, 3, 2, 2, 2, 193, 194, 7, 40, 2, 2, 194, 195, 7, 30, 2, 2, 195,
	196, 7, 31, 2, 2, 196, 200, 7, 3, 2, 2, 197, 199, 5, 14, 8, 2, 198, 197,
	3, 2, 2, 2, 199, 202, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 200, 201, 3, 2,
	2, 2, 201, 206, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 203, 205, 5, 48, 25,
	2, 204, 203, 3, 2, 2, 2, 205, 208, 3, 2, 2, 2, 206, 204, 3, 2, 2, 2, 206,
	207, 3, 2, 2, 2, 207, 209, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2, 209, 210,
	7, 4, 2, 2, 210, 13, 3, 2, 2, 2, 211, 215, 5, 16, 9, 2, 212, 215, 5, 18,
	10, 2, 213, 215, 5, 20, 11, 2, 214, 211, 3, 2, 2, 2, 214, 212, 3, 2, 2,
	2, 214, 213, 3, 2, 2, 2, 215, 15, 3, 2, 2, 2, 216, 217, 7, 46, 2, 2, 217,
	218, 7, 53, 2, 2, 218, 219, 7, 5, 2, 2, 219, 220, 5, 26, 14, 2, 220, 221,
	7, 28, 2, 2, 221, 17, 3, 2, 2, 2, 222, 223, 7, 46, 2, 2, 223, 224, 7, 53,
	2, 2, 224, 225, 7, 5, 2, 2, 225, 226, 7, 6, 2, 2, 226, 227, 7, 13, 2, 2,
	227, 228, 7, 7, 2, 2, 228, 229, 5, 134, 68, 2, 229, 230, 7, 28, 2, 2, 230,
	19, 3, 2, 2, 2, 231, 232, 7, 46, 2, 2, 232, 233, 7, 53, 2, 2, 233, 234,
	7, 5, 2, 2, 234, 235, 7, 6, 2, 2, 235, 236, 7, 13, 2, 2, 236, 237, 7, 8,
	2, 2, 237, 238, 7, 13, 2, 2, 238, 239, 7, 7, 2, 2, 239, 240, 5, 134, 68,
	2, 240, 241, 7, 28, 2, 2, 241, 21, 3, 2, 2, 2, 242, 246, 5, 24, 13, 2,
	243, 246, 5, 18, 10, 2, 244, 246, 5, 20, 11, 2, 245, 242, 3, 2, 2, 2, 245,
	243, 3, 2, 2, 2, 245, 244, 3, 2, 2, 2, 246, 23, 3, 2, 2, 2, 247, 248, 7,
	46, 2, 2, 248, 249, 7, 53, 2, 2, 249, 252, 7, 5, 2, 2, 250, 253, 5, 134,
	68, 2, 251, 253, 7, 53, 2, 2, 252, 250, 3, 2, 2, 2, 252, 251, 3, 2, 2,
	2, 253, 254, 3, 2, 2, 2, 254, 255, 7, 28, 2, 2, 255, 25, 3, 2, 2, 2, 256,
	259, 5, 134, 68, 2, 257, 259, 7, 53, 2, 2, 258, 256, 3, 2, 2, 2, 258, 257,
	3, 2, 2, 2, 259, 261, 3, 2, 2, 2, 260, 262, 5, 28, 15, 2, 261, 260, 3,
	2, 2, 2, 261, 262, 3, 2, 2, 2, 262, 27, 3, 2, 2, 2, 263, 264, 7, 29, 2,
	2, 264, 265, 5, 94, 48, 2, 265, 29, 3, 2, 2, 2, 266, 269, 7, 53, 2, 2,
	267, 268, 7, 9, 2, 2, 268, 270, 7, 53, 2, 2, 269, 267, 3, 2, 2, 2, 269,
	270, 3, 2, 2, 2, 270, 31, 3, 2, 2, 2, 271, 274, 7, 53, 2, 2, 272, 273,
	7, 9, 2, 2, 273, 275, 7, 53, 2, 2, 274, 272, 3, 2, 2, 2, 274, 275, 3, 2,
	2, 2, 275, 276, 3, 2, 2, 2, 276, 277, 7, 6, 2, 2, 277, 278, 5, 94, 48,
	2, 278, 279, 7, 7, 2, 2, 279, 33, 3, 2, 2, 2, 280, 283, 7, 53, 2, 2, 281,
	282, 7, 9, 2, 2, 282, 284, 7, 53, 2, 2, 283, 281, 3, 2, 2, 2, 283, 284,
	3, 2, 2, 2, 284, 289, 3, 2, 2, 2, 285, 286, 7, 6, 2, 2, 286, 287, 5, 94,
	48, 2, 287, 288, 7, 7, 2, 2, 288, 290, 3, 2, 2, 2, 289, 285, 3, 2, 2, 2,
	289, 290, 3, 2, 2, 2, 290, 295, 3, 2, 2, 2, 291, 292, 7, 6, 2, 2, 292,
	293, 5, 94, 48, 2, 293, 294, 7, 7, 2, 2, 294, 296, 3, 2, 2, 2, 295, 291,
	3, 2, 2, 2, 295, 296, 3, 2, 2, 2, 296, 35, 3, 2, 2, 2, 297, 298, 7, 43,
	2, 2, 298, 299, 7, 53, 2, 2, 299, 301, 7, 30, 2, 2, 300, 302, 5, 38, 20,
	2, 301, 300, 3, 2, 2, 2, 301, 302, 3, 2, 2, 2, 302, 303, 3, 2, 2, 2, 303,
	306, 7, 31, 2, 2, 304, 307, 5, 134, 68, 2, 305, 307, 7, 51, 2, 2, 306,
	304, 3, 2, 2, 2, 306, 305, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 309,
	5, 42, 22, 2, 309, 37, 3, 2, 2, 2, 310, 315, 5, 40, 21, 2, 311, 312, 7,
	10, 2, 2, 312, 314, 5, 40, 21, 2, 313, 311, 3, 2, 2, 2, 314, 317, 3, 2,
	2, 2, 315, 313, 3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316, 39, 3, 2, 2, 2,
	317, 315, 3, 2, 2, 2, 318, 319, 7, 53, 2, 2, 319, 320, 7, 5, 2, 2, 320,
	321, 5, 134, 68, 2, 321, 41, 3, 2, 2, 2, 322, 326, 7, 3, 2, 2, 323, 325,
	5, 14, 8, 2, 324, 323, 3, 2, 2, 2, 325, 328, 3, 2, 2, 2, 326, 324, 3, 2,
	2, 2, 326, 327, 3, 2, 2, 2, 327, 332, 3, 2, 2, 2, 328, 326, 3, 2, 2, 2,
	329, 331, 5, 50, 26, 2, 330, 329, 3, 2, 2, 2, 331, 334, 3, 2, 2, 2, 332,
	330, 3, 2, 2, 2, 332, 333, 3, 2, 2, 2, 333, 335, 3, 2, 2, 2, 334, 332,
	3, 2, 2, 2, 335, 336, 7, 4, 2, 2, 336, 43, 3, 2, 2, 2, 337, 339, 7, 44,
	2, 2, 338, 340, 5, 94, 48, 2, 339, 338, 3, 2, 2, 2, 339, 340, 3, 2, 2,
	2, 340, 341, 3, 2, 2, 2, 341, 342, 7, 28, 2, 2, 342, 45, 3, 2, 2, 2, 343,
	347, 7, 3, 2, 2, 344, 346, 5, 50, 26, 2, 345, 344, 3, 2, 2, 2, 346, 349,
	3, 2, 2, 2, 347, 345, 3, 2, 2, 2, 347, 348, 3, 2, 2, 2, 348, 350, 3, 2,
	2, 2, 349, 347, 3, 2, 2, 2, 350, 351, 7, 4, 2, 2, 351, 47, 3, 2, 2, 2,
	352, 360, 5, 52, 27, 2, 353, 360, 5, 66, 34, 2, 354, 360, 5, 72, 37, 2,
	355, 360, 5, 74, 38, 2, 356, 360, 5, 82, 42, 2, 357, 360, 5, 88, 45, 2,
	358, 360, 5, 92, 47, 2, 359, 352, 3, 2, 2, 2, 359, 353, 3, 2, 2, 2, 359,
	354, 3, 2, 2, 2, 359, 355, 3, 2, 2, 2, 359, 356, 3, 2, 2, 2, 359, 357,
	3, 2, 2, 2, 359, 358, 3, 2, 2, 2, 360, 49, 3, 2, 2, 2, 361, 370, 5, 52,
	27, 2, 362, 370, 5, 66, 34, 2, 363, 370, 5, 72, 37, 2, 364, 370, 5, 74,
	38, 2, 365, 370, 5, 82, 42, 2, 366, 370, 5, 88, 45, 2, 367, 370, 5, 92,
	47, 2, 368, 370, 5, 44, 23, 2, 369, 361, 3, 2, 2, 2, 369, 362, 3, 2, 2,
	2, 369, 363, 3, 2, 2, 2, 369, 364, 3, 2, 2, 2, 369, 365, 3, 2, 2, 2, 369,
	366, 3, 2, 2, 2, 369, 367, 3, 2, 2, 2, 369, 368, 3, 2, 2, 2, 370, 51, 3,
	2, 2, 2, 371, 375, 5, 30, 16, 2, 372, 375, 5, 32, 17, 2, 373, 375, 5, 34,
	18, 2, 374, 371, 3, 2, 2, 2, 374, 372, 3, 2, 2, 2, 374, 373, 3, 2, 2, 2,
	375, 376, 3, 2, 2, 2, 376, 377, 7, 29, 2, 2, 377, 378, 5, 94, 48, 2, 378,
	379, 7, 28, 2, 2, 379, 53, 3, 2, 2, 2, 380, 381, 7, 53, 2, 2, 381, 383,
	7, 30, 2, 2, 382, 384, 5, 56, 29, 2, 383, 382, 3, 2, 2, 2, 383, 384, 3,
	2, 2, 2, 384, 385, 3, 2, 2, 2, 385, 386, 7, 31, 2, 2, 386, 55, 3, 2, 2,
	2, 387, 388, 5, 94, 48, 2, 388, 389, 5, 58, 30, 2, 389, 57, 3, 2, 2, 2,
	390, 391, 7, 10, 2, 2, 391, 394, 5, 56, 29, 2, 392, 394, 3, 2, 2, 2, 393,
	390, 3, 2, 2, 2, 393, 392, 3, 2, 2, 2, 394, 59, 3, 2, 2, 2, 395, 396, 7,
	53, 2, 2, 396, 397, 7, 9, 2, 2, 397, 398, 7, 53, 2, 2, 398, 400, 7, 30,
	2, 2, 399, 401, 5, 56, 29, 2, 400, 399, 3, 2, 2, 2, 400, 401, 3, 2, 2,
	2, 401, 402, 3, 2, 2, 2, 402, 403, 7, 31, 2, 2, 403, 61, 3, 2, 2, 2, 404,
	405, 7, 11, 2, 2, 405, 406, 7, 53, 2, 2, 406, 407, 7, 30, 2, 2, 407, 408,
	7, 31, 2, 2, 408, 63, 3, 2, 2, 2, 409, 413, 5, 54, 28, 2, 410, 413, 5,
	60, 31, 2, 411, 413, 5, 62, 32, 2, 412, 409, 3, 2, 2, 2, 412, 410, 3, 2,
	2, 2, 412, 411, 3, 2, 2, 2, 413, 65, 3, 2, 2, 2, 414, 415, 7, 42, 2, 2,
	415, 416, 7, 30, 2, 2, 416, 417, 5, 68, 35, 2, 417, 418, 7, 31, 2, 2, 418,
	419, 7, 28, 2, 2, 419, 67, 3, 2, 2, 2, 420, 424, 5, 30, 16, 2, 421, 424,
	5, 32, 17, 2, 422, 424, 5, 34, 18, 2, 423, 420, 3, 2, 2, 2, 423, 421, 3,
	2, 2, 2, 423, 422, 3, 2, 2, 2, 424, 425, 3, 2, 2, 2, 425, 426, 5, 70, 36,
	2, 426, 69, 3, 2, 2, 2, 427, 428, 7, 10, 2, 2, 428, 431, 5, 68, 35, 2,
	429, 431, 3, 2, 2, 2, 430, 427, 3, 2, 2, 2, 430, 429, 3, 2, 2, 2, 431,
	71, 3, 2, 2, 2, 432, 433, 7, 41, 2, 2, 433, 434, 7, 30, 2, 2, 434, 435,
	5, 56, 29, 2, 435, 436, 7, 31, 2, 2, 436, 437, 7, 28, 2, 2, 437, 73, 3,
	2, 2, 2, 438, 439, 7, 32, 2, 2, 439, 440, 5, 76, 39, 2, 440, 442, 5, 78,
	40, 2, 441, 443, 5, 80, 41, 2, 442, 441, 3, 2, 2, 2, 442, 443, 3, 2, 2,
	2, 443, 75, 3, 2, 2, 2, 444, 445, 7, 30, 2, 2, 445, 446, 5, 94, 48, 2,
	446, 447, 7, 31, 2, 2, 447, 77, 3, 2, 2, 2, 448, 449, 5, 46, 24, 2, 449,
	79, 3, 2, 2, 2, 450, 451, 7, 33, 2, 2, 451, 452, 5, 46, 24, 2, 452, 81,
	3, 2, 2, 2, 453, 454, 7, 35, 2, 2, 454, 455, 5, 84, 43, 2, 455, 456, 7,
	36, 2, 2, 456, 457, 5, 86, 44, 2, 457, 458, 5, 46, 24, 2, 458, 83, 3, 2,
	2, 2, 459, 460, 7, 53, 2, 2, 460, 461, 7, 29, 2, 2, 461, 462, 5, 94, 48,
	2, 462, 85, 3, 2, 2, 2, 463, 464, 5, 94, 48, 2, 464, 87, 3, 2, 2, 2, 465,
	466, 7, 34, 2, 2, 466, 467, 5, 90, 46, 2, 467, 468, 5, 46, 24, 2, 468,
	89, 3, 2, 2, 2, 469, 470, 7, 30, 2, 2, 470, 471, 5, 94, 48, 2, 471, 472,
	7, 31, 2, 2, 472, 91, 3, 2, 2, 2, 473, 474, 5, 94, 48, 2, 474, 475, 7,
	28, 2, 2, 475, 93, 3, 2, 2, 2, 476, 480, 5, 98, 50, 2, 477, 479, 5, 96,
	49, 2, 478, 477, 3, 2, 2, 2, 479, 482, 3, 2, 2, 2, 480, 478, 3, 2, 2, 2,
	480, 481, 3, 2, 2, 2, 481, 95, 3, 2, 2, 2, 482, 480, 3, 2, 2, 2, 483, 484,
	7, 18, 2, 2, 484, 485, 5, 98, 50, 2, 485, 97, 3, 2, 2, 2, 486, 490, 5,
	102, 52, 2, 487, 489, 5, 100, 51, 2, 488, 487, 3, 2, 2, 2, 489, 492, 3,
	2, 2, 2, 490, 488, 3, 2, 2, 2, 490, 491, 3, 2, 2, 2, 491, 99, 3, 2, 2,
	2, 492, 490, 3, 2, 2, 2, 493, 494, 7, 19, 2, 2, 494, 495, 5, 102, 52, 2,
	495, 101, 3, 2, 2, 2, 496, 498, 5, 106, 54, 2, 497, 499, 5, 104, 53, 2,
	498, 497, 3, 2, 2, 2, 498, 499, 3, 2, 2, 2, 499, 103, 3, 2, 2, 2, 500,
	501, 5, 136, 69, 2, 501, 502, 5, 106, 54, 2, 502, 105, 3, 2, 2, 2, 503,
	507, 5, 110, 56, 2, 504, 506, 5, 108, 55, 2, 505, 504, 3, 2, 2, 2, 506,
	509, 3, 2, 2, 2, 507, 505, 3, 2, 2, 2, 507, 508, 3, 2, 2, 2, 508, 107,
	3, 2, 2, 2, 509, 507, 3, 2, 2, 2, 510, 511, 9, 2, 2, 2, 511, 512, 5, 110,
	56, 2, 512, 109, 3, 2, 2, 2, 513, 517, 5, 114, 58, 2, 514, 516, 5, 112,
	57, 2, 515, 514, 3, 2, 2, 2, 516, 519, 3, 2, 2, 2, 517, 515, 3, 2, 2, 2,
	517, 518, 3, 2, 2, 2, 518, 111, 3, 2, 2, 2, 519, 517, 3, 2, 2, 2, 520,
	521, 9, 3, 2, 2, 521, 522, 5, 114, 58, 2, 522, 113, 3, 2, 2, 2, 523, 532,
	5, 116, 59, 2, 524, 532, 5, 118, 60, 2, 525, 529, 5, 30, 16, 2, 526, 529,
	5, 32, 17, 2, 527, 529, 5, 34, 18, 2, 528, 525, 3, 2, 2, 2, 528, 526, 3,
	2, 2, 2, 528, 527, 3, 2, 2, 2, 529, 532, 3, 2, 2, 2, 530, 532, 5, 64, 33,
	2, 531, 523, 3, 2, 2, 2, 531, 524, 3, 2, 2, 2, 531, 528, 3, 2, 2, 2, 531,
	530, 3, 2, 2, 2, 532, 115, 3, 2, 2, 2, 533, 534, 7, 30, 2, 2, 534, 535,
	5, 94, 48, 2, 535, 536, 7, 31, 2, 2, 536, 117, 3, 2, 2, 2, 537, 543, 5,
	120, 61, 2, 538, 543, 5, 122, 62, 2, 539, 543, 5, 124, 63, 2, 540, 543,
	5, 128, 65, 2, 541, 543, 5, 126, 64, 2, 542, 537, 3, 2, 2, 2, 542, 538,
	3, 2, 2, 2, 542, 539, 3, 2, 2, 2, 542, 540, 3, 2, 2, 2, 542, 541, 3, 2,
	2, 2, 543, 119, 3, 2, 2, 2, 544, 545, 7, 13, 2, 2, 545, 121, 3, 2, 2, 2,
	546, 547, 7, 14, 2, 2, 547, 123, 3, 2, 2, 2, 548, 549, 7, 15, 2, 2, 549,
	125, 3, 2, 2, 2, 550, 551, 7, 16, 2, 2, 551, 127, 3, 2, 2, 2, 552, 556,
	7, 12, 2, 2, 553, 555, 11, 2, 2, 2, 554, 553, 3, 2, 2, 2, 555, 558, 3,
	2, 2, 2, 556, 557, 3, 2, 2, 2, 556, 554, 3, 2, 2, 2, 557, 559, 3, 2, 2,
	2, 558, 556, 3, 2, 2, 2, 559, 560, 7, 12, 2, 2, 560, 129, 3, 2, 2, 2, 561,
	562, 7, 45, 2, 2, 562, 563, 7, 30, 2, 2, 563, 564, 7, 31, 2, 2, 564, 565,
	5, 132, 67, 2, 565, 131, 3, 2, 2, 2, 566, 570, 7, 3, 2, 2, 567, 569, 5,
	14, 8, 2, 568, 567, 3, 2, 2, 2, 569, 572, 3, 2, 2, 2, 570, 568, 3, 2, 2,
	2, 570, 571, 3, 2, 2, 2, 571, 576, 3, 2, 2, 2, 572, 570, 3, 2, 2, 2, 573,
	575, 5, 50, 26, 2, 574, 573, 3, 2, 2, 2, 575, 578, 3, 2, 2, 2, 576, 574,
	3, 2, 2, 2, 576, 577, 3, 2, 2, 2, 577, 579, 3, 2, 2, 2, 578, 576, 3, 2,
	2, 2, 579, 580, 7, 4, 2, 2, 580, 133, 3, 2, 2, 2, 581, 582, 9, 4, 2, 2,
	582, 135, 3, 2, 2, 2, 583, 584, 9, 5, 2, 2, 584, 137, 3, 2, 2, 2, 48, 144,
	152, 158, 169, 181, 190, 200, 206, 214, 245, 252, 258, 261, 269, 274, 283,
	289, 295, 301, 306, 315, 326, 332, 339, 347, 359, 369, 374, 383, 393, 400,
	412, 423, 430, 442, 480, 490, 498, 507, 517, 528, 531, 542, 556, 570, 576,
}
var literalNames = []string{
	"", "'{'", "'}'", "':'", "'['", "']'", "']['", "'.'", "','", "'new'", "'\"'",
	"", "", "", "", "", "'||'", "'&&'", "'>'", "'<'", "'=='", "'!='", "'*'",
	"'/'", "'+'", "'-'", "';'", "'='", "'('", "')'", "'if'", "'else'", "'while'",
	"'for'", "'in'", "'class'", "'attributes'", "'methods'", "'init'", "'write'",
	"'read'", "'function'", "'return'", "'main'", "'var'", "'int'", "'float'",
	"'char'", "'bool'", "'void'", "'program'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "INT", "FLOAT", "CHAR", "BOOL",
	"WS", "OR", "AND", "GT", "LT", "EQ", "NEQ", "MUL", "DIV", "ADD", "SUB",
	"SEMI", "ASSIGN", "LPAREN", "RPAREN", "IF", "ELSE", "WHILE", "FOR", "IN",
	"CLASS", "ATTRIBUTES", "METHODS", "INIT", "WRITE", "READ", "FUNCTION",
	"RETURN", "MAIN", "VAR", "INT_TYPE", "FLOAT_TYPE", "CHAR_TYPE", "BOOL_TYPE",
	"VOID", "PROGRAM", "ID",
}

var ruleNames = []string{
	"program", "program2", "classDef", "classBlock", "classAttributes", "classInit",
	"variableDeclaration", "varsDec", "varsDecArray", "varsDecMat", "attributesDeclaration",
	"attributesDec", "varsTypeInit", "varsTypeInit2", "vars", "varArray", "varMat",
	"functions", "parameters", "parameter", "functionBlock", "returnRule",
	"block", "statutesNoReturn", "statutes", "assignation", "functionCall",
	"arguments", "arguments2", "methodCall", "initCall", "call", "read", "read2",
	"read3", "write", "conditional", "conditional2", "conditional3", "conditional4",
	"forLoop", "forLoop2", "forLoop3", "whileLoop", "whileLoop2", "expression",
	"exp", "exp2", "t_exp", "t_exp2", "g_exp", "g_exp2", "m_exp", "m_exp2",
	"term", "term2", "factor", "factor2", "varCte", "cte_i", "cte_f", "cte_c",
	"cte_b", "cte_s", "main", "mainBlock", "typeRule", "relop",
}

type ProyectoParser struct {
	*antlr.BaseParser
}

// NewProyectoParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ProyectoParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewProyectoParser(input antlr.TokenStream) *ProyectoParser {
	this := new(ProyectoParser)
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
	this.GrammarFileName = "Proyecto.g4"

	return this
}

// ProyectoParser tokens.
const (
	ProyectoParserEOF        = antlr.TokenEOF
	ProyectoParserT__0       = 1
	ProyectoParserT__1       = 2
	ProyectoParserT__2       = 3
	ProyectoParserT__3       = 4
	ProyectoParserT__4       = 5
	ProyectoParserT__5       = 6
	ProyectoParserT__6       = 7
	ProyectoParserT__7       = 8
	ProyectoParserT__8       = 9
	ProyectoParserT__9       = 10
	ProyectoParserINT        = 11
	ProyectoParserFLOAT      = 12
	ProyectoParserCHAR       = 13
	ProyectoParserBOOL       = 14
	ProyectoParserWS         = 15
	ProyectoParserOR         = 16
	ProyectoParserAND        = 17
	ProyectoParserGT         = 18
	ProyectoParserLT         = 19
	ProyectoParserEQ         = 20
	ProyectoParserNEQ        = 21
	ProyectoParserMUL        = 22
	ProyectoParserDIV        = 23
	ProyectoParserADD        = 24
	ProyectoParserSUB        = 25
	ProyectoParserSEMI       = 26
	ProyectoParserASSIGN     = 27
	ProyectoParserLPAREN     = 28
	ProyectoParserRPAREN     = 29
	ProyectoParserIF         = 30
	ProyectoParserELSE       = 31
	ProyectoParserWHILE      = 32
	ProyectoParserFOR        = 33
	ProyectoParserIN         = 34
	ProyectoParserCLASS      = 35
	ProyectoParserATTRIBUTES = 36
	ProyectoParserMETHODS    = 37
	ProyectoParserINIT       = 38
	ProyectoParserWRITE      = 39
	ProyectoParserREAD       = 40
	ProyectoParserFUNCTION   = 41
	ProyectoParserRETURN     = 42
	ProyectoParserMAIN       = 43
	ProyectoParserVAR        = 44
	ProyectoParserINT_TYPE   = 45
	ProyectoParserFLOAT_TYPE = 46
	ProyectoParserCHAR_TYPE  = 47
	ProyectoParserBOOL_TYPE  = 48
	ProyectoParserVOID       = 49
	ProyectoParserPROGRAM    = 50
	ProyectoParserID         = 51
)

// ProyectoParser rules.
const (
	ProyectoParserRULE_program               = 0
	ProyectoParserRULE_program2              = 1
	ProyectoParserRULE_classDef              = 2
	ProyectoParserRULE_classBlock            = 3
	ProyectoParserRULE_classAttributes       = 4
	ProyectoParserRULE_classInit             = 5
	ProyectoParserRULE_variableDeclaration   = 6
	ProyectoParserRULE_varsDec               = 7
	ProyectoParserRULE_varsDecArray          = 8
	ProyectoParserRULE_varsDecMat            = 9
	ProyectoParserRULE_attributesDeclaration = 10
	ProyectoParserRULE_attributesDec         = 11
	ProyectoParserRULE_varsTypeInit          = 12
	ProyectoParserRULE_varsTypeInit2         = 13
	ProyectoParserRULE_vars                  = 14
	ProyectoParserRULE_varArray              = 15
	ProyectoParserRULE_varMat                = 16
	ProyectoParserRULE_functions             = 17
	ProyectoParserRULE_parameters            = 18
	ProyectoParserRULE_parameter             = 19
	ProyectoParserRULE_functionBlock         = 20
	ProyectoParserRULE_returnRule            = 21
	ProyectoParserRULE_block                 = 22
	ProyectoParserRULE_statutesNoReturn      = 23
	ProyectoParserRULE_statutes              = 24
	ProyectoParserRULE_assignation           = 25
	ProyectoParserRULE_functionCall          = 26
	ProyectoParserRULE_arguments             = 27
	ProyectoParserRULE_arguments2            = 28
	ProyectoParserRULE_methodCall            = 29
	ProyectoParserRULE_initCall              = 30
	ProyectoParserRULE_call                  = 31
	ProyectoParserRULE_read                  = 32
	ProyectoParserRULE_read2                 = 33
	ProyectoParserRULE_read3                 = 34
	ProyectoParserRULE_write                 = 35
	ProyectoParserRULE_conditional           = 36
	ProyectoParserRULE_conditional2          = 37
	ProyectoParserRULE_conditional3          = 38
	ProyectoParserRULE_conditional4          = 39
	ProyectoParserRULE_forLoop               = 40
	ProyectoParserRULE_forLoop2              = 41
	ProyectoParserRULE_forLoop3              = 42
	ProyectoParserRULE_whileLoop             = 43
	ProyectoParserRULE_whileLoop2            = 44
	ProyectoParserRULE_expression            = 45
	ProyectoParserRULE_exp                   = 46
	ProyectoParserRULE_exp2                  = 47
	ProyectoParserRULE_t_exp                 = 48
	ProyectoParserRULE_t_exp2                = 49
	ProyectoParserRULE_g_exp                 = 50
	ProyectoParserRULE_g_exp2                = 51
	ProyectoParserRULE_m_exp                 = 52
	ProyectoParserRULE_m_exp2                = 53
	ProyectoParserRULE_term                  = 54
	ProyectoParserRULE_term2                 = 55
	ProyectoParserRULE_factor                = 56
	ProyectoParserRULE_factor2               = 57
	ProyectoParserRULE_varCte                = 58
	ProyectoParserRULE_cte_i                 = 59
	ProyectoParserRULE_cte_f                 = 60
	ProyectoParserRULE_cte_c                 = 61
	ProyectoParserRULE_cte_b                 = 62
	ProyectoParserRULE_cte_s                 = 63
	ProyectoParserRULE_main                  = 64
	ProyectoParserRULE_mainBlock             = 65
	ProyectoParserRULE_typeRule              = 66
	ProyectoParserRULE_relop                 = 67
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
	p.RuleIndex = ProyectoParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) PROGRAM() antlr.TerminalNode {
	return s.GetToken(ProyectoParserPROGRAM, 0)
}

func (s *ProgramContext) ID() antlr.TerminalNode {
	return s.GetToken(ProyectoParserID, 0)
}

func (s *ProgramContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ProyectoParserSEMI, 0)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *ProyectoParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ProyectoParserRULE_program)
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
		p.SetState(136)
		p.Match(ProyectoParserPROGRAM)
	}
	{
		p.SetState(137)
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(138)
		p.Match(ProyectoParserSEMI)
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserVAR {
		{
			p.SetState(139)
			p.VariableDeclaration()
		}

		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(145)
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
	p.RuleIndex = ProyectoParserRULE_program2
	return p
}

func (*Program2Context) IsProgram2Context() {}

func NewProgram2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Program2Context {
	var p = new(Program2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_program2

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
	return s.GetToken(ProyectoParserEOF, 0)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterProgram2(s)
	}
}

func (s *Program2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitProgram2(s)
	}
}

func (p *ProyectoParser) Program2() (localctx IProgram2Context) {
	localctx = NewProgram2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ProyectoParserRULE_program2)
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
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserCLASS {
		{
			p.SetState(147)
			p.ClassDef()
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserFUNCTION {
		{
			p.SetState(153)
			p.Functions()
		}

		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(159)
		p.Main()
	}
	{
		p.SetState(160)
		p.Match(ProyectoParserEOF)
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
	p.RuleIndex = ProyectoParserRULE_classDef
	return p
}

func (*ClassDefContext) IsClassDefContext() {}

func NewClassDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDefContext {
	var p = new(ClassDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_classDef

	return p
}

func (s *ClassDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDefContext) CLASS() antlr.TerminalNode {
	return s.GetToken(ProyectoParserCLASS, 0)
}

func (s *ClassDefContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ProyectoParserID)
}

func (s *ClassDefContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ProyectoParserID, i)
}

func (s *ClassDefContext) ClassBlock() IClassBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassBlockContext)
}

func (s *ClassDefContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ProyectoParserSEMI, 0)
}

func (s *ClassDefContext) LT() antlr.TerminalNode {
	return s.GetToken(ProyectoParserLT, 0)
}

func (s *ClassDefContext) GT() antlr.TerminalNode {
	return s.GetToken(ProyectoParserGT, 0)
}

func (s *ClassDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterClassDef(s)
	}
}

func (s *ClassDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitClassDef(s)
	}
}

func (p *ProyectoParser) ClassDef() (localctx IClassDefContext) {
	localctx = NewClassDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ProyectoParserRULE_classDef)
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
		p.SetState(162)
		p.Match(ProyectoParserCLASS)
	}
	{
		p.SetState(163)
		p.Match(ProyectoParserID)
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProyectoParserLT {
		{
			p.SetState(164)
			p.Match(ProyectoParserLT)
		}
		{
			p.SetState(165)
			p.Match(ProyectoParserID)
		}
		{
			p.SetState(166)
			p.Match(ProyectoParserGT)
		}

	}
	{
		p.SetState(169)
		p.ClassBlock()
	}
	{
		p.SetState(170)
		p.Match(ProyectoParserSEMI)
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
	p.RuleIndex = ProyectoParserRULE_classBlock
	return p
}

func (*ClassBlockContext) IsClassBlockContext() {}

func NewClassBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassBlockContext {
	var p = new(ClassBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_classBlock

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
	return s.GetToken(ProyectoParserMETHODS, 0)
}

func (s *ClassBlockContext) AllFunctions() []IFunctionsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionsContext)(nil)).Elem())
	var tst = make([]IFunctionsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionsContext)
		}
	}

	return tst
}

func (s *ClassBlockContext) Functions(i int) IFunctionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionsContext)
}

func (s *ClassBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterClassBlock(s)
	}
}

func (s *ClassBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitClassBlock(s)
	}
}

func (p *ProyectoParser) ClassBlock() (localctx IClassBlockContext) {
	localctx = NewClassBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ProyectoParserRULE_classBlock)
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
		p.SetState(172)
		p.Match(ProyectoParserT__0)
	}
	{
		p.SetState(173)
		p.ClassAttributes()
	}
	{
		p.SetState(174)
		p.ClassInit()
	}
	{
		p.SetState(175)
		p.Match(ProyectoParserMETHODS)
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserFUNCTION {
		{
			p.SetState(176)
			p.Functions()
		}

		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(182)
		p.Match(ProyectoParserT__1)
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
	p.RuleIndex = ProyectoParserRULE_classAttributes
	return p
}

func (*ClassAttributesContext) IsClassAttributesContext() {}

func NewClassAttributesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassAttributesContext {
	var p = new(ClassAttributesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_classAttributes

	return p
}

func (s *ClassAttributesContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassAttributesContext) ATTRIBUTES() antlr.TerminalNode {
	return s.GetToken(ProyectoParserATTRIBUTES, 0)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterClassAttributes(s)
	}
}

func (s *ClassAttributesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitClassAttributes(s)
	}
}

func (p *ProyectoParser) ClassAttributes() (localctx IClassAttributesContext) {
	localctx = NewClassAttributesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ProyectoParserRULE_classAttributes)
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
		p.SetState(184)
		p.Match(ProyectoParserATTRIBUTES)
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserVAR {
		{
			p.SetState(185)
			p.AttributesDeclaration()
		}

		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = ProyectoParserRULE_classInit
	return p
}

func (*ClassInitContext) IsClassInitContext() {}

func NewClassInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassInitContext {
	var p = new(ClassInitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_classInit

	return p
}

func (s *ClassInitContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassInitContext) INIT() antlr.TerminalNode {
	return s.GetToken(ProyectoParserINIT, 0)
}

func (s *ClassInitContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserLPAREN, 0)
}

func (s *ClassInitContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserRPAREN, 0)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterClassInit(s)
	}
}

func (s *ClassInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitClassInit(s)
	}
}

func (p *ProyectoParser) ClassInit() (localctx IClassInitContext) {
	localctx = NewClassInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ProyectoParserRULE_classInit)
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
		p.SetState(191)
		p.Match(ProyectoParserINIT)
	}
	{
		p.SetState(192)
		p.Match(ProyectoParserLPAREN)
	}
	{
		p.SetState(193)
		p.Match(ProyectoParserRPAREN)
	}
	{
		p.SetState(194)
		p.Match(ProyectoParserT__0)
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserVAR {
		{
			p.SetState(195)
			p.VariableDeclaration()
		}

		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProyectoParserT__8)|(1<<ProyectoParserT__9)|(1<<ProyectoParserINT)|(1<<ProyectoParserFLOAT)|(1<<ProyectoParserCHAR)|(1<<ProyectoParserBOOL)|(1<<ProyectoParserLPAREN)|(1<<ProyectoParserIF))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ProyectoParserWHILE-32))|(1<<(ProyectoParserFOR-32))|(1<<(ProyectoParserWRITE-32))|(1<<(ProyectoParserREAD-32))|(1<<(ProyectoParserID-32)))) != 0) {
		{
			p.SetState(201)
			p.StatutesNoReturn()
		}

		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(207)
		p.Match(ProyectoParserT__1)
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
	p.RuleIndex = ProyectoParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_variableDeclaration

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (p *ProyectoParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ProyectoParserRULE_variableDeclaration)

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

	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(209)
			p.VarsDec()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(210)
			p.VarsDecArray()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(211)
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
	p.RuleIndex = ProyectoParserRULE_varsDec
	return p
}

func (*VarsDecContext) IsVarsDecContext() {}

func NewVarsDecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsDecContext {
	var p = new(VarsDecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_varsDec

	return p
}

func (s *VarsDecContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsDecContext) VAR() antlr.TerminalNode {
	return s.GetToken(ProyectoParserVAR, 0)
}

func (s *VarsDecContext) ID() antlr.TerminalNode {
	return s.GetToken(ProyectoParserID, 0)
}

func (s *VarsDecContext) VarsTypeInit() IVarsTypeInitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarsTypeInitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarsTypeInitContext)
}

func (s *VarsDecContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ProyectoParserSEMI, 0)
}

func (s *VarsDecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsDecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsDecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterVarsDec(s)
	}
}

func (s *VarsDecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitVarsDec(s)
	}
}

func (p *ProyectoParser) VarsDec() (localctx IVarsDecContext) {
	localctx = NewVarsDecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ProyectoParserRULE_varsDec)

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
		p.SetState(214)
		p.Match(ProyectoParserVAR)
	}
	{
		p.SetState(215)
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(216)
		p.Match(ProyectoParserT__2)
	}
	{
		p.SetState(217)
		p.VarsTypeInit()
	}
	{
		p.SetState(218)
		p.Match(ProyectoParserSEMI)
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
	p.RuleIndex = ProyectoParserRULE_varsDecArray
	return p
}

func (*VarsDecArrayContext) IsVarsDecArrayContext() {}

func NewVarsDecArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsDecArrayContext {
	var p = new(VarsDecArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_varsDecArray

	return p
}

func (s *VarsDecArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsDecArrayContext) VAR() antlr.TerminalNode {
	return s.GetToken(ProyectoParserVAR, 0)
}

func (s *VarsDecArrayContext) ID() antlr.TerminalNode {
	return s.GetToken(ProyectoParserID, 0)
}

func (s *VarsDecArrayContext) INT() antlr.TerminalNode {
	return s.GetToken(ProyectoParserINT, 0)
}

func (s *VarsDecArrayContext) TypeRule() ITypeRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRuleContext)
}

func (s *VarsDecArrayContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ProyectoParserSEMI, 0)
}

func (s *VarsDecArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsDecArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsDecArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterVarsDecArray(s)
	}
}

func (s *VarsDecArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitVarsDecArray(s)
	}
}

func (p *ProyectoParser) VarsDecArray() (localctx IVarsDecArrayContext) {
	localctx = NewVarsDecArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ProyectoParserRULE_varsDecArray)

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
		p.SetState(220)
		p.Match(ProyectoParserVAR)
	}
	{
		p.SetState(221)
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(222)
		p.Match(ProyectoParserT__2)
	}
	{
		p.SetState(223)
		p.Match(ProyectoParserT__3)
	}
	{
		p.SetState(224)
		p.Match(ProyectoParserINT)
	}
	{
		p.SetState(225)
		p.Match(ProyectoParserT__4)
	}
	{
		p.SetState(226)
		p.TypeRule()
	}
	{
		p.SetState(227)
		p.Match(ProyectoParserSEMI)
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
	p.RuleIndex = ProyectoParserRULE_varsDecMat
	return p
}

func (*VarsDecMatContext) IsVarsDecMatContext() {}

func NewVarsDecMatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsDecMatContext {
	var p = new(VarsDecMatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_varsDecMat

	return p
}

func (s *VarsDecMatContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsDecMatContext) VAR() antlr.TerminalNode {
	return s.GetToken(ProyectoParserVAR, 0)
}

func (s *VarsDecMatContext) ID() antlr.TerminalNode {
	return s.GetToken(ProyectoParserID, 0)
}

func (s *VarsDecMatContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(ProyectoParserINT)
}

func (s *VarsDecMatContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(ProyectoParserINT, i)
}

func (s *VarsDecMatContext) TypeRule() ITypeRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRuleContext)
}

func (s *VarsDecMatContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ProyectoParserSEMI, 0)
}

func (s *VarsDecMatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsDecMatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsDecMatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterVarsDecMat(s)
	}
}

func (s *VarsDecMatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitVarsDecMat(s)
	}
}

func (p *ProyectoParser) VarsDecMat() (localctx IVarsDecMatContext) {
	localctx = NewVarsDecMatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ProyectoParserRULE_varsDecMat)

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
		p.Match(ProyectoParserVAR)
	}
	{
		p.SetState(230)
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(231)
		p.Match(ProyectoParserT__2)
	}
	{
		p.SetState(232)
		p.Match(ProyectoParserT__3)
	}
	{
		p.SetState(233)
		p.Match(ProyectoParserINT)
	}
	{
		p.SetState(234)
		p.Match(ProyectoParserT__5)
	}
	{
		p.SetState(235)
		p.Match(ProyectoParserINT)
	}
	{
		p.SetState(236)
		p.Match(ProyectoParserT__4)
	}
	{
		p.SetState(237)
		p.TypeRule()
	}
	{
		p.SetState(238)
		p.Match(ProyectoParserSEMI)
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
	p.RuleIndex = ProyectoParserRULE_attributesDeclaration
	return p
}

func (*AttributesDeclarationContext) IsAttributesDeclarationContext() {}

func NewAttributesDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributesDeclarationContext {
	var p = new(AttributesDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_attributesDeclaration

	return p
}

func (s *AttributesDeclarationContext) GetParser() antlr.Parser { return s.parser }

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

func (s *AttributesDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributesDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributesDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterAttributesDeclaration(s)
	}
}

func (s *AttributesDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitAttributesDeclaration(s)
	}
}

func (p *ProyectoParser) AttributesDeclaration() (localctx IAttributesDeclarationContext) {
	localctx = NewAttributesDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ProyectoParserRULE_attributesDeclaration)

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

	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(240)
			p.AttributesDec()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(241)
			p.VarsDecArray()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(242)
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
	p.RuleIndex = ProyectoParserRULE_attributesDec
	return p
}

func (*AttributesDecContext) IsAttributesDecContext() {}

func NewAttributesDecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributesDecContext {
	var p = new(AttributesDecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_attributesDec

	return p
}

func (s *AttributesDecContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributesDecContext) VAR() antlr.TerminalNode {
	return s.GetToken(ProyectoParserVAR, 0)
}

func (s *AttributesDecContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ProyectoParserID)
}

func (s *AttributesDecContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ProyectoParserID, i)
}

func (s *AttributesDecContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ProyectoParserSEMI, 0)
}

func (s *AttributesDecContext) TypeRule() ITypeRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRuleContext)
}

func (s *AttributesDecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributesDecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributesDecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterAttributesDec(s)
	}
}

func (s *AttributesDecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitAttributesDec(s)
	}
}

func (p *ProyectoParser) AttributesDec() (localctx IAttributesDecContext) {
	localctx = NewAttributesDecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ProyectoParserRULE_attributesDec)

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
		p.SetState(245)
		p.Match(ProyectoParserVAR)
	}
	{
		p.SetState(246)
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(247)
		p.Match(ProyectoParserT__2)
	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProyectoParserINT_TYPE, ProyectoParserFLOAT_TYPE, ProyectoParserCHAR_TYPE, ProyectoParserBOOL_TYPE:
		{
			p.SetState(248)
			p.TypeRule()
		}

	case ProyectoParserID:
		{
			p.SetState(249)
			p.Match(ProyectoParserID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(252)
		p.Match(ProyectoParserSEMI)
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
	p.RuleIndex = ProyectoParserRULE_varsTypeInit
	return p
}

func (*VarsTypeInitContext) IsVarsTypeInitContext() {}

func NewVarsTypeInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsTypeInitContext {
	var p = new(VarsTypeInitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_varsTypeInit

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
	return s.GetToken(ProyectoParserID, 0)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterVarsTypeInit(s)
	}
}

func (s *VarsTypeInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitVarsTypeInit(s)
	}
}

func (p *ProyectoParser) VarsTypeInit() (localctx IVarsTypeInitContext) {
	localctx = NewVarsTypeInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ProyectoParserRULE_varsTypeInit)
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
	p.SetState(256)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProyectoParserINT_TYPE, ProyectoParserFLOAT_TYPE, ProyectoParserCHAR_TYPE, ProyectoParserBOOL_TYPE:
		{
			p.SetState(254)
			p.TypeRule()
		}

	case ProyectoParserID:
		{
			p.SetState(255)
			p.Match(ProyectoParserID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProyectoParserASSIGN {
		{
			p.SetState(258)
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
	p.RuleIndex = ProyectoParserRULE_varsTypeInit2
	return p
}

func (*VarsTypeInit2Context) IsVarsTypeInit2Context() {}

func NewVarsTypeInit2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsTypeInit2Context {
	var p = new(VarsTypeInit2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_varsTypeInit2

	return p
}

func (s *VarsTypeInit2Context) GetParser() antlr.Parser { return s.parser }

func (s *VarsTypeInit2Context) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserASSIGN, 0)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterVarsTypeInit2(s)
	}
}

func (s *VarsTypeInit2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitVarsTypeInit2(s)
	}
}

func (p *ProyectoParser) VarsTypeInit2() (localctx IVarsTypeInit2Context) {
	localctx = NewVarsTypeInit2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ProyectoParserRULE_varsTypeInit2)

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
		p.SetState(261)
		p.Match(ProyectoParserASSIGN)
	}
	{
		p.SetState(262)
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
	p.RuleIndex = ProyectoParserRULE_vars
	return p
}

func (*VarsContext) IsVarsContext() {}

func NewVarsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsContext {
	var p = new(VarsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_vars

	return p
}

func (s *VarsContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ProyectoParserID)
}

func (s *VarsContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ProyectoParserID, i)
}

func (s *VarsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterVars(s)
	}
}

func (s *VarsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitVars(s)
	}
}

func (p *ProyectoParser) Vars() (localctx IVarsContext) {
	localctx = NewVarsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ProyectoParserRULE_vars)
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
		p.SetState(264)
		p.Match(ProyectoParserID)
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProyectoParserT__6 {
		{
			p.SetState(265)
			p.Match(ProyectoParserT__6)
		}
		{
			p.SetState(266)
			p.Match(ProyectoParserID)
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
	p.RuleIndex = ProyectoParserRULE_varArray
	return p
}

func (*VarArrayContext) IsVarArrayContext() {}

func NewVarArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarArrayContext {
	var p = new(VarArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_varArray

	return p
}

func (s *VarArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *VarArrayContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ProyectoParserID)
}

func (s *VarArrayContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ProyectoParserID, i)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterVarArray(s)
	}
}

func (s *VarArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitVarArray(s)
	}
}

func (p *ProyectoParser) VarArray() (localctx IVarArrayContext) {
	localctx = NewVarArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ProyectoParserRULE_varArray)
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
		p.SetState(269)
		p.Match(ProyectoParserID)
	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProyectoParserT__6 {
		{
			p.SetState(270)
			p.Match(ProyectoParserT__6)
		}
		{
			p.SetState(271)
			p.Match(ProyectoParserID)
		}

	}
	{
		p.SetState(274)
		p.Match(ProyectoParserT__3)
	}
	{
		p.SetState(275)
		p.Exp()
	}
	{
		p.SetState(276)
		p.Match(ProyectoParserT__4)
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
	p.RuleIndex = ProyectoParserRULE_varMat
	return p
}

func (*VarMatContext) IsVarMatContext() {}

func NewVarMatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarMatContext {
	var p = new(VarMatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_varMat

	return p
}

func (s *VarMatContext) GetParser() antlr.Parser { return s.parser }

func (s *VarMatContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ProyectoParserID)
}

func (s *VarMatContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ProyectoParserID, i)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterVarMat(s)
	}
}

func (s *VarMatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitVarMat(s)
	}
}

func (p *ProyectoParser) VarMat() (localctx IVarMatContext) {
	localctx = NewVarMatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ProyectoParserRULE_varMat)
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
		p.SetState(278)
		p.Match(ProyectoParserID)
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProyectoParserT__6 {
		{
			p.SetState(279)
			p.Match(ProyectoParserT__6)
		}
		{
			p.SetState(280)
			p.Match(ProyectoParserID)
		}

	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(283)
			p.Match(ProyectoParserT__3)
		}
		{
			p.SetState(284)
			p.Exp()
		}
		{
			p.SetState(285)
			p.Match(ProyectoParserT__4)
		}

	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProyectoParserT__3 {
		{
			p.SetState(289)
			p.Match(ProyectoParserT__3)
		}
		{
			p.SetState(290)
			p.Exp()
		}
		{
			p.SetState(291)
			p.Match(ProyectoParserT__4)
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
	p.RuleIndex = ProyectoParserRULE_functions
	return p
}

func (*FunctionsContext) IsFunctionsContext() {}

func NewFunctionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionsContext {
	var p = new(FunctionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_functions

	return p
}

func (s *FunctionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionsContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(ProyectoParserFUNCTION, 0)
}

func (s *FunctionsContext) ID() antlr.TerminalNode {
	return s.GetToken(ProyectoParserID, 0)
}

func (s *FunctionsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserLPAREN, 0)
}

func (s *FunctionsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserRPAREN, 0)
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
	return s.GetToken(ProyectoParserVOID, 0)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterFunctions(s)
	}
}

func (s *FunctionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitFunctions(s)
	}
}

func (p *ProyectoParser) Functions() (localctx IFunctionsContext) {
	localctx = NewFunctionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ProyectoParserRULE_functions)
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
		p.SetState(295)
		p.Match(ProyectoParserFUNCTION)
	}
	{
		p.SetState(296)
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(297)
		p.Match(ProyectoParserLPAREN)
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProyectoParserID {
		{
			p.SetState(298)
			p.Parameters()
		}

	}
	{
		p.SetState(301)
		p.Match(ProyectoParserRPAREN)
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProyectoParserINT_TYPE, ProyectoParserFLOAT_TYPE, ProyectoParserCHAR_TYPE, ProyectoParserBOOL_TYPE:
		{
			p.SetState(302)
			p.TypeRule()
		}

	case ProyectoParserVOID:
		{
			p.SetState(303)
			p.Match(ProyectoParserVOID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(306)
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
	p.RuleIndex = ProyectoParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_parameters

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (p *ProyectoParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ProyectoParserRULE_parameters)
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
		p.SetState(308)
		p.Parameter()
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserT__7 {
		{
			p.SetState(309)
			p.Match(ProyectoParserT__7)
		}
		{
			p.SetState(310)
			p.Parameter()
		}

		p.SetState(315)
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
	p.RuleIndex = ProyectoParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) ID() antlr.TerminalNode {
	return s.GetToken(ProyectoParserID, 0)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (p *ProyectoParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ProyectoParserRULE_parameter)

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
		p.SetState(316)
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(317)
		p.Match(ProyectoParserT__2)
	}
	{
		p.SetState(318)
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
	p.RuleIndex = ProyectoParserRULE_functionBlock
	return p
}

func (*FunctionBlockContext) IsFunctionBlockContext() {}

func NewFunctionBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionBlockContext {
	var p = new(FunctionBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_functionBlock

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterFunctionBlock(s)
	}
}

func (s *FunctionBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitFunctionBlock(s)
	}
}

func (p *ProyectoParser) FunctionBlock() (localctx IFunctionBlockContext) {
	localctx = NewFunctionBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ProyectoParserRULE_functionBlock)
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
		p.Match(ProyectoParserT__0)
	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserVAR {
		{
			p.SetState(321)
			p.VariableDeclaration()
		}

		p.SetState(326)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProyectoParserT__8)|(1<<ProyectoParserT__9)|(1<<ProyectoParserINT)|(1<<ProyectoParserFLOAT)|(1<<ProyectoParserCHAR)|(1<<ProyectoParserBOOL)|(1<<ProyectoParserLPAREN)|(1<<ProyectoParserIF))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ProyectoParserWHILE-32))|(1<<(ProyectoParserFOR-32))|(1<<(ProyectoParserWRITE-32))|(1<<(ProyectoParserREAD-32))|(1<<(ProyectoParserRETURN-32))|(1<<(ProyectoParserID-32)))) != 0) {
		{
			p.SetState(327)
			p.Statutes()
		}

		p.SetState(332)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(333)
		p.Match(ProyectoParserT__1)
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
	p.RuleIndex = ProyectoParserRULE_returnRule
	return p
}

func (*ReturnRuleContext) IsReturnRuleContext() {}

func NewReturnRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnRuleContext {
	var p = new(ReturnRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_returnRule

	return p
}

func (s *ReturnRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnRuleContext) RETURN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserRETURN, 0)
}

func (s *ReturnRuleContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ProyectoParserSEMI, 0)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterReturnRule(s)
	}
}

func (s *ReturnRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitReturnRule(s)
	}
}

func (p *ProyectoParser) ReturnRule() (localctx IReturnRuleContext) {
	localctx = NewReturnRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ProyectoParserRULE_returnRule)
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
		p.SetState(335)
		p.Match(ProyectoParserRETURN)
	}
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProyectoParserT__8)|(1<<ProyectoParserT__9)|(1<<ProyectoParserINT)|(1<<ProyectoParserFLOAT)|(1<<ProyectoParserCHAR)|(1<<ProyectoParserBOOL)|(1<<ProyectoParserLPAREN))) != 0) || _la == ProyectoParserID {
		{
			p.SetState(336)
			p.Exp()
		}

	}
	{
		p.SetState(339)
		p.Match(ProyectoParserSEMI)
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
	p.RuleIndex = ProyectoParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_block

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *ProyectoParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ProyectoParserRULE_block)
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
		p.SetState(341)
		p.Match(ProyectoParserT__0)
	}
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProyectoParserT__8)|(1<<ProyectoParserT__9)|(1<<ProyectoParserINT)|(1<<ProyectoParserFLOAT)|(1<<ProyectoParserCHAR)|(1<<ProyectoParserBOOL)|(1<<ProyectoParserLPAREN)|(1<<ProyectoParserIF))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ProyectoParserWHILE-32))|(1<<(ProyectoParserFOR-32))|(1<<(ProyectoParserWRITE-32))|(1<<(ProyectoParserREAD-32))|(1<<(ProyectoParserRETURN-32))|(1<<(ProyectoParserID-32)))) != 0) {
		{
			p.SetState(342)
			p.Statutes()
		}

		p.SetState(347)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(348)
		p.Match(ProyectoParserT__1)
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
	p.RuleIndex = ProyectoParserRULE_statutesNoReturn
	return p
}

func (*StatutesNoReturnContext) IsStatutesNoReturnContext() {}

func NewStatutesNoReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatutesNoReturnContext {
	var p = new(StatutesNoReturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_statutesNoReturn

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterStatutesNoReturn(s)
	}
}

func (s *StatutesNoReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitStatutesNoReturn(s)
	}
}

func (p *ProyectoParser) StatutesNoReturn() (localctx IStatutesNoReturnContext) {
	localctx = NewStatutesNoReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ProyectoParserRULE_statutesNoReturn)

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

	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(350)
			p.Assignation()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(351)
			p.Read()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(352)
			p.Write()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(353)
			p.Conditional()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(354)
			p.ForLoop()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(355)
			p.WhileLoop()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(356)
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
	p.RuleIndex = ProyectoParserRULE_statutes
	return p
}

func (*StatutesContext) IsStatutesContext() {}

func NewStatutesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatutesContext {
	var p = new(StatutesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_statutes

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterStatutes(s)
	}
}

func (s *StatutesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitStatutes(s)
	}
}

func (p *ProyectoParser) Statutes() (localctx IStatutesContext) {
	localctx = NewStatutesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ProyectoParserRULE_statutes)

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

	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(359)
			p.Assignation()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(360)
			p.Read()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(361)
			p.Write()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(362)
			p.Conditional()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(363)
			p.ForLoop()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(364)
			p.WhileLoop()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(365)
			p.Expression()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(366)
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
	p.RuleIndex = ProyectoParserRULE_assignation
	return p
}

func (*AssignationContext) IsAssignationContext() {}

func NewAssignationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignationContext {
	var p = new(AssignationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_assignation

	return p
}

func (s *AssignationContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserASSIGN, 0)
}

func (s *AssignationContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AssignationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ProyectoParserSEMI, 0)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterAssignation(s)
	}
}

func (s *AssignationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitAssignation(s)
	}
}

func (p *ProyectoParser) Assignation() (localctx IAssignationContext) {
	localctx = NewAssignationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ProyectoParserRULE_assignation)

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
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(369)
			p.Vars()
		}

	case 2:
		{
			p.SetState(370)
			p.VarArray()
		}

	case 3:
		{
			p.SetState(371)
			p.VarMat()
		}

	}
	{
		p.SetState(374)
		p.Match(ProyectoParserASSIGN)
	}
	{
		p.SetState(375)
		p.Exp()
	}
	{
		p.SetState(376)
		p.Match(ProyectoParserSEMI)
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
	p.RuleIndex = ProyectoParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) ID() antlr.TerminalNode {
	return s.GetToken(ProyectoParserID, 0)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserRPAREN, 0)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *ProyectoParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ProyectoParserRULE_functionCall)
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
		p.SetState(378)
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(379)
		p.Match(ProyectoParserLPAREN)
	}
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProyectoParserT__8)|(1<<ProyectoParserT__9)|(1<<ProyectoParserINT)|(1<<ProyectoParserFLOAT)|(1<<ProyectoParserCHAR)|(1<<ProyectoParserBOOL)|(1<<ProyectoParserLPAREN))) != 0) || _la == ProyectoParserID {
		{
			p.SetState(380)
			p.Arguments()
		}

	}
	{
		p.SetState(383)
		p.Match(ProyectoParserRPAREN)
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
	p.RuleIndex = ProyectoParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_arguments

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *ProyectoParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ProyectoParserRULE_arguments)

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
		p.SetState(385)
		p.Exp()
	}
	{
		p.SetState(386)
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
	p.RuleIndex = ProyectoParserRULE_arguments2
	return p
}

func (*Arguments2Context) IsArguments2Context() {}

func NewArguments2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arguments2Context {
	var p = new(Arguments2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_arguments2

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterArguments2(s)
	}
}

func (s *Arguments2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitArguments2(s)
	}
}

func (p *ProyectoParser) Arguments2() (localctx IArguments2Context) {
	localctx = NewArguments2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ProyectoParserRULE_arguments2)

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

	p.SetState(391)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProyectoParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(388)
			p.Match(ProyectoParserT__7)
		}
		{
			p.SetState(389)
			p.Arguments()
		}

	case ProyectoParserRPAREN:
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
	p.RuleIndex = ProyectoParserRULE_methodCall
	return p
}

func (*MethodCallContext) IsMethodCallContext() {}

func NewMethodCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallContext {
	var p = new(MethodCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_methodCall

	return p
}

func (s *MethodCallContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ProyectoParserID)
}

func (s *MethodCallContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ProyectoParserID, i)
}

func (s *MethodCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserLPAREN, 0)
}

func (s *MethodCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserRPAREN, 0)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterMethodCall(s)
	}
}

func (s *MethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitMethodCall(s)
	}
}

func (p *ProyectoParser) MethodCall() (localctx IMethodCallContext) {
	localctx = NewMethodCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ProyectoParserRULE_methodCall)
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
		p.SetState(393)
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(394)
		p.Match(ProyectoParserT__6)
	}
	{
		p.SetState(395)
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(396)
		p.Match(ProyectoParserLPAREN)
	}
	p.SetState(398)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProyectoParserT__8)|(1<<ProyectoParserT__9)|(1<<ProyectoParserINT)|(1<<ProyectoParserFLOAT)|(1<<ProyectoParserCHAR)|(1<<ProyectoParserBOOL)|(1<<ProyectoParserLPAREN))) != 0) || _la == ProyectoParserID {
		{
			p.SetState(397)
			p.Arguments()
		}

	}
	{
		p.SetState(400)
		p.Match(ProyectoParserRPAREN)
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
	p.RuleIndex = ProyectoParserRULE_initCall
	return p
}

func (*InitCallContext) IsInitCallContext() {}

func NewInitCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitCallContext {
	var p = new(InitCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_initCall

	return p
}

func (s *InitCallContext) GetParser() antlr.Parser { return s.parser }

func (s *InitCallContext) ID() antlr.TerminalNode {
	return s.GetToken(ProyectoParserID, 0)
}

func (s *InitCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserLPAREN, 0)
}

func (s *InitCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserRPAREN, 0)
}

func (s *InitCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterInitCall(s)
	}
}

func (s *InitCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitInitCall(s)
	}
}

func (p *ProyectoParser) InitCall() (localctx IInitCallContext) {
	localctx = NewInitCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ProyectoParserRULE_initCall)

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
		p.SetState(402)
		p.Match(ProyectoParserT__8)
	}
	{
		p.SetState(403)
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(404)
		p.Match(ProyectoParserLPAREN)
	}
	{
		p.SetState(405)
		p.Match(ProyectoParserRPAREN)
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
	p.RuleIndex = ProyectoParserRULE_call
	return p
}

func (*CallContext) IsCallContext() {}

func NewCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallContext {
	var p = new(CallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_call

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterCall(s)
	}
}

func (s *CallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitCall(s)
	}
}

func (p *ProyectoParser) Call() (localctx ICallContext) {
	localctx = NewCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ProyectoParserRULE_call)

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

	p.SetState(410)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(407)
			p.FunctionCall()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(408)
			p.MethodCall()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(409)
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
	p.RuleIndex = ProyectoParserRULE_read
	return p
}

func (*ReadContext) IsReadContext() {}

func NewReadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReadContext {
	var p = new(ReadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_read

	return p
}

func (s *ReadContext) GetParser() antlr.Parser { return s.parser }

func (s *ReadContext) READ() antlr.TerminalNode {
	return s.GetToken(ProyectoParserREAD, 0)
}

func (s *ReadContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserLPAREN, 0)
}

func (s *ReadContext) Read2() IRead2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRead2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRead2Context)
}

func (s *ReadContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserRPAREN, 0)
}

func (s *ReadContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ProyectoParserSEMI, 0)
}

func (s *ReadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterRead(s)
	}
}

func (s *ReadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitRead(s)
	}
}

func (p *ProyectoParser) Read() (localctx IReadContext) {
	localctx = NewReadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ProyectoParserRULE_read)

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
		p.SetState(412)
		p.Match(ProyectoParserREAD)
	}
	{
		p.SetState(413)
		p.Match(ProyectoParserLPAREN)
	}
	{
		p.SetState(414)
		p.Read2()
	}
	{
		p.SetState(415)
		p.Match(ProyectoParserRPAREN)
	}
	{
		p.SetState(416)
		p.Match(ProyectoParserSEMI)
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
	p.RuleIndex = ProyectoParserRULE_read2
	return p
}

func (*Read2Context) IsRead2Context() {}

func NewRead2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Read2Context {
	var p = new(Read2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_read2

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterRead2(s)
	}
}

func (s *Read2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitRead2(s)
	}
}

func (p *ProyectoParser) Read2() (localctx IRead2Context) {
	localctx = NewRead2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ProyectoParserRULE_read2)

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
	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(418)
			p.Vars()
		}

	case 2:
		{
			p.SetState(419)
			p.VarArray()
		}

	case 3:
		{
			p.SetState(420)
			p.VarMat()
		}

	}
	{
		p.SetState(423)
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
	p.RuleIndex = ProyectoParserRULE_read3
	return p
}

func (*Read3Context) IsRead3Context() {}

func NewRead3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Read3Context {
	var p = new(Read3Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_read3

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterRead3(s)
	}
}

func (s *Read3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitRead3(s)
	}
}

func (p *ProyectoParser) Read3() (localctx IRead3Context) {
	localctx = NewRead3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ProyectoParserRULE_read3)

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

	p.SetState(428)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProyectoParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(425)
			p.Match(ProyectoParserT__7)
		}
		{
			p.SetState(426)
			p.Read2()
		}

	case ProyectoParserRPAREN:
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
	p.RuleIndex = ProyectoParserRULE_write
	return p
}

func (*WriteContext) IsWriteContext() {}

func NewWriteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WriteContext {
	var p = new(WriteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_write

	return p
}

func (s *WriteContext) GetParser() antlr.Parser { return s.parser }

func (s *WriteContext) WRITE() antlr.TerminalNode {
	return s.GetToken(ProyectoParserWRITE, 0)
}

func (s *WriteContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserLPAREN, 0)
}

func (s *WriteContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *WriteContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserRPAREN, 0)
}

func (s *WriteContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ProyectoParserSEMI, 0)
}

func (s *WriteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WriteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WriteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterWrite(s)
	}
}

func (s *WriteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitWrite(s)
	}
}

func (p *ProyectoParser) Write() (localctx IWriteContext) {
	localctx = NewWriteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ProyectoParserRULE_write)

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
		p.SetState(430)
		p.Match(ProyectoParserWRITE)
	}
	{
		p.SetState(431)
		p.Match(ProyectoParserLPAREN)
	}
	{
		p.SetState(432)
		p.Arguments()
	}
	{
		p.SetState(433)
		p.Match(ProyectoParserRPAREN)
	}
	{
		p.SetState(434)
		p.Match(ProyectoParserSEMI)
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
	p.RuleIndex = ProyectoParserRULE_conditional
	return p
}

func (*ConditionalContext) IsConditionalContext() {}

func NewConditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalContext {
	var p = new(ConditionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_conditional

	return p
}

func (s *ConditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalContext) IF() antlr.TerminalNode {
	return s.GetToken(ProyectoParserIF, 0)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterConditional(s)
	}
}

func (s *ConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitConditional(s)
	}
}

func (p *ProyectoParser) Conditional() (localctx IConditionalContext) {
	localctx = NewConditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ProyectoParserRULE_conditional)
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
		p.SetState(436)
		p.Match(ProyectoParserIF)
	}
	{
		p.SetState(437)
		p.Conditional2()
	}
	{
		p.SetState(438)
		p.Conditional3()
	}
	p.SetState(440)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProyectoParserELSE {
		{
			p.SetState(439)
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
	p.RuleIndex = ProyectoParserRULE_conditional2
	return p
}

func (*Conditional2Context) IsConditional2Context() {}

func NewConditional2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Conditional2Context {
	var p = new(Conditional2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_conditional2

	return p
}

func (s *Conditional2Context) GetParser() antlr.Parser { return s.parser }

func (s *Conditional2Context) LPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserLPAREN, 0)
}

func (s *Conditional2Context) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *Conditional2Context) RPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserRPAREN, 0)
}

func (s *Conditional2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Conditional2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Conditional2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterConditional2(s)
	}
}

func (s *Conditional2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitConditional2(s)
	}
}

func (p *ProyectoParser) Conditional2() (localctx IConditional2Context) {
	localctx = NewConditional2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ProyectoParserRULE_conditional2)

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
		p.Match(ProyectoParserLPAREN)
	}
	{
		p.SetState(443)
		p.Exp()
	}
	{
		p.SetState(444)
		p.Match(ProyectoParserRPAREN)
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
	p.RuleIndex = ProyectoParserRULE_conditional3
	return p
}

func (*Conditional3Context) IsConditional3Context() {}

func NewConditional3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Conditional3Context {
	var p = new(Conditional3Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_conditional3

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterConditional3(s)
	}
}

func (s *Conditional3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitConditional3(s)
	}
}

func (p *ProyectoParser) Conditional3() (localctx IConditional3Context) {
	localctx = NewConditional3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ProyectoParserRULE_conditional3)

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
		p.SetState(446)
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
	p.RuleIndex = ProyectoParserRULE_conditional4
	return p
}

func (*Conditional4Context) IsConditional4Context() {}

func NewConditional4Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Conditional4Context {
	var p = new(Conditional4Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_conditional4

	return p
}

func (s *Conditional4Context) GetParser() antlr.Parser { return s.parser }

func (s *Conditional4Context) ELSE() antlr.TerminalNode {
	return s.GetToken(ProyectoParserELSE, 0)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterConditional4(s)
	}
}

func (s *Conditional4Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitConditional4(s)
	}
}

func (p *ProyectoParser) Conditional4() (localctx IConditional4Context) {
	localctx = NewConditional4Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ProyectoParserRULE_conditional4)

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
		p.Match(ProyectoParserELSE)
	}
	{
		p.SetState(449)
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
	p.RuleIndex = ProyectoParserRULE_forLoop
	return p
}

func (*ForLoopContext) IsForLoopContext() {}

func NewForLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForLoopContext {
	var p = new(ForLoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_forLoop

	return p
}

func (s *ForLoopContext) GetParser() antlr.Parser { return s.parser }

func (s *ForLoopContext) FOR() antlr.TerminalNode {
	return s.GetToken(ProyectoParserFOR, 0)
}

func (s *ForLoopContext) ForLoop2() IForLoop2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForLoop2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForLoop2Context)
}

func (s *ForLoopContext) IN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserIN, 0)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterForLoop(s)
	}
}

func (s *ForLoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitForLoop(s)
	}
}

func (p *ProyectoParser) ForLoop() (localctx IForLoopContext) {
	localctx = NewForLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, ProyectoParserRULE_forLoop)

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
		p.SetState(451)
		p.Match(ProyectoParserFOR)
	}
	{
		p.SetState(452)
		p.ForLoop2()
	}
	{
		p.SetState(453)
		p.Match(ProyectoParserIN)
	}
	{
		p.SetState(454)
		p.ForLoop3()
	}
	{
		p.SetState(455)
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
	p.RuleIndex = ProyectoParserRULE_forLoop2
	return p
}

func (*ForLoop2Context) IsForLoop2Context() {}

func NewForLoop2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForLoop2Context {
	var p = new(ForLoop2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_forLoop2

	return p
}

func (s *ForLoop2Context) GetParser() antlr.Parser { return s.parser }

func (s *ForLoop2Context) ID() antlr.TerminalNode {
	return s.GetToken(ProyectoParserID, 0)
}

func (s *ForLoop2Context) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserASSIGN, 0)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterForLoop2(s)
	}
}

func (s *ForLoop2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitForLoop2(s)
	}
}

func (p *ProyectoParser) ForLoop2() (localctx IForLoop2Context) {
	localctx = NewForLoop2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, ProyectoParserRULE_forLoop2)

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
		p.SetState(457)
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(458)
		p.Match(ProyectoParserASSIGN)
	}
	{
		p.SetState(459)
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
	p.RuleIndex = ProyectoParserRULE_forLoop3
	return p
}

func (*ForLoop3Context) IsForLoop3Context() {}

func NewForLoop3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForLoop3Context {
	var p = new(ForLoop3Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_forLoop3

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterForLoop3(s)
	}
}

func (s *ForLoop3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitForLoop3(s)
	}
}

func (p *ProyectoParser) ForLoop3() (localctx IForLoop3Context) {
	localctx = NewForLoop3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, ProyectoParserRULE_forLoop3)

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
		p.SetState(461)
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
	p.RuleIndex = ProyectoParserRULE_whileLoop
	return p
}

func (*WhileLoopContext) IsWhileLoopContext() {}

func NewWhileLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileLoopContext {
	var p = new(WhileLoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_whileLoop

	return p
}

func (s *WhileLoopContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileLoopContext) WHILE() antlr.TerminalNode {
	return s.GetToken(ProyectoParserWHILE, 0)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterWhileLoop(s)
	}
}

func (s *WhileLoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitWhileLoop(s)
	}
}

func (p *ProyectoParser) WhileLoop() (localctx IWhileLoopContext) {
	localctx = NewWhileLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, ProyectoParserRULE_whileLoop)

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
		p.SetState(463)
		p.Match(ProyectoParserWHILE)
	}
	{
		p.SetState(464)
		p.WhileLoop2()
	}
	{
		p.SetState(465)
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
	p.RuleIndex = ProyectoParserRULE_whileLoop2
	return p
}

func (*WhileLoop2Context) IsWhileLoop2Context() {}

func NewWhileLoop2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileLoop2Context {
	var p = new(WhileLoop2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_whileLoop2

	return p
}

func (s *WhileLoop2Context) GetParser() antlr.Parser { return s.parser }

func (s *WhileLoop2Context) LPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserLPAREN, 0)
}

func (s *WhileLoop2Context) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *WhileLoop2Context) RPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserRPAREN, 0)
}

func (s *WhileLoop2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileLoop2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileLoop2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterWhileLoop2(s)
	}
}

func (s *WhileLoop2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitWhileLoop2(s)
	}
}

func (p *ProyectoParser) WhileLoop2() (localctx IWhileLoop2Context) {
	localctx = NewWhileLoop2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, ProyectoParserRULE_whileLoop2)

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
		p.SetState(467)
		p.Match(ProyectoParserLPAREN)
	}
	{
		p.SetState(468)
		p.Exp()
	}
	{
		p.SetState(469)
		p.Match(ProyectoParserRPAREN)
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
	p.RuleIndex = ProyectoParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_expression

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
	return s.GetToken(ProyectoParserSEMI, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ProyectoParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, ProyectoParserRULE_expression)

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
		p.Exp()
	}
	{
		p.SetState(472)
		p.Match(ProyectoParserSEMI)
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
	p.RuleIndex = ProyectoParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_exp

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitExp(s)
	}
}

func (p *ProyectoParser) Exp() (localctx IExpContext) {
	localctx = NewExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, ProyectoParserRULE_exp)
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
		p.SetState(474)
		p.T_exp()
	}
	p.SetState(478)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserOR {
		{
			p.SetState(475)
			p.Exp2()
		}

		p.SetState(480)
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
	p.RuleIndex = ProyectoParserRULE_exp2
	return p
}

func (*Exp2Context) IsExp2Context() {}

func NewExp2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp2Context {
	var p = new(Exp2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_exp2

	return p
}

func (s *Exp2Context) GetParser() antlr.Parser { return s.parser }

func (s *Exp2Context) OR() antlr.TerminalNode {
	return s.GetToken(ProyectoParserOR, 0)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterExp2(s)
	}
}

func (s *Exp2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitExp2(s)
	}
}

func (p *ProyectoParser) Exp2() (localctx IExp2Context) {
	localctx = NewExp2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, ProyectoParserRULE_exp2)

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
		p.Match(ProyectoParserOR)
	}
	{
		p.SetState(482)
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
	p.RuleIndex = ProyectoParserRULE_t_exp
	return p
}

func (*T_expContext) IsT_expContext() {}

func NewT_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *T_expContext {
	var p = new(T_expContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_t_exp

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterT_exp(s)
	}
}

func (s *T_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitT_exp(s)
	}
}

func (p *ProyectoParser) T_exp() (localctx IT_expContext) {
	localctx = NewT_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, ProyectoParserRULE_t_exp)
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
		p.SetState(484)
		p.G_exp()
	}
	p.SetState(488)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserAND {
		{
			p.SetState(485)
			p.T_exp2()
		}

		p.SetState(490)
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
	p.RuleIndex = ProyectoParserRULE_t_exp2
	return p
}

func (*T_exp2Context) IsT_exp2Context() {}

func NewT_exp2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *T_exp2Context {
	var p = new(T_exp2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_t_exp2

	return p
}

func (s *T_exp2Context) GetParser() antlr.Parser { return s.parser }

func (s *T_exp2Context) AND() antlr.TerminalNode {
	return s.GetToken(ProyectoParserAND, 0)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterT_exp2(s)
	}
}

func (s *T_exp2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitT_exp2(s)
	}
}

func (p *ProyectoParser) T_exp2() (localctx IT_exp2Context) {
	localctx = NewT_exp2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, ProyectoParserRULE_t_exp2)

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
		p.Match(ProyectoParserAND)
	}
	{
		p.SetState(492)
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
	p.RuleIndex = ProyectoParserRULE_g_exp
	return p
}

func (*G_expContext) IsG_expContext() {}

func NewG_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *G_expContext {
	var p = new(G_expContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_g_exp

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterG_exp(s)
	}
}

func (s *G_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitG_exp(s)
	}
}

func (p *ProyectoParser) G_exp() (localctx IG_expContext) {
	localctx = NewG_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, ProyectoParserRULE_g_exp)
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
		p.M_exp()
	}
	p.SetState(496)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProyectoParserGT)|(1<<ProyectoParserLT)|(1<<ProyectoParserEQ)|(1<<ProyectoParserNEQ))) != 0 {
		{
			p.SetState(495)
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
	p.RuleIndex = ProyectoParserRULE_g_exp2
	return p
}

func (*G_exp2Context) IsG_exp2Context() {}

func NewG_exp2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *G_exp2Context {
	var p = new(G_exp2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_g_exp2

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterG_exp2(s)
	}
}

func (s *G_exp2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitG_exp2(s)
	}
}

func (p *ProyectoParser) G_exp2() (localctx IG_exp2Context) {
	localctx = NewG_exp2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, ProyectoParserRULE_g_exp2)

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
		p.SetState(498)
		p.Relop()
	}
	{
		p.SetState(499)
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
	p.RuleIndex = ProyectoParserRULE_m_exp
	return p
}

func (*M_expContext) IsM_expContext() {}

func NewM_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *M_expContext {
	var p = new(M_expContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_m_exp

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterM_exp(s)
	}
}

func (s *M_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitM_exp(s)
	}
}

func (p *ProyectoParser) M_exp() (localctx IM_expContext) {
	localctx = NewM_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, ProyectoParserRULE_m_exp)
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
		p.SetState(501)
		p.Term()
	}
	p.SetState(505)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserADD || _la == ProyectoParserSUB {
		{
			p.SetState(502)
			p.M_exp2()
		}

		p.SetState(507)
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
	p.RuleIndex = ProyectoParserRULE_m_exp2
	return p
}

func (*M_exp2Context) IsM_exp2Context() {}

func NewM_exp2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *M_exp2Context {
	var p = new(M_exp2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_m_exp2

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
	return s.GetToken(ProyectoParserADD, 0)
}

func (s *M_exp2Context) SUB() antlr.TerminalNode {
	return s.GetToken(ProyectoParserSUB, 0)
}

func (s *M_exp2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *M_exp2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *M_exp2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterM_exp2(s)
	}
}

func (s *M_exp2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitM_exp2(s)
	}
}

func (p *ProyectoParser) M_exp2() (localctx IM_exp2Context) {
	localctx = NewM_exp2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, ProyectoParserRULE_m_exp2)
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
		p.SetState(508)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ProyectoParserADD || _la == ProyectoParserSUB) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(509)
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
	p.RuleIndex = ProyectoParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_term

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *ProyectoParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, ProyectoParserRULE_term)
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
		p.SetState(511)
		p.Factor()
	}
	p.SetState(515)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserMUL || _la == ProyectoParserDIV {
		{
			p.SetState(512)
			p.Term2()
		}

		p.SetState(517)
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
	p.RuleIndex = ProyectoParserRULE_term2
	return p
}

func (*Term2Context) IsTerm2Context() {}

func NewTerm2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Term2Context {
	var p = new(Term2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_term2

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
	return s.GetToken(ProyectoParserMUL, 0)
}

func (s *Term2Context) DIV() antlr.TerminalNode {
	return s.GetToken(ProyectoParserDIV, 0)
}

func (s *Term2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Term2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Term2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterTerm2(s)
	}
}

func (s *Term2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitTerm2(s)
	}
}

func (p *ProyectoParser) Term2() (localctx ITerm2Context) {
	localctx = NewTerm2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, ProyectoParserRULE_term2)
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
		p.SetState(518)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ProyectoParserMUL || _la == ProyectoParserDIV) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(519)
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
	p.RuleIndex = ProyectoParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_factor

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *ProyectoParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, ProyectoParserRULE_factor)

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

	p.SetState(529)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(521)
			p.Factor2()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(522)
			p.VarCte()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(526)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(523)
				p.Vars()
			}

		case 2:
			{
				p.SetState(524)
				p.VarArray()
			}

		case 3:
			{
				p.SetState(525)
				p.VarMat()
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(528)
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
	p.RuleIndex = ProyectoParserRULE_factor2
	return p
}

func (*Factor2Context) IsFactor2Context() {}

func NewFactor2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Factor2Context {
	var p = new(Factor2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_factor2

	return p
}

func (s *Factor2Context) GetParser() antlr.Parser { return s.parser }

func (s *Factor2Context) LPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserLPAREN, 0)
}

func (s *Factor2Context) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *Factor2Context) RPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserRPAREN, 0)
}

func (s *Factor2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Factor2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Factor2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterFactor2(s)
	}
}

func (s *Factor2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitFactor2(s)
	}
}

func (p *ProyectoParser) Factor2() (localctx IFactor2Context) {
	localctx = NewFactor2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, ProyectoParserRULE_factor2)

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
		p.Match(ProyectoParserLPAREN)
	}
	{
		p.SetState(532)
		p.Exp()
	}
	{
		p.SetState(533)
		p.Match(ProyectoParserRPAREN)
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
	p.RuleIndex = ProyectoParserRULE_varCte
	return p
}

func (*VarCteContext) IsVarCteContext() {}

func NewVarCteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarCteContext {
	var p = new(VarCteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_varCte

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterVarCte(s)
	}
}

func (s *VarCteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitVarCte(s)
	}
}

func (p *ProyectoParser) VarCte() (localctx IVarCteContext) {
	localctx = NewVarCteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, ProyectoParserRULE_varCte)

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

	p.SetState(540)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProyectoParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(535)
			p.Cte_i()
		}

	case ProyectoParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(536)
			p.Cte_f()
		}

	case ProyectoParserCHAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(537)
			p.Cte_c()
		}

	case ProyectoParserT__9:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(538)
			p.Cte_s()
		}

	case ProyectoParserBOOL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(539)
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
	p.RuleIndex = ProyectoParserRULE_cte_i
	return p
}

func (*Cte_iContext) IsCte_iContext() {}

func NewCte_iContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cte_iContext {
	var p = new(Cte_iContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_cte_i

	return p
}

func (s *Cte_iContext) GetParser() antlr.Parser { return s.parser }

func (s *Cte_iContext) INT() antlr.TerminalNode {
	return s.GetToken(ProyectoParserINT, 0)
}

func (s *Cte_iContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cte_iContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cte_iContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterCte_i(s)
	}
}

func (s *Cte_iContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitCte_i(s)
	}
}

func (p *ProyectoParser) Cte_i() (localctx ICte_iContext) {
	localctx = NewCte_iContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, ProyectoParserRULE_cte_i)

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
		p.SetState(542)
		p.Match(ProyectoParserINT)
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
	p.RuleIndex = ProyectoParserRULE_cte_f
	return p
}

func (*Cte_fContext) IsCte_fContext() {}

func NewCte_fContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cte_fContext {
	var p = new(Cte_fContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_cte_f

	return p
}

func (s *Cte_fContext) GetParser() antlr.Parser { return s.parser }

func (s *Cte_fContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ProyectoParserFLOAT, 0)
}

func (s *Cte_fContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cte_fContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cte_fContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterCte_f(s)
	}
}

func (s *Cte_fContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitCte_f(s)
	}
}

func (p *ProyectoParser) Cte_f() (localctx ICte_fContext) {
	localctx = NewCte_fContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, ProyectoParserRULE_cte_f)

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
		p.SetState(544)
		p.Match(ProyectoParserFLOAT)
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
	p.RuleIndex = ProyectoParserRULE_cte_c
	return p
}

func (*Cte_cContext) IsCte_cContext() {}

func NewCte_cContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cte_cContext {
	var p = new(Cte_cContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_cte_c

	return p
}

func (s *Cte_cContext) GetParser() antlr.Parser { return s.parser }

func (s *Cte_cContext) CHAR() antlr.TerminalNode {
	return s.GetToken(ProyectoParserCHAR, 0)
}

func (s *Cte_cContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cte_cContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cte_cContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterCte_c(s)
	}
}

func (s *Cte_cContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitCte_c(s)
	}
}

func (p *ProyectoParser) Cte_c() (localctx ICte_cContext) {
	localctx = NewCte_cContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, ProyectoParserRULE_cte_c)

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
		p.SetState(546)
		p.Match(ProyectoParserCHAR)
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
	p.RuleIndex = ProyectoParserRULE_cte_b
	return p
}

func (*Cte_bContext) IsCte_bContext() {}

func NewCte_bContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cte_bContext {
	var p = new(Cte_bContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_cte_b

	return p
}

func (s *Cte_bContext) GetParser() antlr.Parser { return s.parser }

func (s *Cte_bContext) BOOL() antlr.TerminalNode {
	return s.GetToken(ProyectoParserBOOL, 0)
}

func (s *Cte_bContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cte_bContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cte_bContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterCte_b(s)
	}
}

func (s *Cte_bContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitCte_b(s)
	}
}

func (p *ProyectoParser) Cte_b() (localctx ICte_bContext) {
	localctx = NewCte_bContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, ProyectoParserRULE_cte_b)

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
		p.SetState(548)
		p.Match(ProyectoParserBOOL)
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
	p.RuleIndex = ProyectoParserRULE_cte_s
	return p
}

func (*Cte_sContext) IsCte_sContext() {}

func NewCte_sContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cte_sContext {
	var p = new(Cte_sContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_cte_s

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterCte_s(s)
	}
}

func (s *Cte_sContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitCte_s(s)
	}
}

func (p *ProyectoParser) Cte_s() (localctx ICte_sContext) {
	localctx = NewCte_sContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, ProyectoParserRULE_cte_s)

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
		p.SetState(550)
		p.Match(ProyectoParserT__9)
	}
	p.SetState(554)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(551)
			p.MatchWildcard()

		}
		p.SetState(556)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())
	}
	{
		p.SetState(557)
		p.Match(ProyectoParserT__9)
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
	p.RuleIndex = ProyectoParserRULE_main
	return p
}

func (*MainContext) IsMainContext() {}

func NewMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainContext {
	var p = new(MainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_main

	return p
}

func (s *MainContext) GetParser() antlr.Parser { return s.parser }

func (s *MainContext) MAIN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserMAIN, 0)
}

func (s *MainContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserLPAREN, 0)
}

func (s *MainContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ProyectoParserRPAREN, 0)
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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterMain(s)
	}
}

func (s *MainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitMain(s)
	}
}

func (p *ProyectoParser) Main() (localctx IMainContext) {
	localctx = NewMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, ProyectoParserRULE_main)

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
		p.SetState(559)
		p.Match(ProyectoParserMAIN)
	}
	{
		p.SetState(560)
		p.Match(ProyectoParserLPAREN)
	}
	{
		p.SetState(561)
		p.Match(ProyectoParserRPAREN)
	}
	{
		p.SetState(562)
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
	p.RuleIndex = ProyectoParserRULE_mainBlock
	return p
}

func (*MainBlockContext) IsMainBlockContext() {}

func NewMainBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainBlockContext {
	var p = new(MainBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_mainBlock

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
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterMainBlock(s)
	}
}

func (s *MainBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitMainBlock(s)
	}
}

func (p *ProyectoParser) MainBlock() (localctx IMainBlockContext) {
	localctx = NewMainBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, ProyectoParserRULE_mainBlock)
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
		p.SetState(564)
		p.Match(ProyectoParserT__0)
	}
	p.SetState(568)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserVAR {
		{
			p.SetState(565)
			p.VariableDeclaration()
		}

		p.SetState(570)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(574)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProyectoParserT__8)|(1<<ProyectoParserT__9)|(1<<ProyectoParserINT)|(1<<ProyectoParserFLOAT)|(1<<ProyectoParserCHAR)|(1<<ProyectoParserBOOL)|(1<<ProyectoParserLPAREN)|(1<<ProyectoParserIF))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ProyectoParserWHILE-32))|(1<<(ProyectoParserFOR-32))|(1<<(ProyectoParserWRITE-32))|(1<<(ProyectoParserREAD-32))|(1<<(ProyectoParserRETURN-32))|(1<<(ProyectoParserID-32)))) != 0) {
		{
			p.SetState(571)
			p.Statutes()
		}

		p.SetState(576)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(577)
		p.Match(ProyectoParserT__1)
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
	p.RuleIndex = ProyectoParserRULE_typeRule
	return p
}

func (*TypeRuleContext) IsTypeRuleContext() {}

func NewTypeRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeRuleContext {
	var p = new(TypeRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_typeRule

	return p
}

func (s *TypeRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeRuleContext) INT_TYPE() antlr.TerminalNode {
	return s.GetToken(ProyectoParserINT_TYPE, 0)
}

func (s *TypeRuleContext) FLOAT_TYPE() antlr.TerminalNode {
	return s.GetToken(ProyectoParserFLOAT_TYPE, 0)
}

func (s *TypeRuleContext) CHAR_TYPE() antlr.TerminalNode {
	return s.GetToken(ProyectoParserCHAR_TYPE, 0)
}

func (s *TypeRuleContext) BOOL_TYPE() antlr.TerminalNode {
	return s.GetToken(ProyectoParserBOOL_TYPE, 0)
}

func (s *TypeRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterTypeRule(s)
	}
}

func (s *TypeRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitTypeRule(s)
	}
}

func (p *ProyectoParser) TypeRule() (localctx ITypeRuleContext) {
	localctx = NewTypeRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, ProyectoParserRULE_typeRule)
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
		p.SetState(579)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(ProyectoParserINT_TYPE-45))|(1<<(ProyectoParserFLOAT_TYPE-45))|(1<<(ProyectoParserCHAR_TYPE-45))|(1<<(ProyectoParserBOOL_TYPE-45)))) != 0) {
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
	p.RuleIndex = ProyectoParserRULE_relop
	return p
}

func (*RelopContext) IsRelopContext() {}

func NewRelopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelopContext {
	var p = new(RelopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProyectoParserRULE_relop

	return p
}

func (s *RelopContext) GetParser() antlr.Parser { return s.parser }

func (s *RelopContext) GT() antlr.TerminalNode {
	return s.GetToken(ProyectoParserGT, 0)
}

func (s *RelopContext) LT() antlr.TerminalNode {
	return s.GetToken(ProyectoParserLT, 0)
}

func (s *RelopContext) EQ() antlr.TerminalNode {
	return s.GetToken(ProyectoParserEQ, 0)
}

func (s *RelopContext) NEQ() antlr.TerminalNode {
	return s.GetToken(ProyectoParserNEQ, 0)
}

func (s *RelopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.EnterRelop(s)
	}
}

func (s *RelopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProyectoListener); ok {
		listenerT.ExitRelop(s)
	}
}

func (p *ProyectoParser) Relop() (localctx IRelopContext) {
	localctx = NewRelopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, ProyectoParserRULE_relop)
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
		p.SetState(581)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProyectoParserGT)|(1<<ProyectoParserLT)|(1<<ProyectoParserEQ)|(1<<ProyectoParserNEQ))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
