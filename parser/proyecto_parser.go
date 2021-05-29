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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 51, 528,
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
	60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 3, 2, 3, 2,
	3, 2, 3, 2, 7, 2, 133, 10, 2, 12, 2, 14, 2, 136, 11, 2, 3, 2, 7, 2, 139,
	10, 2, 12, 2, 14, 2, 142, 11, 2, 3, 2, 3, 2, 3, 3, 7, 3, 147, 10, 3, 12,
	3, 14, 3, 150, 11, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5,
	4, 160, 10, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 169, 10,
	5, 12, 5, 14, 5, 172, 11, 5, 3, 5, 3, 5, 3, 6, 3, 6, 7, 6, 178, 10, 6,
	12, 6, 14, 6, 181, 11, 6, 3, 7, 3, 7, 3, 7, 5, 7, 186, 10, 7, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 11, 3, 11, 5, 11, 216, 10, 11, 3, 11, 5, 11, 219, 10,
	11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 5, 13, 227, 10, 13, 3, 14,
	3, 14, 3, 14, 5, 14, 232, 10, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 15, 5, 15, 241, 10, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 247,
	10, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 253, 10, 15, 3, 16, 3, 16, 3,
	16, 3, 16, 5, 16, 259, 10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 264, 10, 16,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 7, 17, 271, 10, 17, 12, 17, 14, 17,
	274, 11, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 7, 19, 282, 10,
	19, 12, 19, 14, 19, 285, 11, 19, 3, 19, 7, 19, 288, 10, 19, 12, 19, 14,
	19, 291, 11, 19, 3, 19, 3, 19, 3, 20, 3, 20, 5, 20, 297, 10, 20, 3, 20,
	3, 20, 3, 21, 3, 21, 7, 21, 303, 10, 21, 12, 21, 14, 21, 306, 11, 21, 3,
	21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22,
	318, 10, 22, 3, 23, 3, 23, 3, 23, 5, 23, 323, 10, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 24, 3, 24, 3, 24, 5, 24, 332, 10, 24, 3, 24, 3, 24, 3, 25,
	3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 5, 26, 342, 10, 26, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 27, 5, 27, 349, 10, 27, 3, 27, 3, 27, 3, 28, 3, 28, 5, 28,
	355, 10, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3,
	30, 5, 30, 366, 10, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 5, 31, 373,
	10, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33,
	3, 33, 5, 33, 385, 10, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3,
	36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38,
	3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3,
	41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 7, 43, 421, 10, 43, 12, 43,
	14, 43, 424, 11, 43, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 7, 45, 431, 10,
	45, 12, 45, 14, 45, 434, 11, 45, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 5,
	47, 441, 10, 47, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 7, 49, 448, 10, 49,
	12, 49, 14, 49, 451, 11, 49, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 7, 51,
	458, 10, 51, 12, 51, 14, 51, 461, 11, 51, 3, 52, 3, 52, 3, 52, 3, 53, 3,
	53, 3, 53, 3, 53, 3, 53, 5, 53, 471, 10, 53, 3, 53, 5, 53, 474, 10, 53,
	3, 54, 3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 5, 55, 485,
	10, 55, 3, 56, 3, 56, 3, 57, 3, 57, 3, 58, 3, 58, 3, 59, 3, 59, 3, 60,
	3, 60, 7, 60, 497, 10, 60, 12, 60, 14, 60, 500, 11, 60, 3, 60, 3, 60, 3,
	61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 62, 3, 62, 7, 62, 511, 10, 62, 12, 62,
	14, 62, 514, 11, 62, 3, 62, 7, 62, 517, 10, 62, 12, 62, 14, 62, 520, 11,
	62, 3, 62, 3, 62, 3, 63, 3, 63, 3, 64, 3, 64, 3, 64, 3, 498, 2, 65, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
	42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76,
	78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110,
	112, 114, 116, 118, 120, 122, 124, 126, 2, 6, 3, 2, 25, 26, 3, 2, 23, 24,
	3, 2, 45, 48, 3, 2, 19, 22, 2, 520, 2, 128, 3, 2, 2, 2, 4, 148, 3, 2, 2,
	2, 6, 154, 3, 2, 2, 2, 8, 164, 3, 2, 2, 2, 10, 175, 3, 2, 2, 2, 12, 185,
	3, 2, 2, 2, 14, 187, 3, 2, 2, 2, 16, 193, 3, 2, 2, 2, 18, 202, 3, 2, 2,
	2, 20, 215, 3, 2, 2, 2, 22, 220, 3, 2, 2, 2, 24, 223, 3, 2, 2, 2, 26, 228,
	3, 2, 2, 2, 28, 237, 3, 2, 2, 2, 30, 254, 3, 2, 2, 2, 32, 267, 3, 2, 2,
	2, 34, 275, 3, 2, 2, 2, 36, 279, 3, 2, 2, 2, 38, 294, 3, 2, 2, 2, 40, 300,
	3, 2, 2, 2, 42, 317, 3, 2, 2, 2, 44, 322, 3, 2, 2, 2, 46, 328, 3, 2, 2,
	2, 48, 335, 3, 2, 2, 2, 50, 341, 3, 2, 2, 2, 52, 343, 3, 2, 2, 2, 54, 354,
	3, 2, 2, 2, 56, 356, 3, 2, 2, 2, 58, 365, 3, 2, 2, 2, 60, 372, 3, 2, 2,
	2, 62, 374, 3, 2, 2, 2, 64, 380, 3, 2, 2, 2, 66, 386, 3, 2, 2, 2, 68, 390,
	3, 2, 2, 2, 70, 392, 3, 2, 2, 2, 72, 395, 3, 2, 2, 2, 74, 401, 3, 2, 2,
	2, 76, 405, 3, 2, 2, 2, 78, 407, 3, 2, 2, 2, 80, 411, 3, 2, 2, 2, 82, 415,
	3, 2, 2, 2, 84, 418, 3, 2, 2, 2, 86, 425, 3, 2, 2, 2, 88, 428, 3, 2, 2,
	2, 90, 435, 3, 2, 2, 2, 92, 438, 3, 2, 2, 2, 94, 442, 3, 2, 2, 2, 96, 445,
	3, 2, 2, 2, 98, 452, 3, 2, 2, 2, 100, 455, 3, 2, 2, 2, 102, 462, 3, 2,
	2, 2, 104, 473, 3, 2, 2, 2, 106, 475, 3, 2, 2, 2, 108, 484, 3, 2, 2, 2,
	110, 486, 3, 2, 2, 2, 112, 488, 3, 2, 2, 2, 114, 490, 3, 2, 2, 2, 116,
	492, 3, 2, 2, 2, 118, 494, 3, 2, 2, 2, 120, 503, 3, 2, 2, 2, 122, 508,
	3, 2, 2, 2, 124, 523, 3, 2, 2, 2, 126, 525, 3, 2, 2, 2, 128, 129, 7, 50,
	2, 2, 129, 130, 7, 51, 2, 2, 130, 134, 7, 27, 2, 2, 131, 133, 5, 6, 4,
	2, 132, 131, 3, 2, 2, 2, 133, 136, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 134,
	135, 3, 2, 2, 2, 135, 140, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 137, 139,
	5, 12, 7, 2, 138, 137, 3, 2, 2, 2, 139, 142, 3, 2, 2, 2, 140, 138, 3, 2,
	2, 2, 140, 141, 3, 2, 2, 2, 141, 143, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2,
	143, 144, 5, 4, 3, 2, 144, 3, 3, 2, 2, 2, 145, 147, 5, 30, 16, 2, 146,
	145, 3, 2, 2, 2, 147, 150, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 148, 149,
	3, 2, 2, 2, 149, 151, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 151, 152, 5, 120,
	61, 2, 152, 153, 7, 2, 2, 3, 153, 5, 3, 2, 2, 2, 154, 155, 7, 36, 2, 2,
	155, 159, 7, 51, 2, 2, 156, 157, 7, 20, 2, 2, 157, 158, 7, 51, 2, 2, 158,
	160, 7, 19, 2, 2, 159, 156, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 161,
	3, 2, 2, 2, 161, 162, 5, 8, 5, 2, 162, 163, 7, 27, 2, 2, 163, 7, 3, 2,
	2, 2, 164, 165, 7, 3, 2, 2, 165, 166, 5, 10, 6, 2, 166, 170, 7, 38, 2,
	2, 167, 169, 5, 30, 16, 2, 168, 167, 3, 2, 2, 2, 169, 172, 3, 2, 2, 2,
	170, 168, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 173, 3, 2, 2, 2, 172,
	170, 3, 2, 2, 2, 173, 174, 7, 4, 2, 2, 174, 9, 3, 2, 2, 2, 175, 179, 7,
	37, 2, 2, 176, 178, 5, 12, 7, 2, 177, 176, 3, 2, 2, 2, 178, 181, 3, 2,
	2, 2, 179, 177, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 11, 3, 2, 2, 2,
	181, 179, 3, 2, 2, 2, 182, 186, 5, 14, 8, 2, 183, 186, 5, 16, 9, 2, 184,
	186, 5, 18, 10, 2, 185, 182, 3, 2, 2, 2, 185, 183, 3, 2, 2, 2, 185, 184,
	3, 2, 2, 2, 186, 13, 3, 2, 2, 2, 187, 188, 7, 44, 2, 2, 188, 189, 7, 51,
	2, 2, 189, 190, 7, 5, 2, 2, 190, 191, 5, 20, 11, 2, 191, 192, 7, 27, 2,
	2, 192, 15, 3, 2, 2, 2, 193, 194, 7, 44, 2, 2, 194, 195, 7, 51, 2, 2, 195,
	196, 7, 5, 2, 2, 196, 197, 7, 6, 2, 2, 197, 198, 7, 12, 2, 2, 198, 199,
	7, 7, 2, 2, 199, 200, 5, 124, 63, 2, 200, 201, 7, 27, 2, 2, 201, 17, 3,
	2, 2, 2, 202, 203, 7, 44, 2, 2, 203, 204, 7, 51, 2, 2, 204, 205, 7, 5,
	2, 2, 205, 206, 7, 6, 2, 2, 206, 207, 7, 12, 2, 2, 207, 208, 7, 8, 2, 2,
	208, 209, 7, 12, 2, 2, 209, 210, 7, 7, 2, 2, 210, 211, 5, 124, 63, 2, 211,
	212, 7, 27, 2, 2, 212, 19, 3, 2, 2, 2, 213, 216, 5, 124, 63, 2, 214, 216,
	7, 51, 2, 2, 215, 213, 3, 2, 2, 2, 215, 214, 3, 2, 2, 2, 216, 218, 3, 2,
	2, 2, 217, 219, 5, 22, 12, 2, 218, 217, 3, 2, 2, 2, 218, 219, 3, 2, 2,
	2, 219, 21, 3, 2, 2, 2, 220, 221, 7, 28, 2, 2, 221, 222, 5, 84, 43, 2,
	222, 23, 3, 2, 2, 2, 223, 226, 7, 51, 2, 2, 224, 225, 7, 9, 2, 2, 225,
	227, 7, 51, 2, 2, 226, 224, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227, 25,
	3, 2, 2, 2, 228, 231, 7, 51, 2, 2, 229, 230, 7, 9, 2, 2, 230, 232, 7, 51,
	2, 2, 231, 229, 3, 2, 2, 2, 231, 232, 3, 2, 2, 2, 232, 233, 3, 2, 2, 2,
	233, 234, 7, 6, 2, 2, 234, 235, 5, 84, 43, 2, 235, 236, 7, 7, 2, 2, 236,
	27, 3, 2, 2, 2, 237, 240, 7, 51, 2, 2, 238, 239, 7, 9, 2, 2, 239, 241,
	7, 51, 2, 2, 240, 238, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 246, 3, 2,
	2, 2, 242, 243, 7, 6, 2, 2, 243, 244, 5, 84, 43, 2, 244, 245, 7, 7, 2,
	2, 245, 247, 3, 2, 2, 2, 246, 242, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247,
	252, 3, 2, 2, 2, 248, 249, 7, 6, 2, 2, 249, 250, 5, 84, 43, 2, 250, 251,
	7, 7, 2, 2, 251, 253, 3, 2, 2, 2, 252, 248, 3, 2, 2, 2, 252, 253, 3, 2,
	2, 2, 253, 29, 3, 2, 2, 2, 254, 255, 7, 41, 2, 2, 255, 256, 7, 51, 2, 2,
	256, 258, 7, 29, 2, 2, 257, 259, 5, 32, 17, 2, 258, 257, 3, 2, 2, 2, 258,
	259, 3, 2, 2, 2, 259, 260, 3, 2, 2, 2, 260, 263, 7, 30, 2, 2, 261, 264,
	5, 124, 63, 2, 262, 264, 7, 49, 2, 2, 263, 261, 3, 2, 2, 2, 263, 262, 3,
	2, 2, 2, 264, 265, 3, 2, 2, 2, 265, 266, 5, 36, 19, 2, 266, 31, 3, 2, 2,
	2, 267, 272, 5, 34, 18, 2, 268, 269, 7, 10, 2, 2, 269, 271, 5, 34, 18,
	2, 270, 268, 3, 2, 2, 2, 271, 274, 3, 2, 2, 2, 272, 270, 3, 2, 2, 2, 272,
	273, 3, 2, 2, 2, 273, 33, 3, 2, 2, 2, 274, 272, 3, 2, 2, 2, 275, 276, 7,
	51, 2, 2, 276, 277, 7, 5, 2, 2, 277, 278, 5, 124, 63, 2, 278, 35, 3, 2,
	2, 2, 279, 283, 7, 3, 2, 2, 280, 282, 5, 12, 7, 2, 281, 280, 3, 2, 2, 2,
	282, 285, 3, 2, 2, 2, 283, 281, 3, 2, 2, 2, 283, 284, 3, 2, 2, 2, 284,
	289, 3, 2, 2, 2, 285, 283, 3, 2, 2, 2, 286, 288, 5, 42, 22, 2, 287, 286,
	3, 2, 2, 2, 288, 291, 3, 2, 2, 2, 289, 287, 3, 2, 2, 2, 289, 290, 3, 2,
	2, 2, 290, 292, 3, 2, 2, 2, 291, 289, 3, 2, 2, 2, 292, 293, 7, 4, 2, 2,
	293, 37, 3, 2, 2, 2, 294, 296, 7, 42, 2, 2, 295, 297, 5, 84, 43, 2, 296,
	295, 3, 2, 2, 2, 296, 297, 3, 2, 2, 2, 297, 298, 3, 2, 2, 2, 298, 299,
	7, 27, 2, 2, 299, 39, 3, 2, 2, 2, 300, 304, 7, 3, 2, 2, 301, 303, 5, 42,
	22, 2, 302, 301, 3, 2, 2, 2, 303, 306, 3, 2, 2, 2, 304, 302, 3, 2, 2, 2,
	304, 305, 3, 2, 2, 2, 305, 307, 3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 307,
	308, 7, 4, 2, 2, 308, 41, 3, 2, 2, 2, 309, 318, 5, 44, 23, 2, 310, 318,
	5, 56, 29, 2, 311, 318, 5, 62, 32, 2, 312, 318, 5, 64, 33, 2, 313, 318,
	5, 72, 37, 2, 314, 318, 5, 78, 40, 2, 315, 318, 5, 82, 42, 2, 316, 318,
	5, 38, 20, 2, 317, 309, 3, 2, 2, 2, 317, 310, 3, 2, 2, 2, 317, 311, 3,
	2, 2, 2, 317, 312, 3, 2, 2, 2, 317, 313, 3, 2, 2, 2, 317, 314, 3, 2, 2,
	2, 317, 315, 3, 2, 2, 2, 317, 316, 3, 2, 2, 2, 318, 43, 3, 2, 2, 2, 319,
	323, 5, 24, 13, 2, 320, 323, 5, 26, 14, 2, 321, 323, 5, 28, 15, 2, 322,
	319, 3, 2, 2, 2, 322, 320, 3, 2, 2, 2, 322, 321, 3, 2, 2, 2, 323, 324,
	3, 2, 2, 2, 324, 325, 7, 28, 2, 2, 325, 326, 5, 84, 43, 2, 326, 327, 7,
	27, 2, 2, 327, 45, 3, 2, 2, 2, 328, 329, 7, 51, 2, 2, 329, 331, 7, 29,
	2, 2, 330, 332, 5, 48, 25, 2, 331, 330, 3, 2, 2, 2, 331, 332, 3, 2, 2,
	2, 332, 333, 3, 2, 2, 2, 333, 334, 7, 30, 2, 2, 334, 47, 3, 2, 2, 2, 335,
	336, 5, 84, 43, 2, 336, 337, 5, 50, 26, 2, 337, 49, 3, 2, 2, 2, 338, 339,
	7, 10, 2, 2, 339, 342, 5, 48, 25, 2, 340, 342, 3, 2, 2, 2, 341, 338, 3,
	2, 2, 2, 341, 340, 3, 2, 2, 2, 342, 51, 3, 2, 2, 2, 343, 344, 7, 51, 2,
	2, 344, 345, 7, 9, 2, 2, 345, 346, 7, 51, 2, 2, 346, 348, 7, 29, 2, 2,
	347, 349, 5, 48, 25, 2, 348, 347, 3, 2, 2, 2, 348, 349, 3, 2, 2, 2, 349,
	350, 3, 2, 2, 2, 350, 351, 7, 30, 2, 2, 351, 53, 3, 2, 2, 2, 352, 355,
	5, 46, 24, 2, 353, 355, 5, 52, 27, 2, 354, 352, 3, 2, 2, 2, 354, 353, 3,
	2, 2, 2, 355, 55, 3, 2, 2, 2, 356, 357, 7, 40, 2, 2, 357, 358, 7, 29, 2,
	2, 358, 359, 5, 58, 30, 2, 359, 360, 7, 30, 2, 2, 360, 361, 7, 27, 2, 2,
	361, 57, 3, 2, 2, 2, 362, 366, 5, 24, 13, 2, 363, 366, 5, 26, 14, 2, 364,
	366, 5, 28, 15, 2, 365, 362, 3, 2, 2, 2, 365, 363, 3, 2, 2, 2, 365, 364,
	3, 2, 2, 2, 366, 367, 3, 2, 2, 2, 367, 368, 5, 60, 31, 2, 368, 59, 3, 2,
	2, 2, 369, 370, 7, 10, 2, 2, 370, 373, 5, 58, 30, 2, 371, 373, 3, 2, 2,
	2, 372, 369, 3, 2, 2, 2, 372, 371, 3, 2, 2, 2, 373, 61, 3, 2, 2, 2, 374,
	375, 7, 39, 2, 2, 375, 376, 7, 29, 2, 2, 376, 377, 5, 48, 25, 2, 377, 378,
	7, 30, 2, 2, 378, 379, 7, 27, 2, 2, 379, 63, 3, 2, 2, 2, 380, 381, 7, 31,
	2, 2, 381, 382, 5, 66, 34, 2, 382, 384, 5, 68, 35, 2, 383, 385, 5, 70,
	36, 2, 384, 383, 3, 2, 2, 2, 384, 385, 3, 2, 2, 2, 385, 65, 3, 2, 2, 2,
	386, 387, 7, 29, 2, 2, 387, 388, 5, 84, 43, 2, 388, 389, 7, 30, 2, 2, 389,
	67, 3, 2, 2, 2, 390, 391, 5, 40, 21, 2, 391, 69, 3, 2, 2, 2, 392, 393,
	7, 32, 2, 2, 393, 394, 5, 40, 21, 2, 394, 71, 3, 2, 2, 2, 395, 396, 7,
	34, 2, 2, 396, 397, 5, 74, 38, 2, 397, 398, 7, 35, 2, 2, 398, 399, 5, 76,
	39, 2, 399, 400, 5, 40, 21, 2, 400, 73, 3, 2, 2, 2, 401, 402, 7, 51, 2,
	2, 402, 403, 7, 28, 2, 2, 403, 404, 5, 84, 43, 2, 404, 75, 3, 2, 2, 2,
	405, 406, 5, 84, 43, 2, 406, 77, 3, 2, 2, 2, 407, 408, 7, 33, 2, 2, 408,
	409, 5, 80, 41, 2, 409, 410, 5, 40, 21, 2, 410, 79, 3, 2, 2, 2, 411, 412,
	7, 29, 2, 2, 412, 413, 5, 84, 43, 2, 413, 414, 7, 30, 2, 2, 414, 81, 3,
	2, 2, 2, 415, 416, 5, 84, 43, 2, 416, 417, 7, 27, 2, 2, 417, 83, 3, 2,
	2, 2, 418, 422, 5, 88, 45, 2, 419, 421, 5, 86, 44, 2, 420, 419, 3, 2, 2,
	2, 421, 424, 3, 2, 2, 2, 422, 420, 3, 2, 2, 2, 422, 423, 3, 2, 2, 2, 423,
	85, 3, 2, 2, 2, 424, 422, 3, 2, 2, 2, 425, 426, 7, 17, 2, 2, 426, 427,
	5, 88, 45, 2, 427, 87, 3, 2, 2, 2, 428, 432, 5, 92, 47, 2, 429, 431, 5,
	90, 46, 2, 430, 429, 3, 2, 2, 2, 431, 434, 3, 2, 2, 2, 432, 430, 3, 2,
	2, 2, 432, 433, 3, 2, 2, 2, 433, 89, 3, 2, 2, 2, 434, 432, 3, 2, 2, 2,
	435, 436, 7, 18, 2, 2, 436, 437, 5, 92, 47, 2, 437, 91, 3, 2, 2, 2, 438,
	440, 5, 96, 49, 2, 439, 441, 5, 94, 48, 2, 440, 439, 3, 2, 2, 2, 440, 441,
	3, 2, 2, 2, 441, 93, 3, 2, 2, 2, 442, 443, 5, 126, 64, 2, 443, 444, 5,
	96, 49, 2, 444, 95, 3, 2, 2, 2, 445, 449, 5, 100, 51, 2, 446, 448, 5, 98,
	50, 2, 447, 446, 3, 2, 2, 2, 448, 451, 3, 2, 2, 2, 449, 447, 3, 2, 2, 2,
	449, 450, 3, 2, 2, 2, 450, 97, 3, 2, 2, 2, 451, 449, 3, 2, 2, 2, 452, 453,
	9, 2, 2, 2, 453, 454, 5, 100, 51, 2, 454, 99, 3, 2, 2, 2, 455, 459, 5,
	104, 53, 2, 456, 458, 5, 102, 52, 2, 457, 456, 3, 2, 2, 2, 458, 461, 3,
	2, 2, 2, 459, 457, 3, 2, 2, 2, 459, 460, 3, 2, 2, 2, 460, 101, 3, 2, 2,
	2, 461, 459, 3, 2, 2, 2, 462, 463, 9, 3, 2, 2, 463, 464, 5, 104, 53, 2,
	464, 103, 3, 2, 2, 2, 465, 474, 5, 106, 54, 2, 466, 474, 5, 108, 55, 2,
	467, 471, 5, 24, 13, 2, 468, 471, 5, 26, 14, 2, 469, 471, 5, 28, 15, 2,
	470, 467, 3, 2, 2, 2, 470, 468, 3, 2, 2, 2, 470, 469, 3, 2, 2, 2, 471,
	474, 3, 2, 2, 2, 472, 474, 5, 54, 28, 2, 473, 465, 3, 2, 2, 2, 473, 466,
	3, 2, 2, 2, 473, 470, 3, 2, 2, 2, 473, 472, 3, 2, 2, 2, 474, 105, 3, 2,
	2, 2, 475, 476, 7, 29, 2, 2, 476, 477, 5, 84, 43, 2, 477, 478, 7, 30, 2,
	2, 478, 107, 3, 2, 2, 2, 479, 485, 5, 110, 56, 2, 480, 485, 5, 112, 57,
	2, 481, 485, 5, 114, 58, 2, 482, 485, 5, 118, 60, 2, 483, 485, 5, 116,
	59, 2, 484, 479, 3, 2, 2, 2, 484, 480, 3, 2, 2, 2, 484, 481, 3, 2, 2, 2,
	484, 482, 3, 2, 2, 2, 484, 483, 3, 2, 2, 2, 485, 109, 3, 2, 2, 2, 486,
	487, 7, 12, 2, 2, 487, 111, 3, 2, 2, 2, 488, 489, 7, 13, 2, 2, 489, 113,
	3, 2, 2, 2, 490, 491, 7, 14, 2, 2, 491, 115, 3, 2, 2, 2, 492, 493, 7, 15,
	2, 2, 493, 117, 3, 2, 2, 2, 494, 498, 7, 11, 2, 2, 495, 497, 11, 2, 2,
	2, 496, 495, 3, 2, 2, 2, 497, 500, 3, 2, 2, 2, 498, 499, 3, 2, 2, 2, 498,
	496, 3, 2, 2, 2, 499, 501, 3, 2, 2, 2, 500, 498, 3, 2, 2, 2, 501, 502,
	7, 11, 2, 2, 502, 119, 3, 2, 2, 2, 503, 504, 7, 43, 2, 2, 504, 505, 7,
	29, 2, 2, 505, 506, 7, 30, 2, 2, 506, 507, 5, 122, 62, 2, 507, 121, 3,
	2, 2, 2, 508, 512, 7, 3, 2, 2, 509, 511, 5, 12, 7, 2, 510, 509, 3, 2, 2,
	2, 511, 514, 3, 2, 2, 2, 512, 510, 3, 2, 2, 2, 512, 513, 3, 2, 2, 2, 513,
	518, 3, 2, 2, 2, 514, 512, 3, 2, 2, 2, 515, 517, 5, 42, 22, 2, 516, 515,
	3, 2, 2, 2, 517, 520, 3, 2, 2, 2, 518, 516, 3, 2, 2, 2, 518, 519, 3, 2,
	2, 2, 519, 521, 3, 2, 2, 2, 520, 518, 3, 2, 2, 2, 521, 522, 7, 4, 2, 2,
	522, 123, 3, 2, 2, 2, 523, 524, 9, 4, 2, 2, 524, 125, 3, 2, 2, 2, 525,
	526, 9, 5, 2, 2, 526, 127, 3, 2, 2, 2, 43, 134, 140, 148, 159, 170, 179,
	185, 215, 218, 226, 231, 240, 246, 252, 258, 263, 272, 283, 289, 296, 304,
	317, 322, 331, 341, 348, 354, 365, 372, 384, 422, 432, 440, 449, 459, 470,
	473, 484, 498, 512, 518,
}
var literalNames = []string{
	"", "'{'", "'}'", "':'", "'['", "']'", "']['", "'.'", "','", "'\"'", "",
	"", "", "", "", "'||'", "'&&'", "'>'", "'<'", "'=='", "'!='", "'*'", "'/'",
	"'+'", "'-'", "';'", "'='", "'('", "')'", "'if'", "'else'", "'while'",
	"'for'", "'in'", "'class'", "'attributes'", "'methods'", "'write'", "'read'",
	"'function'", "'return'", "'main'", "'var'", "'int'", "'float'", "'char'",
	"'bool'", "'void'", "'program'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "INT", "FLOAT", "CHAR", "BOOL",
	"WS", "OR", "AND", "GT", "LT", "EQ", "NEQ", "MUL", "DIV", "ADD", "SUB",
	"SEMI", "ASSIGN", "LPAREN", "RPAREN", "IF", "ELSE", "WHILE", "FOR", "IN",
	"CLASS", "ATTRIBUTES", "METHODS", "WRITE", "READ", "FUNCTION", "RETURN",
	"MAIN", "VAR", "INT_TYPE", "FLOAT_TYPE", "CHAR_TYPE", "BOOL_TYPE", "VOID",
	"PROGRAM", "ID",
}

