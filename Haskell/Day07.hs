import AdventHelper

import Data.List.Split

crabEnergy :: [Int] -> Int -> Int
crabEnergy xs p = sum $ map (\x -> abs (x - p)) xs

crabEnergy2 :: [Int] -> Int -> Int
crabEnergy2 xs p = sum $ map (\x -> mv (abs (x - p))) xs
  where mv n = n * (n + 1) `div` 2 

main = do
  putStrLn "Day 7"
  f <- readFile "../input/input07.txt"
  let ss = lines f
  let cs = map read $ splitOn "," $ head ss :: [Int]

  printSoln 1 $ minimum $ map (crabEnergy cs) $ range ((minimum cs), (maximum cs))
  printSoln 2 $ minimum $ map (crabEnergy2 cs) $ range ((minimum cs), (maximum cs))
