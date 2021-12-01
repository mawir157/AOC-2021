import AdventHelper

diff_off :: [Integer] -> Int -> [(Integer, Integer)]
diff_off x o = filter (\(l,r) -> l > r) (zip (drop o x) x)

main = do
  putStrLn "Day 1"
  f <- readFile "../input/input01.txt"
  let ss = map(read) $ lines f :: [Integer]

  printSoln 1 (length $ diff_off ss 1)
  printSoln 2 (length $ diff_off ss 3)
