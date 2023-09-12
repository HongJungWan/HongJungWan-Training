package com.test.training.next_step.level_2.pg_8;

/**
 * level 2 - N개의 최소공배수
 */
public class Pg_0819_1 {
    public static void main(String[] args) {
        System.out.println(solution(new int[]{1, 2, 3}));
        System.out.println(solution(new int[]{2, 6, 8, 14}));
        System.out.println(solution(new int[]{8, 14}));
    }

    public static int solution(int[] arr) {
        int lcm = arr[0];
        for (int i = 1; i < arr.length; i++) {
            lcm = getLCM(lcm, arr[i]);
        }
        return lcm;
    }

    // 최소 공배수(LCM: Least Common Multiple) 계산 함수
    private static int getLCM(int a, int b) {
        return a * b / getGCD(a, b);
    }

    // 최대 공약수(GCD: Greatest Common Diviso) 계산 함수
    private static int getGCD(int a, int b) {
        while (b != 0) {
            int temp = a % b;
            a = b;
            b = temp;
        }
        return a;
    }
}

// 최소 공배수(LCM) : 두 수의 곱 / 두 수의 최대 공약수

// 최대 공약수(GCD) : 유클리드 호제법
// 2개의 자연수(또는 정식) a, b에 대해서 a % b를 r이라 하면 (단, a>b),
// a와 b의 최대 공약수는 b와 r의 최대 공약수와 같다.