var ruleNames = []string{
	"program", "program2", "classDef", "classBlock", "classAttributes", "variableDeclaration",
	"varsDec", "varsDecArray", "varsDecMat", "varsTypeInit", "varsTypeInit2",
	"vars", "varArray", "varMat", "functions", "parameters", "parameter", "functionBlock",
	"returnRule", "block", "statutes", "assignation", "functionCall", "arguments",
	"arguments2", "methodCall", "call", "read", "read2", "read3", "write",
	"conditional", "conditional2", "conditional3", "conditional4", "forLoop",
	"forLoop2", "forLoop3", "whileLoop", "whileLoop2", "expression", "exp",
	"exp2", "t_exp", "t_exp2", "g_exp", "g_exp2", "m_exp", "m_exp2", "term",
	"term2", "factor", "factor2", "varCte", "cte_i", "cte_f", "cte_c", "cte_b",
	"cte_s", "main", "mainBlock", "typeRule", "relop",
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
	ProyectoParserINT        = 10
	ProyectoParserFLOAT      = 11
	ProyectoParserCHAR       = 12
	ProyectoParserBOOL       = 13
	ProyectoParserWS         = 14
	ProyectoParserOR         = 15
	ProyectoParserAND        = 16
	ProyectoParserGT         = 17
	ProyectoParserLT         = 18
	ProyectoParserEQ         = 19
	ProyectoParserNEQ        = 20
	ProyectoParserMUL        = 21
	ProyectoParserDIV        = 22
	ProyectoParserADD        = 23
	ProyectoParserSUB        = 24
	ProyectoParserSEMI       = 25
	ProyectoParserASSIGN     = 26
	ProyectoParserLPAREN     = 27
	ProyectoParserRPAREN     = 28
	ProyectoParserIF         = 29
	ProyectoParserELSE       = 30
	ProyectoParserWHILE      = 31
	ProyectoParserFOR        = 32
	ProyectoParserIN         = 33
	ProyectoParserCLASS      = 34
	ProyectoParserATTRIBUTES = 35
	ProyectoParserMETHODS    = 36
	ProyectoParserWRITE      = 37
	ProyectoParserREAD       = 38
	ProyectoParserFUNCTION   = 39
	ProyectoParserRETURN     = 40
	ProyectoParserMAIN       = 41
	ProyectoParserVAR        = 42
	ProyectoParserINT_TYPE   = 43
	ProyectoParserFLOAT_TYPE = 44
	ProyectoParserCHAR_TYPE  = 45
	ProyectoParserBOOL_TYPE  = 46
	ProyectoParserVOID       = 47
	ProyectoParserPROGRAM    = 48
	ProyectoParserID         = 49
)

