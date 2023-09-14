package com.test.training.next_step.level_1.pg_9;

/**
 * level 1 - 카드 뭉치
 */
public class Pg_0914_1 {
    public static void main(String[] args) {
        String[] cards1 = {"i", "drink", "water"};
        String[] cards2 = {"want", "to"};
        String[] goal = {"i", "want", "to", "drink", "water"};

        System.out.println(solution(cards1, cards2, goal));
    }

    public static String solution(String[] cards1, String[] cards2, String[] goal) {
        int cards1Index = 0;
        int cards2Index = 0;

        for (int i = 0; i < goal.length; i++) {
            if (cards1Index < cards1.length && goal[i].equals(cards1[cards1Index])) {
                cards1Index++;
            } else if (cards2Index < cards2.length && goal[i].equals(cards2[cards2Index])) {
                cards2Index++;
            }
        }
        if (cards1Index + cards2Index == goal.length) {
            return "Yes";
        }
        return "No";
    }
}