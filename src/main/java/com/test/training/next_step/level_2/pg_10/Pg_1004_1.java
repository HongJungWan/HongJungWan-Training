package com.test.training.next_step.level_2.pg_10;

/**
 * level 2 - 주식가격
 */
public class Pg_1004_1 {
    public static void main(String[] args) {
        for (int answer : solution(new int[]{1, 2, 3, 2, 3})) {
            System.out.print(answer + " ");
        }
    }

    public static int[] solution(int[] prices) {
        int[] answer = new int[prices.length];

        for (int i = 0; i < prices.length; i++) {
            for (int j = i + 1; j < prices.length; j++) {
                answer[i]++;
                if (prices[i] > prices[j]) break;
            }
        }
        return answer;
    }
}