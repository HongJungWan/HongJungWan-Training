package com.test.training.next_step.level_2.pg_7;

/**
 * level 2 - 다음 큰 숫자
 */
public class Pg_0731_1 {
    public static void main(String[] args) {
        System.out.println(solution(15));
    }

    public static int solution(int n) {
        int tempOne = 0;
        int tempTwo = 0;
        int answer = 0;
        String binaryString = "";

        binaryString = Integer.toBinaryString(n);
        for (char c : binaryString.toCharArray()) {
            if (c == '1') {
                tempOne++;
            }
        }

        while (true) {
            n++;
            tempTwo = 0;
            binaryString = Integer.toBinaryString(n);
            for (char c : binaryString.toCharArray()) {
                if (c == '1') {
                    tempTwo++;
                }
            }
            if (tempOne == tempTwo) {
                answer = n;
                break;
            }
        }
        return answer;
    }
}