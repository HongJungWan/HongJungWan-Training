package com.test.training.next_step.level_1.pg_9;

/**
 * level 1 - 추억 점수
 */
public class Pg_0904_1 {
    public static void main(String[] args) {
        String[][] photoGroups = {
                {"may", "kein", "kain", "radi"},
                {"may", "kein", "brin", "deny"},
                {"kon", "kain", "may", "coni"}
        };
        int[] result = solution(
                new String[]{"may", "kein", "kain", "radi"},
                new int[]{5, 10, 1, 3},
                photoGroups
        );

        for (int print : result) {
            System.out.println(print);
        }
    }

    public static int[] solution(String[] name, int[] yearning, String[][] photo) {
        int[] answer = new int[photo.length];

        for (int i = 0; i < photo.length; i++) {
            int temp = 0;
            for (int j = 0; j < photo[i].length; j++) {
                for (int k = 0; k < name.length; k++) {
                    if (photo[i][j].equals(name[k])) {
                        temp += yearning[k];
                    }
                }
                answer[i] = temp;
            }
        }
        return answer;
    }
}