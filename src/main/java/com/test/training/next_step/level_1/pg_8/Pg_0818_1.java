package com.test.training.next_step.level_1.pg_8;

/**
 * level 1 - 푸드 파이트 대회
 */
public class Pg_0818_1 {
    private static final int ZERO = 0;
    private static final String STRING_ZERO = "0";
    private static final int ONE = 1;
    private static final int TWO = 2;

    public static void main(String[] args) {
        System.out.println(solution(new int[]{1, 3, 4, 6}));
        System.out.println(solution(new int[]{1, 7, 1, 2}));
    }

    public static String solution(int[] food) {
        String answer = "";
        StringBuilder stringBuilder = new StringBuilder();

        for (int i = ONE; i < food.length; i++) {
            if (isContinueCondition(food[i])) continue;
            toSubCondition(food, i);
            answer = toAddCondition(food, answer, i);
        }
        return toFoodBatchProcess(answer, stringBuilder);
    }

    private static boolean isContinueCondition(int food) {
        return food == ONE;
    }

    private static void toSubCondition(int[] food, int i) {
        if (isContinueCondition(food[i] % TWO)) {
            food[i] -= ONE;
        }
    }

    private static String toAddCondition(int[] food, String answer, int i) {
        for (int j = ZERO; j < food[i] / TWO; j++) {
            answer += i;
        }
        return answer;
    }

    private static String toFoodBatchProcess(String answer, StringBuilder stringBuilder) {
        stringBuilder.append(answer);
        answer += STRING_ZERO;
        answer += stringBuilder.reverse();
        return answer;
    }
}