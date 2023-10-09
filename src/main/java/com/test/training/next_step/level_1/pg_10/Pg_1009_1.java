package com.test.training.next_step.level_1.pg_10;

import java.util.ArrayList;
import java.util.List;

/**
 * level 1 - 모의고사
 */
public class Pg_1009_1 {
    public static void main(String[] args) {
        for (int print : solution(new int[]{1, 2, 3, 4, 5})) {
            System.out.print(print + " ");
        }

        System.out.println();
        System.out.println();

        for (int print : solution(new int[]{1, 3, 2, 4, 2})) {
            System.out.print(print + " ");
        }
    }

    public static int[] solution(int[] answers) {
        int[] scores = new int[3];
        int[][] patterns = {
                {1, 2, 3, 4, 5},
                {2, 1, 2, 3, 2, 4, 2, 5},
                {3, 3, 1, 1, 2, 2, 4, 4, 5, 5}
        };

        for (int i = 0; i < patterns.length; i++) {
            for (int j = 0; j < answers.length; j++) {
                // answers 10,000 문제 고려
                if (answers[j] == patterns[i][j % patterns[i].length]) {
                    scores[i]++;
                }
            }
        }

        int maxScore = Math.max(scores[0], Math.max(scores[1], scores[2]));
        List<Integer> resultList = new ArrayList<>();

        for (int i = 0; i < scores.length; i++) {
            if (maxScore == scores[i]) {
                resultList.add(i + 1);
            }
        }

        return resultList.stream()
                .mapToInt(e -> e)
                .toArray();
    }
}