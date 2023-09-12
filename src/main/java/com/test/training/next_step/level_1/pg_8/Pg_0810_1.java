package com.test.training.next_step.level_1.pg_8;

/**
 * level 1 - [1차] 비밀지도
 */
public class Pg_0810_1 {
    public static void main(String[] args) {
        int[] arr1 = {9, 20, 28, 18, 11};
        int[] arr2 = {30, 1, 21, 17, 28};
        System.out.println(solution(5, arr1, arr2));
    }

    public static String[] solution(int n, int[] arr1, int[] arr2) {
        int[] temp = new int[n];
        String[] answer = new String[n];

        for (int i = 0; i < n; i++) {
            temp[i] = arr1[i] | arr2[i];
        }

        for (int i = 0; i < n; i++) {
            answer[i] = Integer.toBinaryString(temp[i]);

            while (answer[i].length() < n) {
                answer[i] = "0" + answer[i];
            }

            answer[i] = answer[i].replaceAll("1", "#");
            answer[i] = answer[i].replaceAll("0", " ");
        }
        return answer;
    }
}