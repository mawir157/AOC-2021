import AdventHelper

part2 :: [Integer] -> [Bool]
part2 x = map (\(l,r) -> l > r) $ zip (drop 3 x) x

main = do
  putStrLn "Day 1"
  f <- readFile "../input/input01.txt"
  let ss = map(read) $ lines f :: [Integer]

  printSoln 1 (length $ filter (< 0) $ diff ss)
  printSoln 2 (length $ filter (== True) $ part2 ss)
