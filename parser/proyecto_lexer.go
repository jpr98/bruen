// Code generated from Proyecto.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 52, 322,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 11, 3, 11, 3, 12, 6, 12, 126, 10, 12, 13, 12, 14, 12, 127, 3, 13,
	3, 13, 3, 13, 5, 13, 133, 10, 13, 3, 13, 3, 13, 5, 13, 137, 10, 13, 3,
	13, 5, 13, 140, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 155, 10, 15, 3, 16, 6,
	16, 158, 10, 16, 13, 16, 14, 16, 159, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17,
	3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3,
	22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26,
	3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36,
	3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39,
	3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3,
	41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42,
	3, 43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3,
	45, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47,
	3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3,
	49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50,
	3, 50, 3, 50, 3, 51, 3, 51, 7, 51, 318, 10, 51, 12, 51, 14, 51, 321, 11,
	51, 2, 2, 52, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37,
	20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55,
	29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73,
	38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91,
	47, 93, 48, 95, 49, 97, 50, 99, 51, 101, 52, 3, 2, 8, 3, 2, 50, 59, 4,
	2, 45, 45, 47, 47, 4, 2, 67, 92, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34,
	5, 2, 67, 92, 97, 97, 99, 124, 5, 2, 50, 59, 67, 92, 99, 124, 2, 328, 2,
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2,
	11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2,
	2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2,
	2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2,
	2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3,
	2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49,
	3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2,
	57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2,
	2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2,
	2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2,
	2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3,
	2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95,
	3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 3,
	103, 3, 2, 2, 2, 5, 105, 3, 2, 2, 2, 7, 107, 3, 2, 2, 2, 9, 109, 3, 2,
	2, 2, 11, 111, 3, 2, 2, 2, 13, 113, 3, 2, 2, 2, 15, 116, 3, 2, 2, 2, 17,
	118, 3, 2, 2, 2, 19, 120, 3, 2, 2, 2, 21, 122, 3, 2, 2, 2, 23, 125, 3,
	2, 2, 2, 25, 129, 3, 2, 2, 2, 27, 141, 3, 2, 2, 2, 29, 154, 3, 2, 2, 2,
	31, 157, 3, 2, 2, 2, 33, 163, 3, 2, 2, 2, 35, 166, 3, 2, 2, 2, 37, 169,
	3, 2, 2, 2, 39, 171, 3, 2, 2, 2, 41, 173, 3, 2, 2, 2, 43, 176, 3, 2, 2,
	2, 45, 179, 3, 2, 2, 2, 47, 181, 3, 2, 2, 2, 49, 183, 3, 2, 2, 2, 51, 185,
	3, 2, 2, 2, 53, 187, 3, 2, 2, 2, 55, 189, 3, 2, 2, 2, 57, 191, 3, 2, 2,
	2, 59, 193, 3, 2, 2, 2, 61, 196, 3, 2, 2, 2, 63, 201, 3, 2, 2, 2, 65, 207,
	3, 2, 2, 2, 67, 211, 3, 2, 2, 2, 69, 214, 3, 2, 2, 2, 71, 220, 3, 2, 2,
	2, 73, 231, 3, 2, 2, 2, 75, 239, 3, 2, 2, 2, 77, 245, 3, 2, 2, 2, 79, 250,
	3, 2, 2, 2, 81, 259, 3, 2, 2, 2, 83, 266, 3, 2, 2, 2, 85, 271, 3, 2, 2,
	2, 87, 275, 3, 2, 2, 2, 89, 279, 3, 2, 2, 2, 91, 285, 3, 2, 2, 2, 93, 290,
	3, 2, 2, 2, 95, 297, 3, 2, 2, 2, 97, 302, 3, 2, 2, 2, 99, 307, 3, 2, 2,
	2, 101, 315, 3, 2, 2, 2, 103, 104, 7, 125, 2, 2, 104, 4, 3, 2, 2, 2, 105,
	106, 7, 127, 2, 2, 106, 6, 3, 2, 2, 2, 107, 108, 7, 60, 2, 2, 108, 8, 3,
	2, 2, 2, 109, 110, 7, 93, 2, 2, 110, 10, 3, 2, 2, 2, 111, 112, 7, 95, 2,
	2, 112, 12, 3, 2, 2, 2, 113, 114, 7, 95, 2, 2, 114, 115, 7, 93, 2, 2, 115,
	14, 3, 2, 2, 2, 116, 117, 7, 63, 2, 2, 117, 16, 3, 2, 2, 2, 118, 119, 7,
	48, 2, 2, 119, 18, 3, 2, 2, 2, 120, 121, 7, 46, 2, 2, 121, 20, 3, 2, 2,
	2, 122, 123, 7, 36, 2, 2, 123, 22, 3, 2, 2, 2, 124, 126, 9, 2, 2, 2, 125,
	124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 127, 128,
	3, 2, 2, 2, 128, 24, 3, 2, 2, 2, 129, 132, 5, 23, 12, 2, 130, 131, 7, 48,
	2, 2, 131, 133, 5, 23, 12, 2, 132, 130, 3, 2, 2, 2, 132, 133, 3, 2, 2,
	2, 133, 139, 3, 2, 2, 2, 134, 136, 7, 71, 2, 2, 135, 137, 9, 3, 2, 2, 136,
	135, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138, 140,
	5, 23, 12, 2, 139, 134, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 26, 3, 2,
	2, 2, 141, 142, 7, 36, 2, 2, 142, 143, 9, 4, 2, 2, 143, 144, 7, 36, 2,
	2, 144, 28, 3, 2, 2, 2, 145, 146, 7, 118, 2, 2, 146, 147, 7, 116, 2, 2,
	147, 148, 7, 119, 2, 2, 148, 155, 7, 103, 2, 2, 149, 150, 7, 104, 2, 2,
	150, 151, 7, 99, 2, 2, 151, 152, 7, 110, 2, 2, 152, 153, 7, 117, 2, 2,
	153, 155, 7, 103, 2, 2, 154, 145, 3, 2, 2, 2, 154, 149, 3, 2, 2, 2, 155,
	30, 3, 2, 2, 2, 156, 158, 9, 5, 2, 2, 157, 156, 3, 2, 2, 2, 158, 159, 3,
	2, 2, 2, 159, 157, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 161, 3, 2, 2,
	2, 161, 162, 8, 16, 2, 2, 162, 32, 3, 2, 2, 2, 163, 164, 7, 126, 2, 2,
	164, 165, 7, 126, 2, 2, 165, 34, 3, 2, 2, 2, 166, 167, 7, 40, 2, 2, 167,
	168, 7, 40, 2, 2, 168, 36, 3, 2, 2, 2, 169, 170, 7, 64, 2, 2, 170, 38,
	3, 2, 2, 2, 171, 172, 7, 62, 2, 2, 172, 40, 3, 2, 2, 2, 173, 174, 7, 63,
	2, 2, 174, 175, 7, 63, 2, 2, 175, 42, 3, 2, 2, 2, 176, 177, 7, 35, 2, 2,
	177, 178, 7, 63, 2, 2, 178, 44, 3, 2, 2, 2, 179, 180, 7, 44, 2, 2, 180,
	46, 3, 2, 2, 2, 181, 182, 7, 49, 2, 2, 182, 48, 3, 2, 2, 2, 183, 184, 7,
	45, 2, 2, 184, 50, 3, 2, 2, 2, 185, 186, 7, 47, 2, 2, 186, 52, 3, 2, 2,
	2, 187, 188, 7, 61, 2, 2, 188, 54, 3, 2, 2, 2, 189, 190, 7, 42, 2, 2, 190,
	56, 3, 2, 2, 2, 191, 192, 7, 43, 2, 2, 192, 58, 3, 2, 2, 2, 193, 194, 7,
	107, 2, 2, 194, 195, 7, 104, 2, 2, 195, 60, 3, 2, 2, 2, 196, 197, 7, 103,
	2, 2, 197, 198, 7, 110, 2, 2, 198, 199, 7, 117, 2, 2, 199, 200, 7, 103,
	2, 2, 200, 62, 3, 2, 2, 2, 201, 202, 7, 121, 2, 2, 202, 203, 7, 106, 2,
	2, 203, 204, 7, 107, 2, 2, 204, 205, 7, 110, 2, 2, 205, 206, 7, 103, 2,
	2, 206, 64, 3, 2, 2, 2, 207, 208, 7, 104, 2, 2, 208, 209, 7, 113, 2, 2,
	209, 210, 7, 116, 2, 2, 210, 66, 3, 2, 2, 2, 211, 212, 7, 107, 2, 2, 212,
	213, 7, 112, 2, 2, 213, 68, 3, 2, 2, 2, 214, 215, 7, 101, 2, 2, 215, 216,
	7, 110, 2, 2, 216, 217, 7, 99, 2, 2, 217, 218, 7, 117, 2, 2, 218, 219,
	7, 117, 2, 2, 219, 70, 3, 2, 2, 2, 220, 221, 7, 99, 2, 2, 221, 222, 7,
	118, 2, 2, 222, 223, 7, 118, 2, 2, 223, 224, 7, 116, 2, 2, 224, 225, 7,
	107, 2, 2, 225, 226, 7, 100, 2, 2, 226, 227, 7, 119, 2, 2, 227, 228, 7,
	118, 2, 2, 228, 229, 7, 103, 2, 2, 229, 230, 7, 117, 2, 2, 230, 72, 3,
	2, 2, 2, 231, 232, 7, 111, 2, 2, 232, 233, 7, 103, 2, 2, 233, 234, 7, 118,
	2, 2, 234, 235, 7, 106, 2, 2, 235, 236, 7, 113, 2, 2, 236, 237, 7, 102,
	2, 2, 237, 238, 7, 117, 2, 2, 238, 74, 3, 2, 2, 2, 239, 240, 7, 121, 2,
	2, 240, 241, 7, 116, 2, 2, 241, 242, 7, 107, 2, 2, 242, 243, 7, 118, 2,
	2, 243, 244, 7, 103, 2, 2, 244, 76, 3, 2, 2, 2, 245, 246, 7, 116, 2, 2,
	246, 247, 7, 103, 2, 2, 247, 248, 7, 99, 2, 2, 248, 249, 7, 102, 2, 2,
	249, 78, 3, 2, 2, 2, 250, 251, 7, 104, 2, 2, 251, 252, 7, 119, 2, 2, 252,
	253, 7, 112, 2, 2, 253, 254, 7, 101, 2, 2, 254, 255, 7, 118, 2, 2, 255,
	256, 7, 107, 2, 2, 256, 257, 7, 113, 2, 2, 257, 258, 7, 112, 2, 2, 258,
	80, 3, 2, 2, 2, 259, 260, 7, 116, 2, 2, 260, 261, 7, 103, 2, 2, 261, 262,
	7, 118, 2, 2, 262, 263, 7, 119, 2, 2, 263, 264, 7, 116, 2, 2, 264, 265,
	7, 112, 2, 2, 265, 82, 3, 2, 2, 2, 266, 267, 7, 111, 2, 2, 267, 268, 7,
	99, 2, 2, 268, 269, 7, 107, 2, 2, 269, 270, 7, 112, 2, 2, 270, 84, 3, 2,
	2, 2, 271, 272, 7, 120, 2, 2, 272, 273, 7, 99, 2, 2, 273, 274, 7, 116,
	2, 2, 274, 86, 3, 2, 2, 2, 275, 276, 7, 107, 2, 2, 276, 277, 7, 112, 2,
	2, 277, 278, 7, 118, 2, 2, 278, 88, 3, 2, 2, 2, 279, 280, 7, 104, 2, 2,
	280, 281, 7, 110, 2, 2, 281, 282, 7, 113, 2, 2, 282, 283, 7, 99, 2, 2,
	283, 284, 7, 118, 2, 2, 284, 90, 3, 2, 2, 2, 285, 286, 7, 101, 2, 2, 286,
	287, 7, 106, 2, 2, 287, 288, 7, 99, 2, 2, 288, 289, 7, 116, 2, 2, 289,
	92, 3, 2, 2, 2, 290, 291, 7, 117, 2, 2, 291, 292, 7, 118, 2, 2, 292, 293,
	7, 116, 2, 2, 293, 294, 7, 107, 2, 2, 294, 295, 7, 112, 2, 2, 295, 296,
	7, 105, 2, 2, 296, 94, 3, 2, 2, 2, 297, 298, 7, 100, 2, 2, 298, 299, 7,
	113, 2, 2, 299, 300, 7, 113, 2, 2, 300, 301, 7, 110, 2, 2, 301, 96, 3,
	2, 2, 2, 302, 303, 7, 120, 2, 2, 303, 304, 7, 113, 2, 2, 304, 305, 7, 107,
	2, 2, 305, 306, 7, 102, 2, 2, 306, 98, 3, 2, 2, 2, 307, 308, 7, 114, 2,
	2, 308, 309, 7, 116, 2, 2, 309, 310, 7, 113, 2, 2, 310, 311, 7, 105, 2,
	2, 311, 312, 7, 116, 2, 2, 312, 313, 7, 99, 2, 2, 313, 314, 7, 111, 2,
	2, 314, 100, 3, 2, 2, 2, 315, 319, 9, 6, 2, 2, 316, 318, 9, 7, 2, 2, 317,
	316, 3, 2, 2, 2, 318, 321, 3, 2, 2, 2, 319, 317, 3, 2, 2, 2, 319, 320,
	3, 2, 2, 2, 320, 102, 3, 2, 2, 2, 321, 319, 3, 2, 2, 2, 10, 2, 127, 132,
	136, 139, 154, 159, 319, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'{'", "'}'", "':'", "'['", "']'", "']['", "'='", "'.'", "','", "'\"'",
	"", "", "", "", "", "'||'", "'&&'", "'>'", "'<'", "'=='", "'!='", "'*'",
	"'/'", "'+'", "'-'", "';'", "'('", "')'", "'if'", "'else'", "'while'",
	"'for'", "'in'", "'class'", "'attributes'", "'methods'", "'write'", "'read'",
	"'function'", "'return'", "'main'", "'var'", "'int'", "'float'", "'char'",
	"'string'", "'bool'", "'void'", "'program'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "INT", "FLOAT", "CHAR", "BOOL",
	"WS", "OR", "AND", "GT", "LT", "EQ", "NEQ", "MUL", "DIV", "ADD", "SUB",
	"SEMI", "LPAREN", "RPAREN", "IF", "ELSE", "WHILE", "FOR", "IN", "CLASS",
	"ATTRIBUTES", "METHODS", "WRITE", "READ", "FUNCTION", "RETURN", "MAIN",
	"VAR", "INT_TYPE", "FLOAT_TYPE", "CHAR_TYPE", "STRING_TYPE", "BOOL_TYPE",
	"VOID", "PROGRAM", "ID",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "INT", "FLOAT", "CHAR", "BOOL", "WS", "OR", "AND", "GT", "LT",
	"EQ", "NEQ", "MUL", "DIV", "ADD", "SUB", "SEMI", "LPAREN", "RPAREN", "IF",
	"ELSE", "WHILE", "FOR", "IN", "CLASS", "ATTRIBUTES", "METHODS", "WRITE",
	"READ", "FUNCTION", "RETURN", "MAIN", "VAR", "INT_TYPE", "FLOAT_TYPE",
	"CHAR_TYPE", "STRING_TYPE", "BOOL_TYPE", "VOID", "PROGRAM", "ID",
}

type ProyectoLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewProyectoLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *ProyectoLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewProyectoLexer(input antlr.CharStream) *ProyectoLexer {
	l := new(ProyectoLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Proyecto.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ProyectoLexer tokens.
const (
	ProyectoLexerT__0        = 1
	ProyectoLexerT__1        = 2
	ProyectoLexerT__2        = 3
	ProyectoLexerT__3        = 4
	ProyectoLexerT__4        = 5
	ProyectoLexerT__5        = 6
	ProyectoLexerT__6        = 7
	ProyectoLexerT__7        = 8
	ProyectoLexerT__8        = 9
	ProyectoLexerT__9        = 10
	ProyectoLexerINT         = 11
	ProyectoLexerFLOAT       = 12
	ProyectoLexerCHAR        = 13
	ProyectoLexerBOOL        = 14
	ProyectoLexerWS          = 15
	ProyectoLexerOR          = 16
	ProyectoLexerAND         = 17
	ProyectoLexerGT          = 18
	ProyectoLexerLT          = 19
	ProyectoLexerEQ          = 20
	ProyectoLexerNEQ         = 21
	ProyectoLexerMUL         = 22
	ProyectoLexerDIV         = 23
	ProyectoLexerADD         = 24
	ProyectoLexerSUB         = 25
	ProyectoLexerSEMI        = 26
	ProyectoLexerLPAREN      = 27
	ProyectoLexerRPAREN      = 28
	ProyectoLexerIF          = 29
	ProyectoLexerELSE        = 30
	ProyectoLexerWHILE       = 31
	ProyectoLexerFOR         = 32
	ProyectoLexerIN          = 33
	ProyectoLexerCLASS       = 34
	ProyectoLexerATTRIBUTES  = 35
	ProyectoLexerMETHODS     = 36
	ProyectoLexerWRITE       = 37
	ProyectoLexerREAD        = 38
	ProyectoLexerFUNCTION    = 39
	ProyectoLexerRETURN      = 40
	ProyectoLexerMAIN        = 41
	ProyectoLexerVAR         = 42
	ProyectoLexerINT_TYPE    = 43
	ProyectoLexerFLOAT_TYPE  = 44
	ProyectoLexerCHAR_TYPE   = 45
	ProyectoLexerSTRING_TYPE = 46
	ProyectoLexerBOOL_TYPE   = 47
	ProyectoLexerVOID        = 48
	ProyectoLexerPROGRAM     = 49
	ProyectoLexerID          = 50
)
