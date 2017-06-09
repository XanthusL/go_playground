package com.liyiheng;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.*;

public class Soduku {
    // 9*9 的格子
    private int[][] grid;
    // 每个格子的候选数字
    private HashMap<Integer, Set<Integer>> candidates;

    @SuppressWarnings("WeakerAccess")
    public Soduku() {
        grid = new int[9][9];
        // Initialize candidate data
        candidates = new HashMap<>();
        for (int i = 0; i < 81; i++) {
            HashSet<Integer> set = new HashSet<>();
            set.add(1);
            set.add(2);
            set.add(3);
            set.add(4);
            set.add(5);
            set.add(6);
            set.add(7);
            set.add(8);
            set.add(9);
            candidates.put(i, set);
        }
    }

    @SuppressWarnings("unused")
    public int[][] getGrid() {
        return grid;
    }

    @SuppressWarnings("unused")
    public void setGrid(int[][] grid) {
        this.grid = grid;
    }

    @SuppressWarnings({"WeakerAccess", "UnusedReturnValue"})
    public int[][] execute() {
        legalCheck(true);
        // 排除
        exclude();
        Solution solution = null;
        Solution nodeCusor = null;
        for (int i = 0; i < 9; i++) {
            for (int j = 0; j < 9; j++) {
                int bigIndex = i * 9 + j;
                int n = grid[i][j];
                if (n == 0) {
                    Set<Integer> set = candidates.get(bigIndex);
                    Solution node = new Solution();
                    node.indexInGrid = bigIndex;
                    node.options = new int[set.size()];
                    int _tmp_index = 0;
                    for (Integer _int : set) {
                        node.options[_tmp_index++] = _int;
                    }
                    if (solution == null) {
                        solution = node;
                    } else {
                        nodeCusor.next = node;
                    }
                    nodeCusor = node;
                }
            }
        }
        boolean ok = walk(solution);
        System.out.println(ok);
        System.out.println(Arrays.deepToString(grid));
        return grid;
    }

    @SuppressWarnings("WeakerAccess")
    public void start() {
        inputGrid();
        execute();
    }

    /**
     * 初始化原始格子
     */
    private void inputGrid() {
        InputStreamReader reader = new InputStreamReader(System.in);
        BufferedReader bufferedReader = new BufferedReader(reader);
        for (int i = 0; i < 9; i++) {
            try {
                String line = bufferedReader.readLine();
                if (line == null) {
                    continue;
                }
                String[] numStrs = line.split(" ");
                int indice = 0;
                for (String s : numStrs) {
                    grid[i][indice++] = Integer.parseInt(s);
                }

            } catch (IOException e) {
                e.printStackTrace();
            }
        }

    }

    private boolean legalCheck(boolean strict) {
        for (int i = 0; i < 9; i++) {
            // 检查列
            boolean ok = check(getColumn(i));
            if (!ok) {
                if (strict) {
                    System.exit(0);
                }
                return false;
            }
            // 检查行
            ok = check(grid[i]);
            if (!ok) {
                if (strict) {
                    System.exit(0);
                }
                return false;
            }
        }
        // 检查3*3的格子
        for (int i = 0; i < 3; i++) {
            for (int j = 0; j < 3; j++) {
                int x = i * 3;
                int y = j * 3;
                int[] numbers = new int[]{
                        grid[x][y],
                        grid[x][y + 1],
                        grid[x][y + 2],
                        grid[x + 1][y],
                        grid[x + 1][y + 1],
                        grid[x + 1][y + 2],
                        grid[x + 2][y],
                        grid[x + 2][y + 1],
                        grid[x + 2][y + 2],
                };
                boolean ok = check(numbers);
                if (!ok) {
                    if (strict) {
                        System.exit(0);
                    }
                    return false;
                }
            }
        }
        return true;
    }

    /**
     * 获取列
     *
     * @param index 0-8
     */
    private int[] getColumn(int index) {
        return new int[]{
                grid[0][index],
                grid[1][index],
                grid[2][index],
                grid[3][index],
                grid[4][index],
                grid[5][index],
                grid[6][index],
                grid[7][index],
                grid[8][index],
        };
    }

    private void exclude() {
        for (int i = 0; i < 9; i++) {
            for (int j = 0; j < 9; j++) {
                int bigIndex = i * 9 + j;
                int v = grid[i][j];
                if (v != 0) {
                    candidates.remove(bigIndex);
                    continue;
                }
                removeCandidate(getColumn(i), bigIndex);
                removeCandidate(grid[i], bigIndex);

                int x = i / 3;
                int y = j / 3;
                x *= 3;
                y *= 3;
                int[] ints = {
                        grid[x][y],
                        grid[x][y + 1],
                        grid[x][y + 2],
                        grid[x + 1][y],
                        grid[x + 1][y + 1],
                        grid[x + 1][y + 2],
                        grid[x + 2][y],
                        grid[x + 2][y + 1],
                        grid[x + 2][y + 2]};
                removeCandidate(ints, bigIndex);


                Set<Integer> integers = candidates.get(bigIndex);
                if (integers.size() == 1) {
                    for (Integer intgr : integers) {
                        grid[i][j] = intgr;
                    }
                }
            }

        }
    }

    private void removeCandidate(int[] nums, int index) {
        Set<Integer> integers = candidates.get(index);
        for (int n : nums) {
            if (n != 0) {
                integers.remove(n);
            }
        }
    }


    /**
     * 检查一组数是否是1-9且不重复
     *
     * @param numbers 长度为9的数组
     * @return 是否可用
     */
    private boolean check(int[] numbers) {
        int len = numbers.length;
        int[] copy = Arrays.copyOf(numbers, len);
        Arrays.sort(copy);

        for (int i = 1; i < len; i++) {
            int v = copy[i];
            if (v == 0) {
                continue;
            }
            if (v == copy[i - 1]) {
                System.out.println("重复");
                return false;
            }
            if (v > 9) {
                System.out.println(">9");
                return false;
            }
        }
        return true;

    }


    private boolean walk(Solution node) {
        if (node == null) {
            return legalCheck(false);
        }
        outer:
        for (int v : node.options) {
            int bigIndex = node.indexInGrid;

            int[] ints0 = grid[bigIndex % 9];
            int[] ints1 = getColumn(bigIndex / 9);
            int[] ints2 = getSubGrid(bigIndex / 9, bigIndex % 9);

            int[] relations = new int[27];
            System.arraycopy(ints0, 0, relations, 0, 9);
            System.arraycopy(ints1, 0, relations, 9, 9);
            System.arraycopy(ints2, 0, relations, 18, 9);

            for (int rv : relations) {
                if (rv != 0 && rv == v) {
                    continue outer;
                }
            }
            grid[node.indexInGrid / 9][node.indexInGrid % 9] = v;
            if (walk(node.next)) {
                return true;
            } else {
                grid[node.indexInGrid / 9][node.indexInGrid % 9] = 0;
            }

        }
        return false;
    }

    private int[] getSubGrid(int x, int y) {
        x /= 3;
        x *= 3;
        y /= 3;
        y *= 3;
        return new int[]{
                grid[x][y],
                grid[x][y + 1],
                grid[x][y + 2],
                grid[x + 1][y],
                grid[x + 1][y + 1],
                grid[x + 1][y + 2],
                grid[x + 2][y],
                grid[x + 2][y + 1],
                grid[x + 2][y + 2]};
    }
}