// ProyectoParser rules.
const (
	ProyectoParserRULE_program             = 0
	ProyectoParserRULE_program2            = 1
	ProyectoParserRULE_classDef            = 2
	ProyectoParserRULE_classBlock          = 3
	ProyectoParserRULE_classAttributes     = 4
	ProyectoParserRULE_variableDeclaration = 5
	ProyectoParserRULE_varsDec             = 6
	ProyectoParserRULE_varsDecArray        = 7
	ProyectoParserRULE_varsDecMat          = 8
	ProyectoParserRULE_varsTypeInit        = 9
	ProyectoParserRULE_varsTypeInit2       = 10
	ProyectoParserRULE_vars                = 11
	ProyectoParserRULE_varArray            = 12
	ProyectoParserRULE_varMat              = 13
	ProyectoParserRULE_functions           = 14
	ProyectoParserRULE_parameters          = 15
	ProyectoParserRULE_parameter           = 16
	ProyectoParserRULE_functionBlock       = 17
	ProyectoParserRULE_returnRule          = 18
	ProyectoParserRULE_block               = 19
	ProyectoParserRULE_statutes            = 20
	ProyectoParserRULE_assignation         = 21
	ProyectoParserRULE_functionCall        = 22
	ProyectoParserRULE_arguments           = 23
	ProyectoParserRULE_arguments2          = 24
	ProyectoParserRULE_methodCall          = 25
	ProyectoParserRULE_call                = 26
	ProyectoParserRULE_read                = 27
	ProyectoParserRULE_read2               = 28
	ProyectoParserRULE_read3               = 29
	ProyectoParserRULE_write               = 30
	ProyectoParserRULE_conditional         = 31
	ProyectoParserRULE_conditional2        = 32
	ProyectoParserRULE_conditional3        = 33
	ProyectoParserRULE_conditional4        = 34
	ProyectoParserRULE_forLoop             = 35
	ProyectoParserRULE_forLoop2            = 36
	ProyectoParserRULE_forLoop3            = 37
	ProyectoParserRULE_whileLoop           = 38
	ProyectoParserRULE_whileLoop2          = 39
	ProyectoParserRULE_expression          = 40
	ProyectoParserRULE_exp                 = 41
	ProyectoParserRULE_exp2                = 42
	ProyectoParserRULE_t_exp               = 43
	ProyectoParserRULE_t_exp2              = 44
	ProyectoParserRULE_g_exp               = 45
	ProyectoParserRULE_g_exp2              = 46
	ProyectoParserRULE_m_exp               = 47
	ProyectoParserRULE_m_exp2              = 48
	ProyectoParserRULE_term                = 49
	ProyectoParserRULE_term2               = 50
	ProyectoParserRULE_factor              = 51
	ProyectoParserRULE_factor2             = 52
	ProyectoParserRULE_varCte              = 53
	ProyectoParserRULE_cte_i               = 54
	ProyectoParserRULE_cte_f               = 55
	ProyectoParserRULE_cte_c               = 56
	ProyectoParserRULE_cte_b               = 57
	ProyectoParserRULE_cte_s               = 58
	ProyectoParserRULE_main                = 59
	ProyectoParserRULE_mainBlock           = 60
	ProyectoParserRULE_typeRule            = 61
	ProyectoParserRULE_relop               = 62
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

func (s *ProgramContext) AllClassDef() []IClassDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassDefContext)(nil)).Elem())
	var tst = make([]IClassDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassDefContext)
		}
	}

	return tst
}

