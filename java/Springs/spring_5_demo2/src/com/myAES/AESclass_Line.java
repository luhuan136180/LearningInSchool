package com.myAES;

public class AESclass_Line {
    static int Nr = 10;//轮数
    static int key[] = {0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88,
            0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff};//16进制密钥

    static int text[] = {0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
            0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f};//16进制明文

    static int sbox[] = {0x63, 0x7C, 0x77, 0x7B, 0xF2, 0x6B, 0x6F, 0xC5, 0x30,
            0x01, 0x67, 0x2B, 0xFE, 0xD7, 0xAB, 0x76, 0xCA, 0x82, 0xC9, 0x7D,
            0xFA, 0x59, 0x47, 0xF0, 0xAD, 0xD4, 0xA2, 0xAF, 0x9C, 0xA4, 0x72,
            0xC0, 0xB7, 0xFD, 0x93, 0x26, 0x36, 0x3F, 0xF7, 0xCC, 0x34, 0xA5,
            0xE5, 0xF1, 0x71, 0xD8, 0x31, 0x15, 0x04, 0xC7, 0x23, 0xC3, 0x18,
            0x96, 0x05, 0x9A, 0x07, 0x12, 0x80, 0xE2, 0xEB, 0x27, 0xB2, 0x75,
            0x09, 0x83, 0x2C, 0x1A, 0x1B, 0x6E, 0x5A, 0xA0, 0x52, 0x3B, 0xD6,
            0xB3, 0x29, 0xE3, 0x2F, 0x84, 0x53, 0xD1, 0x00, 0xED, 0x20, 0xFC,
            0xB1, 0x5B, 0x6A, 0xCB, 0xBE, 0x39, 0x4A, 0x4C, 0x58, 0xCF, 0xD0,
            0xEF, 0xAA, 0xFB, 0x43, 0x4D, 0x33, 0x85, 0x45, 0xF9, 0x02, 0x7F,
            0x50, 0x3C, 0x9F, 0xA8, 0x51, 0xA3, 0x40, 0x8F, 0x92, 0x9D, 0x38,
            0xF5, 0xBC, 0xB6, 0xDA, 0x21, 0x10, 0xFF, 0xF3, 0xD2, 0xCD, 0x0C,
            0x13, 0xEC, 0x5F, 0x97, 0x44, 0x17, 0xC4, 0xA7, 0x7E, 0x3D, 0x64,
            0x5D, 0x19, 0x73, 0x60, 0x81, 0x4F, 0xDC, 0x22, 0x2A, 0x90, 0x88,
            0x46, 0xEE, 0xB8, 0x14, 0xDE, 0x5E, 0x0B, 0xDB, 0xE0, 0x32, 0x3A,
            0x0A, 0x49, 0x06, 0x24, 0x5C, 0xC2, 0xD3, 0xAC, 0x62, 0x91, 0x95,
            0xE4, 0x79, 0xE7, 0xC8, 0x37, 0x6D, 0x8D, 0xD5, 0x4E, 0xA9, 0x6C,
            0x56, 0xF4, 0xEA, 0x65, 0x7A, 0xAE, 0x08, 0xBA, 0x78, 0x25, 0x2E,
            0x1C, 0xA6, 0xB4, 0xC6, 0xE8, 0xDD, 0x74, 0x1F, 0x4B, 0xBD, 0x8B,
            0x8A, 0x70, 0x3E, 0xB5, 0x66, 0x48, 0x03, 0xF6, 0x0E, 0x61, 0x35,
            0x57, 0xB9, 0x86, 0xC1, 0x1D, 0x9E, 0xE1, 0xF8, 0x98, 0x11, 0x69,
            0xD9, 0x8E, 0x94, 0x9B, 0x1E, 0x87, 0xE9, 0xCE, 0x55, 0x28, 0xDF,
            0x8C, 0xA1, 0x89, 0x0D, 0xBF, 0xE6, 0x42, 0x68, 0x41, 0x99, 0x2D,
            0x0F, 0xB0, 0x54, 0xBB, 0x16,};//S盒表


    static int RCon[];//密钥的常量数组
    static int word[][];//轮密钥扩展后的数组，大小为 word[44][4]

    public static void main(String[] args) {
        keyExpansion();//密钥
        int i, r = 0;
        text = addRoundKey_28(text, 0);//初始轮密钥加
        r++;
//        System.out.println("");
//        for (i = 0; i < 16; i++) {
//            System.out.print(Integer.toHexString(text[i]) + " ");
//            if (i % 4 == 3) {
//                System.out.println();
//            }
//        }


        for (; r <= Nr; r++) { // 迭代10轮
            for (i = 0; i < 16; i++) {
                text[i] = subByte(text[i]);// 字节替代
            }

            text = shiftRows(text); // 行移位

            if (r != 10) {// 注意：AES加密最后一轮没有列混合
                text = mixColumn(text);// 列混合

            }


            text = addRoundKey_28(text, r);// 轮密钥加


        }

        //输出加密结果
        for (i = 0; i < 16; i++) {
            System.out.print(Integer.toHexString(text[i]) + " ");
            if (i % 4 == 3) {
                System.out.println();
            }
        }


    }

