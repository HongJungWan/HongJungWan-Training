package com.test.training.next_step.level_1;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

/**
 * level 1 - 명예의 전당 (1)
 */
public class Pg_0909_1 {
    public static void main(String[] args) {
        System.out.println(solution(3, new int[]{10, 100, 20, 150, 1, 100, 200}));

        System.out.println(solution(4, new int[]{0, 300, 40, 300, 20, 70, 150, 50, 500, 1000}));
    }

    public static int[] solution(int k, int[] scores) {
        int[] answer = new int[scores.length];
        int i = 0;

        List<Integer> list = new ArrayList();

        for (int score : scores) {
            if (list.size() < k) {
                list.add(score);
                answer[i] = Collections.min(list);
                i++;
            } else {
                Collections.sort(list);
                int min = list.get(0);

                addScoreIfExceedsMin(list, score, min);

                answer[i] = Collections.min(list);
                i++;
            }
        }
        return answer;
    }

    private static void addScoreIfExceedsMin(List<Integer> list, int score, int min) {
        if (score > min) {
            list.remove(min);
            list.add(score);
        }
    }
}