func (s *ProgramContext) ClassDef(i int) IClassDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassDefContext)
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
		p.SetState(126)
		p.Match(ProyectoParserPROGRAM)
	}
	{
		p.SetState(127)
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(128)
		p.Match(ProyectoParserSEMI)
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserCLASS {
		{
			p.SetState(129)
			p.ClassDef()
		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserVAR {
		{
			p.SetState(135)
			p.VariableDeclaration()
		}

		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(141)
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
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserFUNCTION {
		{
			p.SetState(143)
			p.Functions()
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(149)
		p.Main()
	}
	{
		p.SetState(150)
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
		p.SetState(152)
		p.Match(ProyectoParserCLASS)
	}
	{
		p.SetState(153)
		p.Match(ProyectoParserID)
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProyectoParserLT {
		{
			p.SetState(154)
			p.Match(ProyectoParserLT)
		}
		{
			p.SetState(155)
			p.Match(ProyectoParserID)
		}
		{
			p.SetState(156)
			p.Match(ProyectoParserGT)
		}

	}
	{
		p.SetState(159)
		p.ClassBlock()
	}
	{
		p.SetState(160)
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
		p.SetState(162)
		p.Match(ProyectoParserT__0)
	}
	{
		p.SetState(163)
		p.ClassAttributes()
	}
	{
		p.SetState(164)
		p.Match(ProyectoParserMETHODS)
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserFUNCTION {
		{
			p.SetState(165)
			p.Functions()
		}

		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(171)
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

func (s *ClassAttributesContext) AllVariableDeclaration() []IVariableDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem())
	var tst = make([]IVariableDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDeclarationContext)
		}
	}

	return tst
}

func (s *ClassAttributesContext) VariableDeclaration(i int) IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
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
		p.SetState(173)
		p.Match(ProyectoParserATTRIBUTES)
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserVAR {
		{
			p.SetState(174)
			p.VariableDeclaration()
		}

		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 10, ProyectoParserRULE_variableDeclaration)

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

	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.VarsDec()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.VarsDecArray()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(182)
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
	p.EnterRule(localctx, 12, ProyectoParserRULE_varsDec)

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
		p.Match(ProyectoParserVAR)
	}
	{
		p.SetState(186)
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(187)
		p.Match(ProyectoParserT__2)
	}
	{
		p.SetState(188)
		p.VarsTypeInit()
	}
	{
		p.SetState(189)
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
	p.EnterRule(localctx, 14, ProyectoParserRULE_varsDecArray)

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
		p.Match(ProyectoParserVAR)
	}
	{
		p.SetState(192)
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(193)
		p.Match(ProyectoParserT__2)
	}
	{
		p.SetState(194)
		p.Match(ProyectoParserT__3)
	}
	{
		p.SetState(195)
		p.Match(ProyectoParserINT)
	}
	{
		p.SetState(196)
		p.Match(ProyectoParserT__4)
	}
	{
		p.SetState(197)
		p.TypeRule()
	}
	{
		p.SetState(198)
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
	p.EnterRule(localctx, 16, ProyectoParserRULE_varsDecMat)

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
		p.Match(ProyectoParserVAR)
	}
	{
		p.SetState(201)
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(202)
		p.Match(ProyectoParserT__2)
	}
	{
		p.SetState(203)
		p.Match(ProyectoParserT__3)
	}
	{
		p.SetState(204)
		p.Match(ProyectoParserINT)
	}
	{
		p.SetState(205)
		p.Match(ProyectoParserT__5)
	}
	{
		p.SetState(206)
		p.Match(ProyectoParserINT)
	}
	{
		p.SetState(207)
		p.Match(ProyectoParserT__4)
	}
	{
		p.SetState(208)
		p.TypeRule()
	}
	{
		p.SetState(209)
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
	p.EnterRule(localctx, 18, ProyectoParserRULE_varsTypeInit)
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
	p.SetState(213)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProyectoParserINT_TYPE, ProyectoParserFLOAT_TYPE, ProyectoParserCHAR_TYPE, ProyectoParserBOOL_TYPE:
		{
			p.SetState(211)
			p.TypeRule()
		}

	case ProyectoParserID:
		{
			p.SetState(212)
			p.Match(ProyectoParserID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProyectoParserASSIGN {
		{
			p.SetState(215)
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
	p.EnterRule(localctx, 20, ProyectoParserRULE_varsTypeInit2)

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
		p.SetState(218)
		p.Match(ProyectoParserASSIGN)
	}
	{
		p.SetState(219)
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
	p.EnterRule(localctx, 22, ProyectoParserRULE_vars)
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
		p.SetState(221)
		p.Match(ProyectoParserID)
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProyectoParserT__6 {
		{
			p.SetState(222)
			p.Match(ProyectoParserT__6)
		}
		{
			p.SetState(223)
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
	p.EnterRule(localctx, 24, ProyectoParserRULE_varArray)
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
		p.SetState(226)
		p.Match(ProyectoParserID)
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProyectoParserT__6 {
		{
			p.SetState(227)
			p.Match(ProyectoParserT__6)
		}
		{
			p.SetState(228)
			p.Match(ProyectoParserID)
		}

	}
	{
		p.SetState(231)
		p.Match(ProyectoParserT__3)
	}
	{
		p.SetState(232)
		p.Exp()
	}
	{
		p.SetState(233)
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
	p.EnterRule(localctx, 26, ProyectoParserRULE_varMat)
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
		p.SetState(235)
		p.Match(ProyectoParserID)
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProyectoParserT__6 {
		{
			p.SetState(236)
			p.Match(ProyectoParserT__6)
		}
		{
			p.SetState(237)
			p.Match(ProyectoParserID)
		}

	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(240)
			p.Match(ProyectoParserT__3)
		}
		{
			p.SetState(241)
			p.Exp()
		}
		{
			p.SetState(242)
			p.Match(ProyectoParserT__4)
		}

	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProyectoParserT__3 {
		{
			p.SetState(246)
			p.Match(ProyectoParserT__3)
		}
		{
			p.SetState(247)
			p.Exp()
		}
		{
			p.SetState(248)
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
	p.EnterRule(localctx, 28, ProyectoParserRULE_functions)
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
		p.SetState(252)
		p.Match(ProyectoParserFUNCTION)
	}
	{
		p.SetState(253)
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(254)
		p.Match(ProyectoParserLPAREN)
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProyectoParserID {
		{
			p.SetState(255)
			p.Parameters()
		}

	}
	{
		p.SetState(258)
		p.Match(ProyectoParserRPAREN)
	}
	p.SetState(261)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProyectoParserINT_TYPE, ProyectoParserFLOAT_TYPE, ProyectoParserCHAR_TYPE, ProyectoParserBOOL_TYPE:
		{
			p.SetState(259)
			p.TypeRule()
		}

	case ProyectoParserVOID:
		{
			p.SetState(260)
			p.Match(ProyectoParserVOID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(263)
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
	p.EnterRule(localctx, 30, ProyectoParserRULE_parameters)
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
		p.SetState(265)
		p.Parameter()
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserT__7 {
		{
			p.SetState(266)
			p.Match(ProyectoParserT__7)
		}
		{
			p.SetState(267)
			p.Parameter()
		}

		p.SetState(272)
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
	p.EnterRule(localctx, 32, ProyectoParserRULE_parameter)

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
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(274)
		p.Match(ProyectoParserT__2)
	}
	{
		p.SetState(275)
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
	p.EnterRule(localctx, 34, ProyectoParserRULE_functionBlock)
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
		p.SetState(277)
		p.Match(ProyectoParserT__0)
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserVAR {
		{
			p.SetState(278)
			p.VariableDeclaration()
		}

		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProyectoParserT__8)|(1<<ProyectoParserINT)|(1<<ProyectoParserFLOAT)|(1<<ProyectoParserCHAR)|(1<<ProyectoParserBOOL)|(1<<ProyectoParserLPAREN)|(1<<ProyectoParserIF)|(1<<ProyectoParserWHILE))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ProyectoParserFOR-32))|(1<<(ProyectoParserWRITE-32))|(1<<(ProyectoParserREAD-32))|(1<<(ProyectoParserRETURN-32))|(1<<(ProyectoParserID-32)))) != 0) {
		{
			p.SetState(284)
			p.Statutes()
		}

		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(290)
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
	p.EnterRule(localctx, 36, ProyectoParserRULE_returnRule)
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
		p.SetState(292)
		p.Match(ProyectoParserRETURN)
	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProyectoParserT__8)|(1<<ProyectoParserINT)|(1<<ProyectoParserFLOAT)|(1<<ProyectoParserCHAR)|(1<<ProyectoParserBOOL)|(1<<ProyectoParserLPAREN))) != 0) || _la == ProyectoParserID {
		{
			p.SetState(293)
			p.Exp()
		}

	}
	{
		p.SetState(296)
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
	p.EnterRule(localctx, 38, ProyectoParserRULE_block)
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
		p.SetState(298)
		p.Match(ProyectoParserT__0)
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProyectoParserT__8)|(1<<ProyectoParserINT)|(1<<ProyectoParserFLOAT)|(1<<ProyectoParserCHAR)|(1<<ProyectoParserBOOL)|(1<<ProyectoParserLPAREN)|(1<<ProyectoParserIF)|(1<<ProyectoParserWHILE))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ProyectoParserFOR-32))|(1<<(ProyectoParserWRITE-32))|(1<<(ProyectoParserREAD-32))|(1<<(ProyectoParserRETURN-32))|(1<<(ProyectoParserID-32)))) != 0) {
		{
			p.SetState(299)
			p.Statutes()
		}

		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(305)
		p.Match(ProyectoParserT__1)
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
	p.EnterRule(localctx, 40, ProyectoParserRULE_statutes)

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

	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(307)
			p.Assignation()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(308)
			p.Read()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(309)
			p.Write()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(310)
			p.Conditional()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(311)
			p.ForLoop()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(312)
			p.WhileLoop()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(313)
			p.Expression()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(314)
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
	p.EnterRule(localctx, 42, ProyectoParserRULE_assignation)

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
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(317)
			p.Vars()
		}

	case 2:
		{
			p.SetState(318)
			p.VarArray()
		}

	case 3:
		{
			p.SetState(319)
			p.VarMat()
		}

	}
	{
		p.SetState(322)
		p.Match(ProyectoParserASSIGN)
	}
	{
		p.SetState(323)
		p.Exp()
	}
	{
		p.SetState(324)
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
	p.EnterRule(localctx, 44, ProyectoParserRULE_functionCall)
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
		p.SetState(326)
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(327)
		p.Match(ProyectoParserLPAREN)
	}
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProyectoParserT__8)|(1<<ProyectoParserINT)|(1<<ProyectoParserFLOAT)|(1<<ProyectoParserCHAR)|(1<<ProyectoParserBOOL)|(1<<ProyectoParserLPAREN))) != 0) || _la == ProyectoParserID {
		{
			p.SetState(328)
			p.Arguments()
		}

	}
	{
		p.SetState(331)
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
	p.EnterRule(localctx, 46, ProyectoParserRULE_arguments)

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
		p.SetState(333)
		p.Exp()
	}
	{
		p.SetState(334)
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
	p.EnterRule(localctx, 48, ProyectoParserRULE_arguments2)

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

	p.SetState(339)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProyectoParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(336)
			p.Match(ProyectoParserT__7)
		}
		{
			p.SetState(337)
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
	p.EnterRule(localctx, 50, ProyectoParserRULE_methodCall)
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
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(342)
		p.Match(ProyectoParserT__6)
	}
	{
		p.SetState(343)
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(344)
		p.Match(ProyectoParserLPAREN)
	}
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProyectoParserT__8)|(1<<ProyectoParserINT)|(1<<ProyectoParserFLOAT)|(1<<ProyectoParserCHAR)|(1<<ProyectoParserBOOL)|(1<<ProyectoParserLPAREN))) != 0) || _la == ProyectoParserID {
		{
			p.SetState(345)
			p.Arguments()
		}

	}
	{
		p.SetState(348)
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
	p.EnterRule(localctx, 52, ProyectoParserRULE_call)

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

	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(350)
			p.FunctionCall()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(351)
			p.MethodCall()
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
	p.EnterRule(localctx, 54, ProyectoParserRULE_read)

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
		p.SetState(354)
		p.Match(ProyectoParserREAD)
	}
	{
		p.SetState(355)
		p.Match(ProyectoParserLPAREN)
	}
	{
		p.SetState(356)
		p.Read2()
	}
	{
		p.SetState(357)
		p.Match(ProyectoParserRPAREN)
	}
	{
		p.SetState(358)
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
	p.EnterRule(localctx, 56, ProyectoParserRULE_read2)

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
	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(360)
			p.Vars()
		}

	case 2:
		{
			p.SetState(361)
			p.VarArray()
		}

	case 3:
		{
			p.SetState(362)
			p.VarMat()
		}

	}
	{
		p.SetState(365)
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
	p.EnterRule(localctx, 58, ProyectoParserRULE_read3)

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

	p.SetState(370)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProyectoParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(367)
			p.Match(ProyectoParserT__7)
		}
		{
			p.SetState(368)
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
	p.EnterRule(localctx, 60, ProyectoParserRULE_write)

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
		p.SetState(372)
		p.Match(ProyectoParserWRITE)
	}
	{
		p.SetState(373)
		p.Match(ProyectoParserLPAREN)
	}
	{
		p.SetState(374)
		p.Arguments()
	}
	{
		p.SetState(375)
		p.Match(ProyectoParserRPAREN)
	}
	{
		p.SetState(376)
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
	p.EnterRule(localctx, 62, ProyectoParserRULE_conditional)
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
		p.Match(ProyectoParserIF)
	}
	{
		p.SetState(379)
		p.Conditional2()
	}
	{
		p.SetState(380)
		p.Conditional3()
	}
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProyectoParserELSE {
		{
			p.SetState(381)
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
	p.EnterRule(localctx, 64, ProyectoParserRULE_conditional2)

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
		p.SetState(384)
		p.Match(ProyectoParserLPAREN)
	}
	{
		p.SetState(385)
		p.Exp()
	}
	{
		p.SetState(386)
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
	p.EnterRule(localctx, 66, ProyectoParserRULE_conditional3)

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
		p.SetState(388)
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
	p.EnterRule(localctx, 68, ProyectoParserRULE_conditional4)

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
		p.Match(ProyectoParserELSE)
	}
	{
		p.SetState(391)
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
	p.EnterRule(localctx, 70, ProyectoParserRULE_forLoop)

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
		p.Match(ProyectoParserFOR)
	}
	{
		p.SetState(394)
		p.ForLoop2()
	}
	{
		p.SetState(395)
		p.Match(ProyectoParserIN)
	}
	{
		p.SetState(396)
		p.ForLoop3()
	}
	{
		p.SetState(397)
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
	p.EnterRule(localctx, 72, ProyectoParserRULE_forLoop2)

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
		p.SetState(399)
		p.Match(ProyectoParserID)
	}
	{
		p.SetState(400)
		p.Match(ProyectoParserASSIGN)
	}
	{
		p.SetState(401)
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
	p.EnterRule(localctx, 74, ProyectoParserRULE_forLoop3)

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
		p.SetState(403)
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
	p.EnterRule(localctx, 76, ProyectoParserRULE_whileLoop)

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
		p.Match(ProyectoParserWHILE)
	}
	{
		p.SetState(406)
		p.WhileLoop2()
	}
	{
		p.SetState(407)
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
	p.EnterRule(localctx, 78, ProyectoParserRULE_whileLoop2)

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
		p.SetState(409)
		p.Match(ProyectoParserLPAREN)
	}
	{
		p.SetState(410)
		p.Exp()
	}
	{
		p.SetState(411)
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
	p.EnterRule(localctx, 80, ProyectoParserRULE_expression)

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
		p.SetState(413)
		p.Exp()
	}
	{
		p.SetState(414)
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
	p.EnterRule(localctx, 82, ProyectoParserRULE_exp)
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
		p.SetState(416)
		p.T_exp()
	}
	p.SetState(420)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserOR {
		{
			p.SetState(417)
			p.Exp2()
		}

		p.SetState(422)
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
	p.EnterRule(localctx, 84, ProyectoParserRULE_exp2)

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
		p.SetState(423)
		p.Match(ProyectoParserOR)
	}
	{
		p.SetState(424)
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
	p.EnterRule(localctx, 86, ProyectoParserRULE_t_exp)
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
		p.SetState(426)
		p.G_exp()
	}
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserAND {
		{
			p.SetState(427)
			p.T_exp2()
		}

		p.SetState(432)
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
	p.EnterRule(localctx, 88, ProyectoParserRULE_t_exp2)

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
		p.SetState(433)
		p.Match(ProyectoParserAND)
	}
	{
		p.SetState(434)
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
	p.EnterRule(localctx, 90, ProyectoParserRULE_g_exp)
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
		p.M_exp()
	}
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProyectoParserGT)|(1<<ProyectoParserLT)|(1<<ProyectoParserEQ)|(1<<ProyectoParserNEQ))) != 0 {
		{
			p.SetState(437)
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
	p.EnterRule(localctx, 92, ProyectoParserRULE_g_exp2)

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
		p.SetState(440)
		p.Relop()
	}
	{
		p.SetState(441)
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
	p.EnterRule(localctx, 94, ProyectoParserRULE_m_exp)
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
		p.SetState(443)
		p.Term()
	}
	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserADD || _la == ProyectoParserSUB {
		{
			p.SetState(444)
			p.M_exp2()
		}

		p.SetState(449)
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
	p.EnterRule(localctx, 96, ProyectoParserRULE_m_exp2)
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
		p.SetState(450)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ProyectoParserADD || _la == ProyectoParserSUB) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(451)
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
	p.EnterRule(localctx, 98, ProyectoParserRULE_term)
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
		p.SetState(453)
		p.Factor()
	}
	p.SetState(457)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserMUL || _la == ProyectoParserDIV {
		{
			p.SetState(454)
			p.Term2()
		}

		p.SetState(459)
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
	p.EnterRule(localctx, 100, ProyectoParserRULE_term2)
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
		p.SetState(460)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ProyectoParserMUL || _la == ProyectoParserDIV) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(461)
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
	p.EnterRule(localctx, 102, ProyectoParserRULE_factor)

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

	p.SetState(471)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(463)
			p.Factor2()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(464)
			p.VarCte()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(468)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(465)
				p.Vars()
			}

		case 2:
			{
				p.SetState(466)
				p.VarArray()
			}

		case 3:
			{
				p.SetState(467)
				p.VarMat()
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(470)
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
	p.EnterRule(localctx, 104, ProyectoParserRULE_factor2)

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
		p.SetState(473)
		p.Match(ProyectoParserLPAREN)
	}
	{
		p.SetState(474)
		p.Exp()
	}
	{
		p.SetState(475)
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
	p.EnterRule(localctx, 106, ProyectoParserRULE_varCte)

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

	p.SetState(482)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProyectoParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(477)
			p.Cte_i()
		}

	case ProyectoParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(478)
			p.Cte_f()
		}

	case ProyectoParserCHAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(479)
			p.Cte_c()
		}

	case ProyectoParserT__8:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(480)
			p.Cte_s()
		}

	case ProyectoParserBOOL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(481)
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
	p.EnterRule(localctx, 108, ProyectoParserRULE_cte_i)

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
	p.EnterRule(localctx, 110, ProyectoParserRULE_cte_f)

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
		p.SetState(486)
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
	p.EnterRule(localctx, 112, ProyectoParserRULE_cte_c)

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
		p.SetState(488)
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
	p.EnterRule(localctx, 114, ProyectoParserRULE_cte_b)

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
		p.SetState(490)
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
	p.EnterRule(localctx, 116, ProyectoParserRULE_cte_s)

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
		p.SetState(492)
		p.Match(ProyectoParserT__8)
	}
	p.SetState(496)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(493)
			p.MatchWildcard()

		}
		p.SetState(498)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())
	}
	{
		p.SetState(499)
		p.Match(ProyectoParserT__8)
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
	p.EnterRule(localctx, 118, ProyectoParserRULE_main)

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
		p.Match(ProyectoParserMAIN)
	}
	{
		p.SetState(502)
		p.Match(ProyectoParserLPAREN)
	}
	{
		p.SetState(503)
		p.Match(ProyectoParserRPAREN)
	}
	{
		p.SetState(504)
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
	p.EnterRule(localctx, 120, ProyectoParserRULE_mainBlock)
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
		p.SetState(506)
		p.Match(ProyectoParserT__0)
	}
	p.SetState(510)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ProyectoParserVAR {
		{
			p.SetState(507)
			p.VariableDeclaration()
		}

		p.SetState(512)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(516)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProyectoParserT__8)|(1<<ProyectoParserINT)|(1<<ProyectoParserFLOAT)|(1<<ProyectoParserCHAR)|(1<<ProyectoParserBOOL)|(1<<ProyectoParserLPAREN)|(1<<ProyectoParserIF)|(1<<ProyectoParserWHILE))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ProyectoParserFOR-32))|(1<<(ProyectoParserWRITE-32))|(1<<(ProyectoParserREAD-32))|(1<<(ProyectoParserRETURN-32))|(1<<(ProyectoParserID-32)))) != 0) {
		{
			p.SetState(513)
			p.Statutes()
		}

		p.SetState(518)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(519)
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
	p.EnterRule(localctx, 122, ProyectoParserRULE_typeRule)
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
		_la = p.GetTokenStream().LA(1)

		if !(((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(ProyectoParserINT_TYPE-43))|(1<<(ProyectoParserFLOAT_TYPE-43))|(1<<(ProyectoParserCHAR_TYPE-43))|(1<<(ProyectoParserBOOL_TYPE-43)))) != 0) {
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
	p.EnterRule(localctx, 124, ProyectoParserRULE_relop)
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
		p.SetState(523)
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
