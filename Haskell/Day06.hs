import AdventHelper

import Data.List.Split

parseInput :: String -> [Int]
parseInput s = map read $ splitOn "," s :: [Int]

oneDay :: [(Int, Int)] -> [(Int, Int)]
oneDay [] = []
oneDay ((k,v):xs)
  | k == 0    = [(6,v),(8,v)] ++ oneDay xs
  | otherwise = [(k-1,v)] ++ oneDay xs

life :: Int -> [(Int, Int)] -> [(Int, Int)]
life 0 xs = xs
life n xs = life (n-1) $ reduce $ oneDay xs

reduce :: [(Int, Int)] -> [(Int, Int)]
reduce xs = zip days $ map (helper xs) days
  where days = [0,1..8]
        helper xs n = sum $ map snd $ filter (\(k,_) -> k == n) xs

main = do
  putStrLn "Day 6"
  f <- readFile "../input/input06.txt"
  let ss = lines f
  let fs = freq $ parseInput $ head ss

  printSoln 1 $ sum $ map snd $ life 80 fs
  printSoln 2 $ sum $ map snd $ life 256 fs
