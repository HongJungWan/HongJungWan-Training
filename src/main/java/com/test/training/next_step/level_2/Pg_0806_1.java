package com.test.training.next_step.level_2;

/**
 * level 2 - 카펫
 */
public class Pg_0806_1 {
    public static void main(String[] args) {
        int[] temp = solution(10, 2);
        for (int print : temp) {
            System.out.println(print);
        }
    }

    public static int[] solution(int brown, int yellow) {
        int[] answer = new int[2];

        for (int height = 3; height < brown + yellow; height++) {
            int width = (brown + yellow) / height;
            if (width >= height) {
                if ((height - 2) * (width - 2) == yellow) {
                    answer[0] = width;
                    answer[1] = height;
                    break;
                }
            }
        }
        return answer;
    }
}