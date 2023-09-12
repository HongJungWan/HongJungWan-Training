package com.test.training.next_step.level_1.pg_9;

/**
 * level 1 - 콜라 문제
 */
public class Pg_0902_1 {
    public static void main(String[] args) {
        System.out.println(solution(2, 1, 20));
        System.out.println(solution(3, 1, 20));
    }

    public static int solution(int a, int b, int n) {
        int totalCola = 0;
        int remainingBottles = n;

        while (remainingBottles >= a) {
            int newCola = remainingBottles / a * b;
            totalCola += newCola;
            remainingBottles = remainingBottles % a + newCola;
        }
        return totalCola;
    }
}
/**
 * 총 콜라 병 수
 * 남은 빈 병의 수
 * <p>
 * 1. 새로 얻을 수 있는 콜라의 병 수
 * 2. 총 콜라 병 수 업데이트
 * 3. 남은 빈 병의 수 업데이트
 */