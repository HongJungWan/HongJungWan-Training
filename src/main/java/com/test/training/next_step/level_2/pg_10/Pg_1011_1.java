package com.test.training.next_step.level_2.pg_10;

import java.util.HashSet;
import java.util.Set;

public class Pg_1011_1 {
    public static void main(String[] args) {
        System.out.println(solution("17"));
        System.out.println(solution("011"));
    }

    public static int solution(String numbers) {
        Set<Integer> numberSet = new HashSet<>();
        permutation("", numbers, numberSet);

        int count = 0;
        for (int number : numberSet) {
            if (isPrime(number)) {
                count++;
            }
        }
        return count;
    }

    static void permutation(String prefix, String str, Set<Integer> numberSet) {
        int n = str.length();
        if (!prefix.equals("")) {
            numberSet.add(Integer.parseInt(prefix));
        }

        for (int i = 0; i < n; i++) {
            permutation(prefix + str.charAt(i), str.substring(0, i) + str.substring(i + 1, n), numberSet);
        }
    }

    static boolean isPrime(int n) {
        if (n < 2) return false;

        for (int i = 2; i * i <= n; i++) {
            if (n % i == 0) return false;
        }
        return true;
    }
}