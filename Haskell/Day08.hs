import AdventHelper

import Data.List
import Data.List.Split

parseLine :: String -> ([String], [String])
parseLine s = (digits, codes)
  where parts = splitOn " | " s
        codes = map (sort) $ splitOn " " (parts!!1)
        digits = map (sort) $ splitOn " " (parts!!0)

part1 :: [String] -> Int
part1 [] = 0
part1 (s:ss) = v + (part1 ss)
  where v = if' ((length s) `elem` [2,3,4,7]) 1 0

main = do
  putStrLn "Day 8"
  f <- readFile "../input/input08.txt"
  let ss = lines f
  let t = map parseLine ss

  printSoln 1 $ sum $ map part1 $ map snd t
  printSoln 2 2