    /**
     * 密钥数组扩展
     */
    static void keyExpansion() {
        RCon = new int[10];//轮常量为固定值
        RCon[0] = 0x01;
        RCon[1] = 0x02;
        RCon[2] = 0x04;
        RCon[3] = 0x08;
        RCon[4] = 0x10;
        RCon[5] = 0x20;
        RCon[6] = 0x40;
        RCon[7] = 0x80;
        RCon[8] = 0x1B;
        RCon[9] = 0x36;
        word = new int[44][4];//44组轮密钥
        int i, j;
        int temp[];

        for (i = 0; i < 4; i++) {
            for (j = 0; j < 4; j++) {
                word[i][j] = key[i * 4 + j]; //把初始密钥放入数组
            }
        }

        /*通过密钥计算规则计算余下数组
         *
         *1.如果i不是4的倍数，那么第i列由如下等式确定：
         *W[i]=W[i-4]⨁W[i-1]
         *2.如果i是4的倍数，那么第i列由如下等式确定：
         *W[i]=W[i-4]⨁T(W[i-1])
         *其中，T是一个有点复杂的函数。函数T由3部分组成：字循环(每次循环一位)、字节代换（s盒）和轮常量异或。
         */

        for (i = 4; i < 44; i++) {
            temp = word[i - 1];//看作W[i-1]
            if (i % 4 == 0) { //i为4的倍数 进入函数运算 W[i-1]=T(W[i-1])
                temp = subWord(rotWord(temp));
                temp[0] = temp[0] ^ RCon[i / 4 - 1];
            }
            for (j = 0; j < 4; j++) {
                word[i][j] = word[i - 4][j] ^ temp[j];//相当于W[i]=W[i-4]⨁W[i-1]
            }
        }

    }

    //密钥扩展中的移位
    static int[] rotWord(int[] word) {
        int[] rot = new int[4];
        int i;
        for (i = 0; i < 4; i++) {
            rot[i] = word[(i + 1) % 4];
        }
        return rot;
    }

    //密钥扩展中的4个字节的代换（4个字节为一组） 例如 A2 BE C4 D5
    static int[] subWord(int[] word) {
        int sub[] = new int[4];
        int i;
        for (i = 0; i < 4; i++) {
            sub[i] = subByte(word[i]);
        }
        return sub;
    }

    //S盒的单个字节代换 例如 AE
    static int subByte(int w) {
        int x = w / 16;
        int y = w % 16;
        return sbox[x * 16 + y];
    }

    //轮密钥加
    static int[] addRoundKey_28(int[] text, int round) {
        int[] add = new int[16];
        int i, j;
        for (i = 0; i < 4; i++) {
            for (j = 0; j < 4; j++) {
                //System.out.print(Integer.toHexString(word[4 * round + i][j]) + ",");//第round轮的轮密钥
                add[4 * i + j] = text[4 * i + j] ^ word[4 * round + i][j]; //逐比特异或
            }


        }
        return add;
    }

    //行移位（循环移位）规则：第0行移0位  .....  第3行移3位
    static int[] shiftRows(int[] text) {
        int[] shift = new int[16];
        int i, j;
        for (i = 0; i < 4; i++) {
            for (j = 0; j < 4; j++) {
                shift[4 * i + j] = text[4 * ((i + j) % 4) + j];
            }
        }
        return shift;
    }

    //列混合
    static int[] mixColumn(int[] text) {
        int[] mix = new int[16];
        int[] mass = {2, 3, 1, 1};
        int i, j, u;
        for (i = 0; i < 16; i++) {
            u = 0;
            for (j = 0; j < 4; j++) {
                u = u ^ fieldMulit(text[(i / 4) * 4 + j], mass[(4 - i % 4 + j) % 4]);
            }
            mix[i] = u;
        }
        return mix;
    }


    static int fieldMulit(int x, int y) {
        String xString = Integer.toBinaryString(x);
        int i, j, mul = 0, tem = y;
        for (i = 0; i < xString.length(); i++) {
            if (xString.charAt(i) == '1') {
                for (j = 1; j < xString.length() - i; j++) {
                    tem = tem << 1;
                    if (tem > 255) {
                        tem = tem % 256;
                        tem = tem ^ 0x1b;
                    }
                }
                mul = mul ^ tem;
                tem = y;
            }
        }
        return mul;
    }


}

