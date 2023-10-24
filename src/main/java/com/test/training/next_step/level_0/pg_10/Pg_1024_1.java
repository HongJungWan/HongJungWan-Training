package com.test.training.next_step.level_0.pg_10;

public class Pg_1024_1 {
    public static void main(String[] args) {
        for (int result : solution(new int[]{2, 1, 6})) {
            System.out.print(result + " ");
        }

        System.out.println();
        System.out.println();

        for (int result : solution(new int[]{5, 2, 1, 7, 5})) {
            System.out.print(result + " ");
        }
    }

    public static int[] solution(int[] num_list) {
        int size = num_list.length;
        int[] answer = new int[size + 1];

        for (int i = 0; i < size; i++) {
            answer[i] = num_list[i];
        }

        if (num_list[size - 1] > num_list[size - 2]) {
            answer[size] = num_list[size - 1] - num_list[size - 2];
        } else {
            answer[size] = num_list[size - 1] * 2;
        }
        return answer;
    }
}